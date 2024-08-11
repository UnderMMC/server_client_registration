package main

import (
	//"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
)

type RegistrData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func PersonalDataHandler(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusOK)
	data := RegistrData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//err = DataBaseRecording(data.Name, data.Password)
	/*if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}*/

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

/*func DataBaseRecording(name, password string) error {
	connStr := "user=postgres dbname=postgres password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close() // Закрываем базу данных в конце функции

	// Проверяем соединение
	err = db.Ping()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (name, password) VALUES ($1, $2)", name, password)
	if err != nil {
		return err
	}
	return nil
}*/

func main() {

	http.HandleFunc("/api/registr", PersonalDataHandler)
	http.ListenAndServe(":8080", nil)
}
