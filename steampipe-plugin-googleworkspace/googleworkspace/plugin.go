/*
Package googleworkspace implements a steampipe plugin for googleworkspace.

This plugin provides data that Steampipe uses to present foreign
tables that represent Google Workspace resources.
*/
package googleworkspace

import (
	"context"
	essdk "github.com/opengovern/og-util/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-googleworkspace"

// Plugin creates this (googleworkspace) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: pluginName,
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: essdk.ConfigInstance,
			Schema:      essdk.ConfigSchema(),
		},
		DefaultTransform: transform.FromCamel(),
		TableMap: map[string]*plugin.Table{
			//"googleworkspace_calendar":                tableGoogleWorkspaceCalendar(ctx),
			//"googleworkspace_calendar_event":          tableGoogleWorkspaceCalendarEvent(ctx),
			//"googleworkspace_calendar_my_event":       tableGoogleWorkspaceCalendarMyEvent(ctx),
			//"googleworkspace_drive":                   tableGoogleWorkspaceDrive(ctx),
			//"googleworkspace_drive_my_file":           tableGoogleWorkspaceDriveMyFile(ctx),
			//"googleworkspace_gmail_draft":             tableGoogleWorkspaceGmailDraft(ctx),
			//"googleworkspace_gmail_message":           tableGoogleWorkspaceGmailMessage(ctx),
			//"googleworkspace_gmail_my_draft":          tableGoogleWorkspaceGmailMyDraft(ctx),
			//"googleworkspace_gmail_my_message":        tableGoogleWorkspaceGmailMyMessage(ctx),
			//"googleworkspace_gmail_my_settings":       tableGoogleWorkspaceGmailMySettings(ctx),
			//"googleworkspace_gmail_settings":          tableGoogleWorkspaceGmailSettings(ctx),
			//"googleworkspace_people_contact":          tableGoogleWorkspacePeopleContact(ctx),
			//"googleworkspace_people_contact_group":    tableGoogleWorkspacePeopleContactGroup(ctx),
			//"googleworkspace_people_directory_people": tableGoogleWorkspacePeopleDirectoryPeople(ctx),
			"googleworkspace_user":         tableGoogleWorkspaceUser(ctx),
			"googleworkspace_user_alias":   tableGoogleWorkspaceUserAlias(ctx),
			"googleworkspace_group":        tableGoogleWorkspaceGroup(ctx),
			"googleworkspace_group_member": tableGoogleWorkspaceGroupMember(ctx),
			"googleworkspace_org_unit":     tableGoogleWorkspaceOrgUnit(ctx),
		},
	}
	for key, table := range p.TableMap {
		if table == nil {
			continue
		}
		if table.Get != nil && table.Get.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}
		if table.List != nil && table.List.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}

		opengovernanceTable := false
		for _, col := range table.Columns {
			if col != nil && col.Name == "platform_account_id" {
				opengovernanceTable = true
			}
		}

		if opengovernanceTable {
			if table.Get != nil {
				table.Get.KeyColumns = append(table.Get.KeyColumns, plugin.OptionalColumns([]string{"platform_account_id", "platform_resource_id"})...)
			}

			if table.List != nil {
				table.List.KeyColumns = append(table.List.KeyColumns, plugin.OptionalColumns([]string{"platform_account_id", "platform_resource_id"})...)
			}
		}
	}
	return p
}
