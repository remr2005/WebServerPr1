package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type buf struct {
	Href   string `json:"href"`
	Href_  string `json:"method"`
	Href__ bool   `json:"templated"`
}

func even() {
	var a buf
	url := "https://cloud-api.yandex.net/v1/disk/resources/download?path=excel/расписаниеНЧЕТ.xlsx"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "OAuth y0_AgAAAABwu1cYAAsQwAAAAAD2TKlteu5QSy_YSb2JDmuY1k1reRip23c")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&a); err != nil {
		fmt.Println(err)
	}

	req, err = http.NewRequest(method, a.Href, nil)
}

func odd() {
	var a buf
	url := "https://cloud-api.yandex.net/v1/disk/resources/download?path=excel/расписаниеЧЕТ.xlsx"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "OAuth y0_AgAAAABwu1cYAAsQwAAAAAD2TKlteu5QSy_YSb2JDmuY1k1reRip23c")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&a); err != nil {
		fmt.Println(err)
	}

	req, err = http.NewRequest(method, a.Href, nil)
}

func main() {
	even()
	odd()
}
