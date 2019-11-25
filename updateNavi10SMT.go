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

func vbiosUpdated(wd webdriver.WebDriver)(bool, error){

    _, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'Status: Uploaded')]")
    if err != nil {
        return false, err
    }

    return true, nil
}

func uploadVbiosBinary(wd webdriver.WebDriver) (error){

    //click Binaries
    binTab, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'Binaries')]")
    if err != nil {
        log.Fatal( err )
    }
    binTab.Click()

    time.Sleep( 5 * time.Second)

    //check if updated
    if err = wd.WaitWithTimeout(vbiosUpdated, 10 * time.Second); err == nil {
        log.Println("VBIOS has already been updated - " + "BIOS____XXXX")
        return nil
    }

    //*****update vbios since it is not updated
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

    return nil
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

func testReportUploaded(wd webdriver.WebDriver)(bool, error) {

    _, err := wd.FindElement(webdriver.ByXPATH, "//span[contains(text(), 'UPLOAD NEW REPORT')]")
    if err != nil {
        return false, err
    }

    return true, nil
}

func uploadTestReport(wd webdriver.WebDriver)(error) {

    //goto Test Reports tab
    testReportTab, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'Test Reports')]")
    if err != nil {
        log.Fatal( err )
    }
    testReportTab.Click()

    //check test report uploaded or not
    if err = wd.WaitWithTimeout(testReportUploaded, 5 * time.Second); err == nil {
        log.Println("Test Report has alreayd been uploaded yet ..., SKIP" )
        return err
    }

    //add report click
    addReportBtn, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'ADD REPORT')]")
    if err != nil {
        log.Fatal( err )
    }
    addReportBtn.Click()

    //select file
    fileInput, err := wd.FindElement(webdriver.ByXPATH, "//*[@id='mat-dialog-0']//div[@class='upload-file']/input[@type='file']")
    if err != nil {
      log.Fatal( err )
    }
    fileInput.SendKeys(stackConf.StackPath + "/" + stackConf.Version + "/" + stackConf.TestReport)

    //upload click
    uploadBtn, err := wd.FindElement(webdriver.ByXPATH, "//button[@class='save-button mat-button ng-star-inserted']/span[contains(text(), 'UPLOAD')]")
    if err != nil {
        log.Fatal( err )
    }
    uploadBtn.Click()

    //check results
    //TBD

    return nil

}

func UpdateNavi10SMT(wd webdriver.WebDriver) {

    log.Println("Go to stacks")

    gotoSpecNavi10Stack( wd )

    uploadVbiosBinary( wd )

    uploadTestReport( wd )
}
