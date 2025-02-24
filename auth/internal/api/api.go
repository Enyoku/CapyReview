package api

import (
	middleware "authService/internal/api/middlewares"
	"authService/internal/auth"
	"authService/internal/config"
	"authService/internal/db"
	"authService/internal/models"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type API struct {
	router *gin.Engine
	db     *db.DB
	jwt    *config.JWT
}

func New(db *db.DB, jwt *config.JWT) (*API, error) {
	// Initialize Gin
	gin.SetMode(gin.DebugMode)

	api := &API{
		router: gin.New(),
		db:     db,
		jwt:    jwt,
	}
	api.Endpoints()
	return api, nil
}

func (api *API) Endpoints() {
	// Middlewares
	api.router.Use(middleware.HeaderMiddleware())
	api.router.Use(middleware.LoggerMiddleware())

	api.router.Use(gin.Recovery())

	// Public routes
	api.router.POST("/api/login", api.login)
	api.router.POST("/api/register", api.register)

	// Protected routes
	authGroup := api.router.Group("/api/account/")
	authGroup.Use(middleware.RequiredRole(api.jwt, "user", "admin"))
	{
		authGroup.GET("/me", api.getAccountInfo)
		authGroup.PATCH("/me", api.updateAccountInfo)
		authGroup.DELETE("/me", api.deleteAccount)
		authGroup.GET("/logout", api.logout)
	}

	// Private routes
	admin := api.router.Group("/api/admin")
	admin.Use(middleware.RequiredRole(api.jwt, "admin"))
	{
		admin.GET("/dashboard", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"msg": "Admin dashboard"}) })
	}
}

func (api *API) Run(addr string) {
	log.Printf("Starting server on port %v", addr)
	api.router.Run(addr)
}

// POST Хендлер регистрации нового аккаунта
func (api *API) register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверка email и username на уникальность
	if _, exists := db.FindUserByEmail(c.Request.Context(), api.db, user.Email); exists {
		c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
		return
	}

	if _, exists := db.FindUserByUsername(c.Request.Context(), api.db, user.Username); exists {
		c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
		return
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		log.Error().Msg("Unable to hash password")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "something went wrong",
		})
		return
	}

	id, err := db.AddUser(c.Request.Context(), api.db, &user, hashedPassword)
	if err != nil || id == 0 {
		log.Error().Msg("Unable to create user")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg": "registered",
	})
}

// POST Хендлер для входа в аккаунт
func (api *API) login(c *gin.Context) {
	var loginData models.LoginData

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}

	user, exists := db.FindUserByEmail(c.Request.Context(), api.db, loginData.Email)
	if !exists {
		c.JSON(http.StatusConflict, gin.H{"error": "this user does not exist"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"err": "internal server error"})
		}
		return
	}

	tokenString, err := auth.GenerateToken(&user, *api.jwt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
		return
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Now().Add(24 * time.Hour),
	}

	c.SetCookie(cookie.Name, cookie.Value, int(cookie.Expires.Unix()-time.Now().Unix()), cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	c.JSON(http.StatusOK, gin.H{"msg": "Login successful"})
}

// GET Хендлер для выхода из аккаунта
func (api *API) logout(c *gin.Context) {
	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Now().Add(-1 * time.Hour),
	}
	c.SetCookie(cookie.Name, cookie.Value, int(cookie.Expires.Unix()-time.Now().Unix()), cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	c.JSON(http.StatusOK, gin.H{"msg": "Logout successful"})
}

// GET Хендлер предоставляющий информацию о пользователе
func (api *API) getAccountInfo(c *gin.Context) {
	var id models.Uid
	var user *models.UserProfileInfo

	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := db.GetUserInfo(c.Request.Context(), api.db, id.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// PATCH Хендлер обновления данных аккаунта
func (api *API) updateAccountInfo(c *gin.Context) {

}

// DELETE Хендлер удаления аккаунта
func (api *API) deleteAccount(c *gin.Context) {
	var id models.Uid

	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.DeleteUser(c.Request.Context(), api.db, id.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "deleted")
}
