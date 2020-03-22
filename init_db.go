package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/traaction/our-life-before-corona/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type dbinit struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func getCSVReader(file string) *csv.Reader {
	csvFile, _ := os.Open(file)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.LazyQuotes = true
	return reader
}

func getReader(file string) *bufio.Reader {
	rfile, _ := os.Open(file)
	reader := bufio.NewReader(rfile)
	return reader
}

func (d dbinit) readCountries() {
	reader := getCSVReader("database_seed/countries.csv")
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if line[0][0] != '#' {
			d.Logger.Info(fmt.Sprintf("Adding country: %s", line[1]))
			d.DB.Create(&models.Place{Name: line[1], Type: models.Country})
		}

	}
}

func (d dbinit) readCities() {
	reader := getCSVReader("database_seed/cities.csv")
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if line[0][0] != '#' {
			fmt.Println(fmt.Sprintf("Adding city: %s", line[0]))
			d.Logger.Info(&models.Place{Name: line[0], Type: models.City})
		}
	}
}

func (d dbinit) readActivities() {
	reader := getReader("database_seed/activities.md")
	for {
		line, error := reader.ReadString('\n')
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if len(line) > 0 && line[0] != '#' {
			lineStripped := strings.TrimSpace(line)
			if lineStripped != "" {
				d.Logger.Info(fmt.Sprintf("Adding activity: %s", lineStripped))
				d.DB.Create(&models.Activity{Name: lineStripped})
			}
		}
	}
}

func (d dbinit) dbinit(c *gin.Context) {
	if d.DB.HasTable(&models.Activity{}) {
		d.DB.DropTable(&models.Activity{})
	}
	if d.DB.HasTable(&models.Place{}) {
		d.DB.DropTable(&models.Place{})
	}
	if d.DB.HasTable(&models.Sentence{}) {
		d.DB.DropTable(&models.Sentence{})
	}
	if d.DB.HasTable(&models.UserInfo{}) {
		d.DB.DropTable(&models.UserInfo{})
	}

	d.DB.AutoMigrate(&models.Activity{}, &models.Place{}, &models.Sentence{}, &models.UserInfo{})

	d.readActivities()
	d.readCountries()
	d.readCities()

	c.JSON(http.StatusOK, nil)
}
