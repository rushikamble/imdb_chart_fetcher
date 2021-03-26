package models

import (
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// Request : Structure or schema required to defined internal request to server
type Request struct {
	Url         string `json:"chart_url"`
	Items_Count int32  `json:"items_count"`
}

// GetUrlResponse : Function to return the HTTP response and goquery document
func (request *Request) GetUrlResponse() *goquery.Document {
	resp, err := http.Get(request.Url)
	if err != nil {
		log.Fatalf("GetUrlResponse : Error while fetching URL %v:", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("GetUrlResponse : URL Fetching error: %d %s", resp.StatusCode, resp.Status)
	}

	// Translate to the HTML document
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("GetUrlResponse : Error while translating URL %v:", err)
	}
	return document
}

// CheckURLMoviesCount: To check whether the request is execeeding the number of movies from the arguments.
func (request *Request) CheckURLMoviesCount(count int32) error {
	if count < request.Items_Count {
		return errors.New("CheckURLMoviesCount : The input count is greater than the movies present in the URL. URL has only " + string(request.Items_Count) + " movies")
	} else {
		return nil
	}
}

// GetHostName : To get the hostname from the URL.
func (request *Request) GetHostName() string {

	url, err := url.Parse(request.Url)
	if err != nil {
		log.Fatal("GetHostName : ", err)
	}
	return url.Host
}

// GetURLScheme : To get the scheme from the URL.
func (request *Request) GetURLScheme() string {
	url, err := url.Parse(request.Url)
	if err != nil {
		log.Fatal("GetURLScheme : ", err)
	}
	return url.Scheme
}
