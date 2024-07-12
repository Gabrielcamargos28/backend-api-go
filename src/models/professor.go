package models

type Professor struct {
	Id    int    `gorm:"type:int;primary_key`
	Nome  string `gorm:type:varchar(255)`
	Email string `gorm:type:varchar(255)`
	CPF   string `gorm:type:varchar(255)`
}

func (Professor) TableName() string {
	return "professor"
}
