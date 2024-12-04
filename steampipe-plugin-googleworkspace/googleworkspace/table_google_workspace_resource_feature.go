package googleworkspace

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-googleworkspace/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGoogleWorkspaceResourceFeature(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "google_workspace_resource_feature",
		Description: "Details about a resource, including its name, kind, and etags.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListResourceFeature,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetResourceFeature,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the resource."},
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "The kind of resource."},
			{Name: "etags", Type: proto.ColumnType_STRING, Description: "The etag identifier for the resource."},
		}),
	}
}
