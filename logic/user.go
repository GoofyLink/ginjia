package logic

import (
	"blog.com/dao/mysql"
	"blog.com/models"
	"blog.com/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户是否存在
	if err = mysql.CheckUserExists(p.Username); err != nil {
		return err
	}
	// 2. 生成UID
	userID := snowflake.GenID()
	// 3. 构造user实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//3. 插入数据库
	if err = mysql.InsertUsername(user); err != nil {
		return err
	}
	return err
}

func Login(p *models.ParamLogin) (err error) {
	// 直接登录1
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}
