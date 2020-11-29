package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
	// "io"
	// "log"
	"os"
)

// Movie struct for json conversion
type Movie struct {
	UniqueMovieID string
	MovieName     string
	MovieGenre    string
}

// Plot structure
type Plot struct {
	moviePlot     string
	UniqueMovieID string
}

func main() {

	// Read Movie Metadata CSV file
	meta, err := os.Open("data/MovieSummaries/movie.metadata.csv")
	if err != nil {
		fmt.Printf("Error opening the csv file :%s", err)
	}
	reader := csv.NewReader(bufio.NewReader(meta))
	reader.Comma = '\t'
	defer meta.Close()

	var movies []Movie
	var moviePlots []Plot

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		// fmt.Println(line)
		movies = append(movies, Movie{
			UniqueMovieID: line[0],
			MovieName:     line[2],
			MovieGenre:    line[8],
		})
	}

	// Plots of movies
	plots, err := os.Open("data/MovieSummaries/plot_summaries.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer plots.Close()
	scanner := bufio.NewScanner(plots)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		s := strings.Split(scanner.Text(), "\t")
		// fmt.Printf("MovieID is %s and Movie Plot is %s\n", s[0], s[1])
		// fmt.Printf("---------------\n")
		moviePlots = append(moviePlots, Plot{
			UniqueMovieID: s[0],
			moviePlot:     s[1],
		})
	}
	fmt.Println(moviePlots)

	// fmt.Println(string(plots))

	// moviej, _ := json.Marshal(movies)
	// fmt.Println(string(moviej))
	// fmt.Println(string(moviej)[0])
	// fmt.Println(reflect.TypeOf(moviej).String())
	// fmt.Println(movies)
}
