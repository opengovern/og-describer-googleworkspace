package googleworkspace

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-googleworkspace/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGoogleWorkspacePrivilege(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_privilege",
		Description: "Information about privileges, including service details and child privileges.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListPrivilege,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetPrivilege,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "etag", Type: proto.ColumnType_STRING, Description: "The etag identifier for the privilege resource."},
			{Name: "isOuScopable", Type: proto.ColumnType_BOOL, Description: "Indicates whether the privilege is organizational unit scorable."},
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "The kind of privilege resource (e.g., Privilege)."},

			// Specific privilege information
			{Name: "privilegeName", Type: proto.ColumnType_STRING, Description: "The name of the privilege."},
			{Name: "serviceId", Type: proto.ColumnType_STRING, Description: "The ID of the service associated with the privilege."},
			{Name: "serviceName", Type: proto.ColumnType_STRING, Description: "The name of the service associated with the privilege."},

			// JSON column for child privileges (nested list)
			{Name: "childPrivileges", Type: proto.ColumnType_JSON, Description: "A list of child privileges associated with this privilege."},
		}),
	}
}
