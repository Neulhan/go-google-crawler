package main

import (
	"encoding/json"
	"github.com/Neulhan/TravelCrawler/configs"
	"github.com/Neulhan/TravelCrawler/controler"
	"github.com/Neulhan/TravelCrawler/features"
	"github.com/Neulhan/TravelCrawler/models"
	"io/ioutil"
	"log"
)

func main() {
	chromeDriver := controler.RunChromeDriver()
	session, _ := controler.GetNewSession()
	defer chromeDriver.Stop()
	defer session.Delete()

	dataList := make([]models.Data, 0)
	log.Println(configs.CityList)

	for _, city := range configs.CityList {
		dataList = append(dataList, features.GetDataFromCity(city)...)
	}

	file, _ := json.MarshalIndent(dataList, "", " ")
	_ = ioutil.WriteFile("files/data.json", file, 0644)
}
