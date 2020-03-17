package model

import "encoding/json"

//Go的json解析：Marshal与Unmarshal
//https://www.jianshu.com/p/3534532e06ed

//“简单数据”：是指不能再进行二次json解析的数据，如”Msg”，”State”，“Code” 只能进行一次json解析。
//“复合数据”：类似“Data”==> ”User\”:{\”id\”:1,\”username\”:"zs",\”nickname\”:"zs",\”password\”:"123",...}这样的数据，是可进行二次甚至多次json解析的，
// 因为它的value也是个可被解析的独立json。 即第一次解析key为User的value，第二次解析value中的key为id,username....

//对于”复合数据”，如果接收体中配的项被声明为interface{}类型，go都会默认解析成map[string]interface{}类型。
// 如果我们想直接解析到struct User对象中，可以将接受体对应的项定义为 *User 类型。
// 如果不想指定Data变量为具体的类型，仍想保留interface{}类型，但又希望该变量可以解析到struct Data对象中,
// 我们可以将该变量定义为json.RawMessage类型
// 接收体中，被声明为json.RawMessage类型的变量在json解析时，变量值仍保留json的原值，即未被自动解析为map[string]interface{}类型。
// 如变量Data解析后的值为：{\”id\”:1,\”username\”:"zs",\”nickname\”:"zs",\”password\”:"123",...}
// 在第一次json解析时，变量Class的类型是json.RawMessage。此时，我们可以对该变量进行二次json解析，因为其值仍是个独立且可解析的完整json串。我们只需再定义一个新的接受体即可，如json.Unmarshal(u.data,&user)


//http rest api服务端返回的json结构体
type UserInfo struct{
	Msg string `json:"msg"`//成功，失败
	Data json.RawMessage `json:"data"` //Data数据类型： interface{}: *User/string/*UserSign
	State int `json:"state"` //0：成功 1：失败
	Code int `json:"code"` //http返回值
}
