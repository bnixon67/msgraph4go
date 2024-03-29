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

	if len(os.Args) != 2 {
		log.Fatalf("usage: %s path/to/file\n", os.Args[0])
	}

	// get drive item
	driveItem, err := msGraphClient.GetDriveItemByPath("me", os.Args[1], query)
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
