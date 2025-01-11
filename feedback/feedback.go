package feedback

// Feedback represents a feedback object
type Feedback struct {
	ID    string
	Email string
	Title string `json:"title"`
	Body  string `json:"body"`
}
