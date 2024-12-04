package googleworkspace

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-googleworkspace/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGoogleWorkspaceOrgUnit(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "google_workspace_org_unit",
		Description: "Organizational Units in the Google Workspace domain.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListOrgUnit,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("orgUnitId"),
			Hydrate:    opengovernance.GetOrgUnit,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "Resource kind identifier, typically 'admin#directory#orgUnit'."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the organizational unit."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "A description of the organizational unit."},
			{Name: "etag", Type: proto.ColumnType_STRING, Description: "ETag of the resource."},
			{Name: "blockInheritance", Type: proto.ColumnType_BOOL, Description: "Indicates if inheritance of organizational unit settings is blocked."},
			{Name: "orgUnitId", Type: proto.ColumnType_STRING, Description: "The unique ID of the organizational unit."},
			{Name: "orgUnitPath", Type: proto.ColumnType_STRING, Description: "The full path of the organizational unit in the domain."},
			{Name: "parentOrgUnitId", Type: proto.ColumnType_STRING, Description: "The ID of the parent organizational unit."},
			{Name: "parentOrgUnitPath", Type: proto.ColumnType_STRING, Description: "The path of the parent organizational unit."},
		}),
	}
}
