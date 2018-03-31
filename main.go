package main

import (
	"net/http"
	"fmt"
	"BMail/config"
	"github.com/gorilla/mux"
	"log"
	"BMail/ping"
)

func main() {

	configuration := config.Init("develop")

	router := mux.NewRouter()
	router.HandleFunc("/ping", ping.Handle)

	listen := fmt.Sprintf("%s:%s", configuration.Host, configuration.Port)
	fmt.Println("Listening..." + listen)

	log.Fatal(http.ListenAndServe(listen, router))
}
