//get server working
// get response for random numbers
// write algorithm - done
// write database save methods - done
// write unit tests

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
	// w.Write([]byte("fib("+nString+") = (v2)" + fibout))
	jsonfile := json.NewEncoder(w).Encode(fibout)
	log.Println(jsonfile)
	log.Println(nString)
	log.Println(fibout)
	// w.Write([]byte(json.NewEncoder(w).Encode(fibout)))
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/fib/{n}", fibHandler)
	log.Println("Listening.........")
	http.ListenAndServe(":8080",router)
}

