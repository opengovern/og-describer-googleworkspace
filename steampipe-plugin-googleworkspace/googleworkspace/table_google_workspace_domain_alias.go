package googleworkspace

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-googleworkspace/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGoogleWorkspaceDomainAlias(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "google_workspace_domain_alias",
		Description: "Information about a domain alias entry, including the parent domain, alias name, and verification status.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDomainAlias,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetDomainAlias,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "The type or kind of the domain alias entry."},
			{Name: "parentDomainName", Type: proto.ColumnType_STRING, Description: "The parent domain associated with the alias."},
			{Name: "domainAliasName", Type: proto.ColumnType_STRING, Description: "The name of the domain alias."},
			{Name: "verified", Type: proto.ColumnType_BOOL, Description: "Indicates whether the domain alias has been verified."},
			{Name: "etag", Type: proto.ColumnType_STRING, Description: "The etag identifier for the resource."},
			{Name: "creationTime", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the domain alias entry was created."},
		}),
	}
}
