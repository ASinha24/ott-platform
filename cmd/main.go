package main

import (
	"flag"
	"log"

	gohttp "net/http"

	"github.com/gorilla/mux"

	ott "github.com/asinha24/ott-platform"
	"github.com/asinha24/ott-platform/http"
	"github.com/asinha24/ott-platform/movies"
)

var port = flag.String("port", "8080", "port to listen")

func main() {
	flag.Parse()
	router := mux.NewRouter()

	ottPlatform := movies.NewOTTPlatform()
	ottService := ott.NewOttService(movies.NewMovieStore(), ottPlatform)
	ottHandler := http.NewOTTHandler(ottService, ottPlatform)
	ottHandler.IntallRouter(router)

	log.Println("starting http server, listening on port:", *port)
	if err := gohttp.ListenAndServe(":"+*port, router); err != nil {
		log.Fatalf("error in starting server: %v", err)
	}

}
