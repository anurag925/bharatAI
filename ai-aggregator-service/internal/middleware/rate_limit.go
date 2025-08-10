package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

// RateLimiter provides rate limiting functionality
type RateLimiter struct {
	visitors map[string]*visitor
	mu       sync.RWMutex
	rate     int // requests per second
	burst    int // maximum burst size
}

type visitor struct {
	lastSeen time.Time
	count    int
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(rate, burst int) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*visitor),
		rate:     rate,
		burst:    burst,
	}

	// Clean up old entries every minute
	go rl.cleanupVisitors()
	return rl
}

// RateLimitMiddleware returns an Echo middleware for rate limiting
func (rl *RateLimiter) RateLimitMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get client identifier (IP address)
			clientIP := c.RealIP()
			if clientIP == "" {
				clientIP = c.Request().RemoteAddr
			}

			// Check rate limit
			if !rl.allow(clientIP) {
				return c.JSON(http.StatusTooManyRequests, map[string]string{
					"error": "Rate limit exceeded",
				})
			}

			return next(c)
		}
	}
}

// allow checks if a request from the given client is allowed
func (rl *RateLimiter) allow(clientIP string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[clientIP]
	if !exists {
		rl.visitors[clientIP] = &visitor{
			lastSeen: time.Now(),
			count:    1,
		}
		return true
	}

	// Reset counter if window has passed
	if time.Since(v.lastSeen) > time.Second {
		v.count = 0
		v.lastSeen = time.Now()
	}

	// Check if under rate limit
	if v.count >= rl.rate {
		return false
	}

	v.count++
	return true
}

// cleanupVisitors removes old visitor entries
func (rl *RateLimiter) cleanupVisitors() {
	for {
		time.Sleep(time.Minute)
		rl.mu.Lock()
		for ip, v := range rl.visitors {
			if time.Since(v.lastSeen) > 5*time.Minute {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}
