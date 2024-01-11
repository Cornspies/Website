package main

import (
	"Website/config"
	"Website/internal/routing"
	"Website/internal/templates"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	advancedLoggingFlag := flag.Bool("advancedLogging", false, "enables advanced logging")
	flag.Parse()
	config.AdvancedLogging = *advancedLoggingFlag

	getwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	config.WorkingDirectory = getwd
	templates.InitTemplates()

	_, err = os.Stat("cert.pem")
	_, err1 := os.Stat("key.pem")
	if errors.Is(err, os.ErrNotExist) || errors.Is(err1, os.ErrNotExist) {
		log.Println("cert and/or key not found... defaulting to http")
		log.Fatal(http.ListenAndServe(":80", routing.GetServeMux()))
	}
	log.Println("cert and key found... using https")
	go log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(redirectToHttps)))
	log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", routing.GetServeMux()))
}

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	target := "https://" + r.Host + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	if config.AdvancedLogging {
		log.Printf("redirect to: %s", target)
	}
	http.Redirect(w, r, target, http.StatusPermanentRedirect)
}
