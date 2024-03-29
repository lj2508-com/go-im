package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseData struct {
	Msg   string      `json:"msg,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Code  int         `json:"code"`
	Rows  interface{} `json:"rows,omitempty"`
	Total interface{} `json:"total,omitempty"`
}

// 返回成功的信息
func ResponseOk(w http.ResponseWriter, msg string, data interface{}) {
	WriteJSONResponse(w, msg, data, 0)
}

// 返回失败的信息
func ResponseFail(w http.ResponseWriter, msg string) {
	WriteJSONResponse(w, msg, nil, -1)
}

// 返回分页
func RespOkList(w http.ResponseWriter, lists interface{}, total interface{}) {
	//分页数目,
	RespList(w, 0, lists, total)
}
func RespList(w http.ResponseWriter, code int, data interface{}, total interface{}) {

	w.Header().Set("Content-Type", "application/json")
	//设置200状态
	w.WriteHeader(http.StatusOK)

	h := ResponseData{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	//将结构体转化成JSOn字符串
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	//输出
	w.Write(ret)
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
