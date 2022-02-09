package entities

type Tugas struct {
	Id        int64  `gorm:"primary_key;auto_increment" db:"id"`
	Date      string `gorm:"type:date" db:"date"`
	Judul     string `gorm:"type:varchar(255)" db:"judul"`
	Isi       string `gorm:"type:text" db:"isi"`
	Status    uint64 `gorm:"type:int" db:"status"`
	Divisi    string `gorm:"type:varchar(255)" db:"divisi"`
	UserName  string `gorm:"type:varchar(255)" db:"user_name"`
	UserLevel uint64 `gorm:"type:int" db:"user_level"`
	UserID    uint64 `gorm:"column:user_id" db:"user_id"`
	User      User
	Laporans  []Laporan `gorm:"ForeignKey:TugasID"`
}
