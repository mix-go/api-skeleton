package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/mix-go/api-skeleton/controllers"
    "github.com/mix-go/api-skeleton/middleware"
)

func RouteDefinitionCallbacks() (callbacks []func(router *gin.Engine)) {
    callbacks = append(callbacks, func(router *gin.Engine) {
        router.GET("hello",
            middleware.CorsMiddleware(),
            func(ctx *gin.Context) {
                hello := controllers.HelloController{}
                hello.Index(ctx)
            },
        )

        router.POST("users/add",
            middleware.AuthMiddleware(),
            func(ctx *gin.Context) {
                hello := controllers.UserController{}
                hello.Add(ctx)
            },
        )

        router.POST("auth", func(ctx *gin.Context) {
            auth := controllers.AuthController{}
            auth.Index(ctx)
        })
    })
    return
}
