package maps

import (
	"github.com/opengovern/og-describer-googleworkspace/discovery/pkg/es"
)

var ResourceTypesToTables = map[string]string{
  "GoogleWorkspace/User": "googleworkspace_user",
  "GoogleWorkspace/UserAlias": "googleworkspace_user_alias",
  "GoogleWorkspace/Group": "googleworkspace_group",
  "GoogleWorkspace/GroupMember": "googleworkspace_group_member",
  "GoogleWorkspace/GroupAlias": "googleworkspace_group_alias",
  "GoogleWorkspace/OrgUnit": "googleworkspace_org_unit",
  "GoogleWorkspace/MobileDevice": "googleworkspace_mobile_device",
  "GoogleWorkspace/ChromeDevice": "googleworkspace_chrome_device",
  "GoogleWorkspace/Role": "googleworkspace_role",
  "GoogleWorkspace/RoleAssignment": "googleworkspace_role_assignment",
  "GoogleWorkspace/Domain": "googleworkspace_domain",
  "GoogleWorkspace/DomainAlias": "googleworkspace_domain_alias",
  "GoogleWorkspace/Privilege": "googleworkspace_privilege",
  "GoogleWorkspace/ResourceBuilding": "googleworkspace_resource_building",
  "GoogleWorkspace/ResourceCalender": "googleworkspace_resource_calender",
  "GoogleWorkspace/ResourceFeature": "googleworkspace_resource_feature",
}

var ResourceTypeToDescription = map[string]interface{}{
  "GoogleWorkspace/User": opengovernance.User{},
  "GoogleWorkspace/UserAlias": opengovernance.UserAlias{},
  "GoogleWorkspace/Group": opengovernance.Group{},
  "GoogleWorkspace/GroupMember": opengovernance.GroupMember{},
  "GoogleWorkspace/GroupAlias": opengovernance.GroupAlias{},
  "GoogleWorkspace/OrgUnit": opengovernance.OrgUnit{},
  "GoogleWorkspace/MobileDevice": opengovernance.MobileDevice{},
  "GoogleWorkspace/ChromeDevice": opengovernance.ChromeDevice{},
  "GoogleWorkspace/Role": opengovernance.Role{},
  "GoogleWorkspace/RoleAssignment": opengovernance.RoleAssignment{},
  "GoogleWorkspace/Domain": opengovernance.Domain{},
  "GoogleWorkspace/DomainAlias": opengovernance.DomainAlias{},
  "GoogleWorkspace/Privilege": opengovernance.Privilege{},
  "GoogleWorkspace/ResourceBuilding": opengovernance.ResourceBuilding{},
  "GoogleWorkspace/ResourceCalender": opengovernance.ResourceCalender{},
  "GoogleWorkspace/ResourceFeature": opengovernance.ResourceFeature{},
}

var TablesToResourceTypes = map[string]string{
  "googleworkspace_user": "GoogleWorkspace/User",
  "googleworkspace_user_alias": "GoogleWorkspace/UserAlias",
  "googleworkspace_group": "GoogleWorkspace/Group",
  "googleworkspace_group_member": "GoogleWorkspace/GroupMember",
  "googleworkspace_group_alias": "GoogleWorkspace/GroupAlias",
  "googleworkspace_org_unit": "GoogleWorkspace/OrgUnit",
  "googleworkspace_mobile_device": "GoogleWorkspace/MobileDevice",
  "googleworkspace_chrome_device": "GoogleWorkspace/ChromeDevice",
  "googleworkspace_role": "GoogleWorkspace/Role",
  "googleworkspace_role_assignment": "GoogleWorkspace/RoleAssignment",
  "googleworkspace_domain": "GoogleWorkspace/Domain",
  "googleworkspace_domain_alias": "GoogleWorkspace/DomainAlias",
  "googleworkspace_privilege": "GoogleWorkspace/Privilege",
  "googleworkspace_resource_building": "GoogleWorkspace/ResourceBuilding",
  "googleworkspace_resource_calender": "GoogleWorkspace/ResourceCalender",
  "googleworkspace_resource_feature": "GoogleWorkspace/ResourceFeature",
}
