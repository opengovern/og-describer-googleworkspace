package googleworkspace

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-googleworkspace/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGoogleWorkspaceGroupMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_group_member",
		Description: "Group members in the Google Workspace domain.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListGroupMember,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetGroupMember,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "Resource kind identifier, typically 'admin#directory#groupMember'."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "The email address of the group member."},
			{Name: "role", Type: proto.ColumnType_STRING, Description: "The role of the user within the group, such as 'OWNER', 'MEMBER', or 'MANAGER'."},
			{Name: "etag", Type: proto.ColumnType_STRING, Description: "ETag of the resource."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of member (can be 'USER' or 'SERVICE_ACCOUNT')."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The membership status of the user in the group, such as 'ACTIVE' or 'INVITED'."},
			{Name: "delivery_settings", Type: proto.ColumnType_STRING, Description: "The email delivery setting for the member. It can be 'ALL_EMAILS', 'NOTIFICATIONS_ONLY', or 'NONE'."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique ID of the group member."},
		}),
	}
}
