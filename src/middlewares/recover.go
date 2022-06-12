package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
	"tiktok/src/errno"
)

// RecoverMiddleware 全局异常捕捉插件
func RecoverMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//若是非自定义异常，将其进行转换为自定异常统一返回

			err, ok := r.(errno.ErrNo)
			if !ok {
				r = errno.ConvertErr(err)
			}
			//封装通用JSON返回
			c.JSON(http.StatusOK, gin.H{
				"status_code": err.ErrCode,
				"status_msg":  err.ErrMsg,
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续的插件及代码执行
	c.Next()
}

//错误转字符串
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

