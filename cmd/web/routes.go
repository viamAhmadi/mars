package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (app *application) routes() http.Handler {
	//mux := http.NewServeMux()
	mux := pat.New()
	mux.Get("/", app.session.Enable(http.HandlerFunc(app.home)))
	mux.Get("/post/create", app.session.Enable(http.HandlerFunc(app.createPostForm)))
	mux.Post("/post/create", app.session.Enable(http.HandlerFunc(app.createPost)))
	mux.Get("/post/:id", app.session.Enable(http.HandlerFunc(app.showPost)))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
