package describer

import (
	"context"
	"fmt"
	"github.com/opengovern/og-describer-googleworkspace/pkg/sdk/models"
	"github.com/opengovern/og-describer-googleworkspace/provider/model"
	admin "google.golang.org/api/admin/directory/v1"
	"sync"
)

func ListUserAliases(ctx context.Context, handler *GoogleWorkspaceAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	GoogleWorkspaceChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	users, err := getUsers(ctx, handler)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(GoogleWorkspaceChan)
		defer close(errorChan)
		for _, user := range users {
			if err = processUserAliases(ctx, handler, user.Id, GoogleWorkspaceChan, &wg); err != nil {
				errorChan <- err // Send error to the error channel
			}
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
				if err = (*stream)(value); err != nil {
					return nil, err
				}
			} else {
				values = append(values, value)
			}
		case err = <-errorChan:
			return nil, err
		}
	}
}

func processUserAliases(ctx context.Context, handler *GoogleWorkspaceAPIHandler, userID string, GoogleWorkspaceChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var aliasesResp *admin.Aliases

	req := handler.AdminService.Users.Aliases.List(userID)

	requestFunc := func() (*int, error) {
		var e error
		aliasesResp, e = req.Do()
		if e != nil {
			return nil, fmt.Errorf("request execution failed: %w", e)
		}
		return &aliasesResp.HTTPStatusCode, nil
	}

	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return fmt.Errorf("error during request handling: %w", err)
	}
	for _, alias := range aliasesResp.Aliases {
		wg.Add(1)
		if aliasValue, ok := alias.(map[string]interface{}); ok {
			go func(aliasValue map[string]interface{}) {
				defer wg.Done()
				id, _ := aliasValue["id"].(string)
				aliasName, _ := aliasValue["alias"].(string)
				etag, _ := aliasValue["etag"].(string)
				kind, _ := aliasValue["kind"].(string)
				primaryEmail, _ := aliasValue["primaryEmail"].(string)
				value := models.Resource{
					ID:   id,
					Name: aliasName,
					Description: model.UserAliasDescription{
						UserAlias: admin.UserAlias{
							Alias:        aliasName,
							Etag:         etag,
							Id:           id,
							Kind:         kind,
							PrimaryEmail: primaryEmail,
						},
					},
				}
				GoogleWorkspaceChan <- value
			}(aliasValue)
		}
	}
	return nil
}
