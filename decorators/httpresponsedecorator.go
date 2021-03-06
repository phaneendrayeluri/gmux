package decorators

import (
	"encoding/json"
	"net/http"
)

// ResponsdAsJSON - Adds a content-type json
func ResponsdAsJSON(s int, resp interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bees, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(s)
		w.Write(bees)
	}
}
