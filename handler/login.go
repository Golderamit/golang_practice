package handler

import (
	"github/golang_practice/storage"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/csrf"
	"golang.org/x/crypto/bcrypt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)


type Login struct {
	Email    string
	Password string
}

type LoginTempData struct {
	CSRFField  template.HTML
	Form       Login
	FormErrors map[string]string
}
func (l Login) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Email, validation.Required, is.Email),
		validation.Field(&l.Password, validation.Required, validation.Length(6, 12)),
	)
}
func (s *Server) getLogin(w http.ResponseWriter, r *http.Request) {
    
	template := s.templates.Lookup("login.html")
	if template == nil {
		s.logger.Error("lookup template login.html")
		http.Error(w, "unable to load template", http.StatusInternalServerError)
		return
	}
    
	tempData := LoginTempData{
		CSRFField:  csrf.TemplateField(r),
		
	}
    err := template.Execute(w,tempData)

	if err != nil {
		s.logger.Info("error with execute  template: %+v", err)
		
	}
	return
	
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

	var form Login
	if err := s.decoder.Decode(&form, r.PostForm); err != nil {
		s.logger.WithError(err).Error("can not decode login form")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
    if err := form.Validate(); err != nil {
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
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}


func IntToStringConversion(id int32) string {
	t := strconv.Itoa(int(id))
	return t
}
func ComparePassword(result *storage.User, form Login, w http.ResponseWriter, r *http.Request) {
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(form.Password)); err != nil {
		log.Println("Password does not match.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
/* func LoginRedirect(isAdmin bool, w http.ResponseWriter, r *http.Request) {
	if isAdmin == true {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}
} */