package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {

	var help bool
	var port string
	var dir string

	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&help, "help", false, "show help")

	flag.StringVar(&port, "p", "8080", "port to serve from")
	flag.StringVar(&port, "port", "8080", "port to serve from")

	flag.StringVar(&dir, "d", "./", "root directory for server")
	flag.StringVar(&dir, "directory", ".", "root directory for server")

	flag.Parse()

	if help {
		usage()
		os.Exit(0)
	}

	http.Handle("/", http.FileServer(http.Dir(dir)))

	fmt.Print("Starting local server\n")
	fmt.Printf("Port: %s\n", port)
	fmt.Printf("Root directory: %s\n", dir)
	fmt.Println()

	fmt.Printf("Visit \033[34m\033[4mhttp://localhost:%v/\033[m\n", port) // blue color "\033[34m" and underline "\033[4m", then reset "\033[m"

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not start server: %v\n", err)
	}

}

func usage() {
	fmt.Printf("Usage: %s [-p port] [-d directory]\n", os.Args[0])
}
