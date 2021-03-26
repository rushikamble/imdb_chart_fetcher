package models

import (
	"encoding/json"
	"io"
	"log"
)

// Movie : structure or schema for the indiviual records for the movie.
// Also please note the sequence which the variable is defined is descending according
// to each data-type memory size for optimal memory consumption.
type Movie struct {
	Title              string  `json:"title"`
	Summary            string  `json:"summary"`
	Duration           string  `json:"duration"`
	Genre              string  `json:"genre"`
	Imdb_rating        float32 `json:"imdb_rating"`
	Movie_release_year int32   `json:"movie_release_year"`
	Rank               int32   `json:"-"` //Field is ommited from enconding and used 											  //internally as for just say Sorting of movies 										  //rank-wise.
}

// Result : structure required to store to array of object of movie type
type Result struct {
	Movies Movies
}

type Movies []Movie

// InitiateMoviesArray : To initiate the movies array in result with blank objects of movie struct.
func (Result *Result) InitiateMoviesArray(count int) {
	for i := 0; i < count; i++ {
		Result.Movies = append(Result.Movies, Movie{})
	}
}

// AppendMovie : To add the data of the movie rank wise in the blank structs of Movies Array
func (Result *Result) AppendMovie(movie Movie) {
	// Result.Movies = append(Result.Movies, movie)
	Result.Movies[movie.Rank-1] = movie
}

// JsonEncoder : Used with standard output to encode the json and show to the command line.
func (Result *Result) JsonEncoder(writer io.Writer) {
	err := json.NewEncoder(writer).Encode(Result.Movies)
	if err != nil {
		log.Fatal("Error while encoding the json file.")
	}
}
