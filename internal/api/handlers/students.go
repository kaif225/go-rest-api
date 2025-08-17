package handlers

import (
	"fmt"
	"net/http"
)

func StudentsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on  Students route"))
		fmt.Println("Hello GET method on  Students route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on  Students route"))
		fmt.Println("Hello PUT method on  Students route")
	case http.MethodPost:
		w.Write([]byte("Hello POST method on  Students route"))
		fmt.Println("Hello POST method on  Students route")

	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on  Students route"))
		fmt.Println("Hello PATCH method on  Students route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on  Students route"))
		fmt.Println("Hello DELETE method on  Students route")
	}
}
