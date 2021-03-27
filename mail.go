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

// ListMessages gets all the messages in the current users mailbox.
func (c *MSGraphClient) ListMyMessages(query url.Values) (response MessageCollection, err error) {
	body, err := c.Get("/me/messages", query)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}

// ListMessagesInFolder gets all the messages in a folder for the current user.
func (c *MSGraphClient) ListMyMessagesInFolder(folder string, query url.Values) (response MessageCollection, err error) {
	body, err := c.Get("/me/mailFolders/"+folder+"/messages", query)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}

// GetMyMessageByID gets the message for the specified ID for the current user
func (c *MSGraphClient) GetMyMessageByID(messageID string, query url.Values) (response Message, err error) {
	body, err := c.Get("/me/messages/"+messageID, query)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}
