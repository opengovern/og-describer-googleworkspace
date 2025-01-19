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

func ListOrgUnits(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	GoogleWorkspaceChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(GoogleWorkspaceChan)
		defer close(errorChan)
		if err := processOrgUnits(ctx, handler, GoogleWorkspaceChan, &wg); err != nil {
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

func GetOrgUnit(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, resourceID string) (*models.Resource, error) {
	orgUnit, err := processOrgUnit(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   orgUnit.OrgUnitId,
		Name: orgUnit.Name,
		Description: provider.OrgUnitDescription{
			OrgUnit: *orgUnit,
		},
	}
	return &value, nil
}

func processOrgUnits(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, GoogleWorkspaceChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var orgUnits []*admin.OrgUnit
	var orgUnitsResp *admin.OrgUnits
	req := handler.AdminService.Orgunits.List(handler.CustomerID)

	requestFunc := func() (*int, error) {
		var e error
		orgUnitsResp, e = req.Do()
		if e != nil {
			return nil, fmt.Errorf("request execution failed: %w", e)
		}

		orgUnits = append(orgUnits, orgUnitsResp.OrganizationUnits...)
		return &orgUnitsResp.HTTPStatusCode, nil
	}

	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return fmt.Errorf("error during request handling: %w", err)
	}

	for _, orgUnit := range orgUnits {
		wg.Add(1)
		go func(orgUnit *admin.OrgUnit) {
			defer wg.Done()
			value := models.Resource{
				ID:   orgUnit.OrgUnitId,
				Name: orgUnit.Name,
				Description: provider.OrgUnitDescription{
					OrgUnit: *orgUnit,
				},
			}
			GoogleWorkspaceChan <- value
		}(orgUnit)
	}
	return nil
}

func processOrgUnit(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, resourceID string) (*admin.OrgUnit, error) {
	var orgUnit *admin.OrgUnit
	var status *int

	req := handler.AdminService.Orgunits.Get(handler.CustomerID, resourceID)

	requestFunc := func() (*int, error) {
		var e error
		orgUnit, e = req.Do()
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
	return orgUnit, nil
}
