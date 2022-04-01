package handler

import (
	"fmt"
	"github/golang_practice/storage"
	"log"
	"net/http"
)

type adminHomeData struct{
     AdminHome   []storage.AdminHomeDB
	 Alladmincount int32
}

type Users struct {
	ID        int32     `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	IsActive  bool      `db:"is_active"`
	IsAdmin   bool      `db:"is_admin"`
}

type userviewData struct{
	Form []Users
}
func (s *Server) adminHomePage(w http.ResponseWriter, r *http.Request) {
	temp := s.templates.Lookup("admin-home.html")
	if temp == nil{
        s.logger.Error("lookup  template admin-home.html")
		http.Error(w,"unable to load template", http.StatusInternalServerError)
		return
	}
	var form []Users

	userInfoSelect := `SELECT 
	id,
	first_name, 
	last_name, 
	username, 
	password,
	email
	FROM users `

	err := s.db.Select(&form, userInfoSelect)
	if err != nil {
		s.logger.WithError(err).Info("db fetch users data")
		http.Error(w, "Not found.", http.StatusInternalServerError)
		return
	}
	tmpData := userviewData{
		
		Form: form,
	}

	err = temp.Execute(w, tmpData)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
 
