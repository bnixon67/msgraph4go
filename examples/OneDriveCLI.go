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
	"bufio"
	"encoding/csv"
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

	fileType := ""

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
			if item.File != nil {
				fileType = ""
			}
			if item.Folder != nil {
				fileType = "/"
			}
			if item.Package != nil {
				fileType = "*"
			}
			fmt.Printf("%s %16s %s%s\n",
				t.Local().Format("01/02/2006 03:04 PM"),
				CommaFormat(item.Size),
				item.Name,
				fileType,
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

func help() {
	fmt.Print(`CLI commands
exit, quit - exit or quit the program
list, dir, ls [path] - list contents of directory
pwd - current directory
cd [path] - change to directory, or root if no path is provided
get path/to/file path/to/download - download file
`)
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

	user, err := msGraphClient.GetMyProfile(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User", user.UserPrincipalName)

Loop:
	for {
		fmt.Print(pwd + "> ")

		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())

		// nothing to do
		if line == "" {
			continue
		}

		// use csv reader to split command line and handle embedded quotes
		// https://stackoverflow.com/questions/47489745/splitting-a-string-at-space-except-inside-quotation-marks-go
		reader := csv.NewReader(strings.NewReader(line))
		reader.Comma = ' ' // space
		fields, err := reader.Read()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// dispatch based on first word in input
		switch strings.ToLower(fields[0]) {

		case "help":
			help()

		case "exit", "quit":
			break Loop

		case "list", "dir", "ls":
			var filePath string

			if len(fields) == 1 {
				filePath = pwd
			} else {
				if pwd == "" {
					filePath = fields[1]
				} else {
					filePath = path.Join(pwd, fields[1])
				}
			}

			err = list(msGraphClient, filePath)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case "pwd":
			fmt.Println("pwd =", pwd)

		case "cd":
			if len(fields) == 1 {
				pwd = ""
			} else {
				if pwd == "" {
					pwd = fields[1]
				} else {
					pwd = path.Join(pwd, fields[1])
				}
			}
			if (pwd == ".") || (pwd == "..") {
				pwd = ""
			}

		case "get":
			if len(fields) != 3 {
				fmt.Println("get path/to/file download_location")
				break
			}

			driveItem, err := msGraphClient.GetDriveItemByPath("me", fields[1], nil)
			if err != nil {
				fmt.Println("Cannot GetDriveItemByPath", fields[1])
				fmt.Println(err)
				break
			}

			fmt.Println("Name", driveItem.Name, "Size", driveItem.Size)
			fmt.Println("DownloadURL", driveItem.DownloadURL)

			err = msGraphClient.GetFile(driveItem.DownloadURL, fields[2])
			if err != nil {
				fmt.Println("Cannot GetFile:", err)
			}

		default:
			fmt.Println("Unknown command:", fields[0])
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
