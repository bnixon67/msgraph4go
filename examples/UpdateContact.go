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
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bnixon67/msgraph4go"
)

func ParseCommandLine() (tokenFile string, scopes []string, user string) {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s [options] request\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "options:")
		flag.PrintDefaults()
	}

	flag.StringVar(&tokenFile, "token", ".token.json", "path to `file` to use for token")
	flag.StringVar(&user, "user", "me", "user to use for update")

	var scopeString string
	flag.StringVar(&scopeString,
		"scopes", "Contacts.ReadWrite", "comma-seperated `scopes` to use for request")

	flag.Parse()

	scopes = strings.Split(scopeString, ",")

	return tokenFile, scopes, user
}

func main() {
	// Get Microsoft Application (client) ID
	// The ID is not in the source code to avoid someone reusing the ID
	clientID, present := os.LookupEnv("MSCLIENTID")
	if !present {
		log.Fatal("Must set MSCLIENTID")
	}

	// parse command line to get path to the token file and scopes to use in request
	tokenFile, scopes, user := ParseCommandLine()

	// check for no remaining args
	if len(flag.Args()) != 1 {
		flag.Usage()
		return
	}

	msGraphClient := msgraph4go.New(tokenFile, clientID, scopes)

	jsonStr := strings.NewReader("{ \"displayName\": \"Tlast, Tfirst\" }")

	contact, err := msGraphClient.UpdateContact(nil, user, flag.Arg(0), jsonStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msgraph4go.VarToJsonString(contact))
}
