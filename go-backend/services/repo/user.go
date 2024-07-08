package repo

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	model "github.com/mmvergara/biscord/go-backend/models"
)

// CREATE TABLE users (
// 	id UUID PRIMARY KEY,
// 	username VARCHAR(255) NOT NULL UNIQUE,
// 	display_name VARCHAR(255) NOT NULL,
// 	email VARCHAR(255) NOT NULL UNIQUE,
// 	password VARCHAR(255) NOT NULL,
// 	avatar_url TEXT NOT NULL DEFAULT 'https://cdn.discordapp.com/embed/avatars/0.png',
// 	is_bot BOOLEAN NOT NULL DEFAULT FALSE,
// 	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
// 	last_active_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
// );

type UserRepo struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(user model.SignUpForm, hashedPassword model.HashedPassword) error {
	_, err := r.db.Exec("INSERT INTO users (email, username, display_name, password) VALUES ($1, $2, $3, $4)", user.Email, user.Username, user.DisplayName, hashedPassword)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to create user")
	}
	return nil
}

func (r *UserRepo) DoesUsernameExist(username string) (bool, error) {
	rows, err := r.db.Query("SELECT * FROM users WHERE username = $1", username)

	if err != nil {
		log.Println(err)
		return false, fmt.Errorf("failed to check if username exists")
	}

	defer rows.Close()

	for rows.Next() {
		return true, nil
	}

	return false, nil
}

func (r *UserRepo) DoesEmailExist(email string) (bool, error) {
	rows, err := r.db.Query("SELECT * FROM users WHERE email = $1", email)

	if err != nil {
		log.Println(err)
		return false, fmt.Errorf("failed to check if email exists")
	}

	defer rows.Close()

	for rows.Next() {
		return true, nil
	}

	return false, nil
}

func (r *UserRepo) GetUserByID(id model.UserID) (*model.User, error) {
	rows, err := r.db.Query("SELECT * FROM users WHERE id = $1", id.String())

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to get user by id")
	}

	defer rows.Close()

	user := new(model.User)

	for rows.Next() {
		user, err = scanRowsIntoUser(rows)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("failed to scan user")
		}
	}

	if user.ID == uuid.Nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (r *UserRepo) GetUserByEmail(email string) (*model.User, error) {
	rows, err := r.db.Query("SELECT * FROM users WHERE email = $1", email)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to get user by email")
	}

	defer rows.Close()

	user := new(model.User)

	for rows.Next() {
		user, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user")
		}
		log.Println(user)
	}

	if user.ID == uuid.Nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (r *UserRepo) GetUserByUsername(username string) (*model.User, error) {
	rows, err := r.db.Query("SELECT * FROM users WHERE username = $1", username)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to get user by username")
	}

	defer rows.Close()

	user := new(model.User)

	for rows.Next() {
		user, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user")
		}
		log.Println(user)
	}

	if user.ID == uuid.Nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func scanRowsIntoUser(rows *pgx.Rows) (*model.User, error) {
	var user model.User
	err := rows.Scan(&user.ID, &user.Username, &user.DisplayName, &user.Email, &user.Password, &user.AvatarURL, &user.IsBot, &user.CreatedAt, &user.LastActiveAt)
	if err != nil {
		return nil, err
	}
	log.Println("Found user")
	log.Println(user)
	log.Println(user)
	log.Println(user)
	log.Println(user)
	return &user, nil
}
