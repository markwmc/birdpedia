// This is the name of the packge
// Everything with this package can see everything
// else inside the same package, regardless of the file they are in
package main

//These are the libraries we are going to use
//Both "fmt" and "net" are part of the Go standard library
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//the new router function creates the router and
// returns it to us. We can now use this function
// to instatiate and test the router outside of the main function

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	//Declare the static file directory and point it to the
	//directory we just made

	staticFileDirectory := http.Dir("./assets/")

	// Declare the handler, that routes requests to their respective filename.
	//the fileserver is wrapped in the 'stripPrefix  method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// for example, if we type "/assets/index.html" in the browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// if we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html" and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// the "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r
}

func main() {
	// The "HandleFunc" method accepts a path and a function as arguments
	// (yes, we can pass functions as arguements, and even treat them like variables in Go)
	// however, the handler function has to have the appropriate signature (as described by the "handler" function below)")
	//http.HandleFunc("/", handler)

	// The router is now formed by calling the 'newRouter' constructor function
	// that we defined above. the  rest of the code stays the same
	r := newRouter()
	http.ListenAndServe(":8080", r)

	// after defining our server, we finally "listen and serve" on port 8080
	// the second argument is the handler, which we will come to later on, but for now it is left as nil,
	// and the handler defined above in HandleFunc is used
	//http.ListenAndServe(":8080", nil)

	// This is where the router is useful, if allows us to declare methods that
	// this path will be valid for
	//r.HandleFunc("/hello", handler).Methods("GET")

	// We can then pass our router (after declaring all our routes) to this method
	// (where previously, we were leaving the second argument as nil)
	//http.ListenAndServe(":8080", r)
}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func handler(w http.ResponseWriter, r *http.Request) {
	//For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World!")
}
