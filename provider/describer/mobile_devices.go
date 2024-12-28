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

func ListMobileDevices(ctx context.Context, handler *GoogleWorkspaceAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	GoogleWorkspaceChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(GoogleWorkspaceChan)
		defer close(errorChan)
		if err := processMobileDevices(ctx, handler, GoogleWorkspaceChan, &wg); err != nil {
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

func GetMobileDevice(ctx context.Context, handler *GoogleWorkspaceAPIHandler, resourceID string) (*models.Resource, error) {
	mobileDevice, err := processMobileDevice(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   mobileDevice.DeviceId,
		Name: mobileDevice.DeviceId,
		Description: model.MobileDeviceDescription{
			MobileDevice: *mobileDevice,
		},
	}
	return &value, nil
}

func processMobileDevices(ctx context.Context, handler *GoogleWorkspaceAPIHandler, GoogleWorkspaceChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var mobileDevices []*admin.MobileDevice
	var mobileDevicesResp *admin.MobileDevices
	pageToken := ""

	for {
		req := handler.AdminService.Mobiledevices.List(handler.CustomerID).MaxResults(MaxPageResultsMobileDevices)
		if pageToken != "" {
			req.PageToken(pageToken)
		}

		requestFunc := func() (*int, error) {
			var e error
			mobileDevicesResp, e = req.Do()
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}

			mobileDevices = append(mobileDevices, mobileDevicesResp.Mobiledevices...)
			return &mobileDevicesResp.HTTPStatusCode, nil
		}

		err := handler.DoRequest(ctx, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if mobileDevicesResp.NextPageToken == "" {
			break
		}
		pageToken = mobileDevicesResp.NextPageToken
	}

	for _, mobileDevice := range mobileDevices {
		wg.Add(1)
		go func(mobileDevice *admin.MobileDevice) {
			defer wg.Done()
			value := models.Resource{
				ID:   mobileDevice.DeviceId,
				Name: mobileDevice.DeviceId,
				Description: model.MobileDeviceDescription{
					MobileDevice: *mobileDevice,
				},
			}
			GoogleWorkspaceChan <- value
		}(mobileDevice)
	}
	return nil
}

func processMobileDevice(ctx context.Context, handler *GoogleWorkspaceAPIHandler, resourceID string) (*admin.MobileDevice, error) {
	var mobileDevice *admin.MobileDevice
	var status *int

	req := handler.AdminService.Mobiledevices.Get(handler.CustomerID, resourceID)

	requestFunc := func() (*int, error) {
		var e error
		mobileDevice, e = req.Do()
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
	return mobileDevice, nil
}
