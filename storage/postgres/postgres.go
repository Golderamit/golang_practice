package postgres

import (
	"github.com/jmoiron/sqlx"
)

type Storage struct{
	db *sqlx.DB
}

func NewStorage(dbstring string) (*Storage ,error){
	db,err := sqlx.Connect("postgres",dbstring)
	if err != nil {
		return nil, err
	}
	err =db.Ping()
	if err !=nil{
		return nil,err
	}
	return &Storage{db: db}, nil
}