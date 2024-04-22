package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/k2gutierrez/gosql/database/dbtools"
)

type Configuration struct {
	DriverName     string `json:"driverName"`
	DataSourceName string `json:"dataSourceName"`
}

func main() {

	file, err := os.Open("configuration/config.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	conf := new(Configuration)

	json.NewDecoder(file).Decode(conf)

	dbtools.DBInitializer(conf.DriverName, conf.DataSourceName)

	students := dbtools.SelectAllStudents()

	for _, student := range students {

		fmt.Println("ID: ", student.ID, "\tName: ", student.Name, "\tAge: ", student.Age)

	}

}
