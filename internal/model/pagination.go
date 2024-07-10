package model

type Pagination struct {
	Page     int `form:"page,default=1" validate:"omitempty,min=1" example:"1" minimum:"1" `
	PageSize int `form:"page_size,default=10" validate:"omitempty,min=1" example:"10" minimum:"1"`
}
