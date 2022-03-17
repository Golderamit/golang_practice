package handler

import (
	"context"
	"html/template"
	"net/http"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/csrf"
)

type UserSignUp struct{
  ID        int32     `db:"id"`
  FirstName string    `db:"first_name"`
  LastName  string    `db:"last_name"`
  Username  string    `db:"username"`
  Email     string    `db:"email"`
  Password  string    `db:"password"`
  IsAdmin   bool      `db:"is_admin"` 
}
type userFormData struct{
	CSRFField  template.HTML   
	Form       UserSignUp 
	Errors map[string]error
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
func (s *Server) usersignup(w http.ResponseWriter, r *http.Request) {

   template := s.templates.Lookup("signup.html")
   if template == nil{
	   s.logger.Error("lookup template login.html")
	   http.Error(w, "unable to load template",http.StatusInternalServerError)
       return
   }
  
   data := userFormData{
   	CSRFField: csrf.TemplateField(r),
   }

   err := template.Execute(w ,data)

   if err != nil{
	s.logger.Info("error with execute  template: %+v", err)
   }
   
}

func (s *Server) createUserSignUp(w http.ResponseWriter, r *http.Request) {
  
	template := s.templates.Lookup("signup.html")
	if template == nil {
		s.logger.Error("lookup template login.html")
		http.Error(w,"unable to load template",http.StatusInternalServerError)
		return
	}

	if err := r.ParseForm(); err != nil {
      s.logger.WithError(err).Error("cannot parse form")
	  http.Error(w, err.Error(), http.StatusBadRequest)
	  return
	}

	var form UserSignUp
	if err := s.decoder.Decode(&form, r.PostForm); err != nil{
		s.logger.WithError(err).Error("can not decode login form")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
    savedVErrs := validation.Errors{}

	if err:= form.ValidationUserFrom(r.Context()); err != nil{
        if vErrs, ok := (err).(validation.Errors); ok {
			savedVErrs = vErrs
		} else {
			s.logger.WithError(err).Error("validate event form")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		} 
	}

	if len(savedVErrs) > 0 {
		data := userFormData{
			CSRFField: csrf.TemplateField(r),
			Form:         form,
			Errors:       savedVErrs,
		}

		err := template.Execute(w, data)

		if err != nil {
			s.logger.Info("error with template execution: %+v", err)
		}
		return
	}
    http.Redirect(w, r, "/home?success=true", http.StatusSeeOther)
}
