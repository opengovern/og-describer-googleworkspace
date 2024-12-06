package googleworkspace

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	opengovernance "github.com/opengovern/og-describer-googleworkspace/pkg/sdk/es"
)

func tableGoogleWorkspaceDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_domain",
		Description: "Information about domain aliases in Google Workspace, including domain name, verification status, and aliases.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDomain,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetDomain,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "The type or kind of the domain alias."},
			{Name: "domainName", Type: proto.ColumnType_STRING, Description: "The domain name for the alias."},
			{Name: "verified", Type: proto.ColumnType_BOOL, Description: "Indicates whether the domain alias has been verified."},
			{Name: "etag", Type: proto.ColumnType_STRING, Description: "The etag identifier for the resource."},
			{Name: "creationTime", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the domain alias was created."},
			{Name: "isPrimary", Type: proto.ColumnType_BOOL, Description: "Indicates whether this is the primary domain alias."},

			// JSON column for domain aliases
			{Name: "domainAliases", Type: proto.ColumnType_JSON, Description: "List of domain aliases."},
		}),
	}
}
