package main

import (
	"code.oldbody.com/studygolang/mytest/signdemo/siginin-svc/pb"
	"code.oldbody.com/studygolang/mytest/signdemo/signin-http/common"
	"code.oldbody.com/studygolang/mytest/signdemo/signin-http/model"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
	"time"
)

//HTTP REST API服务（ApiGateWay）

var conn *grpc.ClientConn
var client pb.UserServiceExtClient

//注册
func RegisterHandler(w http.ResponseWriter, r *http.Request){
	var data []byte
	result:=new(model.UserInfo)
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	if username != "" && password != "" {
		//3.组装参数
		reqruser := new(pb.RegistUserReq)
		reqruser.UserName = username
		reqruser.Password = password
		reqruser.Age = 18
		reqruser.NickName = "昵称"
		reqruser.IsDel = 0
		reqruser.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		//4.调用接口
		_, err := client.RegistUser(r.Context(), reqruser, grpc.CallCustomCodec(&common.JSONCoder{}))
		if err != nil {
			result.Data=""
			result.Code=0
			result.State=1
			result.Msg="失败"
			data,_=json.Marshal(result)
		}else{
			result.Data=""
			result.Code=0
			result.State=0
			result.Msg="成功"
			data,_=json.Marshal(result)
		}
	}else{
		result.Data=""
		result.Code=0
		result.State=1
		result.Msg="失败"
		data,_=json.Marshal(result)
	}
	w.Write(data)
}

//登录
func LoginHandler(w http.ResponseWriter, r *http.Request){
	var data []byte
	var result=new(model.UserInfo)
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	if username != "" && password != "" {
		//3.组装参数
		reqluser := new(pb.LoginUserReq)
		reqluser.UserName = username
		reqluser.Password = password
		//4.调用接口
		respluser, err := client.LoginUser(r.Context(), reqluser, grpc.CallCustomCodec(&common.JSONCoder{}))
		if err != nil {
			result.Msg="失败"
			result.State=1
			result.Code=0
			result.Data=new(model.User)
			data,_=json.Marshal(result)
		}else{
			var user=new(model.User)
			user.Id=respluser.Id
			user.Username=respluser.UserName
			user.Password=respluser.Password
			user.Age=respluser.Age
			user.Nickname=respluser.NickName
			user.Createtime=respluser.CreateTime
			user.Isdel=respluser.IsDel
			result.Msg="成功"
			result.State=0
			result.Code=0
			result.Data=user
			data,_=json.Marshal(result)
		}
	} else {
		result.Msg="失败"
		result.State=1
		result.Code=0
		result.Data=new(model.User)
		data,_=json.Marshal(result)
	}
	w.Write(data)
}

//获取用户最后一次签到信息
func GetSignUserLastHandler(w http.ResponseWriter, r *http.Request){
	var data []byte
	var result =new(model.UserInfo)
	//获取用户最后一次签到信息
	id:=r.PostFormValue("uid")
	var (
		uid int64
		err error
		respsuserlast *pb.SignUserLastResp
	)
	if uid,err=strconv.ParseInt(id,10,32); err!=nil{
		result.Msg="用户Id有误"
		result.State=1
		result.Code=0
		result.Data=new(model.UserSign)
		data,_=json.Marshal(result)
	}else {
		//3.组装参数
		reqsuserlast := new(pb.SignUserLastReq)
		reqsuserlast.Uid = int32(uid)
		//调用接口
		respsuserlast, err = client.SignUserLast(r.Context(), reqsuserlast, grpc.CallCustomCodec(&common.JSONCoder{}))
		if err != nil {
			result.Msg="查询签到信息有误"
			result.State=1
			result.Code=0
			result.Data=new(model.UserSign)
			data,_=json.Marshal(result)
		} else {
			result.Msg="成功"
			result.State=0
			result.Code=0
			signuserlast:=new(model.UserSign)
			signuserlast.Id=respsuserlast.Id
			signuserlast.Isdel=respsuserlast.IsDel
			signuserlast.CreateTime=respsuserlast.CreateTime
			signuserlast.Uid=respsuserlast.Uid
			signuserlast.SignDate=respsuserlast.SignDate
			signuserlast.SignCount=respsuserlast.SignCount
			result.Data=signuserlast
			data,_=json.Marshal(result)
		}
	}
	w.Write(data)
}

//用户签到操作
func SignHandler(w http.ResponseWriter, r *http.Request){
	var data []byte
	var result =new(model.UserInfo)
	//signCount 今天的签到次数
	var (
		signCount int64
		uid int64
		err error
	)
	id:=r.PostFormValue("uid")
	count:=r.PostFormValue("signcount")
	if uid,err=strconv.ParseInt(id,10,32);err!=nil{
		result.Data=""
		result.Code=0
		result.State=1
		result.Msg="用户Id有误"
		data,_=json.Marshal(result)
	}else{
		if signCount,err=strconv.ParseInt(count,10,32);err!=nil{
			result.Data=""
			result.Code=0
			result.State=1
			result.Msg="用户签到次数有误"
			data,_=json.Marshal(result)
		}
		//3.组装参数
		reqsuser := new(pb.SignUserReq)
		reqsuser.Uid = int32(uid)
		today := time.Now().Format("2006-01-02")
		createday:=time.Now().Format("2006-01-02 15:04:05")
		reqsuser.SignDate = today
		reqsuser.CreateTime = createday
		reqsuser.IsDel = 0
		reqsuser.SignCount = int32(signCount)
		_, err = client.SignUser(r.Context(), reqsuser, grpc.CallCustomCodec(&common.JSONCoder{}))
		if err != nil {
			result.Data=""
			result.Code=0
			result.State=1
			result.Msg="用户签到失败"
			data,_=json.Marshal(result)
		}else {
			result.Data=""
			result.Code=0
			result.State=0
			result.Msg="签到成功"
			data,_=json.Marshal(result)
		}
	}
	w.Write(data)
}

func main(){
	//1.创建与gRPC服务端的连接
	//grpc.WithInsecure() 建立一个安全连接；注：与正常差别，在调用rpc接口时，指定自定义编解码
	var err error
	conn, err = grpc.Dial("127.0.0.1:8090", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Printf("连接异常：err %s\n", err)
		return
	}
	//2.实例化gRPC客户端
	client = pb.NewUserServiceExtClient(conn)

	//http监听
	//注册接口
	http.HandleFunc("/register",RegisterHandler)
	//登陆接口
	http.HandleFunc("/login", LoginHandler)
	//获取用户最后一次签到信息接口
	http.HandleFunc("/getsignuserlast", GetSignUserLastHandler)
	//用户签到接口
	http.HandleFunc("/sign", SignHandler)
	//http监听端口
	if err := http.ListenAndServe(":8082", nil);err!=nil{
		fmt.Println("监听8082端口有误，err:",err)
	}
}