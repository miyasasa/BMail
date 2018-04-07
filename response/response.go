package response

import (
	"net/http"
	"encoding/json"
)

func JsonResponse(w http.ResponseWriter, model interface{}) {
	res, _ := json.Marshal(model)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
