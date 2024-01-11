package templates

import (
	"Website/config"
	"Website/internal/user"
	"html/template"
	"log"
	"net/http"
)

var BaseTemplate *template.Template
var NotFoundTemplate *template.Template

func InitTemplates() {
	var path = config.WorkingDirectory + "/internal/templates/html/"
	BaseTemplate = template.Must(template.ParseFiles(path+"head.html", path+"nav.html", path+"footer.html", path+"baseTemplate.html"))
	NotFoundTemplate = template.Must(template.ParseFiles(path+"head.html", path+"nav.html", path+"footer.html", path+"404.html"))
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request, user user.User) {
	var page = Page{
		Title:           "Site not found",
		CurrentLocation: LocationNone,
	}
	w.WriteHeader(404)
	if err := NotFoundTemplate.ExecuteTemplate(w, "404.html", M{
		"Page":        page,
		"User":        user,
		"URLNotFound": r.Host + r.URL.Path,
	}); err != nil {
		log.Fatal(err)
	}
}
