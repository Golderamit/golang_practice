
 package handler

import (
	"fmt"
	"github/golang_practice/storage"
	"net/http"
)

type adminHomeData struct{
     AdminHome   []storage.AdminHomeDB
	 Alladmincount int32
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
	adminuser, err := s.store.GetAdmin(1)
	
	fmt.Printf("%+v", adminuser)


	data := adminHomeData{
		AdminHome:     adminuser,

	}
	err = temp.Execute(w, data)

	fmt.Printf("%+v", err)
	if err !=nil{
		s.logger.Info("error with template execution: %+v", err)
	}
}
 


	var home []AdminHomeDB

	const query :=


func (s *Server) adminHomePage(w http.ResponseWriter, r *http.Request) {
	

}

