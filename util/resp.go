package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseData struct {
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Code int         `json:"code"`
}

// 返回成功的信息
func ResponseOk(w http.ResponseWriter, msg string, data interface{}) {
	WriteJSONResponse(w, msg, data, 0)
}

// 返回失败的信息
func ResponseFail(w http.ResponseWriter, msg string) {
	WriteJSONResponse(w, msg, nil, -1)
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
