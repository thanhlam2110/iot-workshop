package service

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/labstack/echo"
	"github.com/thanhlam/iot-workshop/model"
	"github.com/thanhlam/iot-workshop/repository"
	"gopkg.in/mgo.v2"
)

func PushMessage(c echo.Context) error {
	pushMessage := new(model.PushMessage)
	err := c.Bind(pushMessage)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, map[string]interface{}{"code": "6", "message": "Body is Invalid", "data": map[string]interface{}{"info": nil}})
	}
	//parse SSO TOKEN
	//parse token ---> get user status --> kiem tra user status
	authResponse, err := BasicAuth(pushMessage.Token)

	if err != nil {
		return c.JSON(200, map[string]interface{}{"code": "10", "message": "Connection refused", "data": map[string]interface{}{"info": nil}})
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(authResponse), &result)
	if result["error"] != nil {
		if result["message"] != nil {
			return c.JSON(200, map[string]interface{}{"code": "7", "message": (result["message"]).(string), "data": map[string]interface{}{"info": nil}})
		}
		return c.JSON(200, map[string]interface{}{"code": "5", "message": (result["error"]).(string), "data": map[string]interface{}{"info": nil}})
	}
	attributes := result["attributes"].(map[string]interface{})
	if attributes["userstatus"] != "ACTIVE" {
		return c.JSON(200, map[string]interface{}{"code": "8", "message": "USER DISABLED", "data": map[string]interface{}{"info": nil}})
	}
	//check map thing to chanel
	profileRepository := repository.NewProfileRepositoryMongo(DB, "things_map_chanels")
	_, err = profileRepository.FindMapThingChanel(pushMessage.Thingid, pushMessage.Chanelid)
	if err == mgo.ErrNotFound {
		return c.JSON(200, map[string]interface{}{"code": "2", "message": "ThingId haven't mapped to Chanel", "data": map[string]interface{}{"info": nil}})
	}
	//connect broker
	connect, err := ConnectMQTT(pushMessage.Thingid, pushMessage.Thingkey)
	if err != nil {
		return c.JSON(200, map[string]interface{}{"code": "0", "message": err.Error(), "data": map[string]interface{}{"info": nil}})
	}
	//b81c709b-9bf1-4d3c-882c-0aea7feaa1bf
	if connect != nil {
		connect.Publish(pushMessage.Chanelid, 0, false, pushMessage.Message)
		//connect.Publish("lam-send", 0, false, pushMessage.Message)
	}
	connect.Disconnect(250)
	return c.JSON(200, map[string]interface{}{"code": "0", "message": "Push Success", "data": map[string]interface{}{"info": nil}})
}
func ConnectMQTT(thingid, thingkey string) (c mqtt.Client, err error) {
	link := "tcp://18.141.164.223:1883"
	opts := mqtt.NewClientOptions().AddBroker(link).SetClientID(thingid)
	c = mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return c, nil
}
