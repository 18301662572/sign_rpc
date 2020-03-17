### 使用rpc创建登陆功能

##### 使用rpc实现签到功能 (实现了 对外REST，对内RPC的服务部署及外部访问http )
##### 目前存在的问题：
```text
 //1.时间计算存在问题？？？签到出国一天清零，两个日期间隔有问题
 //2.json.UnmarshalTypeError 不知道什么意思？？？
 // if err:=json.Unmarshal(u.Data,&user);err!=nil{
 //	if ute, ok := err.(*json.UnmarshalTypeError);
 //  	fmt.Printf("UnmarshalTypeError %v - %v - %v\n", ute.Value, ute.Type, ute.Offset)
 //	}
 //3. encoding/json 在go语言中可以json； 如果跨语言的话应该使用哪个包？？
```
     

架构
```text
signin-svc : -127.0.0.1:8090     rpc 内部服务 (对内 RPC)
signin-http: -:8082              http外部服务 (APIGateway) (对内 RPC,对外 REST)
signin-cli:  -:8083              client客户端访问http外部服务 (对外 REST)
signin:      -:8081              client客户端直接访问rpc
```
