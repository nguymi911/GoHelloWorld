package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/test", HelloServer)
	http.HandleFunc("/", DefaultPath)
	http.ListenAndServe(":8080", nil)
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func DefaultPath(w http.ResponseWriter, r *http.Request) {
	sd := time.Date(2022, 2, 24, 0, 0, 0, 0, time.Now().UTC().Location())
	fmt.Fprintf(w, "Heo <3 Banh Bao, day %v and counting!", int(time.Now().Sub(sd).Hours()/24))
}
