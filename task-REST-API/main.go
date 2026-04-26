package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

	case http.MethodDelete:
		idStr := strings.TrimPrefix(r.URL.Path, "/users/") // сохраняем ID в строку из ссылки
		id, err := strconv.Atoi(idStr) // конвертируем из строки в int и проверяем на ошибки
		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		
		for i, u := range users{ // удаление id из слайса
			if u.ID == id {
				users = append(users[:i], users[i+1:]...)
				w.WriteHeader(204)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/", usersHandler)
	fmt.Println("Сервер на :8080")
	http.ListenAndServe(":8080", nil)
}

