package yasuna

type User struct {
	ID       int64  `json:"id,string"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Verified bool   `json:"verified"`
}
