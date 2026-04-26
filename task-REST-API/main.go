package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var users = []User{
	{ID: 1, Name: "Вика"},
	{ID: 2, Name: "Павел"},
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("→ %s %s\n", r.Method, r.URL.Path)
	w.Header().Set("Content-type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		if user.Name == "" { // проверка на пустое имя
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "имя не может быть пустым")
			return
		}

		user.ID = len(users) + 1
		users = append(users, user)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/users", usersHandler)
	fmt.Println("Сервер на :8080")
	http.ListenAndServe(":8080", nil)
}
