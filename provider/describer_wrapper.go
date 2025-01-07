package provider

import (
	"errors"
	"fmt"
	"strings"
	"time"
	// "unsafe"

	model "github.com/opengovern/og-describer-googleworkspace/pkg/sdk/models"
	"github.com/opengovern/og-describer-googleworkspace/provider/configs"
	"github.com/opengovern/og-describer-googleworkspace/provider/describer"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"golang.org/x/time/rate"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

// DescribeListByGoogleWorkspace A wrapper to pass GoogleWorkspace authorization to describer functions
func DescribeListByGoogleWorkspace(describe func(context.Context, *describer.GoogleWorkspaceAPIHandler, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

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
			//admin.AdminDirectoryResourceCalendarReadonlyScope,
			admin.AdminDirectoryDomainReadonlyScope,
		}

		// Create credentials using the service account key
		
		key_file_string := string(cfg.KeyFile)
		unescapedString := strings.ReplaceAll(key_file_string, `\n`, "<NEWLINE>")
		unescapedString = strings.ReplaceAll(unescapedString, `\<NEWLINE>`, "\\n")
		unescapedString = strings.ReplaceAll(unescapedString, `<NEWLINE>`, "")
		unescapedString = strings.ReplaceAll(unescapedString, `\"`, `"`)
		unescapedString = strings.Trim(unescapedString, `"`)

		keyFileData := []byte(unescapedString)

		// fmt.Println("keyFileData", keyFileData)
		// fmt.Println("Key file converted to bytes")
		config, err := google.JWTConfigFromJSON(keyFileData, scopes...)
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

		googleWorkspaceAPIHandler := describer.NewGoogleWorkspaceAPIHandler(adminService, cfg.CustomerID, rate.Every(time.Minute/200), 1, 10, 5, 5*time.Minute)

		// Get values from describer
		var values []model.Resource
		result, err := describe(ctx, googleWorkspaceAPIHandler, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}

// DescribeSingleByGoogleWorkspace A wrapper to pass GoogleWorkspace authorization to describer functions
func DescribeSingleByGoogleWorkspace(describe func(context.Context, *describer.GoogleWorkspaceAPIHandler, string) (*model.Resource, error)) model.SingleResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, resourceID string) (*model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

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
		key_file_string := string(cfg.KeyFile)
		unescapedString := strings.ReplaceAll(key_file_string, `\n`, "<NEWLINE>")
		unescapedString = strings.ReplaceAll(unescapedString, `\<NEWLINE>`, "\\n")
		unescapedString = strings.ReplaceAll(unescapedString, `<NEWLINE>`, "")
		unescapedString = strings.ReplaceAll(unescapedString, `\"`, `"`)
		unescapedString = strings.Trim(unescapedString, `"`)

		keyFileData := []byte(unescapedString)
		config, err := google.JWTConfigFromJSON(keyFileData, scopes...)
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

		googleWorkspaceAPIHandler := describer.NewGoogleWorkspaceAPIHandler(adminService, cfg.CustomerID, rate.Every(time.Minute/200), 1, 10, 5, 5*time.Minute)
		// Get value from describer
		value, err := describe(ctx, googleWorkspaceAPIHandler, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}
