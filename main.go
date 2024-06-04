package main

import (
	"fmt"
	"net/http"

	ascii "ascii/funcs"
)

const port = ":8280"

func main() {
	// Register handler function with URL pattern "/"
	http.HandleFunc("/", ascii.Handler )
	http.HandleFunc("/ascii-art", ascii.Asciiart)

	fmt.Printf("server started on port (http://localhost%s/) :\n", port)
	fmt.Println(http.ListenAndServe(port, nil))
}
