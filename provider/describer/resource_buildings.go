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

func ListResourceBuildings(ctx context.Context, handler *GoogleWorkspaceAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	GoogleWorkspaceChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(GoogleWorkspaceChan)
		defer close(errorChan)
		if err := processResourceBuildings(ctx, handler, GoogleWorkspaceChan, &wg); err != nil {
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

func GetResourceBuilding(ctx context.Context, handler *GoogleWorkspaceAPIHandler, resourceID string) (*models.Resource, error) {
	building, err := processResourceBuilding(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   building.BuildingId,
		Name: building.BuildingName,
		Description: model.ResourceBuildingDescription{
			Building: *building,
		},
	}
	return &value, nil
}

func processResourceBuildings(ctx context.Context, handler *GoogleWorkspaceAPIHandler, GoogleWorkspaceChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var buildings []*admin.Building
	var buildingsResp *admin.Buildings
	pageToken := ""

	for {
		req := handler.AdminService.Resources.Buildings.List(handler.CustomerID).MaxResults(MaxPageResultsBuildings)
		if pageToken != "" {
			req.PageToken(pageToken)
		}

		requestFunc := func() (*int, error) {
			var e error
			buildingsResp, e = req.Do()
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}

			buildings = append(buildings, buildingsResp.Buildings...)
			return &buildingsResp.HTTPStatusCode, nil
		}

		err := handler.DoRequest(ctx, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if buildingsResp.NextPageToken == "" {
			break
		}
		pageToken = buildingsResp.NextPageToken
	}

	for _, building := range buildings {
		wg.Add(1)
		go func(building *admin.Building) {
			defer wg.Done()
			value := models.Resource{
				ID:   building.BuildingId,
				Name: building.BuildingName,
				Description: model.ResourceBuildingDescription{
					Building: *building,
				},
			}
			GoogleWorkspaceChan <- value
		}(building)
	}
	return nil
}

func processResourceBuilding(ctx context.Context, handler *GoogleWorkspaceAPIHandler, resourceID string) (*admin.Building, error) {
	var building *admin.Building
	var status *int

	req := handler.AdminService.Resources.Buildings.Get(handler.CustomerID, resourceID)

	requestFunc := func() (*int, error) {
		var e error
		building, e = req.Do()
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
	return building, nil
}
