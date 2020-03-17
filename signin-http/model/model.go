package model

type UserInfo struct{
	Msg string `json:"msg"`//成功，失败
	Data *User `json:"data"`
	State int `json:"state"` //0：成功 1：失败
	Code int `json:"code"` //http返回值
}

type User struct {
	Id int32 `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string	`json:"password"`
	Age int32 `json:"age"`
	Createtime string `json:"createtime"`
	Isdel int32 `json:"isdel"`
}

//无返回实体
type SignInfo struct{
	Msg string `json:"msg"`//成功，失败
	Data string `json:"data"` //没有返回值
	State int `json:"state"` //0：成功 1：失败
	Code int `json:"code"` //http返回值
}

type SignUserLast struct {
	Msg string `json:"msg"`//成功，失败
	Data *UserSign `json:"data"`
	State int `json:"state"` //0：成功 1：失败
	Code int `json:"code"` //http返回值
}

type UserSign struct {
	Id         int32    `json:"id"`
	Uid        int32 `json:"uid"`
	SignDate   string `json:"signdate"`
	SignCount  int32  `json:"signcount"`
	CreateTime string `json:"createtime"`
	Isdel      int32   `json:"isdel"`
}

