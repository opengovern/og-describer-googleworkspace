package configs

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "googleworkspace"                                    // example: aws, azure
	IntegrationName      = integration.Type("GOOGLE_WORKSPACE_ACCOUNT")         // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-googleworkspace" // example: github.com/opengovern/og-describer-aws
)
