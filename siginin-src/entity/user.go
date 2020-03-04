package entity

//User 用户实体
type User struct {
	Id         int32  `db:"Id" json:"id"`
	UserName   string `db:"UserName" json:"username"`
	NickName   string `db:"NickName" json:"nickname"`
	Password   string `db:"PassWord" json:"password"`
	Age        int32  `db:"Age" json:"age"`
	CreateTime string `db:"CreateTime" json:"createtime"`
	IsDel      int32  `db:"IsDel" json:"isdel"`
}
