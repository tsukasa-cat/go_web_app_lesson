package main

import (
  "github.com/gin-gonic/gin"
  "./Infrastructure"
  // "encoding/json"
  "net/http"
)

type User struct {
  ID int `json:"id"`
  Username string `json:"username",gorm:"type:varchar(255)"`
  LastName string `json:"last_name",gorm:"type:varchar(255)"`
  FirstName string `json:"first_name",gorm:"type:varchar(255)"`
  Email string `json:"email",gorm:"type:varchar(255)"`
  Phone string `json:"phone",gorm:"type:varchar(255)"`
  // Password string `json:"password",gorm:"type:varchar(255)"`
  UserStatus int`json:"user_status",gorm:"type:int(32)"`
}


func main() {

  d := Infrastructure.NewDB()
  db := d.Connection

  router := gin.Default()
  router.LoadHTMLGlob("templates/*.html")

  router.GET("/", func(ctx *gin.Context){
    ctx.HTML(200, "index.html", gin.H{})
  })


  router.GET("/users", func(c *gin.Context) {
    users := []User{}
    db.Find(&users) // 全レコード
    c.JSON(http.StatusOK, users)
  })

  router.GET("/user/:username", func(c *gin.Context) {
    username := c.Param("username")
    user := User{}
    db.Where("username = ?", username).First(&user)
    c.JSON(http.StatusOK, user)
  })

  router.POST("/user", func(c *gin.Context) {
    /** TODO パスワード処理を入れる */

    user := User{
      Username: c.PostForm("username"),
      LastName: c.PostForm("last_name"),
      FirstName: c.PostForm("first_name"),
      Email: c.PostForm("email"),
      Phone: c.PostForm("phone"),
      UserStatus: 0,
    }
    db.Create(&user)
  })

  router.PUT("/user/:username", func(c *gin.Context) { 
    username := c.Param("username")
    user := User{}
    db.Where("username = ?", username).First(&user)
    user.Username = c.PostForm("username")
    user.LastName = c.PostForm("last_name")
    user.FirstName = c.PostForm("first_name")
    user.Email = c.PostForm("email")
    user.phone = c.PostForm("phone")
  })

  router.DELETE("/user/:username", func(c *gin.Context) {
    username := c.Param("username")
    user := User{}
    db.Where("username = ?", username).First(&user)
    db.Delete(&user)
  })


  router.Run()
}