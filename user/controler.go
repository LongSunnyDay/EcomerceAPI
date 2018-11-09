package user

func roleByGroupId(groupId int) (string) {
	if groupId < 1 {
		return adminRole
	}
	return userRole
}
