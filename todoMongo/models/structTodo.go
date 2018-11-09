package models

type Todo struct {
	ObjectId string `json:"ObjectID" bson:"ObjectID"`
	ID       string `json:"id,omitempty" bson:"id"`
	Title    string `json:"title,omitempty" bson:"title"`
	Category string `json:"category,omitempty" bson:"category"`
	Content  string `json:"content,omitempty" bson:"content"`
	Created  string `json:"created,omitempty" bson:"created"`
	Modified string `json:"modified,omitempty" bson:"modified"`
	State    string `json:"state,omitempty" bson:"state"`
}
