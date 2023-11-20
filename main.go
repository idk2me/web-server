package main

import (
	"fmt"
	"net/http"
)

// So main should just init the server and start accepting HTTP requests.
func main() {
	fmt.Printf("Server started at port 8080\n")
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8080", nil)

}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello, World!\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func index(w http.ResponseWriter, req *http.Request) {

}
