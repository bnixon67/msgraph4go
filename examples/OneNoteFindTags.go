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
	"io"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/bnixon67/msgraph4go"
	"golang.org/x/net/html"
)

// find_tag find the given tag in the source using HTML Tokenizer
//   returns a string array of the tag values
func find_tag(source io.Reader, tag string) (vals []string) {

	z := html.NewTokenizer(source)

	saveText := false
	for {
		tokenType := z.Next()

		switch tokenType {

		case html.ErrorToken:
			return

		case html.TextToken:
			if saveText {
				vals = append(vals, string(z.Text()))
			}

		case html.StartTagToken:
			_, hasAttr := z.TagName()
			if hasAttr {
				key, val, _ := z.TagAttr()
				if string(key) == "data-tag" {
					vals := strings.Split(string(val), ",")
					for _, v := range vals {
						if v == tag {
							saveText = true
						}
					}
				}
			}
		case html.EndTagToken:
			saveText = false
		}
	}
}

func main() {
	// Get Microsoft Application (client) ID
	// The ID is not in the source code to avoid someone reusing the ID
	clientID, present := os.LookupEnv("MSCLIENTID")
	if !present {
		log.Fatal("Must set MSCLIENTID")
	}

	msGraphClient := msgraph4go.New(".token.json", clientID, []string{"User.Read"})

	var query url.Values
	var nextLink *url.URL

	// WaitGroup to fetch multiple pages
	var wg sync.WaitGroup

	// list of page may be returned by multiple queries
	// @odata.nextLink has the link to next set of pages
	for {
		// ----- List Pages

		// first query (no nextLink)
		if nextLink == nil {
			query = url.Values{}

			// total number of pages
			query.Set("$count", "true")

			// sort by page title
			//			query.Set("$orderby", "parentSection/displayName,title")
			query.Set("$orderby", "title")

			// exand parentNotebook to get displayName
			query.Set("$expand", "parentNotebook,parentSection")

			// filter on just one Notebook
			//query.Set("$filter",
			//	"parentNotebook/displayName eq 'UMB Notes'")
		} else {
			// set query value based on nextLink
			query = nextLink.Query()
		}

		// get pages
		pagesResponse, err := msGraphClient.ListPages(query)
		if err != nil {
			log.Fatal(err)
		}

		// get nextLink (if any)
		nextLink, _ = url.Parse(pagesResponse.ODataNextLink)

		// loop thru each page
		for _, page := range pagesResponse.Value {

			// increase WaitGroup counter
			wg.Add(1)

			// run goroutine to get page content and find tags
			go func(page msgraph4go.Page) {
				// ensure we decrease WaitGroup
				defer wg.Done()

				/*
					fmt.Printf("Checking  %s/%s/%s\n",
						page.ParentNotebook.DisplayName,
						page.ParentSection.DisplayName,
						page.Title)
				*/

				// ----- Get Page Content
				content, err := msGraphClient.GetPageContent(page.ID, nil)
				if err != nil {
					fmt.Printf("ERROR  %s/%s/%s\n",
						page.ParentNotebook.DisplayName,
						page.ParentSection.DisplayName,
						page.Title)
					log.Fatal(err)
				}

				// find to-do tags in the page content
				v := find_tag(strings.NewReader(content), "to-do")

				// at least one to-do tag found
				if len(v) > 0 {
					fmt.Printf("----- %3d %s/%s/%s\n",
						len(v),
						page.ParentNotebook.DisplayName,
						page.ParentSection.DisplayName,
						page.Title)
					for n, v := range v {
						fmt.Printf("%3d\t%v\n", n, v)
					}
					fmt.Println()
				}
			}(page)

			// ----- Write Page Content
			//writeContent(page.Id+".html", content)
		}

		// Wait for all page requests complete
		wg.Wait()

		// nextLink is empty, so exit loop
		if pagesResponse.ODataNextLink == "" {
			break
		}

	}
}
