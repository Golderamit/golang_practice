package handler

import (
	"fmt"
	"github/golang_practice/storage"
	"html/template"
	"log"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gorilla/csrf"
)


type Login struct {
	Email    string
	Password string
}

type LoginTempData struct {
	CSRFField  template.HTML
	Form      storage.User
	FormErrors map[string]string
}
func (l Login) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Email, validation.Required.Error("email is required"), is.Email),
		validation.Field(&l.Password, validation.Required.Error("Password is required"), validation.Length(6, 12).Error("Password Lenght must be 3 to 10")),
	)
}
func (s *Server) getLogin(w http.ResponseWriter, r *http.Request) {
    
	template := s.templates.Lookup("login.html")
	if template == nil {
		s.logger.Error("lookup template login.html")
		http.Error(w, "unable to load template", http.StatusInternalServerError)
		return
	}
    session, _ := s.session.Get(r, "practice_project_app")
	userId:=session.Values["user_id"] 

	if _,ok:=userId.(string);ok{
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}

	tempData := LoginTempData{
		CSRFField:  csrf.TemplateField(r),
		
	}
    err := template.Execute(w, tempData)

	if err != nil {
		s.logger.Info("error with execute  template: %+v", err)
		
	}
	
	
}

func (s *Server) postLogin(w http.ResponseWriter, r *http.Request) {

	template := s.templates.Lookup("login.html")
	if template == nil {
		s.logger.Error("lookup template login.html")
		http.Error(w, "unable to load template", http.StatusInternalServerError)
		return
     } 

	if err := r.ParseForm(); err !=nil {
		s.logger.WithError(err).Error("cannot parse form")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var form storage.User
	if err := s.decoder.Decode(&form, r.PostForm); err != nil {
		s.logger.WithError(err).Error("can not decode login form")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
    if err := form.ValidateUser(); err != nil {
		vErrs := map[string]string{}
		if e, ok := err.(validation.Errors); ok {
			if len(e) > 0 {
				for key, value := range e {
					vErrs[key] = value.Error()
				}
			}
		}
		tempData := LoginTempData{
			CSRFField:  csrf.TemplateField(r),
			Form:       form,
			FormErrors: vErrs,
		}
		err := template.Execute(w,tempData)
	
		if err != nil {
			s.logger.Info("error with execute  template: %+v", err)
			
		}
		return
	}
	log.Println(form)

	log.Println("###########=hi=============")
	log.Println(form)
    
	user, err := s.store.GetUser(form.Email, form.Password)


    

	user, err := s.store.GetUser(form.Email, form.Password)


	user, err := s.store.GetUser(form.Email, form.Password)

	email := form.Email
	result := s.store.GetUserInfo(email)
	ComparePassword(result, form, w, r)
	sessionUID := result.ID
	isAdmin := result.IsAdmin
	session, _ := s.session.Get(r, "practice_project_app")
	session.Values["user_id"] = IntToStringConversion(sessionUID)
	session.Values["is_admin"] = isAdmin
	if err := session.Save(r, w); err != nil {
		log.Fatalln("error while saving user id into session")
	}
	http.Redirect(w, r, "/home?success=true", http.StatusSeeOther)
}


	if err != nil {
		log.Fatalln("user not found")
		return
	}


    fmt.Printf("##########  %+v", user)


}
 func LoginRedirect(isAdmin bool, w http.ResponseWriter, r *http.Request) {
	if isAdmin == true {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)

	session, _ := s.session.Get(r, "practice_project_app")
	session.Values["user_id"] = strconv.Itoa(int(user.ID))
	session.Values["user_email"] = user.Email
	if err := session.Save(r, w); err != nil {
		log.Fatalln("saving error session")
	}


	http.Redirect(w, r, "/?success=true", http.StatusTemporaryRedirect)
}

} 


