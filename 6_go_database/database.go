package godatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/db_test?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                  // Pengaturan jumlah koneksi minimal yang dibuat
	db.SetMaxOpenConns(100)                 // Jumlah koneksi maksimal yang dibuat
	db.SetConnMaxIdleTime(5 * time.Minute)  // Pengaturan berapa lama koneksi yang sudah tidak di pakai akan di hapus
	db.SetConnMaxLifetime(60 * time.Minute) // Pengaturan berapa lama koneksi boleh digunakan
	return db
}
