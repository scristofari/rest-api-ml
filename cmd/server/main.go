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
	r.HandleFunc("/result", resultHandler).Methods("GET")
	return r
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	// s, err := ml.LaunchArtifact("./ml/fixture/archive.tar")
	// if err != nil {
	//  	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	// enc := json.NewEncoder(w)
	// enc.Encode(s)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Result Handler")
}
