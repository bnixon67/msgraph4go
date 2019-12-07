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
		fmt.Printf("page[%3d]\t%s\n%s", n, page.Title, page.Id)
		fmt.Printf("\t%s\n", page.ParentNotebook.DisplayName)

		// ----- Get Page Content
		content, err := msGraphClient.GetPageContent(page.Id, nil)
		if err != nil {
			log.Fatal(err)
		}

		// ----- Write Page Content
		writeContent(page.Id+".html", content)

	}
	fmt.Println()
	//showResponse(pagesResponse)

	// ----- Get Page
	query = url.Values{}
	query.Set("$expand", "parentNotebook")
	page, err := msGraphClient.GetPage(
		"0-30121ca3825c4aac8b0719aa50aa3634!1-16BE860D241E39E5!37639",
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
}
