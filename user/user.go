package user

import (
	api "cf_help_bot/api"
	help_func "cf_help_bot/help_func"
)

type User struct {
	Handle              string
	CurrentRating       int
	MaxRating           int
	SolvedQuantity      int
	SubmissionsQuantity int
	Submissions         []api.Problem
	Solved              []api.Problem
	RatingHistory       []int
	SubmissionList      api.SubmissionList
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

// GetMaxRating returns the maximum rating of the user.
func (u *User) GetMaxRating() int {
	return u.MaxRating
}

// SetMaxRating sets the maximum rating of the user.
func (u *User) setMaxRating(data []int) {
	maxRating := help_func.Max(data)
	u.MaxRating = maxRating
}

// GetSolvedQuantity returns the value of the SolvedQuantity field.
func (u *User) GetSolvedQuantity() int {
	return u.SolvedQuantity
}

// SetSolvedQuantity sets the value of the SolvedQuantity field.
func (u *User) setSolvedQuantity(quantity int) {
	u.SolvedQuantity = quantity
}

func (u *User) GetSubmissionsQuantity() int {
	return u.SubmissionsQuantity
}

func (u *User) setSubmissionsQuantity(quantity int) {
	u.SubmissionsQuantity = quantity
}

func (u *User) GetSubmissions() []api.Problem {
	return u.Submissions
}

func (u *User) setSubmissions(submissionList api.SubmissionList) {
	submissions := []api.Problem{}
	for i := 0; i < len(submissionList.Result); i++ {
		submissions = append(submissions, submissionList.Result[i].RequestedProblem)
	}
	u.Submissions = submissions
}

// GetSolved returns the value of the Solved field.
func (u *User) GetSolved() []api.Problem {
	return u.Solved
}

// SetSolved sets the value of the Solved field.
func (u *User) setSolved(submissionList api.SubmissionList) {
	solved := []api.Problem{}
	for i := 0; i < len(submissionList.Result); i++ {
		if submissionList.Result[i].Verdict == "OK" {
			solved = append(solved, submissionList.Result[i].RequestedProblem)
		}
	}
	u.Solved = solved
}

// GetRatingHistory returns the value of the RatingHistory field.
func (u *User) GetRatingHistory() []int {
	return u.RatingHistory
}

// SetRatingHistory sets the value of the RatingHistory field.
func (u *User) setRatingHistory(ratingChangeList api.RatingChangeList) {
	ratingHistory := []int{}
	ratingHistory = append(ratingHistory, 0)
	for i := 0; i < len(ratingChangeList.Result); i++ {
		ratingHistory = append(ratingHistory, ratingChangeList.Result[i].NewRating)
	}
	u.RatingHistory = ratingHistory
}

// sets the submission list of a user as a given thing
func (u *User) setSubmissionList(list api.SubmissionList) {
	u.SubmissionList = list
}

// returns the user's submission list
func (u *User) GetSubmissionList() api.SubmissionList {
	return u.SubmissionList
}
func (u *User) Initialize(handle string) {
	u.setHandle(handle)
	ratingChangeList := api.GetUserRating(u.Handle)
	u.setCurrentRating(ratingChangeList.Result[len(ratingChangeList.Result)-1].NewRating)
	u.setRatingHistory(ratingChangeList)
	u.setMaxRating(u.GetRatingHistory())
	submissionList := api.GetUserStatus(u.Handle)
	u.setSubmissionList(submissionList)
	u.setSubmissions(submissionList)
	u.setSubmissionsQuantity(len(submissionList.Result))
	u.setSolved(submissionList)
	u.setSolvedQuantity(len(u.GetSolved()))
}

// this functions returns a dictionary string-> int, which represents the quantity of tasks with a given tag that the user has solved
func (u *User) Get_solved_quantity_by_tags() map[string]int {
	m := map[string]int{}
	current_user_rating := u.GetCurrentRating()
	for i := 0; i < len(u.Solved); i++ {
		if current_user_rating-u.Solved[i].Rating < 300 {
			for l := 0; l < len(u.Solved[i].Tags); l++ {
				m[u.Solved[i].Tags[l]]++
			}
		}
	}
	return m
}

// this function returns an int - number of tasks solved by user which have a rating much higher than his own
func (u *User) Get_solved_by_big_rating() int {
	k := 0
	current_user_rating := u.GetCurrentRating()
	for i := 0; i < len(u.Solved); i++ {
		if u.Solved[i].Rating-current_user_rating >= 300 {
			k++
		}
	}
	return k
}

// returns the amount of wrong attempts on "interesting" tasks by tags
func (u *User) Get_wrong_attempts_by_task() map[string]int {
	m := map[string]int{}
	current_user_rating := u.GetCurrentRating()
	for i := 0; i < len(u.SubmissionList.Result); i++ {
		current_submission := u.SubmissionList.Result[i]
		if current_submission.Verdict != "OK" && current_submission.RequestedProblem.Rating-current_user_rating >= -300 {
			if current_submission.PassedTestCount != 0 {
				for iter_str := 0; iter_str < len(current_submission.RequestedProblem.Tags); iter_str++ {
					m[current_submission.RequestedProblem.Tags[iter_str]]++
				}
			}
		}
	}
	return m
}
