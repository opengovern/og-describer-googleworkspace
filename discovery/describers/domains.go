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

func ListDomains(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	GoogleWorkspaceChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(GoogleWorkspaceChan)
		defer close(errorChan)
		if err := processDomains(ctx, handler, GoogleWorkspaceChan, &wg); err != nil {
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

func GetDomain(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, resourceID string) (*models.Resource, error) {
	domain, err := processDomain(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   domain.DomainName,
		Name: domain.DomainName,
		Description: provider.DomainDescription{
			Domains: *domain,
		},
	}
	return &value, nil
}

func processDomains(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, GoogleWorkspaceChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var domainsResp *admin.Domains2

	req := handler.AdminService.Domains.List(handler.CustomerID)

	requestFunc := func() (*int, error) {
		var e error
		domainsResp, e = req.Do()
		if e != nil {
			return nil, fmt.Errorf("request execution failed: %w", e)
		}

		return &domainsResp.HTTPStatusCode, nil
	}

	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return fmt.Errorf("error during request handling: %w", err)
	}

	for _, domain := range domainsResp.Domains {
		wg.Add(1)
		go func(domain *admin.Domains) {
			defer wg.Done()
			value := models.Resource{
				ID:   domain.DomainName,
				Name: domain.DomainName,
				Description: provider.DomainDescription{
					Domains: *domain,
				},
			}
			GoogleWorkspaceChan <- value
		}(domain)
	}
	return nil
}

func processDomain(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, resourceID string) (*admin.Domains, error) {
	var domain *admin.Domains
	var status *int

	req := handler.AdminService.Domains.Get(handler.CustomerID, resourceID)

	requestFunc := func() (*int, error) {
		var e error
		domain, e = req.Do()
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
	return domain, nil
}
