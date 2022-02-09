package models

import (
	"github.com/raisunee/tapklgo3/config"
	"github.com/raisunee/tapklgo3/entities"
)

type TugasModel struct {
}

// ================================================================
// ADMIN TUGAS MANAGEMENT
// ================================================================
func (*TugasModel) FindAll() ([]entities.Tugas, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, date, judul, isi, divisi, status, user_id, user_name FROM tugas")
		if err2 != nil {
			return nil, err2
		} else {
			var tugasss []entities.Tugas
			for rows.Next() {
				var tugas entities.Tugas
				rows.Scan(&tugas.Id, &tugas.Date, &tugas.Judul, &tugas.Isi, &tugas.Divisi, &tugas.Status, &tugas.UserID, &tugas.UserName)
				tugasss = append(tugasss, tugas)
			}
			return tugasss, nil
		}
	}
}

func (*TugasModel) CountTugasSedang(divisi string) (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		rows, err2 := db.Query("SELECT COUNT(status) as count_tugas FROM tugas WHERE divisi=? AND status=1", divisi)
		if err2 != nil {
			return 0, err2
		} else {
			var count_tugas int64
			for rows.Next() {
				rows.Scan(&count_tugas)
			}
			return count_tugas, nil
		}
	}
}

func (*TugasModel) CountTugasStaffSelesai(user_name string) (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		rows, err2 := db.Query("SELECT COUNT(status) as count_tugas FROM tugas WHERE user_name=? AND status=2", user_name)
		if err2 != nil {
			return 0, err2
		} else {
			var count_tugas int64
			for rows.Next() {
				rows.Scan(&count_tugas)
			}
			return count_tugas, nil
		}
	}
}

func (*TugasModel) CountTugasSelesai(divisi string) (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		rows, err2 := db.Query("SELECT COUNT(status) as count_tugas FROM tugas WHERE divisi=? AND status=2", divisi)
		if err2 != nil {
			return 0, err2
		} else {
			var count_tugas int64
			for rows.Next() {
				rows.Scan(&count_tugas)
			}
			return count_tugas, nil
		}
	}
}

func (*TugasModel) CountTugasStaffBelum(user_name string) (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		rows, err2 := db.Query("SELECT COUNT(status) as count_tugas FROM tugas WHERE user_name=? AND status=0", user_name)
		if err2 != nil {
			return 0, err2
		} else {
			var count_tugas int64
			for rows.Next() {
				rows.Scan(&count_tugas)
			}
			return count_tugas, nil
		}
	}
}

func (*TugasModel) CountTugasBelum(divisi string) (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		rows, err2 := db.Query("SELECT COUNT(status) as count_tugas FROM tugas WHERE divisi=? AND status=0", divisi)
		if err2 != nil {
			return 0, err2
		} else {
			var count_tugas int64
			for rows.Next() {
				rows.Scan(&count_tugas)
			}
			return count_tugas, nil
		}
	}
}

func (*TugasModel) CountLaporanStaffBelum(user_name string) (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		rows, err2 := db.Query("SELECT COUNT(status) as count_laporan FROM laporan WHERE pemberi=? AND status=0", user_name)
		if err2 != nil {
			return 0, err2
		} else {
			var count_laporan int64
			for rows.Next() {
				rows.Scan(&count_laporan)
			}
			return count_laporan, nil
		}
	}
}

func (*TugasModel) CountLaporanStaffSedang(user_name string) (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		rows, err2 := db.Query("SELECT COUNT(status) as count_laporan FROM laporan WHERE pemberi=? AND status=1", user_name)
		if err2 != nil {
			return 0, err2
		} else {
			var count_laporan int64
			for rows.Next() {
				rows.Scan(&count_laporan)
			}
			return count_laporan, nil
		}
	}
}

func (*TugasModel) CountLaporanStaffSelesai(user_name string) (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		rows, err2 := db.Query("SELECT COUNT(status) as count_laporan FROM laporan WHERE pemberi=? AND status=2", user_name)
		if err2 != nil {
			return 0, err2
		} else {
			var count_laporan int64
			for rows.Next() {
				rows.Scan(&count_laporan)
			}
			return count_laporan, nil
		}
	}
}

func (*TugasModel) CountLaporanSiswaBelum(user_name string) (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		rows, err2 := db.Query("SELECT COUNT(status) as count_laporan FROM laporan WHERE nama_siswa=? AND status=0", user_name)
		if err2 != nil {
			return 0, err2
		} else {
			var count_laporan int64
			for rows.Next() {
				rows.Scan(&count_laporan)
			}
			return count_laporan, nil
		}
	}
}

func (*TugasModel) CountLaporanSiswaSedang(user_name string) (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		rows, err2 := db.Query("SELECT COUNT(status) as count_laporan FROM laporan WHERE nama_siswa=? AND status=1", user_name)
		if err2 != nil {
			return 0, err2
		} else {
			var count_laporan int64
			for rows.Next() {
				rows.Scan(&count_laporan)
			}
			return count_laporan, nil
		}
	}
}

func (*TugasModel) CountLaporanSiswaSelesai(user_name string) (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	} else {
		rows, err2 := db.Query("SELECT COUNT(status) as count_laporan FROM laporan WHERE nama_siswa=? AND status=2", user_name)
		if err2 != nil {
			return 0, err2
		} else {
			var count_laporan int64
			for rows.Next() {
				rows.Scan(&count_laporan)
			}
			return count_laporan, nil
		}
	}
}

func (*TugasModel) StaffFindAll(user_id string) ([]entities.Tugas, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, date, judul, isi, divisi, status, user_id, user_name FROM tugas WHERE user_id=? ORDER BY status ASC", user_id)
		if err2 != nil {
			return nil, err2
		} else {
			var tugasss []entities.Tugas
			for rows.Next() {
				var tugas entities.Tugas
				rows.Scan(&tugas.Id, &tugas.Date, &tugas.Judul, &tugas.Isi, &tugas.Divisi, &tugas.Status, &tugas.UserID, &tugas.UserName)
				tugasss = append(tugasss, tugas)
			}
			return tugasss, nil
		}
	}
}

func (*TugasModel) StaffFindAllBelum(user_id string) ([]entities.Tugas, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, date, judul, isi, divisi, status, user_id, user_name FROM tugas WHERE user_id=? AND (status=0 OR status=1)", user_id)
		if err2 != nil {
			return nil, err2
		} else {
			var tugasss []entities.Tugas
			for rows.Next() {
				var tugas entities.Tugas
				rows.Scan(&tugas.Id, &tugas.Date, &tugas.Judul, &tugas.Isi, &tugas.Divisi, &tugas.Status, &tugas.UserID, &tugas.UserName)
				tugasss = append(tugasss, tugas)
			}
			return tugasss, nil
		}
	}
}

func (*TugasModel) FindAllLaporan(user_name string) ([]entities.Laporan, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, date, judul_lap, isi_lap, status, user_id, nama_siswa, tugas_name, tugas_id, divisi FROM laporan WHERE pemberi=?", user_name)
		if err2 != nil {
			return nil, err2
		} else {
			var laporans []entities.Laporan
			for rows.Next() {
				var laporan entities.Laporan
				rows.Scan(&laporan.Id, &laporan.Date, &laporan.Judul_lap, &laporan.Isi_lap, &laporan.Status, &laporan.UserID, &laporan.Nama_siswa, &laporan.Tugas_name, &laporan.TugasID, &laporan.Divisi)
				laporans = append(laporans, laporan)
			}
			return laporans, nil
		}
	}
}

func (*TugasModel) StaffFindLaporanBelum(user_id string) ([]entities.Tugas, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, judul_lap, isi_lap, status, user_id, tugas_id FROM laporan", user_id)
		if err2 != nil {
			return nil, err2
		} else {
			var tugasss []entities.Tugas
			for rows.Next() {
				var tugas entities.Tugas
				rows.Scan(&tugas.Id, &tugas.Date, &tugas.Judul, &tugas.Isi, &tugas.Divisi, &tugas.Status, &tugas.UserID, &tugas.UserName)
				tugasss = append(tugasss, tugas)
			}
			return tugasss, nil
		}
	}
}

func (*TugasModel) FindById(id int64) (entities.Tugas, error) {
	db, err := config.GetDB()
	if err != nil {
		return entities.Tugas{}, err
	} else {
		rows, err2 := db.Query("SELECT * FROM tugas WHERE id = ?", id)
		if err2 != nil {
			return entities.Tugas{}, err2
		} else {
			var tugas entities.Tugas
			for rows.Next() {
				rows.Scan(&tugas.Id, &tugas.Judul, &tugas.Isi)
			}
			return tugas, nil
		}
	}
}

func (*TugasModel) Update(tugas entities.Tugas) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE tugas SET judul = ?, isi = ? WHERE id = ?", tugas.Judul, tugas.Isi, tugas.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*TugasModel) StaffUpdate(tugas entities.Tugas) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE tugas SET judul = ?, isi = ?, divisi = ? WHERE id = ?", tugas.Judul, tugas.Isi, tugas.Divisi, tugas.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*TugasModel) Create(tugas *entities.Tugas) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("INSERT INTO tugas(judul, isi) VALUES(?,?)", tugas.Judul, tugas.Isi)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*TugasModel) TugasConfirm(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE tugas SET status=1 WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*TugasModel) TugasSelesai(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE tugas SET status=2 WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*TugasModel) TugasCancel(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE tugas SET status=0 WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*TugasModel) LpTugasConfirm(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE laporan SET status=1 WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*TugasModel) LpTugasFinish(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE laporan SET status=2 WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*TugasModel) StaffCreateTugas(tugas *entities.Tugas) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("INSERT INTO tugas(judul, isi, divisi, user_name, user_level, user_id) VALUES(?,?,?,?,?,?)", tugas.Judul, tugas.Isi, tugas.Divisi, tugas.UserName, tugas.UserLevel, tugas.UserID)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*TugasModel) SiswaConfirmLpTugas(laporan *entities.Laporan) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("INSERT INTO laporan(judul_lap, isi_lap, divisi, status, nama_siswa, pemberi, tugas_name, user_id, tugas_id) VALUES(?,?,?,0,?,?,?,?,?)", laporan.Judul_lap, laporan.Isi_lap, laporan.Divisi, laporan.Nama_siswa, laporan.Pemberi, laporan.Tugas_name, laporan.UserID, laporan.TugasID)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*TugasModel) SiswaConfirmTugas(tugas *entities.Tugas) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE tugas SET status=1 WHERE id=?", tugas.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*TugasModel) Delete(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("DELETE FROM tugas WHERE id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

// ================================================================
// SISWA TUGAS MANAGEMENT
// ================================================================
func (*TugasModel) SiswaFindAll(divisi string) ([]entities.Tugas, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, date, judul, isi, divisi, status, user_name FROM tugas WHERE divisi=? ORDER BY status ASC", divisi)
		if err2 != nil {
			return nil, err2
		} else {
			var tugass []entities.Tugas
			for rows.Next() {
				var tugas entities.Tugas
				rows.Scan(&tugas.Id, &tugas.Date, &tugas.Judul, &tugas.Isi, &tugas.Divisi, &tugas.Status, &tugas.UserName)
				tugass = append(tugass, tugas)
			}
			return tugass, nil
		}
	}
}

func (*TugasModel) SiswaFindAllBelum(divisi string) ([]entities.Tugas, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, date, judul, isi, divisi, status, user_name FROM tugas WHERE divisi=? AND status=0 ORDER BY id ASC", divisi)
		if err2 != nil {
			return nil, err2
		} else {
			var tugass []entities.Tugas
			for rows.Next() {
				var tugas entities.Tugas
				rows.Scan(&tugas.Id, &tugas.Date, &tugas.Judul, &tugas.Isi, &tugas.Divisi, &tugas.Status, &tugas.UserName)
				tugass = append(tugass, tugas)
			}
			return tugass, nil
		}
	}
}

func (*TugasModel) SiswaFindLaporan(user_id string) ([]entities.Laporan, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, date, judul_lap, isi_lap, status, user_id, tugas_name, pemberi FROM laporan WHERE user_id=?", user_id)
		if err2 != nil {
			return nil, err2
		} else {
			var laporans []entities.Laporan
			for rows.Next() {
				var laporan entities.Laporan
				rows.Scan(&laporan.Id, &laporan.Date, &laporan.Judul_lap, &laporan.Isi_lap, &laporan.Status, &laporan.UserID, &laporan.Tugas_name, &laporan.Pemberi)
				laporans = append(laporans, laporan)
			}
			return laporans, nil
		}
	}
}

func (*TugasModel) SiswaFindLaporanBelum(user_id string) ([]entities.Laporan, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, date, judul_lap, isi_lap, status, user_id, tugas_name, pemberi FROM laporan WHERE user_id=? AND status=0 ORDER BY id ASC", user_id)
		if err2 != nil {
			return nil, err2
		} else {
			var laporans []entities.Laporan
			for rows.Next() {
				var laporan entities.Laporan
				rows.Scan(&laporan.Id, &laporan.Date, &laporan.Judul_lap, &laporan.Isi_lap, &laporan.Status, &laporan.UserID, &laporan.Tugas_name, &laporan.Pemberi)
				laporans = append(laporans, laporan)
			}
			return laporans, nil
		}
	}
}

func (*TugasModel) UpdateLaporan(laporan entities.Laporan) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE laporan SET judul_lap = ?, isi_lap = ? WHERE id = ?", laporan.Judul_lap, laporan.Isi_lap, laporan.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}
