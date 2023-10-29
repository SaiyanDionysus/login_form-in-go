package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "your_user"
	password = "your_password"
	dbname   = "your_database"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			username := r.FormValue("username")
			email := r.FormValue("email")
			password := r.FormValue("password")

			// Сохраняем данные в базу данных
			db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			// Определяем структуру таблицы
			_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id serial PRIMARY KEY, username text, email text, password text)")
			if err != nil {
				log.Fatal(err)
			}

			// Вставляем данные в таблицу
			_, err = db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", username, email, password)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintf(w, "Регистрация успешно завершена для пользователя: %s", username)
		} else {
			http.ServeFile(w, r, "index.html")
		}
	})

	http.ListenAndServe(":8080", nil)
}
