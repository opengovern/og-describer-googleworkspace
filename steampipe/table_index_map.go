package steampipe

import (
	"github.com/opengovern/og-describer-googleworkspace/pkg/sdk/es"
)

var Map = map[string]string{
  "GoogleWorkspace/User": "google_workspace_user",
  "GoogleWorkspace/UserAlias": "google_workspace_user_alias",
  "GoogleWorkspace/Group": "google_workspace_group",
  "GoogleWorkspace/GroupMember": "google_workspace_group_member",
  "GoogleWorkspace/GroupAlias": "google_workspace_group_alias",
  "GoogleWorkspace/OrgUnit": "google_workspace_org_unit",
  "GoogleWorkspace/MobileDevice": "google_workspace_mobile_device",
  "GoogleWorkspace/ChromeDevice": "google_workspace_chrome_device",
  "GoogleWorkspace/Role": "google_workspace_role",
  "GoogleWorkspace/RoleAssignment": "google_workspace_role_assignment",
  "GoogleWorkspace/Domain": "google_workspace_domain",
  "GoogleWorkspace/DomainAlias": "google_workspace_domain_alias",
  "GoogleWorkspace/Privilege": "google_workspace_privilege",
  "GoogleWorkspace/ResourceBuilding": "google_workspace_resource_building",
  "GoogleWorkspace/ResourceCalender": "google_workspace_resource_calender",
  "GoogleWorkspace/ResourceFeature": "google_workspace_resource_feature",
}

var DescriptionMap = map[string]interface{}{
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

var ReverseMap = map[string]string{
  "google_workspace_user": "GoogleWorkspace/User",
  "google_workspace_user_alias": "GoogleWorkspace/UserAlias",
  "google_workspace_group": "GoogleWorkspace/Group",
  "google_workspace_group_member": "GoogleWorkspace/GroupMember",
  "google_workspace_group_alias": "GoogleWorkspace/GroupAlias",
  "google_workspace_org_unit": "GoogleWorkspace/OrgUnit",
  "google_workspace_mobile_device": "GoogleWorkspace/MobileDevice",
  "google_workspace_chrome_device": "GoogleWorkspace/ChromeDevice",
  "google_workspace_role": "GoogleWorkspace/Role",
  "google_workspace_role_assignment": "GoogleWorkspace/RoleAssignment",
  "google_workspace_domain": "GoogleWorkspace/Domain",
  "google_workspace_domain_alias": "GoogleWorkspace/DomainAlias",
  "google_workspace_privilege": "GoogleWorkspace/Privilege",
  "google_workspace_resource_building": "GoogleWorkspace/ResourceBuilding",
  "google_workspace_resource_calender": "GoogleWorkspace/ResourceCalender",
  "google_workspace_resource_feature": "GoogleWorkspace/ResourceFeature",
}
