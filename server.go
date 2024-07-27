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
	connStr := "user=postgres dbname=postgres password=4125 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка при открытии базы данных: %v", err)
	}
	defer db.Close() // Закрываем базу данных в конце функции

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
	if err != nil {
		log.Fatalf("Ошибка при вставке данных: %v", err)
	}

	log.Println("Данные успешно вставлены в таблицу users")
}

func main() {
	http.HandleFunc("/api/registr", PersonalDataHandler)
	DataBaseRecording()
	http.ListenAndServe(":8080", nil)
}
