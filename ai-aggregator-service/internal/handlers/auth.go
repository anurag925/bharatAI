package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// LoginRequest represents the login request structure
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// LoginResponse represents the login response structure
type LoginResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	TokenType    string   `json:"token_type"`
	ExpiresIn    int      `json:"expires_in"`
	User         UserInfo `json:"user"`
}

// UserInfo represents user information
type UserInfo struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

// RegisterRequest represents the registration request structure
type RegisterRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	Name            string `json:"name" validate:"required"`
	Username        string `json:"username" validate:"required,min=3,max=30"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

// RegisterResponse represents the registration response structure
type RegisterResponse struct {
	User UserInfo `json:"user"`
}

// RefreshTokenRequest represents the refresh token request structure
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// RefreshTokenResponse represents the refresh token response structure
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// LogoutRequest represents the logout request structure
type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// Login godoc
// @Summary User login
// @Description Authenticate user with email and password to obtain access and refresh tokens
// @Tags Authentication
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse "Successfully authenticated"
// @Failure 400 {object} map[string]interface{} "Bad request - Invalid request format"
// @Failure 401 {object} map[string]interface{} "Unauthorized - Invalid credentials"
// @Failure 422 {object} map[string]interface{} "Unprocessable entity - Validation errors"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/login [post]
func (h *handler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"code":    "INVALID_REQUEST",
				"message": "Invalid request format",
			},
		})
	}

	// TODO: Validate request
	// TODO: Check user credentials
	// TODO: Generate tokens
	// TODO: Update last login

	// Mock response for now
	response := LoginResponse{
		AccessToken:  "mock_access_token_" + generateID(),
		RefreshToken: "mock_refresh_token_" + generateID(),
		TokenType:    "Bearer",
		ExpiresIn:    3600,
		User: UserInfo{
			ID:       "user_" + generateID(),
			Email:    req.Email,
			Name:     "Test User",
			Username: "testuser",
		},
	}

	return c.JSON(http.StatusOK, response)
}

// Register godoc
// @Summary User registration
// @Description Register a new user account with email, password, and personal details
// @Tags Authentication
// @Accept json
// @Produce json
// @Param register body RegisterRequest true "Registration details"
// @Success 201 {object} RegisterResponse "User successfully registered"
// @Failure 400 {object} map[string]interface{} "Bad request - Invalid request format"
// @Failure 409 {object} map[string]interface{} "Conflict - Email or username already exists"
// @Failure 422 {object} map[string]interface{} "Unprocessable entity - Validation errors"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/register [post]
func (h *handler) Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"code":    "INVALID_REQUEST",
				"message": "Invalid request format",
			},
		})
	}

	// TODO: Validate request
	// TODO: Check if email already exists
	// TODO: Check if username already exists
	// TODO: Hash password
	// TODO: Create user
	// TODO: Generate tokens
	// TODO: Send welcome email

	// Mock response for now
	response := RegisterResponse{
		User: UserInfo{
			ID:       "user_" + generateID(),
			Email:    req.Email,
			Name:     req.Name,
			Username: req.Username,
		},
	}

	return c.JSON(http.StatusCreated, response)
}

// RefreshToken godoc
// @Summary Refresh access token
// @Description Generate a new access token using a valid refresh token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param refresh body RefreshTokenRequest true "Refresh token details"
// @Success 200 {object} RefreshTokenResponse "New access token generated successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - Invalid request format"
// @Failure 401 {object} map[string]interface{} "Unauthorized - Invalid or expired refresh token"
// @Failure 422 {object} map[string]interface{} "Unprocessable entity - Validation errors"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/refresh [post]
func (h *handler) RefreshToken(c echo.Context) error {
	var req RefreshTokenRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"code":    "INVALID_REQUEST",
				"message": "Invalid request format",
			},
		})
	}

	// TODO: Validate refresh token
	// TODO: Check if token is blacklisted
	// TODO: Generate new access token
	// TODO: Update token expiry

	// Mock response for now
	response := RefreshTokenResponse{
		AccessToken: "new_mock_access_token_" + generateID(),
		TokenType:   "Bearer",
		ExpiresIn:   3600,
	}

	return c.JSON(http.StatusOK, response)
}

// Logout godoc
// @Summary User logout
// @Description Logout user by invalidating refresh token and clearing session
// @Tags Authentication
// @Accept json
// @Produce json
// @Param logout body LogoutRequest true "Logout details"
// @Success 200 {object} map[string]interface{} "Successfully logged out"
// @Failure 400 {object} map[string]interface{} "Bad request - Invalid request format"
// @Failure 401 {object} map[string]interface{} "Unauthorized - Invalid refresh token"
// @Failure 422 {object} map[string]interface{} "Unprocessable entity - Validation errors"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/logout [post]
func (h *handler) Logout(c echo.Context) error {
	var req LogoutRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"code":    "INVALID_REQUEST",
				"message": "Invalid request format",
			},
		})
	}

	// TODO: Validate refresh token
	// TODO: Blacklist refresh token
	// TODO: Blacklist access token
	// TODO: Clear user session

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully logged out",
	})
}

// Me godoc
// @Summary Get current user
// @Description Retrieve information about the currently authenticated user
// @Tags Authentication
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} UserInfo "Current user information"
// @Failure 401 {object} map[string]interface{} "Unauthorized - Missing or invalid access token"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/me [get]
func (h *handler) Me(c echo.Context) error {
	// TODO: Get user from context (after authentication middleware)
	// Mock response for now
	user := UserInfo{
		ID:       "user_" + generateID(),
		Email:    "user@example.com",
		Name:     "Current User",
		Username: "currentuser",
	}

	return c.JSON(http.StatusOK, user)
}

// ForgotPassword godoc
// @Summary Request password reset
// @Description Send password reset email to the provided email address
// @Tags Authentication
// @Accept json
// @Produce json
// @Param forgotPassword body struct{Email string "json:\"email\" validate:\"required,email\""} true "Email for password reset"
// @Success 200 {object} map[string]interface{} "Password reset email sent if email exists"
// @Failure 400 {object} map[string]interface{} "Bad request - Invalid request format"
// @Failure 422 {object} map[string]interface{} "Unprocessable entity - Validation errors"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/forgot-password [post]
func (h *handler) ForgotPassword(c echo.Context) error {
	var req struct {
		Email string `json:"email" validate:"required,email"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"code":    "INVALID_REQUEST",
				"message": "Invalid request format",
			},
		})
	}

	// TODO: Validate email
	// TODO: Generate password reset token
	// TODO: Send password reset email

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "If the email exists, a password reset link has been sent",
	})
}

// ResetPassword godoc
// @Summary Reset user password
// @Description Reset user password using a valid reset token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param resetPassword body struct{Token string "json:\"token\" validate:\"required\""; NewPassword string "json:\"new_password\" validate:\"required,min=8\""} true "Password reset details"
// @Success 200 {object} map[string]interface{} "Password successfully reset"
// @Failure 400 {object} map[string]interface{} "Bad request - Invalid request format"
// @Failure 401 {object} map[string]interface{} "Unauthorized - Invalid or expired reset token"
// @Failure 422 {object} map[string]interface{} "Unprocessable entity - Validation errors"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/reset-password [post]
func (h *handler) ResetPassword(c echo.Context) error {
	var req struct {
		Token       string `json:"token" validate:"required"`
		NewPassword string `json:"new_password" validate:"required,min=8"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"code":    "INVALID_REQUEST",
				"message": "Invalid request format",
			},
		})
	}

	// TODO: Validate reset token
	// TODO: Update user password
	// TODO: Invalidate reset token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Password has been reset successfully",
	})
}
