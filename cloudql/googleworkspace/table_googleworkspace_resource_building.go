package googleworkspace

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-googleworkspace/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGoogleWorkspaceResourceBuilding(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_resource_building",
		Description: "Details about a building, including address, coordinates, and floor information.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListResourceBuilding,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetResourceBuilding,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "buildingId", Type: proto.ColumnType_STRING, Description: "The unique ID of the building."},
			{Name: "buildingName", Type: proto.ColumnType_STRING, Description: "The name of the building."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description or purpose of the building."},
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "The kind of building resource."},
			{Name: "etags", Type: proto.ColumnType_STRING, Description: "The etag identifier for the building resource."},

			// Nested JSON columns
			{Name: "coordinates", Type: proto.ColumnType_JSON, Description: "The geographic coordinates of the building (latitude, longitude)."},
			{Name: "address", Type: proto.ColumnType_JSON, Description: "The address of the building, including street, city, state, and postal code."},

			// Array columns
			{Name: "floorNames", Type: proto.ColumnType_JSON, Description: "A list of floor names in the building."},
		}),
	}
}
