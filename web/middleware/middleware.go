package middleware

import (
	"github.com/firstrow/logvoyage/web/render"
	"github.com/martini-contrib/sessions"
)

// Check user authentication
func Authorize(r *render.Render, sess sessions.Session) {
	email := sess.Get("email")
	if email == nil {
		r.Redirect("/login")
	}
}

// Redirect user to Dashboard if authorized
func RedirectIfAuthorized(r *render.Render, sess sessions.Session) {
	email := sess.Get("email")
	if email != nil {
		r.Redirect("/dashboard")
	}
}
