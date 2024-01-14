package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func odd(reader io.Reader) {
	drop_odd()
	file, err := excelize.OpenReader(reader)
	if err != nil {
		fmt.Println(err)
	}
	name_days := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	for _, d := range name_days {
		var day kday
		day.Day = d
		rows := file.GetRows(d)
		if err != nil {
			continue
		}
		var res string = "{\"day\":\"" + d + "\",\"array\":["
		var _para_ para
		for j, row := range rows {
			if j == 0 {
				continue
			}
			for i, col := range row {
				if i == 0 {
					_para_.Number = col
				} else if i == 1 {
					_para_.Class = col
				} else if i == 2 {
					_para_.Teacher = col
				} else if i == 3 {
					_para_.Comment = col
				}
			}
			day.paras[j] = _para_

			a, _ := json.Marshal(_para_)
			res = res + string(a) + ","

		}

		res = res[:len(res)-1]
		res = res + "]}"
		//fmt.Println(res)
		r := bytes.NewReader([]byte(res))

		_, err = http.Post("http://127.0.0.1:8080/schedule/odd/add", "application/json", r)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func even(reader io.Reader) {
	drop_even()
	file, err := excelize.OpenReader(reader)
	if err != nil {
		fmt.Println(err)
	}
	name_days := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	for _, d := range name_days {
		var day kday
		day.Day = d
		rows := file.GetRows(d)
		if err != nil {
			continue
		}
		var res string = "{\"day\":\"" + d + "\",\"array\":["
		var _para_ para
		for j, row := range rows {
			if j == 0 {
				continue
			}
			for i, col := range row {
				if i == 0 {
					_para_.Number = col
				} else if i == 1 {
					_para_.Class = col
				} else if i == 2 {
					_para_.Teacher = col
				} else if i == 3 {
					_para_.Comment = col
				}
			}
			day.paras[j] = _para_

			a, _ := json.Marshal(_para_)
			res = res + string(a) + ","

		}

		res = res[:len(res)-1]
		res = res + "]}"
		//fmt.Println(res)
		r := bytes.NewReader([]byte(res))

		_, err = http.Post("http://127.0.0.1:8080/schedule/even/add", "application/json", r)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func drop_even() {
	url := "http://127.0.0.1:8080/schedule/drop/even"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

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
}

func drop_odd() {
	url := "http://127.0.0.1:8080/schedule/drop/odd"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

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
}
