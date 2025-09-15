package modernroutes

import (
	"fmt"
	"net/http"
)

func ModernRoutes() {

	mux := http.NewServeMux()

	// Method based routing
	mux.HandleFunc("POST /items/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "item created")
	})

	mux.HandleFunc("DELETE /items/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "item deleted")
	})

	// wildcard in pattern -path pattern
	mux.HandleFunc("GET /teachers/90", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Teacher id: %s", r.PathValue("id"))
	})

	// wildcard with "..."
	mux.HandleFunc("/files/{path...}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Path: %s", r.PathValue("path"))
	})

	mux.HandleFunc("/path1/{param1}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Param1: %s", r.PathValue("param"))
	})

	http.ListenAndServe(":8080", mux)

}