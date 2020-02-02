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
