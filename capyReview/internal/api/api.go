package api

import "APIGateway/internal/db"

type API struct {
	// router *gin.Engine
	// db *db.DB
	// jwt
}

func New(db *db.DB) *API {
	return &API{}
}

func (api *API) Run(addr string) {

}
