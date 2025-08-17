package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func RequestTimeMiddleware(next http.Handler) http.Handler {
	fmt.Println("Request time Middleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request Time Middleware being return")
		fmt.Println("Received Request in ResponseTime")
		start := time.Now()

		// custom responseWriter to capture teh status code
		wrappedWritter := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		// calculate the duration

		duration := time.Since(start)

		w.Header().Set("X-Response-Time", duration.String()) // to ge the response time in header
		next.ServeHTTP(wrappedWritter, r)
		fmt.Println("Ratelimiter Middleware End")
		// Log the request details
		fmt.Printf("Method : %s, URL : %s, Status : %v, Duration : %v \n", r.Method, r.URL, wrappedWritter.status, duration.String())
		fmt.Println("Send Request from Response time Middleware")
	})
}

// for response writer
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code

	rw.ResponseWriter.WriteHeader(code)
}

/*
Code from 32 to 42 - how the instructor come with this Idea

when you hover over  http.ResponseWriter - we see REsponseWriter is an interface
that takes WriteHandler of type int

*/
