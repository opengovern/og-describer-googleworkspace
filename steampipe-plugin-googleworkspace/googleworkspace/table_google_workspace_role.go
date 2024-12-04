package googleworkspace

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGoogleWorkspaceRole(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "google_workspace_role",
		Description: "Details about roles in the Google Workspace system, including privileges and system settings.",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "roleId", Type: proto.ColumnType_STRING, Description: "The unique identifier of the role."},
			{Name: "roleName", Type: proto.ColumnType_STRING, Description: "The name of the role."},
			{Name: "roleDescription", Type: proto.ColumnType_STRING, Description: "A description of what the role entails."},
			{Name: "isSystemRole", Type: proto.ColumnType_BOOL, Description: "Indicates if the role is a system role."},
			{Name: "isSuperAdminRole", Type: proto.ColumnType_BOOL, Description: "Indicates if the role is a super admin role."},
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "The type of role (e.g., system, custom)."},
			{Name: "etag", Type: proto.ColumnType_STRING, Description: "The etag identifier for the resource."},

			// JSON column for role privileges
			{Name: "rolePrivileges", Type: proto.ColumnType_JSON, Description: "List of privileges associated with the role."},
		},
	}
}
