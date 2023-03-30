package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
)

func NewDatabase()(*gorm.DB, error){
     fmt.Println("Setting up new DB Connection")
	 dbUsername := "postgres"
	 dbPassword := "abayomi"
	 dbHost := "localhost"
	 //dbHost := "host.docker.internal"
	 dbTable:= "commentsApi"
	 dbPort := 5432
 
 
 
	 connectString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)  
	 db, err := gorm.Open("postgres", connectString) 
	 if err != nil {
		return db, err
	   }

	   if err := db.DB().Ping(); err != nil{
		return db, err
	}
	fmt.Println("Setting up new DB Connection 22")
		return db, nil
}