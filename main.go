package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	bind = flag.String("bind", os.Getenv("PORT"), "http server binding")
)

func main() {
	flag.Parse()
	var port string
	if *bind == "" {
		port = "8000"
	} else {
		port = *bind
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("version 1"))
		return
	})

	fmt.Println("App serving on the port", port)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(err)
}
