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
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/bnixon67/msgraph4go"
)

// writeContent writes out the content
func writeContent(fileName string, content string) {
	// create file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// write content
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

	return
}

// showResponse pretty prints the response using MarshalIndent
func showResponse(v interface{}) {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		fmt.Println("error: err")
	}
	fmt.Printf("==========\n%s\n===========\n", b)
}

func main() {
	// Get Microsoft Application (client) ID
	// The ID is not in the source code to avoid someone reusing the ID
	clientID, present := os.LookupEnv("MSCLIENTID")
	if !present {
		log.Fatal("Must set MSCLIENTID")
	}

	msGraphClient := msgraph4go.New(".token.json", clientID, []string{"User.Read"})

	// ----- List Notebooks
	query := url.Values{}
	query.Set("$count", "true")
	query.Set("$top", "2")
	//query.Set("$filter", "startswith(displayName, 'U')")
	notebooksResponse, err := msGraphClient.ListNotebooks(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("total notebooks = %d\n", notebooksResponse.ODataCount)
	fmt.Printf("response notebooks = %d\n\n", len(notebooksResponse.Value))
	//showResponse(notebooksResponse)

	for n, notebook := range notebooksResponse.Value {
		fmt.Printf("notebook[%d]\t%s\n", n, notebook.DisplayName)
	}
	fmt.Println()

	// ----- List Pages
	query = url.Values{}
	query.Set("$count", "true")
	//query.Set("$top", "5")
	query.Set("$expand", "parentNotebook")
	pagesResponse, err := msGraphClient.ListPages(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("count of pages = %d\n", pagesResponse.ODataCount)
	fmt.Printf("pages in response = %d\n\n", len(pagesResponse.Value))

	for n, page := range pagesResponse.Value {
		fmt.Printf("page[%3d]\t%s\n%s", n, page.Title, page.ID)
		fmt.Printf("\t%s\n", page.ParentNotebook.DisplayName)

		// ----- Get Page Content
		content, err := msGraphClient.GetPageContent(page.ID, nil)
		if err != nil {
			log.Fatal(err)
		}

		// ----- Write Page Content
		writeContent(page.ID+".html", content)

	}
	fmt.Println()
	//showResponse(pagesResponse)

	// ----- Get Page
	/*
		query = url.Values{}
		query.Set("$expand", "parentNotebook")
		page, err := msGraphClient.GetPage(
			page.Id,
			query,
		)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("id=%v\n", page.Id)
		fmt.Printf("title=%v\n", page.Title)
		fmt.Printf("link=%v\n", page.Links.OneNoteWebUrl.Href)
		fmt.Printf("notebook=%v\n", page.ParentNotebook.DisplayName)
		//showResponse(page)
	*/
}
