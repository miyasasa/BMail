package bmail

import (
	"net/http"
	"encoding/json"
	"BMail/response"
)

func Send(w http.ResponseWriter, r *http.Request) {

	var mail Bmail
	err := json.NewDecoder(r.Body).Decode(&mail)

	if err != nil {
		w.Write([]byte("Invalid B-mail type"))
	}

	response.JsonResponse(w, mail)
	defer r.Body.Close()
}
