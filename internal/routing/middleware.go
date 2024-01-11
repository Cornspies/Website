package routing

import (
	"Website/config"
	"log"
	"net/http"
)

func applyMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if config.AdvancedLogging {
			log.Println(r.Host + r.URL.Path)
		}
		handlerFunc(w, r)
	}
}
