// This is the name of the packge
// Everything with this package can see everything
// else inside the same package, regardless of the file they are in
package main

//These are the libraries we are going to use
//Both "fmt" and "net" are part of the Go standard library
import (
	"fmt"
	"net/http"
)

func main() {
	// The "HandleFunc" method accepts a path and a function as arguments
	// (yes, we can pass functions as arguements, and even treat them like variables in Go)
	// however, the handler function has to have the appropriate signature (as described by the "handler" function below)")
	http.HandleFunc("/", handler)

	// after defining our server, we finally "listen and serve" on port 8080
	// the second argument is the handler, which we will come to later on, but for now it is left as nil,
	// and the handler defined above in HandleFunc is used
	http.ListenAndServe(":8080", nil)
}
// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func handler(w http.ResponseWriter, r *http.Request) {
	//For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World!")
}