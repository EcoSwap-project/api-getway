package handlers

import (
	auth "api_getway/genproto/authentication_service"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type AuthHendler interface {
	Register(*gin.Context)
	Login(*gin.Context)
	GetProfile(*gin.Context)
	EditProfile(*gin.Context)
	ListUsers(*gin.Context)
	DeleteUser(*gin.Context)
	ResetPassword(*gin.Context)
	RefreshToken(*gin.Context)
	Logout(*gin.Context)
	GetEcoPoints(*gin.Context)
	AddEcoPoints(*gin.Context)
	GetEcoPointsHistory(*gin.Context)
}

type authHandlerImpl struct {
	authService auth.EcoServiceClient
}

func NewAuthHendler(authSer auth.EcoServiceClient) AuthHendler {
	return &authHandlerImpl{authService: authSer}
}

// @Tags Authentication
// @Summary Register a new user
// @Description Register a new user with username, password, and email.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body authentication_service.RegisterRequest true "Register request"
// @Success 200 {object} authentication_service.RegisterResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /auth/register [post]
func (h *authHandlerImpl) Register(c *gin.Context) {
	var req auth.RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.authService.Register(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Tags Authentication
// @Summary Login a user
// @Description Login a user with username and password.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body authentication_service.LoginRequest true "Login request"
// @Success 200 {object} authentication_service.LoginResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /auth/login [post]
func (h *authHandlerImpl) Login(c *gin.Context) {
	var req auth.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.authService.Login(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	Accesstoken, err := GenerateAccessToken(req.Email)
	if err != nil {
		log.Println("Failed to generate JWT", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication token"})
		return
	}

	refresh_token, err := GenerateRefreshToken(req.Email)
	if err != nil{
		log.Println("Failed to generate JWT", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication token"})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"accsess_token": Accesstoken,
		"refresh_token": refresh_token,
	

		"Login": gin.H{
			"Success": res,
		},
	})
}

// @Tags Authentication
// @Summary Get user profile
// @Description Get user profile by ID.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param userId path string true "User ID"
// @Success 200 {object} authentication_service.ProfileResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /auth/profile [get]
func (h *authHandlerImpl) GetProfile(c *gin.Context) {
	userId := c.GetString("userId") // Assuming userId is extracted from JWT in middleware

	res, err := h.authService.GetProfile(c, &auth.ProfilRequest{Id: userId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Tags Authentication
// @Summary Edit user profile
// @Description Edit user profile by ID.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body authentication_service.EditProfileRequest true "Edit profile request"
// @Success 200 {object} authentication_service.EditProfileResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /auth/profile [put]
func (h *authHandlerImpl) EditProfile(c *gin.Context) {
	var req auth.EditProfileRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetString("userId") // Assuming userId is extracted from JWT in middleware
	req.Id = userId

	res, err := h.authService.EditProfile(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Tags Authentication
// @Summary List users
// @Description List all users.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body authentication_service.UsersListRequest true "List users request"
// @Success 200 {object} authentication_service.UsersListResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /auth/users [get]
func (h *authHandlerImpl) ListUsers(c *gin.Context) {
	res, err := h.authService.ListUsers(c, &auth.UsersListRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *authHandlerImpl) DeleteUser(c *gin.Context) {
	userId := c.GetString("userId") // Assuming userId is extracted from JWT in middleware

	res, err := h.authService.DeleteUser(c, &auth.DeleteUserRequest{Id: userId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Tags Authentication
// @Summary Reset user password
// @Description Reset user password by email.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body authentication_service.ResetPasswordRequest true "Reset password request"
// @Success 200 {object} authentication_service.ResetPasswordResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /auth/reset-password [post]
func (h *authHandlerImpl) ResetPassword(c *gin.Context) {
	var req auth.ResetPasswordRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.authService.ResetPassword(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Tags Authentication
// @Summary Refresh token
// @Description Refresh token.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body authentication_service.RefreshTokenRequest true "Refresh token request"
// @Success 200 {object} authentication_service.RefreshTokenResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /auth/refresh [post]
func (h *authHandlerImpl) RefreshToken(c *gin.Context) {
	var req auth.RefreshTokenRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.authService.RefreshToken(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Tags Authentication
// @Summary Logout a user
// @Description Logout a user.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} authentication_service.LogoutResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /auth/logout [post]
func (h *authHandlerImpl) Logout(c *gin.Context) {
	userId := c.GetString("userId") // Assuming userId is extracted from JWT in middleware

	res, err := h.authService.Logout(c, &auth.LogoutRequest{AccessToken: userId}) // ! tokin kerak
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Tags Authentication
// @Summary Get user eco points
// @Description Get user eco points by ID.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} authentication_service.EcoPointsResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /auth/eco-points [get]
func (h *authHandlerImpl) GetEcoPoints(c *gin.Context) {
	userId := c.GetString("userId") // Assuming userId is extracted from JWT in middleware

	res, err := h.authService.GetEcoPoints(c, &auth.EcoPointsRequest{UserId: userId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Tags Authentication
// @Summary Add eco points to user
// @Description Add eco points to user by ID.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body authentication_service.AddEcoPointsRequest true "Add eco points request"
// @Success 200 {object} authentication_service.AddEcoPointsResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /auth/eco-points [post]
func (h *authHandlerImpl) AddEcoPoints(c *gin.Context) {
	var req auth.AddEcoPointsRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetString("userId") // Assuming userId is extracted from JWT in middleware
	req.UserId = userId

	res, err := h.authService.AddEcoPoints(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Tags Authentication
// @Summary Get user eco points history
// @Description Get user eco points history by ID.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} authentication_service.EcoPointsHistoryResponse
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /auth/eco-points/history [get]
func (h *authHandlerImpl) GetEcoPointsHistory(c *gin.Context) {
	userId := c.GetString("userId") // Assuming userId is extracted from JWT in middleware

	res, err := h.authService.GetEcoPointsHistory(c, &auth.GetEcoPointsHistoryRequest{UserId: userId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Access token yaratish
func GenerateAccessToken(username string) (string, error) {
	expirationTime := time.Now().Add(120 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Refresh token yaratish
func GenerateRefreshToken(username string) (string, error) {
	expirationTime := time.Now().Add(48 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
