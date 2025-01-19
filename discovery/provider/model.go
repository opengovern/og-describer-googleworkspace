//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package provider

import admin "google.golang.org/api/admin/directory/v1"

type Metadata struct{}

type UserDescription struct {
	admin.User
}

type UserAliasDescription struct {
	admin.UserAlias
}

type GroupDescription struct {
	admin.Group
}

type GroupMemberDescription struct {
	admin.Member
}

type OrgUnitDescription struct {
	admin.OrgUnit
}

type MobileDeviceDescription struct {
	admin.MobileDevice
}

type ChromeDeviceDescription struct {
	admin.ChromeOsDevice
}

type RoleDescription struct {
	admin.Role
}

type RoleAssignmentDescription struct {
	admin.RoleAssignment
}

type DomainDescription struct {
	admin.Domains
}

type DomainAliasDescription struct {
	admin.DomainAlias
}

type GroupAliasDescription struct {
	admin.GroupAlias
}

type PrivilegeDescription struct {
	admin.Privilege
}

type ResourceBuildingDescription struct {
	admin.Building
}

type ResourceCalenderDescription struct {
	admin.CalendarResource
}

type ResourceFeatureDescription struct {
	admin.Feature
}
