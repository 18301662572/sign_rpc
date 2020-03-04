package main
import (
	"code.oldbody.com/studygolang/mytest/signdemo/siginin-src/common"
	"code.oldbody.com/studygolang/mytest/signdemo/siginin-src/db"
	_ "code.oldbody.com/studygolang/mytest/signdemo/siginin-src/entity"
	"code.oldbody.com/studygolang/mytest/signdemo/siginin-src/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

//user服务，用户业务逻辑类Bll

type UserServiceExtServer struct {
}

var u = UserServiceExtServer{}

//注册用户
func (u *UserServiceExtServer) RegistUser(c context.Context, req *pb.RegistUserReq) (resp *pb.RegistUserResp, err error) {
	userName := req.UserName
	nickName := req.NickName
	password := req.Password
	age := req.Age
	createTime := req.CreateTime
	isDel := req.IsDel
	err = db.InsertUser(userName, nickName, password, age, createTime, isDel)
	return
}

//用户登录
func (u *UserServiceExtServer) LoginUser(c context.Context, req *pb.LoginUserReq) (resp *pb.LoginUserResp, err error) {
	userName := req.UserName
	password := req.Password
	user, err := db.SelectUserByNameAndPwd(userName, password)
	if err == nil && user!=nil {
		resp=new(pb.LoginUserResp)
		resp.Id = user.Id
		resp.UserName = user.UserName
		resp.Password = user.Password
		resp.NickName = user.NickName
		resp.Age = user.Age
		resp.IsDel = user.IsDel
	}
	return
}

//用户签到
func (u *UserServiceExtServer) SignUser(c context.Context, req *pb.SignUserReq) (resp *pb.SignUserResp, err error) {
	uid := req.Uid
	signDate := req.SignDate
	createTime := req.CreateTime
	isDel := req.IsDel
	signCount := req.SignCount
	err = db.InsertUserSign(uid, signDate, signCount, createTime, isDel)
	return
}

//获取用户最后一次的签到信息
func (u *UserServiceExtServer) SignUserLast(c context.Context, req *pb.SignUserLastReq) (resp *pb.SignUserLastResp, err error) {
	uid := req.Uid
	userSign, err := db.SelectUserSignByUId(uid)
	if err == nil && userSign != nil {
		resp=new(pb.SignUserLastResp)
		resp.Id= userSign.Id
		resp.Uid = userSign.Uid
		resp.SignDate = userSign.SignDate
		resp.SignCount = userSign.SignCount
		resp.CreateTime = userSign.CreateTime
		resp.IsDel = userSign.IsDel
	}
	return
}

//登陆服务
func main() {
	//连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/signin"
	err := db.InitDB(dsn)
	if err != nil {
		fmt.Printf("数据库连接异常：%s \n", err)
		return
	}
	//1.监听
	addr := "127.0.0.1:8090"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常：%s \n", err)
		return
	}
	fmt.Printf("开始监听：%s \n", addr)
	//2、实例化gRPC
	s := grpc.NewServer(grpc.CustomCodec(&common.JSONCoder{}))
	//3.在gRPC上注册微服务
	//第二个参数要接口类型的变量
	pb.RegisterUserServiceExtServer(s, &u)
	//4.启动gRPC服务端
	if err := s.Serve(lis); err != nil {
		//log.Fatalf("failed to serve: %v", err)
		fmt.Printf("failed to serve: %v", err)
	}
}
