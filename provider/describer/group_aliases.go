package describer

import (
	"context"
	"fmt"
	"github.com/opengovern/og-describer-googleworkspace/pkg/sdk/models"
	"github.com/opengovern/og-describer-googleworkspace/provider/model"
	admin "google.golang.org/api/admin/directory/v1"
	"log"
	"sync"
)

func ListGroupAliases(ctx context.Context, handler *GoogleWorkspaceAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	GoogleWorkspaceChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	groups, err := getGroups(ctx, handler)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(GoogleWorkspaceChan)
		defer close(errorChan)
		for _, group := range groups {
			if err = processGroupAliases(ctx, handler, group.Id, GoogleWorkspaceChan, &wg); err != nil {
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

func processGroupAliases(ctx context.Context, handler *GoogleWorkspaceAPIHandler, GroupID string, GoogleWorkspaceChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var aliasesResp *admin.Aliases

	req := handler.AdminService.Groups.Aliases.List(GroupID)

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
		log.Println(alias)
		wg.Add(1)
		if aliasValue, ok := alias.(admin.GroupAlias); ok {
			go func(aliasValue admin.GroupAlias) {
				defer wg.Done()
				value := models.Resource{
					ID:   aliasValue.Id,
					Name: aliasValue.Alias,
					Description: JSONAllFieldsMarshaller{
						Value: model.GroupAliasDescription{
							GroupAlias: aliasValue,
						},
					},
				}
				GoogleWorkspaceChan <- value
			}(aliasValue)
		} else if GroupAliasValuePtr, ok := alias.(*admin.GroupAlias); ok {
			go func(GroupAliasValuePtr *admin.GroupAlias) {
				defer wg.Done()
				value := models.Resource{
					ID:   GroupAliasValuePtr.Id,
					Name: GroupAliasValuePtr.Alias,
					Description: JSONAllFieldsMarshaller{
						Value: model.GroupAliasDescription{
							GroupAlias: *GroupAliasValuePtr,
						},
					},
				}
				GoogleWorkspaceChan <- value
			}(GroupAliasValuePtr)
		}
	}
	return nil
}
