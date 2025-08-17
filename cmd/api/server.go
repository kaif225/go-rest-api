package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/router"
	"restapi/internal/repository/sqlconnect"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		return
	}
	_, err = sqlconnect.ConnectDb()

	if err != nil {
		fmt.Println("Error --- ", err)
	}
	port := os.Getenv("API_PORT")

	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	/*
		hppOptions := mw.HPPOptions{
			CheckQuery:                  true,
			CheckBody:                   true,
			CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
			Whitelist:                   []string{"allowedParam"},
		}
	*/
	//rl := mw.NewRateLimitor(5, time.Minute)

	//secureMux := mw.Cors(rl.Middleware(mw.RequestTimeMiddleware(mw.SecurityHeader(mw.Compression(mw.Hpp(hppOptions)(mux))))))
	//secureMux := applyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeader, mw.RequestTimeMiddleware, rl.Middleware, mw.Cors)

	router := router.Router()
	secureMux := mw.SecurityHeader(router)
	server := &http.Server{
		Addr: port,
		//Handler:   middlewares.SecurityHeader(mux),
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port ", port)
	err = server.ListenAndServeTLS(cert, key)
	//err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting server :", err)
	}

}
