package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Email    string
	Password string
	Level    int
}

type Response struct {
	Status  bool
	Message string
	Data    []Users
}

func Koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/tapklgodb")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Add Users
func Register(email string, password string, level int) Response {
	db, err := Koneksi()
	if err != nil {
		return Response{
			Status:  false,
			Message: "Connection Failed: " + err.Error(),
		}
	}
	defer db.Close()
	_, err = db.Exec("insert into users values (?, ?, ?)", email, password, level)
	if err != nil {
		return Response{
			Status:  false,
			Message: "Failed to Insert: " + err.Error(),
			Data:    []Users{},
		}
	}
	return Response{
		Status:  true,
		Message: "Success to Insert User",
		Data:    []Users{},
	}
}
