package AccountController

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/gorilla/sessions"
	"github.com/raisunee/tapklgo3/entities"
	"github.com/raisunee/tapklgo3/models"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("mysession"))

// ================================================================
// ADMIN ACCOUNT
// ================================================================
func AdminAccount(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	user_id := fmt.Sprint(userId)
	var userModel models.UserModel
	users, _ := userModel.FindAccAdmin(user_id)
	title := "Account Management"
	data := map[string]interface{}{
		"users":      users,
		"userId":     userId,
		"user_Name":  userName,
		"userEmail":  userEmail,
		"userPhoto":  userPhoto,
		"userLevel":  userLevel,
		"userDivisi": userDivisi,
		"title":      title,
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/admin/accA.html",
		"templates/layouts/admin/navDashA.html",
		"templates/layouts/admin/header/admindash-headerAcc.html",
		"templates/layouts/admin/footer/admindash-footerAcc.html",
	))
	tmp.ExecuteTemplate(response, "accA", data)
}

// ================================================================
// STAFF ACCOUNT
// ================================================================
func StaffAccount(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	userId := session.Values["id"]
	userName := session.Values["name"]
	userEmail := session.Values["email"]
	userPhoto := session.Values["photo"]
	userLevel := session.Values["level"]
	userDivisi := session.Values["divisi"]
	user_id := fmt.Sprint(userId)
	var userModel models.UserModel
	users, _ := userModel.FindAccAdmin(user_id)
	title := "Account Management"
	data := map[string]interface{}{
		"users":      users,
		"userId":     userId,
		"user_Name":  userName,
		"userEmail":  userEmail,
		"userPhoto":  userPhoto,
		"userLevel":  userLevel,
		"userDivisi": userDivisi,
		"title":      title,
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/staff/accS.html",
		"templates/layouts/staff/navDashSt.html",
		"templates/layouts/staff/header/staffdash-headerAcc.html",
		"templates/layouts/staff/footer/staffdash-footerAcc.html",
	))
	tmp.ExecuteTemplate(response, "accS", data)
}

// ================================================================
// SISWA ACCOUNT
// ================================================================
func SiswaAccount(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	sess_id := session.Values["id"]
	sess_name := session.Values["name"]
	sess_email := session.Values["email"]
	sess_photo := session.Values["photo"]
	user_id := fmt.Sprint(sess_id)
	var userModel models.UserModel
	users, _ := userModel.FindAccAdmin(user_id)
	title := "Account Management"
	data := map[string]interface{}{
		"users":      users,
		"sess_id":    sess_id,
		"sess_name":  sess_name,
		"sess_email": sess_email,
		"sess_photo": sess_photo,
		"title":      title,
	}
	var tmp = template.Must(template.ParseFiles(
		"templates/siswa/accS.html",
		"templates/layouts/siswa/navDashS.html",
		"templates/layouts/siswa/header/siswadash-headerAcc.html",
		"templates/layouts/siswa/footer/siswadash-footerAcc.html",
	))
	tmp.ExecuteTemplate(response, "accS", data)
}

func UpdateAccountAdm(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)
	var user entities.User
	user.Id, _ = strconv.ParseInt(request.Form.Get("user_id"), 10, 64)
	user.Name = request.Form.Get("name")
	user.Email = request.Form.Get("email")

	var userModel models.UserModel
	userModel.UpdateAccountAdm(user)
	http.Redirect(response, request, "/dashboard/admin/account", http.StatusSeeOther)
}

func UpdatePhotoAdm(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)
	var user entities.User
	user.Id, _ = strconv.ParseInt(request.Form.Get("user_id"), 10, 64)
	user.Photo = request.Form.Get("photo")
	file, handler, err := request.FormFile("photo")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	dst, err := os.Create("C:/xampp/htdocs/tapklgo3/assets/img/" + handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	var userModel models.UserModel
	userModel.UpdatePhotoStaff(user)
	http.Redirect(response, request, "/dashboard/admin/account", http.StatusSeeOther)
}

func UpdateAccountStaff(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)
	var user entities.User
	user.Id, _ = strconv.ParseInt(request.Form.Get("user_id"), 10, 64)
	user.Name = request.Form.Get("name")
	user.Email = request.Form.Get("email")

	var userModel models.UserModel
	userModel.UpdateAccountStaff(user)
	http.Redirect(response, request, "/dashboard/staff/account", http.StatusSeeOther)
}

func UpdatePhotoStaff(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)
	var user entities.User
	user.Id, _ = strconv.ParseInt(request.Form.Get("user_id"), 10, 64)
	user.Photo = request.Form.Get("photo")
	file, handler, err := request.FormFile("photo")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	dst, err := os.Create("C:/xampp/htdocs/tapklgo3/assets/img/" + handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	var userModel models.UserModel
	userModel.UpdatePhotoStaff(user)
	http.Redirect(response, request, "/dashboard/staff/account", http.StatusSeeOther)
}

func UpdatePhotoSiswa(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)
	var user entities.User
	user.Id, _ = strconv.ParseInt(request.Form.Get("user_id"), 10, 64)
	user.Photo = request.Form.Get("photo")
	file, handler, err := request.FormFile("photo")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	dst, err := os.Create("C:/xampp/htdocs/tapklgo3/assets/img/" + handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	var userModel models.UserModel
	userModel.UpdatePhotoSis(user)
	http.Redirect(response, request, "/dashboard/siswa/account", http.StatusSeeOther)
}

func UpdateAccountSiswa(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var user entities.User
	user.Id, _ = strconv.ParseInt(request.Form.Get("user_id"), 10, 64)
	user.Name = request.Form.Get("name")
	user.Email = request.Form.Get("email")

	var userModel models.UserModel
	userModel.UpdateAccountSis(user)
	http.Redirect(response, request, "/dashboard/siswa/account", http.StatusSeeOther)
}

func UpdatePasswordAdm(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	session, _ := store.Get(request, "mysession")
	sess_email := session.Values["email"]
	email := fmt.Sprint(sess_email)
	user_id, _ := strconv.ParseInt(request.Form.Get("user_id"), 10, 64)
	password_old := request.FormValue("password_old")
	password_new := request.FormValue("password_new")
	password_new_c := request.FormValue("password_confirm")
	var userModel models.UserModel
	users, _ := userModel.QueryUser(email)

	errf := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password_old))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("ERROR!")
	} else {
		if sess_email == users.Email {
			var user entities.User
			user.Password = password_new
			if user.Password == password_new_c {
				fmt.Println(password_old)
				fmt.Println(password_new)
				fmt.Println(password_new_c)
				hashpass, err := bcrypt.GenerateFromPassword([]byte(password_new), bcrypt.DefaultCost)
				if err != nil {
					log.Fatal(err)
				}
				user.Id = user_id
				user.Password = string(hashpass)
				var userModel models.UserModel
				userModel.UpdatePass(user)
				http.Redirect(response, request, "/dashboard/admin/account", http.StatusSeeOther)
			} else {
				fmt.Println("Error!")
				return
			}
		}
	}
}

func UpdatePasswordStaff(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	session, _ := store.Get(request, "mysession")
	sess_email := session.Values["email"]
	email := fmt.Sprint(sess_email)
	user_id, _ := strconv.ParseInt(request.Form.Get("user_id"), 10, 64)
	password_old := request.FormValue("password_old")
	password_new := request.FormValue("password_new")
	password_new_c := request.FormValue("password_confirm")
	var userModel models.UserModel
	users, _ := userModel.QueryUser(email)

	errf := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password_old))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("ERROR!")
	} else {
		if sess_email == users.Email {
			var user entities.User
			user.Password = password_new
			if user.Password == password_new_c {
				fmt.Println(password_old)
				fmt.Println(password_new)
				fmt.Println(password_new_c)
				hashpass, err := bcrypt.GenerateFromPassword([]byte(password_new), bcrypt.DefaultCost)
				if err != nil {
					log.Fatal(err)
				}
				user.Id = user_id
				user.Password = string(hashpass)
				var userModel models.UserModel
				userModel.UpdatePass(user)
				http.Redirect(response, request, "/dashboard/admin/account", http.StatusSeeOther)
			} else {
				fmt.Println("Error!")
				return
			}
		}
	}
}

func UpdatePasswordSiswa(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	session, _ := store.Get(request, "mysession")
	sess_email := session.Values["email"]
	email := fmt.Sprint(sess_email)
	user_id, _ := strconv.ParseInt(request.Form.Get("user_id"), 10, 64)
	password_old := request.FormValue("password_old")
	password_new := request.FormValue("password_new")
	password_new_c := request.FormValue("password_confirm")
	var userModel models.UserModel
	users, _ := userModel.QueryUser(email)

	errf := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password_old))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("ERROR!")
	} else {
		if sess_email == users.Email {
			var user entities.User
			user.Password = password_new
			if user.Password == password_new_c {
				fmt.Println(password_old)
				fmt.Println(password_new)
				fmt.Println(password_new_c)
				hashpass, err := bcrypt.GenerateFromPassword([]byte(password_new), bcrypt.DefaultCost)
				if err != nil {
					log.Fatal(err)
				}
				user.Id = user_id
				user.Password = string(hashpass)
				var userModel models.UserModel
				userModel.UpdatePass(user)
				http.Redirect(response, request, "/dashboard/siswa/account", http.StatusSeeOther)
			} else {
				fmt.Println("Error!")
				return
			}
		}
	}
}
