package service

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo"
	"github.com/thanhlam/iot-workshop/model"
)

func ParseSSOToken(c echo.Context) error {
	authenRequestBody := new(model.AuthenRequestBody)
	err := c.Bind(authenRequestBody)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, map[string]interface{}{"code": "6", "message": "Body is Invalid", "data": map[string]interface{}{"info": nil}})
	}
	token := authenRequestBody.Token
	authResponse, err := BasicAuth(token)
	//
	if err != nil {
		return c.JSON(200, map[string]interface{}{"code": "10", "message": "Connection refused", "data": map[string]interface{}{"info": nil}})
	}
	//
	var result map[string]interface{}
	json.Unmarshal([]byte(authResponse), &result)

	if result["error"] != nil {
		if result["message"] != nil {
			return c.JSON(200, map[string]interface{}{"code": "7", "message": (result["message"]).(string), "data": map[string]interface{}{"info": nil}})
		}
		return c.JSON(200, map[string]interface{}{"code": "5", "message": (result["error"]).(string), "data": map[string]interface{}{"info": nil}})
	}
	//test
	attributes := result["attributes"].(map[string]interface{})
	attributes["username"] = result["id"]
	return c.JSON(200, map[string]interface{}{"code": "0", "message": "success", "data": map[string]interface{}{"info": attributes}})
}
