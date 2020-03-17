package model

//User 用户实体
type User struct {
	Id int `db:"Id" json:"id"`
	Username string `db:"UserName" json:"username"`
	Nickname string `db:"NickName" json:"nickname"`
	Password string	`db:"PassWord" json:"password"`
	Age int `db:"Age" json:"age"`
	Createtime string `db:"CreateTime" json:"createtime"`
	Isdel int8 `db:"IsDel" json:"isdel"`
}
