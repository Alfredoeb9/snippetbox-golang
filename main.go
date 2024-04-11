package main

import (
	"fmt" // New import
	"log"
	"net/http"
	"strconv" // New import
)


func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
    // register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
    mux.HandleFunc("GET /snippet/view/{id}", snippetView)
    mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

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


// Define a home handler aka (controller) function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Use the Header().Add() method to add a 'Server: Go' header to the 
	// response header map. The first parameter is the header name, and
	// the second parameter is the header value
	w.Header().Add("Server", "Go")
    w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	
    id, err := strconv.Atoi(r.PathValue("id"))
	
    if err != nil || id < 1 {
		
        http.NotFound(w, r)
        return
    }

	

	// This will allow us to interpolate the wildcard id value in our response body message 
	// and write the response in a single line, like so:

    fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)

}

// Add a snippetCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Use the w.WriteHeader() method to send a 201 status code
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}
