package main

import (
	"code.oldbody.com/studygolang/mytest/signdemo/signin-cli/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//设计一个连续签到的任务，连续签到7天，中间不能中断，如果中断了，就重新从第一天开始签到，连续签到三天奖励，连续签到7天奖励
//使用gin框架, 模仿客户端 访问 http API服务（ApiGateWay）
//目前存在的问题：
//1.时间计算存在问题？？？签到出国一天清零，两个日期间隔有问题
//2.json.UnmarshalTypeError 不知道什么意思？？？
// if err:=json.Unmarshal(u.Data,&user);err!=nil{
//	if ute, ok := err.(*json.UnmarshalTypeError);
//  	fmt.Printf("UnmarshalTypeError %v - %v - %v\n", ute.Value, ute.Type, ute.Offset)
//	}
//3. encoding/json 在go语言中可以json； 如果跨语言的话应该使用哪个包？？


var conn *grpc.ClientConn
var client = &http.Client{}

//go client 访问http
func httpDo(r,params string) (body []byte,err error) {
	//client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:8082%s",r),
		strings.NewReader(params))
	if err != nil {
		return nil,err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body)) //打印返回文本
	return
}

//登陆
func loginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username != "" && password != "" {
			resp,err:=httpDo("/login",fmt.Sprintf("username=%s&password=%s",username,password))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg":   "服务器返回有误",
					"error": err,
				})
				return
			}
			//解析json
			var u=new(model.UserInfo)
			if err:= json.Unmarshal(resp,&u);err!=nil{
				c.JSON(http.StatusOK, gin.H{
					"msg":   "解析用户信息失败",
					"error": err,
				})
				return
			}
			if u.State==1{
				c.JSON(http.StatusOK, gin.H{
					"msg": "用户名或密码有误",
				})
				return
			}
			//二次解析
			user:=new(model.User)
			if err:=json.Unmarshal(u.Data,&user);err!=nil{
				c.JSON(http.StatusOK, gin.H{
					"msg":   "二次解析用户信息失败",
					"error": err,
				})
				return
			}
			//将用户信息加入缓存
			c.SetCookie("ckusername", user.Username, 20, "/", "127.0.0.1", false, true)
			//跳转到index页面
			index := fmt.Sprintf("/index?uid=%s", strconv.Itoa(int(user.Id)))
			c.Redirect(http.StatusMovedPermanently, index)
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
	uid:=c.Query("uid")
	//获取用户最后一次签到信息
	resp,err:=httpDo("/getsignuserlast",fmt.Sprintf("uid=%s",uid))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":   "服务器返回有误",
			"error": err,
		})
		return
	}
	var u =new(model.UserInfo)
	if err:=json.Unmarshal(resp,&u);err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "解析用户签到信息有误！",
			"error": err,
		})
		return
	}
	if u.State==1{
		c.JSON(http.StatusOK, gin.H{
			"msg": "解析用户签到信息失败！",
			"error": err,
		})
		return
	}
	var user=new(model.UserSign)
	if err:=json.Unmarshal(u.Data,&user);err!=nil{
		if ute, ok := err.(*json.UnmarshalTypeError); ok {
			fmt.Printf("UnmarshalTypeError %v - %v - %v\n", ute.Value, ute.Type, ute.Offset)
		}else{
			c.JSON(http.StatusOK, gin.H{
				"msg": "二次解析用户签到信息失败！",
				"error": err,
			})
			return
		}
	}
	if user.SignDate < time.Now().Format("2006-01-02"){
		c.HTML(http.StatusOK, "web/index.html", gin.H{
			"uid": c.Query("uid"),
			"state":"true",
		})
	}else {
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
			resp,err:=httpDo("/register",fmt.Sprintf("username=%s&password=%s",username,password))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg":   "服务器返回有误",
					"error": err,
				})
				return
			}
			//解析json
			var u=new(model.UserInfo)
			if err:= json.Unmarshal(resp,&u);err!=nil{
				c.JSON(http.StatusOK, gin.H{
					"msg":   "解析注册返回信息有误",
					"error": err,
				})
				return
			}
			if u.State==1{
				c.JSON(http.StatusOK, gin.H{
					"msg":   "注册失败",
					"error": err,
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
	uid:=c.Query("uid")
	//signCount 今天的签到次数
	var signCount int32 = 1
	//获取用户最后一次签到信息
	resp,err:=httpDo("/getsignuserlast",fmt.Sprintf("uid=%s",uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "服务器返回有误",
			"error": err,
		})
		return
	}
	//解析json
	var u=new(model.UserInfo)
	if err:= json.Unmarshal(resp,&u);err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"msg":   "解析用户最后一次签到信息有误",
			"error": err,
		})
		return
	}
	if u.State==1{
		c.JSON(http.StatusOK, gin.H{
			"msg": "查询用户最后一次签到信息失败！",
			"error": err,
		})
		return
	}
	var user=new(model.UserSign)
	if err:=json.Unmarshal(u.Data,&user);err!=nil{
		if ute, ok := err.(*json.UnmarshalTypeError); ok {
			fmt.Printf("UnmarshalTypeError %v - %v - %v\n", ute.Value, ute.Type, ute.Offset)
		}else{
			c.JSON(http.StatusOK, gin.H{
				"msg": "二次解析用户最后一次签到信息失败！",
				"error": err,
			})
			return
		}
	}
	//判断用户的签到日期是否是昨天，如果是的话，签到次数+1；如果不是，签到次数=1
	//时间计算存在问题？？？
	today := time.Now().Format("2006-01-02")
	timeLayout := "2016/01/02"
	loc, _ := time.LoadLocation("Local")
	// 转成时间戳
	startUnix, _ := time.ParseInLocation(timeLayout, user.SignDate, loc)
	endUnix, _ := time.ParseInLocation(timeLayout, today, loc)
	startTime := startUnix.Unix()
	endTime := endUnix.Unix()
	// 求相差天数
	date := (endTime - startTime) / 86400
	if date <= 1 {
		signCount += user.SignCount
	}
	//用户签到
	_,err=httpDo("/sign",fmt.Sprintf("uid=%s&signcount=%d",uid,signCount))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "服务器返回有误",
			"error": err,
		})
		return
	}
	//解析json
	var s=new(model.UserInfo)
	if err:= json.Unmarshal(resp,&s);err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"msg":   s.Msg,
			"error": err,
		})
		return
	}
	if s.State==1{
		c.JSON(http.StatusOK, gin.H{
			"msg":   s.Msg,
			"error": err,
		})
		return
	}
	fmt.Println("今日签到成功")
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
		msg := fmt.Sprintf("您已经连续签到%d天", signCount)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	}
}

func main() {
	//前台页面逻辑
	//创建一个默认的路由引擎
	r := gin.Default()
	r.Any("/register", registerHandler)
	r.Any("/login", loginHandler)
	r.GET("/index", indexHandler)
	r.GET("/sign",signHandler)
	r.Run(":8083")
}
