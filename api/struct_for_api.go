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
