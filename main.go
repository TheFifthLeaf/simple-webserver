package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Initialize the server
	fileServer := http.FileServer(http.Dir("./static"))

	// Set handlers
	http.Handle("/", fileServer)

}
