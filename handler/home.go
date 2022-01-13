package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type UserDB struct {
	ID        int32  `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Username  string `db:"username"`
	Email     string `db:"email"`
}
type userviewdata struct {
	user []UserDB
}
type DBWithNamed struct {
	*sqlx.DB
}

func (s *Server) getHome(res http.ResponseWriter, req *http.Request) {

	template := s.templates.Lookup("./assets/templates/home.html")

	if template != nil {
		s.logger.Error("lookup template ./assets/templates/home.html")
		http.Error(res, "unable to load template", http.StatusInternalServerError)
	}
	var users []UserDB
	userquery := `SELECT id,
	                     first_name, 
	                     last_name,
	                     username,
	                     email
	                     FROM users 
	 `
	 err :=s.db.Select(&users,userquery)

	 s.logger.Infoln("#########   %+v",users)
	 if err != nil{
		 s.logger.WithError(err).Info("db fetch users data")
		 http.Error(res,  "Not found.", http.StatusInternalServerError)
		 return
	 }
	 data:=userviewdata{
		user:   users,
	 }
	 err = template.Execute(res, data)

	 if err != nil {
		 s.logger.Info("error with template execution: %+v ",err)
	 }
}
