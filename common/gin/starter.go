package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"medium-server-go/common/errno"
	"net/http"
	"strconv"
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
				context.JSON(http.StatusBadRequest, errno.InternalError.WithData(err))
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
