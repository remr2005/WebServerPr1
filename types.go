package main

// строка принимаемых данных
type jwt_string struct {
	ID   string `json: "ID"`
	TEXT string `json: "TEXT"`
}

// ссылка в виде json
type links struct {
	LINK string `json: "LINK"`
}

type person struct {
	GITID       string `json: "GITID"`
	TELID       string `json: "TELID"`
	SURNAME     string `json: "SURNAME"`
	NAMEP       string `json: "NAMEP"`
	FATHER_NAME string `json: "FATHER_NAME"`
	GROUPP      string `json: "GROUPP"`
	STUDENT     uint8  `json: "STUDENT"`
	LEHRER      uint8  `json: "LEHRER"`
	ADMINP      uint8  `json: "ADMINP"`
}
