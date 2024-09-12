package main

import (
	"database/sql"
	"fmt"
	"go-training/flight"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initDatabase() *sql.DB {
	db, err := sql.Open(viper.GetString("db.driver"), viper.GetString("db.url"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	fmt.Println("Connect to database success")
	return db
}

func main() {
	initConfig()

	r := gin.Default()
	db := initDatabase()
	defer db.Close()

	r.GET("/ping", flight.PingHandler)
	getAllHandler := flight.NewGetAllHandler(db)
	r.GET("/flights", getAllHandler.GetAll)
	getByIDHandler := flight.NewGetByIDHandler(db)
	r.GET("/flights/:id", getByIDHandler.GetByID)
	createHandler := flight.NewCreateHandler(db)
	r.POST("/flights", createHandler.Create)
	updateHandler := flight.NewUpdateHandler(db)
	r.PUT("/flights/:id", updateHandler.Update)
	deleteHandler := flight.NewDeleteHandler(db)
	r.DELETE("/flights/:id", deleteHandler.Delete)

	r.Run()
}
