package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port, dir string

func main() {
	flag.StringVar(&port, "port", "4000", "port to serve on")
	flag.StringVar(&dir, "directory", ".", "directory to serve")
	flag.Parse()

	fmt.Printf("Serving %s on port %s.\n", dir, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port),
		http.FileServer(http.Dir(dir))))
}
