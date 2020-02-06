/*
Copyright 2020 Bill Nixon

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
