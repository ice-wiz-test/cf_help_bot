package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func check() {
	fmt.Println("ok")
}

// This function gets rating of user by handle from codeforces api
func GetUserRating(handle string) RatingChangeList {
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
	var data RatingChangeList
	// Unmarshal the JSON data into the data variable
	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal(err)
	}
	return data
}

// This function returns information about user by handle from codeforces api
func GetUserStatus(handle string) SubmissionList {
	// Make an HTTP GET request to the API endpoint
	resp, err := http.Get("https://codeforces.com/api/user.status?handle=" + handle)

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
	var data SubmissionList
	// Unmarshal the JSON data into the data variable
	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal(err)
	}
	return data
}
