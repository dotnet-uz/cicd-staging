package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type App struct {
	DB *sql.DB
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	app := &App{DB: db}

	http.HandleFunc("/health", app.health)
	http.HandleFunc("/users", app.getUsers)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func connectDB() (*sql.DB, error) {
	connStr := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"

	return sql.Open("postgres", connStr)
}

func (a *App) health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{Message: "OK"})
}

func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := a.DB.Query("SELECT id, name FROM users")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	var users []User

	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)
}
