package handler

import (
	"fmt"
	"github/golang_practice/storage"
	"html/template"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/csrf"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
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



	fmt.Printf("****************  %+v", template)

	 /*  session, _ := s.session.Get(r, "practice_project_app")

type Storage struct {
	db *sqlx.DB
}
func (f *UserSignUp) ValidationUserFrom(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, f,
		validation.Field(&f.FirstName, validation.Required.Error("FirstName is required")),
		validation.Field(&f.LastName, validation.Required.Error("LastName is required")),
		validation.Field(&f.Username, validation.Required.Error("Username is required")),
		validation.Field(&f.Email, validation.Required.Error("Email is required")),
		validation.Field(&f.Password, validation.Required.Error("Password is required")),
	)
}
func (f *UserSignUp) UserDB(id int) *UserSignUp{

	return &UserSignUp{
		ID:        id,
		FirstName: f.FirstName,
		LastName:  f.LastName,
		Username:  f.Username,
		Email:     f.Email,
		Password:  f.Password,
	}
}


	session, _ := s.session.Get(r, "practice_project_app")

	userId := session.Values["user_id"]


	if _, ok := userId.(string); ok {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

	}   */


	template := s.templates.Lookup("signup.html")
	if template == nil {
		s.logger.Error("lookup template signup.html")
		http.Error(w, "unable to load template", http.StatusInternalServerError)
		return

	}


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

	if err := form.Validate(); err != nil {
		vErros := map[string]string{}

		if e, ok := err.(validation.Errors); ok {
			if len(e) > 0 {
				for key, value := range e {
					vErros[key] = value.Error()
				}
			}

	savedVErrs := validation.Errors{}

	if err := form.ValidationUserFrom(r.Context()); err != nil {
		if vErrs, ok := (err).(validation.Errors); ok {
			savedVErrs = vErrs
		} else {
			s.logger.WithError(err).Error("validate user form")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
    form.ID = 0 
	if len(savedVErrs) > 0 {
		data := userFormData{
			CSRFField: csrf.TemplateField(r),
			Form:      form,
			Errors:    savedVErrs,

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

	http.Redirect(w, r, "/?Success=True", http.StatusTemporaryRedirect)


} 


	pass := form.Password
	hashed, err := HashAndSalt(pass)
	form.Password = hashed

	const createUserQuery = `
	INSERT INTO users(
		first_name,
		last_name,
		username,
		email,
		password
	)
	VALUES(
		:first_name,
		:last_name,
		:username,
		:email,
		:password
	 )
	 RETURNING id
	`
	_, err = s.db.Exec(createUserQuery, form.UserDB(0))
	if err != nil{
      s.logger.WithError(err).Error("failed to insert users")
	  http.Error(w, "unable to insert users", http.StatusInternalServerError)
	  return
	} 


}

func (s *Server) SignupTemplate(w http.ResponseWriter, r *http.Request, form SignUpFormData) {
	temp := s.templates.Lookup("signup.html")


}

