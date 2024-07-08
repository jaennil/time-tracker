package model

type User struct {
	Id             int64
	PassportSeries string `json:"passport_series" validate:"omitempty,number,len=4"`
	PassportNumber string `json:"passport_number" validate:"omitempty,number,len=6"`
	Name           string `json:"name" validate:"omitempty,max=255"`
	Surname        string `json:"surname" validate:"omitempty,max=255"`
	Patronymic     string `json:"patronymic" validate:"omitempty,max=255"`
	Address        string `json:"address"`
}
