package handlers

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Hello from Server")
	w.Write([]byte("Hello from Server")) // does the same thing as of above fmt
	fmt.Println("Hello from Server")
}
