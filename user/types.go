package user

type User struct {
	Id          uint32
	Username    string
	Password    string
	Email       string
	DisplayName string
	Status      int
}

type Permission struct {
	Id string
	Title string
}

type Role struct {
	Name        string
	Permissions []string
}
