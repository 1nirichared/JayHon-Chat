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
	result := Result{}
	result.Code = APIcode.Success
	result.Message = APIcode.GetMessage(APIcode.Success)

	if data == nil {
		// 如果没有传递数据，默认返回一个空的 gin.H
		result.Data = gin.H{}
	} else if msg, ok := data.(string); ok {
		// 如果传递的是字符串，则将其作为 message 返回
		result.Data = gin.H{"message": msg, "success": true}
	} else if res, ok := data.(gin.H); ok {
		// 如果传递的是 gin.H 类型，则直接使用
		res["success"] = true
		result.Data = res
	} else {
		// 否则保持原有数据并返回 success 标志
		result.Data = gin.H{"data": data, "success": true}
	}

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
