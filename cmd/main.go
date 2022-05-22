package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string
	Isbn     string
	Director *Director
	Title    string
}

type Director struct {
	Firstname string
	Lastname  string
}

var Movies []Movie

func main() {
	r := mux.NewRouter()
	port := "8000"
	r.HandleFunc("/getallmovies", GetAllMovies).Methods(http.MethodGet)
	r.HandleFunc("/getmovie/{id}", GetMovie).Methods(http.MethodGet)
	r.HandleFunc("/createmovie", CreateMovie).Methods(http.MethodPost)
	r.HandleFunc("/updatemovie/{id}", UpdateMovie).Methods(http.MethodPatch)
	r.HandleFunc("/deletemovie/{id}", DeleteMovie).Methods(http.MethodDelete)

	fmt.Printf("Starting port at: %s\n", port)
	fmt.Println("Press ctrl + c to stop the server")
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/getmovie"))
}
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	m := Movie{
		Id:   "Hash1",
		Isbn: "ASDWQ1223431",
		Director: &Director{
			Firstname: "ABC",
			Lastname:  "XYZ",
		},
		Title: "Horror Nigts",
	}
	Movies = append(Movies, m)
	resp, err := json.Marshal(m)
	if err != nil {
		w.Write([]byte("Error marshalling data"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/createmovie"))

}
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/updatemovie"))

}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/deletemovie"))

}
