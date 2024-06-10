package v1UserRepo

import "go.mongodb.org/mongo-driver/mongo"

type UserRepo struct {
	mdb *mongo.Database
}

func NewUserRepo(mdb *mongo.Database) *UserRepo {
	return &UserRepo{
		mdb: mdb,
	}
}

func (u *UserRepo) CreateUser() {
	
}