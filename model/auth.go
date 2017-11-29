package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

var db *gorm.DB
func init()  {
	var err error
	db, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/mysite?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(db)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func Login(userName , password string) (string , bool)  {
	var user User
	db.Debug().Where("name = ? ",userName).First(&user)
	fmt.Println("userName is :",userName , "password is :" ,password)
	fmt.Println("user is :",user)
	if user.Name != "" && user.Password == password {
		return userName , true
	}
	defer db.Close()
	return userName, false
}

func Register()  {
	
}