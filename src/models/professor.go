package models

type Professor struct {
	Id    int    `gorm:"type:int;primary_key`
	Nome  string `gorm:type:varchar(255)`
	Email string
	CPF   string
}
