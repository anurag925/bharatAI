package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetupMiddleware configures global middleware for the API gateway
func (h *handler) SetupMiddleware(e *echo.Echo) {
	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "X-API-Key"},
		ExposeHeaders:    []string{"X-Total-Count", "X-RateLimit-Limit", "X-RateLimit-Remaining", "X-RateLimit-Reset"},
		MaxAge:           86400,
		AllowCredentials: true,
	}))

	// Request ID middleware
	e.Use(middleware.RequestID())

	// Request/Response logging middleware
	e.Use(h.loggingMiddleware)

	// Rate limiting middleware
	e.Use(h.rateLimitMiddleware)

	// Recover middleware for panic handling
	e.Use(middleware.Recover())

	// Body limit middleware
	e.Use(middleware.BodyLimit("10M"))

	// Security headers
	e.Use(middleware.Secure())
}

// loggingMiddleware logs HTTP requests and responses
func (h *handler) loggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		// Log request
		req := c.Request()
		res := c.Response()

		// Use Echo's built-in logger for now
		c.Logger().Infof("HTTP request started: %s %s from %s", req.Method, req.URL.Path, c.RealIP())

		// Process request
		err := next(c)

		// Log response
		duration := time.Since(start)
		c.Logger().Infof("HTTP request completed: %s %s - %d (%v)", req.Method, req.URL.Path, res.Status, duration)

		return err
	}
}

// rateLimitMiddleware implements rate limiting
func (h *handler) rateLimitMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Implement actual rate limiting based on user/API key
		// For now, we'll use a simple IP-based rate limiter

		// Get client identifier (IP or API key)
		_ = h.getClientIdentifier(c)

		// TODO: Check rate limit in Redis/database
		// For now, we'll allow all requests

		// Set rate limit headers
		c.Response().Header().Set("X-RateLimit-Limit", "1000")
		c.Response().Header().Set("X-RateLimit-Remaining", "999")
		c.Response().Header().Set("X-RateLimit-Reset", "3600")

		return next(c)
	}
}

// getClientIdentifier returns a unique identifier for rate limiting
func (h *handler) getClientIdentifier(c echo.Context) string {
	// Check for API key in header
	apiKey := c.Request().Header.Get("X-API-Key")
	if apiKey != "" {
		return "api_key:" + apiKey
	}

	// Check for Authorization header
	auth := c.Request().Header.Get("Authorization")
	if strings.HasPrefix(auth, "Bearer ") {
		return "token:" + auth[7:]
	}

	// Fallback to IP address
	return "ip:" + c.RealIP()
}

// HealthCheckHandler handles health check endpoint
// @Summary Health check endpoint
// @Description Returns the current health status of the API gateway service. This endpoint is used for monitoring, load balancer health checks, and service discovery. It provides basic service information including current status, timestamp, and version.
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Service health status"
// @Success 200 {string} object.status "healthy" "Service status indicator"
// @Success 200 {string} object.timestamp "2023-08-10T13:39:20Z" "Current UTC timestamp in ISO 8601 format"
// @Success 200 {string} object.version "1.0.0" "Current API version"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /health [get]
func (h *handler) HealthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().UTC(),
		"version":   "1.0.0",
	})
}

// NotFoundHandler handles 404 errors
// @Summary Handle 404 Not Found errors
// @Description Returns a standardized error response when a requested resource or endpoint does not exist. This handler is triggered automatically by the Echo router when no matching route is found for the requested path. It provides detailed error information including the error code, descriptive message, and the actual path that was not found.
// @Tags error
// @Accept json
// @Produce json
// @Param path path string true "The requested path that was not found" default(/nonexistent/endpoint)
// @Success 404 {object} map[string]interface{} "Not found error response"
// @Success 404 {object} object.error "Error details"
// @Success 404 {string} object.error.code "NOT_FOUND" "Error code identifier"
// @Success 404 {string} object.error.message "The requested resource was not found" "Human-readable error description"
// @Success 404 {string} object.error.path "/nonexistent/endpoint" "The actual path that triggered the 404 error"
// @Router /* [get]
// @Router /* [post]
// @Router /* [put]
// @Router /* [delete]
// @Router /* [patch]
// @Router /* [head]
// @Router /* [options]
func (h *handler) NotFoundHandler(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"error": map[string]interface{}{
			"code":    "NOT_FOUND",
			"message": "The requested resource was not found",
			"path":    c.Request().URL.Path,
		},
	})
}

// MethodNotAllowedHandler handles 405 errors
// @Summary Handle 405 Method Not Allowed errors
// @Description Returns a standardized error response when an HTTP method is not supported for a valid endpoint. This handler is triggered automatically by the Echo router when a client attempts to use an HTTP method (GET, POST, PUT, DELETE, etc.) that is not allowed for the requested endpoint. It provides detailed error information including the error code, descriptive message, the attempted HTTP method, and the path.
// @Tags error
// @Accept json
// @Produce json
// @Param method path string true "The HTTP method that is not allowed" default(PATCH)
// @Param path path string true "The endpoint path where the method is not allowed" default(/api/v1/resource)
// @Success 405 {object} map[string]interface{} "Method not allowed error response"
// @Success 405 {object} object.error "Error details"
// @Success 405 {string} object.error.code "METHOD_NOT_ALLOWED" "Error code identifier"
// @Success 405 {string} object.error.message "Method not allowed for this endpoint" "Human-readable error description"
// @Success 405 {string} object.error.method "PATCH" "The HTTP method that was attempted"
// @Success 405 {string} object.error.path "/api/v1/resource" "The endpoint path where the method was attempted"
// @Router /* [get]
// @Router /* [post]
// @Router /* [put]
// @Router /* [delete]
// @Router /* [patch]
// @Router /* [head]
// @Router /* [options]
func (h *handler) MethodNotAllowedHandler(c echo.Context) error {
	return c.JSON(http.StatusMethodNotAllowed, map[string]interface{}{
		"error": map[string]interface{}{
			"code":    "METHOD_NOT_ALLOWED",
			"message": "Method not allowed for this endpoint",
			"method":  c.Request().Method,
			"path":    c.Request().URL.Path,
		},
	})
}
