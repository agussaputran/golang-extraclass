package connection

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//sementara kosong
var dbConn *gorm.DB

func Connect() {
	psqlHost := os.Getenv("PSQL_HOST")
	psqlUser := os.Getenv("PSQL_USER")
	psqlPw := os.Getenv("PSQL_PASSWORD")
	psqlDb := os.Getenv("PSQL_DB")
	psqlPort := os.Getenv("PSQL_PORT")

	dsn := "host=" + psqlHost +
		" user=" + psqlUser +
		" password=" + psqlPw +
		" dbname=" + psqlDb +
		" port=" + psqlPort +
		" sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to db..")
	} else {
		log.Println("Berhasil connect")
	}

	dbConn = db
}

func GetConnection() *gorm.DB {
	return dbConn
}
