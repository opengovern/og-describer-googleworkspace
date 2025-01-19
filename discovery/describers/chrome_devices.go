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

func ListChromeDevices(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	GoogleWorkspaceChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(GoogleWorkspaceChan)
		defer close(errorChan)
		if err := processChromeDevices(ctx, handler, GoogleWorkspaceChan, &wg); err != nil {
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

func GetChromeDevice(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, resourceID string) (*models.Resource, error) {
	chromeDevice, err := processChromeDevice(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   chromeDevice.DeviceId,
		Name: chromeDevice.DeviceId,
		Description: provider.ChromeDeviceDescription{
			ChromeOsDevice: *chromeDevice,
		},
	}
	return &value, nil
}

func processChromeDevices(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, GoogleWorkspaceChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var chromeDevices []*admin.ChromeOsDevice
	var chromeDevicesResp *admin.ChromeOsDevices
	pageToken := ""

	for {
		req := handler.AdminService.Chromeosdevices.List(handler.CustomerID).MaxResults(provider.MaxPageResultsChromeDevices)
		if pageToken != "" {
			req.PageToken(pageToken)
		}

		requestFunc := func() (*int, error) {
			var e error
			chromeDevicesResp, e = req.Do()
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}

			chromeDevices = append(chromeDevices, chromeDevicesResp.Chromeosdevices...)
			return &chromeDevicesResp.HTTPStatusCode, nil
		}

		err := handler.DoRequest(ctx, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if chromeDevicesResp.NextPageToken == "" {
			break
		}
		pageToken = chromeDevicesResp.NextPageToken
	}

	for _, chromeDevice := range chromeDevices {
		wg.Add(1)
		go func(chromeDevice *admin.ChromeOsDevice) {
			defer wg.Done()
			value := models.Resource{
				ID:   chromeDevice.DeviceId,
				Name: chromeDevice.DeviceId,
				Description: provider.ChromeDeviceDescription{
					ChromeOsDevice: *chromeDevice,
				},
			}
			GoogleWorkspaceChan <- value
		}(chromeDevice)
	}
	return nil
}

func processChromeDevice(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, resourceID string) (*admin.ChromeOsDevice, error) {
	var chromeDevice *admin.ChromeOsDevice
	var status *int

	req := handler.AdminService.Chromeosdevices.Get(handler.CustomerID, resourceID)

	requestFunc := func() (*int, error) {
		var e error
		chromeDevice, e = req.Do()
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
	return chromeDevice, nil
}
