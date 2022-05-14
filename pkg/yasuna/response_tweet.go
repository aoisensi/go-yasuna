package yasuna

import "time"

type Tweet struct {
	ID        int64     `json:"id,string"`
	Text      string    `json:"text"`
	AuthorID  int64     `json:"author_id,string"`
	CreatedAt time.Time `json:"created_at"`
}
