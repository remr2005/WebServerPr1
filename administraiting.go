package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

				client := &http.Client{}
				req, err := http.NewRequest("GET", "http://localhost:8001/getPersons/count", nil)

				if err != nil {
					fmt.Println(err)
					return
				}
				res, err := client.Do(req)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer res.Body.Close()

				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					fmt.Println(err)
					return
				}
				int_a, err := strconv.Atoi(string(body))
				if err != nil {
					fmt.Println(err)
					return
				}
				int_a *= 30
				fmt.Fprint(w, `<!DOCTYPE html>
				<html style="margin-top: 0;padding: 0" >
					<head>
						<meta charset="UTF-8">
						<meta name="viewport" content="width=device-width, initial-scale=1.0">
						<title>Администрирование</title>
						<style>
						label {
							display: inline-block;
							background: #ddd;
							border: 1px outset #ccc;
							border-radius: .3em;
							padding: .3em 1em;
							margin: .5em;
						}
						
						label:active {
						  border-style: inset;
						}
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
								width: 98%;height:`)
				fmt.Println(int_a)
				fmt.Fprint(w, int_a)
				fmt.Fprint(w, `px;
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
						<div id="header"><a href="`)
				fmt.Fprint(w, "http://localhost:8000/admins/excel?token="+token)
				fmt.Fprint(w, `">Excel расписание<br></a>
 						</div>
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
						fmt.Fprint(w, "<th><a href=\"http://localhost:8001/changeVar/tel?tel_id="+i.TELID+"&vars=ADMINP&sets=0&state="+token+"\">отобрать</a></th>")
					} else {
						fmt.Fprint(w, "<th><a href=\"http://localhost:8001/changeVar/tel?tel_id="+i.TELID+"&vars=ADMINP&sets=1&state="+token+"\">дать</a></th>")
					}
					if i.LEHRER == 1 {
						fmt.Fprint(w, "<th><a href=\"http://localhost:8001/changeVar/tel?tel_id="+i.TELID+"&vars=LEHRER&sets=0&state="+token+"\">отобрать</a></th>")
					} else {
						fmt.Fprint(w, "<th><a href=\"http://localhost:8001/changeVar/tel?tel_id="+i.TELID+"&vars=LEHRER&sets=1&state="+token+"\">дать</a></th>")
					}
					if i.STUDENT == 1 {
						fmt.Fprint(w, "<th><a href=\"http://localhost:8001/changeVar/tel?tel_id="+i.TELID+"&vars=STUDENT&sets=0&state="+token+"\">отобрать</a></th>")
					} else {
						fmt.Fprint(w, "<th><a href=\"http://localhost:8001/changeVar/tel?tel_id="+i.TELID+"&vars=STUDENT&sets=1&state="+token+"\">дать</a></th>")
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
