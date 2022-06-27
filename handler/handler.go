package handler

import (
	"github/golang_practice/storage/postgres"
	"html/template"
	"log"
	"net/http"

	"github.com/Masterminds/sprig"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Server struct {
	templates *template.Template
	store     *postgres.Storage
	logger    *logrus.Logger
	decoder   *schema.Decoder
	session   *sessions.CookieStore
	db         *sqlx.DB
}

func NewServer(st *postgres.Storage, decoder *schema.Decoder, session *sessions.CookieStore, db *sqlx.DB) (*mux.Router, error) {

	s := &Server{
		store: st,
		decoder: decoder,
		session: session,
		db:     db,
	}

	if err := s.parseTemplates(); err != nil {
		return nil, err
	}

	r := mux.NewRouter()

    r.Use(csrf.Protect([]byte("1234")))






    csrf.Protect([]byte("go-secret-go-safe-----"), csrf.Secure(false))(r)


	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))
	
	r.HandleFunc("/", s.getHome).Methods("GET")


	r.HandleFunc("/login", s.getLogin).Methods("GET")
	r.HandleFunc("/login", s.postLogin).Methods("POST")
	r.HandleFunc("/logout", s.logOut)
	r.HandleFunc("/signup", s.getSignup).Methods("GET")
	 r.HandleFunc("/signup", s.postSignup).Methods("POST") 

	r.HandleFunc("/login/", s.getLogin).Methods("GET")
	r.HandleFunc("/login/", s.postLogin).Methods("POST")


	r.HandleFunc("/signup/", s.getSignup).Methods("GET")
	r.HandleFunc("/signup/", s.postSignup).Methods("POST")


	
	r.HandleFunc("/admin-home", s.adminHomePage).Methods("GET")
	return r, nil
}
func (s *Server) parseTemplates() error {
	templates := template.New("templates").Funcs(template.FuncMap{
		"strrev": func(str string) string {
			n := len(str)
			runes := make([]rune, n)
			for _, rune := range str {
				n--
				runes[n] = rune
			}
			return string(runes[n:])
		},
	}).Funcs(sprig.FuncMap())

	tmpl, err := templates.ParseGlob("assets/templates/*.html")
	if err != nil {
		return err
	}
	s.templates = tmpl
	return nil
}


 func (s *Server) DefaultTemplate(w http.ResponseWriter, r *http.Request, temp_name string, data interface{}) {

	temp := s.templates.Lookup(temp_name)

	if err := temp.Execute(w, data); err != nil {
		log.Fatalln("executing template: ", err)
		return
	}

} 

} 

func SessionCheckAndRedirect(s *Server, r *http.Request, next http.Handler, w http.ResponseWriter, user bool) {
	uid, user_type := GetSetSessionValue(s, r)
	if uid != "" && user_type == user {
		next.ServeHTTP(w, r)
	} else {
		http.Redirect(w, r, "/forbidden", http.StatusSeeOther)
	}
}

func GetSetSessionValue(s *Server, r *http.Request) (interface{}, interface{}) {
	session, _ := s.session.Get(r, "practice_project_app")
	uid := session.Values["user_id"]
	user_type := session.Values["is_admin"]
	return uid, user_type
}


