package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/nu50218/go-cacher"
	"github.com/nu50218/nuinfo-syllabus-go/syllabus"
)

type Server struct {
	handler       http.Handler
	client        *syllabus.Client
	subjectsCache cacher.Cacher
	subjectCache  *cacher.Map
}

type Config struct {
	Endpoint string        `env:"ENDPOINT" envDefault:"https://syllabus.i.nagoya-u.ac.jp/i/"`
	Expires  time.Duration `env:"EXPIRES" envDefault:"1h"`
	Interval time.Duration `env:"INTERVAL" envDefault:"500ms"`
}

func New(c Config) *Server {
	s := &Server{
		client: syllabus.NewClient(c.Endpoint, c.Interval),
	}

	// cacher
	s.subjectsCache = cacher.New(c.Expires, func() interface{} {
		subjects, err := s.client.GetAllConciseSubjects()
		if err != nil {
			return err
		}
		return subjects
	})
	s.subjectCache = cacher.NewMap(func(key interface{}) cacher.Cacher {
		return cacher.New(c.Expires, func() interface{} {
			subject, err := s.client.GetSubject(key.(string))
			if err != nil {
				return err
			}
			return subject
		})
	})

	// router
	r := chi.NewRouter()
	r.Get("/subjects", s.getAllSubjects)
	r.Get("/subjects/{code}", s.getSubject)
	s.handler = r

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}

func (s *Server) getAllSubjects(w http.ResponseWriter, r *http.Request) {
	i := s.subjectsCache.Load()
	if err, ok := i.(error); ok {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(i); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (s *Server) getSubject(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	i := s.subjectCache.Get(code).Load()
	if err, ok := i.(error); ok {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(i); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
