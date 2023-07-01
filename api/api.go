package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseUserRating struct {
	ContestID               int    `json:"contestId"`
	ContestName             string `json:"contestName"`
	Handle                  string `json:"handle"`
	Rank                    int    `json:"rank"`
	RatingUpdateTimeSeconds int    `json:"ratingUpdateTimeSeconds"`
	OldRating               int    `json:"oldRating"`
	NewRating               int    `json:"newRating"`
}

type ResponseUserRatingUpdates struct {
	Status string               `json:"status"`
	Result []ResponseUserRating `json:"result"`
}

func check() {
	fmt.Println("ok")
}

func GetUserRating(handle string) ResponseUserRatingUpdates {
	// Make an HTTP GET request to the API endpoint
	resp, err := http.Get("https://codeforces.com/api/user.rating?handle=" + handle)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	// Create a variable to hold the parsed response
	var data ResponseUserRatingUpdates
	// Unmarshal the JSON data into the data variable
	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal(err)
	}
	return data
}
