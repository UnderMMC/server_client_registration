package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type RegistrData struct {
	Login    string
	Password string
}

func PersonalDataHandler(w http.ResponseWriter, r *http.Request) {
	data := RegistrData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func DataBaseRecording() {
	connStr := "user=username dbname=mydb password=mysecretpassword host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	// Проверяем соединение
	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}

	// Создание таблицы
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username TEXT, password TEXT)")
	if err != nil {
		log.Fatalf("Ошибка при создании таблицы: %v", err)
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", "user1", "password1")

	/*database, _ := sql.Open("sqlite3", "./gopher.db")

	statemant, _ := database.Prepare("CREATE TABLE IF NOT EXISTS registrdatadb (login TEXT PRIMARY KEY, password TEXT)")
	statemant.Exec()*/
}

func main() {
	http.ListenAndServe(":8080", nil)
	DataBaseRecording()
	//http.HandleFunc("/api/registr", PersonalDataHandler)
}
