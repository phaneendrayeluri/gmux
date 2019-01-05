package decorators

import "net/http"

// TypeDecoratorJSON - Adds a content-type json
func TypeDecoratorJSON(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		f(w, r)
	}
}
