package v1Handler

import "go.mongodb.org/mongo-driver/mongo"

type UserHandler struct {
	mdb *mongo.Database
}

func NewUserHandler(mdb *mongo.Database) *UserHandler {
	return &UserHandler{
		mdb: mdb,
	}
}


// registerUser

type RegisterUserRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`

}

func (u *UserHandler) RegisterUser() {
}