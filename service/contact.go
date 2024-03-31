package service

import (
	"ImApplication.go/model"
	"errors"
	"time"
)

type ContactService struct {
}

// 添加好友
func (service *ContactService) AddFriend(
	userid,
	dstid int64) error {
	//不能添加自己为好友。
	if userid == dstid {
		return errors.New("无法加自己为好友！")
	}
	//判断这个记录是否已经存在，已存在则无需二次添加
	tmp := model.Contact{}
	DbEngine.Where("ownerid = ? and dstobj = ?", userid, dstid).Get(&tmp)
	if tmp.Id > 0 {
		return errors.New("已经是好友了，无法重复添加！")
	}
	//插入数据，开启事物来保证一致性
	session := DbEngine.NewSession()
	//开启事物
	session.Begin()
	//插入自己和对方的好友关系
	_, err1 := DbEngine.InsertOne(model.Contact{
		Ownerid:  userid,
		Dstobj:   dstid,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})
	//插入对方和自己的好友关系
	_, err2 := DbEngine.InsertOne(model.Contact{
		Ownerid:  dstid,
		Dstobj:   userid,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})
	if err1 != nil && err2 != nil {
		//提交事务。将操作执行到数据库中
		session.Commit()
		return nil
	} else {
		//代表发生了操作，回滚事物，并返回错误信息
		session.Rollback()
		if err1 != nil {
			return err1
		}
		return err2
	}

}

// 查找好友列表
func (serive *ContactService) SearchFriend(userId int64) []model.User {
	contacts := make([]model.Contact, 0)
	//获取所有的用户id
	DbEngine.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_USER).Find(&contacts)
	//根据用户id来组成预查询条件 新建一个数组来保存用户id
	objIds := make([]int64, 0)
	for _, v := range contacts {
		objIds = append(objIds, v.Dstobj)
	}
	//新建返回的数据
	users := make([]model.User, 0)
	if len(objIds) == 0 {
		return users
	}
	DbEngine.In("id", objIds).Find(&users)
	return users
}

// 获取群列表
func (service *ContactService) SearchComunity(userId int64) []model.Community {
	conconts := make([]model.Contact, 0)
	comIds := make([]int64, 0)

	DbEngine.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_COMUNITY).Find(&conconts)
	for _, v := range conconts {
		comIds = append(comIds, v.Dstobj)
	}
	coms := make([]model.Community, 0)
	if len(comIds) == 0 {
		return coms
	}
	DbEngine.In("id", comIds).Find(&coms)
	return coms
}

// 加入
func (service *ContactService) JoinCommunity(userId, comId int64) error {
	cot := model.Contact{
		Ownerid: userId,
		Dstobj:  comId,
		Cate:    model.CONCAT_CATE_COMUNITY,
	}
	DbEngine.Get(&cot)
	if cot.Id == 0 {
		cot.Createat = time.Now()
		_, err := DbEngine.InsertOne(cot)
		return err
	} else {
		return nil
	}
}

// 建群
func (service *ContactService) CreateCommunity(comm model.Community) (ret model.Community, err error) {
	if len(comm.Name) == 0 {
		err = errors.New("缺少群名称")
		return ret, err
	}
	if comm.Ownerid == 0 {
		err = errors.New("请先登录")
		return ret, err
	}
	com := model.Community{
		Ownerid: comm.Ownerid,
	}
	num, err := DbEngine.Count(&com)

	if num > 5 {
		err = errors.New("一个用户最多只能创见5个群")
		return com, err
	} else {
		comm.Createat = time.Now()
		session := DbEngine.NewSession()
		session.Begin()
		_, err = session.InsertOne(&comm)
		if err != nil {
			session.Rollback()
			return com, err
		}
		_, err = session.InsertOne(
			model.Contact{
				Ownerid:  comm.Ownerid,
				Dstobj:   comm.Id,
				Cate:     model.CONCAT_CATE_COMUNITY,
				Createat: time.Now(),
			})
		if err != nil {
			session.Rollback()
		} else {
			session.Commit()
		}
		return com, err
	}
}

// 获取到用户加入的群id列表
func (service *ContactService) SearchComunityIds(userId int64) (comIds []int64) {
	conconts := make([]model.Contact, 0)
	comIds = make([]int64, 0)

	DbEngine.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_COMUNITY).Find(&conconts)
	for _, v := range conconts {
		comIds = append(comIds, v.Dstobj)
	}
	return comIds
}
