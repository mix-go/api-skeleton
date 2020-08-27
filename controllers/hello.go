package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type HelloController struct {
}

func (t *HelloController) Index(c *gin.Context) {
    c.String(http.StatusOK, "hello, world!")
}
