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

// GetMyDrive returns the current user's OneDrive
func (c *MSGraphClient) GetMyDrive(query url.Values) (drive Drive, err error) {
	var body []byte
	body, err = c.Get("/me/drive", query)
	if err != nil {
		return drive, err
	}

	err = json.Unmarshal(body, &drive)

	return drive, err
}

// ListMyDrives retrieve a list of Drives available for the current user
func (c *MSGraphClient) ListMyDrives(query url.Values) (drives DriveResponse, err error) {
	var body []byte
	body, err = c.Get("/me/drives", query)
	if err != nil {
		return drives, err
	}

	err = json.Unmarshal(body, &drives)

	return drives, err
}

// ListRecentFiles is a collection of DriveItems that have been recently used by the signed in user.
//
// This collection includes items that are in the user's drive as well as
// items they have access to from other drives.
func (c *MSGraphClient) ListRecentFiles(query url.Values) (driveItems DriveItemResponse, err error) {
	var body []byte
	body, err = c.Get("/drive/recent", query)
	if err != nil {
		return driveItems, err
	}

	err = json.Unmarshal(body, &driveItems)

	return driveItems, err
}

// ListDriveItemChildren Return a collection of DriveItems in the children relationship of a DriveItem.
//
// DriveItems with a non-null folder or package facet can have one or more child DriveItems.
func (c *MSGraphClient) ListDriveItemChildren(query url.Values) (driveItems DriveItemResponse, err error) {
	var body []byte
	body, err = c.Get("/me/drive/root/children", query)
	if err != nil {
		return driveItems, err
	}

	err = json.Unmarshal(body, &driveItems)

	return driveItems, err
}