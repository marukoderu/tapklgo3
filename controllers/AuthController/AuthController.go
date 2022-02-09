package AuthController

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/raisunee/tapklgo3/entities"
	"github.com/raisunee/tapklgo3/models"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func Login(response http.ResponseWriter, request *http.Request) {
	var tmp = template.Must(template.ParseFiles(
		"templates/auth/login.html",
		"templates/layouts/auth/header/auth-header.html",
		"templates/layouts/auth/footer/auth-footer.html",
	))
	tmp.ExecuteTemplate(response, "login", nil)
}

func LoginProcess(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	email := request.FormValue("email")
	password := fmt.Sprint(request.FormValue("password"))
	var userModel models.UserModel
	users, _ := userModel.QueryUser(email)

	errf := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		data := map[string]interface{}{
			"errM": "0",
		}
		var tmp = template.Must(template.ParseFiles(
			"templates/auth/login.html",
			"templates/layouts/auth/header/auth-header.html",
			"templates/layouts/auth/footer/auth-footer.html",
		))
		tmp.ExecuteTemplate(response, "login", data)
	} else {
		if users.Level == "3" {
			if email == users.Email {
				session, _ := store.Get(request, "mysession")
				session.Values["email"] = users.Email
				session.Values["name"] = users.Name
				session.Values["id"] = users.Id
				session.Values["photo"] = users.Photo
				session.Values["divisi"] = users.Divisi
				session.Values["level"] = users.Level
				session.Save(request, response)
				http.Redirect(response, request, "/dashboard/admin", http.StatusSeeOther)
			} else {
				data := map[string]interface{}{
					"errM": "0",
				}
				tmp, _ := template.ParseFiles("templates/auth/login.html")
				tmp.Execute(response, data)
				http.Redirect(response, request, "/login", http.StatusSeeOther)
			}
		} else if users.Level == "1" {
			if email == users.Email {
				session, _ := store.Get(request, "mysession")
				session.Values["email"] = users.Email
				session.Values["name"] = users.Name
				session.Values["id"] = users.Id
				session.Values["photo"] = users.Photo
				session.Values["divisi"] = users.Divisi
				session.Values["level"] = users.Level
				session.Save(request, response)
				http.Redirect(response, request, "/dashboard/siswa", http.StatusSeeOther)
			} else {
				data := map[string]interface{}{
					"errM": "0",
				}
				var tmp = template.Must(template.ParseFiles(
					"templates/auth/login.html",
					"templates/layouts/auth/header/auth-header.html",
					"templates/layouts/auth/footer/auth-footer.html",
				))
				tmp.ExecuteTemplate(response, "login", data)
				http.Redirect(response, request, "/login", http.StatusSeeOther)
			}
		} else if users.Level == "2" {
			if email == users.Email {
				session, _ := store.Get(request, "mysession")
				session.Values["email"] = users.Email
				session.Values["name"] = users.Name
				session.Values["id"] = users.Id
				session.Values["photo"] = users.Photo
				session.Values["divisi"] = users.Divisi
				session.Values["level"] = users.Level
				session.Save(request, response)
				http.Redirect(response, request, "/dashboard/staff", http.StatusSeeOther)
			} else {
				data := map[string]interface{}{
					"errM": "0",
				}
				tmp, _ := template.ParseFiles("templates/auth/login.html")
				tmp.Execute(response, data)
				http.Redirect(response, request, "/login", http.StatusSeeOther)
			}
		} else {
			data := map[string]interface{}{
				"errM": "1",
			}
			var tmp = template.Must(template.ParseFiles(
				"templates/auth/login.html",
				"templates/layouts/auth/header/auth-header.html",
				"templates/layouts/auth/footer/auth-footer.html",
			))
			tmp.ExecuteTemplate(response, "login", data)
		}
	}
}

func Register(response http.ResponseWriter, request *http.Request) {
	var tmp = template.Must(template.ParseFiles(
		"templates/auth/register.html",
		"templates/layouts/auth/header/auth-header.html",
		"templates/layouts/auth/footer/auth-footer.html",
	))
	tmp.ExecuteTemplate(response, "register", nil)
}

func RegisterProcess(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var user entities.User
	user.Name = request.Form.Get("name")
	user.Email = request.Form.Get("email")
	pass := fmt.Sprint(request.Form.Get("password"))
	hashpass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	user.Password = string(hashpass)
	user.Level = "1"

	if err != nil {
		log.Fatal(err)
	}

	var userModel models.UserModel
	userModel.Create(&user)
	http.Redirect(response, request, "/login", http.StatusSeeOther)
}

func Logout(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	session.Options.MaxAge = -1
	session.Save(request, response)
	http.Redirect(response, request, "/login", http.StatusSeeOther)
}
