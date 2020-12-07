//
// To start the server:
//
//   $ go run simple_server_mux.go
//
// Do the curls on port 9000
//
//   $ curl localhost:9000/home/subscriptions
//   $ curl localhost:9000
//   404 page not found
//   $ curl localhost:9000/home
//   endpoint /home
//   $ curl localhost:9000/home/subscription 
//   endpoint /home/subscriptions


package main

import (
	"fmt"
	"net/http"
)

func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("/home",func (res http.ResponseWriter, req *http.Request){

		fmt.Fprint(res,"endpoint /home\n")
	})

	mux.HandleFunc("/home/accounts",func (res http.ResponseWriter, req *http.Request){

		fmt.Fprint(res,"endpoint /home/accounts\n")
	})

	mux.HandleFunc("/home/subscription",func (res http.ResponseWriter, req *http.Request){

		fmt.Fprint(res,"endpoint /home/subscriptions\n")
	})

	http.ListenAndServe(":9000",mux)
}

