package model

type User struct {
	Id             int
	PassportSeries string
	PassportNumber string
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
}
