package model

type UserInfo struct{
	Msg string `json:"msg"`//成功，失败
	Data *User `json:"data"`
	State int `json:"state"` //0：成功 1：失败
	Code int `json:"code"` //http返回值
}

type SignUserLastInfo struct {
	Msg string `json:"msg"`//成功，失败
	Data *UserSign `json:"data"`
	State int `json:"state"` //0：成功 1：失败
	Code int `json:"code"` //http返回值
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