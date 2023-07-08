package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/test", HelloServer)
	http.HandleFunc("/", DefaultPath)
	http.HandleFunc("/appreciations", Appreciation)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error spinning up server")
	}

}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func DefaultPath(w http.ResponseWriter, r *http.Request) {
	sd := time.Date(2022, 2, 24, 0, 0, 0, 0, time.Now().UTC().Location())
	fmt.Fprintf(w, "Heo <3 Banh Bao, day %v and counting!", int(time.Since(sd).Hours()/24))
}

func Appreciation(w http.ResponseWriter, r *http.Request) {
	t := []string{"Understanding", "Loving", "Caring", "Sympathizing", "Collaborating", "Patient",
		"Family-orienting", "Resilient", "Strong-willed", "Honest"}
	fmt.Fprintf(w, "%v", t[rand.Intn(len(t))])
}
