package transport

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	userHandler        *UserHandler
	transactionHandler *TransactionHandler
}

func NewRouter(userHandler *UserHandler, transactionHandler *TransactionHandler) *Router {
	return &Router{
		userHandler:        userHandler,
		transactionHandler: transactionHandler,
	}
}

func (router *Router) SetupRoutes() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/api/users", router.userHandler.HandleCreateUser).Methods("POST")
	muxRouter.HandleFunc("/api/users/{id}", router.userHandler.HandleGetUser).Methods("GET")

	muxRouter.HandleFunc("/api/users/{userId}/transactions", router.transactionHandler.HandleCreateTransaction).Methods("POST")
	muxRouter.HandleFunc("/api/users/{userId}/transactions", router.transactionHandler.HandleGetTransactions).Methods("GET")
	muxRouter.HandleFunc("/api/users/{userId}/transactions/{id}", router.transactionHandler.HandleUpdateTransaction).Methods("PUT")
	muxRouter.HandleFunc("/api/users/{userId}/transactions/{id}", router.transactionHandler.HandleDeleteTransaction).Methods("DELETE")
	muxRouter.HandleFunc("/api/users/{userId}/statistics", router.transactionHandler.HandleGetStatistics).Methods("GET")

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
