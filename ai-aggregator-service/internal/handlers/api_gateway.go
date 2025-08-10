package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// APIGatewayHandler handles API gateway functionality
type APIGatewayHandler struct{}

// NewAPIGatewayHandler creates a new API gateway handler
func NewAPIGatewayHandler() *APIGatewayHandler {
	return &APIGatewayHandler{}
}

// SetupMiddleware configures global middleware for the API gateway
func (h *APIGatewayHandler) SetupMiddleware(e *echo.Echo) {
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
func (h *APIGatewayHandler) loggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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
func (h *APIGatewayHandler) rateLimitMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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
func (h *APIGatewayHandler) getClientIdentifier(c echo.Context) string {
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
func (h *APIGatewayHandler) HealthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().UTC(),
		"version":   "1.0.0",
	})
}

// NotFoundHandler handles 404 errors
func (h *APIGatewayHandler) NotFoundHandler(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"error": map[string]interface{}{
			"code":    "NOT_FOUND",
			"message": "The requested resource was not found",
			"path":    c.Request().URL.Path,
		},
	})
}

// MethodNotAllowedHandler handles 405 errors
func (h *APIGatewayHandler) MethodNotAllowedHandler(c echo.Context) error {
	return c.JSON(http.StatusMethodNotAllowed, map[string]interface{}{
		"error": map[string]interface{}{
			"code":    "METHOD_NOT_ALLOWED",
			"message": "Method not allowed for this endpoint",
			"method":  c.Request().Method,
			"path":    c.Request().URL.Path,
		},
	})
}
