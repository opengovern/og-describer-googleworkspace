package googleworkspace

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGoogleWorkspaceUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "google_workspace_user",
		Description: "Users in the Google Workspace domain.",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique ID of the user."},
			{Name: "primaryEmail", Type: proto.ColumnType_STRING, Description: "The primary email address of the user."},
			{Name: "password", Type: proto.ColumnType_STRING, Description: "Password associated with the user."},
			{Name: "hashFunction", Type: proto.ColumnType_STRING, Description: "Hash function used to store the password."},
			{Name: "isAdmin", Type: proto.ColumnType_BOOL, Description: "Indicates if the user is an administrator."},
			{Name: "isDelegatedAdmin", Type: proto.ColumnType_BOOL, Description: "Indicates if the user is a delegated administrator."},
			{Name: "agreedToTerms", Type: proto.ColumnType_BOOL, Description: "Indicates if the user has agreed to the terms of service."},
			{Name: "suspended", Type: proto.ColumnType_BOOL, Description: "Indicates if the user's account is suspended."},
			{Name: "changePasswordAtNextLogin", Type: proto.ColumnType_BOOL, Description: "Indicates if the user needs to change their password at the next login."},
			{Name: "ipWhitelisted", Type: proto.ColumnType_BOOL, Description: "Indicates if the user is IP whitelisted."},

			// Name object (UserName)
			{Name: "name", Type: proto.ColumnType_JSON, Description: "The name of the user, including given name and family name."},

			// Other fields
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "Resource kind identifier, typically 'admin#directory#user'."},
			{Name: "etag", Type: proto.ColumnType_STRING, Description: "ETag of the resource."},
			{Name: "emails", Type: proto.ColumnType_JSON, Description: "Email addresses associated with the user."},
			{Name: "externalIds", Type: proto.ColumnType_JSON, Description: "External IDs associated with the user."},
			{Name: "relations", Type: proto.ColumnType_JSON, Description: "Relationships associated with the user."},
			{Name: "aliases", Type: proto.ColumnType_JSON, Description: "Email aliases for the user."},
			{Name: "isMailboxSetup", Type: proto.ColumnType_BOOL, Description: "Indicates if the user's mailbox is set up."},
			{Name: "customerId", Type: proto.ColumnType_STRING, Description: "The customer ID associated with the user."},
			{Name: "addresses", Type: proto.ColumnType_JSON, Description: "Addresses associated with the user."},
			{Name: "organizations", Type: proto.ColumnType_JSON, Description: "Organizations associated with the user."},
			{Name: "lastLoginTime", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of the user's last login."},
			{Name: "phones", Type: proto.ColumnType_JSON, Description: "Phone numbers associated with the user."},
			{Name: "suspensionReason", Type: proto.ColumnType_STRING, Description: "Reason for suspending the user's account."},
			{Name: "thumbnailPhotoUrl", Type: proto.ColumnType_STRING, Description: "URL to the user's thumbnail photo."},
			{Name: "languages", Type: proto.ColumnType_JSON, Description: "Languages spoken by the user."},
			{Name: "posixAccounts", Type: proto.ColumnType_JSON, Description: "POSIX account information associated with the user."},
			{Name: "creationTime", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the user was created."},
			{Name: "nonEditableAliases", Type: proto.ColumnType_JSON, Description: "Non-editable aliases for the user."},
			{Name: "sshPublicKeys", Type: proto.ColumnType_JSON, Description: "SSH public keys associated with the user."},
			{Name: "notes", Type: proto.ColumnType_JSON, Description: "Notes associated with the user."},
			{Name: "websites", Type: proto.ColumnType_JSON, Description: "Websites associated with the user."},
			{Name: "locations", Type: proto.ColumnType_JSON, Description: "Locations associated with the user."},
			{Name: "includeInGlobalAddressList", Type: proto.ColumnType_BOOL, Description: "Indicates if the user is included in the global address list."},
			{Name: "keywords", Type: proto.ColumnType_JSON, Description: "Keywords associated with the user."},
			{Name: "deletionTime", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the user's account was deleted."},
			{Name: "gender", Type: proto.ColumnType_JSON, Description: "Gender information associated with the user."},
			{Name: "thumbnailPhotoEtag", Type: proto.ColumnType_STRING, Description: "ETag for the user's thumbnail photo."},
			{Name: "ims", Type: proto.ColumnType_JSON, Description: "Instant messaging service accounts associated with the user."},
			{Name: "customSchemas", Type: proto.ColumnType_JSON, Description: "Custom schema information associated with the user."},
			{Name: "isEnrolledIn2Sv", Type: proto.ColumnType_BOOL, Description: "Indicates if the user is enrolled in two-step verification."},
			{Name: "isEnforcedIn2Sv", Type: proto.ColumnType_BOOL, Description: "Indicates if two-step verification is enforced for the user."},
			{Name: "archived", Type: proto.ColumnType_BOOL, Description: "Indicates if the user account is archived."},
			{Name: "orgUnitPath", Type: proto.ColumnType_STRING, Description: "The organizational unit path for the user."},
			{Name: "recoveryEmail", Type: proto.ColumnType_STRING, Description: "The user's recovery email address."},
			{Name: "recoveryPhone", Type: proto.ColumnType_STRING, Description: "The user's recovery phone number."},
		},
	}
}
