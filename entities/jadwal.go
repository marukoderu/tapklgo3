package entities

type Jadwal struct {
	Id   int64  `gorm:"primary_key;auto_increment" db:"id"`
	Day  string `gorm:"type:varchar(255)" db:"judul"`
	Name string `gorm:"type:varchar(255)" db:"status"`
}
