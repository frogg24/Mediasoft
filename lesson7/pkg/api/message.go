package api

type MessageIn struct {
	To   string `json:"to"`
	Body string `json:"body"`
}
type MessageOut struct {
	From string `json:"from"`
	Body string `json:"body"`
}
