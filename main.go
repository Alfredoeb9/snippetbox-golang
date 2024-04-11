package main

import (
	"log"
	"net/http"
)

// Define a home handler aka (controller) function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
    // register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Print a log message to say that the server is starting
	log.Print("starting server on :4000")

	// use the http.ListenAndServer() function to start a new web sserver. we pass in 
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servermux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nul.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}