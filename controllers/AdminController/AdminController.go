package AdminController

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

func AdminDashboard(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]

	var userModel models.UserModel
	staff_aktif, _ := userModel.StaffAktif()
	siswa_aktif, _ := userModel.SiswaAktif()

	data := map[string]interface{}{
		"staff_aktif": staff_aktif,
		"siswa_aktif": siswa_aktif,
		"userName":    userName,
		"userId":      userId,
		"user_Name":   userName,
		"userEmail":   userEmail,
		"userPhoto":   userPhoto,
		"userLevel":   userLevel,
		"userDivisi":  userDivisi,
		"title":       "Admin Dashboard",
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/admin/dashA.html",
		"templates/layouts/admin/navDashA.html",
		"templates/layouts/admin/header/admindash-header.html",
		"templates/layouts/admin/footer/admindash-footer.html",
	))
	tmp.ExecuteTemplate(response, "dashA", data)
}

func AdminAccount(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	sess_name := session.Values["name"]
	title := "Account Management"
	data := map[string]interface{}{
		"sess_name": sess_name,
		"title":     title,
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/admin/accA.html",
		"templates/layouts/admin/navDashA.html",
		"templates/layouts/admin/header/admindash-headerTugas.html",
		"templates/layouts/admin/footer/admindash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "accA", data)
}

func Pengumuman(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	title := "Pengumuman Management"
	var userModel models.UserModel
	pengumumans, _ := userModel.FindPengumuman()
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
		"templates/admin/pengumuman.html",
		"templates/layouts/admin/navDashA.html",
		"templates/layouts/admin/header/admindash-headerTugas.html",
		"templates/layouts/admin/footer/admindash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "pengumuman", data)
}

func ActivePengumuman(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.ActivePengumuman(id)
	http.Redirect(response, request, "/dashboard/admin/pengumuman", http.StatusSeeOther)
}

func InactivePengumuman(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.InactivePengumuman(id)
	http.Redirect(response, request, "/dashboard/admin/pengumuman", http.StatusSeeOther)
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
	http.Redirect(response, request, "/dashboard/admin/pengumuman", http.StatusSeeOther)
}

func UpdatePengumuman(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var pengumuman entities.Pengumuman
	pengumuman.Id, _ = strconv.ParseInt(request.Form.Get("peng_id"), 10, 64)
	pengumuman.Judul = request.Form.Get("judul")
	pengumuman.Isi = request.Form.Get("isi")
	pengumuman.Divisi = request.Form.Get("divisi")

	var userModel models.UserModel
	userModel.UpdatePengumuman(pengumuman)
	http.Redirect(response, request, "/dashboard/admin/pengumuman", http.StatusSeeOther)
}

func DeletePengumuman(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.DeletePengumuman(id)
	http.Redirect(response, request, "/dashboard/admin/pengumuman", http.StatusSeeOther)
}

func Jadwal(response http.ResponseWriter, request *http.Request) {
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
		"userName":    userName,
		"userId":      userId,
		"user_Name":   userName,
		"userEmail":   userEmail,
		"userPhoto":   userPhoto,
		"userLevel":   userLevel,
		"userDivisi":  userDivisi,
		"title":       title,
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/admin/jadwal.html",
		"templates/layouts/admin/navDashA.html",
		"templates/layouts/admin/header/admindash-headerTugas.html",
		"templates/layouts/admin/footer/admindash-footerTugas.html",
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
	http.Redirect(response, request, "/dashboard/admin/jadwal", http.StatusSeeOther)
}

func DeleteJadwal(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.DeleteJadwal(id)
	http.Redirect(response, request, "/dashboard/admin/jadwal", http.StatusSeeOther)
}

func Staff(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	title := "Staff Management"

	var userModel models.UserModel
	users, _ := userModel.FindStaff()
	staff_aktif, _ := userModel.StaffAktif()
	staff_inaktif, _ := userModel.StaffInAktif()
	data := map[string]interface{}{
		"users":         users,
		"staff_aktif":   staff_aktif,
		"staff_inaktif": staff_inaktif,
		"userName":      userName,
		"userId":        userId,
		"user_Name":     userName,
		"userEmail":     userEmail,
		"userPhoto":     userPhoto,
		"userLevel":     userLevel,
		"userDivisi":    userDivisi,
		"title":         title,
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/admin/staff.html",
		"templates/layouts/admin/navDashA.html",
		"templates/layouts/admin/header/admindash-headerTugas.html",
		"templates/layouts/admin/footer/admindash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "staff", data)
}

func ActiveStaff(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.ActiveStaff(id)
	http.Redirect(response, request, "/dashboard/admin/staff", http.StatusSeeOther)
}

func InactiveStaff(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.InactiveStaff(id)
	http.Redirect(response, request, "/dashboard/admin/staff", http.StatusSeeOther)
}

func CreateStaffProcess(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var user entities.User
	user.Name = request.Form.Get("name")
	user.Email = request.Form.Get("email")
	user.Password = request.Form.Get("password")

	var userModel models.UserModel
	userModel.CreateUserStaff(&user)
	http.Redirect(response, request, "/dashboard/admin/staff", http.StatusSeeOther)
}

func UpdateStaff(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var user entities.User
	user.Id, _ = strconv.ParseInt(request.Form.Get("user_id"), 10, 64)
	user.Name = request.Form.Get("name")
	user.Email = request.Form.Get("email")

	var userModel models.UserModel
	userModel.UpdateSiswa(user)
	http.Redirect(response, request, "/dashboard/admin/staff", http.StatusSeeOther)
}

func DeleteStaff(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.DeleteSiswa(id)
	http.Redirect(response, request, "/dashboard/admin/staff", http.StatusSeeOther)
}

// ================================================================
// SISWA MANAGEMENT
// ================================================================
func Divisi(response http.ResponseWriter, request *http.Request) {
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
		"userName":   userName,
		"userId":     userId,
		"user_Name":  userName,
		"userEmail":  userEmail,
		"userPhoto":  userPhoto,
		"userLevel":  userLevel,
		"userDivisi": userDivisi,
		"title":      title,
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/admin/divisi.html",
		"templates/layouts/admin/navDashA.html",
		"templates/layouts/admin/header/admindash-headerTugas.html",
		"templates/layouts/admin/footer/admindash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "divisi", data)
}

func CreateDivisiProcess(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userID := session.Values["id"]
	request.ParseForm()
	var divisi entities.Divisi
	divisi.Div = request.Form.Get("divisi")

	fmt.Println("userID : ", userID)

	var userModel models.UserModel
	userModel.CreateDivisi(&divisi)
	http.Redirect(response, request, "/dashboard/admin/divisi", http.StatusSeeOther)
}

func UpdateDivisi(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var divisi entities.Divisi
	divisi.Id, _ = strconv.ParseInt(request.Form.Get("divisi_id"), 10, 64)
	divisi.Div = request.Form.Get("divisi")

	var userModel models.UserModel
	userModel.UpdateDivisi(divisi)
	http.Redirect(response, request, "/dashboard/admin/divisi", http.StatusSeeOther)
}

func DeleteDivisi(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.DeleteDivisi(id)
	http.Redirect(response, request, "/dashboard/admin/divisi", http.StatusSeeOther)
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
	siswa_aktif, _ := userModel.SiswaAktif()
	siswa_inaktif, _ := userModel.SiswaInAktif()
	data := map[string]interface{}{
		"users":         users,
		"siswa_aktif":   siswa_aktif,
		"siswa_inaktif": siswa_inaktif,
		"userName":      userName,
		"userId":        userId,
		"user_Name":     userName,
		"userEmail":     userEmail,
		"userPhoto":     userPhoto,
		"userLevel":     userLevel,
		"userDivisi":    userDivisi,
		"title":         title,
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/admin/siswa.html",
		"templates/layouts/admin/navDashA.html",
		"templates/layouts/admin/header/admindash-headerTugas.html",
		"templates/layouts/admin/footer/admindash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "siswa", data)
}

func ActiveSiswa(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.ActiveSiswa(id)
	http.Redirect(response, request, "/dashboard/admin/siswa", http.StatusSeeOther)
}

func InactiveSiswa(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.InactiveSiswa(id)
	http.Redirect(response, request, "/dashboard/admin/siswa", http.StatusSeeOther)
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
	http.Redirect(response, request, "/dashboard/admin/siswa", http.StatusSeeOther)
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
	http.Redirect(response, request, "/dashboard/admin/siswa", http.StatusSeeOther)
}

func DeleteSiswa(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var userModel models.UserModel
	userModel.DeleteSiswa(id)
	http.Redirect(response, request, "/dashboard/admin/siswa", http.StatusSeeOther)
}

// ================================================================
// TUGAS MANAGEMENT
// ================================================================
func Tugas(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	var tugasModel models.TugasModel
	tugass, _ := tugasModel.FindAll()
	// tugas_selesai, _ := tugasModel.CountTugasStaffSelesai(user_name)
	// tugas_belum, _ := tugasModel.CountTugasStaffBelum(user_name)
	// laporan_belum, _ := tugasModel.CountLaporanStaffBelum(user_name)
	data := map[string]interface{}{
		"tugass": tugass,
		// "tugas_selesai": tugas_selesai,
		// "tugas_belum":   tugas_belum,
		// "laporan_belum": laporan_belum,
		"userId":     userId,
		"user_Name":  userName,
		"userEmail":  userEmail,
		"userPhoto":  userPhoto,
		"userLevel":  userLevel,
		"userDivisi": userDivisi,
		"title":      "Tugas Management",
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/admin/tugas.html",
		"templates/layouts/admin/navDashA.html",
		"templates/layouts/admin/header/admindash-headerTugas.html",
		"templates/layouts/admin/footer/admindash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "tugas", data)
}

func AdmTugasConfirm(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var tugasModel models.TugasModel
	tugasModel.TugasConfirm(id)
	http.Redirect(response, request, "/dashboard/admin/tugas", http.StatusSeeOther)
}

func AdmTugasCancel(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var tugasModel models.TugasModel
	tugasModel.TugasCancel(id)
	http.Redirect(response, request, "/dashboard/admin/tugas", http.StatusSeeOther)
}

func AdmTugasSelesai(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var tugasModel models.TugasModel
	tugasModel.TugasSelesai(id)
	http.Redirect(response, request, "/dashboard/admin/tugas", http.StatusSeeOther)
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
	http.Redirect(response, request, "/dashboard/admin/tugas", http.StatusSeeOther)
}

func AdmUpdateTugas(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var tugas entities.Tugas
	tugas.Id, _ = strconv.ParseInt(request.Form.Get("tugas_id"), 10, 64)
	tugas.Judul = request.Form.Get("judul")
	tugas.Isi = request.Form.Get("isi")
	tugas.Divisi = request.Form.Get("divisi")

	var tugasModel models.TugasModel
	tugasModel.StaffUpdate(tugas)
	http.Redirect(response, request, "/dashboard/admin/tugas", http.StatusSeeOther)
}

func AdmDeleteTugas(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var tugasModel models.TugasModel
	tugasModel.Delete(id)
	http.Redirect(response, request, "/dashboard/admin/tugas", http.StatusSeeOther)
}

func LaporanTugas(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	sess_name := session.Values["name"]
	user_name := fmt.Sprint(sess_name)
	var tugasModel models.TugasModel
	laporans, _ := tugasModel.FindAllLaporan(user_name)
	data := map[string]interface{}{
		"laporans":  laporans,
		"sess_name": sess_name,
		"title":     "Laporan Tugas Management",
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/admin/laporan-tugas.html",
		"templates/layouts/admin/navDashA.html",
		"templates/layouts/admin/header/admindash-headerTugas.html",
		"templates/layouts/admin/footer/admindash-footerTugas.html",
	))
	tmp.ExecuteTemplate(response, "laporan-tugas", data)
}

func AdmLpTugasConfirm(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var tugasModel models.TugasModel
	tugasModel.LpTugasConfirm(id)
	http.Redirect(response, request, "/dashboard/admin/laporan-tugas", http.StatusSeeOther)
}

func AdmLpTugasFinish(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var tugasModel models.TugasModel
	tugasModel.LpTugasFinish(id)
	http.Redirect(response, request, "/dashboard/admin/laporan-tugas", http.StatusSeeOther)
}
