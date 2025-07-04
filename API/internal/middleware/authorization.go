package middleware

import (
	"errors"
	"net/http"

	"github.com/Rishi2600/Go-tutorial/api"
	"github.com/Rishi2600/Go-tutorial/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("invalid Username or wrong token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHanlder(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()

		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHanlder(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
