/*
Copyright 2019 Bill Nixon

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published
by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

This program is distributed in the hope that it will be useful, but
WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/bnixon67/msgraph4go"
)

func ParseCommandLine() (tokenFile string, scopes []string, user string) {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s [options] request\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "options:")
		flag.PrintDefaults()
	}

	flag.StringVar(&tokenFile, "token", ".token.json", "path to `file` to use for token")
	flag.StringVar(&user, "user", "me", "user to get contacts for")

	var scopeString string
	flag.StringVar(&scopeString,
		"scopes", "Contacts.ReadWrite", "comma-seperated `scopes` to use for request")

	flag.Parse()

	scopes = strings.Split(scopeString, ",")

	return tokenFile, scopes, user
}

func main() {
	// Get Microsoft Application (client) ID
	// The ID is not in the source code to avoid someone reusing the ID
	clientID, present := os.LookupEnv("MSCLIENTID")
	if !present {
		log.Fatal("Must set MSCLIENTID")
	}

	// parse command line to get path to the token file and scopes to use in request
	tokenFile, scopes, user := ParseCommandLine()

	// check for no remaining args
	if len(flag.Args()) != 0 {
		flag.Usage()
		return
	}

	msGraphClient := msgraph4go.New(tokenFile, clientID, scopes)

	// nextLink will contain the link to the next set of items, if any
	var nextLink *url.URL

	query := url.Values{}

	n := 1

	// loop until no more results
	for {
		contacts, err := msGraphClient.ListContacts(query, user)
		if err != nil {
			log.Fatal(err)
		}

		for _, contact := range contacts.Value {
			fmt.Printf("Updating contact %d\n", n)
			//fmt.Printf("\tID: %s\n", contact.ID)
			fmt.Printf("\tSurname: %s GivenName: %s\n",
				contact.Surname, contact.GivenName)
			fmt.Printf("\tOld DisplayName: %s\n", contact.DisplayName)
			n++

			updateStr := fmt.Sprintf(`{ "displayName": "%s, %s" }`,
				contact.Surname, contact.GivenName)
			reader := strings.NewReader(updateStr)

			updatedContact, err := msGraphClient.UpdateContact(
				nil, user, contact.ID, reader)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\tNew DisplayName: %s\n", updatedContact.DisplayName)
			fmt.Println()
		}

		// check if additional items
		if contacts.ODataNextLink == "" {
			// if ODataNextLink is empty, then no more items
			break
		} else {
			// parse nextLink for query parameters
			nextLink, err = url.Parse(contacts.ODataNextLink)
			if err != nil {
				log.Fatal(err)
			}

			// set query parameters for nextLink
			query = nextLink.Query()
		}
	}
}
