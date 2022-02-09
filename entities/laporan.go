package entities

type Laporan struct {
	Id         int64  `gorm:"primary_key;auto_increment" db:"id"`
	Date       string `gorm:"type:date" db:"date"`
	Judul_lap  string `gorm:"type:varchar(255)" db:"judul_lap"`
	Isi_lap    string `gorm:"type:text" db:"isi_lap"`
	Divisi     string `gorm:"type:varchar(255)" db:"divisi"`
	Status     uint64 `gorm:"type:int" db:"status"`
	Nama_siswa string `gorm:"type:varchar(255)" db:"nama_siswa"`
	Pemberi    string `gorm:"type:varchar(255)" db:"pemberi"`
	Tugas_name string `gorm:"type:varchar(255)" db:"tugas_name"`
	UserID     uint64 `gorm:"column:user_id" db:"user_id"`
	TugasID    uint64 `gorm:"column:tugas_id" db:"tugas_id"`
	User       User
	Tugas      Tugas
}
