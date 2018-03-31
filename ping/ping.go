package ping

import (
	"fmt"
	"time"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	info := fmt.Sprintf("{ Time: %s }", time.Now().Format(time.RFC3339))
	w.Write([]byte(info))
}
