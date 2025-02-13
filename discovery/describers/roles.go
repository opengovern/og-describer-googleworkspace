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
	"strconv"
	"sync"
)

func ListRoles(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	GoogleWorkspaceChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(GoogleWorkspaceChan)
		defer close(errorChan)
		if err := processRoles(ctx, handler, GoogleWorkspaceChan, &wg); err != nil {
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

func GetRole(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, resourceID string) (*models.Resource, error) {
	role, err := processRole(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(int(role.RoleId)),
		Name: role.RoleName,
		Description: provider.RoleDescription{
			Role: *role,
		},
	}
	return &value, nil
}

func processRoles(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, GoogleWorkspaceChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var roles []*admin.Role
	var rolesResp *admin.Roles
	pageToken := ""

	for {
		req := handler.AdminService.Roles.List(handler.CustomerID).MaxResults(provider.MaxPageResultsRoles)
		if pageToken != "" {
			req.PageToken(pageToken)
		}

		requestFunc := func() (*int, error) {
			var e error
			rolesResp, e = req.Do()
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}

			roles = append(roles, rolesResp.Items...)
			return &rolesResp.HTTPStatusCode, nil
		}

		err := handler.DoRequest(ctx, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if rolesResp.NextPageToken == "" {
			break
		}
		pageToken = rolesResp.NextPageToken
	}

	for _, role := range roles {
		wg.Add(1)
		go func(role *admin.Role) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(int(role.RoleId)),
				Name: role.RoleName,
				Description: provider.RoleDescription{
					Role: *role,
				},
			}
			GoogleWorkspaceChan <- value
		}(role)
	}
	return nil
}

func processRole(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, resourceID string) (*admin.Role, error) {
	var role *admin.Role
	var status *int

	req := handler.AdminService.Roles.Get(handler.CustomerID, resourceID)

	requestFunc := func() (*int, error) {
		var e error
		role, e = req.Do()
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
	return role, nil
}
