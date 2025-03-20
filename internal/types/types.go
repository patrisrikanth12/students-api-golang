package types

type Student struct {
	Id int  `json:"id"`
	Name string `validate:"required" json:"name"`
	Email string  `validate:"required" json:"email"`
	Mobile string  `validate:"required" json:"mobile"`
}