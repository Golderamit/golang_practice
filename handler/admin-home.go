
package handler


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



package handler

import (
	"net/http"
)

type AdminHomeDB struct {
	ID               int32     `db:"id"`
	Title            string     `db:"title"`
	Venue            string    `db:"venue"`
	Address          string    `db:"address"`
	Country          string     `db:"country"`
	Email            string    `db:"email"`
	PhoneNumber      string    `db:"phone_number"`
	ShortDescription string    `db:"short_description"`
	Description      string     `db:"description"`
	Image            string     `db:"image"`
	Status           bool      `db:"status"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
	FromDate         time.Time `db:"from_date"`
	ToDate           time.Time `db:"to_date"`	
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
 



	var home []AdminHomeDB

	const query :=


func (s *Server) adminHomePage(w http.ResponseWriter, r *http.Request) {
	

}



