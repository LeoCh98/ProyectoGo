package main

import (
	"Backend/db"
	"Backend/router"
	"log"
	"net/http"
)

func main() {
	// Inicializar la base de datos
	db.InitDB()
	defer db.CloseDB()

	// Iniciar el servidor en el puerto 9080
	log.Println("Starting the HTTP server on port 9080")
	http.ListenAndServe(":9080", router.NewRouter())
}
