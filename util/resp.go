package util

import (
	"encoding/json"
	"log"
	"net/http"
)

// ResponseData is the response data structure
type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func RespJson(writer http.ResponseWriter, code int, msg string, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	resp := ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	// struct to json
	ret, err := json.Marshal(resp)
	if err != nil {
		log.Panicln(err.Error())
	}
	// return json
	writer.Write(ret)
}
