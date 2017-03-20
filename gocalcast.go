package gocalcast

import (
	"flag"
	"log"
	"net/http"
)

var port = flag.String("port", ":1801", "http server address")

func New() {
	flag.Parse()
	err := http.ListenAndServe(*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
