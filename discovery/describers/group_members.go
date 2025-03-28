package describers

import (
	"context"
	"fmt"
	"github.com/opengovern/og-describer-googleworkspace/discovery/pkg/models"
	"github.com/opengovern/og-describer-googleworkspace/discovery/provider"
	admin "google.golang.org/api/admin/directory/v1"
	"sync"
)

func ListGroupMembers(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	GoogleWorkspaceChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	groups, err := provider.GetGroups(ctx, handler)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(GoogleWorkspaceChan)
		defer close(errorChan)
		for _, group := range groups {
			if err = processGroupMembers(ctx, handler, group.Id, GoogleWorkspaceChan, &wg); err != nil {
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

func processGroupMembers(ctx context.Context, handler *provider.GoogleWorkspaceAPIHandler, groupID string, GoogleWorkspaceChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var groupMembers []*admin.Member
	var groupMembersResp *admin.Members
	pageToken := ""

	for {
		req := handler.AdminService.Members.List(groupID).MaxResults(provider.MaxPageResultsGroupMembers)
		if pageToken != "" {
			req.PageToken(pageToken)
		}

		requestFunc := func() (*int, error) {
			var e error
			groupMembersResp, e = req.Do()
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}

			groupMembers = append(groupMembers, groupMembersResp.Members...)
			return &groupMembersResp.HTTPStatusCode, nil
		}

		err := handler.DoRequest(ctx, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if groupMembersResp.NextPageToken == "" {
			break
		}
		pageToken = groupMembersResp.NextPageToken
	}

	for _, groupMember := range groupMembers {
		wg.Add(1)
		go func(member *admin.Member) {
			defer wg.Done()
			value := models.Resource{
				ID:   member.Id,
				Name: member.Email,
				Description: provider.GroupMemberDescription{
					Member: *member,
				},
			}
			GoogleWorkspaceChan <- value
		}(groupMember)
	}
	return nil
}
