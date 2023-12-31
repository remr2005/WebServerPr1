package main

import (
	"fmt"
	"net/http"
)

func excel(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token") // принимание токена из параметров
	if token != "" {
		for _, cookie := range r.Cookies() {
			if cookie.Name == token {
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
			}
		}
	} else {
		http.Redirect(w, r, "/admins/forbidden", http.StatusSeeOther)
	}

}
