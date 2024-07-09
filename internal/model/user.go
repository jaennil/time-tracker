package model

type User struct {
	Id             int64   `json:"id" db:"user_id"`
	PassportSeries string  `json:"passport_series" validate:"omitempty,number,len=4" db:"passport_series"`
	PassportNumber string  `json:"passport_number" validate:"omitempty,number,len=6" db:"passport_number"`
	Name           string  `json:"name" validate:"omitempty,max=255" db:"name"`
	Surname        string  `json:"surname" validate:"omitempty,max=255" db:"surname"`
	Patronymic     *string `json:"patronymic" validate:"omitempty,max=255" db:"patronymic"`
	Address        string  `json:"address" db:"address"`
}
