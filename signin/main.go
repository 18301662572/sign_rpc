package main

import (
	"code.oldbody.com/studygolang/mytest/signdemo/siginin-svc/common"
	"code.oldbody.com/studygolang/mytest/signdemo/siginin-svc/pb"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
	"time"
)

//设计一个连续签到的任务，连续签到7天，中间不能中断，如果中断了，就重新从第一天开始签到，连续签到三天奖励，连续签到7天奖励

//使用gin框架, http访问rpc

var conn *grpc.ClientConn

//登陆
func loginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username != "" && password != "" {
			//1.实例化gRPC客户端
			client := pb.NewUserServiceExtClient(conn)
			//2.用户登录验证
			reqluser := new(pb.LoginUserReq)
			reqluser.UserName = username
			reqluser.Password = password
			respluser, err := client.LoginUser(c, reqluser, grpc.CallCustomCodec(&common.JSONCoder{}))
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"msg": "用户名或密码错误！",
				})
				return
			}
			//将用户信息加入缓存
			c.SetCookie("ckusername",respluser.UserName,20,"/","127.0.0.1",false,true)
			//跳转到index页面
			index:=fmt.Sprintf("/index?uid=%s",strconv.Itoa(int(respluser.Id)))
			c.Redirect(http.StatusMovedPermanently,index)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": "用户名或密码不能为空",
			})
			return
		}
	} else {
		c.HTML(http.StatusOK, "web/login.html", gin.H{
			"msg": "登录页",
		})
	}
}

//主页面
func indexHandler(c *gin.Context) {
	//1.实例化gRPC客户端
	client := pb.NewUserServiceExtClient(conn)
	//获取用户最后一次签到信息
	reqsuserlast := new(pb.SignUserLastReq)
	uid,_:=strconv.ParseInt(c.Query("uid"),10,32)
	reqsuserlast.Uid= int32(uid)
	respsuserlast, err := client.SignUserLast(c, reqsuserlast, grpc.CallCustomCodec(&common.JSONCoder{}))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "查询签到信息有误！",
		})
		return
	}
	if respsuserlast != nil {
		if respsuserlast.SignDate < time.Now().Format("2006-01-02"){
			c.HTML(http.StatusOK, "web/index.html", gin.H{
				"uid": c.Query("uid"),
				"state":"true",
			})
		}else {
			c.HTML(http.StatusOK, "web/index.html", gin.H{
				"uid": c.Query("uid"),
			})
		}
	}else{
		c.HTML(http.StatusOK, "web/index.html", gin.H{
			"uid": c.Query("uid"),
		})
	}
}

//注册
func registerHandler(c *gin.Context) {
	if c.Request.Method=="POST"{
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username != "" && password != "" {
			//1.实例化gRPC客户端
			client := pb.NewUserServiceExtClient(conn)
			//2.将用户注册到用户表中
			reqruser := new(pb.RegistUserReq)
			reqruser.UserName = username
			reqruser.Password = password
			reqruser.Age = 18
			reqruser.NickName = "昵称"
			reqruser.IsDel = 0
			reqruser.CreateTime = time.Now().Format("2006-01-02 15:04:05")
			_, err := client.RegistUser(c, reqruser, grpc.CallCustomCodec(&common.JSONCoder{}))
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"msg": "注册失败",
					"error":err,
				})
				return
			}
			//3.跳转到login页面
			fmt.Println("注册成功")
			c.Redirect(http.StatusMovedPermanently, "/login")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": "用户名或密码不能为空",
			})
			return
		}
	}else{
		c.HTML(http.StatusOK,"web/register.html",gin.H{
			"msg":"注册页",
		})
	}
}

//签到
func signHandler(c *gin.Context){
	//1.实例化gRPC客户端
	client := pb.NewUserServiceExtClient(conn)
	//2.用户签到
	//signCount 今天的签到次数
	var signCount int32 = 1
	//查询用户最后一次的签到信息
	reqsuserlast := new(pb.SignUserLastReq)
	fmt.Println(c.Query("uid"))
	uid,_:=strconv.ParseInt(c.Query("uid"),10,32)
	reqsuserlast.Uid= int32(uid)
	respsuserlast, err := client.SignUserLast(c, reqsuserlast, grpc.CallCustomCodec(&common.JSONCoder{}))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "查询签到信息有误！",
		})
		return
	}
	if respsuserlast != nil && respsuserlast.Id!=0 {
		//判断用户的签到日期是否是昨天，如果是的话，签到次数+1；如果不是，签到次数=1
		today := time.Now().Format("2006-01-02")
		timeLayout := "2016/01/02"
		loc, _ := time.LoadLocation("Local")
		// 转成时间戳
		startUnix, _ := time.ParseInLocation(timeLayout, respsuserlast.SignDate, loc)
		endUnix, _ := time.ParseInLocation(timeLayout, today, loc)
		startTime := startUnix.Unix()
		endTime := endUnix.Unix()
		// 求相差天数
		date := (endTime - startTime) / 86400
		if date <= 1 {
			signCount += respsuserlast.SignCount
		}
	}
	reqsuser := new(pb.SignUserReq)
	reqsuser.Uid = reqsuserlast.Uid
	today := time.Now().Format("2006-01-02")
	createday:=time.Now().Format("2006-01-02 15:04:05")
	reqsuser.SignDate = today
	reqsuser.CreateTime = createday
	reqsuser.IsDel = 0
	reqsuser.SignCount = signCount
	_, err = client.SignUser(c, reqsuser, grpc.CallCustomCodec(&common.JSONCoder{}))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户签到失败！",
		})
		return
	}
	//4.如果连续签到3天，给用户10元优惠券，如果连续签到7天，给用户20元优惠券
	if signCount >= 7 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "您已经连续签到7天，获得一张20元优惠券！",
		})
	} else if signCount >= 3 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "您已经连续签到3天，获得一张10元优惠券！",
		})
	} else {
		msg := fmt.Sprintf("您已经连续签到%s天", signCount)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	}
}

func main() {
	//访问服务
	//1.创建与gRPC服务端的连接
	var err error
	// 与正常差别，在调用rpc接口时，指定自定义编解码
	conn, err = grpc.Dial("127.0.0.1:8090", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Printf("连接异常：err %s\n", err)
		return
	}

	//前台页面逻辑
	//创建一个默认的路由引擎
	r := gin.Default()
	//r.LoadHTMLGlob("tempaltes/*")
	r.LoadHTMLGlob("G:/GOWork/src/code.oldbody.com/studygolang/mytest/sign/signin/tempaltes/*")
	//加载静态页面(代码里使用的路径，实际保存静态文件的路径)
	r.Static("/script", "G:/GOWork/src/code.oldbody.com/studygolang/mytest/sign/signin/script")
	r.Static("/img", "G:/GOWork/src/code.oldbody.com/studygolang/mytest/sign/signin/icon")
	r.Any("/register", registerHandler)
	r.Any("/login", loginHandler)
	r.GET("/index", indexHandler)
	r.GET("/sign",signHandler)
	r.Run(":8081")
}
