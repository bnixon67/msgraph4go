/*
Copyright 2020 Bill Nixon

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
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/bnixon67/msgraph4go"
)

func main() {
	// Get Microsoft Application (client) ID
	// The ID is not in the source code to avoid someone reusing the ID
	clientID, present := os.LookupEnv("MSCLIENTID")
	if !present {
		log.Fatal("Must set MSCLIENTID")
	}

	msGraphClient := msgraph4go.New(
		".token.json",
		clientID,
		[]string{"User.Read", "Files.Read", "Files.Read.All", "Sites.Read.All"},
	)

	query := url.Values{}
	query.Set("$top", "5")

	// nextLink will contain the link to the next set of driveItems, if any
	var nextLink *url.URL

	// loop until no more results
	for {
		// get drive items
		driveItems, err := msGraphClient.ListDriveItemChildrenByID("me", "root", query)
		if err != nil {
			log.Fatal(err)
		}

		// loop thru and display each driveItem
		for _, item := range driveItems.Value {
			fmt.Println("FILE:", item.ID, item.Name, item.LastModifiedDateTime)
		}

		if driveItems.ODataNextLink == "" {
			// if ODataNextLink is empty, then no more items
			break
		} else {
			// parse nextLink for query parameters
			nextLink, err = url.Parse(driveItems.ODataNextLink)
			if err != nil {
				log.Fatal(err)
			}

			// set query parameters for nextLink
			query = nextLink.Query()
			fmt.Println(query)
		}
	}
}
