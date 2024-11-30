package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-googleworkspace/pkg/sdk/models"
	"github.com/opengovern/og-describer-googleworkspace/provider/model"
	"net/http"
	"net/url"
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
			Value: user,
		},
	}
	return &value, nil
}

func processUsers(ctx context.Context, handler *GoogleWorkspaceAPIHandler, openaiChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var users []model.UserDescription
	var userListResponse model.UserListResponse
	var resp *http.Response
	baseURL := "https://admin.googleapis.com/admin/directory/v1/users"
	pageToken := ""

	for {
		params := url.Values{}
		params.Set("pageToken", pageToken)
		params.Set("maxResults", "500")
		finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

		req, err := http.NewRequest("GET", finalURL, nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		requestFunc := func(req *http.Request) (*http.Response, error) {
			var e error
			resp, e = handler.Client.Do(req)
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}
			defer resp.Body.Close()

			if e = json.NewDecoder(resp.Body).Decode(&userListResponse); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			users = append(users, userListResponse.Users...)
			return resp, nil
		}

		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if userListResponse.NextPageToken == "" {
			break
		}
		pageToken = userListResponse.NextPageToken
	}
	for _, user := range users {
		wg.Add(1)
		go func(user model.UserDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   user.Id,
				Name: user.Name.FullName,
				Description: JSONAllFieldsMarshaller{
					Value: user,
				},
			}
			openaiChan <- value
		}(user)
	}
	return nil
}

func processUser(ctx context.Context, handler *GoogleWorkspaceAPIHandler, resourceID string) (*model.UserDescription, error) {
	var user model.UserDescription
	var resp *http.Response
	baseURL := "https://admin.googleapis.com/admin/directory/v1/users/"

	finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, err
	}

	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		resp, e = handler.Client.Do(req)
		if e != nil {
			return nil, fmt.Errorf("request execution failed: %w", e)
		}
		defer resp.Body.Close()

		if e = json.NewDecoder(resp.Body).Decode(&user); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &user, nil
}
