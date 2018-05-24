package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not load the .env file")
	}

	http.Handle("/", handlers())

	port := os.Getenv("PORT")
	log.Printf("Listening on port %s ...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/upload", uploadHandler).Methods("POST")
	r.HandleFunc("/artifact/{uuid}", artifactHandler).Methods("GET")
	r.HandleFunc("/artifact/{uuid}/logs/{event}", logHandler).Methods("GET")
	return r
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	// @TODO Call the function in background, with a goroutine.
	// s, err := ml.LaunchArtifact("./ml/fixture/archive.tar")
	// if err != nil {
	//  	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	// enc := json.NewEncoder(w)
	// enc.Encode(s)
}

func artifactHandler(w http.ResponseWriter, r *http.Request) {
	// Call GetArtifactInfo from the ml package
	http.Error(w, "Not Implemented !", http.StatusInternalServerError)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	// Call GetArtifactLogs from the ml package
	http.Error(w, "Not Implemented !", http.StatusInternalServerError)
}
