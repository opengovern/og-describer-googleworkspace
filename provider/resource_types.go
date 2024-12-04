package provider

import (
	model "github.com/opengovern/og-describer-googleworkspace/pkg/sdk/models"
	"github.com/opengovern/og-describer-googleworkspace/provider/configs"
	"github.com/opengovern/og-describer-googleworkspace/provider/describer"
)

var ResourceTypes = map[string]model.ResourceType{

	"GoogleWorkspace/User": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "GoogleWorkspace/User",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByGoogleWorkspace(describer.ListUsers),
		GetDescriber:    DescribeSingleByGoogleWorkspace(describer.GetUser),
	},

	"GoogleWorkspace/UserAlias": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "GoogleWorkspace/UserAlias",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByGoogleWorkspace(describer.ListUserAliases),
		GetDescriber:    nil,
	},

	"GoogleWorkspace/Group": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "GoogleWorkspace/Group",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByGoogleWorkspace(describer.ListGroups),
		GetDescriber:    DescribeSingleByGoogleWorkspace(describer.GetGroup),
	},

	"GoogleWorkspace/GroupMember": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "GoogleWorkspace/GroupMember",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByGoogleWorkspace(describer.ListGroupMembers),
		GetDescriber:    nil,
	},

	"GoogleWorkspace/GroupAlias": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "GoogleWorkspace/GroupAlias",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByGoogleWorkspace(describer.ListGroupAliases),
		GetDescriber:    nil,
	},

	"GoogleWorkspace/OrgUnit": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "GoogleWorkspace/OrgUnit",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByGoogleWorkspace(describer.ListOrgUnits),
		GetDescriber:    DescribeSingleByGoogleWorkspace(describer.GetOrgUnit),
	},

	"GoogleWorkspace/MobileDevice": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "GoogleWorkspace/MobileDevice",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByGoogleWorkspace(describer.ListMobileDevices),
		GetDescriber:    DescribeSingleByGoogleWorkspace(describer.GetMobileDevice),
	},

	"GoogleWorkspace/ChromeDevice": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "GoogleWorkspace/ChromeDevice",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByGoogleWorkspace(describer.ListChromeDevices),
		GetDescriber:    DescribeSingleByGoogleWorkspace(describer.GetChromeDevice),
	},

	"GoogleWorkspace/Role": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "GoogleWorkspace/Role",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByGoogleWorkspace(describer.ListRoles),
		GetDescriber:    DescribeSingleByGoogleWorkspace(describer.GetRole),
	},

	"GoogleWorkspace/RoleAssignment": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "GoogleWorkspace/RoleAssignment",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByGoogleWorkspace(describer.ListRoleAssignments),
		GetDescriber:    DescribeSingleByGoogleWorkspace(describer.GetRoleAssignment),
	},

	"GoogleWorkspace/Domain": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "GoogleWorkspace/Domain",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByGoogleWorkspace(describer.ListDomains),
		GetDescriber:    DescribeSingleByGoogleWorkspace(describer.GetDomain),
	},

	"GoogleWorkspace/DomainAlias": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "GoogleWorkspace/DomainAlias",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByGoogleWorkspace(describer.ListDomainAliases),
		GetDescriber:    DescribeSingleByGoogleWorkspace(describer.GetDomainAlias),
	},

	"GoogleWorkspace/Privilege": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "GoogleWorkspace/Privilege",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByGoogleWorkspace(describer.ListPrivileges),
		GetDescriber:    nil,
	},

	"GoogleWorkspace/ResourceBuilding": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "GoogleWorkspace/ResourceBuilding",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByGoogleWorkspace(describer.ListResourceBuildings),
		GetDescriber:    DescribeSingleByGoogleWorkspace(describer.GetResourceBuilding),
	},

	"GoogleWorkspace/ResourceCalender": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "GoogleWorkspace/ResourceCalender",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByGoogleWorkspace(describer.ListResourceCalenders),
		GetDescriber:    DescribeSingleByGoogleWorkspace(describer.GetResourceCalender),
	},

	"GoogleWorkspace/ResourceFeature": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "GoogleWorkspace/ResourceFeature",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByGoogleWorkspace(describer.ListResourceFeatures),
		GetDescriber:    DescribeSingleByGoogleWorkspace(describer.GetResourceFeature),
	},
}
