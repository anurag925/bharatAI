package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

// LoggerMiddleware provides structured logging for HTTP requests
func LoggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			// Process request
			err := next(c)

			// Log request details
			stop := time.Now()
			latency := stop.Sub(start)
			status := c.Response().Status
			method := c.Request().Method
			path := c.Request().URL.Path
			ip := c.RealIP()
			if ip == "" {
				ip = c.Request().RemoteAddr
			}
			userAgent := c.Request().UserAgent()

			// Format log message
			logMessage := fmt.Sprintf(
				"[%s] %s %s %d %s %s \"%s\"",
				stop.Format("2006-01-02 15:04:05"),
				method,
				path,
				status,
				latency,
				ip,
				userAgent,
			)

			// Log to console (can be extended to use structured logging)
			c.Logger().Info(logMessage)

			return err
		}
	}
}

// RequestIDMiddleware adds a unique request ID to each request
func RequestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Generate or get request ID from header
			requestID := c.Request().Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = fmt.Sprintf("%d", time.Now().UnixNano())
			}

			// Set request ID in response header
			c.Response().Header().Set("X-Request-ID", requestID)

			// Store in context for use in handlers
			c.Set("requestID", requestID)

			return next(c)
		}
	}
}

// CORSMiddleware handles Cross-Origin Resource Sharing
func CORSMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

			// Handle preflight requests
			if c.Request().Method == "OPTIONS" {
				return c.NoContent(200)
			}

			return next(c)
		}
	}
}
