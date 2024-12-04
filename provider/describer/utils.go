package describer

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/time/rate"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/gmail/v1"
	"net/http"
	"time"
)

const (
	MaxPageResultsUsers           = 500
	MaxPageResultsGroups          = 200
	MaxPageResultsGroupMembers    = 200
	MaxPageResultsMobileDevices   = 100
	MaxPageResultsChromeDevices   = 300
	MaxPageResultsRoles           = 100
	MaxPageResultsRoleAssignments = 100
	MaxPageResultsBuildings       = 100
	MaxPageResultsCalenders       = 100
	MaxPageResultsFeatures        = 100
)

type GoogleWorkspaceAPIHandler struct {
	AdminService *admin.Service
	GmailService *gmail.Service
	CustomerID   string
	RateLimiter  *rate.Limiter
	Semaphore    chan struct{}
	MaxRetries   int
	RetryBackoff time.Duration
}

func NewGoogleWorkspaceAPIHandler(adminService *admin.Service, gmailService *gmail.Service, customerID string, rateLimit rate.Limit, burst int, maxConcurrency int, maxRetries int, retryBackoff time.Duration) *GoogleWorkspaceAPIHandler {
	return &GoogleWorkspaceAPIHandler{
		AdminService: adminService,
		GmailService: gmailService,
		CustomerID:   customerID,
		RateLimiter:  rate.NewLimiter(rateLimit, burst),
		Semaphore:    make(chan struct{}, maxConcurrency),
		MaxRetries:   maxRetries,
		RetryBackoff: retryBackoff,
	}
}

func getUsers(ctx context.Context, handler *GoogleWorkspaceAPIHandler) ([]*admin.User, error) {
	var users []*admin.User
	var usersResp *admin.Users
	pageToken := ""

	for {
		req := handler.AdminService.Users.List().Customer(handler.CustomerID).MaxResults(MaxPageResultsUsers)
		if pageToken != "" {
			req.PageToken(pageToken)
		}

		requestFunc := func() (*int, error) {
			var e error
			usersResp, e = req.Do()
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}

			users = append(users, usersResp.Users...)
			return &usersResp.HTTPStatusCode, nil
		}

		err := handler.DoRequest(ctx, requestFunc)
		if err != nil {
			return nil, fmt.Errorf("error during request handling: %w", err)
		}

		if usersResp.NextPageToken == "" {
			break
		}
		pageToken = usersResp.NextPageToken
	}

	return users, nil
}

func getGroups(ctx context.Context, handler *GoogleWorkspaceAPIHandler) ([]*admin.Group, error) {
	var groups []*admin.Group
	var groupsResp *admin.Groups
	pageToken := ""

	for {
		req := handler.AdminService.Groups.List().Customer(handler.CustomerID).MaxResults(MaxPageResultsGroups)
		if pageToken != "" {
			req.PageToken(pageToken)
		}

		requestFunc := func() (*int, error) {
			var e error
			groupsResp, e = req.Do()
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}

			groups = append(groups, groupsResp.Groups...)
			return &groupsResp.HTTPStatusCode, nil
		}

		err := handler.DoRequest(ctx, requestFunc)
		if err != nil {
			return nil, fmt.Errorf("error during request handling: %w", err)
		}

		if groupsResp.NextPageToken == "" {
			break
		}
		pageToken = groupsResp.NextPageToken
	}

	return groups, nil
}

// DoRequest executes the googleWorkspace API request with rate limiting, retries, and concurrency control.
func (h *GoogleWorkspaceAPIHandler) DoRequest(ctx context.Context, requestFunc func() (*int, error)) error {
	h.Semaphore <- struct{}{}
	defer func() { <-h.Semaphore }()
	var status *int
	var err error
	for attempt := 0; attempt <= h.MaxRetries; attempt++ {
		// Wait based on rate limiter
		if err = h.RateLimiter.Wait(ctx); err != nil {
			return err
		}
		// Execute the request function
		status, err = requestFunc()
		if err == nil {
			return nil
		}
		// Handle rate limit errors
		if status != nil && *status == http.StatusTooManyRequests {
			// Exponential backoff if headers are missing
			backoff := h.RetryBackoff * (1 << attempt)
			time.Sleep(backoff)
			continue
		}
		// Handle temporary network errors
		if isTemporary(err) {
			backoff := h.RetryBackoff * (1 << attempt)
			time.Sleep(backoff)
			continue
		}
		break
	}
	return err
}

// isTemporary checks if an error is temporary.
func isTemporary(err error) bool {
	if err == nil {
		return false
	}
	var netErr interface{ Temporary() bool }
	if errors.As(err, &netErr) {
		return netErr.Temporary()
	}
	return false
}
