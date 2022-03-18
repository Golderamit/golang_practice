package handler

import (
	"log"
	"net/http"
)
type HomePage struct{
	UserLoggedIn bool
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
	
	tmpData := HomePage{
		UserLoggedIn: false,
	}
	err := tmp.Execute(w, tmpData)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
    
	
	