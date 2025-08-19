/*
In router.go file modify the mux part like this

*/

package main

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /teachers/", handlers.TeachersHandler)
	mux.HandleFunc("GET /teachers/{id}", handlers.TeachersHandler)
	mux.HandleFunc("POST /teachers/", handlers.TeachersHandler)
	mux.HandleFunc("PUT /teachers/", handlers.TeachersHandler)
	mux.HandleFunc("PATCH /teachers/", handlers.TeachersHandler)
	mux.HandleFunc("PATCH /teachers/{id}", handlers.TeachersHandler)
	mux.HandleFunc("DELETE /teachers/", handlers.TeachersHandler)
	mux.HandleFunc("DELETE /teachers/{id}", handlers.TeachersHandler)

	return mux
}
