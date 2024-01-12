package user

import (
	"Website/internal/templates"
	"log"
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    "SESSION",
		Value:   "",
		Expires: time.Now(),
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var user = GetUserData(r)
	if user.IsLoggedIn {
		http.Redirect(w, r, "/profile/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		var err error
		if err = r.ParseForm(); err == nil {
			if r.Form.Has("username") && r.Form.Has("password") {
				if err = login(w, r.Form.Get("username"), r.Form.Get("password")); err == nil {
					http.Redirect(w, r, "/profile/", http.StatusSeeOther)
					return
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		var page = templates.Page{
			Title:           "Login",
			CurrentLocation: templates.LocationProfile,
		}
		if err = templates.LoginPage.ExecuteTemplate(w, "login.html", templates.M{
			"Page":  page,
			"User":  user,
			"Error": err,
		}); err != nil {
			log.Fatal(err)
		}
		return
	} else {
		var page = templates.Page{
			Title:           "Login",
			CurrentLocation: templates.LocationProfile,
		}
		if err := templates.LoginPage.ExecuteTemplate(w, "login.html", templates.M{
			"Page": page,
			"User": user,
		}); err != nil {
			log.Fatal(err)
		}
	}
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var user = GetUserData(r)
	if user.IsLoggedIn {
		http.Redirect(w, r, "/profile/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		var err error
		if err = r.ParseForm(); err == nil {
			if r.Form.Has("username") && r.Form.Has("email") && r.Form.Has("password") && r.Form.Has("repeatPassword") {
				if err = register(r.Form.Get("username"), r.Form.Get("email"), r.Form.Get("password"), r.Form.Get("repeatPassword")); err == nil {
					http.Redirect(w, r, "/login", http.StatusSeeOther)
					return
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		var page = templates.Page{
			Title:           "Register",
			CurrentLocation: templates.LocationProfile,
		}
		if err = templates.RegisterPage.ExecuteTemplate(w, "register.html", templates.M{
			"Page":  page,
			"User":  user,
			"Error": err,
		}); err != nil {
			log.Fatal(err)
		}
	} else {
		var page = templates.Page{
			Title:           "Register",
			CurrentLocation: templates.LocationProfile,
		}
		if err := templates.RegisterPage.ExecuteTemplate(w, "register.html", templates.M{
			"Page": page,
			"User": user,
		}); err != nil {
			log.Fatal(err)
		}
	}
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	var user = GetUserData(r)
	var page = templates.Page{
		Title:           "Site not found",
		CurrentLocation: templates.LocationNone,
	}
	w.WriteHeader(404)
	if err := templates.NotFoundPage.ExecuteTemplate(w, "404.html", templates.M{
		"Page":        page,
		"User":        user,
		"URLNotFound": r.Host + r.URL.Path,
	}); err != nil {
		log.Fatal(err)
	}
}
