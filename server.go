package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os/exec"
)

func getWhoisPlain(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
		fmt.Fprintf(w, "POST method is not allowed")
	}

	vars := mux.Vars(r)
	var err error
	var out []byte

	port:= r.URL.Query().Get("port")

	if port != "" {
	    out, err = exec.Command("whois", "-H", vars["domain"], "-p", port).Output()
	} else {
	    out, err = exec.Command("whois", "-H", vars["domain"]).Output()
	}

    checkError(err, "Whois failed. Err:")

    // send data to client side
	fmt.Fprintf(w, "%s", string(out))
}

func main() {
	//routing
	router := mux.NewRouter()
	router.HandleFunc("/whois/plain/{domain}", getWhoisPlain)
	http.Handle("/", router)

	err := http.ListenAndServe(":9091", nil) // set listen ip and port
	if err != nil {
		fmt.Println("\n\n", err)
	}
}

func checkError(err error, mes string) {
	if err != nil {
		fmt.Println(mes, err)
	}
}
