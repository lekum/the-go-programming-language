package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type movie struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Metascore  string
	imdbRating string
	imdbVotes  string
	imdbID     string
	Type       string
	Response   string
}

func getPosterURL(title string) (string, error) {
	baseURL := "http://www.omdbapi.com/?r=json&t="
	t := url.QueryEscape(title)
	resp, err := http.Get(baseURL + t)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return "", fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return "", err
	}
	resp.Body.Close()
	return result.Poster, nil
}

func downloadPoster(posterURL string, filename string) error {
	resp, err := http.Get(posterURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, bs, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	movieName := strings.Join(os.Args[1:], " ")
	posterURL, err := getPosterURL(movieName)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
	err = downloadPoster(posterURL, "poster.jpg")
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
}
