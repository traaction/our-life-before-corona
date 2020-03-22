package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"

	"github.com/traaction/our-life-before-corona/pkg"
)

func main() {
	router := gin.Default()
	host := flag.String("host", "localhost", "Host of the DB.")
	port := flag.String("port", "5432", "Port of the DB.")
	user := flag.String("user", "demopostgresadmin", "User of the DB.")
	name := flag.String("dbname", "demopostgresdb", "Name of the DB.")
	password := flag.String("password", "demopostgrespwd", "Password of the DB.")
	sslmode := flag.String("sslmode", "disable", "sslmode of the DB.")
	flag.Parse()

	var logger = logrus.New()

	logger.Out = os.Stdout
	logger.SetLevel(logrus.InfoLevel)

	dbcon := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		*host, *port, *user, *name, *password, *sslmode)
	logger.Info(dbcon)

	db, err := gorm.Open("postgres", dbcon)
	if err != nil {
		logger.Error(err)
		panic("failed to connect database")
	}
	defer db.Close()

	a := pkg.Activity{
		DB:     db,
		Logger: logger,
	}

	s := pkg.Sentence{
		DB:     db,
		Logger: logger,
	}

	p := pkg.Place{
		DB:     db,
		Logger: logger,
	}

	stats := pkg.Stats{
		DB:     db,
		Logger: logger,
	}

	d := dbinit{
		DB:     db,
		Logger: logger,
	}

	c := cors.DefaultConfig()
	c.AllowAllOrigins = true
	handlefun := cors.New(c)
	router.Use(handlefun)

	router.GET("/dev/init", d.dbinit)

	router.GET("/activities/:activity", a.List)
	router.GET("/activities", a.ListAll)
	router.POST("/activities", a.Add)

	router.GET("/sentences/:user", s.Get)
	router.POST("/sentences", s.Add)

	router.GET("/places/:place", p.List)
	router.GET("/places", p.ListAll)
	router.POST("/places", p.Add)

	router.GET("/stats", stats.Get)

	router.Run(":8080")
}
