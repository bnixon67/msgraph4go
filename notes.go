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

// ListNotebooks retrives a list of Notebook objects
func (c *MSGraphClient) ListNotebooks(query url.Values) (response NotebookCollection, err error) {
	body, err := c.Get("/me/onenote/notebooks", query)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}

// ListPages retrives a list of Page objects
func (c *MSGraphClient) ListPages(query url.Values) (response PageCollection, err error) {
	body, err := c.Get("/me/onenote/pages", query)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}

// ListSectionPages retrieve a list of page objects from the specified section.
func (c *MSGraphClient) ListSectionPages(sectionID string, query url.Values) (response PageCollection, err error) {
	body, err := c.Get("/me/onenote/sections/"+sectionID+"/pages", query)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}

// ListSections retrives a list of Section objects
func (c *MSGraphClient) ListSections(query url.Values) (response SectionResponse, err error) {
	body, err := c.Get("/me/onenote/sections", query)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}

func (c *MSGraphClient) GetPage(id string, query url.Values) (response Page, err error) {
	body, err := c.Get("/me/onenote/pages/"+id, query)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}

func (c *MSGraphClient) GetPageContent(id string, query url.Values) (response string, err error) {
	body, err := c.Get("/me/onenote/pages/"+id+"/content", query)
	if err != nil {
		return response, err
	}

	return string(body), err
}
