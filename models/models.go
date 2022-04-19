package models

// Trivia Model
type Trivia struct {
	Text   string `json:"text" bson:"text,omitempty"`
	Number int64  `json:"number" bson:"number,omitempty"`
	Found  bool   `json:"found"`
	Type   string `json:"type"`
}

// Trivia create new instance of Trivia
// todo
func createTrivia(text string, number int64) Trivia {
	trivia := Trivia{}
	trivia.Text = text
	trivia.Number = number
	trivia.Found = true
	trivia.Type = "trivia"
	return trivia
}
