package main

import (
	"net/http"

	"github.com/raisunee/tapklgo3/controllers/AccountController"
	"github.com/raisunee/tapklgo3/controllers/AdminController"
	"github.com/raisunee/tapklgo3/controllers/AuthController"
	"github.com/raisunee/tapklgo3/controllers/SiswaController"
	"github.com/raisunee/tapklgo3/controllers/StaffController"
	"github.com/raisunee/tapklgo3/middlewares/AuthMiddleware"
)

func main() {

	// fs := http.FileServer(http.Dir("assets/valex/HTML/html/assets/css"))
	// http.Handle("/assets/valex/HTML/html/assets/css/", http.StripPrefix("/assets/valex/HTML/html/assets/css/", fs))
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// AUTH
	http.HandleFunc("/login", AuthController.Login)
	http.HandleFunc("/login/", AuthController.Login)
	http.HandleFunc("/login/process", AuthController.LoginProcess)
	http.HandleFunc("/register", AuthController.Register)
	http.HandleFunc("/register/process", AuthController.RegisterProcess)
	http.HandleFunc("/logout", AuthController.Logout)

	// DASHBOARD
	// ---- LEVEL ADMIN
	http.HandleFunc("/dashboard/admin", AuthMiddleware.AuthMiddleware(AdminController.AdminDashboard))
	http.HandleFunc("/dashboard/admin/account", AuthMiddleware.AuthMiddleware(AccountController.AdminAccount))
	http.HandleFunc("/dashboard/admin/account/update", AuthMiddleware.AuthMiddleware(AccountController.UpdateAccountAdm))
	http.HandleFunc("/dashboard/admin/account/photo/update", AuthMiddleware.AuthMiddleware(AccountController.UpdatePhotoAdm))
	http.HandleFunc("/dashboard/admin/account/password/update", AuthMiddleware.AuthMiddleware(AccountController.UpdatePasswordAdm))
	http.HandleFunc("/dashboard/admin/pengumuman", AuthMiddleware.AuthMiddleware(AdminController.Pengumuman))
	http.HandleFunc("/dashboard/admin/pengumuman/active", AuthMiddleware.AuthMiddleware(AdminController.ActivePengumuman))
	http.HandleFunc("/dashboard/admin/pengumuman/inactive", AuthMiddleware.AuthMiddleware(AdminController.InactivePengumuman))
	http.HandleFunc("/dashboard/admin/pengumuman/create", AuthMiddleware.AuthMiddleware(AdminController.CreatePengumuman))
	http.HandleFunc("/dashboard/admin/pengumuman/update", AuthMiddleware.AuthMiddleware(AdminController.UpdatePengumuman))
	http.HandleFunc("/dashboard/admin/pengumuman/delete", AuthMiddleware.AuthMiddleware(AdminController.DeletePengumuman))
	http.HandleFunc("/dashboard/admin/jadwal", AuthMiddleware.AuthMiddleware(AdminController.Jadwal))
	http.HandleFunc("/dashboard/admin/jadwal/create", AuthMiddleware.AuthMiddleware(AdminController.CreateJadwalProcess))
	http.HandleFunc("/dashboard/admin/jadwal/delete", AuthMiddleware.AuthMiddleware(AdminController.DeleteJadwal))
	// ------- STAFF MANAGEMENT
	http.HandleFunc("/dashboard/admin/staff", AuthMiddleware.AuthMiddleware(AdminController.Staff))
	http.HandleFunc("/dashboard/admin/staff/active", AuthMiddleware.AuthMiddleware(AdminController.ActiveStaff))
	http.HandleFunc("/dashboard/admin/staff/inactive", AuthMiddleware.AuthMiddleware(AdminController.InactiveStaff))
	http.HandleFunc("/dashboard/admin/staff/create", AuthMiddleware.AuthMiddleware(AdminController.CreateStaffProcess))
	http.HandleFunc("/dashboard/admin/staff/update", AuthMiddleware.AuthMiddleware(AdminController.UpdateStaff))
	http.HandleFunc("/dashboard/admin/staff/delete", AuthMiddleware.AuthMiddleware(AdminController.DeleteStaff))
	// ------- SISWA MANAGEMENT
	http.HandleFunc("/dashboard/admin/divisi", AuthMiddleware.AuthMiddleware(AdminController.Divisi))
	http.HandleFunc("/dashboard/admin/divisi/create/process", AuthMiddleware.AuthMiddleware(AdminController.CreateDivisiProcess))
	http.HandleFunc("/dashboard/admin/divisi/update", AuthMiddleware.AuthMiddleware(AdminController.UpdateDivisi))
	http.HandleFunc("/dashboard/admin/divisi/delete", AuthMiddleware.AuthMiddleware(AdminController.DeleteDivisi))
	http.HandleFunc("/dashboard/admin/siswa", AuthMiddleware.AuthMiddleware(AdminController.Siswa))
	http.HandleFunc("/dashboard/admin/siswa/active", AuthMiddleware.AuthMiddleware(AdminController.ActiveSiswa))
	http.HandleFunc("/dashboard/admin/siswa/inactive", AuthMiddleware.AuthMiddleware(AdminController.InactiveSiswa))
	http.HandleFunc("/dashboard/admin/siswa/create", AuthMiddleware.AuthMiddleware(AdminController.CreateSiswaProcess))
	http.HandleFunc("/dashboard/admin/siswa/update", AuthMiddleware.AuthMiddleware(AdminController.UpdateSiswa))
	http.HandleFunc("/dashboard/admin/siswa/delete", AuthMiddleware.AuthMiddleware(AdminController.DeleteSiswa))
	// ------- TUGAS MANAGEMENT
	http.HandleFunc("/dashboard/admin/tugas", AuthMiddleware.AuthMiddleware(AdminController.Tugas))
	http.HandleFunc("/dashboard/admin/tugas/confirm", AuthMiddleware.AuthMiddleware(AdminController.AdmTugasConfirm))
	http.HandleFunc("/dashboard/admin/tugas/cancel", AuthMiddleware.AuthMiddleware(AdminController.AdmTugasCancel))
	http.HandleFunc("/dashboard/admin/tugas/finish", AuthMiddleware.AuthMiddleware(AdminController.AdmTugasSelesai))
	http.HandleFunc("/dashboard/admin/tugas/create", AuthMiddleware.AuthMiddleware(AdminController.CreateTugasProcess))
	http.HandleFunc("/dashboard/admin/tugas/update", AuthMiddleware.AuthMiddleware(AdminController.AdmUpdateTugas))
	http.HandleFunc("/dashboard/admin/tugas/delete", AuthMiddleware.AuthMiddleware(AdminController.AdmDeleteTugas))
	http.HandleFunc("/dashboard/admin/laporan-tugas", AuthMiddleware.AuthMiddleware(AdminController.LaporanTugas))
	http.HandleFunc("/dashboard/admin/laporan-tugas/confirm", AuthMiddleware.AuthMiddleware(AdminController.AdmLpTugasConfirm))
	http.HandleFunc("/dashboard/admin/laporan-tugas/finish", AuthMiddleware.AuthMiddleware(AdminController.AdmLpTugasFinish))
	http.HandleFunc("/dashboard/admin/laporan-tugas/create", AuthMiddleware.AuthMiddleware(StaffController.CreateTugasProcess))

	// ---- LEVEL STAFF
	http.HandleFunc("/dashboard/staff", AuthMiddleware.AuthMiddleware(StaffController.StaffDashboard))
	http.HandleFunc("/dashboard/staff/account", AuthMiddleware.AuthMiddleware(AccountController.StaffAccount))
	http.HandleFunc("/dashboard/staff/account/update", AuthMiddleware.AuthMiddleware(AccountController.UpdateAccountStaff))
	http.HandleFunc("/dashboard/staff/account/photo/update", AuthMiddleware.AuthMiddleware(AccountController.UpdatePhotoStaff))
	http.HandleFunc("/dashboard/staff/account/password/update", AuthMiddleware.AuthMiddleware(AccountController.UpdatePasswordStaff))
	http.HandleFunc("/dashboard/staff/pengumuman", AuthMiddleware.AuthMiddleware(StaffController.StaffPengumuman))
	http.HandleFunc("/dashboard/staff/pengumuman/active", AuthMiddleware.AuthMiddleware(StaffController.ActivePengumuman))
	http.HandleFunc("/dashboard/staff/pengumuman/inactive", AuthMiddleware.AuthMiddleware(StaffController.InactivePengumuman))
	http.HandleFunc("/dashboard/staff/pengumuman/create", AuthMiddleware.AuthMiddleware(StaffController.CreatePengumuman))
	http.HandleFunc("/dashboard/staff/pengumuman/update", AuthMiddleware.AuthMiddleware(StaffController.UpdatePengumuman))
	http.HandleFunc("/dashboard/staff/pengumuman/delete", AuthMiddleware.AuthMiddleware(StaffController.DeletePengumuman))
	http.HandleFunc("/dashboard/staff/jadwal", AuthMiddleware.AuthMiddleware(StaffController.StaffJadwal))
	http.HandleFunc("/dashboard/staff/jadwal/create", AuthMiddleware.AuthMiddleware(StaffController.CreateJadwalProcess))
	http.HandleFunc("/dashboard/staff/jadwal/delete", AuthMiddleware.AuthMiddleware(StaffController.DeleteJadwal))
	http.HandleFunc("/dashboard/staff/divisi", AuthMiddleware.AuthMiddleware(StaffController.StaffDivisi))
	http.HandleFunc("/dashboard/staff/divisi/create", AuthMiddleware.AuthMiddleware(StaffController.CreateDivisiProcess))
	http.HandleFunc("/dashboard/staff/divisi/update", AuthMiddleware.AuthMiddleware(StaffController.UpdateDivisi))
	http.HandleFunc("/dashboard/staff/divisi/delete", AuthMiddleware.AuthMiddleware(StaffController.DeleteDivisi))
	http.HandleFunc("/dashboard/staff/siswa", AuthMiddleware.AuthMiddleware(StaffController.Siswa))
	http.HandleFunc("/dashboard/staff/siswa/active", AuthMiddleware.AuthMiddleware(StaffController.ActiveSiswa))
	http.HandleFunc("/dashboard/staff/siswa/inactive", AuthMiddleware.AuthMiddleware(StaffController.InactiveSiswa))
	http.HandleFunc("/dashboard/staff/siswa/create", AuthMiddleware.AuthMiddleware(StaffController.CreateSiswaProcess))
	http.HandleFunc("/dashboard/staff/siswa/update", AuthMiddleware.AuthMiddleware(StaffController.UpdateSiswa))
	http.HandleFunc("/dashboard/staff/siswa/delete", AuthMiddleware.AuthMiddleware(StaffController.DeleteSiswa))
	http.HandleFunc("/dashboard/staff/tugas", AuthMiddleware.AuthMiddleware(StaffController.Tugas))
	http.HandleFunc("/dashboard/staff/tugas/confirm", AuthMiddleware.AuthMiddleware(StaffController.StaffTugasConfirm))
	http.HandleFunc("/dashboard/staff/tugas/cancel", AuthMiddleware.AuthMiddleware(StaffController.StaffTugasCancel))
	http.HandleFunc("/dashboard/staff/tugas/finish", AuthMiddleware.AuthMiddleware(StaffController.StaffTugasSelesai))
	http.HandleFunc("/dashboard/staff/tugas/create", AuthMiddleware.AuthMiddleware(StaffController.CreateTugasProcess))
	http.HandleFunc("/dashboard/staff/tugas/update", AuthMiddleware.AuthMiddleware(StaffController.StaffUpdateTugas))
	http.HandleFunc("/dashboard/staff/tugas/delete", AuthMiddleware.AuthMiddleware(StaffController.StaffDeleteTugas))
	http.HandleFunc("/dashboard/staff/laporan-tugas", AuthMiddleware.AuthMiddleware(StaffController.LaporanTugas))
	http.HandleFunc("/dashboard/staff/laporan-tugas/confirm", AuthMiddleware.AuthMiddleware(StaffController.StaffLpTugasConfirm))
	http.HandleFunc("/dashboard/staff/laporan-tugas/finish", AuthMiddleware.AuthMiddleware(StaffController.StaffLpTugasFinish))
	http.HandleFunc("/dashboard/staff/laporan-tugas/create", AuthMiddleware.AuthMiddleware(StaffController.CreateTugasProcess))

	// ---- LEVEL USER / SISWA
	http.HandleFunc("/dashboard/siswa", AuthMiddleware.AuthMiddleware(SiswaController.SiswaDashboard))
	http.HandleFunc("/dashboard/siswa/account", AuthMiddleware.AuthMiddleware(AccountController.SiswaAccount))
	http.HandleFunc("/dashboard/siswa/account/photo/update", AuthMiddleware.AuthMiddleware(AccountController.UpdatePhotoSiswa))
	http.HandleFunc("/dashboard/siswa/account/update", AuthMiddleware.AuthMiddleware(AccountController.UpdateAccountSiswa))
	http.HandleFunc("/dashboard/siswa/account/password/update", AuthMiddleware.AuthMiddleware(AccountController.UpdatePasswordSiswa))
	http.HandleFunc("/dashboard/siswa/tugas", AuthMiddleware.AuthMiddleware(SiswaController.TugasSiswa))
	http.HandleFunc("/dashboard/siswa/tugas/confirm", AuthMiddleware.AuthMiddleware(SiswaController.ConfirmTugas))
	http.HandleFunc("/dashboard/siswa/laporan-tugas", AuthMiddleware.AuthMiddleware(SiswaController.LaporanTugasSiswa))
	http.HandleFunc("/dashboard/siswa/laporan-tugas/edit", AuthMiddleware.AuthMiddleware(SiswaController.EditLaporan))
	http.HandleFunc("/dashboard/siswa/laporan-tugas/delete", AuthMiddleware.AuthMiddleware(SiswaController.DeleteLaporan))
	http.HandleFunc("/dashboard/siswa/jadwal", AuthMiddleware.AuthMiddleware(SiswaController.JadwalSiswa))
	http.HandleFunc("/dashboard/siswa/pengumuman", AuthMiddleware.AuthMiddleware(SiswaController.SiswaPengumuman))

	http.ListenAndServe(":8080", nil)
}
