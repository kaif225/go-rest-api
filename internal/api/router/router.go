package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	// root route
	mux.HandleFunc("/", handlers.RootHandler)

	// teacher route
	mux.HandleFunc("/teachers/", handlers.TeachersHandler)

	// student route
	mux.HandleFunc("/students/", handlers.StudentsHandler)

	// executive routes
	mux.HandleFunc("/execs/", handlers.ExecsHandler)

	return mux
}
