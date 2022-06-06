package yasuna

type ResponseError struct {
	Errors []struct {
		Parameters struct {
			Expansions []string `json:"expansions"`
		} `json:"parameters"`
		Message string `json:"message"`
	} `json:"errors"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Type   string `json:"type"`
}

func (e *ResponseError) Error() string {
	return "yasuna: response error: " + e.Detail
}
