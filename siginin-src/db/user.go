package db

import (
	"code.oldbody.com/studygolang/mytest/sign/siginin-src/entity"
	"database/sql"
)

//用户Dal

//InsertUser 添加用户实体
func InsertUser(username string,nickname string,password string,age int32,createtime string,isdel int32) error{
	_,err:=db.Exec("insert into User (UserName,NickName,PassWord,Age,CreateTime,IsDel) values(?,?,?,?,?,?)",username,nickname,password,age,createtime,isdel)
	return err
}

//SelectUser 通过用户ID查询用户信息
func SelectUser(id int)(*entity.User,error){
	user:=entity.User{}
	err:=db.Get(&user,"select Id,UserName,NickName,PassWord,Age,CreateTime,IsDel from User where id=?",id)
	return &user,err
}

//SelectUserByNameAndPwd 通过用户名密码查询用户信息
func SelectUserByNameAndPwd(username,password string) (*entity.User,error){
	var user entity.User
	err:=db.Get(&user,"select Id,UserName,NickName,PassWord,Age,CreateTime,IsDel from User where UserName=? and Password=? and IsDel=0 limit 1",username,password)
	if err!=nil && err.Error() == sql.ErrNoRows.Error(){
		return nil, nil
	}
	return &user,err
}

