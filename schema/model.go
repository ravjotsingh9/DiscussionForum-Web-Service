package schema

type Comment struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	PID     string `json:"pid"`
	TID     string `json:"tid"`
}
