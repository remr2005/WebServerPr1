package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func auth_and_administrating(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token") // принимание токена из параметров
	if token != "" {
		for _, cookie := range r.Cookies() {
			if cookie.Name == token {
				resp, err := http.Get("http://localhost:8001/getPersons")
				if err != nil {
					fmt.Println(err)
					http.Redirect(w, r, "/admins/forbidden", http.StatusSeeOther)
				}
				persons := make([]person, 0)
				_ = json.NewDecoder(resp.Body).Decode(&persons)
				fmt.Fprint(w, `<!DOCTYPE html>
				<html style="margin-top: 0;padding: 0" >
					<head>
						<meta charset="UTF-8">
						<meta name="viewport" content="width=device-width, initial-scale=1.0">
						<title>Администрирование</title>
						<style>
							/*Убирает отступы*/
							html,body {
								margin:0;
								padding:0;
							}
	
							/*Айди для таблички с exel*/
							#header{
								height: 80px;
								background-color: #B0E0E6;
								width: 100%;
							}
	
							/*Список пользователей*/
							#content{
								margin-top: 10px;
								height: 180px;
								background-color: #87CEFA;
							}
	
							/*Человек*/
							.person_table{
								border-width:2px;
								margin-left: 1%;
								margin-right: 1%;
								margin-top:3px;
								width: 98%;
								height: 30px;
								background-color: rgb(48, 173, 132);
							}
							
							/* внешние границы таблицы серого цвета толщиной 1px */
							table {
	   							border: 1px solid grey;
							}
							/* границы ячеек первого ряда таблицы */
							th {
	   						border: 1px solid grey;
							}
								/* границы ячеек тела таблицы */
							td {
	  					 border: 1px solid grey;
											}
							</style>
						</head>
							<body >
						<div id="header">Excel расписание</div>
						<div id="content">
							<table>
								<caption>Список пользователей</caption>
								  <tr>
									<th>GITID</th>
									<th>TELID</th>
									<th>Фамилия</th>
									<th>Имя</th>
									<th>Отчество</th>
									<th>Группа</th>
									<th>Студент</th>
									<th>Учитель</th>
									<th>Админ</th>
									<th>Дать/забрать админку</th>
									<th>Дать/забрать роль учителя</th>
									<th>Дать/забрать роль студента</th>
								  </tr>`)
				for _, i := range persons {
					fmt.Fprint(w, "<tr><th>"+i.GITID+"</th>")
					fmt.Fprint(w, "<th>"+i.TELID+"</th>")
					fmt.Fprint(w, "<th>"+i.SURNAME+"</th>")
					fmt.Fprint(w, "<th>"+i.NAMEP+"</th>")
					fmt.Fprint(w, "<th>"+i.FATHER_NAME+"</th>")
					fmt.Fprint(w, "<th>"+i.GROUPP+"</th>")
					if i.STUDENT == 1 {
						fmt.Fprint(w, `<th>1</th>`)
					} else {
						fmt.Fprint(w, `<th>0</th>`)
					}
					if i.LEHRER == 1 {
						fmt.Fprint(w, `<th>1</th>`)
					} else {
						fmt.Fprint(w, `<th>0</th>`)
					}
					if i.ADMINP == 1 {
						fmt.Fprint(w, `<th>1</th>`)
					} else {
						fmt.Fprint(w, `<th>0</th>`)
					}
					if i.ADMINP == 1 {
						fmt.Fprint(w, "<th><a href=\"http://localhost:8001/changeVar/tel?tel_id="+i.TELID+"&vars=ADMINP&sets=0\">отобрать</a></th>")
					} else {
						fmt.Fprint(w, "<th><a href=\"http://localhost:8001/changeVar/tel?tel_id="+i.TELID+"&vars=ADMINP&sets=1\">дать</a></th>")
					}
					if i.LEHRER == 1 {
						fmt.Fprint(w, "<th><a href=\"http://localhost:8001/changeVar/tel?tel_id="+i.TELID+"&vars=LEHRER&sets=0\">отобрать</a></th>")
					} else {
						fmt.Fprint(w, "<th><a href=\"http://localhost:8001/changeVar/tel?tel_id="+i.TELID+"&vars=LEHRER&sets=1\">дать</a></th>")
					}
					if i.STUDENT == 1 {
						fmt.Fprint(w, "<th><a href=\"http://localhost:8001/changeVar/tel?tel_id="+i.TELID+"&vars=STUDENT&sets=0\">отобрать</a></th>")
					} else {
						fmt.Fprint(w, "<th><a href=\"http://localhost:8001/changeVar/tel?tel_id="+i.TELID+"&vars=STUDENT&sets=1\">дать</a></th>")
					}
					fmt.Fprint(w, `</tr>`)
				}
				fmt.Fprintf(w, `</table>
								  </div>
							  </body>
						  </html>`)
			} else {
				http.Redirect(w, r, "/admins/forbidden", http.StatusSeeOther)
			}
		}
	} else {
		http.Redirect(w, r, "/admins/forbidden", http.StatusSeeOther)
	}

}
