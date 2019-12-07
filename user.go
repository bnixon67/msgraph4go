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
	"net/url"
)

// GetMyProfile returns the user profile of the current user.
func (c *MSGraphClient) GetMyProfile(query url.Values) (response User, err error) {
	var body []byte

	body, err = c.Get("/me", query)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}

// GetMyPhotoInfo returns the user profile of the current user.
//
// Photos are not supported on personal (consumer) accounts.
func (c *MSGraphClient) GetMyPhotoInfo(query url.Values) (response ProfilePhoto, err error) {
	var body []byte

	body, err = c.Get("/me/photo", query)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}

// GetMyPhoto returns the photo of the current user.
func (c *MSGraphClient) GetMyPhoto(query url.Values) (response []byte, err error) {
	var body []byte

	body, err = c.Get("/me/photo/$value", query)
	if err != nil {
		return nil, err
	}

	return body, err
}
