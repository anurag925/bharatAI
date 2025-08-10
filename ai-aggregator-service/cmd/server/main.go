package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create Echo instance
	e := echo.New()

	// Setup basic middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Simple CORS
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if c.Request().Method == "OPTIONS" {
				return c.NoContent(200)
			}
			return next(c)
		}
	})

	// Setup routes for testing
	setupTestRoutes(e)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting test server on port %s\n", port)
	if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
		log.Fatal("Failed to start server:", err)
	}
}

func setupTestRoutes(e *echo.Echo) {
	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "ok",
			"service": "ai-aggregator-service",
		})
	})

	// Test endpoints for each handler type
	e.GET("/test/api-gateway", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"handler": "api-gateway",
			"message": "API Gateway handler is working",
			"endpoints": []string{
				"GET /api/v1/gateway/models",
				"POST /api/v1/gateway/models/:provider/chat",
				"POST /api/v1/gateway/models/:provider/completion",
				"POST /api/v1/gateway/models/:provider/embeddings",
			},
		})
	})

	e.GET("/test/unified-api", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"handler": "unified-api",
			"message": "Unified API handler is working",
			"endpoints": []string{
				"POST /api/v1/unified/chat",
				"POST /api/v1/unified/completion",
				"POST /api/v1/unified/embeddings",
			},
		})
	})

	e.GET("/test/auth", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"handler": "auth",
			"message": "Auth handler is working",
			"endpoints": []string{
				"POST /api/v1/auth/login",
				"POST /api/v1/auth/register",
				"POST /api/v1/auth/refresh",
				"POST /api/v1/auth/logout",
				"GET /api/v1/auth/me",
			},
		})
	})

	e.GET("/test/user", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"handler": "user",
			"message": "User handler is working",
			"endpoints": []string{
				"GET /api/v1/users/profile",
				"PUT /api/v1/users/profile",
				"GET /api/v1/users/organizations",
				"POST /api/v1/users/api-keys",
				"GET /api/v1/users/api-keys",
				"DELETE /api/v1/users/api-keys/:id",
			},
		})
	})

	e.GET("/test/billing", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"handler": "billing",
			"message": "Billing handler is working",
			"endpoints": []string{
				"GET /api/v1/billing/account",
				"PUT /api/v1/billing/account",
				"GET /api/v1/billing/usage",
				"GET /api/v1/billing/invoices",
				"GET /api/v1/billing/invoices/:id",
				"POST /api/v1/billing/payment-methods",
				"GET /api/v1/billing/payment-methods",
				"DELETE /api/v1/billing/payment-methods/:id",
			},
		})
	})

	// Test POST endpoints
	e.POST("/test/echo", func(c echo.Context) error {
		var body map[string]interface{}
		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid JSON",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"received": body,
			"method":   c.Request().Method,
			"path":     c.Request().URL.Path,
		})
	})
}
