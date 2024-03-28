package ctrl

import (
	"ImApplication.go/model"
	"ImApplication.go/service"
	"ImApplication.go/util"
	"fmt"
	"math/rand"
	"net/http"
)

func UserLogin(writer http.ResponseWriter,
	request *http.Request) {

	request.ParseForm()

	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	//模拟
	user, err := userService.Login(mobile, passwd)

	if err != nil {
		util.ResponseFail(writer, err.Error())
	} else {
		util.ResponseOk(writer, "", user)
	}

}

var userService service.UserService

func UserRegister(writer http.ResponseWriter,
	request *http.Request) {

	request.ParseForm()
	//
	mobile := request.PostForm.Get("mobile")
	//
	plainpwd := request.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW

	user, err := userService.Register(mobile, plainpwd, nickname, avatar, sex)
	if err != nil {
		util.ResponseFail(writer, err.Error())
	} else {
		util.ResponseOk(writer, "", user)

	}

}
