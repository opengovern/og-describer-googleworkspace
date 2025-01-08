package googleworkspace

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-googleworkspace/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGoogleWorkspaceUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_user",
		Description: "Users in the Google Workspace domain.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListUser,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetUser,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Id"),
				Description: "The unique ID of the user."},
			{
				Name:        "primaryEmail",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrimaryEmail"),
				Description: "The primary email address of the user."},
			{
				Name:        "password",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Password"),
				Description: "Password associated with the user."},
			{
				Name:        "hashFunction",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HashFunction"),
				Description: "Hash function used to store the password."},
			{
				Name:        "isAdmin",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.IsAdmin"),
				Description: "Indicates if the user is an administrator."},
			{
				Name:        "isDelegatedAdmin",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.IsDelegatedAdmin"),
				Description: "Indicates if the user is a delegated administrator."},
			{
				Name:        "agreedToTerms",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.AgreedToTerms"),
				Description: "Indicates if the user has agreed to the terms of service."},
			{
				Name:        "suspended",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Suspended"),
				Description: "Indicates if the user's account is suspended."},
			{
				Name:        "changePasswordAtNextLogin",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ChangePasswordAtNextLogin"),
				Description: "Indicates if the user needs to change their password at the next login."},
			{
				Name:        "ipWhitelisted",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.IpWhitelisted"),
				Description: "Indicates if the user is IP whitelisted."},

			// Name object (UserName)
			{
				Name:        "name",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Name"),
				Description: "The name of the user, including given name and family name."},

			// Other fields
			{
				Name:      "kind",
				Transform: transform.FromField("Description.Kind"),
				Type:      proto.ColumnType_STRING, Description: "Resource kind identifier, typically 'admin#directory#user'."},
			{
				Name:      "etag",
				Transform: transform.FromField("Description.Etag"),
				Type:      proto.ColumnType_STRING, Description: "ETag of the resource."},
			{
				Name:      "emails",
				Transform: transform.FromField("Description.Emails"),
				Type:      proto.ColumnType_JSON, Description: "Email addresses associated with the user."},
			{
				Name:      "externalIds",
				Transform: transform.FromField("Description.ExternalIds"),
				Type:      proto.ColumnType_JSON, Description: "External IDs associated with the user."},
			{
				Name:      "relations",
				Transform: transform.FromField("Description.Relations"),
				Type:      proto.ColumnType_JSON, Description: "Relationships associated with the user."},
			{
				Name:      "aliases",
				Transform: transform.FromField("Description.Aliases"),
				Type:      proto.ColumnType_JSON, Description: "Email aliases for the user."},
			{
				Name:      "isMailboxSetup",
				Transform: transform.FromField("Description.IsMailboxSetup"),
				Type:      proto.ColumnType_BOOL, Description: "Indicates if the user's mailbox is set up."},
			{
				Name:      "customerId",
				Transform: transform.FromField("Description.CustomerId"),
				Type:      proto.ColumnType_STRING, Description: "The customer ID associated with the user."},
			{
				Name:      "addresses",
				Transform: transform.FromField("Description.Addresses"),
				Type:      proto.ColumnType_JSON, Description: "Addresses associated with the user."},
			{
				Name:      "organizations",
				Transform: transform.FromField("Description.Organizations"),
				Type:      proto.ColumnType_JSON, Description: "Organizations associated with the user."},
			{
				Name:      "lastLoginTime",
				Transform: transform.FromField("Description.LastLoginTime"),
				Type:      proto.ColumnType_TIMESTAMP, Description: "Timestamp of the user's last login."},
			{
				Name:      "phones",
				Transform: transform.FromField("Description.Phones"),
				Type:      proto.ColumnType_JSON, Description: "Phone numbers associated with the user."},
			{
				Name:      "suspensionReason",
				Transform: transform.FromField("Description.SuspensionReason"),
				Type:      proto.ColumnType_STRING, Description: "Reason for suspending the user's account."},
			{
				Name:      "thumbnailPhotoUrl",
				Transform: transform.FromField("Description.ThumbnailPhotoUrl"),
				Type:      proto.ColumnType_STRING, Description: "URL to the user's thumbnail photo."},
			{
				Name:      "languages",
				Transform: transform.FromField("Description.Languages"),
				Type:      proto.ColumnType_JSON, Description: "Languages spoken by the user."},
			{
				Name:      "posixAccounts",
				Transform: transform.FromField("Description.PosixAccounts"),
				Type:      proto.ColumnType_JSON, Description: "POSIX account information associated with the user."},
			{
				Name:      "creationTime",
				Transform: transform.FromField("Description.CreationTime"),
				Type:      proto.ColumnType_TIMESTAMP, Description: "Timestamp when the user was created."},
			{
				Name:      "nonEditableAliases",
				Transform: transform.FromField("Description.NonEditableAliases"),
				Type:      proto.ColumnType_JSON, Description: "Non-editable aliases for the user."},
			{
				Name:      "sshPublicKeys",
				Transform: transform.FromField("Description.SshPublicKeys"),
				Type:      proto.ColumnType_JSON, Description: "SSH public keys associated with the user."},
			{
				Name:      "notes",
				Transform: transform.FromField("Description.Notes"),
				Type:      proto.ColumnType_JSON, Description: "Notes associated with the user."},
			{
				Name:      "websites",
				Transform: transform.FromField("Description.Websites"),
				Type:      proto.ColumnType_JSON, Description: "Websites associated with the user."},
			{
				Name:      "locations",
				Transform: transform.FromField("Description.Locations"),
				Type:      proto.ColumnType_JSON, Description: "Locations associated with the user."},
			{
				Name:      "includeInGlobalAddressList",
				Transform: transform.FromField("Description.IncludeInGlobalAddressList"),
				Type:      proto.ColumnType_BOOL, Description: "Indicates if the user is included in the global address list."},
			{
				Name:      "keywords",
				Transform: transform.FromField("Description.Keywords"),
				Type:      proto.ColumnType_JSON, Description: "Keywords associated with the user."},
			{
				Name:      "deletionTime",
				Transform: transform.FromField("Description.DeletionTime"),
				Type:      proto.ColumnType_TIMESTAMP, Description: "Timestamp when the user's account was deleted."},
			{
				Name:      "gender",
				Transform: transform.FromField("Description.Gender"),
				Type:      proto.ColumnType_JSON, Description: "Gender information associated with the user."},
			{
				Name:      "thumbnailPhotoEtag",
				Transform: transform.FromField("Description.ThumbnailPhotoEtag"),
				Type:      proto.ColumnType_STRING, Description: "ETag for the user's thumbnail photo."},
			{
				Name:      "ims",
				Transform: transform.FromField("Description.Ims"),
				Type:      proto.ColumnType_JSON, Description: "Instant messaging service accounts associated with the user."},
			{
				Name:      "customSchemas",
				Transform: transform.FromField("Description.CustomSchemas"),
				Type:      proto.ColumnType_JSON, Description: "Custom schema information associated with the user."},
			{
				Name:      "isEnrolledIn2Sv",
				Transform: transform.FromField("Description.IsEnrolledIn2Sv"),
				Type:      proto.ColumnType_BOOL, Description: "Indicates if the user is enrolled in two-step verification."},
			{
				Name:      "isEnforcedIn2Sv",
				Transform: transform.FromField("Description.IsEnforcedIn2Sv"),
				Type:      proto.ColumnType_BOOL, Description: "Indicates if two-step verification is enforced for the user."},
			{
				Name: "archived", Type: proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Archived"),
				Description: "Indicates if the user account is archived."},
			{
				Name: "orgUnitPath", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OrgUnitPath"),
				Description: "The organizational unit path for the user."},
			{
				Name:        "recoveryEmail",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryEmail"),
				Description: "The user's recovery email address."},
			{
				Name:        "recoveryPhone",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RecoveryPhone"),
				Description: "The user's recovery phone number."},
		}),
	}
}
