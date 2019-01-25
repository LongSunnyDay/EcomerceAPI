package tax

type (
	Rules struct {
		GroupId int     `json:"group_id"`
		Rates   Rate    `json:"rates"`
		Amount  float64 `json:"amount"`
	}

	Rate struct {
		Percent string `json:"percent"`
		Title   string `json:"title"`
	}
)
