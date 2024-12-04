package googleworkspace

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGoogleWorkspaceGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "google_workspace_group",
		Description: "Groups in the Google Workspace domain.",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique ID of the group."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "The email address of the group."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the group."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "A description of the group."},
			{Name: "adminCreated", Type: proto.ColumnType_BOOL, Description: "Indicates if the group was created by an admin."},
			{Name: "directMembersCount", Type: proto.ColumnType_STRING, Description: "The number of direct members in the group."},
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "Resource kind identifier, typically 'admin#directory#group'."},
			{Name: "etag", Type: proto.ColumnType_STRING, Description: "ETag of the resource."},

			// Aliases columns
			{Name: "aliases", Type: proto.ColumnType_JSON, Description: "List of email aliases associated with the group."},
			{Name: "nonEditableAliases", Type: proto.ColumnType_JSON, Description: "List of email aliases that cannot be edited."},
		},
	}
}
