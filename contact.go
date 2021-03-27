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

package msgraph4go

import (
	"encoding/json"
	"io"
	"net/url"
)

// ListMyContacts gets all contacts in a user's mailbox.
//
// user must be "me", userPrincipalName, or id
func (c *MSGraphClient) ListContacts(query url.Values, user string) (response ContactResponse, err error) {
	var body []byte

	var url string

	if user == "me" {
		url = "/me/contacts"
	} else {
		url = "/users/" + user + "/contacts"
	}

	body, err = c.Get(url, query)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}

// UpdateContact updates the properties of a contact object.
//
// user must be "me", userPrincipalName, or id
func (c *MSGraphClient) UpdateContact(query url.Values, user string, contactID string, data io.Reader) (contact Contact, err error) {
	var body []byte
	var url string

	if user == "me" {
		url = "/me/contacts/" + contactID
	} else {
		url = "/users/" + user + "/contacts/" + contactID
	}

	body, err = c.Patch(url, query, data)
	if err != nil {
		return contact, err
	}

	err = json.Unmarshal(body, &contact)

	return contact, err
}
