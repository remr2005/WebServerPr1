package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func login_from_bot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //формат json

	var juest jwt_string
	_ = json.NewDecoder(r.Body).Decode(&juest) // принимаю данные json записывая их в juest

	tokenString := juest.TEXT // строка с токеном
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	}) // расшифровка jwt

	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
	}
	payload, ok := token.Claims.(jwt.MapClaims) // запись информмации из jwt в payload

	if ok && token.Valid && payload["admin"].(bool) && (payload["expires_at"].(float64) >= float64(time.Now().Unix())) {
		tokeExpiresAt := time.Now().Add(time.Minute * time.Duration(15)) // время жизни токена
		payload := jwt.MapClaims{
			"id":         juest.ID,
			"admin":      true,
			"expires_at": tokeExpiresAt.Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) // создание нового токена
		NewTokenString, err := token.SignedString([]byte(SECRET))   // подпись токена ключём
		if err != nil {
			fmt.Println(err)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var linkk links
		linkk.LINK = "/set?token=" + NewTokenString //создание ссылки
		json.NewEncoder(w).Encode(linkk)            // отправление ссылки
	} else {
		var linkk links
		linkk.LINK = "/admins/forbidden" //создание ссылки
		json.NewEncoder(w).Encode(linkk)
	}
}
