package controler

import (
	"github.com/fedesog/webdriver"
	"log"
)

var ChromeDriver *webdriver.ChromeDriver
var Session *webdriver.Session

func RunChromeDriver() *webdriver.ChromeDriver {
	ChromeDriver = webdriver.NewChromeDriver("./chromedriver")

	err := ChromeDriver.Start()
	if err != nil {
		log.Println(err)
	}

	return ChromeDriver
}

func GetNewSession() (*webdriver.Session, error) {
	desired := webdriver.Capabilities{}
	required := webdriver.Capabilities{}
	var err error
	Session, err = ChromeDriver.NewSession(desired, required)
	return Session, err
}
