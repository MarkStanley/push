package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func main() {

	// Start up the webserver
	startServer()
}

func startServer() {
	CertDir := os.Getenv("CERT_DIR")
	// Instatiate the router
	r := newRouter()

	// start the server on port 8080 with TLS
	serverCert := CertDir + "/fullchain.pem"
	serverKey := CertDir + "/privkey.pem"
	fmt.Println("I am running HTTPS")
	log.Fatal(http.ListenAndServeTLS(":9500", serverCert, serverKey, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))
}

func newRouter() *mux.Router {
	r := mux.NewRouter()

	// As a general principle of URL structure we will be consistent with /entity/verb/{id}
	// For example /friends/add/{friendID}

	// These first few handlers deal with signin and signup via email/password or google.
	// They result in the passing back of a cookie containing a JWT
	r.HandleFunc("/push", ShowPushButton).Methods("GET")
	r.HandleFunc("/push", PushMessage).Methods("POST")
	r.HandleFunc("/service-worker.js", SendSW).Methods("POST")

	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	return r
}
