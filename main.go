package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	MovieId       int       `json:movieid`
	MovieName     string    `json:moviename`
	MovieLength   string    `json:movielength`
	MovieDirector *Director `json:moviedirector`
}
type Director struct {
	Fullname string `json:fullname`
	Nmovies  int    `json:nmovies` //no of movies directed by the director
}

var movies []Movie

func main() {
	movies = append(movies, Movie{MovieId: 1, MovieName: "Django Unchained", MovieLength: "2h 45min",
		MovieDirector: &Director{Fullname: "Quentin Tarantino", Nmovies: 10}})
	movies = append(movies, Movie{MovieId: 2, MovieName: "Avatar", MovieLength: "2h 42min",
		MovieDirector: &Director{Fullname: "James Cameron", Nmovies: 9}})

	r := mux.NewRouter()
	r.HandleFunc("/", getall).Methods("GET")
	fmt.Println("Listening on port 4000")
	log.Panic(http.ListenAndServe(":4000", r))

}

func getall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all movie details")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
