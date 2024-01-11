package routing

import (
	"Website/config"
	"Website/internal/templates"
	"Website/internal/user"
	"log"
	"net/http"
	"strings"
)

func GetServeMux() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", applyMiddleware(handleRequest))
	return serveMux
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var requestURLParts = strings.Split(r.URL.Path, "/")
	var userData = user.GetUserData()
	switch strings.ToLower(requestURLParts[1]) {
	case "favicon.ico":
		http.Redirect(w, r, "/static/img/favicon.png", http.StatusTemporaryRedirect)
		break
	case "static":
		//TODO: potentially unsafe, needs input sanitation
		http.ServeFile(w, r, strings.TrimPrefix(config.WorkingDirectory+r.URL.Path, "/"))
		break
	case "":
		w.WriteHeader(200)
		//TODO: get Homepage template
		var page = templates.Page{
			Title:           "Homepage",
			CurrentLocation: templates.LocationHome,
		}
		if err := templates.BaseTemplate.ExecuteTemplate(w, "baseTemplate.html", templates.M{
			"Page": page,
			"User": userData,
		}); err != nil {
			log.Fatal(err)
		}
		break
	default:
		templates.NotFoundResponse(w, r, userData)
		break
	}
}
