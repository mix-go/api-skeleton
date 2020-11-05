package api

import (
    "github.com/gin-gonic/gin"
    "github.com/mix-go/api-skeleton/controllers"
    "github.com/mix-go/api-skeleton/middleware"
)

var RouteDefinitionCallbacks []func(router *gin.Engine)

func init() {
    RouteDefinitionCallbacks = append(RouteDefinitionCallbacks, func(router *gin.Engine) {
        router.GET("hello",
            middleware.CorsMiddleware(),
            func(ctx *gin.Context) {
                hello := controllers.HelloController{}
                hello.Index(ctx)
            },
        )

        router.POST("users/add",
            middleware.CorsMiddleware(),
            func(ctx *gin.Context) {
                hello := controllers.AddUserController{}
                hello.Index(ctx)
            },
        )

        router.POST("auth", func(ctx *gin.Context) {
            auth := controllers.AuthController{}
            auth.Index(ctx)
        })
    })
}
