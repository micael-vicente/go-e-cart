package main

import (
	"cart/config"
	"cart/pkg/controllers"
	"cart/pkg/db"
	"cart/pkg/models"
	"cart/pkg/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func main() {

	conf := getConfig("dev")

	engine := gin.Default()
	service := services.CartService{Repo: &db.CartRepositoryImpl{Db: getDB(conf.Database)}}
	controller := controllers.CrudController{Service: service}

	engine.GET("/carts", controller.ListCarts)
	engine.GET("/carts/:id", controller.GetCart)
	engine.POST("/carts", controller.CreateCart)
	engine.DELETE("/carts/:id", controller.DeleteCart)
	engine.PUT("/carts/:id", controller.UpdateCart)

	err := engine.Run(":" + strconv.Itoa(conf.Server.Port))

	if err != nil {
		log.Fatal(err)
	}
}

func getConfig(env string) config.Config {
	viper.SetConfigName("config-" + env)
	viper.AddConfigPath("config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	var conf config.Config

	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatal(err)
	}
	return conf
}

func getDB(conf config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		conf.Host, conf.Port, conf.Username, conf.Name, conf.Password)

	dbInst, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if conf.AutoMigrate && err == nil {
		migrate(dbInst)
		return dbInst
	} else {
		log.Fatal(err)
		return nil
	}
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Cart{}, &models.CartItem{})

	if err != nil {
		log.Fatal(err)
	}
}
