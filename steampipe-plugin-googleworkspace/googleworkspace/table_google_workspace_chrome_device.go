package googleworkspace

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGoogleWorkspaceChromeDevice(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "google_workspace_chrome_device",
		Description: "Device details including model, status, system reports, and active time ranges.",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "deviceId", Type: proto.ColumnType_STRING, Description: "Unique device identifier."},
			{Name: "serialNumber", Type: proto.ColumnType_STRING, Description: "The serial number of the device."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The current status of the device."},
			{Name: "lastSync", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the device was last synced."},
			{Name: "supportEndDate", Type: proto.ColumnType_TIMESTAMP, Description: "The date when the device's support ends."},
			{Name: "annotatedUser", Type: proto.ColumnType_STRING, Description: "User associated with the device."},
			{Name: "annotatedLocation", Type: proto.ColumnType_STRING, Description: "Location annotated for the device."},
			{Name: "notes", Type: proto.ColumnType_STRING, Description: "Additional notes about the device."},
			{Name: "model", Type: proto.ColumnType_STRING, Description: "The model of the device."},
			{Name: "meid", Type: proto.ColumnType_STRING, Description: "MEID number of the device."},
			{Name: "orderNumber", Type: proto.ColumnType_STRING, Description: "The order number for the device."},
			{Name: "willAutoRenew", Type: proto.ColumnType_BOOL, Description: "Indicates if the device will auto-renew."},
			{Name: "osVersion", Type: proto.ColumnType_STRING, Description: "The operating system version on the device."},
			{Name: "platformVersion", Type: proto.ColumnType_STRING, Description: "The platform version of the device."},
			{Name: "firmwareVersion", Type: proto.ColumnType_STRING, Description: "The firmware version of the device."},
			{Name: "macAddress", Type: proto.ColumnType_STRING, Description: "The MAC address of the device."},
			{Name: "bootMode", Type: proto.ColumnType_STRING, Description: "The boot mode of the device."},
			{Name: "lastEnrollmentTime", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the device was last enrolled."},
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "The type of the device (e.g., mobile, desktop)."},

			// JSON columns
			{Name: "recentUsers", Type: proto.ColumnType_JSON, Description: "List of recent users associated with the device."},
			{Name: "activeTimeRanges", Type: proto.ColumnType_JSON, Description: "Array of active time ranges for the device."},
			{Name: "diskVolumeReports", Type: proto.ColumnType_JSON, Description: "Disk volume usage reports."},
			{Name: "cpuStatusReports", Type: proto.ColumnType_JSON, Description: "CPU utilization and temperature reports."},
			{Name: "cpuInfo", Type: proto.ColumnType_JSON, Description: "Detailed CPU information including model, architecture, and scaling."},
			{Name: "deviceFiles", Type: proto.ColumnType_JSON, Description: "Files associated with the device."},
			{Name: "systemRamFreeReports", Type: proto.ColumnType_JSON, Description: "Reports on system RAM free space."},
			{Name: "lastKnownNetwork", Type: proto.ColumnType_JSON, Description: "Information on the last known network connection."},
			{Name: "fanInfo", Type: proto.ColumnType_JSON, Description: "Fan status and information on the device."},
			{Name: "backlightInfo", Type: proto.ColumnType_JSON, Description: "Information on the device's backlight."},
			{Name: "screenshotFiles", Type: proto.ColumnType_JSON, Description: "List of screenshot files from the device."},

			// Other columns for specific configurations
			{Name: "autoUpdateExpiration", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of auto-update expiration."},
			{Name: "ethernetMacAddress0", Type: proto.ColumnType_STRING, Description: "Ethernet MAC address of the device."},
			{Name: "dockMacAddress", Type: proto.ColumnType_STRING, Description: "Docking station MAC address."},
			{Name: "manufactureDate", Type: proto.ColumnType_TIMESTAMP, Description: "The manufacture date of the device."},
			{Name: "orgUnitPath", Type: proto.ColumnType_STRING, Description: "Organizational unit path where the device resides."},
			{Name: "tpmVersionInfo", Type: proto.ColumnType_JSON, Description: "TPM version and related information."},
			{Name: "orgUnitId", Type: proto.ColumnType_STRING, Description: "Organizational unit ID for the device."},
			{Name: "osUpdateStatus", Type: proto.ColumnType_JSON, Description: "Status of the device's OS updates."},
			{Name: "firstEnrollmentTime", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the device was first enrolled."},
			{Name: "lastDeprovisionTimestamp", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the device was last deprovisioned."},
			{Name: "deprovisionReason", Type: proto.ColumnType_STRING, Description: "Reason for deprovisioning the device."},
			{Name: "deviceLicenseType", Type: proto.ColumnType_STRING, Description: "License type for the device."},
			{Name: "extendedSupportEligible", Type: proto.ColumnType_BOOL, Description: "Indicates if extended support is eligible for the device."},
			{Name: "extendedSupportStart", Type: proto.ColumnType_TIMESTAMP, Description: "Start date for extended support."},
			{Name: "extendedSupportEnabled", Type: proto.ColumnType_BOOL, Description: "Indicates if extended support is enabled."},
			{Name: "chromeOsType", Type: proto.ColumnType_STRING, Description: "The type of Chrome OS running on the device."},
			{Name: "diskSpaceUsage", Type: proto.ColumnType_JSON, Description: "Disk space usage details for the device."},
		},
	}
}
