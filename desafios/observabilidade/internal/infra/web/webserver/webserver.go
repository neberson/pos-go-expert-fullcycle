package webserver

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type WebServer struct {
	Router        chi.Router
	Routes        []Route
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	return &WebServer{
		Router:        r,
		Routes:        make([]Route, 0),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	s.Routes = append(s.Routes, Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	})
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)

	for _, r := range s.Routes {
		switch r.Method {
		case http.MethodGet:
			s.Router.Get(r.Path, r.Handler)
		case http.MethodPost:
			s.Router.Post(r.Path, r.Handler)
		default:
			panic("método HTTP não suportado: " + r.Method)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
