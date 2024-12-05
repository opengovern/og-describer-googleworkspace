package googleworkspace

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-googleworkspace/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGoogleWorkspaceResourceCalender(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_resource_calender",
		Description: "Details about resources in a building, including capacity, floor, and category.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListResourceCalender,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetResourceCalender,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "The kind of resource."},
			{Name: "etags", Type: proto.ColumnType_STRING, Description: "The etag identifier for the resource."},
			{Name: "resourceName", Type: proto.ColumnType_STRING, Description: "The name of the resource."},
			{Name: "resourceDescription", Type: proto.ColumnType_STRING, Description: "The description of the resource."},
			{Name: "resourceType", Type: proto.ColumnType_STRING, Description: "The type of the resource."},
			{Name: "resourceEmail", Type: proto.ColumnType_STRING, Description: "The email associated with the resource."},
			{Name: "resourceCategory", Type: proto.ColumnType_STRING, Description: "The category of the resource."},
			{Name: "userVisibleDescription", Type: proto.ColumnType_STRING, Description: "A description that is visible to the user."},
			{Name: "generatedResourceName", Type: proto.ColumnType_STRING, Description: "The generated name of the resource."},
			{Name: "resourceId", Type: proto.ColumnType_STRING, Description: "The unique identifier for the resource."},

			// Integer and JSON columns
			{Name: "capacity", Type: proto.ColumnType_INT, Description: "The capacity of the resource."},
			{Name: "floorName", Type: proto.ColumnType_STRING, Description: "The name of the floor where the resource is located."},
			{Name: "buildingId", Type: proto.ColumnType_STRING, Description: "The unique identifier for the building."},
			{Name: "floorSection", Type: proto.ColumnType_STRING, Description: "The section of the floor where the resource is located."},

			// Feature instances - JSON column
			{Name: "featureInstances", Type: proto.ColumnType_JSON, Description: "Feature instances associated with the resource."},
		}),
	}
}
