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
	"os"
	"path/filepath"

	"github.com/bnixon67/msgraph4go"
)

func main() {
	// Get Microsoft Application (client) ID
	// The ID is not in the source code to avoid someone reusing the ID
	clientID, present := os.LookupEnv("MSCLIENTID")
	if !present {
		log.Fatal("Must set MSCLIENTID")
	}

	if len(os.Args) != 2 {
		log.Fatalf("usage: %s path/to/file\n", os.Args[0])
	}

	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	msGraphClient := msgraph4go.New(".token.json", clientID, []string{
		"User.Read", "Files.ReadWrite", "Files.ReadWrite.All", "Sites.ReadWrite.All"})

	_, fileName := filepath.Split(os.Args[1])

	driveItem, err := msGraphClient.UploadNewFile(nil, "me", "root", fileName, data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msgraph4go.VarToJsonString(driveItem))
}
