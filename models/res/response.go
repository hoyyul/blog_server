package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 0
	ERROR   = 1
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// encapuslate response
func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(data any, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "Succeeded", c)
}

func OkWithSuccess(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "Succeeded", c)
}

/*func Fail(data any, msg string, c *gin.Context) {
	Result(ERROR, data, msg, c)
}*/

func FailWithMessage(msg string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, msg, c)
}

func FailWithCode(code int, c *gin.Context) {
	if msg, ok := ErrorMap[ErrorCode(code)]; ok { // check error map
		Result(code, map[string]interface{}{}, msg, c)
		return
	}
	Result(SUCCESS, map[string]interface{}{}, "Unknown Error", c)
}
