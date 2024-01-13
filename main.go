package main

import (
	_ "go-sql-driver/mysql"
	"gorilla/mux"
	"log"
	"net/http"

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

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// http://localhost:8000/admins?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzX2F0IjoxNzAyMTE2MTQ4LCJpZCI6IjEyMzEzMiJ9.tETrU4Z_fqje1L-qsq1hATZ5qN39WPFQTt9Emg0oTq8
// Точка входа
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/set", setCookieHandler).Methods("GET")
	r.HandleFunc("/delete", cookieDelete).Methods("GET")
	r.HandleFunc("/shedule/odd", auth_and_administrating).Methods("GET")
	r.HandleFunc("/admins", auth_and_administrating).Methods("GET")
	r.HandleFunc("/admins/excel", excel).Methods("GET")
	//r.HandleFunc("/shedule/url", urlExcel_to_json).Methods("GET")
	r.HandleFunc("/admins/forbidden", forbidden_page).Methods("GET")
	r.HandleFunc("/admins/login", login_from_bot).Methods("POST")
	r.HandleFunc("/upload", uploadFile).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}

// admin/login&user=(user_jwt)
//Если jwt есть в cooke то допустить к администрированию, иначе на мороз

// admin если cooke есть на продолжить осмотр, иначе на страницу login

// во вкладке admin должна быть html страничка со всеми пользователями и excel
