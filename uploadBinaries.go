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

        if notFound := binNotExisted(wd, stackConf.LnxStack[ index].VBIOS); notFound == true {
            binSlice[count] = stackConf.LnxStack[index].VBIOS
            count ++
        }
    }


    //get new OSDB
    for index, _ := range stackConf.LnxStack {

        if notFound := binNotExisted(wd, stackConf.LnxStack[index].OSDB); notFound == true {
            binSlice[count] = stackConf.LnxStack[index].OSDB
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

func uploadBIOS(wd webdriver.WebDriver, vbiosSlice []string) {

    log.Println("==> New VBIOS Lisst, \n", vbiosSlice)

    if err := wd.Get("http://smt.amd.com/#/upload?uploadID="); err != nil {
        log.Fatal( err )
    }

    biosRadioBtn, err := wd.FindElement(webdriver.ByID, "mat-radio-3-input")
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

    clickHereBtn, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'CLICKE HERE TO SELECT A FILE')]")
    if err != nil {
        log.Fatal( err )
    }
    clickHereBtn.Click()

    fileInput, err := wd.FindElement(webdriver.ByXPATH, "//*[@id='mat-tab-content-1-3']/div/div/div/input")
    if err != nil {
        log.Fatal( err )
    }
    fileInput.SendKeys("/opt/shares/Navi10_Stack/WW47/D1880201.102")

    versionInput, err := wd.FindElement(webdriver.ByXPATH, "//*[@id='alias']")
    if err != nil {
        log.Fatal( err )
    }
    versionInput.SendKeys("D1880201_102")

    osInput, err := wd.FindElement(webdriver.ByXPATH, "//*[@id='os']")
    if err != nil {
        log.Fatal( err )
    }
    osInput.SendKeys("Linux Ubuntu 18.04 LTS")

    uploadBtn, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'UPLOAD')]")
    if err != nil {
        log.Fatal( err )
    }
    uploadBtn.Clear()
    //uploadBtn.Click()

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


    binSlice := []string{
        "D1880201_102",
        "D1890101_066",
        "19.50-949708",
        "19.40-948413",
    } // getNewBinToUpload(wd)

    uploadBIOS(wd, binSlice[:2])

    //uploadOSDB(wd, binSlice[2:])

    return wd

}
