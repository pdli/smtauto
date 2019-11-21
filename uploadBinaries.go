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

    versionQuery, err := wd.FindElement(webdriver.ByID, "mat-input-1")
    if err != nil {
        log.Fatal( err )
    }
    versionQuery.SendKeys("D1880201_101")

    searchMatIcon, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='mat-form-field-suffix ng-tns-c11-3 ng-star-inserted']/*[contains(text(), 'search')]")
    if err != nil {
        log.Fatal( err )
    }
    searchMatIcon.Click()

    _, err = wd.FindElement(webdriver.ByXPATH, "//div[@class='query-results ng-star-inserted']/*[contains(text(), 'D1880201_101')]")
    if err != nil {
        log.Fatal( err )
    }else {
        log.Println("Catch up you")
    }

    return wd

}
