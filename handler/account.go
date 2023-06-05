package handler

import (
	"bank/errs"
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const CONTENT_TYPE_HEADER = "content-type"
const CONTENT_TYPE_VALUE = "application/json"

type accountHandler struct {
	accSrv service.AccountService
}

func NewAccountHandler(accSrv service.AccountService) accountHandler {
	return accountHandler{accSrv: accSrv}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	// validate header-type
	if r.Header.Get(CONTENT_TYPE_HEADER) != CONTENT_TYPE_VALUE {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	// validate invalid json format
	request := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	// sent to newAccount for creating
	response, err := h.accSrv.NewAccount(customerID, request)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set(CONTENT_TYPE_HEADER, CONTENT_TYPE_VALUE)
	json.NewEncoder(w).Encode(response)
}

func (h accountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	// handle cannot get account
	responses, err := h.accSrv.GetAccounts(customerID)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set(CONTENT_TYPE_HEADER, CONTENT_TYPE_VALUE)
	json.NewEncoder(w).Encode(responses)
}
