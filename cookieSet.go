package main

import "net/http"

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	// Создаём и заполняем новую структуру http.Cookie
	cookie := http.Cookie{
		Name:     token,
		Value:    token,
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/admins?token="+token, http.StatusSeeOther)
}
