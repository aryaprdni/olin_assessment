package app

import (
	"aryaprdni/assessment_golang/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres dbname=olin_assessment sslmode=disable password=root")

	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
