package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (app *application) routes() http.Handler {
	//mux := http.NewServeMux()
	mux := pat.New()

	mux.Get("/", app.session.Enable(noSurf(app.authenticate(http.HandlerFunc(app.home)))))

	mux.Get("/post/create", app.session.Enable(noSurf(app.authenticate(app.requireAuthenticatedUser(http.HandlerFunc(app.createPostForm))))))
	mux.Post("/post/create", app.session.Enable(noSurf(app.authenticate(app.requireAuthenticatedUser(http.HandlerFunc(app.createPost))))))
	mux.Get("/post/:id", app.session.Enable(noSurf(app.authenticate(http.HandlerFunc(app.showPost)))))

	mux.Get("/user/signup", app.session.Enable(noSurf(app.authenticate(http.HandlerFunc(app.signupUserForm)))))
	mux.Post("/user/signup", app.session.Enable(noSurf(app.authenticate(http.HandlerFunc(app.signupUser)))))
	mux.Get("/user/login", app.session.Enable(noSurf(app.authenticate(http.HandlerFunc(app.loginUserForm)))))
	mux.Post("/user/login", app.session.Enable(noSurf(app.authenticate(http.HandlerFunc(app.loginUser)))))
	mux.Post("/user/logout", app.session.Enable(noSurf(app.authenticate(app.requireAuthenticatedUser(http.HandlerFunc(app.logoutUser))))))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
