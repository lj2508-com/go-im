package ctrl

import (
	"ImApplication.go/args"
	"ImApplication.go/model"
	"ImApplication.go/service"
	"ImApplication.go/util"
	"net/http"
)

var contactService service.ContactService

func LoadFriend(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	util.Bind(req, &arg)

	users := contactService.SearchFriend(arg.Userid)
	util.RespOkList(w, users, len(users))
}

func Addfriend(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	util.Bind(req, &arg)
	//调用service
	err := contactService.AddFriend(arg.Userid, arg.Dstid)
	//
	if err != nil {
		util.ResponseFail(w, err.Error())
	} else {
		util.ResponseOk(w, "好友添加成功", nil)
	}
}

func LoadCommunity(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	util.Bind(req, &arg)
	comunitys := contactService.SearchComunity(arg.Userid)
	util.RespOkList(w, comunitys, len(comunitys))
}
func JoinCommunity(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	util.Bind(req, &arg)
	err := contactService.JoinCommunity(arg.Userid, arg.Dstid)
	//todo 刷新用户的群组信息 AddGroupId(arg.Userid,arg.Dstid)

	if err != nil {
		util.ResponseFail(w, err.Error())
	} else {
		util.ResponseOk(w, "", nil)
	}
}
func CreateCommunity(w http.ResponseWriter, req *http.Request) {
	var arg model.Community
	//如果这个用的上,那么可以直接
	util.Bind(req, &arg)
	com, err := contactService.CreateCommunity(arg)
	if err != nil {
		util.ResponseFail(w, err.Error())
	} else {
		util.ResponseOk(w, "", com)
	}
}
