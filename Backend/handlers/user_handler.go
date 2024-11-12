package handlers

import (
	"Backend/db"
	"Backend/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// GetUsers: obtiene todos los usuarios
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	result, err := db.GetDB().Query("SELECT id, first_name, last_name, email FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer result.Close()

	for result.Next() {
		var user models.User
		if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

// CreateUser: crea un nuevo usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var user models.User
	json.Unmarshal(body, &user)

	stmt, err := db.GetDB().Prepare("INSERT INTO users(first_name, last_name, email) VALUES(?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "New user was created")
}

// GetUser: obtiene un usuario por ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user models.User

	result := db.GetDB().QueryRow("SELECT id, first_name, last_name, email FROM users WHERE id = ?", params["id"])
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// UpdateUser: actualiza un usuario existente
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	body, _ := ioutil.ReadAll(r.Body)
	var user models.User
	json.Unmarshal(body, &user)

	stmt, err := db.GetDB().Prepare("UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User with ID = %s was updated", params["id"])
}

// DeleteUser: elimina un usuario por ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	stmt, err := db.GetDB().Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User with ID = %s was deleted", params["id"])
}
