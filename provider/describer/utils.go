package describer

import (
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

type GoogleWorkspaceAPIHandler struct {
	Client       *http.Client
	Token        string
	RateLimiter  *rate.Limiter
	Semaphore    chan struct{}
	MaxRetries   int
	RetryBackoff time.Duration
}

func NewGoogleWorkspaceAPIHandler(token string, rateLimit rate.Limit, burst int, maxConcurrency int, maxRetries int, retryBackoff time.Duration) *GoogleWorkspaceAPIHandler {
	return &GoogleWorkspaceAPIHandler{
		Client:       http.DefaultClient,
		Token:        token,
		RateLimiter:  rate.NewLimiter(rateLimit, burst),
		Semaphore:    make(chan struct{}, maxConcurrency),
		MaxRetries:   maxRetries,
		RetryBackoff: retryBackoff,
	}
}
