package model

type HashedPassword = string

type SignInForm struct {
	Email    string         `json:"email"`
	Password HashedPassword `json:"password"`
}

type SignUpForm struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
