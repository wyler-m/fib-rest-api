//get server working
// get response for random numbers
// write algorithm - done
// write database save methods - done
// write unit tests

package main

import ("github.com/wyler-m/fibbalgo"
		"log"
		"net/http"
		"math/big"
		"github.com/gorilla/mux"
		// "fmt"
		)

// func timeHandler(w http.ResponseWriter, r *http.Request) {
//   n := new(big.Int).SetUint64(20)
//   fib := fibbalgo.TwokFib(n)
//   w.Write([]byte("fib("+n.String()+") = " +fib.String()))
// }


func fibHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    nString := vars["n"]
    n := new(big.Int)
    n.SetString(nString,10)
    fib := fibbalgo.TwokFib(n)
    w.Write([]byte("fib("+nString+") = " + fib.String()))
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/fib/{n}", fibHandler)
	log.Println("Listening.........")
	http.ListenAndServe(":8080",router)
}

