package addresses

import (
	"go-api-ws/config"
	"go-api-ws/helpers"
)

// MySQL operations

func GetAddressesFromMySQL(userId int64) []Address {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var addresses []Address
	rows, err := db.Query("SELECT "+
		"id, "+
		"customer_id, "+
		"region_id, "+
		"country_id, "+
		"telephone, "+
		"postcode, "+
		"city, "+
		"firstname, "+
		"lastname, "+
		"default_shipping, "+
		"street_line_0, "+
		"street_line_1 FROM addresses WHERE customer_id = ? AND default_shipping = true", userId)
	for rows.Next() {
		var address Address
		if err := rows.Scan(&address.ID, &address.CustomerID, &address.RegionID,
			&address.CountryID, &address.Telephone, &address.Postcode, &address.City,
			&address.Firstname, &address.Lastname, &address.DefaultShipping,
			&address.StreetLine0, &address.StreetLine1); err != nil {
			helpers.PanicErr(err)
		}
		address.FormatStreetArray()
		address.GetRegion()
		addresses = append(addresses, address)
	}
	if len(addresses) > 0 {
		return addresses
	}
	return []Address{}
}

func (address *Address) InsertOrUpdateAddressIntoMySQL(customerId int64) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	if len(address.Street) == 1 {
		address.Street = []string{address.Street[0], ""}
	} else if len(address.Street) == 0 {
		address.Street = []string{"", ""}
	}

	res, err := db.Exec("INSERT INTO addresses("+
		"id, "+
		"customer_id, "+
		"region_id, "+
		"country_id, "+
		"telephone, "+
		"postcode, "+
		"city, "+
		"firstname, "+
		"lastname, "+
		"default_shipping, "+
		"default_billing, "+
		"email, "+
		"street_line_0, "+
		"street_line_1)"+
		" VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE "+
		"region_id=VALUES(region_id), "+
		"country_id=VALUES(country_id), "+
		"telephone=VALUES(telephone), "+
		"postcode=VALUES(postcode), "+
		"city=VALUES(city), "+
		"firstname=VALUES(firstname), "+
		"lastname=VALUES(lastname), "+
		"default_shipping=VALUES(default_shipping), "+
		"default_billing=VALUES(default_billing), "+
		"email=VALUES(email), "+
		"street_line_0=VALUES(street_line_0), "+
		"street_line_1=VALUES(street_line_1)", +
		address.ID,
		customerId,
		address.RegionID,
		address.CountryID,
		address.Telephone,
		address.Postcode,
		address.City,
		address.Firstname,
		address.Lastname,
		address.DefaultShipping,
		address.DefaultBilling,
		address.Email,
		address.Street[0],
		address.Street[1])
	helpers.PanicErr(err)
	addressId, err := res.LastInsertId()
	helpers.PanicErr(err)
	if addressId != 0 {
		address.ID = addressId
	}
}

func (address *Address) GetRegion() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var region Region
	err = db.QueryRow("SELECT region_id, region_code, region "+
		"FROM region WHERE region_id=?", address.RegionID).Scan(&region.RegionID, &region.RegionCode, &region.Region)
	helpers.PanicErr(err)
	address.Region = region
}

func (address *Address) FormatStreetArray() {
	address.Street = []string{address.StreetLine0, address.StreetLine1}
	address.StreetLine0 = ""
	address.StreetLine1 = ""
}

func UpdateAddressInMySQL(address Address) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("UPDATE addresses SET "+
		"region_id=?, "+
		"country_id=?, "+
		"telephone=?, "+
		"postcode=?, "+
		"city=?, "+
		"firstname=?, "+
		"lastname=?, "+
		"default_shipping=? "+
		"WHERE id=?",
		address.RegionID,
		address.CountryID,
		address.Telephone,
		address.Postcode,
		address.City,
		address.Firstname,
		address.Lastname,
		address.DefaultShipping)
	helpers.PanicErr(err)
}
