package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.RootHandler)
	// // root router
	// mux.HandleFunc("/", handlers.RootHandler)

	// // teacher route
	// mux.HandleFunc("/teachers/", handlers.TeachersHandler)

	// // student route
	// mux.HandleFunc("/students/", handlers.StudentsHandler)

	// // executive routes
	// mux.HandleFunc("/execs/", handlers.ExecsHandler)

	mux.HandleFunc("GET /teachers", handlers.GetTeachersHandler)
	mux.HandleFunc("POST /teachers", handlers.AddTeacherHandler)
	mux.HandleFunc("PUT /teachers", handlers.UpdateTeacherhandler)
	mux.HandleFunc("PATCH /teachers", handlers.PatchTeachersHandler)
	mux.HandleFunc("DELETE /teachers", handlers.DeleteTeachersHandler)

	mux.HandleFunc("GET /teachers/{id}", handlers.GetOneTeacherHandler)
	mux.HandleFunc("PATCH /teachers/{id}", handlers.PatchOneTeacherHandler)
	mux.HandleFunc("DELETE /teachers/{id}", handlers.DeleteOneTeacherHandler)
	/////////////// For student //////

	mux.HandleFunc("GET /students", handlers.GetStudentsHandler)
	mux.HandleFunc("POST /students", handlers.AddStudentsHandler)
	mux.HandleFunc("PUT /students", handlers.UpdateStudentshandler)
	mux.HandleFunc("PATCH /students", handlers.PatchStudentsHandler)
	mux.HandleFunc("DELETE /students", handlers.DeleteStudentsHandler)

	mux.HandleFunc("GET /students/{id}", handlers.GetOneStudentHandler)
	mux.HandleFunc("PATCH /students/{id}", handlers.PatchOneStudentHandler)
	mux.HandleFunc("DELETE /students/{id}", handlers.DeleteOneStudentHandler)
	return mux
}
