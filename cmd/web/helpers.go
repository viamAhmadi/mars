package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exist", name))
		return
	}

	//Catching Runtime Errors
	buf := new(bytes.Buffer)

	err := ts.Execute(buf, app.addDefaultDate(td, r))
	if err != nil {
		app.serverError(w, err)
	}

	buf.WriteTo(w)
}

func (app *application) addDefaultDate(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.AuthenticatedUser = app.authentication(r)
	td.CurrentYear = time.Now().Year()
	td.Flash = app.session.PopString(r, "flash")
	return td
}

func (app *application) authentication(r *http.Request) int {
	return app.session.GetInt(r, "userID")
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
