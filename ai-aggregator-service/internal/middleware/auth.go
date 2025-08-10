package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware handles JWT token validation
func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Skip auth for public endpoints
			if isPublicEndpoint(c.Path()) {
				return next(c)
			}

			// Get token from Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Authorization header required",
				})
			}

			// Extract token from Bearer scheme
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid authorization format",
				})
			}

			// TODO: Validate JWT token
			// For now, we'll just check if token exists
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid token",
				})
			}

			// TODO: Extract user claims from token and set in context
			// userID := extractUserIDFromToken(tokenString)
			// c.Set("userID", userID)

			return next(c)
		}
	}
}

// isPublicEndpoint checks if the endpoint doesn't require authentication
func isPublicEndpoint(path string) bool {
	publicEndpoints := []string{
		"/api/v1/auth/login",
		"/api/v1/auth/register",
		"/api/v1/auth/forgot-password",
		"/api/v1/auth/reset-password",
		"/health",
	}

	for _, endpoint := range publicEndpoints {
		if path == endpoint {
			return true
		}
	}
	return false
}
