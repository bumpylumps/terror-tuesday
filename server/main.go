package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

var movies = []Movie{
	{ID: 1, Title: "Blue Velvet", Director: "David Lynch", Runtime: "1hr, 30mins", DateWatched: "4/20/1992"},
	{ID: 2, Title: "Friday the 13th pt 4", Director: "", Runtime: "1h 30m", DateWatched: "1/8/25"},
}

// get movies responds with all movies
func getAll(w http.ResponseWriter, req *http.Request) {
	//set the content type
	w.Header().Set("Content-Type", "application/json")

	//
	slcMovies, err := json.Marshal(movies)

	if err != nil {
		http.Error(w, "Failed to encode movies", http.StatusInternalServerError)
		return
	}

	w.Write(slcMovies)
	fmt.Fprintf(w, "\n hello little bro")
}

// post new movie
func postMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//add movie to slice
	testelem := Movie{ID: 3, Title: "The Blob", Director: "George Romero", Runtime: "1h 30m", DateWatched: "1/8/25"}
	movies = append(movies, testelem)
	//write slice to browser

	movies, err := json.Marshal(movies)

	if err != nil {
		http.Error(w, "Failed to encode movies", http.StatusInternalServerError)
		return
	}

	w.Write(movies)
	fmt.Fprintf(w, "\n Adding new movie to movies[]")

}

func deleteMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	i := 2

	//delete by creating a new slice and omitting desired index
	//: used for before and after index in slice
	movies = append(movies[:i], movies[i+1:]...)

	movies, err := json.Marshal(movies)

	if err != nil {
		http.Error(w, "Failed to encode movies", http.StatusInternalServerError)
	}

	w.Write(movies)
	fmt.Fprintf(w, "\n removed The blob")

}

func main() {
	http.HandleFunc("/hello", hello)
	//get all movies
	http.HandleFunc("/movies/all", getAll)
	//add movies to list
	http.HandleFunc("/movies/add", postMovie)
	//delete movie from list
	http.HandleFunc("/movies/delete", deleteMovie)
	//update movie from list

	http.ListenAndServe(":8090", nil)
}

// movie struct
type Movie struct {
	ID          int    `json: "id"`
	Title       string `json: "title"`
	Director    string `json: "director"`
	Runtime     string `json: "runtime"`
	DateWatched string `json: "dateWatched"`
}
