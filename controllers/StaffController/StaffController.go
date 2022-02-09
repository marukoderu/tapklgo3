package StaffController

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

func StaffDashboard(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	user_id := fmt.Sprint(userId)
	user_name := fmt.Sprint(userName)
	var tugasModel models.TugasModel
	var userModel models.UserModel
	pengumumans, _ := userModel.FindPengumumanStaffDash()
	pengumumans_gen, _ := userModel.FindPengumumanGeneral()
	siswa_aktif, _ := userModel.SiswaAktif()
	tugass, _ := tugasModel.StaffFindAllBelum(user_id)
	tugas_belum, _ := tugasModel.CountTugasStaffBelum(user_name)
	laporans, _ := tugasModel.StaffFindAllBelum(user_id)
	laporan_belum, _ := tugasModel.CountLaporanStaffBelum(user_name)
	data := map[string]interface{}{
		"pengumumans":     pengumumans,
		"pengumumans_gen": pengumumans_gen,
		"siswa_aktif":     siswa_aktif,
		"tugass":          tugass,
		"laporans":        laporans,
		"laporan_belum":   laporan_belum,
		"tugas_belum":     tugas_belum,
		"userId":          userId,
		"user_Name":       userName,
		"userEmail":       userEmail,
		"userPhoto":       userPhoto,
		"userLevel":       userLevel,
		"userDivisi":      userDivisi,
		"title":           "Staff Dashboard",
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/staff/dashboard.html",
		"templates/layouts/staff/navDashSt.html",
		"templates/layouts/staff/header/staffdash-header.html",
		"templates/layouts/staff/footer/staffdash-footer.html",
	))
	tmp.ExecuteTemplate(response, "dashboard", data)
}

func StaffPengumuman(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	username := fmt.Sprint(userName)
	title := "Pengumuman Management"
	var userModel models.UserModel
	pengumumans_self, _ := userModel.FindPengumumanStaff(username)
	pengumumans_gen, _ := userModel.FindPengumumanGeneral()
	divisis, _ := userModel.FindDivisi()
	data := map[string]interface{}{
		"pengumumans_self": pengumumans_self,
		"pengumumans_gen":  pengumumans_gen,
		"divisis":          divisis,
		"userId":           userId,
		"user_Name":        userName,
		"userEmail":        userEmail,
		"userPhoto":        userPhoto,
		"userLevel":        userLevel,
		"userDivisi":       userDivisi,
		"title":            title,
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/staff/pengumuman.html",
		"templates/layouts/staff/navDashSt.html",
		"templates/layouts/staff/header/staffdash-headerTugas.html",
		"templates/layouts/staff/footer/staffdash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "pengumuman", data)
}

func ActivePengumuman(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.ActivePengumuman(id)
	http.Redirect(response, request, "/dashboard/staff/pengumuman", http.StatusSeeOther)
}

func InactivePengumuman(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.InactivePengumuman(id)
	http.Redirect(response, request, "/dashboard/staff/pengumuman", http.StatusSeeOther)
}

func CreatePengumuman(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var pengumuman entities.Pengumuman
	pengumuman.Judul = request.Form.Get("judul")
	pengumuman.Isi = request.Form.Get("isi")
	pengumuman.Divisi = request.Form.Get("divisi")
	pengumuman.Name = request.Form.Get("name")
	pengumuman.UserID, _ = strconv.ParseUint(request.Form.Get("user_id"), 10, 64)

	var userModel models.UserModel
	userModel.CreatePengumuman(&pengumuman)
	http.Redirect(response, request, "/dashboard/staff/pengumuman", http.StatusSeeOther)
}

func UpdatePengumuman(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var pengumuman entities.Pengumuman
	pengumuman.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
	pengumuman.Judul = request.Form.Get("judul")
	pengumuman.Isi = request.Form.Get("isi")
	pengumuman.Divisi = request.Form.Get("divisi")

	var userModel models.UserModel
	userModel.UpdatePengumuman(pengumuman)
	http.Redirect(response, request, "/dashboard/staff/pengumuman", http.StatusSeeOther)
}

func DeletePengumuman(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.DeletePengumuman(id)
	http.Redirect(response, request, "/dashboard/staff/pengumuman", http.StatusSeeOther)
}

func StaffJadwal(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	title := "Jadwal Management"
	var userModel models.UserModel
	users, _ := userModel.FindByLevel()
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
		"templates/staff/jadwal.html",
		"templates/layouts/staff/navDashSt.html",
		"templates/layouts/staff/header/staffdash-headerTugas.html",
		"templates/layouts/staff/footer/staffdash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "jadwal", data)
}

func CreateJadwalProcess(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var jadwal entities.Jadwal
	jadwal.Day = request.Form.Get("day")
	jadwal.Name = request.Form.Get("name")

	var userModel models.UserModel
	userModel.CreateJadwal(&jadwal)
	http.Redirect(response, request, "/dashboard/staff/jadwal", http.StatusSeeOther)
}

func DeleteJadwal(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.DeleteJadwal(id)
	http.Redirect(response, request, "/dashboard/staff/jadwal", http.StatusSeeOther)
}

func StaffDivisi(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	title := "Divisi Management"
	var userModel models.UserModel
	divisis, _ := userModel.FindDivisi()
	data := map[string]interface{}{
		"divisis":    divisis,
		"userId":     userId,
		"user_Name":  userName,
		"userEmail":  userEmail,
		"userPhoto":  userPhoto,
		"userLevel":  userLevel,
		"userDivisi": userDivisi,
		"title":      title,
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/staff/divisi.html",
		"templates/layouts/staff/navDashSt.html",
		"templates/layouts/staff/header/staffdash-headerTugas.html",
		"templates/layouts/staff/footer/staffdash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "divisi", data)
}

func CreateDivisiProcess(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var divisi entities.Divisi
	divisi.Div = request.Form.Get("divisi")

	var userModel models.UserModel
	userModel.CreateDivisi(&divisi)
	http.Redirect(response, request, "/dashboard/staff/divisi", http.StatusSeeOther)
}

func UpdateDivisi(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var divisi entities.Divisi
	divisi.Id, _ = strconv.ParseInt(request.Form.Get("divisi_id"), 10, 64)
	divisi.Div = request.Form.Get("divisi")

	var userModel models.UserModel
	userModel.UpdateDivisi(divisi)
	http.Redirect(response, request, "/dashboard/staff/divisi", http.StatusSeeOther)
}

func DeleteDivisi(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.DeleteDivisi(id)
	http.Redirect(response, request, "/dashboard/staff/divisi", http.StatusSeeOther)
}

func Siswa(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	title := "Siswa Management"
	var userModel models.UserModel
	users, _ := userModel.FindByLevel()
	divisis, _ := userModel.FindDivisi()
	siswa_aktif, _ := userModel.SiswaAktif()
	siswa_inaktif, _ := userModel.SiswaInAktif()
	data := map[string]interface{}{
		"users":         users,
		"siswa_aktif":   siswa_aktif,
		"siswa_inaktif": siswa_inaktif,
		"divisis":       divisis,
		"userId":        userId,
		"user_Name":     userName,
		"userEmail":     userEmail,
		"userPhoto":     userPhoto,
		"userLevel":     userLevel,
		"userDivisi":    userDivisi,
		"title":         title,
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/staff/siswa.html",
		"templates/layouts/staff/navDashSt.html",
		"templates/layouts/staff/header/staffdash-headerTugas.html",
		"templates/layouts/staff/footer/staffdash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "siswa", data)
}

func CreateSiswaProcess(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var user entities.User
	user.Name = request.Form.Get("name")
	user.Email = request.Form.Get("email")
	user.Divisi = request.Form.Get("divisi")
	user.Password = request.Form.Get("password")

	var userModel models.UserModel
	userModel.CreateUser(&user)
	http.Redirect(response, request, "/dashboard/staff/siswa", http.StatusSeeOther)
}

func UpdateSiswa(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var user entities.User
	user.Id, _ = strconv.ParseInt(request.Form.Get("user_id"), 10, 64)
	user.Name = request.Form.Get("name")
	user.Email = request.Form.Get("email")
	user.Divisi = request.Form.Get("divisi")

	var userModel models.UserModel
	userModel.UpdateSiswa(user)
	http.Redirect(response, request, "/dashboard/staff/siswa", http.StatusSeeOther)
}

func ActiveSiswa(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.ActiveSiswa(id)
	http.Redirect(response, request, "/dashboard/staff/siswa", http.StatusSeeOther)
}

func InactiveSiswa(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.InactiveSiswa(id)
	http.Redirect(response, request, "/dashboard/staff/siswa", http.StatusSeeOther)
}

func DeleteSiswa(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.DeleteSiswa(id)
	http.Redirect(response, request, "/dashboard/staff/siswa", http.StatusSeeOther)
}

func Tugas(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	user_id := fmt.Sprint(userId)
	user_name := fmt.Sprint(userName)
	var tugasModel models.TugasModel
	var userModel models.UserModel
	tugass, _ := tugasModel.StaffFindAll(user_id)
	tugas_selesai, _ := tugasModel.CountTugasStaffSelesai(user_name)
	tugas_belum, _ := tugasModel.CountTugasStaffBelum(user_name)
	laporan_belum, _ := tugasModel.CountLaporanStaffBelum(user_name)
	divisis, _ := userModel.FindDivisi()
	data := map[string]interface{}{
		"divisis":       divisis,
		"tugass":        tugass,
		"tugas_selesai": tugas_selesai,
		"tugas_belum":   tugas_belum,
		"laporan_belum": laporan_belum,
		"userId":        userId,
		"user_Name":     userName,
		"userEmail":     userEmail,
		"userPhoto":     userPhoto,
		"userLevel":     userLevel,
		"userDivisi":    userDivisi,
		"title":         "Tugas Management",
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/staff/tugas.html",
		"templates/layouts/staff/navDashSt.html",
		"templates/layouts/staff/header/staffdash-headerTugas.html",
		"templates/layouts/staff/footer/staffdash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "tugas", data)
}

func StaffTugasConfirm(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var tugasModel models.TugasModel
	tugasModel.TugasConfirm(id)
	http.Redirect(response, request, "/dashboard/staff/tugas", http.StatusSeeOther)
}

func StaffTugasCancel(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var tugasModel models.TugasModel
	tugasModel.TugasCancel(id)
	http.Redirect(response, request, "/dashboard/staff/tugas", http.StatusSeeOther)
}

func StaffTugasSelesai(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var tugasModel models.TugasModel
	tugasModel.TugasSelesai(id)
	http.Redirect(response, request, "/dashboard/staff/tugas", http.StatusSeeOther)
}

func CreateTugasProcess(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var tugas entities.Tugas
	tugas.Judul = request.Form.Get("judul")
	tugas.Isi = request.Form.Get("isi")
	tugas.Divisi = request.Form.Get("divisi")
	id, _ := strconv.ParseInt(request.Form.Get("user_id"), 10, 64)
	tugas.UserID = uint64(id)
	tugas.UserName = request.Form.Get("user_name")
	level, _ := strconv.ParseInt(request.Form.Get("user_level"), 10, 64)
	tugas.UserLevel = uint64(level)

	var tugasModel models.TugasModel
	tugasModel.StaffCreateTugas(&tugas)
	http.Redirect(response, request, "/dashboard/staff/tugas", http.StatusSeeOther)
}

func StaffUpdateTugas(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var tugas entities.Tugas
	tugas.Id, _ = strconv.ParseInt(request.Form.Get("tugas_id"), 10, 64)
	tugas.Judul = request.Form.Get("judul")
	tugas.Isi = request.Form.Get("isi")
	tugas.Divisi = request.Form.Get("divisi")

	var tugasModel models.TugasModel
	tugasModel.StaffUpdate(tugas)
	http.Redirect(response, request, "/dashboard/staff/tugas", http.StatusSeeOther)
}

func StaffDeleteTugas(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var tugasModel models.TugasModel
	tugasModel.Delete(id)
	http.Redirect(response, request, "/dashboard/staff/tugas", http.StatusSeeOther)
}

func LaporanTugas(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	user_name := fmt.Sprint(userName)
	var tugasModel models.TugasModel
	laporans, _ := tugasModel.FindAllLaporan(user_name)
	laporan_selesai, _ := tugasModel.CountLaporanStaffSelesai(user_name)
	laporan_sedang, _ := tugasModel.CountLaporanStaffSedang(user_name)
	laporan_belum, _ := tugasModel.CountLaporanStaffBelum(user_name)
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
		"title":           "Laporan Tugas Management",
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/staff/laporan-tugas.html",
		"templates/layouts/staff/navDashSt.html",
		"templates/layouts/staff/header/staffdash-headerTugas.html",
		"templates/layouts/staff/footer/staffdash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "laporan-tugas", data)
}

func StaffLpTugasConfirm(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var tugasModel models.TugasModel
	tugasModel.LpTugasConfirm(id)
	http.Redirect(response, request, "/dashboard/staff/laporan-tugas", http.StatusSeeOther)
}

func StaffLpTugasFinish(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var tugasModel models.TugasModel
	tugasModel.LpTugasFinish(id)
	http.Redirect(response, request, "/dashboard/staff/laporan-tugas", http.StatusSeeOther)
}
