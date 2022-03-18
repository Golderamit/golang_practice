package handler

import (
	"github/golang_practice/storage"
	"html/template"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/csrf"
)

type SignUpFormData struct {
	CSRFField  template.HTML
	Form       storage.User
	FormErrors map[string]string
}

func (s *Server) getSignup(w http.ResponseWriter, r *http.Request) {

	session, _ := s.session.Get(r, "practice_project_app")
	userId := session.Values["user_id"]

	if _, ok := userId.(string); ok {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}

	data := SignUpFormData{
		CSRFField: csrf.TemplateField(r),
	}

	s.SignupTemplate(w, r, data)

}

func (s *Server) postSignup(w http.ResponseWriter, r *http.Request) {

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
        s.SignupTemplate(w, r, data)
		return
	}
	id, err := s.store.SaveUser(form)
	if err != nil {
		log.Println("data not saved")
	}
    log.Println(id)

	log.Printf("\n %#v", form)

	http.Redirect(w, r, "/login/?Success=True", http.StatusTemporaryRedirect)
	

}

func (s *Server) SignupTemplate(w http.ResponseWriter, r *http.Request, form SignUpFormData) {
	temp := s.templates.Lookup("signup.html")

	if err := temp.Execute(w, form); err != nil {
		log.Fatalln("executing template: ", err)
		return
	}
}
