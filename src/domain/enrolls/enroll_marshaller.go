package enrolls

type PublicUserCourse struct {
	UserID int `json:"user_id"`
}

func (user User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUserCourse{
			UserID: user.UserID,
		}
	}

	return user
}
