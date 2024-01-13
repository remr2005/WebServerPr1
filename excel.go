package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func excel(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token") // принимание токена из параметров
	if token != "" {
		for _, cookie := range r.Cookies() {
			if cookie.Name == token {
				dec_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
					return []byte(SECRET), nil
				})
				if err != nil {
					fmt.Println(err)
				}
				payload, ok := dec_token.Claims.(jwt.MapClaims)
				if ok && (payload["expires_at"].(float64) >= float64(time.Now().Unix())) && payload["admin"].(bool) {
					fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>Upload File</title>
  </head>
  <body>
    <form
      enctype="multipart/form-data"
	  
      action="`)
					fmt.Fprint(w, "http://localhost:8000/upload?token="+token)
					fmt.Fprint(w, `"
      method="post"
    >
      <input type="file" name="myFile" />
      <input type="submit" value="upload" />
    </form>
  </body>
</html>`)
				} else {
					http.Redirect(w, r, "/admins/forbidden", http.StatusSeeOther)
				}
			} else {
				http.Redirect(w, r, "/admins/forbidden", http.StatusSeeOther)
			}

		}
	} else {
		http.Redirect(w, r, "/admins/forbidden", http.StatusSeeOther)
	}

}
