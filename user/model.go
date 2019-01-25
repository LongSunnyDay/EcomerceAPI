package user

import (
	"go-api-ws/addresses"
	"time"
)

const (
	collectionName = "users"
)

type (
	User struct {
		ID       string   `json:"id,omitempty"`
		Customer Customer `json:"customer,omitempty"`
		Password string   `json:"password,omitempty"`
		GroupId  int      `json:"group_id,omitempty"`
	}

	Customer struct {
		FirstName string `json:"firstname,omitempty"`
		LastName  string `json:"lastname,omitempty"`
		Email     string `json:"email,omitempty"`
	}
	UpdatePassword struct {
		Password    string `json:"password,omitempty"`
		NewPassword string `json:"newPassword,omitempty"`
	}

	LoginForm struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}

	UpdatedCustomer struct {
		UpdateUser CustomerData `json:"customer,omitempty" bson:"customer,omitempty"`
	}

	CustomerData struct {
		Addresses              []addresses.Address `json:"addresses" bson:"addresses"`
		CreatedAt              time.Time           `json:"created_at,omitempty" bson:"created_at,omitempty"`
		CreatedIn              string              `json:"created_in,omitempty" bson:"created_in,omitempty"`
		DisableAutoGroupChange int32               `json:"disable_auto_group_change,omitempty" bson:"disable_auto_group_change,omitempty"`
		Email                  string              `json:"email,omitempty" bson:"email,omitempty"`
		FirstName              string              `json:"firstname,omitempty" bson:"firstname,omitempty"`
		GroupID                int64               `json:"group_id,omitempty" bson:"group_id,omitempty"`
		ID                     int64               `json:"id,omitempty" bson:"id,omitempty"`
		LastName               string              `json:"lastname,omitempty" bson:"lastname,omitempty"`
		StoreID                int32               `json:"store_id,omitempty" bson:"store_id,omitempty"`
		UpdatedAt              time.Time           `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
		WebsiteID              int32               `json:"website_id,omitempty" bson:"website_id,omitempty"`
		DefaultShipping        string              `json:"default_shipping,omitempty" bson:"default_shipping,omitempty"`
	}

	OrderHistory struct {
		Items          []Item `json:"items" bson:"items"`
		SearchCriteria string `json:"search_criteria" bson:"search_criteria"`
		TotalCount     int    `json:"total_count" bson:"omitempty"`
	}

	Item struct {
		SKU string `json:"sku,omitempty" bson:"sku"`
	}
)
