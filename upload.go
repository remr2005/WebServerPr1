package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)
	state := r.URL.Query().Get("token")
	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	if handler.Filename == "расписаниеЧЕТ.xlsx" {
		odd(file)
	} else if handler.Filename == "расписаниеНЧЕТ.xlsx" {
		even(file)
	}

	dec_token, err := jwt.Parse(state, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		fmt.Println(err)
	}
	payload, ok := dec_token.Claims.(jwt.MapClaims)

	if ok && payload["admin"].(bool) && (payload["expires_at"].(float64) >= float64(time.Now().Unix())) {
		tokeExpiresAt := time.Now().Add(time.Minute * time.Duration(15)) // время жизни токена
		payload_second := jwt.MapClaims{
			"id":         payload["id"].(string),
			"admin":      true,
			"expires_at": tokeExpiresAt.Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload_second) // создание нового токена
		NewTokenString, err := token.SignedString([]byte(SECRET))
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "http://localhost:8000/delete?tokenDel="+state+"&tokenSet="+NewTokenString, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "http://localhost:8000/admins/forbidden"+state, http.StatusSeeOther)
	}
}
