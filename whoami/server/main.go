package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LEGO/ip-verification/handlers"
)


func main() {
	http.HandleFunc("/ip", handlers.Handler)
	http.HandleFunc("/test", handlers.TestHandler)

	kubernetesClient()


	port := os.Getenv("PORT")  
	if port == "" {  
		port = "8888" // Default port if not specified
	}
	fmt.Printf("Server is starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
