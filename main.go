package main

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"main.go/helpers"
	"main.go/models"
	"main.go/services"
)

var (
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	// Call the validator helper for the validations for the input data.
	chart_url, items_count := helpers.ValidateInputArguments(os.Args)

	// Create and return the instance of request.
	request := services.CreateRequest(chart_url, items_count)

	// Call the helper for with mechanism for calling request.
	document := request.GetUrlResponse()

	// Get the count of movies on the webpage.
	count, er := strconv.Atoi(document.Find("div.desc span").Text())
	if er != nil {
		log.Fatal(er)
	}

	// Check whether input value is greater than the count of movies on the webpage.
	err := request.CheckURLMoviesCount(int32(count))
	if err != nil {
		log.Fatal(err)
	}

	movies := &models.Result{}

	movies.InitiateMoviesArray(int(request.Items_Count))

	var counter int
	wg.Add(int(request.Items_Count))
	// Call for the goroutine for every data from the row value of movies.
	document.Find("tbody").Find("tr").EachWithBreak(func(index int, selection *goquery.Selection) bool {
		counter = index
		go scrapewebpages(selection, request, movies, &wg)
		if counter == int(request.Items_Count)-1 {
			return false
		} else {
			return true
		}
	})
	// When items_count is greater than available
	if counter != int(request.Items_Count)-1 {
		wg.Add(-int(request.Items_Count) + counter + 1)
		mutex.Lock()
		movies.Movies = movies.Movies[:counter+1]
		mutex.Unlock()
	}
	wg.Wait()
	movies.JsonEncoder(os.Stdout)
}

// scrapewebpages : Function which will be scrapping the webpages with the help of helpers and models.
func scrapewebpages(selection *goquery.Selection, request *models.Request, movies *models.Result, wg *sync.WaitGroup) {
	title := selection.Find("td.titleColumn a").Text()
	rank := selection.Find("td.titleColumn").Text()
	year := selection.Find("td.titleColumn span").Text()
	rating := selection.Find("td.imdbRating strong").Text()
	path, ok := selection.Find("td.titleColumn a").Attr("href")
	var duration, summary, genre string
	// Only if the Href value is present for the record.
	if ok {
		// Create new request instance for subdocument crawling.
		subDocumentRequest := services.CreateRequest(request.GetURLScheme()+"://"+request.GetHostName()+path, 0)
		// Call the helper for with mechanism for calling request.
		subDocument := subDocumentRequest.GetUrlResponse()
		duration = strings.TrimSpace(subDocument.Find("div.subtext time").Text())
		summary = strings.TrimSpace(subDocument.Find("div.summary_text").Text())
		children := subDocument.Find("div.subtext")
		// Extracting the data from 1 '|' and 3 '|' as genre in between them.
		Span1 := children.Find("span.ghost").Eq(1).NextAllFiltered("a")
		Span2 := children.Find("span.ghost").Eq(2).PrevAllFiltered("a")
		text := Span1.Intersection(Span2).Text()
		pattern := regexp.MustCompile("[A-Z][a-z]+")
		byteSlice := pattern.FindAll([]byte(text), -1)
		// Multiple genres are joined together by comma in single string.
		genre = string(bytes.Join(byteSlice, []byte(", ")))
	}

	// To Get the movie instance with required data filled.
	movie := services.GetMovieDetailsInstance(rank, title, year, rating, duration, summary, genre)
	mutex.Lock()
	movies.AppendMovie(movie)
	mutex.Unlock()
	wg.Done()
}
