package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LEGO/novus-echo/pkg/novus-echo"
)

func main() {
	http.HandleFunc("/", novus_echo.GetIP)

	port := os.Getenv("PORT")  
	if port == "" {  
		port = "7777"  
	}
	fmt.Printf("Server is starting on port %s\n", port)  
	log.Fatal(http.ListenAndServe(":" + port, nil)) 
}
