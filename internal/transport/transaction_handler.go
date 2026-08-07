package transport

import (
	"FinanceApp/internal/dto"
	"FinanceApp/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	*service.TransactionService
}

func NewTransactionHandler(svc *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		TransactionService: svc,
	}
}

func (h *TransactionHandler) HandleCreateTransaction(w http.ResponseWriter, r *http.Request) {
	transactionRequest := dto.TransactionRequest{}

	if err := json.NewDecoder(r.Body).Decode(&transactionRequest); err != nil {
		errResp := dto.NewErrorResponse(err)

		http.Error(w, errResp.ToString(), http.StatusBadRequest)

		return
	}

	userID, err := strconv.Atoi(mux.Vars(r)["userId"])
	if err != nil {
		errResp := dto.NewErrorResponse(fmt.Errorf("некорректный userId"))
		http.Error(w, errResp.ToString(), http.StatusBadRequest)
		return
	}

	transaction, err := h.CreateTransaction(
		r.Context(),
		userID,
		transactionRequest.Type,
		transactionRequest.Amount,
		transactionRequest.Category,
		transactionRequest.Description,
	)
	if err != nil {
		errResp := dto.NewErrorResponse(err)
		http.Error(w, errResp.ToString(), http.StatusInternalServerError)
		return
	}

	transactionResponse := dto.TransactionResponse{
		ID:          transaction.ID,
		UserID:      transaction.UserID,
		Type:        transaction.Type,
		Amount:      transaction.Amount,
		Category:    transaction.Category,
		Description: transaction.Description,
		CreatedAt:   transaction.CreatedAt,
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(transactionResponse); err != nil {
		errResp := dto.NewErrorResponse(err)

		http.Error(w, errResp.ToString(), http.StatusInternalServerError)
		return
	}
}

func (h *TransactionHandler) HandleGetTransactions(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["userId"])
	if err != nil {
		errResp := dto.NewErrorResponse(err)
		http.Error(w, errResp.ToString(), http.StatusBadRequest)
		return
	}

	transactions, err := h.GetTransactions(r.Context(), userID)
	if err != nil {
		errResp := dto.NewErrorResponse(err)
		http.Error(w, errResp.ToString(), http.StatusInternalServerError)
		return
	}

	var resp []dto.TransactionResponse
	for _, t := range transactions {
		resp = append(resp, dto.TransactionResponse{
			ID:          t.ID,
			UserID:      t.UserID,
			Type:        t.Type,
			Amount:      t.Amount,
			Category:    t.Category,
			Description: t.Description,
			CreatedAt:   t.CreatedAt,
		})
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		errResp := dto.NewErrorResponse(err)
		http.Error(w, errResp.ToString(), http.StatusInternalServerError)
		return
	}
}
