package postgres

import "github/practice22/storage/postgres"

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