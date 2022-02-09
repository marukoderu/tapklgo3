package entities

type Divisi struct {
	Id  int64  `gorm:"primary_key;auto_increment" db:"id" form:"id"`
	Div string `gorm:"type:varchar(255)" db:"divisi" form:"divisi"`
}
