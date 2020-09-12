package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-crew/Bolierplate-CRUD-Gingonic/docs"
	"github.com/golang-crew/Bolierplate-CRUD-Gingonic/models"
	"github.com/golang-crew/Bolierplate-CRUD-Gingonic/routers"
	"github.com/spf13/viper"

	_ "github.com/sirupsen/logrus"

	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/swaggerFiles"
)

// @APIVersion 1.0
// @Title Bolierplate CRUD API
// @Description This documentation describes bolierplate CRUD APIs
func main() {
	r := gin.Default()

	conf := &cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "x-user-auth-token"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	conf.AllowOrigins = []string{"*"}
	r.Use(cors.New(*conf))

	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = ioutil.Discard

	viper.SetConfigFile(`./config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}

	err = models.Init()
	if err != nil {
		panic(err)
	}

	routers.Init(r)

	docs.SwaggerInfo.Host = "127.0.0.1:8080"

	r.Run(":8080")
}
