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
