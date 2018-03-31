package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
	"time"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
	info := fmt.Sprintf("{ Time: %s }", time.Now().Format(time.RFC3339))
	w.Write([]byte(info))
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/ping", handlePing)

	fmt.Println("Listining Localhost:8080 port")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
