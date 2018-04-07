package main

import (
	"fmt"
	"log"
	"net/http"
	"BMail/ping"
	"BMail/config"
	"BMail/bmail"
	"github.com/gorilla/mux"
)

func main() {

	configuration := config.Init("develop")

	router := mux.NewRouter()
	router.HandleFunc("/ping", ping.Pong).Methods("GET")
	router.HandleFunc("/send", bmail.Send).Methods("POST")

	listen := fmt.Sprintf("%s:%s", configuration.Host, configuration.Port)
	log.Println("Listening..." + listen)

	log.Fatal(http.ListenAndServe(listen, router))
}
