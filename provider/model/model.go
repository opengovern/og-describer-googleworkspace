//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package model

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
