package handler

import (
	"github/golang_practice/storage/postgres"
	"html/template"
	"net/http"

	"github.com/Masterminds/sprig"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
)

type Server struct {
	templates *template.Template
	store     *postgres.Storage
	logger    *logrus.Logger
	decoder   *schema.Decoder
	session   *sessions.CookieStore
}

func NewServer(st *postgres.Storage, decoder *schema.Decoder, session *sessions.CookieStore) (*mux.Router, error) {

	s := &Server{
		store: st,
		decoder: decoder,
		session: session,
	}

	if err := s.parseTemplates(); err != nil {
		return nil, err
	}

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))
	r.HandleFunc("/", s.getHome).Methods("GET")

	r.HandleFunc("/login", s.getLogin).Methods("GET")
	r.HandleFunc("/login", s.postLogin).Methods("POST")

	r.HandleFunc("/signup", s.usersignup).Methods("GET")
	r.HandleFunc("/signup", s.createUserSignUp).Methods("POST")
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
