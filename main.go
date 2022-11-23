package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func allUserCount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("gsvged")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("mhd")
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "123", Title: "Yüzüklerin Efendisi", Director: &Director{Firstname: "Arzu", Lastname: "Yıldız"}})
	movies = append(movies, Movie{ID: "2", Isbn: "456", Title: "Esaretin Bedeli", Director: &Director{Firstname: "Ahmet", Lastname: "Mehmet"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/create", createMovie).Methods("POST")
	r.HandleFunc("/movies/update/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/delete/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/allUserCount", allUserCount).Methods("GET")

	fmt.Println("Starting Server t port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
