package middlewares

import (
	"errors"
	"net/http"

	"github.com/ihsanbhee/mu-golang/api/auth"
	"github.com/ihsanbhee/mu-golang/api/responses"
)

func SetMiddlewareJSON( next http.HandlerFunc ) http.HandlerFunc {
	return func( w http.ResponseWriter, r *http.Request ) {
		w.Header()
			.Set( "Content-Type", "application/json" )

			next( w, r )
	}
}

func SetMiddlewareAuthentication( next http.HandlerFunc ) http.HandlerFunc {
	return func( w http.ResponseWriter, r *http.Request ) {
		err := auth.TokenValid( r )

		if err != nil {
			response.ERROR( w, http.StatusUnauthorized, errors.New( "Unauthorized" ) )
			return
		}

		next( w, r )
	}
}