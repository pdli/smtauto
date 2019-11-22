package smtauto

import (
//    "fmt"
    "flag"
    "log"
    "time"
    "github.com/radutopala/webdriver"
    "github.com/radutopala/webdriver/chrome"
)

func newChromeDriver() (webdriver.WebDriver ) {

    chromeBinary := flag.String("chrome_binary", "/usr/bin/chromium-browser", "The name of the Chrome binary or the path to it. If name is not an exact path, the PATH will be searched.")
    chrCaps := chrome.Capabilities{
        Path: *chromeBinary,
        Args: []string{
            //"--headless",
            //"--disable-gpu",
            //"--no-sandbox",
            //"--window-size=800,600",
        },
    }
    caps := webdriver.Capabilities{
        "browserName": "chrome",
        "path": chromeBinary,
    }
    caps.AddChrome( chrCaps )

    //_, err := webdriver.NewService()
    //if err != nil {
    //    log.Fatal( err )
    //}
    //defer s.Stop()

    wd, err := webdriver.NewRemote(caps, "")
    if err != nil {
        log.Fatal( err )
    }
    //defer wd.Quit()

    if err := wd.SetImplicitWaitTimeout( 30 * time.Second ); err != nil {
        log.Fatal( err )
    }

    return wd
}

func mainPageLoaded(wd webdriver.WebDriver) (bool, error){

    _, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'ACTIVE PROGRAMS LIST')]")
    if err != nil{
        return false, err
    }

    return true, nil
}

func LogonSMT( smtUrl string) (wd webdriver.WebDriver){

    wd = newChromeDriver()

    if err := wd.Get( smtUrl ); err != nil {
        log.Fatal( err )
    }

    if err := wd.WaitWithTimeout(mainPageLoaded, 90 * time.Second); err != nil {
        log.Fatal( err )
    } else {
        log.Println("Log in successfully")
    }

    return wd
}

