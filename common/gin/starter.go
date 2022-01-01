package gin

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"medium-server-go/common/errno"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

type Starter struct {
	engine *gin.Engine
}

func New() *Starter {
	engine := gin.Default()

	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusBadRequest, errno.NotFound.WithKeyPair("uri", context.Request.URL.RequestURI()))
	})

	engine.HandleMethodNotAllowed = true
	engine.NoMethod(func(context *gin.Context) {
		context.JSON(http.StatusBadRequest, errno.MethodNotAllowed.WithKeyPair("uri", context.Request.URL.RequestURI()))
	})

	engine.Use(func(context *gin.Context) {
		fmt.Println("use handler")
	})

	return &Starter{
		engine: engine,
	}
}

func cacheHandler(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				funcName, file, line, _ := runtime.Caller(3)
				dataMap := make(map[string]interface{})
				dataMap["reason"] = err

				stack := make(map[string]string)
				funcForPCName := runtime.FuncForPC(funcName).Name()
				funcShortName := funcForPCName[strings.LastIndex(funcForPCName, "/")+1:]
				stack["function"] = funcShortName
				file += fmt.Sprintf(":%d", line)
				stack["file"] = file

				dataMap["stack"] = stack
				errNo := errno.InternalError.WithData(dataMap)

				marshal, _ := json.MarshalIndent(errNo, "", "    ")
				fmt.Println(string(marshal))

				file = file[strings.LastIndex(file, "/controller"):]
				stack["file"] = file

				context.JSON(http.StatusBadRequest, errNo)
			}
		}()

		handlerFunc(context)
	}
}

func (server *Starter) RegisterHandler(httpMethod string, relativePath string, handlerFunc gin.HandlerFunc) {
	switch httpMethod {
	case "GET":
		server.engine.GET(relativePath, cacheHandler(handlerFunc))
	case "POST":
		server.engine.POST(relativePath, cacheHandler(handlerFunc))
	case "PUT":
		server.engine.PUT(relativePath, cacheHandler(handlerFunc))
	case "DELETE":
		server.engine.DELETE(relativePath, cacheHandler(handlerFunc))
	default:
		panic(errno.InternalError.WithKeyPair("httpMethod", httpMethod))
	}
}

func (server *Starter) Run(port int) {
	err := server.engine.Run(":" + strconv.Itoa(port))
	if err != nil {
		panic(errno.InternalError.WithData(err))
	}
}
