package user


type User struct {
	Handle          string
	Rating          int
	SolvedQuantity  int
	Solved          []int
	RatingHistory   []int
}

 // GetHandle returns the value of the Handle field.
func (u *User) GetHandle() string {
	return u.Handle
}

 // SetHandle sets the value of the Handle field.
func (u *User) SetHandle(handle string) {
	u.Handle = handle
}

 // GetRating returns the value of the Rating field.
func (u *User) GetRating() int {
	return u.Rating
}

 // SetRating sets the value of the Rating field.
func (u *User) SetRating(rating int) {
	u.Rating = rating
}

 // GetSolvedQuantity returns the value of the SolvedQuantity field.
func (u *User) GetSolvedQuantity() int {
	return u.SolvedQuantity
}

 // SetSolvedQuantity sets the value of the SolvedQuantity field.
func (u *User) SetSolvedQuantity(quantity int) {
	u.SolvedQuantity = quantity
}

 // GetSolved returns the value of the Solved field.
func (u *User) GetSolved() []int {
	return u.Solved
}

 // SetSolved sets the value of the Solved field.
func (u *User) SetSolved(solved []int) {
	u.Solved = solved
}

// GetRatingHistory returns the value of the RatingHistory field.
func (u *User) GetRatingHistory() []int {
	return u.RatingHistory
}

// SetRatingHistory sets the value of the RatingHistory field.
func (u *User) SetRatingHistory(history []int) {
	u.RatingHistory = history
}
