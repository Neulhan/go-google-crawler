package features

import (
	"github.com/Neulhan/TravelCrawler/configs"
	"github.com/Neulhan/TravelCrawler/controler"
	"github.com/Neulhan/TravelCrawler/models"
	"github.com/tebeka/selenium"
	"log"
)

func GetDataFromCity(city string) []models.Data {
	dataList := make([]models.Data, 0)
	for _, key := range configs.KeyWordList {
		dataList = append(dataList, ExtractData(city, key)...)
	}
	return dataList
}

func ExtractData(city string, keyWord string) []models.Data {
	var err error

	err = controler.Session.Url(
		"https://www.google.com/search??tbs=lf:1,lf_ui:9&tbm=lcl&q=" +
			city + "+" + keyWord + "&oq=" + city + "+" + keyWord,
	)
	if err != nil {
		log.Println(err)
	}

	// 셀레니움 관련 제어 부분
	links, _ := controler.Session.FindElements(selenium.ByCSSSelector, "a[role=\"link\"]")

	var dataList []models.Data

	for _, link := range links {
		data := new(models.Data)

		var imgSrc string
		imgEl, err := link.FindElement(selenium.ByCSSSelector, "img.tLipRb")
		if err == nil {
			imgSrc, _ = imgEl.GetAttribute("src")
		}
		data.Img = imgSrc

		var name string
		nameEl, err := link.FindElement(selenium.ByCSSSelector, "div.dbg0pd")
		if err == nil {
			name, _ = nameEl.Text()
		}
		data.Name = name

		var rating string
		ratingEl, err := link.FindElement(selenium.ByCSSSelector, "span.BTtC6e")
		if err == nil {
			rating, _ = ratingEl.Text()
		}
		data.Rating = rating

		var rateNum string
		rateNumEl, err := link.FindElement(selenium.ByCSSSelector, "span.BTtC6e + g-review-stars + span")
		if err == nil {
			rateNum, _ = rateNumEl.Text()
		}
		data.RateNum = rateNum

		data.HashTags = []string{city, keyWord}

		dataList = append(dataList, *data)
	}

	return dataList
}
