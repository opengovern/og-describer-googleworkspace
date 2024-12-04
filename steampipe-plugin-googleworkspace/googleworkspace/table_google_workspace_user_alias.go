package googleworkspace

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGoogleWorkspaceUserAlias(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "google_workspace_user_alias",
		Description: "User aliases in the Google Workspace domain.",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique ID of the user alias."},
			{Name: "primaryEmail", Type: proto.ColumnType_STRING, Description: "The primary email address of the user who owns this alias."},
			{Name: "alias", Type: proto.ColumnType_STRING, Description: "The alias email address."},
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "Resource kind identifier, typically 'admin#directory#userAlias'."},
			{Name: "etag", Type: proto.ColumnType_STRING, Description: "ETag of the resource."},
		}),
	}
}
