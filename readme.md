# IMDB Chart Fetcher Guide

## Motivation:

1. The main motivation for the this project is to able to learn scrapping of website using Go programming language with some support from internal packages and some support from external package such as "goquery".
2. The project can be implemented using two ways as follows:
   > 1. Using normal sequential approach,which has its drawbacks such as high amount of data which needs processing with take higher amount of time.
   > 2. To eradicate the above sequential approach's problem we can implement it using Goroutines and respective terminologies assosiated with it.

### **User Guidance:**

1. [Running_Project](##running_project)
1. [Building_Project](##building_project)
1. [Executing_Build_Project](##executing_build_project)

## [Running_Project](##running_project)

> `cd /user_path/imdb_chart_fetcher`

> `go run main.go 'https://www.imdb.com/india/top-rated-indian-movies' 5`

## [Building_Project](##building_project)

> `cd /user_path/imdb_chart_fetcher`

> `go build -o imdb_chart_fetcher main.go`

## [Executing_Build_Project](##executing_build_project)

> `cd /user_path/imdb_chart_fetcher`

> `./imdb_chart_fetcher 'https://www.imdb.com/india/top-rated-indian-movies' 1`
