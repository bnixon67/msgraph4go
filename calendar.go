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
