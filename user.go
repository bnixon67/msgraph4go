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
