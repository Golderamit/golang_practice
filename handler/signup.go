package handler

import (
	"fmt"
	"github/golang_practice/storage"
	"html/template"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/csrf"
)

type SignUpFormData struct {
	CSRFField  template.HTML
	Form       storage.User
	FormErrors map[string]string
}

func (s *Server) getSignup(w http.ResponseWriter, r *http.Request) {
	template := s.templates.Lookup("signup.html")
	if template == nil {
		s.logger.Error("lookup template signup.html")
		http.Error(w, "unable to load template", http.StatusInternalServerError)
		return
	}
    fmt.Printf("$$$$$$$$$$$$    %+v",template)
	
	 /*  session, _ := s.session.Get(r, "practice_project_app")
	userId := session.Values["user_id"]

	if _, ok := userId.(string); ok {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}   */

	data := SignUpFormData{
		CSRFField: csrf.TemplateField(r),
	}

	err := template.Execute(w, data)
	

	if err != nil {
		s.logger.Info("error with execute  template: %+v", err)

	}
}

 func (s *Server) postSignup(w http.ResponseWriter, r *http.Request) {

	template := s.templates.Lookup("signup.html")
	if template == nil {
		s.logger.Error("lookup template signup.html")
		http.Error(w, "unable to load template", http.StatusInternalServerError)
		return
	}
    fmt.Printf("$$$$$$$$$$$$    %+v",template)
	if err := r.ParseForm(); err != nil {
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
	fmt.Printf("$$$$$$$$$$$$    %+v",form)
	if err := form.Validate(); err != nil {
		vErros := map[string]string{}

		if e, ok := err.(validation.Errors); ok {
			if len(e) > 0 {
				for key, value := range e {
					vErros[key] = value.Error()
				}
			}
		}
		data := SignUpFormData{
			CSRFField:  csrf.TemplateField(r),
			Form:       form,
			FormErrors: vErros,
		}
		err := template.Execute(w, data)

        fmt.Printf("$$$$$$$$$$$$    %+v",data)		
		if err != nil {
			s.logger.Info("error with execute  template: %+v", err)
		}
		return
	}

	id, err := s.store.SaveUser(form)
	if err != nil {
		log.Println("data not saved")
	}
	fmt.Printf("$$$$$$$$$$$$    %+v",id)
	log.Println(id)

	log.Printf("\n %#v", form)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

} 
