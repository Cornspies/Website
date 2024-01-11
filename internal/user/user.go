package user

type User struct {
	UserID         string
	IsLoggedIn     bool
	HasProfileIcon bool
}

func GetUserData() User {
	return User{
		UserID:         "",
		IsLoggedIn:     true,
		HasProfileIcon: false,
	}
}
