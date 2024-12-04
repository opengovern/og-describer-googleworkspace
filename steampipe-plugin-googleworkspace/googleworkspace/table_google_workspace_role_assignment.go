package googleworkspace

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGoogleWorkspaceRoleAssignment(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "google_workspace_role_assignment",
		Description: "Details about role assignments in Google Workspace, including assignee, scope, and conditions.",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "roleAssignmentId", Type: proto.ColumnType_STRING, Description: "The unique identifier of the role assignment."},
			{Name: "roleId", Type: proto.ColumnType_STRING, Description: "The ID of the assigned role."},
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "The type of the role assignment."},
			{Name: "etag", Type: proto.ColumnType_STRING, Description: "The etag identifier for the resource."},
			{Name: "assignedTo", Type: proto.ColumnType_STRING, Description: "The entity to which the role is assigned (e.g., user or group)."},
			{Name: "assigneeType", Type: proto.ColumnType_STRING, Description: "The type of assignee (e.g., USER, GROUP, etc.)."},
			{Name: "scopeType", Type: proto.ColumnType_STRING, Description: "The scope type of the assignment (e.g., ORGANIZATION, UNIT, etc.)."},
			{Name: "orgUnitId", Type: proto.ColumnType_STRING, Description: "The ID of the organizational unit to which the role is assigned."},
			{Name: "condition", Type: proto.ColumnType_STRING, Description: "A condition that must be met for the role assignment to apply."},
		}),
	}
}
