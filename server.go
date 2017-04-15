package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
)

func main() {
	var port = flag.String("p", "8080", "Set port.")
	flag.Parse()

	fmt.Println("Starting server on port " + *port + "!")
	fmt.Println("Set port with flag -p=[port]!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
		fmt.Println("/" + r.URL.Path[1:] + " was requested!")
	})

	log.Fatal(http.ListenAndServe(":" + *port, nil))

}
