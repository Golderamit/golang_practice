package handler

import (
	"context"
	"html/template"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/csrf"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserSignUp struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Username  string `db:"username"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	IsAdmin   bool   `db:"is_admin"`
}
type userFormData struct {
	CSRFField template.HTML
	Form      UserSignUp
	Errors    map[string]error
}
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

func (s *Server) usersignup(w http.ResponseWriter, r *http.Request) {

	template := s.templates.Lookup("signup.html")
	if template == nil {
		s.logger.Error("lookup template signup.html")
		http.Error(w, "unable to load template", http.StatusInternalServerError)
		return
	}

	data := userFormData{
		CSRFField: csrf.TemplateField(r),
	}

	err := template.Execute(w, data)

	if err != nil {
		s.logger.Info("error with execute  template: %+v", err)
	}

}

func (s *Server) createUserSignUp(w http.ResponseWriter, r *http.Request) {

	template := s.templates.Lookup("signup.html")
	if template == nil {
		s.logger.Error("lookup template signup.html")
		http.Error(w, "unable to load template", http.StatusInternalServerError)
		return
	}

	if err := r.ParseForm(); err != nil {
		s.logger.WithError(err).Error("cannot parse form")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var form UserSignUp
	if err := s.decoder.Decode(&form, r.PostForm); err != nil {
		s.logger.WithError(err).Error("can not decode login form")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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

		err := template.Execute(w, data)

		if err != nil {
			s.logger.Info("error with template execution: %+v", err)
		}
		return
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

	http.Redirect(w, r, "/?success=true", http.StatusSeeOther)
}

func HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}
	return string(hash), nil

}
