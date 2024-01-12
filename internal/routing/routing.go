package routing

import (
	"Website/config"
	"Website/internal/templates"
	"Website/internal/user"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var staticRegex = regexp.MustCompile("css|img|js")

func GetServeMux() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/logout", applyMiddleware(user.Logout))
	serveMux.HandleFunc("/register", applyMiddleware(user.HandleRegister))
	serveMux.HandleFunc("/login", applyMiddleware(user.HandleLogin))
	serveMux.HandleFunc("/", applyMiddleware(handleRequest))
	return serveMux
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var requestURLParts = strings.Split(r.URL.Path, "/")
	switch strings.ToLower(requestURLParts[1]) {
	case "favicon.ico":
		http.Redirect(w, r, "/static/img/favicon.png", http.StatusTemporaryRedirect)
		break
	case "static":
		if len(requestURLParts) < 3 || !staticRegex.MatchString(requestURLParts[2]) {
			user.NotFoundResponse(w, r)
			return
		}
		//TODO: sanitize user input
		http.ServeFile(w, r, strings.TrimPrefix(config.WorkingDirectory+r.URL.Path, "/"))
		break
	case "":
		var userData = user.GetUserData(r)
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
		user.NotFoundResponse(w, r)
		break
	}
}
