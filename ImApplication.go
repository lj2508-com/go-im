package main

import (
	"ImApplication.go/ctrl"
	"html/template"
	"log"
	"net/http"
)

func main() {
	//注册动态模板
	registerTemplates()
	//加载静态文件
	http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("asset"))))
	http.Handle("/mnt/", http.FileServer(http.Dir(".")))
	//初始化数据库

	//用户相关
	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/user/register", ctrl.UserRegister)
	//好友相关
	http.HandleFunc("/contact/loadfriend", ctrl.LoadFriend)
	http.HandleFunc("/contact/addfriend", ctrl.Addfriend)
	//群聊相关
	http.HandleFunc("/contact/loadcommunity", ctrl.LoadCommunity)
	http.HandleFunc("/contact/joincommunity", ctrl.JoinCommunity)
	http.HandleFunc("/contact/createcommunity", ctrl.CreateCommunity)
	http.HandleFunc("/chat", ctrl.Chat)
	http.HandleFunc("/attach/upload", ctrl.Upload)
	http.ListenAndServe(":9001", nil)
}

// 注册模板
func registerTemplates() {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplname := v.Name()
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplname, nil)
		})
	}
}
