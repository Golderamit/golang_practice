package handler

import (
	"database/sql"
	"fmt"
	"github/golang_practice/storage"
	"log"
	"net/http"
)
type HomePage struct{
	UserLoggedIn bool
	Form      storage.User
}

func (s *Server) getHome(w http.ResponseWriter, r *http.Request) {
	tmp := s.templates.Lookup("home.html")
	if tmp == nil {
		log.Println("unable to look home.html")
		return
	}
    session, _ := s.session.Get(r, "practice_project_app")
	userID := session.Values["user_id"]

	if _, ok := userID.(string); ok {
		data := HomePage{
			UserLoggedIn: true,
		}
		if err := tmp.Execute(w, data); err != nil {
			log.Fatalln("Session Execution error")
		}
		return
	}
	var form storage.User

	userInfoSelect := `SELECT 
	id,
	email
	FROM users WHERE email=$1`

   if  err := s.db.Get(&form, userInfoSelect, form.Email); err == sql.ErrNoRows {
	} else if err != nil {
		s.logger.WithError(err).Error("db select users")
		http.Error(w, "failed to get users", http.StatusInternalServerError)
		return
	
}
	tmpData := HomePage{
		UserLoggedIn: false,
		Form: form,
	}

	fmt.Printf("$$$$$$$$$$$$    %+v",form)
	err := tmp.Execute(w, tmpData)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
    
	
	