package services

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"main.go/models"
)

// GetMovieDetailsInstance : Function to return a object of Request type.
func GetMovieDetailsInstance(rank, title, year, rating, duration, summary, genre string) models.Movie {
	return models.Movie{
		Title:              title,
		Movie_release_year: getMovieYear(year),
		Imdb_rating:        getMovieRating(rating),
		Duration:           duration,
		Summary:            summary,
		Genre:              genre,
		Rank:               getMovieRank(rank),
	}
}

// getMovieYear : Function to return a int32 value of year.
func getMovieYear(year string) int32 {
	yearPattern := regexp.MustCompile("[0-9]{4}")
	onlyYear := string(yearPattern.Find([]byte(year)))
	number, er := strconv.Atoi(onlyYear)
	if er != nil {
		log.Fatal(er)
	}
	return int32(number)
}

// getMovieRank : Function to return a int32 value of rank.
func getMovieRank(rank string) int32 {

	pattern := regexp.MustCompile(`^([0-9]+)`)
	text := strings.TrimSpace(rank)
	Onlyrank, er := strconv.Atoi(pattern.FindString(text))
	if er != nil {
		log.Fatal(er)
	}
	return int32(Onlyrank)
}

// getMovieRating : Function to return a float32 value of rating.
func getMovieRating(rating string) float32 {
	number, er := strconv.ParseFloat(rating, 32)
	if er != nil {
		log.Fatal(er)
	}
	return float32(number)
}
