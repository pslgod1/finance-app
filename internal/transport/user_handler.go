package transport

import (
	"FinanceApp/internal/dto"
	"FinanceApp/internal/service"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserHandler struct {
	*service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: svc,
	}
}

func (h *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	userRequest := dto.CreateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		errResp := dto.NewErrorResponse(err)

		http.Error(w, errResp.ToString(), http.StatusBadRequest)
		return
	}

	// проверка на пользователя существует или не
	user, err := h.CreateUser(r.Context(), userRequest.Username, userRequest.Password)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			errResp := dto.NewErrorResponse(fmt.Errorf("пользователь уже существует"))
			http.Error(w, errResp.ToString(), http.StatusConflict)

			return
		}
		errResp := dto.NewErrorResponse(err)
		http.Error(w, errResp.ToString(), http.StatusInternalServerError)

		return
	}

	userResp := dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(userResp); err != nil {
		errResp := dto.NewErrorResponse(err)

		http.Error(w, errResp.ToString(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errResp := dto.NewErrorResponse(fmt.Errorf("некорректный id"))
		http.Error(w, errResp.ToString(), http.StatusBadRequest)
		return
	}

	user, err := h.GetUser(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			errResp := dto.NewErrorResponse(fmt.Errorf("пользователь не найден"))
			http.Error(w, errResp.ToString(), http.StatusNotFound)
			return
		}
		errResp := dto.NewErrorResponse(err)
		http.Error(w, errResp.ToString(), http.StatusInternalServerError)
		return
	}

	userResp := dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userResp); err != nil {
		errResp := dto.NewErrorResponse(err)
		http.Error(w, errResp.ToString(), http.StatusInternalServerError)
		return
	}

}
