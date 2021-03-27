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

// ListMyCalendars gets all the user's calendars.
func (c *MSGraphClient) ListMyCalendars(query url.Values) (response CalendarResponse, err error) {
	var body []byte

	body, err = c.Get("/me/calendars", query)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}

// GetMyDefaultCalendar gets the current users default calendar.
func (c *MSGraphClient) GetMyDefaultCalendar(query url.Values) (response Calendar, err error) {
	var body []byte

	body, err = c.Get("/me/calendar", query)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}

// ListMyCalendarGroups gets the curent user's calendar groups.
func (c *MSGraphClient) ListMyCalendarGroups(query url.Values) (response CalendarGroupResponse, err error) {
	var body []byte

	body, err = c.Get("/me/calendarGroups", query)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}
