package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/thanhlam/iot-workshop/model"
	"github.com/thanhlam/iot-workshop/repository"
	"gopkg.in/mgo.v2"
)

var DB *mgo.Database
var Session *mgo.Session

func init() {
	Session, err := mgo.Dial("mongodb://root:Vnpt123@54.169.219.90:27017/cas?authSource=admin")
	if err != nil {
		panic(err)
	}
	if err = Session.Ping(); err != nil {
		panic(err)
	}
	DB = Session.DB("users")
	fmt.Println("You connected to your mongo database.")
}

func CreateUser(c echo.Context) error {
	userSSO := new(model.UserSSO)
	err := c.Bind(userSSO)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, map[string]interface{}{"code": "6", "message": "Body is Invalid", "data": map[string]interface{}{"info": nil}})
	}
	profileRepository := repository.NewProfileRepositoryMongo(DB, "users")
	_, err = profileRepository.FindByUser(userSSO.Username)
	if err != mgo.ErrNotFound {
		return c.JSON(200, map[string]interface{}{"code": "2", "message": "User is Exists", "data": map[string]interface{}{"info": nil}})
	}
	userSSO.Userid = uuid.New().String()
	err = profileRepository.SaveUser(userSSO)
	if err != nil {
		fmt.Println(err)
		return c.JSON(200, map[string]interface{}{"code": "10", "message": "MongoDB connection refused", "data": map[string]interface{}{"info": nil}})
	}
	Session.Close()
	fmt.Println("Saved User success")
	return c.JSON(200, map[string]interface{}{"code": "0", "message": "Success", "data": map[string]interface{}{"info": nil}})
}
