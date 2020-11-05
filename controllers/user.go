package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/mix-go/api-skeleton/globals"
    "github.com/mix-go/api-skeleton/models"
    "net/http"
    "time"
)

type AddUserController struct {
}

func (t *AddUserController) Index(c *gin.Context) {
    db := globals.DB()
    if err := db.Create(&models.User{
        Name:     c.Request.PostFormValue("name"),
        CreateAt: time.Now(),
    }).Error; err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status":  http.StatusInternalServerError,
            "message": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "status":  http.StatusOK,
        "message": "ok",
    })
}
