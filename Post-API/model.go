package posts

type Post struct {
	Id        string     `json:"id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Tags      []string   `json:"tags,omitempty"`
	Comments  []Comment  `json:"comments,omitempty"`
	Reactions []Reaction `json:"reactions,omitempty"`
}

type Comment struct {
	ID       string     `json:"id"`
	Body     string     `json:"body"`
	Reaction []Reaction `json:"reaction,omitempty"`
}

type Reaction struct {
	Emoji rune `json:"emoji"`
}
