package models

type Todo struct {
	ObjectId string `json:"ObjectID,omitempty" bson:"ObjectID,omitempty"`
	ID       string `json:"id,omitempty" bson:"id,omitempty"`
	Title    string `json:"title,omitempty" bson:"title,omitempty"`
	Category string `json:"category,omitempty" bson:"category,omitempty"`
	Content  string `json:"content,omitempty" bson:"content,omitempty"`
	Created  string `json:"created,omitempty" bson:"created,omitempty"`
	Modified string `json:"modified,omitempty" bson:"modified,omitempty"`
	State    string `json:"state,omitempty" bson:"state,omitempty"`
}
