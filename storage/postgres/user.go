package postgres

import (
	"fmt"
	"github/golang_practice/storage"
)

const getUser = `
	SELECT id, first_name, last_name, username, email, password from users
	WHERE id = $1
`

func (s *Storage) GetUser(id int32) (*storage.User, error) {
	user := storage.User{}
	if err := s.db.Get(&user, getUser, id); err != nil {
		return nil, err
	}
	return &user, nil
}
const createUserQuery = `
	INSERT INTO users(
		first_name,
		last_name,
		username,
		email,
		password
	)
	VALUES(
		:first_name,
		:last_name,
		:username,
		:email,
		:password
	 )
	RETURNING id
	`
	func (s *Storage) CreateUser(usr storage.User) (int32, error) {
		stmt, err := s.db.PrepareNamed(createUserQuery)
		if err != nil {
			return 0, err
		}
		var id int32
		if err := stmt.Get(&id, usr); err != nil {
			return 0, err
		}
		return id, nil
	}

	const getUserEmailAndPass = `
	SELECT * from users
	WHERE email = $1 AND password = $2
      `

func (s *Storage) GetUserEmailAndPass(email, password string) *storage.User {
	user := storage.User{}
	if err := s.db.Get(&user, getUserEmailAndPass, email, password); err != nil {
		return &user
	}
	fmt.Print("Get email and pass  = ", user)
	return &user
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