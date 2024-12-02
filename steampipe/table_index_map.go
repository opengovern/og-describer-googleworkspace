package steampipe

import (
	"github.com/opengovern/og-describer-googleworkspace/pkg/sdk/es"
)

var Map = map[string]string{
  "GoogleWorkspace/User": "google_workspace_user",
  "GoogleWorkspace/UserAlias": "google_workspace_user_alias",
  "GoogleWorkspace/Group": "google_workspace_group",
  "GoogleWorkspace/GroupMember": "google_workspace_group_member",
  "GoogleWorkspace/OrgUnit": "google_workspace_org_unit",
}

var DescriptionMap = map[string]interface{}{
  "GoogleWorkspace/User": opengovernance.User{},
  "GoogleWorkspace/UserAlias": opengovernance.UserAlias{},
  "GoogleWorkspace/Group": opengovernance.Group{},
  "GoogleWorkspace/GroupMember": opengovernance.GroupMember{},
  "GoogleWorkspace/OrgUnit": opengovernance.OrgUnit{},
}

var ReverseMap = map[string]string{
  "google_workspace_user": "GoogleWorkspace/User",
  "google_workspace_user_alias": "GoogleWorkspace/UserAlias",
  "google_workspace_group": "GoogleWorkspace/Group",
  "google_workspace_group_member": "GoogleWorkspace/GroupMember",
  "google_workspace_org_unit": "GoogleWorkspace/OrgUnit",
}
