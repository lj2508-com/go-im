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
	//初始化数据库

	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/user/register", ctrl.UserRegister)
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
