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
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bnixon67/msgraph4go"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("OneDrive CLI")

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())

		words := strings.Fields(line)

		fmt.Println(words)

		if strings.ToLower(line) == "exit" {
			break
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Get Microsoft Application (client) ID
	// The ID is not in the source code to avoid someone reusing the ID
	clientID, present := os.LookupEnv("MSCLIENTID")
	if !present {
		log.Fatal("Must set MSCLIENTID")
	}

	msGraphClient := msgraph4go.New(".token.json", clientID, []string{"User.Read"})

	fmt.Println("foo")
	resp, err := msGraphClient.Get("/me/drive/root/children", nil)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, resp, "", " ")
	fmt.Println("=== BEGIN ===")
	fmt.Println(out.String())
	fmt.Println("=== END ===")
}
