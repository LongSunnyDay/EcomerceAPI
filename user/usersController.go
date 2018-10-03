package user

import "net/http"

func getUsersHandler(w http.ResponseWriter, r * *http.Request, user AuthorizedUser) {

}

func getUserHandler(w http.ResponseWriter, r *http.Request) {

}

func postUserHandler(w http.ResponseWriter, r *http.Request) {

}

func postUserUpdateHandler(w http.ResponseWriter, r *http.Request) {

}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
}

func updateUerPermissions(w http.ResponseWriter, r *http.Request) {

}


func deleteUserPermission(w http.ResponseWriter, r *http.Request) {

}

func deleteUserMeta(w http.ResponseWriter, r *http.Request) {

}

func deleteUserRole(w http.ResponseWriter, r *http.Request) {

}

