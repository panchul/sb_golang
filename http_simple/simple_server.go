//
// run it like so:
//
//   $ go run simple_server.go
//
// In another terminal:
//
//   $ curl localhost:9000
//   simple http server
//

package main

import(
	"net/http"
)


type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request){
	data := []byte("simple http server\n")
	res.Write(data)
}


func main(){

	handler:=HttpHandler{}
	http.ListenAndServe(":9000",handler)
}

