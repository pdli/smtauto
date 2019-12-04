package smtauto

import (
	"log"
	"os"
	"time"

	"github.com/radutopala/webdriver"
	"github.com/radutopala/webdriver/chrome"
)

func newChromeDriver() webdriver.WebDriver {

	homeDir := os.Getenv("HOME")

	//    chromeBinary := flag.String("chrome_binary", "/usr/bin/chromium-browser", "The name of the Chrome binary or the path to it. If name is not an exact path, the PATH will be searched.")
	//	chromeBinary := flag.String("chrome_binary", "C:/Program Files (x86)/Google/Chrome/Application", "")
	chrCaps := chrome.Capabilities{
		//Path: *chromeBinary,
		Args: []string{
			"--headless",
			"--disable-gpu",
			"--no-sandbox",
			"--log-level=1",
			//"--window-size=800,600",
			"--remote-debugging-port=9222",
			"--disable-dev-shm-usage",
			"--user-data-dir=" + homeDir + "/.config/chromium/",
		},
	}
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

	_, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'ACTIVE PROGRAMS LIST')]")
	if err != nil {
		return false, err
	}

	return true, nil
}

func NewWebService() *webdriver.Service {

	service, err := webdriver.NewService()
	if err != nil {
		log.Fatal(err)
	}

	return service
}

func LogonSMT(smtUrl string) (wd webdriver.WebDriver) {

	wd = newChromeDriver()

	if err := wd.Get(smtUrl); err != nil {
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
