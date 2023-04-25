package handlers

import "github.com/realwebdev/clockify/datastore"

type Handler struct {
	DB datastore.DBController
}

func New(dbHandler datastore.DBController) *Handler {
	return &Handler{DB: dbHandler}
}
