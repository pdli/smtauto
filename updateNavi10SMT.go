package smtauto

import(

    "log"
    "time"
    "github.com/radutopala/webdriver"
)

func stackLoaded(wd webdriver.WebDriver) (bool, error) {

    //NV10-D18801W1947LN5
    _, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='info-holder']/*[contains(text(), 'NV10-D18801W1947LN5')]")
    if err != nil {
        return false, err
    }

    return true, nil
}

func binaryLinked(wd webdriver.WebDriver) (bool, error) {

    _, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'Binary linked successfully')]")
    if err != nil {
        return false, err
    }

    return true, nil
}

func uploadVbiosBinary(wd webdriver.WebDriver) {

    //click Binaries
    binTab, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'Binaries')]")
    if err != nil {
        log.Fatal( err )
    }
    binTab.Click()

    time.Sleep( 5 * time.Second)

    //select Action
    selectActBtn, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'Select Action')]")
    if err != nil {
        log.Fatal( err )
    }
    selectActBtn.Click()

    //link binary
    linkBinBtn, err := wd.FindElement(webdriver.ByXPATH, "//div[@id='cdk-overlay-0']//*[contains(text(), 'Link Binary')]")
    if err != nil {
        log.Fatal( err )
    }
    linkBinBtn.Click()

    versionSearchInput, err := wd.FindElement(webdriver.ByID, "mat-input-3")
    if err != nil {
        log.Fatal( err )
    }
    versionSearchInput.Clear()
    versionSearchInput.SendKeys("D1880201_102")

    searchBtn, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='mat-form-field-suffix ng-tns-c11-29 ng-star-inserted']/*[contains(text(), 'search')]")
    if err != nil {
        log.Fatal( err )
    }
    searchBtn.Click()

    resultSpan, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='query-results ng-star-inserted']/*[contains(text(), 'D1880201_102')]")
    if err != nil {
        log.Fatal( err )
    }
    resultSpan.Click()

    linkBinaryBtn, err := wd.FindElement(webdriver.ByXPATH, "//button/*[contains(text(), 'LINK BINARY')]")
    if err != nil {
        log.Fatal( err )
    }
    linkBinaryBtn.Click()

    //Binary linked successfully
    if err = wd.WaitWithTimeout(binaryLinked, 20 * time.Second); err != nil {
        log.Fatal("Binary linked failed...", err)
    } else {
        log.Println("Binary linked successful")
    }
}

func gotoSpecNavi10Stack(wd webdriver.WebDriver) {

    viewStackBtn, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'VIEW STACKS')]")
    if err != nil {
        log.Fatal( err )
    }
    viewStackBtn.Click()

    time.Sleep( 5 * time.Second ) //wait for stacks page loading

    lnxStackSpan, err := wd.FindElement(webdriver.ByXPATH, "//span[@class='progress-text']/*[contains(text(), 'D18801W1947LN5')]")
    if err != nil {
        log.Fatal( err )
    }
    lnxStackSpan.Click()

    //wait for loading
    if err = wd.WaitWithTimeout(stackLoaded, 10 * time.Second); err != nil {
        log.Fatal( "Time out for stack loading ==>" , err)
    }

}

func UpdateNavi10SMT(wd webdriver.WebDriver) {

    log.Println("Go to stacks")

    gotoSpecNavi10Stack( wd )

    uploadVbiosBinary( wd )
}
