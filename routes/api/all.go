package api

import (
    "github.com/gin-gonic/gin"
    "github.com/mix-go/mix-api-skeleton/controllers"
)

var RouteDefinitions []func(router *gin.Engine)

func init() {
    RouteDefinitions = append(RouteDefinitions, func(router *gin.Engine) {
        router.GET("hello", func(ctx *gin.Context) {
            hello := controllers.HelloController{}
            hello.Index(ctx)
        })
    })
}
