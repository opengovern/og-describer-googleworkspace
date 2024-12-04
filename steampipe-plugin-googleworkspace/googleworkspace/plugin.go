/*
Package googleworkspace implements a steampipe plugin for googleworkspace.

This plugin provides data that Steampipe uses to present foreign
tables that represent Google Workspace resources.
*/
package googleworkspace

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-googleworkspace"

// Plugin creates this (googleworkspace) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		TableMap: map[string]*plugin.Table{
			//"googleworkspace_calendar":                tableGoogleWorkspaceCalendar(ctx),
			//"googleworkspace_drive":                   tableGoogleWorkspaceDrive(ctx),
			//"googleworkspace_gmail_settings":          tableGoogleWorkspaceGmailSettings(ctx),
			"googleworkspace_user":              tableGoogleWorkspaceUser(ctx),
			"googleworkspace_user_alias":        tableGoogleWorkspaceUserAlias(ctx),
			"googleworkspace_group":             tableGoogleWorkspaceGroup(ctx),
			"googleworkspace_group_member":      tableGoogleWorkspaceGroupMember(ctx),
			"googleworkspace_group_alias":       tableGoogleWorkspaceGroupAlias(ctx),
			"googleworkspace_org_unit":          tableGoogleWorkspaceOrgUnit(ctx),
			"googleworkspace_mobile_device":     tableGoogleWorkspaceMobileDevice(ctx),
			"googleworkspace_chrome_device":     tableGoogleWorkspaceChromeDevice(ctx),
			"googleworkspace_role":              tableGoogleWorkspaceRole(ctx),
			"googleworkspace_role_assignment":   tableGoogleWorkspaceRoleAssignment(ctx),
			"googleworkspace_domain":            tableGoogleWorkspaceDomain(ctx),
			"googleworkspace_domain_alias":      tableGoogleWorkspaceDomainAlias(ctx),
			"googleworkspace_privilege":         tableGoogleWorkspacePrivilege(ctx),
			"googleworkspace_resource_building": tableGoogleWorkspaceResourceBuilding(ctx),
			"googleworkspace_resource_calender": tableGoogleWorkspaceResourceCalender(ctx),
			"googleworkspace_resource_feature":  tableGoogleWorkspaceResourceFeature(ctx),
		},
	}

	return p
}
