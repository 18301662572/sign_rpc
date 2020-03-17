package model

//UserSign 登录实体
type UserSign struct {
	Id         int    `db:"Id" json:"id"`
	Uid        string `db:"Uid" json:"uid"`
	SignDate   string `db:"SignDate" json:"signdate"`
	SignCount  int32  `db:"SignCount" json:"signcount"`
	CreateTime string `db:"CreateTime json:"createtime"`
	Isdel      int8   `db:"IsDel" json:"isdel"`
}
