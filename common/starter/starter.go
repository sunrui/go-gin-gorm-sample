package starter

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"medium-server-go/common/errno"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Starter struct {
	engine *gin.Engine
}

func New() *Starter {
	engine := gin.Default()

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, errno.NotFound.WithKeyPair("uri", ctx.Request.URL.RequestURI()))
	})

	engine.HandleMethodNotAllowed = true
	engine.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, errno.MethodNotAllowed.WithKeyPair("uri", ctx.Request.URL.RequestURI()))
	})

	engine.Use(rateLimitMiddleware(time.Second, 10000, 10000))

	return &Starter{
		engine: engine,
	}
}

type RouterPath struct {
	HttpMethod   string
	RelativePath string
	HandlerFunc  gin.HandlerFunc
}

type Router struct {
	GroupName   string
	NeedAuth    bool
	RouterPaths []RouterPath
}

func (starter *Starter) RegisterRouter(router Router) {
	groupRouter := starter.engine.Group(router.GroupName)

	if router.NeedAuth {
		groupRouter.Use(authMiddleware)
	}

	for _, routerPath := range router.RouterPaths {
		switch routerPath.HttpMethod {
		case "GET":
			groupRouter.GET(routerPath.RelativePath, catchHandler(routerPath.HandlerFunc))
		case "POST":
			groupRouter.POST(routerPath.RelativePath, catchHandler(routerPath.HandlerFunc))
		case "PUT":
			groupRouter.PUT(routerPath.RelativePath, catchHandler(routerPath.HandlerFunc))
		case "DELETE":
			groupRouter.DELETE(routerPath.RelativePath, catchHandler(routerPath.HandlerFunc))
		default:
			panic(errno.InternalError.WithKeyPair("httpMethod", routerPath.HttpMethod))
		}
	}
}

func (starter *Starter) Run(port int) {
	err := starter.engine.Run(":" + strconv.Itoa(port))
	if err != nil {
		panic(errno.InternalError.WithData(err.Error()))
	}
}

func authMiddleware(ctx *gin.Context) {
	fmt.Println("authMiddleware = ")
}

func rateLimitMiddleware(fillInterval time.Duration, cap, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(ctx *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			ctx.JSON(http.StatusBadRequest, errno.RateLimit)
			return
		}

		ctx.Next()
	}
}

func catchHandler(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

				ctx.JSON(http.StatusBadRequest, errNo)
			}
		}()

		handlerFunc(ctx)
	}
}
