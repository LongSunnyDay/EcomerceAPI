package tax

import (
	"go-api-ws/config"
	"go-api-ws/helpers"
)

func (r *Rules) GetRates() Rules {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var rules Rules
	err = db.QueryRow("SELECT * FROM tax_rates t WHERE t.groupId = ?", r.GroupId).
		Scan(&rules.GroupId, &rules.Rates.Percent, &rules.Rates.Title)
	helpers.PanicErr(err)
	return rules
}
