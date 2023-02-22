package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	MovieId       string    `json:movieid`
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
	movies = append(movies, Movie{MovieId: "1", MovieName: "Django Unchained", MovieLength: "2h 45min",
		MovieDirector: &Director{Fullname: "Quentin Tarantino", Nmovies: 10}})
	movies = append(movies, Movie{MovieId: "2", MovieName: "Avatar", MovieLength: "2h 42min",
		MovieDirector: &Director{Fullname: "James Cameron", Nmovies: 9}})

	r := mux.NewRouter()
	r.HandleFunc("/", welcome).Methods("GET")
	r.HandleFunc("/movies", getall).Methods("GET")
	r.HandleFunc("/movies/{id}", getone).Methods("GET")
	r.HandleFunc("/movies", createone).Methods("POST")
	r.HandleFunc("/movies/{id}", updateone).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteone).Methods("DELETE")

	fmt.Println("Listening on port 4000")

	log.Panic(http.ListenAndServe(":4000", r))

}
func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome msg")
	w.Header().Set("content-type", "application/json")
	w.Write([]byte("WELCOME TO MOVIES USER"))
}
func getall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all movie details")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func getone(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting one movie")
	w.Header().Set("content-type", "application-json")
	params := mux.Vars(r)
	var count int = 0
	for _, movie := range movies {
		if movie.MovieId == params["id"] {
			json.NewEncoder(w).Encode(movie)
			count++
		}
	}
	if count == 0 {
		json.NewEncoder(w).Encode("Invalid Input")
	}
}
func createone(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating movie details")
	w.Header().Set("content-type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please enter some data")
	}
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}

func updateone(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updation of a movie")
	params := mux.Vars(r)
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)

	for index, mov := range movies {
		if mov.MovieId == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
		}
	}

	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}
func deleteone(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting a movie")
	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.MovieId == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(movies)
		}

	}
}
