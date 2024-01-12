package templates

import (
	"Website/config"
	"html/template"
)

var BaseTemplate *template.Template
var NotFoundPage *template.Template
var LoginPage *template.Template
var RegisterPage *template.Template

func InitTemplates() {
	var path = config.WorkingDirectory + "/static/html/"
	BaseTemplate = template.Must(template.ParseFiles(path+"head.html", path+"nav.html", path+"footer.html", path+"baseTemplate.html"))
	NotFoundPage = template.Must(template.ParseFiles(path+"head.html", path+"nav.html", path+"footer.html", path+"404.html"))
	LoginPage = template.Must(template.ParseFiles(path+"head.html", path+"nav.html", path+"footer.html", path+"login.html"))
	RegisterPage = template.Must(template.ParseFiles(path+"head.html", path+"nav.html", path+"footer.html", path+"register.html"))
}
