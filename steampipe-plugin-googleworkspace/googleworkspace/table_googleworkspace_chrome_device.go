package googleworkspace

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGoogleWorkspaceChromeDevice(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_chrome_device",
		Description: "Details of mobile devices, including model, status, hardware, and software information.",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: commonColumns([]*plugin.Column{
			// Basic details columns
			{Name: "deviceId", Type: proto.ColumnType_STRING, Description: "Unique device identifier."},
			{Name: "serialNumber", Type: proto.ColumnType_STRING, Description: "The serial number of the device."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The current status of the device."},
			{Name: "lastSync", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of the last device sync."},
			{Name: "supportEndDate", Type: proto.ColumnType_TIMESTAMP, Description: "The support end date for the device."},
			{Name: "model", Type: proto.ColumnType_STRING, Description: "Model of the device."},
			{Name: "meid", Type: proto.ColumnType_STRING, Description: "MEID number of the device."},
			{Name: "orderNumber", Type: proto.ColumnType_STRING, Description: "Order number associated with the device."},
			{Name: "willAutoRenew", Type: proto.ColumnType_BOOL, Description: "Indicates if the device will auto-renew."},
			{Name: "osVersion", Type: proto.ColumnType_STRING, Description: "The OS version installed on the device."},
			{Name: "platformVersion", Type: proto.ColumnType_STRING, Description: "The platform version of the device."},
			{Name: "firmwareVersion", Type: proto.ColumnType_STRING, Description: "The firmware version installed on the device."},
			{Name: "macAddress", Type: proto.ColumnType_STRING, Description: "The MAC address of the device."},
			{Name: "bootMode", Type: proto.ColumnType_STRING, Description: "The boot mode of the device."},
			{Name: "lastEnrollmentTime", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the device was last enrolled."},
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "The type of the device."},

			// Recent users columns
			{Name: "recentUsers", Type: proto.ColumnType_JSON, Description: "List of recent users associated with the device."},

			// Time range columns
			{Name: "activeTimeRanges", Type: proto.ColumnType_JSON, Description: "Active time ranges associated with the device."},

			// Network columns
			{Name: "lastKnownNetwork", Type: proto.ColumnType_JSON, Description: "Last known network information for the device."},

			// System and hardware columns
			{Name: "systemRamTotal", Type: proto.ColumnType_STRING, Description: "Total system RAM."},
			{Name: "diskVolumeReports", Type: proto.ColumnType_JSON, Description: "Reports on the device's disk volumes."},
			{Name: "cpuStatusReports", Type: proto.ColumnType_JSON, Description: "Reports on the CPU status."},
			{Name: "cpuInfo", Type: proto.ColumnType_JSON, Description: "CPU information."},
			{Name: "deviceFiles", Type: proto.ColumnType_JSON, Description: "List of device files."},

			// Additional information
			{Name: "autoUpdateExpiration", Type: proto.ColumnType_TIMESTAMP, Description: "Expiration date for auto-update."},
			{Name: "ethernetMacAddress", Type: proto.ColumnType_STRING, Description: "Ethernet MAC address."},
			{Name: "orgUnitPath", Type: proto.ColumnType_STRING, Description: "Organizational unit path for the device."},
			{Name: "tpmVersionInfo", Type: proto.ColumnType_JSON, Description: "TPM version details."},
			{Name: "osUpdateStatus", Type: proto.ColumnType_JSON, Description: "Operating system update status."},
			{Name: "extendedSupportEligible", Type: proto.ColumnType_BOOL, Description: "Indicates if extended support is eligible."},
			{Name: "extendedSupportStart", Type: proto.ColumnType_TIMESTAMP, Description: "Start date of extended support."},
			{Name: "extendedSupportEnabled", Type: proto.ColumnType_BOOL, Description: "Indicates if extended support is enabled."},

			// Fan and backlight info
			{Name: "fanInfo", Type: proto.ColumnType_JSON, Description: "Information about the device's fan."},
			{Name: "backlightInfo", Type: proto.ColumnType_JSON, Description: "Information about the device's backlight."},
		}),
	}
}
