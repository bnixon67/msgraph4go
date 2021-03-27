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
