package main

import (
	//"database/sql"
	//"encoding/json"
	//"fmt"
	_ "go-sql-driver/mysql"
	"gorilla/mux"
	"log"
	"net/http"

	//"time"

	//"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Глобальные переменные
// Куки
var SECRET = "880041"

// Буферная строка
var buffer_string = "exampleCookie"

// var notСonfirmedSessions map[string]*sessions.Session
var store = sessions.NewCookieStore([]byte(SECRET))

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Точка входа
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/set", setCookieHandler).Methods("GET")
	r.HandleFunc("/admins", auth_and_administrating).Methods("GET")
	r.HandleFunc("/admins/forbidden", forbidden_page).Methods("GET")
	r.HandleFunc("/admins/login", login_from_bot).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}

// admin/login&user=(user_jwt)
//Если jwt есть в cooke то допустить к администрированию, иначе на мороз

// admin если cooke есть на продолжить осмотр, иначе на страницу login

// во вкладке admin должна быть html страничка со всеми пользователями и excel
