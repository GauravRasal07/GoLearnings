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
	Firstname  string `json:"firstname"`
	Laststname string `json:"lastname"`
}

var movies []Movie

func getMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}

}

func getMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func createMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	
	movie.ID = strconv.Itoa(rand.Intn(9999))

	movies = append(movies, movie)

	json.NewEncoder(res).Encode(movie)

}

func updateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for index, item := range movies {
		if item.ID == params["id"] {
			//Deleting the movie with given ID
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie
			_ = json.NewDecoder(req.Body).Decode(&movie)

			movie.ID = params["id"]

			//Adding new movie
			movies = append(movies, movie)

			json.NewEncoder(res).Encode(movie)
			return
		}
	}
}

func deleteMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(res).Encode(movies)
}

func main() {
	fmt.Println("Heyy There!")

	router := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "7911", Title: "Animal", Director: &Director{Firstname: "Gaurav", Laststname: "Rasal"}})
	movies = append(movies, Movie{ID: "2", Isbn: "7912", Title: "Kabir Singh", Director: &Director{Firstname: "Sandeep", Laststname: "Reddy"}})
	movies = append(movies, Movie{ID: "3", Isbn: "7913", Title: "Salaar", Director: &Director{Firstname: "Akshay", Laststname: "Kumar"}})
	movies = append(movies, Movie{ID: "4", Isbn: "7914", Title: "12th Fail", Director: &Director{Firstname: "The", Laststname: "Director"}})
	movies = append(movies, Movie{ID: "5", Isbn: "7915", Title: "Wolf of Wall Street", Director: &Director{Firstname: "No", Laststname: "Need"}})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at PORT: 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
