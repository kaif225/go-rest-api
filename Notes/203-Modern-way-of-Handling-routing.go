/*
var students []models.Student for rows.Next() { var s models.Student rows.Scan(&s.ID, &s.FirstName, &s.LastName, &s.Title) students = append(students, s) } in clear words we define students array of slice of type Student struct, now row.Next loop over all rows then it stroes the row value in var s and
 then we appends that s value to students var that is of tyep map of struct ?


*/

package main

import (
	"fmt"
	"net/http"
)

func modernroute() {
	mux := http.NewServeMux()
	//Method based routing

	mux.HandleFunc("POST /items/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from post")
	})

	mux.HandleFunc("DELETE /items/delete", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Item deleted")
	})

	// Wildcard in patterns  - path patterns

	mux.HandleFunc("GET /teacher/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Teacher id : %v", r.PathValue("id"))
	})

	// Wildcard with "..."

	mux.HandleFunc("GET /files/{path...}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Path is : %v", r.PathValue("path"))
	})
	err := http.ListenAndServe(":8097", mux)
	if err != nil {
		fmt.Println(err)
		return
	}

}
