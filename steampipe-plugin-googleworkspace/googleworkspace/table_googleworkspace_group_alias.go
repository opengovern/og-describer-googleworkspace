package googleworkspace

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-googleworkspace/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGoogleWorkspaceGroupAlias(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "google_workspace_group_alias",
		Description: "Details about user aliases in Google Workspace, including alias, primary email, and alias type.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListGroupAlias,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetGroupAlias,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the user alias."},
			{Name: "primaryEmail", Type: proto.ColumnType_STRING, Description: "The primary email address associated with the alias."},
			{Name: "alias", Type: proto.ColumnType_STRING, Description: "The alias email address."},
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "The type or kind of the user alias."},
			{Name: "etag", Type: proto.ColumnType_STRING, Description: "The etag identifier for the resource."},
		}),
	}
}
