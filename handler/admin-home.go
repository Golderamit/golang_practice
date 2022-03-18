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
 
