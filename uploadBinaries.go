package smtauto

import(

  "log"
  "time"
  "github.com/radutopala/webdriver"
)

func uploadBtnLoaded(wd webdriver.WebDriver)(bool, error){

    _, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'UPLOAD BINARIES')]");
    if err != nil {
        return false, err
    }

    return true, err
}

func binNotExisted(wd webdriver.WebDriver, binVersion string)(bool) {

    notFound := true

    versionQuery, err := wd.FindElement(webdriver.ByID, "mat-input-1")
    if err != nil {
        log.Fatal( err )
    }
    versionQuery.SendKeys( binVersion )

    searchMatIcon, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='mat-form-field-suffix ng-tns-c11-3 ng-star-inserted']/*[contains(text(), 'search')]")
    if err != nil {
        log.Fatal( err )
    }
    searchMatIcon.Click()

    _, err = wd.FindElement(webdriver.ByXPATH, "//div[@class='query-results ng-star-inserted']/*[contains(text(), '" + binVersion + "')]")
    if err != nil {
        log.Println( "Not found binary ==> ", binVersion )
    }else {
        log.Println("Catch up you")
        notFound = false
    }

    wd.Refresh()

    return notFound
}

func unique(strSlice []string) []string {
    keys := make(map[string]bool)
    list := []string{}
    for _, entry := range strSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

func getNewBinToUpload(wd webdriver.WebDriver) ([]string){

    log.Println(stackConf.LnxStack)

    binSlice := make([]string, 2*len(stackConf.LnxStack))

    count :=0

    //Get new VBIOS
    for  index,_ := range stackConf.LnxStack {

        if notFound := binNotExisted(wd, stackConf.LnxStack[ index].VbiosVersion); notFound == true {
            binSlice[count] = stackConf.LnxStack[index].VbiosVersion
            count ++
        }
    }


    //get new OSDB
    for index, _ := range stackConf.LnxStack {

        if notFound := binNotExisted(wd, stackConf.LnxStack[index].OsdbVersion); notFound == true {
            binSlice[count] = stackConf.LnxStack[index].OsdbVersion
            count ++
        }
    }

    //unique binary slice

    binSlice = append(binSlice[:count])
    binSlice = unique(binSlice)

    log.Println(binSlice)

    return binSlice
//    return stackConf.AsicConf
}

func uploadBIOS(wd webdriver.WebDriver, asicConf AsicConf) {

    log.Println("==> Conf of the ASIC, ", asicConf)

    if err := wd.Get("http://smt.amd.com/#/upload?uploadID="); err != nil {
        log.Fatal( err )
    }
    time.Sleep(5 * time.Second)

    biosRadioBtn, err := wd.FindElement(webdriver.ByID, "mat-radio-3")
    if err != nil {
        log.Fatal( err )
    }
    biosRadioBtn.Click()

    programInput, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='mat-form-field-infix']/input[@id='cont']")
    if err != nil {
        log.Fatal( err )
    }
    programInput.SendKeys("Navi 10")

    fileUploadTab, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'File Upload')]")
    if err != nil {
        log.Fatal( err )
    }
    fileUploadTab.Click()

    clickHereBtn, err := wd.FindElement(webdriver.ByXPATH, "//*[@id='mat-tab-content-1-3']/div/div/div/button")
    if err != nil {
        log.Fatal( err )
    }
    clickHereBtn.Click()

    fileInput, err := wd.FindElement(webdriver.ByXPATH, "//*[@id='mat-tab-content-1-3']/div/div/div/input")
    if err != nil {
        log.Fatal( err )
    }
    fileInput.SendKeys("/opt/shares/Navi10_Stack/WW47/" + asicConf.VbiosFileName )

    versionInput, err := wd.FindElement(webdriver.ByXPATH, "//*[@id='alias']")
    if err != nil {
        log.Fatal( err )
    }
    versionInput.SendKeys( asicConf.VbiosVersion )

    osInput, err := wd.FindElement(webdriver.ByXPATH, "//*[@id='os']")
    if err != nil {
        log.Fatal( err )
    }
    osInput.SendKeys("Linux Ubuntu 18.04 LTS")

    ////*[@id="mat-option-17"]/span
    osListBox, err := wd.FindElement(webdriver.ByXPATH, "//div[@id='mat-autocomplete-1']//span[contains(text(), 'Linux Ubuntu 18.04')]")
    if err != nil {
        log.Fatal( err )
    }
    osListBox.Click()

    ////*[@id="mat-tab-content-0-0"]/div/app-binary-form/div/div[5]/button/span
    uploadBtn, err := wd.FindElement(webdriver.ByXPATH, "//*[@class='submit-button mat-raised-button']")
    if err != nil {
        log.Fatal( err )
    }
    uploadBtn.Click()

}

func UploadBinaries(wd webdriver.WebDriver)( webdriver.WebDriver ){

    if err := wd.WaitWithTimeout(uploadBtnLoaded, 10 * time.Second); err != nil {
        log.Fatal( err )
    }else {
        log.Println("Will upload binaries")
    }

    uploadBtn, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'UPLOAD BINARIES')]")
    if err != nil {
        log.Fatal( err )
    }
    uploadBtn.Click()


    lnxStack := []AsicConf {
        {
            AsicName: "Navi10 Pro-XL",
            StackName: "D18801W1947LN5",
            TargetRelease: "19.50",
            VbiosVersion: "D1880201_102",
            VbiosFileName: "D1880201.102",
        },
        {
            AsicName: "Navi10 XLE",
            StackName: "D18901W1947LN5",
            TargetRelease: "19.50",
            VbiosVersion: "D1890101_066",
            VbiosFileName: "D1890101.066",
            OsdbVersion: "19.50-949708",
        },
    } // getNewBinToUpload(wd)

    for index,_ := range lnxStack {
        uploadBIOS(wd, lnxStack[index])
    }

    //uploadOSDB(wd, binSlice[2:])

    return wd

}
