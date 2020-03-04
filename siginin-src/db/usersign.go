package db

import (
	"code.oldbody.com/studygolang/mytest/signdemo/siginin-src/entity"
	"github.com/golang/go/src/pkg/database/sql"
)

//用户签到Dal

//添加用户签到次数
func InsertUserSign(uid int32, signDate string, signcount int32, createTime string, isDel int32) error {
	_, err := db.Exec("insert into UserSign (Uid,SignDate,SignCount,CreateTime,IsDel) values(?,?,?,?,?)", uid, signDate, signcount, createTime, isDel)
	return err
}


//通过UID查询用户最后一次添加的签到信息
func SelectUserSignByUId(uid int32) (*entity.UserSign, error) {
	var usersign entity.UserSign
	err := db.Get(&usersign, "select Id,Uid,SignDate,SignCount,CreateTime,IsDel from usersign where uid=? ORDER BY SignDate DESC limit 1", uid)
	if err!=nil && err.Error() == sql.ErrNoRows.Error(){
		return nil, nil
	}
	return &usersign, err
}
