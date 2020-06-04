package smtauto

import (
	"log"
	"os"
	"time"

	"github.com/radutopala/webdriver"
)

func uploadBtnLoaded(wd webdriver.WebDriver) (bool, error) {

	_, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'UPLOAD BINARIES')]")
	if err != nil {
		return false, err
	}

	return true, err
}

func binNotExisted(wd webdriver.WebDriver, binVersion string) bool {

	notFound := true

	if err := wd.Get("http://smt.amd.com/#/upload?uploadID="); err != nil {
		log.Fatal(err)
	}

	versionQuery, err := wd.FindElement(webdriver.ByID, "mat-input-1")
	if err != nil {
		log.Fatal(err)
	}
	versionQuery.Clear()
	versionQuery.SendKeys(binVersion)

	searchMatIcon, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='mat-form-field-suffix ng-tns-c11-3 ng-star-inserted']/*[contains(text(), 'search')]")
	if err != nil {
		log.Fatal(err)
	}
	searchMatIcon.Click()

	_, err = wd.FindElement(webdriver.ByXPATH, "//div[@class='query-results ng-star-inserted']/*[contains(text(), '"+binVersion+"')]")
	if err != nil {
		log.Println("Not found binary ==> ", binVersion)
	} else {
		log.Println("Catch up you ==> ", binVersion)
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

func getNewBinToUpload(wd webdriver.WebDriver) []string {

	log.Println(stackConf.LnxStack)

	binSlice := make([]string, 2*len(stackConf.LnxStack))

	count := 0

	//Get new VBIOS
	for index := range stackConf.LnxStack {

		if notFound := binNotExisted(wd, stackConf.LnxStack[index].VbiosVersion); notFound == true {
			binSlice[count] = stackConf.LnxStack[index].VbiosVersion
			count++
		}
	}

	//get new OSDB
	for index := range stackConf.LnxStack {

		if notFound := binNotExisted(wd, stackConf.LnxStack[index].OsdbVersion); notFound == true {
			binSlice[count] = stackConf.LnxStack[index].OsdbVersion
			count++
		}
	}

	//unique binary slice

	binSlice = append(binSlice[:count])
	binSlice = unique(binSlice)

	log.Println(binSlice)

	return binSlice
	//    return stackConf.AsicConf
}

func uploadOSDB(wd webdriver.WebDriver, asicConf AsicConf) {

	log.Println("==> To upload OSDB, ", asicConf)

	if err := wd.Get("http://smt.amd.com/#/upload?uploadID="); err != nil {
		log.Fatal(err)
	}

	//refresh webpage for loop
	wd.Refresh()

	//Click Software in upload type
	osdbRadioBtn, err := wd.FindElement(webdriver.ByID, "mat-radio-2")
	if err != nil {
		log.Fatal(err)
	}
	osdbRadioBtn.Click()

	//SW Name
	swNameInput, err := wd.FindElement(webdriver.ByID, "sw")
	if err != nil {
		log.Fatal(err)
	}
	swNameInput.Clear()
	swNameInput.SendKeys("AMD GPU DRIVER")

	swListBox, err := wd.FindElement(webdriver.ByXPATH, "//div[@id='mat-autocomplete-0']//span[contains(text(), 'AMD GPU DRIVER')]")
	if err != nil {
		log.Fatal(err)
	}
	swListBox.Click()

	httpLinkInput, err := wd.FindElement(webdriver.ByID, "link")
	if err != nil {
		log.Fatal(err)
	}
	httpLinkInput.Clear()
	//URL - http://lnx-jfrog/artifactory/osibuild-packages-cache/949708/hybrid_rel.u1804_64/amdgpu-pro-19.50-949708-ubuntu-18.04.tar.xz
	ubuntuLink := "http://lnx-jfrog/artifactory/osibuild-packages-cache/" +
		asicConf.OsdbID +
		"/hybrid_rel." + asicConf.DistroShortName + "/amdgpu-pro-" +
		asicConf.OsdbVersion +
		"-" + asicConf.DistroName + ".tar.xz"
	httpLinkInput.SendKeys(ubuntuLink)

	versionInput, err := wd.FindElement(webdriver.ByXPATH, "//*[@id='alias']")
	if err != nil {
		log.Fatal(err)
	}
	versionInput.Clear()
	versionInput.SendKeys(asicConf.OsdbVersion)

	osInput, err := wd.FindElement(webdriver.ByXPATH, "//*[@id='os']")
	if err != nil {
		log.Fatal(err)
	}
	osInput.Clear()
	osInput.SendKeys("Linux")

	osListBox, err := wd.FindElement(webdriver.ByXPATH, "//div[@id='mat-autocomplete-1']//span[contains(text(), 'Linux')]")
	if err != nil {
		log.Fatal(err)
	}
	osListBox.Click()
	os.Exit(1)
	uploadBtn, err := wd.FindElement(webdriver.ByXPATH, "//*[@class='submit-button mat-raised-button']")
	if err != nil {
		log.Fatal(err)
	}
	uploadBtn.Click()

}

func uploadBIOS(wd webdriver.WebDriver, asicConf AsicConf) {

	log.Println("==> To upload VBIOS, ", asicConf)

	if err := wd.Get("http://smt.amd.com/#/upload?uploadID="); err != nil {
		log.Fatal(err)
	}

	//refresh webpage for loop
	wd.Refresh()

	//Click BIOS for upload Type
	biosRadioBtn, err := wd.FindElement(webdriver.ByID, "mat-radio-3")
	if err != nil {
		log.Fatal(err)
	}
	biosRadioBtn.Click()

	//Input Program
	programInput, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='mat-form-field-infix']/input[@id='cont']")
	if err != nil {
		log.Fatal(err)
	}
	programInput.Clear()
	programInput.SendKeys(asicConf.ProgramName)

	//File Upload
	fileUploadTab, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'File Upload')]")
	if err != nil {
		log.Fatal(err)
	}
	fileUploadTab.Click()

	//    clickHereBtn, err := wd.FindElement(webdriver.ByXPATH, "//*[@id='mat-tab-content-1-3']/div/div/div/button")
	//    if err != nil {
	//        log.Fatal( err )
	//    }
	//    clickHereBtn.Click()

	fileInput, err := wd.FindElement(webdriver.ByXPATH, "//*[@id='mat-tab-content-1-3']/div/div/div/input")
	if err != nil {
		log.Fatal(err)
	}
	fileInput.Clear()
	fileInput.SendKeys(stackConf.StackPath + "/" + stackConf.Version + "/" + asicConf.VbiosFileName)

	versionInput, err := wd.FindElement(webdriver.ByXPATH, "//*[@id='alias']")
	if err != nil {
		log.Fatal(err)
	}
	versionInput.Clear()
	versionInput.SendKeys(asicConf.VbiosVersion)

	osInput, err := wd.FindElement(webdriver.ByXPATH, "//*[@id='os']")
	if err != nil {
		log.Fatal(err)
	}
	osInput.Clear()
	osInput.SendKeys("Linux")

	osListBox, err := wd.FindElement(webdriver.ByXPATH, "//div[@id='mat-autocomplete-1']//span[contains(text(), 'Linux')]")
	if err != nil {
		log.Fatal(err)
	}
	osListBox.Click()

	uploadBtn, err := wd.FindElement(webdriver.ByXPATH, "//*[@class='submit-button mat-raised-button']")
	if err != nil {
		log.Fatal(err)
	}
	uploadBtn.Click()

}

//UploadBinaries will upload all listed binaries(VBIOS/OSDB) into SMT website
func UploadBinaries(wd webdriver.WebDriver) {

	log.Println("****** To upload binaries ******")

	/*for index := range stackConf.LnxStack {
		uploadBIOS(wd, stackConf.LnxStack[index])
		time.Sleep(10 * time.Second)
	}*/

	for index := range stackConf.LnxStack {
		uploadOSDB(wd, stackConf.LnxStack[index])
		time.Sleep(20 * time.Second)
	}

	unUploadSlice := getNewBinToUpload(wd)
	if len(unUploadSlice) > 0 {
		log.Println("There are still binaries not uplaoded!!! ====> ", unUploadSlice)
	} else {
		log.Println("Upload binaries successfully! Congratulations!")
	}

}
