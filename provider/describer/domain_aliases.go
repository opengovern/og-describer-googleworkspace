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

func ListDomainAliases(ctx context.Context, handler *GoogleWorkspaceAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	GoogleWorkspaceChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(GoogleWorkspaceChan)
		defer close(errorChan)
		if err := processDomainAliases(ctx, handler, GoogleWorkspaceChan, &wg); err != nil {
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

func GetDomainAlias(ctx context.Context, handler *GoogleWorkspaceAPIHandler, resourceID string) (*models.Resource, error) {
	domainAlias, err := processDomainAlias(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   domainAlias.DomainAliasName,
		Name: domainAlias.DomainAliasName,
		Description: JSONAllFieldsMarshaller{
			Value: model.DomainAliasDescription{
				DomainAlias: *domainAlias,
			},
		},
	}
	return &value, nil
}

func processDomainAliases(ctx context.Context, handler *GoogleWorkspaceAPIHandler, GoogleWorkspaceChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var domainAliasesResp *admin.DomainAliases

	req := handler.AdminService.DomainAliases.List(handler.CustomerID)

	requestFunc := func() (*int, error) {
		var e error
		domainAliasesResp, e = req.Do()
		if e != nil {
			return nil, fmt.Errorf("request execution failed: %w", e)
		}

		return &domainAliasesResp.HTTPStatusCode, nil
	}

	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return fmt.Errorf("error during request handling: %w", err)
	}

	for _, domainAlias := range domainAliasesResp.DomainAliases {
		wg.Add(1)
		go func(domainAlias *admin.DomainAlias) {
			defer wg.Done()
			value := models.Resource{
				ID:   domainAlias.DomainAliasName,
				Name: domainAlias.DomainAliasName,
				Description: JSONAllFieldsMarshaller{
					Value: model.DomainAliasDescription{
						DomainAlias: *domainAlias,
					},
				},
			}
			GoogleWorkspaceChan <- value
		}(domainAlias)
	}
	return nil
}

func processDomainAlias(ctx context.Context, handler *GoogleWorkspaceAPIHandler, resourceID string) (*admin.DomainAlias, error) {
	var domainAlias *admin.DomainAlias
	var status *int

	req := handler.AdminService.DomainAliases.Get(handler.CustomerID, resourceID)

	requestFunc := func() (*int, error) {
		var e error
		domainAlias, e = req.Do()
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
	return domainAlias, nil
}
