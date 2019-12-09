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

// Package msgraph4go provides a Go interface for the Microsoft Graph API.
// See https://developer.microsoft.com/en-us/graph for more details on the Graph API.
package msgraph4go

import (
	"bufio"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/oauth2"
)

const (
	msGraphBase   = "https://graph.microsoft.com/v1.0"
	msAuthBase    = "https://login.microsoftonline.com/common/oauth2/v2.0"
	msAuthURL     = msAuthBase + "/authorize"
	msTokenURL    = msAuthBase + "/token"
	myRedirectURL = "https://login.microsoftonline.com/common/oauth2/nativeclient"
)

// init sets default logging flags
func init() {
	// log with date, time, file name, and line number
	log.SetFlags(log.Lshortfile)
}

func (c *MSGraphClient) GetFile(urlString string, filepath string) (err error) {

	// parse the URL string
	url, err := url.Parse(urlString)
	if err != nil {
		return err
	}

	// execute the request
	resp, err := c.httpClient.Get(url.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// write the body to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return err
}

// Get executes the MS Graph API call, returning the response body.
//
// Query parmeters can be included to specify and control the amount of data returned in a response.
//
// Exact query parameters varies from one API operation to another.
//
// More information can be found at https://docs.microsoft.com/en-us/graph/query-parameters
func (c *MSGraphClient) Get(urlString string, query url.Values) (body []byte, err error) {

	// parse the URL string
	url, err := url.Parse(msGraphBase + urlString)
	if err != nil {
		return body, err
	}

	// add the query parameters to the URL
	url.RawQuery = query.Encode()

	//fmt.Println("DEBUG:", url.String())
	// execute the request
	resp, err := c.httpClient.Get(url.String())
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	// read the body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}

	// check if a MS Graph error occured and return a GraphErrorResponse
	if codeIsError(resp.StatusCode) {
		resError := GraphErrorResponse{}

		err = json.Unmarshal(body, &resError)
		if err != nil {
			return nil, err
		}

		return nil, &resError
	}

	return body, err
}

// MSGraphClient is a client connection to the MS Graph API
type MSGraphClient struct {
	httpClient *http.Client
}

// New creates an initialized MSGraphClient using the token from tokenFileName.
//
// If tokenFileName doesn't exist, then a token is requested and saved in the file.
//
// The current approach assumes the client runs on a host without a
// browser. The user is instructed to vist a URL to login and authorize the
// client. Once the login is successful, the user must copy the response
// URL and provide to the client program.
//
// The token is requested for offline access, which should include a refresh
// token to allow access for a long period of time.
//
// clientID must be an application registered with the Microsoft Identity paltform.
// See https://docs.microsoft.com/en-us/graph/auth-register-app-v2 for more information.
//
// scopes should include the permissions required to call the precding APIs.
// See https://docs.microsoft.com/en-us/graph/permissions-reference for more information.
func New(tokenFileName string, clientID string, scopes []string) *MSGraphClient {
	// default Context that is never canceled, has no values, and has no deadline
	ctx := context.Background()

	scopes = append(scopes, "offline_access")

	// OAuth2 configuration object
	conf := &oauth2.Config{
		ClientID: clientID,
		Scopes:   scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  msAuthURL,
			TokenURL: msTokenURL,
		},
		RedirectURL: myRedirectURL,
	}

	client := &MSGraphClient{}

	// try to get a token from the file
	token, _ := readTokenFromFile(tokenFileName)

	// if token couldn't be retrived, then get a new token
	if token == nil {
		// generate random state to detect Cross-Site Request Forgery
		state := randomBytesBase64(32)

		// get authentication URL for offline access
		authURL := conf.AuthCodeURL(state, oauth2.AccessTypeOffline)

		// instruct the user to vist the authentication URL
		fmt.Println("Vist the following URL in a browser to authenticate this application")
		fmt.Println("After authentication, copy the response URL from the browser")
		fmt.Println(authURL)

		// read the response URL
		fmt.Println("Enter the response URL:")
		responseString, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		responseString = strings.TrimSpace(responseString)

		// parse the response URL
		responseURL, err := url.Parse(responseString)
		if err != nil {
			log.Fatal(err)
		}
		// get and compare state to prevent Cross-Site Request Forgery
		responseState := responseURL.Query().Get("state")
		if responseState != state {
			log.Fatalln("state mismatch, potenial Cross-Site Request Forgery (CSRF)")
		}

		// get authorization code
		code := responseURL.Query().Get("code")

		// exchange authorize code for token
		token, err = conf.Exchange(ctx, code)
		if err != nil {
			log.Fatal(err)
		}

		// save the token to a file
		writeTokenToFile(tokenFileName, token)
	}

	// create HTTP client using the provided token
	client.httpClient = conf.Client(ctx, token)

	return client
}

// randomBytesBase64 returns n bytes encoded in URL friendly base64.
func randomBytesBase64(n int) string {
	// buffer to store n bytes
	b := make([]byte, n)

	// get b random bytes
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	// convert to URL friendly base64
	return base64.URLEncoding.EncodeToString(b)
}

// readTokenFromFile reads the json encoded token from a file.
func readTokenFromFile(filename string) (*oauth2.Token, error) {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read json encoded token
	token := &oauth2.Token{}
	err = json.NewDecoder(file).Decode(token)

	return token, err
}

// writeTokenToFile writes a josn encoded token to a file.
//
// If file already exists, it is replaced.
func writeTokenToFile(fileName string, token *oauth2.Token) {
	// create file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// write access token string
	json.NewEncoder(file).Encode(token)

	return
}

func prettyPrintJson(src []byte) {
	var out bytes.Buffer
	json.Indent(&out, src, "", " ")
	fmt.Println("=== BEGIN ===")
	fmt.Println(out.String())
	fmt.Println("=== END ===")
}

// VarToJsonString converts any varaible (interface{}) to an indented JSON string.
//
// Only public (capitalized) fields will be visible.
func VarToJsonString(v interface{}) string {
	var buffer bytes.Buffer

	// using Encoder rather than MarshalIndent to not escape HTML
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(v)
	if err != nil {
		return "ERROR"
	}

	return string(buffer.Bytes())
}

// WriteContent writes the []byte content to a file
func WriteContentToFile(content []byte, fileName string) (err error) {
	// create file
	var file *os.File
	file, err = os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// write content
	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}
