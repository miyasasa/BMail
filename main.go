package main

import (
	"fmt"
	"log"
	"net/http"
	"BMail/ping"
	"BMail/config"
	"github.com/gorilla/mux"
)

func main() {

	configuration := config.Init("develop")

	router := mux.NewRouter()
	router.HandleFunc("/ping", ping.Handle)

	listen := fmt.Sprintf("%s:%s", configuration.Host, configuration.Port)
	log.Println("Listening..." + listen)

	log.Fatal(http.ListenAndServe(listen, router))
}
