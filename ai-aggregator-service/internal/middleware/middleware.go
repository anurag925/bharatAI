package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// SetupMiddleware configures all middleware for the Echo server
func SetupMiddleware(e *echo.Echo) {
	// Request ID middleware (first to ensure ID is available)
	e.Use(RequestIDMiddleware())

	// CORS middleware
	e.Use(CORSMiddleware())

	// Logger middleware
	e.Use(LoggerMiddleware())

	// Rate limiting middleware
	rateLimiter := NewRateLimiter(100, 200) // 100 requests per second, burst of 200
	e.Use(rateLimiter.RateLimitMiddleware())

	// Authentication middleware
	e.Use(AuthMiddleware())

	// Recovery middleware (built-in Echo)
	e.Use(RecoverMiddleware())
}

// RecoverMiddleware handles panics gracefully
func RecoverMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					c.Logger().Error("Panic recovered:", err)
					c.JSON(500, map[string]string{
						"error": "Internal server error",
					})
				}
			}()
			return next(c)
		}
	}
}
