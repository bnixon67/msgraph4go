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
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/bnixon67/msgraph4go"
)

// CommaFormat formats an integer into a string, with commas
//
// See https://stackoverflow.com/questions/13020308/how-to-fmt-printf-an-integer-with-thousands-comma
// https://stackoverflow.com/users/1705598/icza
func CommaFormat(n int64) string {
	in := strconv.FormatInt(n, 10)
	out := make([]byte, len(in)+(len(in)-2+int(in[0]/'0'))/3)
	if in[0] == '-' {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}

// list prints out the children of a DriveItem
func list(c *msgraph4go.MSGraphClient, path string) (err error) {
	query := url.Values{}
	//query.Set("$top", "5")
	query.Set("$orderby", "Name")

	// nextLink will contain the link to the next set of driveItems, if any
	var nextLink *url.URL

	// loop until no more results
	for {
		// get drive items
		driveItems, err := c.ListDriveItemChildrenByPath("me", path, query)
		if err != nil {
			return err
		}

		// loop thru and display each driveItem
		for _, item := range driveItems.Value {
			t, _ := time.Parse(time.RFC3339, item.LastModifiedDateTime)
			fmt.Printf("%s %17s %s\t%s\n",
				t.Local().Format("01/02/2006 03:04 PM"),
				CommaFormat(item.Size),
				item.Id,
				item.Name,
			)
		}

		if driveItems.ODataNextLink == "" {
			// if ODataNextLink is empty, then no more items
			break
		} else {
			// parse nextLink for query parameters
			nextLink, err = url.Parse(driveItems.ODataNextLink)
			if err != nil {
				return err
			}

			// set query parameters for nextLink
			query = nextLink.Query()
		}
	}
	return err
}

func main() {
	fmt.Println("OneDrive CLI")

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

	scanner := bufio.NewScanner(os.Stdin)

	pwd := ""

Loop:
	for {
		fmt.Print(pwd + "> ")

		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())

		words := strings.Fields(line)

		var err error
		//var itemID string

		switch strings.ToLower(words[0]) {
		case "exit", "quit":
			break Loop
		case "list", "dir", "ls":
			var path string

			if len(words) == 1 {
				path = pwd
			} else {
				path = words[1]
			}

			_ = list(msGraphClient, path)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "pwd":
			fmt.Println("pwd =", pwd)
		case "cd":
			if len(words) == 1 {
				pwd = ""
			} else {
				if pwd == "" {
					pwd = words[1]
				} else {
					pwd = path.Join(pwd, words[1])
				}
			}
		default:
			fmt.Println("Unknown command:", words[0])
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
