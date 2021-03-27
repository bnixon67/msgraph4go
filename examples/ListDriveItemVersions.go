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
		[]string{"User.Read", "Files.Read", "Files.Read.All", "Sites.Read.All"},
	)

	query := url.Values{}
	query.Set("$top", "2")

	// nextLink will contain the link to the next set of driveItems, if any
	var nextLink *url.URL

	// loop until no more results
	for {
		// get drive items
		driveItemVersionResponse, err :=
			msGraphClient.ListDriveItemVersions(
				"me", "014SHJLEPLWAIM2NIZYZHLCOF6GHUOQ4RM", query)
		if err != nil {
			log.Fatal(err)
		}

		// loop thru and display each driveItem
		for _, item := range driveItemVersionResponse.Value {
			fmt.Println(item.ID, item.LastModifiedDateTime)
		}

		if driveItemVersionResponse.ODataNextLink == "" {
			// if ODataNextLink is empty, then no more items
			break
		} else {
			// parse nextLink for query parameters
			nextLink, err = url.Parse(driveItemVersionResponse.ODataNextLink)
			if err != nil {
				log.Fatal(err)
			}

			// set query parameters for nextLink
			query = nextLink.Query()
			fmt.Println(query)
		}
	}
}
