package database

import (
	"context"
	"database/sql"
	"example/model"
	"fmt"
	"log"
	"strconv"

	"github.com/microsoft/go-mssqldb/azuread"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	var err error
	var DB *gorm.DB
	var sqlDB *sql.DB
	p := viper.GetString("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("failed to parse database port")
	}

	connStr := fmt.Sprintf("Server=tcp:%s,%d;Initial Catalog=%s;Persist Security Info=False;User ID=%s;Password=%s;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;",
		//local
		// connStr := fmt.Sprintf("Server=tcp:%s,%d;Database=%s;User ID=%s;Password=%s;TrustServerCertificate=true;Trusted_Connection=False;Encrypt=True;",
		viper.GetString("DB_HOST"),
		port,
		viper.GetString("DB_NAME"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"))

	sqlDB, err = sql.Open(azuread.DriverName, connStr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
	}

	ctx := context.Background()
	err = sqlDB.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	DB, err = gorm.Open(sqlserver.New(sqlserver.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		log.Fatal("GORM failed to connect database: ", err.Error())
	}

	if err != nil {
		panic("failed to connect database")
	}

	if viper.GetBool("DB_MIGRATION") {
		fmt.Println("Connection Opened to Database")
		// DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
		DB.AutoMigrate(
			&model.SmartSolution{},
			&model.SolutionGroup{},
			&model.Product{},
			&model.ProductDetail{},
			&model.Template{},
		)
		fmt.Println("Database Migrated")
	}

	return DB
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
