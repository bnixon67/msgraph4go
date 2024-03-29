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

// A pointer is used for some variable to differentiate between a nil vs empty values
// (see https://stackoverflow.com/questions/33447334/golang-json-marshal-how-to-omit-empty-nested-struct)

package msgraph4go

// Attachment is the base resource for the following derived types of attachments:
//   A file (fileAttachment resource)
//   An item (contact, event or message, represented by an itemAttachment resource)
//   A link to a file (referenceAttachment resource)
type Attachment struct {
	// ContentType is the MIME type.
	ContentType string `json:"contentType,omitempty"`

	// ID is the id of the attachment
	ID string `json:"id,omitempty"`

	// IsInline is true if the attachment is an inline attachment; otherwise, false.
	IsInline bool `json:"isInline,omitempty"`

	// LastModifiedDateTime is when the attachment was last modified.
	LastModifiedDateTime string `json:"lastModifiedDateTime,omitempty"`

	// Name is the attachment's file name.
	Name string `json:"name,omitempty"`

	// Size if the length of the attachment in bytes.
	Size int `json:"size,omitempty"`
}

// BaseItem is an abstract resource that contains a common set of
// properties shared among several other resources types.
// Resources that derive from baseItem include: drive, driveItem, site, sharedDriveItem
type BaseItem struct {
	// The unique identifier of the drive. Read-only.
	ID string `json:"id"`

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
	LastModifiedDateTime string `json:"lastModifiedDateTime,omitempty"`

	// The name of the item. Read-write.
	Name string `json:"name,omitempty"`

	// Parent information, if the item has a parent. Read-write.
	ParentReference *ItemReference `json:"parentReference,omitempty"`

	// URL that displays the resource in the browser. Read-only.
	WebURL string `json:"webUrl,omitempty"`
}

// Calendar is a container for events.
// It can be a calendar for a user, or the default calendar of an Office 365 group.
type Calendar struct {
	OData

	// CanEdit is True if the user can write to the calendar, false otherwise.
	// This property is true for the user who created the calendar.
	// This property is also true for a user who has been shared a calendar and granted write access.
	CanEdit bool `json:"canEdit"`

	// CanShare is True if the user has the permission to share the calendar, false otherwise.
	// Only the user who created the calendar can share it.
	CanShare bool `json:"canShare"`

	// CanViewPrivateItems is True if the user can read calendar
	// items that have been marked private, false otherwise.
	CanViewPrivateItems bool `json:"canViewPrivateItems"`

	// ChangeKey identifies the version of the calendar object.
	// Every time the calendar is changed, changeKey changes as well.
	// This allows Exchange to apply changes to the correct version of the object. Read-only.
	ChangeKey string `json:"changeKey"`

	// Color specifies the color theme to distinguish the calendar from other calendars in a UI.
	// The property values are:
	//   LightBlue=0, LightGreen=1, LightOrange=2, LightGray=3, LightYellow=4, LightTeal=5,
	//   LightPink=6, LightBrown=7, LightRed=8, MaxColor=9, Auto=-1
	Color string `json:"color"`

	// ID is the claendars's unique identifier. Read-only.
	ID string `json:"id"`

	// Name is the calendar name.
	Name string `json:"name"`

	// Ownmer if set, represents the user who created or added the calendar.
	// For a calendar that the user created or added, the owner property is set to the user.
	// For a calendar shared with the user, the owner property is set to the person who shared
	// that calendar with the user.
	Owner EmailAddress `json:"owner"`
}

// CalendarResponse is a collection of Calendar types
type CalendarResponse struct {
	OData
	Value []Calendar `json:"value"`
}

// CalendarGroup is a group of user calendars
type CalendarGroup struct {
	OData

	// Identifies the version of the calendar group.
	// Every time the calendar group is changed, ChangeKey changes as well.
	// This allows Exchange to apply changes to the correct version of the object. Read-only.
	ChangeKey string `json:"changeKey"`

	// The class identifier. Read-only.
	ClassID string `json:"classId"`

	// The group's unique identifier. Read-only.
	ID string `json:"id"`

	// The group name.
	Name string `json:"name"`
}

// CalendarGroupResponse is a collection of CalendarGroup types
type CalendarGroupResponse struct {
	OData
	Value []CalendarGroup `json:"value"`
}

// Contact is an item in Outlook where you can organize and save information about the
// people and organizations you communicate with. Contacts are contained in contact folders.
type Contact struct {
	// The name of the contact's assistant.
	AssistantName string `json:"assistantName"`

	// The contact's birthday. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on
	// Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	Birthday string `json:"birthday"`

	// The contact's business address.
	BusinessAddress struct {
		PhysicalAddress
	} `json:"businessAddress"`

	// The business home page of the contact.
	BusinessHomePage string `json:"businessHomePage"`

	// The contact's business phone numbers.
	BusinessPhones []string `json:"businessPhones"`

	// The categories associated with the contact.
	Categories []string `json:"categories"`

	// Identifies the version of the contact. Every time the contact is changed,
	// ChangeKey changes as well. This allows Exchange to apply changes to the
	// correct version of the object.
	ChangeKey string `json:"changeKey"`

	// The names of the contact's children.
	Children []string `json:"children"`

	// The name of the contact's company.
	CompanyName string `json:"companyName"`

	// The time the contact was created. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	CreatedDateTime string `json:"createdDateTime"`

	// The contact's department.
	Department string `json:"department"`

	// The contact's display name. You can specify the display name in a create or update
	// operation. Note that later updates to other properties may cause an automatically
	// generated value to overwrite the displayName value you have specified. To preserve
	// a pre-existing value, always include it as displayName in an update operation.
	DisplayName string `json:"displayName"`

	// The contact's email addresses.
	EmailAddresses []struct {
		EmailAddress
	} `json:"emailAddresses"`

	// The name the contact is filed under.
	FileAs string `json:"fileAs"`

	// The contact's generation.
	Generation string `json:"generation"`

	// The contact's given name.
	GivenName string `json:"givenName"`

	// The contact's home address.
	HomeAddress struct {
		PhysicalAddress
	} `json:"homeAddress"`

	// The contact's home phone numbers.
	HomePhones []string `json:"homePhones"`

	// The contact's unique identifier. Read-only.
	ID string `json:"id"`

	// The contact's instant messaging (IM) addresses.
	ImAddresses []string `json:"imAddresses"`

	// The contact's initials.
	Initials string `json:"initials"`

	// The contactâs job title.
	JobTitle string `json:"jobTitle"`

	// The time the contact was modified. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight
	// UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
	LastModifiedDateTime string `json:"lastModifiedDateTime"`

	// The name of the contact's manager.
	Manager string `json:"manager"`

	// The contact's middle name.
	MiddleName string `json:"middleName"`

	// The contact's mobile phone number.
	MobilePhone string `json:"mobilePhone"`

	// The contact's nickname.
	NickName string `json:"nickName"`

	// The location of the contact's office.
	OfficeLocation string `json:"officeLocation"`

	// Other addresses for the contact.
	OtherAddress struct {
		PhysicalAddress
	} `json:"otherAddress"`

	// The ID of the contact's parent folder.
	ParentFolderID string `json:"parentFolderId"`

	// The user's notes about the contact.
	PersonalNotes string `json:"personalNotes"`

	// The contact's profession.
	Profession string `json:"profession"`

	// The name of the contact's spouse/partner.
	SpouseName string `json:"spouseName"`

	// The contact's surname.
	Surname string `json:"surname"`

	// The contact's title.
	Title string `json:"title"`

	// The phonetic Japanese company name of the contact.
	YomiCompanyName string `json:"yomiCompanyName"`

	// The phonetic Japanese given name (first name) of the contact.
	YomiGivenName string `json:"yomiGivenName"`

	// The phonetic Japanese surname (last name) of the contact.
	YomiSurname string `json:"yomiSurname"`

	// Optional contact picture. You can get or set a photo for a contact.
	Photo struct {
		ProfilePhoto
	} `json:"photo"`
}

// ContactResponse is a collection of Calendar types
type ContactResponse struct {
	OData
	Value []Contact `json:"value"`
}

// DateTimeTimeZone describes the date, time, and time zone of a point in time.
type DateTimeTimeZone struct {
	// DateTime is a single point of time in a combined date and time representation ({date}T{time}.
	// For example, 2017-08-29T04:00:00.0000000).
	DateTime string `json:"dateTime,omitempty"`

	// TimeZone represents a time zone, for example, "Pacific Standard Time".
	// See https://docs.microsoft.com/en-us/graph/api/resources/datetimetimezone?view=graph-rest-1.0
	// for more possible values.
	TimeZone string `json:"timeZone,omitempty"`
}

// Drive is the top level object representing a user's OneDrive or a document library in SharePoint.
type Drive struct {
	OData

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
	OData
	DownloadURL string `json:"@microsoft.graph.downloadUrl,omitempty"`

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
	ID string `json:"id,omitempty"`

	// TODO: image

	// Identity of the user, device, and application which last modified the item. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime string `json:"lastModifiedDateTime,omitempty"`

	// TODO: location

	// The name of the item (filename and extension). Read-write.
	Name string `json:"name,omitempty"`

	// If present, indicates that this item is a package instead of a folder or file.
	// Packages are treated like files in some contexts and folders in others. Read-only.
	Package *Package `json:"package,omitempty"`

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
	Size int64 `json:"size,omitempty"`

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

// DriveItemVersion represents a specific version of a DriveItem.
type DriveItemVersion struct {
	// ID is the id of the version. Read-only.
	ID string `json:"id"`

	// Identity of the user which last modified the version. Read-only.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// LastModifiedDateTime is Date and time the version was last modified. Read-only.
	LastModifiedDateTime string `json:"lastModifiedDateTime"`

	// Publication indicates the publication status of this particular version. Read-only.
	Publication *PublicationFacet `json:"publication,omitempty"`

	// Size indicates the size of the content stream for this version of the item.
	Size int `json:"size"`

	// TODO: Content
}

// DriveItemVersionResponse is a collection of DriveItemVersion types
type DriveItemVersionResponse struct {
	OData
	Value []DriveItemVersion `json:"value"`
}

// EmailAddress is the name and email address of a contact or message recipient.
type EmailAddress struct {
	// Address is the email address of the person or entity.
	Address string `json:"address,omitempty"`

	// Name is the display name of the person or entity.
	Name string `json:"name,omitempty"`
}

// Extension is an abstract type to support the OData v4 open type openTypeExtension.
type Extension struct {
	ID string `json:"id"`
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

// FollowupFlag allows setting a flag in an item for the user to follow up on later.
type FollowupFlag struct {
	// CompletedDateTime is the date and time that the follow-up was finished.
	CompletedDateTime *DateTimeTimeZone `json:"completedDateTime,omitempty"`

	// DueDateTime is the date and time that the follow-up is to be finished.
	DueDateTime *DateTimeTimeZone `json:"dueDateTime,omitempty"`

	// FlagStatus is the status for follow-up for an item.
	// Possible values are notFlagged, complete, and flagged.
	FlagStatus string `json:"flagStatus,omitempty"`

	// StartDateTime is the date and time that the follow-up is to begin.
	StartDateTime *DateTimeTimeZone `json:"startDateTime,omitempty"`
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
	ID string `json:"id,omitempty"`
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

// InternetMessageHeader is a key-value pair that represents an Internet message header, as defined
// by RFC5322, that provides details of the network path taken by a message from the sender to the
// recipient.
type InternetMessageHeader struct {
	// Name represents the key in a key-value pair.
	Name string `json:"name"`

	// Value is the value in a key-value pair.
	Value string `json:"value"`
}

// ItemBody represents properties of the body of an item, such as a message, event or group post.
type ItemBody struct {
	// Content is the content of the item.
	Content string `json:"content,omitempty"`

	// ContentType is the type of the content. Possible values are text and html.
	ContentType string `json:"contentType,omitempty"`
}

// ItemReference provides information necessary to address a DriveItem via the API.
type ItemReference struct {
	// Unique identifier of the drive instance that contains the item. Read-only.
	DriveId string `json:"driveId,omitempty"`

	// Identifies the type of drive. See drive resource for values.
	DriveType string `json:"driveType,omitempty"`

	// Unique identifier of the item in the drive. Read-only.
	ID string `json:"id,omitempty"`

	// The name of the item being referenced. Read-only.
	Name string `json:"name,omitempty"`

	// Path that can be used to navigate to the item. Read-only.
	Path string `json:"path,omitempty"`

	// A unique identifier for a shared resource that can be accessed via the Shares API.
	ShareId string `json:"shareId,omitempty"`

	// Returns identifiers useful for SharePoint REST compatibility. Read-only.
	SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`
}

// Message is a message in a mailFolder.
type Message struct {
	OData

	// BccRecipients for the message.
	BccRecipients []Recipient `json:"bccRecipients"`

	// Body is the body of the message in HTML or text format.
	Body ItemBody `json:"body,omitempty"`

	// BodyPreview is the first 255 characters of the message body in text format.
	BodyPreview string `json:"bodyPreview"`

	// Categories associated with the message.
	Categories []string `json:"categories"`

	// CcRecipients for the message.
	CcRecipients []Recipient `json:"ccRecipients"`

	// ChangeKey is the version of the message.
	ChangeKey string `json:"changeKey,omitempty"`

	// ConversationID is the ID of the conversation the email belongs to.
	ConversationID string `json:"conversationId,omitempty"`

	// ConversationIndex indicates the position of the message within the conversation.
	ConversationIndex string `json:"conversationIndex,omitempty"`

	// CreatedDateTime when the message was created.
	CreatedDateTime string `json:"createdDateTime,omitempty"`

	// Flag indicates the status, start date, due date, or completion date for the message.
	Flag FollowupFlag `json:"flag,omitempty"`

	// From is the mailbox owner and sender of the message.
	From Recipient `json:"from,omitempty"`

	// HasAttachments indicates whether the message has attachments.
	// This property doesn't include inline attachments, so if a message contains only inline
	// attachments, this property is false.
	// To verify the existence of inline attachments, parse the body property to look for a src
	// attribute, such as <IMG src="cid:image001.jpg@01D26CD8.6C05F070">.
	HasAttachments bool `json:"hasAttachments"`

	// ID is the unique identifier for the message.
	// This value may change if a message is moved or altered.
	ID string `json:"id,omitempty"`

	// Importance the importance of the message: Low, Normal, High.
	Importance string `json:"importance,omitempty"`

	// InferenceClassification is the classification of the message for the user,
	// based on inferred relevance or importance, or on an explicit override.
	// The possible values are: focused or other.
	InferenceClassification string `json:"inferenceClassification,omitempty"`

	// InternetMessageHeaders is a collection of message headers defined by RFC5322.
	// The set includes message headers indicating the network path taken by a message from the
	// sender to the recipient. It can also contain custom message headers that hold app data for
	// the message.
	InternetMessageHeaders []InternetMessageHeader `json:"internetMessageHeaders,omitempty"`

	// InternetMessageID is the message ID in the format specified by RFC2822.
	InternetMessageID string `json:"internetMessageId,omitempty"`

	// IsDeliveryReceiptRequested indicates whether a read receipt is requested for the message.
	IsDeliveryReceiptRequested bool `json:"isDeliveryReceiptRequested"`

	// IsDraft indicates whether the message is a draft.
	// A message is a draft if it hasn't been sent yet.
	IsDraft bool `json:"isDraft"`

	// IsRead indicates whether the message has been read.
	IsRead bool `json:"isRead"`

	// IsReadReceiptRequested indicates whether a read receipt is requested for the message.
	IsReadReceiptRequested bool `json:"isReadReceiptRequested"`

	// LastModifiedDateTime is the date and time the message was last changed.
	LastModifiedDateTime string `json:"lastModifiedDateTime,omitempty"`

	// ParentFolderID is the unique identifier for the message's parent mailFolder.
	ParentFolderID string `json:"parentFolderId,omitempty"`

	// ReceivedDateTime is the date and time the message was received.
	ReceivedDateTime string `json:"receivedDateTime,omitempty"`

	// ReplyTo is the email addresses to use when replying.
	ReplyTo []Recipient `json:"replyTo"`

	// Sender is the account that is actually used to generate the message.
	// In most cases, this value is the same as the from property.
	// You can set this property to a different value when sending a message from a shared mailbox,
	// or sending a message as a delegate.
	// In any case, the value must correspond to the actual mailbox used.
	Sender *Recipient `json:"sender,omitempty"`

	// SentDateTime is the date and time the message was sent.
	SentDateTime string `json:"sentDateTime,omitempty"`

	// Subject is the subject of the message.
	Subject string `json:"subject,omitempty"`

	// ToRecipients for the message.
	ToRecipients []Recipient `json:"toRecipients,omitempty"`

	// UniqueBody is the part of the body of the message that is unique to the current message.
	// uniqueBody is not returned by default but can be retrieved for a given message by use of
	// the ?$select=uniqueBody query. It can be in HTML or text format.
	UniqueBody *ItemBody `json:"uniqueBody,omitempty"`

	// WebLink is the URL to open the message in Outlook Web App.
	//
	// You can append an ispopout argument to the end of the URL to
	// change how the message is displayed. If ispopout is not present or
	// if it is set to 1, then the message is shown in a popout window. If
	// ispopout is set to 0, then the browser will show the message in the
	// Outlook Web App review pane.
	//
	// The message will open in the browser if you are logged in to your
	// mailbox via Outlook Web App. You will be prompted to login if you
	// are not already logged in with the browser.
	//
	// This URL can be accessed from within an iFrame.
	WebLink string `json:"webLink,omitempty"`

	// Attachments are the fileAttachment and itemAttachment attachments for the message.
	Attachments []Attachment `json:"attachments,omitempty"`

	// Extensions are the collection of open extensions defined for the message. Nullable.
	Extensions []Extension `json:"extensions,omitempty"`

	// MultiValueExtendedProperties are the collection of multi-value extended properties
	// defined for the message. Nullable.
	MultiValueExtendedProperties []MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`

	// SingleValueExtendedProperties are the collection of single-value extended properties
	// defined for the message. Nullable.
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}

// MessageCollection is a collection of Notebook types
type MessageCollection struct {
	OData
	Value []Message
}

// MultiValueLegacyExtendedProperty is an extended property that contains a collection of values.
type MultiValueLegacyExtendedProperty struct {
	// ID is the property identifier. Read-only.
	ID string `json:"id"`

	// Value is a collection of property values.
	Value []string `json:"value"`
}

// Notebook represents a OneNote notebook
type Notebook struct {
	// Identity of the user, device, and application which created the item. Read-only.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date and time when the notebook was created. Read-only.
	CreatedDateTime string `json:"createdDateTime"`

	// The unique identifier of the notebook. Read-only.
	ID string `json:"id"`

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
	ODataETag     string `json:"@odata.etag,omitempty"`
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
	ID string `json:"id"`

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

// PageCollection is a collection of Page types
type PageCollection struct {
	OData
	Value []Page `json:"value"`
}

// PageLinks contain links for opening a OneNote page.
type PageLinks struct {
	// Opens the page in the OneNote native client if it's installed.
	OneNoteClientUrl ExternalLink `json:"oneNoteClientUrl"`

	// Opens the page in OneNote on the web.
	OneNoteWebUrl ExternalLink `json:"oneNoteWebUrl"`
}

// Permission provides information about a sharing permission granted for a DriveItem resource.
type Permission struct {
	OData

	// The unique identifier of the permission among all permissions on the item. Read-only.
	ID string `json:"id"`

	// For user type permissions, the details of the users & applications for this permission.
	// Read-only.
	GrantedTo *IdentitySet `json:"grantedTo,omitempty"`

	// Provides a reference to the ancestor of the current permission,
	// if it is inherited from an ancestor. Read-only.
	InheritedFrom *ItemReference `json:"inheritedFrom,omitempty"`

	// Details of any associated sharing invitation for this permission. Read-only.
	Invitation *SharingInvitation `json:"invitation,omitempty"`

	// Provides the link details of the current permission, if it is a link type permissions.
	// Read-only.
	Link *SharingLink `json:"link,omitempty"`

	// The type of permission, e.g. read. Read-only.
	//   read Provides the ability to read the metadata and contents of the item.
	//   write Provides the ability to read and modify the metadata and contents of the item.
	//   sp.owner For SharePoint and OneDrive for Business this represents the owner role.
	//   sp.member For SharePoint and OneDrive for Business this represents the member role.
	Roles []string `json:"roles"`

	// A unique token that can be used to access this shared item via the shares API. Read-only.
	ShareID string `json:"shareId,omitempty"`
}

// PermissionsResponse is a collection of Permission types
type PermissionsResponse struct {
	OData
	Value []Permission `json:"value"`
}

// PhysicalAddress represents the street address of a resource such as a contact or event.
type PhysicalAddress struct {
	// The city.
	City string `json:"city"`

	// The country or region. It's a free-format string value, for example, "United States".
	CountryOrRegion string `json:"countryOrRegion"`

	// The postal code.
	PostalCode string `json:"postalCode"`

	// The state.
	State string `json:"state"`

	// The street.
	Street string `json:"street"`
}

// ProfilePhoto is a profile photo of a user, group or an Outlook contact accessed from Exchange Online.
// It's binary data not encoded in base-64.
//
// The supported sizes of HD photos on Exchange Online are as follows:
// '48x48', '64x64', '96x96', '120x120', '240x240', '360x360','432x432', '504x504', and '648x648'.
type ProfilePhoto struct {
	// Id of the photo. Read-only.
	ID string `json:"id"`

	// Height of the photo. Read-only.
	Height int `json:"height"`

	// Width of the photo. Read-only.
	Width int `json:"width"`

	ODataContext          string `json:"@odata.context"`
	ODataMediaContentType string `json:"@odata.mediaContentType"`
	ODataMediaEtag        string `json:"@odata.mediaEtag"`
}

// PublicationFacet provides details on the published status of a
// driveItemVersion or driveItem resource.
type PublicationFacet struct {
	// Level is the state of publication for this document.
	// Either published or checkout. Read-only.
	Level string `json:"level"`

	// VersionID is The unique identifier for the version that is
	// visible to the current caller. Read-only.
	VersionID string `json:"versionId"`
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

// Recipient represents information about a user in the sending or receiving end of an event,
// message or group post.
type Recipient struct {
	EmailAddress *EmailAddress `json:"emailaddress,omitempty"`
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
	ID string `json:"id,omitempty"`

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
	ID string `json:"id"`

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

	// The notebook that contains the section. Read-only.
	ParentNotebook Notebook `json:"parentNotebook"`
}

// SectionResponse is a collection of Section types
type SectionResponse struct {
	OData
	Value []Section `json:"value"`
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

// SharingInvitation groups invitation-related data items into a single structure.
type SharingInvitation struct {
	// The email address provided for the recipient of the sharing invitation. Read-only.
	Email string `json:"email"`

	// Provides information about who sent the invitation that created this permission,
	// if that information is available. Read-only.
	InvitedBy *IdentitySet `json:"invitedBy,omitempty"`

	// If true the recipient of the invitation needs to sign in in order to access the shared item.
	// Read-only.
	SignInRequired bool `json:"signInRequired"`
}

// SharingLink groups link-related data items into a single structure.
//
// If a Permission resource has a non-null sharingLink facet, the permission represents a sharing link
// (as opposed to permissions granted to a person or group).
type SharingLink struct {
	// The app the link is associated with.
	Application *Identity `json:"application,omitempty"`

	// The type of the link created.
	Type string `json:"type,omitempty"`

	// The scope of the link represented by this permission. Value anonymous indicates the link
	// is usable by anyone, organization indicates the link is only usable for users signed into
	// the same tenant.
	Scope string `json:"scope,omitempty"`

	// For embed links, this property contains the HTML code for an <iframe> element that will
	// embed the item in a webpage.
	WebHTML string `json:"webHtml,omitempty"`

	// A URL that opens the item in the browser on the OneDrive website.
	WebURL string `json:"webUrl,omitempty"`
}

// SingleValueLegacyExtendedProperty is an extended property that contains a single value.
type SingleValueLegacyExtendedProperty struct {
	// ID is the property ID used to identify the property. Read-only.
	ID string `json:"id"`

	// Value is a property value.
	Value string `json:"value"`
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
	ID string `json:"id,omitempty"`

	// The userâs job title.
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
