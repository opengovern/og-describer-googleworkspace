package global

import (
	"github.com/opengovern/og-util/pkg/integration"
)

const (
	IntegrationTypeLower = "googleworkspace"                                    // example: aws, azure
	IntegrationName      = integration.Type("google_workspace_account")         // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-googleworkspace" // example: github.com/opengovern/og-describer-aws
)

type IntegrationCredentials struct {
	AdminEmail string `json:"admin_email"`
	CustomerID string `json:"customer_id"`
	KeyFile    string `json:"key_file"`
}
