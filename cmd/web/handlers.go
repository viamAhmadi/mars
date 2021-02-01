package main

import (
	"fmt"
	"github.com/viamAhmadi/mars/pkg/forms"
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
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
	// The check of r.Method !="POST" is now superfluous and can be removed
	//if r.Method != "POST" {
	//	w.Header().Set("Allow", "POST")
	//	app.clientError(w, http.StatusMethodNotAllowed)
	//	return
	//}
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}

	//title := r.PostForm.Get("title")
	//content := r.PostForm.Get("content")
	//expires := r.PostForm.Get("expires")

	id, err := app.posts.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", "Post successfully created!")

	//http.Redirect(w, r, fmt.Sprintf("/post?id=%d", id), http.StatusSeeOther)
	http.Redirect(w, r, fmt.Sprintf("/post/%d", id), http.StatusSeeOther)
}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MatchPattern("email", forms.EmailRx)
	form.MinLength("password", 10)
	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{
			Form: form,
		})
		return
	}

	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
	if err == models.ErrDuplicateEmail {
		form.Errors.Add("email", "Address is already in use")
		app.render(w, r,
			"signup.page.tmpl", &templateData{
				Form: form,
			})
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", "Your signup was successfully. Please log in.")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {

}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {

}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {

}