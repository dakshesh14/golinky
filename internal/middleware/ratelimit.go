package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/httprate"
)

func RateLimit(requestsPerMinute int, windowLength time.Duration) func(http.Handler) http.Handler {
	return httprate.LimitByIP(requestsPerMinute, windowLength)
}

func GlobalRateLimit(requestsPerMinute int) func(http.Handler) http.Handler {
	return RateLimit(requestsPerMinute, time.Minute)
}
