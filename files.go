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
	"fmt"
	"io"
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

// ListDriveItemChildrenByID return a collection of DriveItems in the children relationship
// of a DriveItem.
//
// driveID should be a valid driveID or could be "me"
//
// itemID should be a valid itemID or could be "root"
//
// DriveItems with a non-null folder or package facet can have one or more child DriveItems.
func (c *MSGraphClient) ListDriveItemChildrenByID(driveID string, itemID string, query url.Values) (driveItems DriveItemResponse, err error) {
	var body []byte
	body, err = c.Get("/drives/"+driveID+"/items/"+itemID+"/children", query)
	if err != nil {
		return driveItems, err
	}

	err = json.Unmarshal(body, &driveItems)

	return driveItems, err
}

// ListDriveItemPermissionsByID returns a collection of Permission objects.
//
// driveID should be a valid driveID or could be "me"
//
// itemID should be a valid itemID or could be "root"
func (c *MSGraphClient) ListDriveItemPermissionsByID(driveID string, itemID string, query url.Values) (permissions PermissionsResponse, err error) {
	var body []byte
	body, err = c.Get("/drives/"+driveID+"/items/"+itemID+"/permissions", query)
	if err != nil {
		return permissions, err
	}

	err = json.Unmarshal(body, &permissions)

	return permissions, err
}

// GetDriveItemPermission returns a collection of Permission objects.
//
// driveID should be a valid driveID or could be "me"
//
// itemID should be a valid itemID or could be "root"
//
// permID should be a valid permission ID
func (c *MSGraphClient) GetDriveItemPermission(driveID string, itemID string, permID string, query url.Values) (permission Permission, err error) {
	var body []byte
	fmt.Println("/drives/" + driveID + "/items/" + itemID + "/permissions/" + permID)
	body, err = c.Get("/drives/"+driveID+"/items/"+itemID+"/permissions/"+permID, query)
	if err != nil {
		return permission, err
	}

	err = json.Unmarshal(body, &permission)

	return permission, err
}

// ListDriveItemVersion lists versions of a DriveItem.
//
// driveID should be a valid driveID or could be "me"
//
// OneDrive and SharePoint can be configured to retain the history for
// files. Depending on the service and configuration, a new version can be
// created for each edit, each time the file is saved, manually, or never.
//
// Previous versions of a document may be retained for a finite period of
// time depending on admin settings which may be unique per user or location.
func (c *MSGraphClient) ListDriveItemVersions(driveID string, itemID string, query url.Values) (driveItemVersionResponse DriveItemVersionResponse, err error) {
	var body []byte
	body, err = c.Get("/drives/"+driveID+"/items/"+itemID+"/versions", query)
	if err != nil {
		return driveItemVersionResponse, err
	}

	err = json.Unmarshal(body, &driveItemVersionResponse)

	return driveItemVersionResponse, err
}

// ListDriveItemChildrenByPath return a collection of DriveItems in the children relationship
// of a DriveItem.
//
// driveID should be a valid driveID or could be "me"
//
// itemID should be a valid itemID or could be "root"
//
// DriveItems with a non-null folder or package facet can have one or more child DriveItems.
func (c *MSGraphClient) ListDriveItemChildrenByPath(driveID string, path string, query url.Values) (driveItems DriveItemResponse, err error) {
	var body []byte
	var url string

	if path == "" || path == "/" {
		url = "/drives/" + driveID + "/items/root/children"
	} else {
		url = "/drives/" + driveID + "/root:/" + path + ":/children"
	}

	body, err = c.Get(url, query)
	if err != nil {
		return driveItems, err
	}

	err = json.Unmarshal(body, &driveItems)

	return driveItems, err
}

// GetDriveItemByID return a DriveItem
//
// driveID should be a valid driveID or could be "me"
//
// itemID should be a valid itemID or could be "root"
func (c *MSGraphClient) GetDriveItemByID(driveID string, itemID string, query url.Values) (driveItem DriveItem, err error) {
	var body []byte
	body, err = c.Get("/drives/"+driveID+"/items/"+itemID, query)
	if err != nil {
		return driveItem, err
	}

	err = json.Unmarshal(body, &driveItem)

	return driveItem, err
}

// GetDriveItemByPath return a DriveItem
//
// driveID should be a valid driveID or could be "me"
//
// itemID should be a valid itemID or could be "root"
func (c *MSGraphClient) GetDriveItemByPath(driveID string, path string, query url.Values) (driveItem DriveItem, err error) {
	var body []byte
	var url string

	if path == "" || path == "/" {
		url = "/drives/" + driveID + "/items/root"
	} else {
		url = "/drives/" + driveID + "/root:/" + path + ":"
	}

	body, err = c.Get(url, query)
	if err != nil {
		return driveItem, err
	}

	err = json.Unmarshal(body, &driveItem)

	return driveItem, err
}

// UploadNewFile is a simple upload API allows you to provide the
// contents of a new file or update the contents of an existing file in a
// single API call. This method only supports files up to 4MB in size.
func (c *MSGraphClient) UploadNewFile(query url.Values, driveID string, parentID string, fileName string, data io.Reader) (driveItem DriveItem, err error) {
	var body []byte
	var url string

	url = "/drives/" + driveID + "/items/" + parentID + ":/" + fileName + ":/content"

	body, err = c.Put(url, query, data)
	if err != nil {
		return driveItem, err
	}

	err = json.Unmarshal(body, &driveItem)

	return driveItem, err
}
