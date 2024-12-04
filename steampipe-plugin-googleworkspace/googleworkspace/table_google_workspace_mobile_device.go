package googleworkspace

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGoogleWorkspaceMobileDevice(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "google_workspace_mobile_device",
		Description: "Device details including model, status, hardware, and applications.",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "The type of the device (e.g., android_device)."},
			{Name: "etag", Type: proto.ColumnType_STRING, Description: "The etag identifier for the resource."},
			{Name: "resourceId", Type: proto.ColumnType_STRING, Description: "Unique resource identifier for the device."},
			{Name: "deviceId", Type: proto.ColumnType_STRING, Description: "Unique device identifier."},
			{Name: "model", Type: proto.ColumnType_STRING, Description: "The model of the device."},
			{Name: "os", Type: proto.ColumnType_STRING, Description: "The operating system running on the device."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of device (e.g., smartphone)."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The current status of the device (e.g., active, inactive)."},
			{Name: "hardwareId", Type: proto.ColumnType_STRING, Description: "Hardware identifier of the device."},
			{Name: "firstSync", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the device was first synced."},
			{Name: "lastSync", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the device was last synced."},
			{Name: "userAgent", Type: proto.ColumnType_STRING, Description: "The user agent string for the device."},
			{Name: "serialNumber", Type: proto.ColumnType_STRING, Description: "The serial number of the device."},
			{Name: "imei", Type: proto.ColumnType_STRING, Description: "IMEI number of the device."},
			{Name: "meid", Type: proto.ColumnType_STRING, Description: "MEID number of the device."},
			{Name: "wifiMacAddress", Type: proto.ColumnType_STRING, Description: "MAC address of the device's Wi-Fi."},
			{Name: "networkOperator", Type: proto.ColumnType_STRING, Description: "The network operator associated with the device."},
			{Name: "defaultLanguage", Type: proto.ColumnType_STRING, Description: "The default language setting on the device."},
			{Name: "managedAccountIsOnOwnerProfile", Type: proto.ColumnType_BOOL, Description: "Indicates if the managed account is on the owner's profile."},
			{Name: "deviceCompromisedStatus", Type: proto.ColumnType_STRING, Description: "Status of device compromise (e.g., none, compromised)."},
			{Name: "buildNumber", Type: proto.ColumnType_STRING, Description: "The build number of the device."},
			{Name: "kernelVersion", Type: proto.ColumnType_STRING, Description: "The kernel version of the device."},
			{Name: "basebandVersion", Type: proto.ColumnType_STRING, Description: "The baseband version of the device."},
			{Name: "unknownSourcesStatus", Type: proto.ColumnType_BOOL, Description: "Indicates if unknown sources are allowed."},
			{Name: "developerOptionsStatus", Type: proto.ColumnType_BOOL, Description: "Indicates if developer options are enabled."},
			{Name: "adbStatus", Type: proto.ColumnType_BOOL, Description: "Indicates if ADB (Android Debug Bridge) is enabled."},
			{Name: "supportsWorkProfile", Type: proto.ColumnType_BOOL, Description: "Indicates if the device supports work profiles."},
			{Name: "manufacturer", Type: proto.ColumnType_STRING, Description: "The manufacturer of the device."},
			{Name: "releaseVersion", Type: proto.ColumnType_STRING, Description: "The release version of the OS."},
			{Name: "securityPatchLevel", Type: proto.ColumnType_STRING, Description: "The latest security patch level installed on the device."},
			{Name: "brand", Type: proto.ColumnType_STRING, Description: "The brand name of the device."},
			{Name: "bootloaderVersion", Type: proto.ColumnType_STRING, Description: "The bootloader version of the device."},
			{Name: "hardware", Type: proto.ColumnType_STRING, Description: "Hardware name of the device."},
			{Name: "encryptionStatus", Type: proto.ColumnType_STRING, Description: "Indicates the encryption status of the device."},
			{Name: "devicePasswordStatus", Type: proto.ColumnType_STRING, Description: "Status of device password (e.g., set, not set)."},
			{Name: "privilege", Type: proto.ColumnType_STRING, Description: "The privilege level of the device."},

			// JSON column for applications
			{Name: "applications", Type: proto.ColumnType_JSON, Description: "List of installed applications on the device."},
		},
	}
}
