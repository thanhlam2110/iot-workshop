package model

//push
type PushMessage struct {
	Token    string `json:"token"`
	Thingid  string `json:"thingid"`
	Thingkey string `json:"thingkey"`
	Chanelid string `json:"chanelid"`
	Message  string `json:"message"`
}

//User struct
type UserSSO struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Userid       string `json:"lastname"`
	Usermail     string `json:"usermail"`
	Userstatus   string `json:"userstatus"`
	Userparentid string `json:"userparentid"`
	Usertype     string `json:"usertype"`
}
type UserSSOs []UserSSO
type AuthenRequestBody struct {
	Token string `json:"token"`
}
type MapThingChanel struct {
	Thingid      string `json:"thingid"`
	Chanelid     string `json:"chanelid"`
	Mapstatus    string `json:"mapstatus"`
	Userparentid string `json:"userparentid"`
}
