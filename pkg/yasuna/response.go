package yasuna

type Response[T any] struct {
	Data     T `json:"data"`
	Includes *struct {
		Users []*User `json:"users,omitempty"`
	} `json:"includes,omitempty"`
	Meta *struct {
		ResultCount int    `json:"result_count"`
		NextToken   string `json:"next_token"`
	} `json:"meta,omitempty"`
}

func (r *Response[T]) IncludeUser(id int64) *User {
	for _, user := range r.Includes.Users {
		if user.ID == id {
			return user
		}
	}
	return nil
}
