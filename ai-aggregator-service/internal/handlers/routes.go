package handlers

import (
	"net/http"

	"ai-aggregator-service/internal/middleware"

	"github.com/labstack/echo/v4"
)

// SetupRoutes configures all API routes for the AI Aggregator Service
func SetupRoutes(e *echo.Echo) {
	// Initialize handlers
	handler := NewHandler()

	// Health check endpoint (public)
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":    "ok",
			"service":   "ai-aggregator-service",
			"timestamp": c.Get("request_id"),
		})
	})

	// API versioning
	v1 := e.Group("/api/v1")

	// Public routes (no authentication required)
	public := v1.Group("")
	{
		// Auth routes
		auth := public.Group("/auth")
		{
			auth.POST("/login", handler.Login)
			auth.POST("/register", handler.Register)
			auth.POST("/refresh", handler.RefreshToken)
			auth.POST("/logout", handler.Logout)
			auth.POST("/forgot-password", handler.ForgotPassword)
			auth.POST("/reset-password", handler.ResetPassword)
		}

		// OpenAI-compatible API routes (public access with API key)
		openai := public.Group("/openai")
		{
			openai.GET("/models", handler.ListModels)
			openai.POST("/chat/completions", handler.ChatCompletions)
			openai.POST("/completions", handler.Completions)
			openai.POST("/embeddings", handler.Embeddings)
		}
	}

	// Protected routes (require authentication)
	protected := v1.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// User management routes
		users := protected.Group("/users")
		{
			users.GET("/profile", handler.GetProfile)
			users.PUT("/profile", handler.UpdateProfile)

			// API Key management
			apiKeys := users.Group("/api-keys")
			{
				apiKeys.GET("", handler.ListAPIKeys)
				apiKeys.POST("", handler.CreateAPIKey)
				apiKeys.PUT("/:id", handler.UpdateAPIKey)
			}
		}

		// Billing routes
		billing := protected.Group("/billing")
		{
			billing.GET("/usage", handler.GetUsage)

			// Invoice management
			invoices := billing.Group("/invoices")
			{
				invoices.GET("/:id", handler.GetInvoice)
			}

			// Payment method management
			paymentMethods := billing.Group("/payment-methods")
			{
				paymentMethods.POST("", handler.AddPaymentMethod)
			}

			// // Subscription management
			// subscriptions := billing.Group("/subscriptions")
			// {
			// 	subscriptions.GET("", handler.GetSubscriptions)
			// 	subscriptions.POST("", handler.CreateSubscription)
			// 	subscriptions.PUT("/:id", handler.UpdateSubscription)
			// 	subscriptions.DELETE("/:id", handler.CancelSubscription)
			// }
		}

		// API Gateway routes (for direct provider access)
		gateway := protected.Group("/gateway")
		{
			gateway.GET("/models", handler.ListModels)

			// Provider-specific endpoints
			providers := gateway.Group("/providers")
			{
				providers.POST("/:provider/chat/completions", handler.ChatCompletions)

				// Completion endpoints
				providers.POST("/:provider/completions", handler.Completions)

				// Embedding endpoints
				providers.POST("/:provider/embeddings", handler.Embeddings)
			}
		}

		// Unified API routes (provider-agnostic)
		unified := protected.Group("/unified")
		{
			unified.GET("/models", handler.ListModels)
			unified.POST("/chat/completions", handler.ChatCompletions)
			unified.POST("/completions", handler.Completions)
			unified.POST("/embeddings", handler.Embeddings)
		}
	}

	// Admin routes (require admin role)
	admin := v1.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
}

// SetupTestRoutes configures test routes for development/testing
func SetupTestRoutes(e *echo.Echo) {
	// Test endpoints for development
	test := e.Group("/test")
	{
		test.GET("/health", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]string{
				"status":  "ok",
				"service": "test-endpoint",
			})
		})

		test.GET("/echo", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"method": c.Request().Method,
				"path":   c.Request().URL.Path,
				"query":  c.QueryParams(),
				"headers": map[string]string{
					"User-Agent": c.Request().UserAgent(),
					"IP":         c.RealIP(),
				},
			})
		})

		test.POST("/echo", func(c echo.Context) error {
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
}
