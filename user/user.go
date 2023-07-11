package user

import (
	api "cf_help_bot/api"
	help_func "cf_help_bot/help_func"
	"fmt"
	"log"
)

type User struct {
	Handle              string
	CurrentRating       int
	MaxRating           int
	SolvedQuantity      int
	SubmissionsQuantity int
	Submissions         []api.Submission
	Solved              []api.Problem
	RatingHistory       []int
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

// GetSubmissionsQuantity returns quantity of submissions by user
func (u *User) GetSubmissionsQuantity() int {
	return u.SubmissionsQuantity
}

// setSubmissionsQuantity sets quantity of submissions
func (u *User) setSubmissionsQuantity(quantity int) {
	u.SubmissionsQuantity = quantity
}

// GetSubmissions gets all submissions from user
func (u *User) GetSubmissions() []api.Submission {
	return u.Submissions
}

// setSubmissions sets value of Submissions field(Submission class from api)
func (u *User) setSubmissions(submissionList api.SubmissionList) {
	submissions := []api.Submission{}
	for i := 0; i < len(submissionList.Result); i++ {
		submissions = append(submissions, submissionList.Result[i])
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

// this function sets data for user by handle
func (u *User) Initialize(handle string) {
	u.setHandle(handle)
	log.Println("Handle:")
	log.Println(u.GetHandle())
	ratingChangeList := api.GetUserRating(u.Handle)
	u.setCurrentRating(ratingChangeList.Result[len(ratingChangeList.Result)-1].NewRating)
	u.setRatingHistory(ratingChangeList)
	u.setMaxRating(u.GetRatingHistory())
	submissionList := api.GetUserStatus(u.Handle)
	u.setSubmissions(submissionList)
	u.setSubmissionsQuantity(len(submissionList.Result))
	u.setSolved(submissionList)
	u.setSolvedQuantity(len(u.GetSolved()))
}

// this functions returns a dictionary string-> int, which represents the quantity of tasks with a given tag that the user has solved
func (u *User) Get_solved_quantity_by_tags() map[string]int {
	solved_quantity_by_tags := map[string]int{}
	current_user_rating := u.GetCurrentRating()
	for i := 0; i < len(u.Solved); i++ {
		if current_user_rating-u.Solved[i].Rating < 300 {
			for l := 0; l < len(u.Solved[i].Tags); l++ {
				solved_quantity_by_tags[u.Solved[i].Tags[l]]++
			}
		}
	}
	return solved_quantity_by_tags
}

// this function returns indexes of solved problems sorted by tags
func (U *User) Get_solved_indexes_by_tags() map[string][]string {
	solved_indexes_by_tags := map[string][]string{}
	for i := 0; i < len(U.Solved); i++ {
		for l := 0; l < len(U.Solved[i].Tags); l++ {
			contestId := fmt.Sprint(U.Solved[i].ContestId)
			solved_indexes_by_tags[U.Solved[i].Tags[l]] = append(solved_indexes_by_tags[U.Solved[i].Tags[l]],
				"/"+contestId+"/"+U.Solved[i].Index)
		}
	}
	return solved_indexes_by_tags
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
// ебучая хуйня, убрать при первой возможности магические константы
func (u *User) Get_wrong_attempts_by_task() map[string]int {
	m := map[string]int{}
	current_user_rating := u.GetCurrentRating()
	for i := 0; i < len(u.Submissions); i++ {
		current_submission := u.Submissions[i]
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
