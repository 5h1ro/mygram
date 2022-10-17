package database

import (
	"database/sql"
	"fmt"
	"log"
	"mygram/comment"
	"mygram/photo"
	"mygram/social_media"
	"mygram/user"

	"gorm.io/driver/postgres"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	db       *sql.DB
)

func InitializeDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=mygram port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database", err.Error())
	}

	if err != nil {
		log.Fatal("error while tyring to ping the database connection", err.Error())
	}

	fmt.Println("successfully connected to my database")

	database.Migrator().DropTable(
		&user.User{},
		&photo.Photo{},
		&comment.Comment{},
		&social_media.SocialMedia{},
	)

	database.AutoMigrate(
		&user.User{},
		&photo.Photo{},
		&comment.Comment{},
		&social_media.SocialMedia{},
	)

	fmt.Println("successfully migrating table")

	db, err = sql.Open("postgres", "postgresql://postgres:postgres@localhost:5432/mygram?sslmode=disable")

}

func GetDB() (*sql.DB, *gorm.DB) {
	return db, database
}
