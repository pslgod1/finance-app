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

func (r *Router) SetupRoutes() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/api/users", r.userHandler.HandleCreateUser).Methods("POST")
	muxRouter.HandleFunc("/api/users/{id}", r.userHandler.HandleGetUser).Methods("GET")

	muxRouter.HandleFunc("/api/users/{userId}/transactions", r.transactionHandler.HandleCreateTransaction).Methods("POST")
	muxRouter.HandleFunc("/api/users/{userId}/transactions", r.transactionHandler.HandleGetTransactions).Methods("GET")
	muxRouter.HandleFunc("/api/users/{userId}/transactions/{id}", r.transactionHandler.HandleUpdateTransaction).Methods("PUT")
	muxRouter.HandleFunc("/api/users/{userId}/transactions/{id}", r.transactionHandler.HandleDeleteTransaction).Methods("DELETE")
	muxRouter.HandleFunc("/api/users/{userId}/statistics", r.transactionHandler.HandleGetStatistics).Methods("GET")

	muxRouter.HandleFunc("/api/admin/users/{id}", r.userHandler.HandleDeleteUser).Methods("DELETE")
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
