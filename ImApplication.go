package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type ResponseData struct {
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Code int         `json:"code"`
}

func main() {
	//注册动态模板
	registerTemplates()
	//加载静态文件
	http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("asset"))))

	http.HandleFunc("/user/login", userLogin)
	http.ListenAndServe(":9001", nil)
}

func WriteJSONResponse(w http.ResponseWriter, msg string, data interface{}, code int) {
	response := ResponseData{
		Msg:  msg,
		Data: data,
		Code: code,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json, err := json.Marshal(response)
	if err != nil {
		log.Println(err.Error())
	}
	w.Write(json)
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

func userLogin(writer http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	moblie := request.PostForm.Get("moblie")
	password := request.PostForm.Get("password")

	loginOk := false
	if moblie == "18163750583" && password == "123456" {
		loginOk = true
	}
	if loginOk {
		data := make(map[string]interface{})
		data["id"] = "12"
		data["token"] = "test"
		WriteJSONResponse(writer, "登录成功", data, 0)
	} else {
		WriteJSONResponse(writer, "账号或密码错误", nil, -1)
	}
}
