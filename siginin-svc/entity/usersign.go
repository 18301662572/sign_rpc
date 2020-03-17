package entity

//UserSign 登录实体
type UserSign struct {
	Id         int32    `db:"Id" json:"id"`
	Uid        int32  `db:"Uid" json:"uid"`
	SignDate   string `db:"SignDate" json:"signdate"`
	SignCount  int32  `db:"SignCount" json:"signcount"`
	CreateTime string `db:"CreateTime" json:"createtime"`
	IsDel      int32  `db:"IsDel" json:"isdel"`
}
