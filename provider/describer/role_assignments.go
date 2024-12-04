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
	"strconv"
	"sync"
)

func ListRoleAssignments(ctx context.Context, handler *GoogleWorkspaceAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	GoogleWorkspaceChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(GoogleWorkspaceChan)
		defer close(errorChan)
		if err := processRoleAssignments(ctx, handler, GoogleWorkspaceChan, &wg); err != nil {
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

func GetRoleAssignment(ctx context.Context, handler *GoogleWorkspaceAPIHandler, resourceID string) (*models.Resource, error) {
	roleAssignment, err := processRoleAssignment(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   strconv.Itoa(int(roleAssignment.RoleAssignmentId)),
		Name: strconv.Itoa(int(roleAssignment.RoleAssignmentId)),
		Description: JSONAllFieldsMarshaller{
			Value: model.RoleAssignmentDescription{
				RoleAssignment: *roleAssignment,
			},
		},
	}
	return &value, nil
}

func processRoleAssignments(ctx context.Context, handler *GoogleWorkspaceAPIHandler, GoogleWorkspaceChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var roleAssignments []*admin.RoleAssignment
	var roleAssignmentsResp *admin.RoleAssignments
	pageToken := ""

	for {
		req := handler.Service.RoleAssignments.List(handler.CustomerID).MaxResults(MaxPageResultsRoleAssignments)
		if pageToken != "" {
			req.PageToken(pageToken)
		}

		requestFunc := func() (*int, error) {
			var e error
			roleAssignmentsResp, e = req.Do()
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}

			roleAssignments = append(roleAssignments, roleAssignmentsResp.Items...)
			return &roleAssignmentsResp.HTTPStatusCode, nil
		}

		err := handler.DoRequest(ctx, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if roleAssignmentsResp.NextPageToken == "" {
			break
		}
		pageToken = roleAssignmentsResp.NextPageToken
	}

	for _, roleAssignment := range roleAssignments {
		wg.Add(1)
		go func(roleAssignment *admin.RoleAssignment) {
			defer wg.Done()
			value := models.Resource{
				ID:   strconv.Itoa(int(roleAssignment.RoleAssignmentId)),
				Name: strconv.Itoa(int(roleAssignment.RoleAssignmentId)),
				Description: JSONAllFieldsMarshaller{
					Value: model.RoleAssignmentDescription{
						RoleAssignment: *roleAssignment,
					},
				},
			}
			GoogleWorkspaceChan <- value
		}(roleAssignment)
	}
	return nil
}

func processRoleAssignment(ctx context.Context, handler *GoogleWorkspaceAPIHandler, resourceID string) (*admin.RoleAssignment, error) {
	var roleAssignment *admin.RoleAssignment
	var status *int

	req := handler.Service.RoleAssignments.Get(handler.CustomerID, resourceID)

	requestFunc := func() (*int, error) {
		var e error
		roleAssignment, e = req.Do()
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
	return roleAssignment, nil
}
