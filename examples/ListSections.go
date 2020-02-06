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
		[]string{"User.Read", "Notes.Read", "Notes.Read.All"},
	)

	query := url.Values{}
	query.Set("$count", "true")

	sectionsResponse, err := msGraphClient.ListSections(query)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("total sections = %d\n", sectionsResponse.ODataCount)
	fmt.Printf("response sections = %d\n\n", len(sectionsResponse.Value))

	for n, section := range sectionsResponse.Value {
		fmt.Printf("%d %s %s-%s\n", n, section.ID, section.ParentNotebook.DisplayName,
			section.DisplayName)

		query := url.Values{}
		query.Set("$orderby", "order")

		pages, err := msGraphClient.ListSectionPages(section.ID, query)
		if err != nil {
			log.Fatal(err)
		}

		for n, page := range pages.Value {
			fmt.Printf("%2d %s %s\n",
				n,
				page.LastModifiedDateTime,
				page.Title,
			)
		}

		fmt.Println()
	}
}
