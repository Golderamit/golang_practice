package postgres

import (
	"github/golang_practice/storage"
)

const getUser = `
	SELECT id, first_name, last_name, username, email from users
	WHERE id = $1
`

func (s *Storage) GetUser(id int32) (*storage.User, error) {
	user := storage.User{}
	if err := s.db.Get(&user, getUser, id); err != nil {
		return nil, err
	}
	return &user, nil
}

const userQuery = `SELECT * FROM users WHERE email=$1`

func (s *Storage) GetUserInfo(email string) *storage.User {
	user := storage.User{}
	err := s.db.Get(&user, userQuery, email)
	if err != nil {
		return &user
	}
	return &user
}