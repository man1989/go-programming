package main

import (
	"fmt"
)

//Movie represent a movie type
type Movie struct {
	title  string
	year   int
	genres []string
}

func main() {
	movie := Movie{title: "Batman Begins", year: 2005, genres: []string{"Action", "Adventure"}}
	fmt.Println(movie.genres[1])
}
