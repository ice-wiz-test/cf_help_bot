package user

import (
	api "cf_help_bot/api"
)

type User struct {
	Handle         string
	CurrentRating  int
	SolvedQuantity int
	Solved         []int
	RatingHistory  []int
}

// GetHandle returns the value of the Handle field.
func (u *User) GetHandle() string {
	return u.Handle
}

// SetHandle sets the value of the Handle field.
func (u *User) setHandle(handle string) {
	u.Handle = handle
}

// GetRating returns the value of the Rating field.
func (u *User) GetCurrentRating() int {
	return u.CurrentRating
}

// SetRating sets the value of the Rating field.
func (u *User) setCurrentRating(rating int) {
	u.CurrentRating = rating
}

// GetSolvedQuantity returns the value of the SolvedQuantity field.
func (u *User) GetSolvedQuantity() int {
	return u.SolvedQuantity
}

// SetSolvedQuantity sets the value of the SolvedQuantity field.
func (u *User) setSolvedQuantity(quantity int) {
	u.SolvedQuantity = quantity
}

// GetSolved returns the value of the Solved field.
func (u *User) GetSolved() []int {
	return u.Solved
}

// SetSolved sets the value of the Solved field.
func (u *User) setSolved(solved []int) {
	u.Solved = solved
}

// GetRatingHistory returns the value of the RatingHistory field.
func (u *User) GetRatingHistory() []int {
	return u.RatingHistory
}

// SetRatingHistory sets the value of the RatingHistory field.
func (u *User) setRatingHistory(data api.ResponseUserRatingUpdates) {
	ratingHistory := []int{}
	ratingHistory = append(ratingHistory, 0)
	for i := 0; i < len(data.Result); i++ {
		u.RatingHistory = append(u.RatingHistory, data.Result[i].NewRating)
	}
	u.RatingHistory = ratingHistory
}

func (u *User) Initialize(handle string) {
	u.setHandle(handle)
	data := api.GetUserRating(u.Handle)
	u.setCurrentRating(data.Result[len(data.Result)-1].NewRating)
	u.setRatingHistory(data)
}
