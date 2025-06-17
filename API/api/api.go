package api

import (
	"encoding/json"
	"net/http"
)

/*few structs to represent parameters for endpoints and responses*/

// coin balance params
type CoinBalanceParams struct {
	Username string
}

// coin balance response
type CoinBalanceResponse struct {
	//status code
	Code int

	//account balance
	Balance int64
}

// error response
type Error struct {
	//status code
	Code int

	//error message
	Message string
}

func WriteError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHanlder = func(w http.ResponseWriter, err error) {
		WriteError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		WriteError(w, "An unexpected error occured.", http.StatusInternalServerError)
	}
)
