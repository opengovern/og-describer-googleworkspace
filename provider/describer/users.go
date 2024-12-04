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

func ListUsers(ctx context.Context, handler *GoogleWorkspaceAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	GoogleWorkspaceChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(GoogleWorkspaceChan)
		defer close(errorChan)
		if err := processUsers(ctx, handler, GoogleWorkspaceChan, &wg); err != nil {
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

func GetUser(ctx context.Context, handler *GoogleWorkspaceAPIHandler, resourceID string) (*models.Resource, error) {
	user, err := processUser(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   user.Id,
		Name: user.Name.FullName,
		Description: JSONAllFieldsMarshaller{
			Value: model.UserDescription{
				User: *user,
			},
		},
	}
	return &value, nil
}

func processUsers(ctx context.Context, handler *GoogleWorkspaceAPIHandler, GoogleWorkspaceChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var users []*admin.User
	var usersResp *admin.Users
	pageToken := ""

	for {
		req := handler.AdminService.Users.List().Customer(handler.CustomerID).MaxResults(MaxPageResultsUsers)
		if pageToken != "" {
			req.PageToken(pageToken)
		}

		requestFunc := func() (*int, error) {
			var e error
			usersResp, e = req.Do()
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}

			users = append(users, usersResp.Users...)
			return &usersResp.HTTPStatusCode, nil
		}

		err := handler.DoRequest(ctx, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if usersResp.NextPageToken == "" {
			break
		}
		pageToken = usersResp.NextPageToken
	}

	for _, user := range users {
		wg.Add(1)
		go func(user *admin.User) {
			defer wg.Done()
			value := models.Resource{
				ID:   user.Id,
				Name: user.Name.FullName,
				Description: JSONAllFieldsMarshaller{
					Value: model.UserDescription{
						User: *user,
					},
				},
			}
			GoogleWorkspaceChan <- value
		}(user)
	}
	return nil
}

func processUser(ctx context.Context, handler *GoogleWorkspaceAPIHandler, resourceID string) (*admin.User, error) {
	var user *admin.User
	var status *int

	req := handler.AdminService.Users.Get(resourceID)

	requestFunc := func() (*int, error) {
		var e error
		user, e = req.Do()
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
	return user, nil
}
