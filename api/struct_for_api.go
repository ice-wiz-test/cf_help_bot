package api

type Problem struct {
	ContestId      int      `json:"contestId"`
	ProblemsetName string   `json:"problemsetName"`
	Index          string   `json:"index"`
	Name           string   `json:"name"`
	Type           string   `json:"type"`
	Points         float64  `json:"points"`
	Rating         int      `json:"rating"`
	Tags           []string `json:"tags"`
}

type Member struct {
	Handle string `json:"string"`
	Name   string `json:"name"`
}

type Party struct {
	ContestId        int      `json:"contestId"`
	Members          []Member `json:"members"`
	ParticipantType  string   `json:"participantType"`
	TeamID           int      `json:"teamId"`
	TeamName         string   `json:"teamName"`
	Ghost            bool     `json:"ghost"`
	Room             int      `json:"room"`
	StartTimeSeconds int64    `json:"startTimeSeconds"`
}

type Submission struct {
	ID                  int64   `json:"id"`
	ContestId           int     `json:"contestId"`
	CreationTimeSeconds int64   `json:"creationTimeSeconds"`
	RelativeTimeSeconds int64   `json:"relativeTimeSeconds"`
	RequestedProblem    Problem `json:"problem"`
	Author              Party   `json:"party"`
	ProgrammingLanguage string  `json:"programmingLanguage"`
	Verdict             string  `json:"verdict"`
	PassedTestCount     int     `json:"passedTestCount"`
	TimeConsumedMillis  int     `json:"timeConsumedMillis"`
	MemoryConsumedBytes int     `json:"memoryConsumedBytes"`
	Points              float64 `json:"points"`
}

type SubmissionList struct {
	Status string       `json:"status"`
	Result []Submission `json:"result"`
}

type RatingChange struct {
	ContestID               int    `json:"contestId"`
	ContestName             string `json:"contestName"`
	Handle                  string `json:"handle"`
	Rank                    int    `json:"rank"`
	RatingUpdateTimeSeconds int    `json:"ratingUpdateTimeSeconds"`
	OldRating               int    `json:"oldRating"`
	NewRating               int    `json:"newRating"`
}

type RatingChangeList struct {
	Status string         `json:"status"`
	Result []RatingChange `json:"result"`
}

type User struct {
	LastName                string `json:"lastName"`
	LastOnlineTimeSeconds   int    `json:"lastOnlineTimeSeconds"`
	Rating                  int    `json:"rating"`
	FriendOfCount           int    `json:"friendOfCount"`
	TitlePhoto              string `json:"titlePhoto"`
	Handle                  string `json:"handle"`
	Avatar                  string `json:"avatar"`
	FirstName               string `json:"firstName"`
	Contribution            int    `json:"contribution"`
	Organization            string `json:"organization"`
	Rank                    string `json:"rank"`
	MaxRating               int    `json:"maxRating"`
	RegistrationTimeSeconds int    `json:"registrationTimeSeconds"`
	MaxRank                 string `json:"maxRank"`
}
type UserList struct {
	Status string `json:"status"`
	Result []User `json:"result"`
}
