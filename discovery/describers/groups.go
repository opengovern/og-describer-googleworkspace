package describers

import (
	"context"
	"errors"
	"fmt"
	"github.com/opengovern/og-describer-googleworkspace/discovery/pkg/models"
	"github.com/opengovern/og-describer-googleworkspace/discovery/provider"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
	"net/http"
	"sync"
)

func ListGroups(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	GoogleWorkspaceChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(GoogleWorkspaceChan)
		defer close(errorChan)
		if err := processGroups(ctx, handler, GoogleWorkspaceChan, &wg); err != nil {
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

func GetGroup(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, resourceID string) (*models.Resource, error) {
	group, err := processGroup(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   group.Id,
		Name: group.Name,
		Description: provider.GroupDescription{
			Group: *group,
		},
	}
	return &value, nil
}

func processGroups(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, GoogleWorkspaceChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var groups []*admin.Group
	var groupsResp *admin.Groups
	pageToken := ""

	for {
		req := handler.AdminService.Groups.List().Customer(handler.CustomerID).MaxResults(provider.MaxPageResultsGroups)
		if pageToken != "" {
			req.PageToken(pageToken)
		}

		requestFunc := func() (*int, error) {
			var e error
			groupsResp, e = req.Do()
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}

			groups = append(groups, groupsResp.Groups...)
			return &groupsResp.HTTPStatusCode, nil
		}

		err := handler.DoRequest(ctx, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if groupsResp.NextPageToken == "" {
			break
		}
		pageToken = groupsResp.NextPageToken
	}

	for _, group := range groups {
		wg.Add(1)
		go func(group *admin.Group) {
			defer wg.Done()
			value := models.Resource{
				ID:   group.Id,
				Name: group.Name,
				Description: provider.GroupDescription{
					Group: *group,
				},
			}
			GoogleWorkspaceChan <- value
		}(group)
	}
	return nil
}

func processGroup(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, resourceID string) (*admin.Group, error) {
	var group *admin.Group
	var status *int

	req := handler.AdminService.Groups.Get(resourceID)

	requestFunc := func() (*int, error) {
		var e error
		group, e = req.Do()
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
	return group, nil
}
