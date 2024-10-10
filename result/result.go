// 结构数据定义
package result

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Result结构体
type Result struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// success
func Success(data interface{}, c *gin.Context) {
	if data == nil {
		data = gin.H{}
	}
	result := Result{}
	result.Code = APIcode.Success
	result.Message = APIcode.GetMessage(APIcode.Success)
	result.Data = data
	c.JSON(http.StatusOK, result)
}

//failure

func Failture(code uint, message string, c *gin.Context, err *error) {
	result := Result{}
	result.Code = code
	result.Message = message
	if err != nil && *err != nil { // 确保 err 和 *err 都不为 nil
		result.Data = *err
	} else {
		result.Data = gin.H{}
	}
	if c != nil {
		c.JSON(http.StatusBadRequest, result)
		log.Printf("Error %d: %s, Data: %v", code, message, result.Data)
	} else {
		log.Printf("Error %d: %s, Data: %v", code, message, result.Data)
	}

}
