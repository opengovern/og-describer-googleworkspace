package provider

import (
	"errors"
	model "github.com/opengovern/og-describer-googleworkspace/pkg/sdk/models"
	"github.com/opengovern/og-describer-googleworkspace/provider/configs"
	"github.com/opengovern/og-describer-googleworkspace/provider/describer"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"time"
)

// DescribeListByGoogleWorkspace A wrapper to pass GoogleWorkspace authorization to describer functions
func DescribeListByGoogleWorkspace(describe func(context.Context, *describer.GoogleWorkspaceAPIHandler, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

		var err error
		// Check for the token
		if cfg.Token == "" {
			return nil, errors.New("token must be configured")
		}

		googleWorkspaceAPIHandler := describer.NewGoogleWorkspaceAPIHandler(cfg.Token, rate.Every(time.Minute/200), 1, 10, 5, 5*time.Minute)

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
		// Check for the token
		if cfg.Token == "" {
			return nil, errors.New("token must be configured")
		}

		googleWorkspaceAPIHandler := describer.NewGoogleWorkspaceAPIHandler(cfg.Token, rate.Every(time.Minute/200), 1, 10, 5, 5*time.Minute)

		// Get value from describer
		value, err := describe(ctx, googleWorkspaceAPIHandler, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}
