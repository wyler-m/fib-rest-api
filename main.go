//get server working - done
// get response for random numbers - done
// write algorithm - done
// write unit tests
// logging
// write error handleing

package main

import (
		"encoding/json"
		"log"
		"net/http"
		"github.com/gorilla/mux"
		)

func fibHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nString := vars["n"]
	fibout := fiboutput{ N: nString, Nthfibonacci: fib_processer(nString)}
	json.NewEncoder(w).Encode(fibout)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/fib/{n}", fibHandler)
	log.Println("Listening.........")
	http.ListenAndServe(":8080",router)
}

