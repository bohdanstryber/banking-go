package app

import (
	"github.com/bohdanstryber/banking-go/dto"
	"github.com/bohdanstryber/banking-go/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appErr := h.service.NewAccount(request)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func (h AccountHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["id"]
	accountId := vars["account_id"]

	var request dto.NewTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != 	nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		request.CustomerId = customerId
	}

	account, appErr := h.service.NewTransaction(request)

	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, account)
	}
}

