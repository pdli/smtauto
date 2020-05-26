package smtauto

import (
	"log"
	"time"

	"github.com/radutopala/webdriver"
)

func newChromeDriver() webdriver.WebDriver {

	caps := webdriver.Capabilities{
		"browserName": "chrome",
		//		"path":        chromeBinary,
	}
	caps.AddChrome(chrCaps)

	wd, err := webdriver.NewRemote(caps, "")
	if err != nil {
		log.Fatal(err)
	}
	//defer wd.Quit()

	if err := wd.SetImplicitWaitTimeout(30 * time.Second); err != nil {
		log.Fatal(err)
	}

	return wd
}

func mainPageLoaded(wd webdriver.WebDriver) (bool, error) {

	_, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'CLONE')]")
	if err != nil {
		return false, err
	}

	return true, nil
}

//NewWebService return a chromedriver service to be launched
func NewWebService() *webdriver.Service {

	service, err := webdriver.NewService()
	if err != nil {
		log.Fatal(err)
	}

	return service
}

//LogonSMT will visit SMT website automatically
func LogonSMT(smtURL string) (wd webdriver.WebDriver) {

	wd = newChromeDriver()

	if err := wd.Get(smtURL); err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	//TODO: automatically login SMT

	if err := wd.WaitWithTimeout(mainPageLoaded, 90*time.Second); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Log in successfully")
	}

	return wd
}
