package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/joshuabl97/chichichi/handlers"
	"github.com/joshuabl97/chichichi/middleware"
	"github.com/joshuabl97/chichichi/routers"
)

type Redirect struct {
	URL      string `json:"url"`
	Endpoint string `json:"endpoint"`
}

var redirectMap = make(map[string]string)

func main() {
	jsonFlag := flag.String("json", "", "JSON payload for creating the redirects")
	flag.Parse()

	if *jsonFlag == "" {
		log.Fatal("JSON payload is required")
	}

	var redirects []Redirect
	err := json.Unmarshal([]byte(*jsonFlag), &redirects)
	if err != nil {
		log.Fatal("Invalid JSON payload")
	}

	for _, redirect := range redirects {
		redirectMap[redirect.Endpoint] = redirect.URL
	}

	handler := &handlers.Handler{
		RedirectMap: redirectMap,
	}

	router := routers.NewRouter(handler)
	middleware := middleware.NewMiddleware(router)

	http.ListenAndServe(":8080", middleware)
}
