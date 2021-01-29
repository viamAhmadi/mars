package main

import (
	"fmt"
	"github.com/viamAhmadi/mars/pkg/models"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Pat matches the / path
	//if r.URL.Path != "/" {
	//	app.notFound(w)
	//	return
	//}

	p, err := app.posts.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{Posts: p})
}

func (app *application) showPost(w http.ResponseWriter, r *http.Request) {
	//id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// Pat doesn't strip the colon from the named capture key
	// we need get the value of :id from the query string instead id
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	p, err := app.posts.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Post: p,
	})
	//fmt.Fprintf(w, "%v", p)
}

func (app *application) createPostForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a now post..."))
}

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
	// The check of r.Method !="POST" is now superfluous and can be removed
	//if r.Method != "POST" {
	//	w.Header().Set("Allow", "POST")
	//	app.clientError(w, http.StatusMethodNotAllowed)
	//	return
	//}

	title := "title"
	content := "content"
	expires := "?"

	id, err := app.posts.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	//http.Redirect(w, r, fmt.Sprintf("/post?id=%d", id), http.StatusSeeOther)
	http.Redirect(w, r, fmt.Sprintf("/post/%d", id), http.StatusSeeOther)
}
