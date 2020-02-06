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

	if len(os.Args) != 2 {
		log.Fatalf("usage: %s path/to/file\n", os.Args[0])
	}

	// get drive item
	driveItem, err := msGraphClient.GetDriveItemByID("me", os.Args[1], query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name", driveItem.Name)
	fmt.Println("Size", driveItem.Size)
	fmt.Println("DownloadURL", driveItem.DownloadURL)

	err = msGraphClient.GetFile(driveItem.DownloadURL, driveItem.Name+".download")
	if err != nil {
		log.Fatal(err)
	}
}
