package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	// // root route
	// mux.HandleFunc("/", handlers.RootHandler)

	// // teacher route
	// mux.HandleFunc("/teachers/", handlers.TeachersHandler)

	// // student route
	// mux.HandleFunc("/students/", handlers.StudentsHandler)

	// // executive routes
	// mux.HandleFunc("/execs/", handlers.ExecsHandler)

	mux.HandleFunc("GET /teachers/", handlers.GetTeachersHandler)
	mux.HandleFunc("POST /teachers/", handlers.AddTeacherHandler)
	mux.HandleFunc("PUT /teachers/", handlers.UpdateTeacherhandler)
	mux.HandleFunc("PATCH /teachers/", handlers.PatchTeachersHandler)
	mux.HandleFunc("DELETE /teachers/", handlers.DeleteTeachersHandler)

	mux.HandleFunc("GET /teachers/{id}", handlers.GetOneTeacherHandler)
	mux.HandleFunc("PATCH /teachers/{id}", handlers.PatchOneTeacherHandler)
	mux.HandleFunc("DELETE /teachers/{id}", handlers.DeleteOneTeacherHandler)

	return mux
}
