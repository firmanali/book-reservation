package main

import (
	"fmt"
	config2 "github.com/firmanali/book-reservation/internal/config"
	handlers2 "github.com/firmanali/book-reservation/internal/handlers"
	render2 "github.com/firmanali/book-reservation/internal/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8888"

var app config2.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render2.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers2.NewRepo(&app)
	handlers2.NewHandlers(repo)

	render2.NewTemplates(&app)

	fmt.Println("Starting on port", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
