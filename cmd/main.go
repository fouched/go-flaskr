package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":9080"

func main() {

	srv := &http.Server{
		Addr:    port,
		Handler: routes(),
	}

	fmt.Printf("Starting application on %s\n", port)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}
}
