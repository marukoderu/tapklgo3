package models

import (
	"fmt"

	"github.com/raisunee/tapklgo3/config"
	"github.com/raisunee/tapklgo3/entities"
)

type UserModel struct {
}

// func (*UserModel) QueryUser(email string) ([]entities.User, error) {
// 	db, err := config.GetDB()
// 	if err != nil {
// 		return nil, err
// 	} else {
// 		var users []entities.User
// 		db.Table("users").Where("email = ?", email).Select("id, name, email, password, level").Scan(&users)
// 		return users, nil
// 	}
// }

func (*UserModel) QueryUser(email string) (entities.User, error) {
	db, err := config.GetDB()
	if err != nil {
		return entities.User{}, err
	} else {
		var users = entities.User{}
		err = db.QueryRow("SELECT * FROM users WHERE email=?", email).
			Scan(
				&users.Id,
				&users.Date,
				&users.Name,
				&users.Email,
				&users.Password,
				&users.Photo,
				&users.Divisi,
				&users.Level,
			)
		if err != nil {
			return entities.User{}, err
		} else {
			return users, nil
		}
	}
}

// Show By Level
func (*UserModel) FindByLevel() ([]entities.User, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, date, name, email, divisi, level, password FROM users WHERE level = 0 OR level = 1")
		if err2 != nil {
			return nil, err2
		} else {
			var users []entities.User
			for rows.Next() {
				var user entities.User
				rows.Scan(&user.Id, &user.Date, &user.Name, &user.Email, &user.Divisi, &user.Level, &user.Password)
				users = append(users, user)
			}
			return users, nil
		}
	}
}

func (*UserModel) FindStaff() ([]entities.User, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, date, name, email, divisi, level, password FROM users WHERE level = 2 OR level = 4")
		if err2 != nil {
			return nil, err2
		} else {
			var users []entities.User
			for rows.Next() {
				var user entities.User
				rows.Scan(&user.Id, &user.Date, &user.Name, &user.Email, &user.Divisi, &user.Level, &user.Password)
				users = append(users, user)
			}
			return users, nil
		}
	}
}

// Show By Level
func (*UserModel) FindAccAdmin(user_id string) ([]entities.User, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, name, email, photo, divisi, level FROM users WHERE id=?", user_id)
		if err2 != nil {
			return nil, err2
		} else {
			var users []entities.User
			for rows.Next() {
				var user entities.User
				rows.Scan(&user.Id, &user.Name, &user.Email, &user.Photo, &user.Divisi, &user.Level)
				users = append(users, user)
			}
			return users, nil
		}
	}
}

func (*UserModel) UpdatePass(user entities.User) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE users SET password = ? WHERE id = ?", user.Password, user.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

// Show By Level
func (*UserModel) FindDivisi() ([]entities.Divisi, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, divisi FROM divisi")
		if err2 != nil {
			return nil, err2
		} else {
			var divisis []entities.Divisi
			for rows.Next() {
				var divisi entities.Divisi
				rows.Scan(&divisi.Id, &divisi.Div)
				divisis = append(divisis, divisi)
			}
			return divisis, nil
		}
	}
}

// Show By Id
func (*UserModel) FindUser(email string) (entities.User, error) {
	db, err := config.GetDB()
	if err != nil {
		return entities.User{}, err
	} else {
		var users = entities.User{}
		sqlState := "SELECT * FROM users WHERE email = $1"
		row := db.QueryRow(sqlState, email)
		err := row.Scan(&users.Id, &users.Email, &users.Name, &users.Password)

		if err != nil {
			return entities.User{}, err
		} else {
			return users, nil
		}
	}
}

func (*UserModel) CreateStaff(user *entities.User) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("INSERT INTO users(name, email, password, level) VALUES(?,?,?,2)", user.Name, user.Email, user.Password)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) CreateUser(user *entities.User) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("INSERT INTO users(name, email, password, divisi, level) VALUES(?,?,?,?, 1)", user.Name, user.Email, user.Password, user.Divisi)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) CreateUserStaff(user *entities.User) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("INSERT INTO users(name, email, password, level) VALUES(?,?,?,2)", user.Name, user.Email, user.Password)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) Create(user *entities.User) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("INSERT INTO users(name, email, password, level) VALUES(?,?,?,?)", user.Name, user.Email, user.Password, user.Level)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) UpdateSiswa(user entities.User) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE users SET name = ?, email = ?, divisi = ? WHERE id = ?", user.Name, user.Email, user.Divisi, user.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) UpdateAccountAdm(user entities.User) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) UpdateAccountStaff(user entities.User) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) UpdatePhotoStaff(user entities.User) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE users SET photo = ? WHERE id = ?", user.Photo, user.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) UpdatePhotoSis(user entities.User) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE users SET photo = ? WHERE id = ?", user.Photo, user.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) UpdateAccountSis(user entities.User) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) ActiveStaff(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE users SET level=2 WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) InactiveStaff(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE users SET level=4 WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) ActiveSiswa(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE users SET level=1 WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) InactiveSiswa(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE users SET level=0 WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) DeleteSiswa(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("DELETE FROM users WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) CountSiswaAktif() (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		rows, err2 := db.Query("SELECT COUNT(level) as count_siswa FROM users WHERE level=1")
		if err2 != nil {
			return 0, err2
		} else {
			var count_siswa int64
			for rows.Next() {
				rows.Scan(&count_siswa)
			}
			return count_siswa, nil
		}
	}
}

func (*UserModel) CreateDivisi(divisi *entities.Divisi) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("INSERT INTO divisi(divisi) VALUES(?)", divisi.Div)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) UpdateDivisi(divisi entities.Divisi) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE divisi SET divisi = ? WHERE id = ?", divisi.Div, divisi.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) DeleteDivisi(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("DELETE FROM divisi WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) DeleteJadwal(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("DELETE FROM jadwal WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

// JADWAL BY DAY
func (*UserModel) FindJadwalMon() ([]entities.Jadwal, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, day, name FROM jadwal WHERE day='Monday'")
		if err2 != nil {
			return nil, err2
		} else {
			var jadwals []entities.Jadwal
			for rows.Next() {
				var jadwal entities.Jadwal
				rows.Scan(&jadwal.Id, &jadwal.Day, &jadwal.Name)
				jadwals = append(jadwals, jadwal)
			}
			return jadwals, nil
		}
	}
}

func (*UserModel) FindJadwalTue() ([]entities.Jadwal, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, day, name FROM jadwal WHERE day='Tuesday'")
		if err2 != nil {
			return nil, err2
		} else {
			var jadwals []entities.Jadwal
			for rows.Next() {
				var jadwal entities.Jadwal
				rows.Scan(&jadwal.Id, &jadwal.Day, &jadwal.Name)
				jadwals = append(jadwals, jadwal)
			}
			return jadwals, nil
		}
	}
}

func (*UserModel) FindJadwalWed() ([]entities.Jadwal, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, day, name FROM jadwal WHERE day='Wednesday'")
		if err2 != nil {
			return nil, err2
		} else {
			var jadwals []entities.Jadwal
			for rows.Next() {
				var jadwal entities.Jadwal
				rows.Scan(&jadwal.Id, &jadwal.Day, &jadwal.Name)
				jadwals = append(jadwals, jadwal)
			}
			return jadwals, nil
		}
	}
}

func (*UserModel) FindJadwalThu() ([]entities.Jadwal, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, day, name FROM jadwal WHERE day='Thursday'")
		if err2 != nil {
			return nil, err2
		} else {
			var jadwals []entities.Jadwal
			for rows.Next() {
				var jadwal entities.Jadwal
				rows.Scan(&jadwal.Id, &jadwal.Day, &jadwal.Name)
				jadwals = append(jadwals, jadwal)
			}
			return jadwals, nil
		}
	}
}

func (*UserModel) FindJadwalFri() ([]entities.Jadwal, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, day, name FROM jadwal WHERE day='Friday'")
		if err2 != nil {
			return nil, err2
		} else {
			var jadwals []entities.Jadwal
			for rows.Next() {
				var jadwal entities.Jadwal
				rows.Scan(&jadwal.Id, &jadwal.Day, &jadwal.Name)
				jadwals = append(jadwals, jadwal)
			}
			return jadwals, nil
		}
	}
}

// JADWAL BY DAY AND SISWA
func (*UserModel) FindJadwaSisMon(user_id string) ([]entities.Jadwal, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, day, name FROM jadwal WHERE day='Monday' AND user_id=?", user_id)
		if err2 != nil {
			return nil, err2
		} else {
			var jadwals []entities.Jadwal
			for rows.Next() {
				var jadwal entities.Jadwal
				rows.Scan(&jadwal.Id, &jadwal.Day, &jadwal.Name)
				jadwals = append(jadwals, jadwal)
			}
			return jadwals, nil
		}
	}
}

func (*UserModel) FindJadwalSisTue(user_id string) ([]entities.Jadwal, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, day, name FROM jadwal WHERE day='Tuesday' AND user_id=?", user_id)
		if err2 != nil {
			return nil, err2
		} else {
			var jadwals []entities.Jadwal
			for rows.Next() {
				var jadwal entities.Jadwal
				rows.Scan(&jadwal.Id, &jadwal.Day, &jadwal.Name)
				jadwals = append(jadwals, jadwal)
			}
			return jadwals, nil
		}
	}
}

func (*UserModel) FindJadwalSisWed(user_id string) ([]entities.Jadwal, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, day, name FROM jadwal WHERE day='Wednesday' AND user_id=?", user_id)
		if err2 != nil {
			return nil, err2
		} else {
			var jadwals []entities.Jadwal
			for rows.Next() {
				var jadwal entities.Jadwal
				rows.Scan(&jadwal.Id, &jadwal.Day, &jadwal.Name)
				jadwals = append(jadwals, jadwal)
			}
			return jadwals, nil
		}
	}
}

func (*UserModel) FindJadwalSisThu(user_id string) ([]entities.Jadwal, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, day, name FROM jadwal WHERE day='Thursday' AND user_id=?", user_id)
		if err2 != nil {
			return nil, err2
		} else {
			var jadwals []entities.Jadwal
			for rows.Next() {
				var jadwal entities.Jadwal
				rows.Scan(&jadwal.Id, &jadwal.Day, &jadwal.Name)
				jadwals = append(jadwals, jadwal)
			}
			return jadwals, nil
		}
	}
}

func (*UserModel) FindJadwalSisFri(user_id string) ([]entities.Jadwal, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, day, name FROM jadwal WHERE day='Friday' AND user_id=?", user_id)
		if err2 != nil {
			return nil, err2
		} else {
			var jadwals []entities.Jadwal
			for rows.Next() {
				var jadwal entities.Jadwal
				rows.Scan(&jadwal.Id, &jadwal.Day, &jadwal.Name)
				jadwals = append(jadwals, jadwal)
			}
			return jadwals, nil
		}
	}
}

func (*UserModel) CreateJadwal(jadwal *entities.Jadwal) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("INSERT INTO jadwal(day, name) SELECT day, name FROM (SELECT ? AS day, ? AS name) AS tmp WHERE NOT EXISTS (SELECT day, name FROM jadwal WHERE day = ? AND name = ?) LIMIT 1;", jadwal.Day, jadwal.Name, jadwal.Day, jadwal.Name)
		if err2 != nil {
			fmt.Println("DATA EXISTS!")
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) FindPengumuman() ([]entities.Pengumuman, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, date, judul, isi, name, user_id, divisi, status FROM pengumuman ORDER BY id DESC")
		if err2 != nil {
			return nil, err2
		} else {
			var pengumumans []entities.Pengumuman
			for rows.Next() {
				var pengumuman entities.Pengumuman
				rows.Scan(&pengumuman.Id, &pengumuman.Date, &pengumuman.Judul, &pengumuman.Isi, &pengumuman.Name, &pengumuman.UserID, &pengumuman.Divisi, &pengumuman.Status)
				pengumumans = append(pengumumans, pengumuman)
			}
			return pengumumans, nil
		}
	}
}

func (*UserModel) FindPengumumanGeneral() ([]entities.Pengumuman, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, judul, isi, name, user_id, divisi, status FROM pengumuman WHERE divisi='General' AND status=1 ORDER BY id DESC")
		if err2 != nil {
			return nil, err2
		} else {
			var pengumumans []entities.Pengumuman
			for rows.Next() {
				var pengumuman entities.Pengumuman
				rows.Scan(&pengumuman.Id, &pengumuman.Judul, &pengumuman.Isi, &pengumuman.Name, &pengumuman.UserID, &pengumuman.Divisi, &pengumuman.Status)
				pengumumans = append(pengumumans, pengumuman)
			}
			return pengumumans, nil
		}
	}
}

func (*UserModel) FindPengumumanStaff(username string) ([]entities.Pengumuman, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, date, judul, isi, name, user_id, divisi, status FROM pengumuman WHERE name=?", username)
		if err2 != nil {
			return nil, err2
		} else {
			var pengumumans []entities.Pengumuman
			for rows.Next() {
				var pengumuman entities.Pengumuman
				rows.Scan(&pengumuman.Id, &pengumuman.Date, &pengumuman.Judul, &pengumuman.Isi, &pengumuman.Name, &pengumuman.UserID, &pengumuman.Divisi, &pengumuman.Status)
				pengumumans = append(pengumumans, pengumuman)
			}
			return pengumumans, nil
		}
	}
}

func (*UserModel) FindPengumumanSiswa(divisi string) ([]entities.Pengumuman, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, date, judul, isi, name, user_id, divisi, status FROM pengumuman WHERE divisi=? ORDER BY id DESC", divisi)
		if err2 != nil {
			return nil, err2
		} else {
			var pengumumans []entities.Pengumuman
			for rows.Next() {
				var pengumuman entities.Pengumuman
				rows.Scan(&pengumuman.Id, &pengumuman.Date, &pengumuman.Judul, &pengumuman.Isi, &pengumuman.Name, &pengumuman.UserID, &pengumuman.Divisi, &pengumuman.Status)
				pengumumans = append(pengumumans, pengumuman)
			}
			return pengumumans, nil
		}
	}
}

func (*UserModel) FindPengumumanStaffDash() ([]entities.Pengumuman, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, judul, isi, name, user_id, divisi, status FROM pengumuman WHERE name='Admin' ORDER BY id DESC LIMIT 1")
		if err2 != nil {
			return nil, err2
		} else {
			var pengumumans []entities.Pengumuman
			for rows.Next() {
				var pengumuman entities.Pengumuman
				rows.Scan(&pengumuman.Id, &pengumuman.Judul, &pengumuman.Isi, &pengumuman.Name, &pengumuman.UserID, &pengumuman.Divisi, &pengumuman.Status)
				pengumumans = append(pengumumans, pengumuman)
			}
			return pengumumans, nil
		}
	}
}

func (*UserModel) FindPengumumanSiswaDash(divisi string) ([]entities.Pengumuman, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, judul, isi, name, user_id, divisi, status FROM pengumuman WHERE divisi=? AND status=1 ORDER BY id DESC LIMIT 1", divisi)
		if err2 != nil {
			return nil, err2
		} else {
			var pengumumans []entities.Pengumuman
			for rows.Next() {
				var pengumuman entities.Pengumuman
				rows.Scan(&pengumuman.Id, &pengumuman.Judul, &pengumuman.Isi, &pengumuman.Name, &pengumuman.UserID, &pengumuman.Divisi, &pengumuman.Status)
				pengumumans = append(pengumumans, pengumuman)
			}
			return pengumumans, nil
		}
	}
}

func (*UserModel) ActivePengumuman(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE pengumuman SET status=1 WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) InactivePengumuman(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE pengumuman SET status=2 WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) CreatePengumuman(pengumuman *entities.Pengumuman) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("INSERT INTO pengumuman(judul, isi, divisi, name, user_id, status) VALUES(?,?,?,?,?,1)", pengumuman.Judul, pengumuman.Isi, pengumuman.Divisi, pengumuman.Name, pengumuman.UserID)
		if err2 != nil {
			fmt.Println("DATA EXISTS!")
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) UpdatePengumuman(pengumuman entities.Pengumuman) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE pengumuman SET judul = ?, isi = ?, divisi = ? WHERE id = ?", pengumuman.Judul, pengumuman.Isi, pengumuman.Divisi, pengumuman.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) DeletePengumuman(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("DELETE FROM pengumuman WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) DeleteLaporan(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("DELETE FROM laporan WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UserModel) StaffAktif() (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		rows, err2 := db.Query("SELECT COUNT(level) as count_staff FROM users WHERE level=2")
		if err2 != nil {
			return 0, err2
		} else {
			var count_staff int64
			for rows.Next() {
				rows.Scan(&count_staff)
			}
			return count_staff, nil
		}
	}
}

func (*UserModel) StaffInAktif() (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		rows, err2 := db.Query("SELECT COUNT(level) as count_siswa FROM users WHERE level=4")
		if err2 != nil {
			return 0, err2
		} else {
			var count_siswa int64
			for rows.Next() {
				rows.Scan(&count_siswa)
			}
			return count_siswa, nil
		}
	}
}

func (*UserModel) SiswaAktif() (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		rows, err2 := db.Query("SELECT COUNT(level) as count_siswa FROM users WHERE level=1")
		if err2 != nil {
			return 0, err2
		} else {
			var count_siswa int64
			for rows.Next() {
				rows.Scan(&count_siswa)
			}
			return count_siswa, nil
		}
	}
}

func (*UserModel) SiswaInAktif() (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		rows, err2 := db.Query("SELECT COUNT(level) as count_siswa FROM users WHERE level=0")
		if err2 != nil {
			return 0, err2
		} else {
			var count_siswa int64
			for rows.Next() {
				rows.Scan(&count_siswa)
			}
			return count_siswa, nil
		}
	}
}
