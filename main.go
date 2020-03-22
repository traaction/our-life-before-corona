package main

import (
	"flag"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"gitlab/wirvsvirus/our-life-before-corona/models"
	"gitlab/wirvsvirus/our-life-before-corona/pkg"
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

	dbcon := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", *host, *port, *user, *name, *password, *sslmode)
	fmt.Println(dbcon)

	db, err := gorm.Open("postgres", dbcon)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()
	// Migrate the schema
	db.AutoMigrate(&models.Place{})
	db.AutoMigrate(&models.Activity{})
	db.AutoMigrate(&models.Sentence{})

	db.Model(&models.Sentence{}).AddForeignKey("place_id", "places(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Sentence{}).AddForeignKey("activity_id", "activities(id)", "RESTRICT", "RESTRICT")

	a := pkg.Activity{
		DB: db,
	}

	s := pkg.Sentence{
		DB: db,
	}

	p := pkg.Place{
		DB: db,
	}

	d := dbinit{
		DB: db,
	}

	c := cors.Config{
		AllowOrigins: []string{
			"*",
		},
	}
	handlefun := cors.New(c)
	router.Use(handlefun)

	router.GET("/dev/init", d.dbinit)

	router.GET("/activities/:activity", a.List)
	router.GET("/activities", a.ListAll)
	router.POST("/activities", a.Add)

	router.POST("/sentences", s.Add)

	router.GET("/places/:place", p.List)
	router.GET("/places", p.ListAll)
	router.POST("/places", p.Add)

	router.Run(":8080")
}
