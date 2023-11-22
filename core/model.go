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

func FindByEmail(email string) (*UserModel, error) {
	var prev []UserModel
	err := getFile("user", &prev)
	if err != nil {
		return nil, err
	}

	for _, user := range prev {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, nil
}
