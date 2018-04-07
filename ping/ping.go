package ping

import (
	"time"
	"net/http"
	"BMail/response"
)

func Pong(w http.ResponseWriter, _ *http.Request) {
	response.JsonResponse(w, time.Now().Format(time.RFC3339))
}
