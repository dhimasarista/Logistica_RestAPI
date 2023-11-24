package main

import (
	"database/sql"
	"fmt"
	"golang-api/config"
	"golang-api/handlers"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Import driver MySQL
	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	// Membuat koneksi ke database
	db = config.GetDBConnect()
	defer db.Close()

	fmt.Println("Server is running on http://localhost:8080")

	r := mux.NewRouter()

	// Rute untuk halaman utama
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to LOGISTICA API with GOLANG")
	}).Methods("GET")

	// Rute untuk data users
	r.HandleFunc("/users", handlers.GetUsersData).Methods("GET")

	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
