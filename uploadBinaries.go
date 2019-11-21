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

func getNewBinToUpload(wd webdriver.WebDriver) ([]string){

    log.Println(stackConf.LnxStack)

    binSlice := make([]string, len(stackConf.LnxStack))

    count :=0
    for  index,_ := range stackConf.LnxStack {

        if notFound := binNotExisted(wd, stackConf.LnxStack[ index].VBIOS); notFound == true {
            binSlice[count] = stackConf.LnxStack[index].VBIOS
            count ++
        }
    }

    binSlice = append(binSlice[:count])

    log.Println(binSlice)

    return binSlice
//    return stackConf.AsicConf
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


    getNewBinToUpload(wd)

    return wd

}
