package handler

import (
	"fmt"
	"github/practice22/storage"
	"log"
	"net/http"
)

/* type UserDB struct {
	ID        int32  `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Username  string `db:"username"`
	Email     string `db:"email"`
}
*/
type (
	templateData struct {
		User storage.User
	}
)
func (s *Server) getHome(w http.ResponseWriter, r *http.Request) {
	tmp := s.templates.Lookup("home.html")
	if tmp == nil {
		log.Println("unable to look home.html")
		return
	}

	user, err := s.store.GetUser(1)
	if err != nil {
		log.Println("unable to get user: ", err)
		return
	}

	fmt.Printf("%+v", user)

	tmpData := templateData{
		User: *user,
	}

	err = tmp.Execute(w, tmpData)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}