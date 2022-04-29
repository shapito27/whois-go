package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

const (
	addr = ":9091"
	whoisRoute = "/whois/plain/{domain}"
)

// main runs server on fiven ip and port
func main() {
	// setup routing
	router := mux.NewRouter()
	router.HandleFunc(whoisRoute, getWhoisPlain)
	http.Handle("/", router)

	err := http.ListenAndServe(addr, nil) 
	if err != nil {
		fmt.Println("\n\n", err)
	}
}

// getWhoisPlain outputs whois information by given domain name and port.
// Example: /whois/plain/google.com?port=9999
func getWhoisPlain(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintf(w, "Only GET requests are allowed")
		return
	}

	vars := mux.Vars(r)
	var err error
	var out []byte

	port := r.URL.Query().Get("port")

	if port != "" {
		out, err = exec.Command("whois", "-H", vars["domain"], "-p", port).Output()
	} else {
		out, err = exec.Command("whois", "-H", vars["domain"]).Output()
	}

	checkError(err, "Whois failed. Err:")

	// send data to client side
	fmt.Fprintf(w, "%s", string(out))
}

// checkError check if error and output it
func checkError(err error, mes string) {
	if err != nil {
		fmt.Println(mes, err)
	}
}
