package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"gitlab/wirvsvirus/our-life-before-corona/models"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type dbinit struct {
	DB *gorm.DB
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
			fmt.Println(fmt.Sprintf("Adding country: %s", line[1]))
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
			d.DB.Create(&models.Place{Name: line[0], Type: models.City})
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
			fmt.Println(fmt.Sprintf("Adding activity: %s", lineStripped))
			d.DB.Create(&models.Activity{Name: lineStripped})
		}
	}
}

func (d dbinit) dbinit(c *gin.Context) {
	d.DB.DropTable(&models.Activity{})
	d.DB.DropTable(&models.Place{})
	d.DB.DropTable(&models.Sentence{})
	d.readActivities()
	d.readCountries()
	d.readCities()

	c.JSON(http.StatusOK, nil)
}
