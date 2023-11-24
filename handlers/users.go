package handlers

import (
	"encoding/json"
	"golang-api/auth"
	"golang-api/config"
	"golang-api/models"
	"log"
	"net/http"
)

var db = config.GetDBConnect()

func GetUsersData(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan token dari header request
	tokenString := r.Header.Get("Authorization")

	if !auth.ValidateToken(tokenString) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Mengambil data pengguna dari database MySQL
	script := "SELECT id, username, password, is_user, is_superuser FROM users"

	// Membuat channel untuk komunikasi antara goroutine
	resultChan := make(chan []models.User)
	errorChan := make(chan error)

	// Menjalankan goroutine untuk pengambilan data
	var users []models.User
	go func() {
		rows, err := db.Query(script)
		if err != nil {
			errorChan <- err
			return
		}
		defer rows.Close()

		for rows.Next() {
			var user models.User
			err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.IsUser, &user.IsSuperuser)
			if err != nil {
				errorChan <- err
				return
			}

			users = append(users, user)
		}

		resultChan <- users
	}()

	select {
	case materials := <-resultChan:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(materials)
	case err := <-errorChan:
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// func GetUserById(w http.ResponseWriter, r *http.Request) {
// 	// Mendapatkan token dari header request
// 	tokenString := r.Header.Get("Authorization")

// 	if !auth.ValidateToken(tokenString) {
// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		return
// 	}

// 	// Mengambil data pengguna dari database MySQL
// 	script := "SELECT id, username, password, is_user, is_superuser FROM users WHERE id = ?"
// }
