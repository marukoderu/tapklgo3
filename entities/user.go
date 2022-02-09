package entities

type User struct {
	Id       int64     `gorm:"primary_key;auto_increment" db:"id" form:"id"`
	Date     string    `gorm:"type:date" db:"date"`
	Name     string    `gorm:"type:varchar(255)" db:"name" form:"name"`
	Email    string    `gorm:"type:varchar(255)" db:"email" form:"email"`
	Password string    `gorm:"type:longtext" db:"password" form:"password"`
	Photo    string    `gorm:"type:varchar(255)" db:"photo" form:"photo"`
	Divisi   string    `gorm:"type:varchar(255)" db:"divisi" form:"divisi"`
	Level    string    `gorm:"type:varchar(1)" db:"level" form:"level"`
	Tugass   []Tugas   `gorm:"ForeignKey:UserID"`
	Laporans []Laporan `gorm:"ForeignKey:UserID"`
	Jadwals  []Jadwal  `gorm:"ForeignKey:UserID"`
}
