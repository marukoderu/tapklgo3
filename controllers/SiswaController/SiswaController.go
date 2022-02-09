package SiswaController

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/sessions"
	"github.com/raisunee/tapklgo3/entities"
	"github.com/raisunee/tapklgo3/models"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func SiswaDashboard(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	user_id := fmt.Sprint(userId)
	user_name := fmt.Sprint(userName)
	divisi := fmt.Sprint(userDivisi)
	var tugasModel models.TugasModel
	var userModel models.UserModel
	pengumumans, _ := userModel.FindPengumumanSiswaDash(divisi)
	pengumumans_gen, _ := userModel.FindPengumumanGeneral()
	laporans, _ := tugasModel.SiswaFindLaporanBelum(user_id)
	laporan_belum, _ := tugasModel.CountLaporanSiswaBelum(user_name)
	tugass, _ := tugasModel.SiswaFindAllBelum(divisi)
	tugas_selesai, _ := tugasModel.CountTugasStaffSelesai(user_name)
	tugas_belum, _ := tugasModel.CountTugasBelum(divisi)
	data := map[string]interface{}{
		"pengumumans":     pengumumans,
		"pengumumans_gen": pengumumans_gen,
		"laporans":        laporans,
		"laporan_belum":   laporan_belum,
		"tugass":          tugass,
		"tugas_selesai":   tugas_selesai,
		"tugas_belum":     tugas_belum,
		"userId":          userId,
		"user_Name":       userName,
		"userEmail":       userEmail,
		"userPhoto":       userPhoto,
		"userLevel":       userLevel,
		"userDivisi":      userDivisi,
		"title":           "Siswa Dashboard",
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/siswa/dashboard.html",
		"templates/layouts/siswa/navDashS.html",
		"templates/layouts/siswa/header/siswadash-header.html",
		"templates/layouts/siswa/footer/siswadash-footer.html",
	))
	tmp.ExecuteTemplate(response, "dashboard", data)
}

// ================================================================
// TUGAS MANAGEMENT
// ================================================================
func TugasSiswa(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	divisi := fmt.Sprint(userDivisi)
	fmt.Println(divisi)
	var tugasModel models.TugasModel
	tugass, _ := tugasModel.SiswaFindAll(divisi)
	tugas_selesai, _ := tugasModel.CountTugasSelesai(divisi)
	tugas_sedang, _ := tugasModel.CountTugasSedang(divisi)
	tugas_belum, _ := tugasModel.CountTugasBelum(divisi)
	data := map[string]interface{}{
		"tugass":        tugass,
		"tugas_selesai": tugas_selesai,
		"tugas_sedang":  tugas_sedang,
		"tugas_belum":   tugas_belum,
		"userId":        userId,
		"user_Name":     userName,
		"userEmail":     userEmail,
		"userPhoto":     userPhoto,
		"userLevel":     userLevel,
		"userDivisi":    userDivisi,
		"title":         "Tugas Siswa",
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/siswa/tugas.html",
		"templates/layouts/siswa/navDashS.html",
		"templates/layouts/siswa/header/siswadash-headerTugas.html",
		"templates/layouts/siswa/footer/siswadash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "tugas", data)
}

func ConfirmTugas(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	user_id := fmt.Sprint(userId)
	user_name := fmt.Sprint(userName)
	request.ParseForm()
	var tugas entities.Tugas
	id_for_tugas, _ := strconv.ParseInt(request.Form.Get("tugas_id"), 10, 64)
	tugas.Id = int64(id_for_tugas)
	var laporan entities.Laporan
	laporan.Judul_lap = request.Form.Get("judul_lap")
	laporan.Isi_lap = request.Form.Get("isi_lap")
	laporan.Divisi = request.Form.Get("divisi")
	id_user, _ := strconv.ParseInt(user_id, 10, 64)
	laporan.UserID = uint64(id_user)
	id_tugas, _ := strconv.ParseInt(request.Form.Get("tugas_id"), 10, 64)
	laporan.TugasID = uint64(id_tugas)
	laporan.Nama_siswa = user_name
	laporan.Pemberi = request.Form.Get("pemberi")
	laporan.Tugas_name = request.Form.Get("tugas_name")

	var tugasModel models.TugasModel
	tugasModel.SiswaConfirmLpTugas(&laporan)
	tugasModel.SiswaConfirmTugas(&tugas)
	http.Redirect(response, request, "/dashboard/siswa/laporan-tugas", http.StatusSeeOther)
}

func LaporanTugasSiswa(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	user_name := fmt.Sprint(userName)
	user_id := fmt.Sprint(userId)
	var tugasModel models.TugasModel
	laporans, _ := tugasModel.SiswaFindLaporan(user_id)
	laporan_selesai, _ := tugasModel.CountLaporanSiswaSelesai(user_name)
	laporan_sedang, _ := tugasModel.CountLaporanSiswaSedang(user_name)
	laporan_belum, _ := tugasModel.CountLaporanSiswaBelum(user_name)
	data := map[string]interface{}{
		"laporans":        laporans,
		"laporan_selesai": laporan_selesai,
		"laporan_sedang":  laporan_sedang,
		"laporan_belum":   laporan_belum,
		"userId":          userId,
		"user_Name":       userName,
		"userEmail":       userEmail,
		"userPhoto":       userPhoto,
		"userLevel":       userLevel,
		"userDivisi":      userDivisi,
		"title":           "Laporan Tugas Siswa",
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/siswa/laporan-tugas.html",
		"templates/layouts/siswa/navDashS.html",
		"templates/layouts/siswa/header/siswadash-headerTugas.html",
		"templates/layouts/siswa/footer/siswadash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "laporan-tugas", data)
}

func EditLaporan(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var laporan entities.Laporan
	laporan.Id, _ = strconv.ParseInt(request.Form.Get("laporan_id"), 10, 64)
	laporan.Judul_lap = request.Form.Get("judul_lap")
	laporan.Isi_lap = request.Form.Get("isi_lap")

	var tugasModel models.TugasModel
	tugasModel.UpdateLaporan(laporan)
	http.Redirect(response, request, "/dashboard/siswa/laporan-tugas", http.StatusSeeOther)
}

func DeleteLaporan(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.DeleteLaporan(id)
	http.Redirect(response, request, "/dashboard/siswa/laporan-tugas", http.StatusSeeOther)
}

// ================================================================
// JADWAL MANAGEMENT
// ================================================================
func JadwalSiswa(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	email := fmt.Sprint(userEmail)
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	title := "Jadwal Siswa"
	var userModel models.UserModel
	users, _ := userModel.FindAccAdmin(email)
	divisis, _ := userModel.FindDivisi()
	// JADWAL BY DAY
	jadwals_mon, _ := userModel.FindJadwalMon()
	jadwals_tue, _ := userModel.FindJadwalTue()
	jadwals_wed, _ := userModel.FindJadwalWed()
	jadwals_thu, _ := userModel.FindJadwalThu()
	jadwals_fri, _ := userModel.FindJadwalFri()
	data := map[string]interface{}{
		"users":       users,
		"divisis":     divisis,
		"jadwals_mon": jadwals_mon,
		"jadwals_tue": jadwals_tue,
		"jadwals_wed": jadwals_wed,
		"jadwals_thu": jadwals_thu,
		"jadwals_fri": jadwals_fri,
		"userId":      userId,
		"user_Name":   userName,
		"userEmail":   userEmail,
		"userPhoto":   userPhoto,
		"userLevel":   userLevel,
		"userDivisi":  userDivisi,
		"title":       title,
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/siswa/jadwal.html",
		"templates/layouts/siswa/navDashS.html",
		"templates/layouts/siswa/header/siswadash-headerTugas.html",
		"templates/layouts/siswa/footer/siswadash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "jadwal", data)
}

// ================================================================
// PENGUMUMAN MANAGEMENT
// ================================================================
func SiswaPengumuman(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	divisi := fmt.Sprint(userDivisi)
	title := "Pengumuman Siswa"
	var userModel models.UserModel
	pengumumans, _ := userModel.FindPengumumanSiswa(divisi)
	data := map[string]interface{}{
		"pengumumans": pengumumans,
		"userId":      userId,
		"user_Name":   userName,
		"userEmail":   userEmail,
		"userPhoto":   userPhoto,
		"userLevel":   userLevel,
		"userDivisi":  userDivisi,
		"title":       title,
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/siswa/pengumuman.html",
		"templates/layouts/siswa/navDashS.html",
		"templates/layouts/siswa/header/siswadash-headerTugas.html",
		"templates/layouts/siswa/footer/siswadash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "pengumuman", data)
}
