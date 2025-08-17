package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func Compression(next http.Handler) http.Handler {
	fmt.Println("Compressions Middleware...")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Compressions Middleware being return ")
		// check if client accepts gzip encoding

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		// Set the response header
		w.Header().Set("Content-Encoding", "gzip") // This tells the browser: “Hey, the data I’m sending is gzip-compressed, so unzip it before showing.”

		gz := gzip.NewWriter(w)
		defer gz.Close()

		// Wrap the responseWriter
		w = &gzipResponseWriter{
			ResponseWriter: w,
			Writer:         gz,
		}

		next.ServeHTTP(w, r)
		fmt.Println("Compressions Middleware End ")
	})
}

// gzip ResponseWriter wraps http.ResponseWriter to write

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (g *gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}
