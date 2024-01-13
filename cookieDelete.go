package main

import "net/http"

func cookieDelete(w http.ResponseWriter, r *http.Request) {

	tokenDel := r.URL.Query().Get("tokenDel")
	token := r.URL.Query().Get("tokenSet")

	new_cookie := http.Cookie{
		Name:   tokenDel,
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(w, &new_cookie)

	http.Redirect(w, r, "http://localhost:8000/set?token="+token, http.StatusSeeOther)
}
