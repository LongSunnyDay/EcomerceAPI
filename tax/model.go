package tax

import (
	"go-api-ws/config"
	"go-api-ws/helpers"
)

type Rules struct {
	GroupId int     `json:"group_id"`
	Rates   Rate    `json:"rates"`
	Amount  float64 `json:"amount"`
}

type Rate struct {
	Percent string `json:"percent"`
	Title   string `json:"title"`
}

func (r *Rules) GetRates() Rules {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var rules Rules
	err = db.QueryRow("SELECT * FROM taxRates t WHERE t.groupId = ?", r.GroupId).
		Scan(&rules.GroupId, &rules.Rates.Percent, &rules.Rates.Title)
	helpers.PanicErr(err)
	return rules
}
