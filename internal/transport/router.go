package transport

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	userHandler *UserHandler
}

func NewRouter(userHandler *UserHandler) *Router {
	return &Router{userHandler: userHandler}
}

func (router *Router) SetupRoutes() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/api/users", router.userHandler.HandleCreateUser).Methods("POST")
	muxRouter.HandleFunc("/api/users/{id}", router.userHandler.HandleGetUser).Methods("GET")
	return muxRouter
}

func (r *Router) Start(port string) error {
	router := r.SetupRoutes()
	if err := http.ListenAndServe(":"+port, router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
