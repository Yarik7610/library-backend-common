package event

type BookAdded struct {
	ID             uint   `json:"id"`
	AuthorID       uint   `json:"authorId"`
	AuthorFullname string `json:"authorFullname"`
	Title          string `json:"title"`
	Year           int    `json:"year"`
	Category       string `json:"category"`
}
