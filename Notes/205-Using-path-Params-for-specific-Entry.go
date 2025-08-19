/*
In Handlers folder in teachers.go Capitalize Add the handler like GetHandler, AddHandler ... etc

Then in updated router.go updated the handler based on the Request method

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

// Modified teachers.go
