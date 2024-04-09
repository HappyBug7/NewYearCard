package model

type Card struct {
	BaseModel
	FromUser string
	ToUser   string
	Content  string
}
