package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// UserProfile represents user profile information
type UserProfile struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Avatar    string    `json:"avatar,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Company   string    `json:"company,omitempty"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UpdateProfileRequest represents the update profile request structure
type UpdateProfileRequest struct {
	Name     string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Username string `json:"username,omitempty" validate:"omitempty,min=3,max=30"`
	Avatar   string `json:"avatar,omitempty" validate:"omitempty,url"`
	Phone    string `json:"phone,omitempty" validate:"omitempty,e164"`
	Company  string `json:"company,omitempty" validate:"omitempty,max=100"`
}

// Organization represents an organization
type Organization struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description,omitempty"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// APIKey represents an API key
type APIKey struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Key         string    `json:"key"`
	Prefix      string    `json:"prefix"`
	LastUsed    time.Time `json:"last_used,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at,omitempty"`
	IsActive    bool      `json:"is_active"`
	Permissions []string  `json:"permissions"`
}

// CreateAPIKeyRequest represents the create API key request structure
type CreateAPIKeyRequest struct {
	Name        string   `json:"name" validate:"required,min=3,max=50"`
	ExpiresIn   int      `json:"expires_in,omitempty"` // days
	Permissions []string `json:"permissions,omitempty"`
}

// CreateAPIKeyResponse represents the create API key response structure
type CreateAPIKeyResponse struct {
	APIKey APIKey `json:"api_key"`
	Key    string `json:"key"` // Only returned on creation
}

// GetProfile handles GET /users/profile
// @Summary Get user profile
// @Description Retrieves the authenticated user's profile information including personal details, role, and timestamps
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} UserProfile "User profile retrieved successfully"
// @Failure 401 {object} map[string]interface{} "Unauthorized - Invalid or missing authentication token"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /users/profile [get]
func (h *handler) GetProfile(c echo.Context) error {
	// TODO: Get user ID from context (after authentication middleware)
	userID := "user_" + generateID()

	// TODO: Fetch user profile from database
	// Mock response for now
	profile := UserProfile{
		ID:        userID,
		Email:     "user@example.com",
		Name:      "Test User",
		Username:  "testuser",
		Avatar:    "https://example.com/avatar.jpg",
		Phone:     "+1234567890",
		Company:   "Test Company",
		Role:      "user",
		CreatedAt: time.Now().AddDate(-1, 0, 0),
		UpdatedAt: time.Now(),
	}

	return c.JSON(http.StatusOK, profile)
}

// UpdateProfile handles PUT /users/profile
// @Summary Update user profile
// @Description Updates the authenticated user's profile information. Only provided fields will be updated. Username must be unique if provided.
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param profile body UpdateProfileRequest true "Profile update request"
// @Success 200 {object} UserProfile "Profile updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - Invalid request format or validation errors"
// @Failure 401 {object} map[string]interface{} "Unauthorized - Invalid or missing authentication token"
// @Failure 409 {object} map[string]interface{} "Conflict - Username already taken"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /users/profile [put]
func (h *handler) UpdateProfile(c echo.Context) error {
	var req UpdateProfileRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"code":    "INVALID_REQUEST",
				"message": "Invalid request format",
			},
		})
	}

	// TODO: Validate request
	// TODO: Get user ID from context
	// TODO: Update user profile in database
	// TODO: Check username uniqueness if provided

	// Mock response for now
	profile := UserProfile{
		ID:        "user_" + generateID(),
		Email:     "user@example.com",
		Name:      req.Name,
		Username:  req.Username,
		Avatar:    req.Avatar,
		Phone:     req.Phone,
		Company:   req.Company,
		Role:      "user",
		CreatedAt: time.Now().AddDate(-1, 0, 0),
		UpdatedAt: time.Now(),
	}

	return c.JSON(http.StatusOK, profile)
}

// GetOrganizations handles GET /users/organizations
// @Summary Get user organizations
// @Description Retrieves all organizations that the authenticated user is a member of, including their role in each organization
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Organizations retrieved successfully"
// @Success 200 {object} map[string]interface{} "Schema: {\"organizations\": []Organization, \"total\": integer}"
// @Failure 401 {object} map[string]interface{} "Unauthorized - Invalid or missing authentication token"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /users/organizations [get]
func (h *handler) GetOrganizations(c echo.Context) error {
	// TODO: Get user ID from context
	// TODO: Fetch user's organizations from database

	// Mock response for now
	organizations := []Organization{
		{
			ID:          "org_" + generateID(),
			Name:        "Personal Workspace",
			Slug:        "personal-workspace",
			Description: "Your personal workspace",
			Role:        "owner",
			CreatedAt:   time.Now().AddDate(-1, 0, 0),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          "org_" + generateID(),
			Name:        "Test Company",
			Slug:        "test-company",
			Description: "Test company organization",
			Role:        "member",
			CreatedAt:   time.Now().AddDate(-6, 0, 0),
			UpdatedAt:   time.Now(),
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"organizations": organizations,
		"total":         len(organizations),
	})
}

// CreateAPIKey handles POST /users/api-keys
// @Summary Create API key
// @Description Creates a new API key for the authenticated user with specified permissions and optional expiration. The full API key is only returned on creation.
// @Tags api-keys
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param api_key body CreateAPIKeyRequest true "API key creation request"
// @Success 201 {object} CreateAPIKeyResponse "API key created successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - Invalid request format or validation errors"
// @Failure 401 {object} map[string]interface{} "Unauthorized - Invalid or missing authentication token"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /users/api-keys [post]
func (h *handler) CreateAPIKey(c echo.Context) error {
	var req CreateAPIKeyRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"code":    "INVALID_REQUEST",
				"message": "Invalid request format",
			},
		})
	}

	// TODO: Validate request
	// TODO: Get user ID from context
	// TODO: Generate API key
	// TODO: Store API key in database
	// TODO: Set permissions

	// Generate mock API key
	apiKey := "sk_" + generateID() + generateID()
	prefix := apiKey[:8]

	// Mock response for now
	response := CreateAPIKeyResponse{
		APIKey: APIKey{
			ID:        "key_" + generateID(),
			Name:      req.Name,
			Key:       apiKey,
			Prefix:    prefix,
			CreatedAt: time.Now(),
			ExpiresAt: time.Now().AddDate(0, 0, req.ExpiresIn),
			IsActive:  true,
			Permissions: []string{
				"chat:read",
				"chat:write",
				"completions:read",
				"completions:write",
			},
		},
		Key: apiKey,
	}

	return c.JSON(http.StatusCreated, response)
}

// ListAPIKeys handles GET /users/api-keys
// @Summary List API keys
// @Description Retrieves all API keys belonging to the authenticated user. Returns a paginated list of API keys with their metadata (excluding the actual key values for security).
// @Tags api-keys
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Maximum number of results to return (default: 50, max: 100)" default(50) minimum(1) maximum(100)
// @Param offset query int false "Number of results to skip for pagination (default: 0)" default(0) minimum(0)
// @Success 200 {object} map[string]interface{} "API keys retrieved successfully"
// @Success 200 {object} map[string]interface{} "Schema: {\"api_keys\": []APIKey, \"total\": integer, \"limit\": integer, \"offset\": integer}"
// @Failure 401 {object} map[string]interface{} "Unauthorized - Invalid or missing authentication token"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /users/api-keys [get]
func (h *handler) ListAPIKeys(c echo.Context) error {
	// TODO: Get user ID from context
	// TODO: Fetch user's API keys from database
	// TODO: Support pagination

	// Mock response for now
	apiKeys := []APIKey{
		{
			ID:        "key_" + generateID(),
			Name:      "Production Key",
			Prefix:    "sk_abc123",
			LastUsed:  time.Now().AddDate(0, 0, -1),
			CreatedAt: time.Now().AddDate(-1, 0, 0),
			ExpiresAt: time.Now().AddDate(0, 6, 0),
			IsActive:  true,
			Permissions: []string{
				"chat:read",
				"chat:write",
				"completions:read",
				"completions:write",
			},
		},
		{
			ID:        "key_" + generateID(),
			Name:      "Development Key",
			Prefix:    "sk_def456",
			LastUsed:  time.Now().AddDate(0, 0, -7),
			CreatedAt: time.Now().AddDate(-2, 0, 0),
			ExpiresAt: time.Now().AddDate(0, 3, 0),
			IsActive:  true,
			Permissions: []string{
				"chat:read",
				"chat:write",
			},
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"api_keys": apiKeys,
		"total":    len(apiKeys),
	})
}

// RevokeAPIKey handles DELETE /users/api-keys/:key_id
// @Summary Revoke API key
// @Description Revokes (deactivates) an API key by its ID. The key must belong to the authenticated user. This action is irreversible.
// @Tags api-keys
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param key_id path string true "API key ID to revoke" example("key_1234567890abcdef")
// @Success 200 {object} map[string]interface{} "API key revoked successfully"
// @Success 200 {object} map[string]interface{} "Schema: {\"message\": string, \"id\": string}"
// @Failure 400 {object} map[string]interface{} "Bad request - Invalid API key ID format"
// @Failure 401 {object} map[string]interface{} "Unauthorized - Invalid or missing authentication token"
// @Failure 403 {object} map[string]interface{} "Forbidden - API key does not belong to user"
// @Failure 404 {object} map[string]interface{} "Not found - API key not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /users/api-keys/{key_id} [delete]
func (h *handler) RevokeAPIKey(c echo.Context) error {
	keyID := c.Param("key_id")
	if keyID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"code":    "INVALID_REQUEST",
				"message": "API key ID is required",
			},
		})
	}

	// TODO: Validate key ID format
	// TODO: Get user ID from context
	// TODO: Check if key belongs to user
	// TODO: Revoke API key in database

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "API key revoked successfully",
		"id":      keyID,
	})
}

// GetAPIKey handles GET /users/api-keys/:key_id
// @Summary Get API key details
// @Description Retrieves detailed information about a specific API key by its ID. The key must belong to the authenticated user. Returns metadata excluding the actual key value for security.
// @Tags api-keys
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param key_id path string true "API key ID to retrieve" example("key_1234567890abcdef")
// @Success 200 {object} APIKey "API key details retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - Invalid API key ID format"
// @Failure 401 {object} map[string]interface{} "Unauthorized - Invalid or missing authentication token"
// @Failure 403 {object} map[string]interface{} "Forbidden - API key does not belong to user"
// @Failure 404 {object} map[string]interface{} "Not found - API key not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /users/api-keys/{key_id} [get]
func (h *handler) GetAPIKey(c echo.Context) error {
	keyID := c.Param("key_id")
	if keyID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"code":    "INVALID_REQUEST",
				"message": "API key ID is required",
			},
		})
	}

	// TODO: Get user ID from context
	// TODO: Check if key belongs to user
	// TODO: Fetch API key from database

	// Mock response for now
	apiKey := APIKey{
		ID:        keyID,
		Name:      "Production Key",
		Prefix:    "sk_abc123",
		LastUsed:  time.Now().AddDate(0, 0, -1),
		CreatedAt: time.Now().AddDate(-1, 0, 0),
		ExpiresAt: time.Now().AddDate(0, 6, 0),
		IsActive:  true,
		Permissions: []string{
			"chat:read",
			"chat:write",
			"completions:read",
			"completions:write",
		},
	}

	return c.JSON(http.StatusOK, apiKey)
}

// UpdateAPIKey handles PUT /users/api-keys/:key_id
// @Summary Update API key
// @Description Updates an existing API key's properties. Only provided fields will be updated. The key must belong to the authenticated user.
// @Tags api-keys
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param key_id path string true "API key ID to update" example("key_1234567890abcdef")
//
//	@Param api_key body struct {
//	    Name        string   `json:"name,omitempty" validate:"omitempty,min=3,max=50"`
//	    IsActive    *bool    `json:"is_active,omitempty"`
//	    Permissions []string `json:"permissions,omitempty"`
//	} true "API key update request"
//
// @Success 200 {object} APIKey "API key updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - Invalid request format, validation errors, or invalid API key ID"
// @Failure 401 {object} map[string]interface{} "Unauthorized - Invalid or missing authentication token"
// @Failure 403 {object} map[string]interface{} "Forbidden - API key does not belong to user"
// @Failure 404 {object} map[string]interface{} "Not found - API key not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /users/api-keys/{key_id} [put]
func (h *handler) UpdateAPIKey(c echo.Context) error {
	keyID := c.Param("key_id")
	if keyID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"code":    "INVALID_REQUEST",
				"message": "API key ID is required",
			},
		})
	}

	var req struct {
		Name        string   `json:"name,omitempty" validate:"omitempty,min=3,max=50"`
		IsActive    *bool    `json:"is_active,omitempty"`
		Permissions []string `json:"permissions,omitempty"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"code":    "INVALID_REQUEST",
				"message": "Invalid request format",
			},
		})
	}

	// TODO: Validate request
	// TODO: Get user ID from context
	// TODO: Check if key belongs to user
	// TODO: Update API key in database

	// Mock response for now
	apiKey := APIKey{
		ID:          keyID,
		Name:        req.Name,
		Prefix:      "sk_abc123",
		LastUsed:    time.Now().AddDate(0, 0, -1),
		CreatedAt:   time.Now().AddDate(-1, 0, 0),
		ExpiresAt:   time.Now().AddDate(0, 6, 0),
		IsActive:    *req.IsActive,
		Permissions: req.Permissions,
	}

	return c.JSON(http.StatusOK, apiKey)
}
