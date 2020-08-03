package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := getPORT()
	fmt.Printf("serve on %s\n", port)
	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir("./public"))))
}

func getPORT() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	port = ":" + port
	return port
}
