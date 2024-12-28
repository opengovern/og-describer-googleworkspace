package describer

import (
	"context"
	"errors"
	"fmt"
	"github.com/opengovern/og-describer-googleworkspace/pkg/sdk/models"
	"github.com/opengovern/og-describer-googleworkspace/provider/model"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
	"net/http"
	"sync"
)

func ListResourceFeatures(ctx context.Context, handler *GoogleWorkspaceAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	GoogleWorkspaceChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(GoogleWorkspaceChan)
		defer close(errorChan)
		if err := processResourceFeatures(ctx, handler, GoogleWorkspaceChan, &wg); err != nil {
			errorChan <- err // Send error to the error channel
		}
		wg.Wait()
	}()

	var values []models.Resource
	for {
		select {
		case value, ok := <-GoogleWorkspaceChan:
			if !ok {
				return values, nil
			}
			if stream != nil {
				if err := (*stream)(value); err != nil {
					return nil, err
				}
			} else {
				values = append(values, value)
			}
		case err := <-errorChan:
			return nil, err
		}
	}
}

func GetResourceFeature(ctx context.Context, handler *GoogleWorkspaceAPIHandler, resourceID string) (*models.Resource, error) {
	feature, err := processResourceFeature(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   feature.Name,
		Name: feature.Name,
		Description: model.ResourceFeatureDescription{
			Feature: *feature,
		},
	}
	return &value, nil
}

func processResourceFeatures(ctx context.Context, handler *GoogleWorkspaceAPIHandler, GoogleWorkspaceChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var features []*admin.Feature
	var featuresResp *admin.Features
	pageToken := ""

	for {
		req := handler.AdminService.Resources.Features.List(handler.CustomerID).MaxResults(MaxPageResultsFeatures)
		if pageToken != "" {
			req.PageToken(pageToken)
		}

		requestFunc := func() (*int, error) {
			var e error
			featuresResp, e = req.Do()
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}

			features = append(features, featuresResp.Features...)
			return &featuresResp.HTTPStatusCode, nil
		}

		err := handler.DoRequest(ctx, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if featuresResp.NextPageToken == "" {
			break
		}
		pageToken = featuresResp.NextPageToken
	}

	for _, feature := range features {
		wg.Add(1)
		go func(feature *admin.Feature) {
			defer wg.Done()
			value := models.Resource{
				ID:   feature.Name,
				Name: feature.Name,
				Description: model.ResourceFeatureDescription{
					Feature: *feature,
				},
			}
			GoogleWorkspaceChan <- value
		}(feature)
	}
	return nil
}

func processResourceFeature(ctx context.Context, handler *GoogleWorkspaceAPIHandler, resourceID string) (*admin.Feature, error) {
	var feature *admin.Feature
	var status *int

	req := handler.AdminService.Resources.Features.Get(handler.CustomerID, resourceID)

	requestFunc := func() (*int, error) {
		var e error
		feature, e = req.Do()
		if e != nil {
			var apiErr *googleapi.Error
			if errors.As(e, &apiErr) {
				*status = apiErr.Code
			}
			return status, fmt.Errorf("request execution failed: %w", e)
		}
		*status = http.StatusOK
		return status, e
	}

	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return feature, nil
}
