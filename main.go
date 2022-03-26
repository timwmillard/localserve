package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {

	port := flag.String("p", "6000", "port to serve from")
	dir := flag.String("d", ".", "root directory for server")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))

	fmt.Print("Starting local server\n")
	fmt.Printf("Port: %s\n", *port)
	fmt.Printf("Root directory: %s\n", *dir)
	fmt.Println()

	fmt.Printf("Visit \033[34m\033[4mhttp://localhost:%v/\033[m\n", *port) // blue color "\033[34m" and underline "\033[4m", then reset "\033[m"

	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not start server: %v\n", err)
	}
}
