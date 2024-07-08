package repo

import "github.com/jackc/pgx"

type UserRepo struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) GetUserByID(id int) int {
	return 1
}
