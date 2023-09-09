package main

import (
	"demo/config"
	_ "demo/config"
	"demo/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", listCategory)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func listCategory(c *gin.Context) {
	reply, err := config.RedisConn.Do("PING")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	fmt.Println(reply)
	var categoryList []model.Category
	if err := config.MysqlConn.Find(&categoryList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "data": categoryList})

}
