package rsp

import (
	"github.com/gin-gonic/gin"
	"goadmin/common/log"

	//"encoding/json"
	"net/http"
)

type ResultRsp struct {
	code int         `json:"code"`
	msg  string      `json:"success"`
	data interface{} `json:"data"`
}

// 正确状态处理
func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":       200,
		"msg":        msg,
		"data":       data,
		"request_id": log.GetNewUid(),
	})
}

// 错误状态处理
func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest,
		gin.H{
			"error":      msg,
			"code":       -1,
			"request_id": log.GetNewUid(),
		})
}
