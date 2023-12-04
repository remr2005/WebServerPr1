package main

import (
	"fmt"
	"net/http"
)

func forbidden_page(w http.ResponseWriter, r *http.Request) {
	page := `<!DOCTYPE html>
                <html>
                <head>
                    <meta charset="UTF-8">
                    <meta name="viewport" content="width=device-width, initial-scale=1.0">
                    <title>Доступ запрещен</title>
                </head>
                <body>
                    <div>Доступ запрещен</div>
                </body>
                </html>`
	fmt.Fprint(w, page)
}
