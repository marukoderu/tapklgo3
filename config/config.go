package config

import "database/sql"

// func GetDB() (*gorm.DB, error) {
// 	dbDriver := "mysql"
// 	dbName := "tapklgodb"
// 	dbUser := "root"
// 	dbPassword := ""
// 	db, err := gorm.Open(dbDriver, dbUser+":"+dbPassword+"@tcp(localhost:3306)/"+dbName+"?charset=utf8&parseTime=True")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

func GetDB() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/tapklgodb")
	if err != nil {
		return nil, err
	}
	return db, nil
}
