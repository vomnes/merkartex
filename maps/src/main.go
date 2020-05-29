package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	libHTTP "github.com/vomnes/go-library/http"
	libPretty "github.com/vomnes/go-library/pretty"
)

// HandleAPIRoutes instantiates and populates the router
func handleAPIRoutes() *mux.Router {
	// instantiating the router
	api := mux.NewRouter()

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		libHTTP.RespondWithJSON(w, http.StatusOK, "OK")
	})
	return api
}

func main() {
	portPtr := flag.String("port", "3000", "port your want to listen on")
	flag.Parse()
	if *portPtr != "" {
		fmt.Printf("running on port: %s\n", *portPtr)
	}
	router := handleAPIRoutes()
	enhancedRouter := enhanceHandlers(router)
	if err := http.ListenAndServe(":"+*portPtr, enhancedRouter); err != nil {
		log.Fatal(libPretty.Error(err.Error()))
	}
}
