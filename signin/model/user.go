package model

//User 用户实体
type User struct {
	id int `db:"Id" json:"id"`
	username string `db:"UserName" json:"username"`
	nickname string `db:"NickName" json:"nickname"`
	password string	`db:"PassWord" json:"password"`
	age int `db:"Age" json:"age"`
	createtime string `db:"CreateTime" json:"createtime"`
	isdel int8 `db:"IsDel" json:"isdel"`
}
