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

// A pointer is used for some variable to differentiate between a nil vs empty values
// (see https://stackoverflow.com/questions/33447334/golang-json-marshal-how-to-omit-empty-nested-struct)

// BaseItem is an abstract resource that contains a common set of
// properties shared among several other resources types.
// Resources that derive from baseItem include: drive, driveItem, site, sharedDriveItem
type BaseItem struct {
	// The unique identifier of the drive. Read-only.
	Id string `json:"id"`

	// Identity of the user, device, or application which created the item. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// Date and time of item creation. Read-only.
	CreatedDateTime string `json:"createdDateTime,omitempty"`

	// Provides a user-visible description of the item. Optional.
	Description string `json:"description,omitempty"`

	// ETag for the item. Read-only.
	ETag string `json:"eTag,omitempty"`

	// Identity of the user, device, and application which last modified the item. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime string `json:"LastModifiedDateTime,omitempty"`

	// The name of the item. Read-write.
	Name string `json:"name,omitempty"`

	// Parent information, if the item has a parent. Read-write.
	ParentReference *ItemReference `json:"parentReference,omitempty"`

	// URL that displays the resource in the browser. Read-only.
	WebURL string `json:"webUrl,omitempty"`
}

// Drive is the top level object representing a user's OneDrive or a document library in SharePoint.
type Drive struct {
	BaseItem

	// Describes the type of drive represented by this resource.
	// OneDrive personal drives will return personal.
	// OneDrive for Business will return business.
	// SharePoint document libraries will return documentLibrary.
	// Read-only.
	DriveType string `json:"driveType"`

	// Optional. The user account that owns the drive. Read-only.
	Owner *IdentitySet `json:"owner,omitempty"`

	// Optional. Information about the drive's storage space quota. Read-only.
	Quota *Quota `json:"quota,omitempty"`

	// TODO: sharepointids
	// TODO: system
}

// DriveResponse is a collection of Drive types
type DriveResponse struct {
	OData
	Value []Drive `json:"value"`
}

// DriveItem represents a file, folder, or other item stored in a drive.
type DriveItem struct {
	// TODO: audio
	// TODO: content

	// Identity of the user, device, and application which created the item. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// Date and time of item creation. Read-only.
	CreatedDateTime string `json:"createdDateTime,omitempty"`

	// An eTag for the content of the item.
	// This eTag is not changed if only the metadata is changed.
	CTag string `json:"cTag,omitempty"`

	// TODO: deleted

	// Provide a user-visible description of the item.
	Description string `json:"description,omitempty"`

	// eTag for the entire item (metadata + content). Read-only.
	ETag string `json:"eTag,omitempty"`

	// File metadata, if the item is a file. Read-only.
	File *File `json:"file,omitempty"`

	// File system information on client. Read-write.
	FileSystemInfo FileSystemInfo `json:"fileSystemInfo,omitempty"`

	// Folder metadata, if the item is a folder. Read-only.
	Folder *Folder `json:"folder,omitempty"`

	// The unique identifier of the item within the Drive. Read-only.
	Id string `json:"id,omitempty"`

	// TODO: image

	// Identity of the user, device, and application which last modified the item. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime string `json:"lastModifiedDateTime,omitempty"`

	// TODO: location

	// The name of the item (filename and extension). Read-write.
	Name string `json:"name,omitempty"`

	// TODO: package

	// Parent information, if the item has a parent. Read-write.
	ParentReference *ItemReference `json:"parentReference,omitempty"`

	// TODO: photo
	// TODO: publication

	// Remote item data, if the item is shared from a drive other than the one being accessed.
	// Read-only.
	RemoteItem RemoteItem `json:"remoteItem,omitempty"`

	// TODO: root
	// TODO: searchResult
	// TODO: shared

	// Returns identifiers useful for SharePoint REST compatibility. Read-only.
	SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`

	// Size of the remote item. Read-only.
	Size *int64 `json:"size,omitempty"`

	// TODO: specialFolder
	// TODO: video
	// TODO: webDavUrl

	// URL that displays the resource in the browser. Read-only.
	WebURL string `json:"webUrl,omitempty"`
}

// DriveItemResponse is a collection of DriveItem types
type DriveItemResponse struct {
	OData
	Value []DriveItem `json:"value"`
}

// ExternalLink is a url that opens a OneNote page or notebook.
type ExternalLink struct {
	// The url of the link.
	Href string `json:"href"`
}

// File groups file-related data items into a single structure.
type File struct {
	// TODO: hashes

	// The MIME type for the file. This is determined by logic on
	// the server and might not be the value provided when the file
	// was uploaded. Read-only.
	MimeType string `json:"mimeType,omitempty"`
}

// FileSystemInfo contains properties that are reported by the device's
// local file system for the local version of an item.
type FileSystemInfo struct {
	// The UTC date and time the file was created on a client.
	CreatedDateTime string `json:"createdDateTime,omitempty"`

	// The UTC date and time the file was last accessed. Available for the recent file list only.
	LastAccessedDateTime string `json:"lastAccessedDateTime,omitempty"`

	// The UTC date and time the file was last modified on a client.
	LastModifiedDateTime string `json:"lastModifiedDateTime,omitempty"`
}

// Folder groups folder-related data on an item into a single structure.
// DriveItems with a non-null folder facet are containers for other DriveItems.
type Folder struct {
	// Number of children contained immediately within this container.
	ChildCount int32 `json:"childCount,omitempty"`

	// A collection of properties defining the recommended view for the folder.
	View *FolderView `json:"view,omitempty"`
}

// FolderView provides or sets recommendations on the user-experience of a folder.
type FolderView struct {
	// The method by which the folder should be sorted.
	SortBy string `json:"sortBy,omitempty"`

	// If true, indicates that items should be sorted in descending order.
	// Otherwise, items should be sorted ascending.
	SortOrder string `json:"sortOrder,omitempty"`

	// The type of view that should be used to represent the folder.
	ViewType string `json:"viewType,omitempty"`
}

// The Identity resource represents an identity of an actor.
// For example, an actor can be a user, device, or application.
type Identity struct {
	// The identity's display name. Note that this may not always
	// be available or up to date. For example, if a user changes their
	// display name, the API may show the new value in a future response,
	// but the items associated with the user won't show up as having
	// changed when using delta.
	DisplayName string `json:"displayName,omitempty"`

	EMail string `json:"email,omitempty"`

	// Unique identifier for the identity.
	Id string `json:"id,omitempty"`
}

// IdentitySet is a keyed collection of identity resources. It is used
// to represent a set of identities associated with various events for an
// item, such as created by or last modified by.
type IdentitySet struct {
	// Optional. The application associated with this action.
	Application *Identity `json:"application,omitempty"`

	// Optional. The device associated with this action.
	Device *Identity `json:"device,omitempty"`

	// Optional. The user associated with this action.
	User *Identity `json:"user,omitempty"`
}

// ItemReference provides information necessary to address a DriveItem via the API.
type ItemReference struct {
	// Unique identifier of the drive instance that contains the item. Read-only.
	DriveId string `json:"driveId,omitempty"`

	// Identifies the type of drive. See drive resource for values.
	DriveType string `json:"driveType,omitempty"`

	// Unique identifier of the item in the drive. Read-only.
	Id string `json:"id,omitempty"`

	// The name of the item being referenced. Read-only.
	Name string `json:"name,omitempty"`

	// Path that can be used to navigate to the item. Read-only.
	Path string `json:"path,omitempty"`

	// A unique identifier for a shared resource that can be accessed via the Shares API.
	ShareId string `json:"shareId,omitempty"`

	// Returns identifiers useful for SharePoint REST compatibility. Read-only.
	SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`
}

// Notebook represents a OneNote notebook
type Notebook struct {
	// Identity of the user, device, and application which created the item. Read-only.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date and time when the notebook was created. Read-only.
	CreatedDateTime string `json:"createdDateTime"`

	// The unique identifier of the notebook. Read-only.
	Id string `json:"id"`

	// Indicates whether this is the user's default notebook. Read-only.
	IsDefault bool `json:"isDefault"`

	// Indicates whether the notebook is shared. Read-only.
	IsShared bool `json:"isShared"`

	// Identity of the user, device, and application which created the item. Read-only.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The date and time when the notebook was last modified. Read-only.
	LastModifiedDateTime string `json:"lastModifiedDateTime"`

	// Links for opening the notebook.
	Links NotebookLinks `json:"links"`

	// The name of the notebook.
	DisplayName string `json:"displayName"`

	// The URL for the sectionGroups navigation property, which
	// returns all the section groups in the notebook. Read-only.
	SectionGroupsUrl string `json:"sectionsGroupsUrl"`

	// The URL for the sections navigation property, which returns
	// all the sections in the notebook. Read-only.
	SectionsUrl string `json:"sectionsUrl"`

	// The endpoint where you can get details about the notebook. Read-only.
	Self string `json:"self"`

	// Possible values are: Owner, Contributor, Reader, None. Read-only.
	//   Owner represents owner-level access to the notebook.
	//   Contributor represents read/write access to the notebook.
	//   Reader represents read-only access to the notebook.
	UserRole string `json:"userRole"`
}

// NotebookCollection is a collection of Notebook types
type NotebookCollection struct {
	OData
	Value []Notebook
}

// NotebookLinks are links for opening a OneNote notebook.
type NotebookLinks struct {
	// Opens the notebook in the OneNote native client if it's installed.
	OneNoteClientUrl ExternalLink `json:"oneNoteClientUrl"`

	// Opens the notebook in OneNote on the web.
	OneNoteWebUrl ExternalLink `json:"oneNoteWebUrl"`
}

type OData struct {
	ODataContext  string `json:"@odata.context,omitempty"`
	ODataCount    int    `json:"@odata.count,omitempty"`
	ODataNextLink string `json:"@odata.nextLink,omitempty"`
}

// Package indicates that a DriveItem is the top level item in a "package"
// or a collection of items that should be treated as a collection instead
// of individual items.
type Package struct {
	// A string indicating the type of package.
	// While oneNote is the only currently defined value, you should expect
	// other package types to be returned and handle them accordingly.
	Type string `json:"type,omitempty"`
}

// Page represents a page in a OneNote notebook
type Page struct {
	// The page's HTML content.
	Content string `json:"content"`

	// The URL for the page's HTML content. Read-only.
	ContentUrl string `json:"contentUrl"`

	// The unique identifier of the application that created the page. Read-only.
	CreatedByAppId string `json:"createdByAppId"`

	// The date and time when the page was created.
	CreatedDateTime string `json:"createdDateTime"`

	// The unique identifier of the page. Read-only.
	Id string `json:"id"`

	// The date and time when the page was last modified.
	LastModifiedDateTime string `json:"lastModifiedDateTime"`

	// The indentation level of the page. Read-only.
	Level int32 `json:"level"`

	// Links for opening the page.
	Links PageLinks `json:"links"`

	// The order of the page within its parent section. Read-only.
	Order int32 `json:"order"`

	// The endpoint where you can get details about the page. Read-only.
	Self string `json:"self"`

	// The title of the page.
	Title string `json:"title"`

	// The notebook that contains the page. Read-only.
	ParentNotebook Notebook `json:"parentNotebook"`

	// The section that contains the page. Read-only.
	ParentSection Section `json:"parentSection"`
}

// PageLinks contain links for opening a OneNote page.
type PageLinks struct {
	// Opens the page in the OneNote native client if it's installed.
	OneNoteClientUrl ExternalLink `json:"oneNoteClientUrl"`

	// Opens the page in OneNote on the web.
	OneNoteWebUrl ExternalLink `json:"oneNoteWebUrl"`
}

type PageResponse struct {
	OData
	Value        []Page `json:"value"`
	ODataContext string `json:"@odata.context"`
}

type ProfilePhoto struct {
	// Id of the photo. Read-only.
	Id string `json:"id"`

	// Height of the photo. Read-only.
	Height int `json:"height"`

	// Width of the photo. Read-only.
	Width int `json:"width"`

	ODataContext          string `json:"@odata.context"`
	ODataMediaContentType string `json:"@odata.mediaContentType"`
	ODataMediaEtag        string `json:"@odata.mediaEtag"`
}

// Quota provides details about space constrains on a Drive resource.
type Quota struct {
	// Total allowed storage space, in bytes. Read-only.
	Total int64 `json:"total,omitempty"`

	// Total space used, in bytes. Read-only.
	Used int64 `json:"used,omitempty"`

	// Total space remaining before reaching the quota limit, in bytes. Read-only.
	Remaining int64 `json:"remaining,omitempty"`

	// Total space consumed by files in the recycle bin, in bytes. Read-only.
	Deleted int64 `json:"deleted,omitempty"`

	// Enumeration value that indicates the state of the storage space. Read-only.
	State string `json:"state,omitempty"`
}

// RemoteItem indicates that a driveItem references an item that exists in another drive.
// This resource provides the unique IDs of the source drive and target item.
type RemoteItem struct {
	// Identity of the user, device, and application which created the item. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// Date and time of item creation. Read-only.
	CreatedDateTime string `json:"createdDateTime,omitempty"`

	// Indicates that the remote item is a file. Read-only.
	File *File `json:"file,omitempty"`

	// Information about the remote item from the local file system. Read-only.
	FileSystemInfo FileSystemInfo `json:"fileSystemInfo,omitempty"`

	// Indicates that the remote item is a folder. Read-only.
	Folder *Folder `json:"folder,omitempty"`

	// Unique identifier for the remote item in its drive. Read-only.
	Id string `json:"id,omitempty"`

	// Identity of the user, device, and application which last modified the item. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime string `json:"lastModifiedDateTime,omitempty"`

	// Optional. Filename of the remote item. Read-only.
	Name string `json:"name,omitempty"`

	// If present, indicates that this item is a package instead of a folder or file.
	// Packages are treated like files in some contexts and folders in others. Read-only.
	Package *Package `json:"package,omitempty"`

	// Properties of the parent of the remote item. Read-only.
	ParentReference *ItemReference `json:"parentReference,omitempty"`

	// TODO: shared

	// Provides interop between items in OneDrive for Business and
	// SharePoint with the full set of item identifiers. Read-only.
	SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`

	// Size of the remote item. Read-only.
	Size int64 `json:"size,omitempty"`

	// TODO: specialFolder

	// DAV compatible URL for the item.
	WebDavURL string `json:"webDavUrl,omitempty"`

	// URL that displays the resource in the browser. Read-only.
	WebURL string `json:"webUrl,omitempty"`
}

// Section represents a section in a OneNote notebook. Sections can contain pages.
type Section struct {
	// Identity of the user, device, and application which created the item. Read-only.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date and time when the section was created. Read-only.
	CreatedDateTime string `json:"createdDateTime"`

	// The unique identifier of the section. Read-only.
	Id string `json:"id"`

	// Indicates whether this is the user's default section. Read-only.
	IsDefault bool `json:"isDefault"`

	// Identity of the user, device, and application which created the item. Read-only.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The date and time when the section was last modified.
	LastModifiedDateTime string `json:"lastModifiedDateTime"`

	// Links for opening the section.
	Links NotebookLinks `json:"links"`

	// The name of the section.
	DisplayName string `json:"displayName"`

	// The pages endpoint where you can get details for all the pages in the section. Read-only.
	PagesUrl string `json:"pagesUrl"`

	// The endpoint where you can get details about the section. Read-only.
	Self string `json:"self"`
}

// SharePointIds groups the various identifiers for an item stored in
// a SharePoint site or OneDrive for Business into a single structure.
type SharepointIds struct {
	// The unique identifier (guid) for the item's list in SharePoint.
	ListId string `json:"listId,omitempty"`

	// An integer identifier for the item within the containing list.
	ListItemId string `json:"listItemId,omitempty"`

	// The unique identifier (guid) for the item within OneDrive for Business or a SharePoint site.
	ListItemUniqueId string `json:"listItemUniqueId,omitempty"`

	// The unique identifier (guid) for the item's site collection (SPSite).
	SiteId string `json:"siteId,omitempty"`

	// The SharePoint URL for the site that contains the item.
	SiteUrl string `json:"siteUrl,omitempty"`

	// The unique identifier (guid) for the item's site (SPWeb).
	WebId string `json:"webId,omitempty"`
}

// User represents an Azure AD user account
// Not all of the properties have been included from
// https://docs.microsoft.com/en-us/graph/api/resources/user?view=graph-rest-1.0
type User struct {
	OData

	// The telephone numbers for the user.
	// Although this is a string collection, only one number can be set for this property.
	BusinessPhones []string `json:"businessPhones,omitempty"`

	// The name displayed in the address book for the user.
	DisplayName string `json:"displayName,omitempty"`

	// The given name (first name) of the user.
	GivenName string `json:"givenName,omitempty"`

	// The unique identifier for the user. Read-only.
	Id string `json:"id,omitempty"`

	// The user’s job title.
	JobTitle string `json:"jobTitle,omitempty"`

	// The SMTP address for the user, for example, "jeff@contoso.onmicrosoft.com". Read-Only.
	Mail string `json:"mail,omitempty"`

	// The primary cellular telephone number for the user.
	MobilePhone string `json:"mobilePhone,omitempty"`

	// The office location in the user's place of business.
	OfficeLocation string `json:"officeLocation,omitempty"`

	// The preferred language for the user. Should follow ISO 639-1 Code; for example "en-US".
	PreferredLanguage interface{} `json:"preferredLanguage,omitempty"`

	// The user's surname (family name or last name).
	Surname string `json:"surname,omitempty"`

	// The user principal name (UPN) of the user.
	UserPrincipalName string `json:"userPrincipalName,omitempty"`
}
