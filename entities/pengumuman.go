package entities

type Pengumuman struct {
	Id     int64  `gorm:"primary_key;auto_increment" db:"id"`
	Date   string `gorm:"type:date" db:"date"`
	Judul  string `gorm:"type:varchar(255)" db:"judul"`
	Isi    string `gorm:"type:text" db:"isi"`
	Name   string `gorm:"type:varchar(255)" db:"name"`
	Divisi string `gorm:"type:varchar(255)" db:"divisi"`
	Status uint64 `gorm:"type:int" db:"status"`
	UserID uint64 `gorm:"column:user_id" db:"user_id"`
	User   User
}
