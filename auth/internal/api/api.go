package api

import (
	middleware "authService/internal/api/middlewares"
	"authService/internal/db"
	"authService/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type API struct {
	router *gin.Engine
	db     *db.DB
}

func New(db *db.DB) (*API, error) {
	// Initialize Gin
	gin.SetMode(gin.DebugMode)

	api := &API{
		router: gin.New(),
		db:     db,
	}
	api.Endpoints()
	return api, nil
}

func (api *API) Endpoints() {
	// Middlewares
	api.router.Use(middleware.HeaderMiddleware())
	api.router.Use(middleware.LoggerMiddleware())
	api.router.Use(gin.Recovery())

	// Handlers
	authGroup := api.router.Group("/api/account/")
	authGroup.GET("", api.addUser)
}

func (api *API) Run(addr string) {
	log.Printf("Starting server on port %v", addr)
	api.router.Run(addr)
}

func (api *API) addUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		log.Error().Msg("Unable to hash password")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "something went wrong",
		})
		return
	}

	db.AddUser(c.Request.Context(), api.db, &user, hashedPassword)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func hashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}
	return string(hash), err
}
