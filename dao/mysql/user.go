package mysql

import (
	"blog.com/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

const secret = "goofy"

// CheckUserExists 检查用户是否存在
func CheckUserExists(username string) (err error) {
	// 查询
	sqlStr := `select count(user_id) from users where username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

// InsertUsername 向数据库中插入一条新的记录
func InsertUsername(user *models.User) error {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 入库
	sqlStr := `insert into users(user_id,username,password) values (?,?,?)`
	_, err := db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))

	// 转换成十六进制字符串
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}