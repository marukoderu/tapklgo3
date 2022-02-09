package AuthMiddleware

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		session, _ := store.Get(request, "mysession")
		email := session.Values["email"]
		if email == nil {
			http.Redirect(response, request, "/login", http.StatusSeeOther)
		} else {
			h.ServeHTTP(response, request)
		}
	})
}
