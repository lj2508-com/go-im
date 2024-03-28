package service

import (
	"ImApplication.go/model"
	"ImApplication.go/util"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type UserService struct {
}

func (userService *UserService) Register(
	mobile, //手机
	plainpwd, //明文密码
	nickname, //昵称
	avatar, sex string) (user model.User, err error) {
	tmp := model.User{}
	//先检查手机号是否已注册，如果已注册，则返回错误，否则插入数据
	_, err = DbEngine.Where("mobile=? ", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已经注册")
	}
	tmp.Mobile = mobile
	//进行密码加密
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = util.MakePasswd(plainpwd, tmp.Salt)
	tmp.Sex = sex
	tmp.Nickname = nickname
	tmp.Avatar = avatar
	tmp.Createat = time.Now()
	//token 先预设
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())

	_, err = DbEngine.InsertOne(&tmp)
	//前端恶意插入特殊字符
	//数据库连接操作失败
	return tmp, err
}

func (userService *UserService) Login(
	mobile, //手机
	plainpwd string) (user model.User, err error) {
	//通过手机号来查询是否存在账号，然后判断密码是否正确，密码错误和手机号不存在返回同一个错误
	tmp := model.User{}
	DbEngine.Where("mobile = ?", mobile).Get(&tmp)
	if tmp.Id == 0 {
		return tmp, errors.New("账号未注册，或密码输入错误")
	}
	if !util.ValidatePasswd(tmp.Passwd, tmp.Salt, plainpwd) {
		return tmp, errors.New("密码不正确")
	}
	//生成临时token
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)
	tmp.Token = token
	//返回数据
	DbEngine.ID(tmp.Id).Cols("token").Update(&tmp)
	return tmp, nil
}
