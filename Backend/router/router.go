package router

import (
	"Backend/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter configura las rutas y devuelve un *mux.Router con soporte para CORS
func NewRouter() *CORSRouterDecorator {
	router := mux.NewRouter()

	// Configurar las rutas
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	return &CORSRouterDecorator{router}
}

// CORSRouterDecorator aplica encabezados CORS a las respuestas
type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
	}

	if req.Method == "OPTIONS" {
		rw.WriteHeader(http.StatusOK)
		return
	}

	c.R.ServeHTTP(rw, req)
}
