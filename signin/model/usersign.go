package model

//UserSign 登录实体
type UserSign struct {
	id         int    `db:"Id" json:"id"`
	uid        string `db:"Uid" json:"uid"`
	signDate   string `db:"SignDate" json:"signdate"`
	signCount  int32  `db:"SignCount" json:"signcount"`
	createTime string `db:"CreateTime json:"createtime"`
	isdel      int8   `db:"IsDel" json:"isdel"`
}
