/*
 // Basic Middleware Skeleton

func securityHeader(next http.Handler) http.Handler { // the return is because when it returns the http handler that means it is running
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

*/

/*
1 - X-DNS-Prefetch-Control : disable dns request prefetch , so that we are safe from DNS related attacks
2. X-Frame-Options: //This header prevents the web page from being displayed in an iframe on other websites
3- X-XSS-Protection:  This enable cross-site scripting filter build into most moderb web browsers and instructs the browser to block
   the page if an XSS attack is detected internally.
4. X-Content-Type-Options : prevents browsers from Mine sniffing from Mime a response away from the declared content type .
   MIME (Multipurpose Internet Mail Extensions) refers to a standard that identifies the type of data being transmitted, such as HTML, images, or other files

5. Strict-Transport-Security: to tell browser to open the pages over https only

6. X-Powered-By : prevent from knowling the backend technology
*/

package middlewares

import (
	"fmt"
	"net/http"
)

// this is going to accept the handler as their argument
func SecurityHeader(next http.Handler) http.Handler { // the return is because when it returns the http handler that means it is running
	fmt.Println("Security Header Middleware ")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Security-Header Middleware being return ")
		w.Header().Set("X-DNS-Prefetch-Control", "off")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1;mode=block")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Strict-Transport-Security", "max-age=6307000;include:SubDomains;preload")
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("Referred-Policy", "no-referrer")
		w.Header().Set("X-Powered-By", "Django")
		w.Header().Set("Server", "")
		w.Header().Set("X-Permitted-Cross-Domain-Policies", "none")
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		w.Header().Set("Permission-Policy", "geolocation=(self), microphone=()")
		next.ServeHTTP(w, r)
		fmt.Println("Security-Header Middleware End ")
	})
}
