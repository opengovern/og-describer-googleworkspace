package describer

import (
	"context"
	"fmt"
	"github.com/opengovern/og-describer-googleworkspace/pkg/sdk/models"
	"github.com/opengovern/og-describer-googleworkspace/provider/model"
	admin "google.golang.org/api/admin/directory/v1"
	"sync"
)

func ListPrivileges(ctx context.Context, handler *GoogleWorkspaceAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	GoogleWorkspaceChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(GoogleWorkspaceChan)
		defer close(errorChan)
		if err := processPrivileges(ctx, handler, GoogleWorkspaceChan, &wg); err != nil {
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

func processPrivileges(ctx context.Context, handler *GoogleWorkspaceAPIHandler, GoogleWorkspaceChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var privilegesResp *admin.Privileges

	req := handler.Service.Privileges.List(handler.CustomerID)

	requestFunc := func() (*int, error) {
		var e error
		privilegesResp, e = req.Do()
		if e != nil {
			return nil, fmt.Errorf("request execution failed: %w", e)
		}

		return &privilegesResp.HTTPStatusCode, nil
	}

	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return fmt.Errorf("error during request handling: %w", err)
	}

	for _, privilege := range privilegesResp.Items {
		wg.Add(1)
		go func(privilege *admin.Privilege) {
			defer wg.Done()
			value := models.Resource{
				ID:   privilege.PrivilegeName,
				Name: privilege.PrivilegeName,
				Description: JSONAllFieldsMarshaller{
					Value: model.PrivilegeDescription{
						Privilege: *privilege,
					},
				},
			}
			GoogleWorkspaceChan <- value
		}(privilege)
	}
	return nil
}
