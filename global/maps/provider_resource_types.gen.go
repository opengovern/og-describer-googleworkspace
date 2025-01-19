package maps

import (
	"github.com/opengovern/og-describer-googleworkspace/discovery/describers"
	model "github.com/opengovern/og-describer-googleworkspace/discovery/pkg/models"
	"github.com/opengovern/og-describer-googleworkspace/discovery/provider"
	"github.com/opengovern/og-describer-googleworkspace/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
)

var ResourceTypes = map[string]model.ResourceType{

	"GoogleWorkspace/User": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "GoogleWorkspace/User",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByGoogleWorkspace(describers.ListUsers),
		GetDescriber:    provider.DescribeSingleByGoogleWorkspace(describers.GetUser),
	},

	"GoogleWorkspace/UserAlias": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "GoogleWorkspace/UserAlias",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByGoogleWorkspace(describers.ListUserAliases),
		GetDescriber:    nil,
	},

	"GoogleWorkspace/Group": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "GoogleWorkspace/Group",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByGoogleWorkspace(describers.ListGroups),
		GetDescriber:    provider.DescribeSingleByGoogleWorkspace(describers.GetGroup),
	},

	"GoogleWorkspace/GroupMember": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "GoogleWorkspace/GroupMember",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByGoogleWorkspace(describers.ListGroupMembers),
		GetDescriber:    nil,
	},

	"GoogleWorkspace/GroupAlias": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "GoogleWorkspace/GroupAlias",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByGoogleWorkspace(describers.ListGroupAliases),
		GetDescriber:    nil,
	},

	"GoogleWorkspace/OrgUnit": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "GoogleWorkspace/OrgUnit",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByGoogleWorkspace(describers.ListOrgUnits),
		GetDescriber:    provider.DescribeSingleByGoogleWorkspace(describers.GetOrgUnit),
	},

	"GoogleWorkspace/MobileDevice": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "GoogleWorkspace/MobileDevice",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByGoogleWorkspace(describers.ListMobileDevices),
		GetDescriber:    provider.DescribeSingleByGoogleWorkspace(describers.GetMobileDevice),
	},

	"GoogleWorkspace/ChromeDevice": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "GoogleWorkspace/ChromeDevice",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByGoogleWorkspace(describers.ListChromeDevices),
		GetDescriber:    provider.DescribeSingleByGoogleWorkspace(describers.GetChromeDevice),
	},

	"GoogleWorkspace/Role": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "GoogleWorkspace/Role",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByGoogleWorkspace(describers.ListRoles),
		GetDescriber:    provider.DescribeSingleByGoogleWorkspace(describers.GetRole),
	},

	"GoogleWorkspace/RoleAssignment": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "GoogleWorkspace/RoleAssignment",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByGoogleWorkspace(describers.ListRoleAssignments),
		GetDescriber:    provider.DescribeSingleByGoogleWorkspace(describers.GetRoleAssignment),
	},

	"GoogleWorkspace/Domain": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "GoogleWorkspace/Domain",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByGoogleWorkspace(describers.ListDomains),
		GetDescriber:    provider.DescribeSingleByGoogleWorkspace(describers.GetDomain),
	},

	"GoogleWorkspace/DomainAlias": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "GoogleWorkspace/DomainAlias",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByGoogleWorkspace(describers.ListDomainAliases),
		GetDescriber:    provider.DescribeSingleByGoogleWorkspace(describers.GetDomainAlias),
	},

	"GoogleWorkspace/Privilege": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "GoogleWorkspace/Privilege",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByGoogleWorkspace(describers.ListPrivileges),
		GetDescriber:    nil,
	},

	"GoogleWorkspace/ResourceBuilding": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "GoogleWorkspace/ResourceBuilding",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByGoogleWorkspace(describers.ListResourceBuildings),
		GetDescriber:    provider.DescribeSingleByGoogleWorkspace(describers.GetResourceBuilding),
	},

	"GoogleWorkspace/ResourceCalender": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "GoogleWorkspace/ResourceCalender",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByGoogleWorkspace(describers.ListResourceCalenders),
		GetDescriber:    provider.DescribeSingleByGoogleWorkspace(describers.GetResourceCalender),
	},

	"GoogleWorkspace/ResourceFeature": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "GoogleWorkspace/ResourceFeature",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByGoogleWorkspace(describers.ListResourceFeatures),
		GetDescriber:    provider.DescribeSingleByGoogleWorkspace(describers.GetResourceFeature),
	},
}

var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{

	"GoogleWorkspace/User": {
		Name:            "GoogleWorkspace/User",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"GoogleWorkspace/UserAlias": {
		Name:            "GoogleWorkspace/UserAlias",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"GoogleWorkspace/Group": {
		Name:            "GoogleWorkspace/Group",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"GoogleWorkspace/GroupMember": {
		Name:            "GoogleWorkspace/GroupMember",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"GoogleWorkspace/GroupAlias": {
		Name:            "GoogleWorkspace/GroupAlias",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"GoogleWorkspace/OrgUnit": {
		Name:            "GoogleWorkspace/OrgUnit",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"GoogleWorkspace/MobileDevice": {
		Name:            "GoogleWorkspace/MobileDevice",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"GoogleWorkspace/ChromeDevice": {
		Name:            "GoogleWorkspace/ChromeDevice",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"GoogleWorkspace/Role": {
		Name:            "GoogleWorkspace/Role",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"GoogleWorkspace/RoleAssignment": {
		Name:            "GoogleWorkspace/RoleAssignment",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"GoogleWorkspace/Domain": {
		Name:            "GoogleWorkspace/Domain",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"GoogleWorkspace/DomainAlias": {
		Name:            "GoogleWorkspace/DomainAlias",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"GoogleWorkspace/Privilege": {
		Name:            "GoogleWorkspace/Privilege",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"GoogleWorkspace/ResourceBuilding": {
		Name:            "GoogleWorkspace/ResourceBuilding",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"GoogleWorkspace/ResourceCalender": {
		Name:            "GoogleWorkspace/ResourceCalender",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"GoogleWorkspace/ResourceFeature": {
		Name:            "GoogleWorkspace/ResourceFeature",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},
}

var ResourceTypesList = []string{
	"GoogleWorkspace/User",
	"GoogleWorkspace/UserAlias",
	"GoogleWorkspace/Group",
	"GoogleWorkspace/GroupMember",
	"GoogleWorkspace/GroupAlias",
	"GoogleWorkspace/OrgUnit",
	"GoogleWorkspace/MobileDevice",
	"GoogleWorkspace/ChromeDevice",
	"GoogleWorkspace/Role",
	"GoogleWorkspace/RoleAssignment",
	"GoogleWorkspace/Domain",
	"GoogleWorkspace/DomainAlias",
	"GoogleWorkspace/Privilege",
	"GoogleWorkspace/ResourceBuilding",
	"GoogleWorkspace/ResourceCalender",
	"GoogleWorkspace/ResourceFeature",
}
