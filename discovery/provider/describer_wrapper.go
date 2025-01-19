package provider

import (
	"errors"
	"fmt"
	"github.com/opengovern/og-describer-googleworkspace/discovery/pkg/models"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"golang.org/x/time/rate"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
	"time"
)

// DescribeListByGoogleWorkspace A wrapper to pass GoogleWorkspace authorization to describer functions
func DescribeListByGoogleWorkspace(describe func(context.Context, *GoogleWorkspaceAPIHandler, *models.StreamSender) ([]models.Resource, error)) models.ResourceDescriber {
	return func(ctx context.Context, cfg models.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *models.StreamSender) ([]models.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)

		var err error
		// Check for the keyFile content
		if string(cfg.KeyFile) == "" {
			return nil, errors.New("key file must be configured")
		}

		// Check for the admin email
		if string(cfg.AdminEmail) == "" {
			return nil, errors.New("admin email must be configured")
		}

		// Check for the customer id
		if string(cfg.CustomerID) == "" {
			return nil, errors.New("customer ID must be configured")
		}

		scopes := []string{
			admin.AdminDirectoryUserReadonlyScope,
			admin.AdminDirectoryGroupReadonlyScope,
			admin.AdminDirectoryDeviceMobileReadonlyScope,
			admin.AdminDirectoryDeviceChromeosReadonlyScope,
			admin.AdminDirectoryOrgunitReadonlyScope,
			admin.AdminDirectoryRolemanagementReadonlyScope,
			admin.AdminDirectoryResourceCalendarReadonlyScope,
			admin.AdminDirectoryDomainReadonlyScope,
		}

		// Create credentials using the service account key
		config, err := google.JWTConfigFromJSON([]byte(cfg.KeyFile), scopes...)
		if err != nil {
			return nil, fmt.Errorf("error creating JWT config: %v", err)
		}

		// Set the Subject to the specified admin email
		config.Subject = cfg.AdminEmail

		// Create the Admin SDK service using the credentials
		adminService, err := admin.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx)))
		if err != nil {
			return nil, fmt.Errorf("error creating Admin SDK service: %v", err)
		}

		googleWorkspaceAPIHandler := NewGoogleWorkspaceAPIHandler(adminService, cfg.CustomerID, rate.Every(time.Minute/200), 1, 10, 5, 5*time.Minute)

		// Get values from describer
		var values []models.Resource
		result, err := describe(ctx, googleWorkspaceAPIHandler, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}

// DescribeSingleByGoogleWorkspace A wrapper to pass GoogleWorkspace authorization to describer functions
func DescribeSingleByGoogleWorkspace(describe func(context.Context, *GoogleWorkspaceAPIHandler, string) (*models.Resource, error)) models.SingleResourceDescriber {
	return func(ctx context.Context, cfg models.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, resourceID string, stream *models.StreamSender) (*models.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)

		var err error
		// Check for the keyFile content
		if string(cfg.KeyFile) == "" {
			return nil, errors.New("key file must be configured")
		}

		// Check for the admin email
		if string(cfg.AdminEmail) == "" {
			return nil, errors.New("admin email must be configured")
		}

		// Check for the customer id
		if string(cfg.CustomerID) == "" {
			return nil, errors.New("customer ID must be configured")
		}

		scopes := []string{
			admin.AdminDirectoryUserReadonlyScope,
			admin.AdminDirectoryGroupReadonlyScope,
			admin.AdminDirectoryDeviceMobileReadonlyScope,
			admin.AdminDirectoryOrgunitReadonlyScope,
			admin.AdminDirectoryRolemanagementReadonlyScope,
		}

		// Create credentials using the service account key
		config, err := google.JWTConfigFromJSON([]byte(cfg.KeyFile), scopes...)
		if err != nil {
			return nil, fmt.Errorf("error creating JWT config: %v", err)
		}

		// Set the Subject to the specified admin email
		config.Subject = cfg.AdminEmail

		// Create the Admin SDK service using the credentials
		adminService, err := admin.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx)))
		if err != nil {
			return nil, fmt.Errorf("error creating Admin SDK service: %v", err)
		}

		googleWorkspaceAPIHandler := NewGoogleWorkspaceAPIHandler(adminService, cfg.CustomerID, rate.Every(time.Minute/200), 1, 10, 5, 5*time.Minute)
		// Get value from describer
		value, err := describe(ctx, googleWorkspaceAPIHandler, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}
