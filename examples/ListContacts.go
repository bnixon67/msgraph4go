/*
Copyright 2021 Bill Nixon

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
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
		"scopes", "Contacts.Read", "comma-seperated `scopes` to use for request")

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
			fmt.Printf("Contact %d %s\n", n, contact.ID)
			fmt.Printf("GivenName = %s Surname = %s\n",
				contact.GivenName, contact.Surname)
			fmt.Printf("DisplayName = %s\n", contact.DisplayName)
			fmt.Println()
			n++
		}

		// check if additional results
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
