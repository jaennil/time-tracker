package model

// TODO: find out why validate have omitempty here
type User struct {
	Id             int64   `json:"id" db:"user_id" example:"1" minimum:"1"`
	PassportSeries string  `json:"passport_series" validate:"omitempty,number,len=4" db:"passport_series" example:"1234" minLength:"4" maxLength:"4"`
	PassportNumber string  `json:"passport_number" validate:"omitempty,number,len=6" db:"passport_number" example:"567890" minLength:"6" maxLength:"6"`
	Name           string  `json:"name" validate:"omitempty,max=255" db:"name" example:"Иван"`
	Surname        string  `json:"surname" validate:"omitempty,max=255" db:"surname" example:"Иванов"`
	Patronymic     *string `json:"patronymic" validate:"omitempty,max=255" db:"patronymic" example:"Иванович"`
	Address        string  `json:"address" db:"address" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
}

type CreateUser struct {
	Passport string `json:"passportNumber" binding:"required" validate:"passport" example:"1234 567890" minLength:"11" maxLength:"11"`
}
