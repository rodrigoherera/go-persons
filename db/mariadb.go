package db

import (
	"fmt"
	"go-persons/models"
	"log"
	"os"
	"strconv"
	"time"

	//Mysql dialect from gorm pkg
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Client ..
var Client *gorm.DB

//CreateMariaDB create a new instance
func CreateMariaDB() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbHost := os.Getenv("DB_HOST")
	DbPort := "3306"

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	retryCount := 30
	for {
		conn, err := gorm.Open("mysql", DBURL)
		if err != nil {
			log.Println(err)
			retryCount--
			time.Sleep(2 * time.Second)
		} else {

			conn.SingularTable(true)
			conn.Debug()
			conn.LogMode(true)
			Client = conn
			return
		}
	}
}

//AddPerson add new person into mariadb
func AddPerson(person *models.Person) (int, error) {
	tx := Client.Begin()

	if err := tx.Error; err != nil {
		tx.Rollback()
		fmt.Println(err)
		return 500, err
	}

	err := tx.Create(person).Error
	if err != nil {
		tx.Rollback()
		return 500, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return 500, err
	}
	return 201, nil
}

//AddUser add new user into mariadb
func AddUser(user *models.User) (int, error) {
	tx := Client.Begin()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return 500, err
	}

	err := tx.Create(user).Error
	if err != nil {
		tx.Rollback()
		return 500, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return 500, err
	}
	return 201, nil
}

//GetPerson get a existing person
func GetPerson(id string) (models.Person, error) {
	var p models.Person
	idInt, err := strconv.Atoi(id)
	if err != nil {
		//
	}
	err = Client.Select("id, name, lastname, age, dni").Model(&models.Person{}).Where("id = ?", idInt).Find(&p).Error

	if err != nil {
		return p, err
	}
	return p, nil
}

//GetAllPerson get all persons
func GetAllPerson() ([]models.Person, error) {
	var p []models.Person

	err := Client.Select("id, name, lastname, age, dni").Model(&models.Person{}).Find(&p).Error

	if err != nil {
		return p, err
	}
	return p, nil
}

//UpdatePerson update a person
func UpdatePerson(person models.Person, newPerson models.Person, c *gorm.DB) (int, error) {
	tx := c.Begin()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return 500, err
	}

	tx.First(&person)
	person.Name = newPerson.Name
	person.LastName = newPerson.LastName
	person.Age = newPerson.Age
	person.Dni = newPerson.Dni

	err := tx.Save(&person).Error
	if err != nil {
		tx.Rollback()
		return 500, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return 500, err
	}
	return 200, nil
}

//DeletePerson delete a person
func DeletePerson(person models.Person, c *gorm.DB) (int, error) {
	tx := c.Begin()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return 500, err
	}

	err := tx.Delete(&person).Error
	if err != nil {
		tx.Rollback()
		return 500, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return 500, err
	}
	return 200, nil
}

//CheckUserExistence check if a user exists
func CheckUserExistence(u *models.User) (bool, error) {

	err := Client.Select("id, password").Model(&models.User{}).Where("email = ?", u.Email).Find(&u).Error

	if err != nil {
		return false, err
	}
	if u.ID == 0 {
		return false, nil
	}
	return true, nil
}
