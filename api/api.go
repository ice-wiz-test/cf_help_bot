package api

 import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	user "cf_help_bot/user"
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

func getUserRating(u user.User) ResponseUserRating{
	// Make an HTTP GET request to the API endpoint
	resp, err := http.Get("https://codeforces.com/api/user.rating?handle=" + u.Handle)

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
	var data ResponseUserRating
 	// Unmarshal the JSON data into the data variable
	err = json.Unmarshal(body, &data)
	
	if err != nil {
		log.Fatal(err)
	}
	return data
}
