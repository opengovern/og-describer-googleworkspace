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

func ListResourceCalenders(ctx context.Context, handler *GoogleWorkspaceAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	GoogleWorkspaceChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(GoogleWorkspaceChan)
		defer close(errorChan)
		if err := processResourceCalenders(ctx, handler, GoogleWorkspaceChan, &wg); err != nil {
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

func GetResourceCalender(ctx context.Context, handler *GoogleWorkspaceAPIHandler, resourceID string) (*models.Resource, error) {
	calender, err := processResourceCalender(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   calender.ResourceId,
		Name: calender.ResourceName,
		Description: JSONAllFieldsMarshaller{
			Value: model.CalenderDescription{
				CalendarResource: *calender,
			},
		},
	}
	return &value, nil
}

func processResourceCalenders(ctx context.Context, handler *GoogleWorkspaceAPIHandler, GoogleWorkspaceChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var calenders []*admin.CalendarResource
	var calendersResp *admin.CalendarResources
	pageToken := ""

	for {
		req := handler.Service.Resources.Calendars.List(handler.CustomerID).MaxResults(MaxPageResultsCalenders)
		if pageToken != "" {
			req.PageToken(pageToken)
		}

		requestFunc := func() (*int, error) {
			var e error
			calendersResp, e = req.Do()
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}

			calenders = append(calenders, calendersResp.Items...)
			return &calendersResp.HTTPStatusCode, nil
		}

		err := handler.DoRequest(ctx, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if calendersResp.NextPageToken == "" {
			break
		}
		pageToken = calendersResp.NextPageToken
	}

	for _, calender := range calenders {
		wg.Add(1)
		go func(calender *admin.CalendarResource) {
			defer wg.Done()
			value := models.Resource{
				ID:   calender.ResourceId,
				Name: calender.ResourceName,
				Description: JSONAllFieldsMarshaller{
					Value: model.CalenderDescription{
						CalendarResource: *calender,
					},
				},
			}
			GoogleWorkspaceChan <- value
		}(calender)
	}
	return nil
}

func processResourceCalender(ctx context.Context, handler *GoogleWorkspaceAPIHandler, resourceID string) (*admin.CalendarResource, error) {
	var calender *admin.CalendarResource
	var status *int

	req := handler.Service.Resources.Calendars.Get(handler.CustomerID, resourceID)

	requestFunc := func() (*int, error) {
		var e error
		calender, e = req.Do()
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
	return calender, nil
}
