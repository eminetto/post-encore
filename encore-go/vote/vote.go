package vote

// Vote is a struct that represents a vote
type Vote struct {
	ID       string
	Email    string
	TalkName string `json:"talk_name"`
	Score    int    `json:"score,string"`
}
