package core

type UserModel struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(name, email, password string) *UserModel {
	u := &UserModel{Name: name, Email: email, Password: password}

	Save("user", u)

	return u
}
