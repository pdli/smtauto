package smtauto

import (
	"log"
	"time"

	"github.com/radutopala/webdriver"
)

func stackLoaded(wd webdriver.WebDriver) (bool, error) {

	//NV10-D18801W1947LN5
	_, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'Overview')]")
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

func vbiosUpdated(wd webdriver.WebDriver) (bool, error) {

	_, err := wd.FindElement(webdriver.ByXPATH, "//p-panel[1]//*[contains(text(), 'Status: Uploaded')]")
	if err != nil {
		return false, err
	}

	return true, nil
}

func uploadVbiosBinary(wd webdriver.WebDriver, biosVersion string) error {

	//check if updated
	if err := wd.WaitWithTimeout(vbiosUpdated, 1*time.Second); err == nil {
		log.Println("- SKIP - VBIOS has already been updated - ", biosVersion)
		return nil
	}

	//*****update vbios since it is not updated
	//select Action
	selectActBtn, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='row'][1]/div[@class='element-name']/button[@class='mat-raised-button']")
	if err != nil {
		log.Fatal(err)
	}
	selectActBtn.Click()

	//link binary
	linkBinBtn, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='mat-menu-content']/button[1]")
	if err != nil {
		log.Fatal(err)
	}
	linkBinBtn.Click()

	//search vbios - //*[@id="mat-input-5"]
	versionSearchInput, err := wd.FindElement(webdriver.ByXPATH, "//mat-dialog-container//input[@name='query']")
	if err != nil {
		log.Fatal(err)
	}
	versionSearchInput.Clear()
	versionSearchInput.SendKeys(biosVersion)

	searchBtn, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='cdk-overlay-pane']//mat-icon[contains(text(), 'search')]")
	if err != nil {
		log.Fatal(err)
	}
	searchBtn.Click()

	resultSpan, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='query-results ng-star-inserted']/*[contains(text(), '"+biosVersion+"')]")
	if err != nil {
		log.Fatal(err)
	}
	resultSpan.Click()

	linkBinaryBtn, err := wd.FindElement(webdriver.ByXPATH, "//button/*[contains(text(), 'LINK BINARY')]")
	if err != nil {
		log.Fatal(err)
	}
	linkBinaryBtn.Click()

	//Binary linked successfully
	if err = wd.WaitWithTimeout(binaryLinked, 20*time.Second); err != nil {
		log.Fatal("Binary linked failed...", err)
	} else {
		log.Println("Binary linked successful")
	}

	return nil
}

func gotoSpecNavi10Stack(wd webdriver.WebDriver, stackName string) bool {

	var loaded = true //false if not loaded

	log.Println("***** Preparing to update Stack - ", stackName)

	err := wd.Get("http://smt.amd.com/#/view/program/1258")
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second) //wait for stacks page loading

	lnxStackSpan, err := wd.FindElement(webdriver.ByXPATH, "//span[@class='progress-text']/*[contains(text(), '"+stackName+"')]")
	if err != nil {
		log.Println("- SKIP - Cant' find spec Navi10 Stack - ", stackName)
		loaded = false
		return loaded
	}
	lnxStackSpan.Click()

	//wait for loading
	if err = wd.WaitWithTimeout(stackLoaded, 10*time.Second); err != nil {
		log.Fatal("Time out for stack loading ==>", err)
	}

	//load to spec stack successfully
	return loaded
}

func osdbUpdated(wd webdriver.WebDriver) (bool, error) {

	_, err := wd.FindElement(webdriver.ByXPATH, "//p-panel[position()=2]//*[contains(text(), 'Status: Uploaded')]")
	if err != nil {
		return false, err
	}

	return true, nil
}

func uploadOsdbBinary(wd webdriver.WebDriver, osdbVersion string) error {

	//check if updated
	if _, err := wd.WaitWithTimeout(vbiosUpdated, 1*time.Second); err == nil {
		log.Println("- SKIP - OSDB has already been updated - ", osdbVersion)
		return nil
	}

	//*****update vbios since it is not updated
	//select Action
	selectActBtn, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='row'][2]/div[@class='element-name']/button[@class='mat-raised-button']")
	if err != nil {
		log.Fatal(err)
	}
	selectActBtn.Click()

	//link binary
	linkBinBtn, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='mat-menu-content']/button[1]")
	if err != nil {
		log.Fatal(err)
	}
	linkBinBtn.Click()

	//search vbios - //*[@id="mat-input-5"]
	versionSearchInput, err := wd.FindElement(webdriver.ByXPATH, "//mat-dialog-container//input[@name='query']")
	if err != nil {
		log.Fatal(err)
	}
	versionSearchInput.Clear()
	versionSearchInput.SendKeys(osdbVersion)

	searchBtn, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='cdk-overlay-pane']//mat-icon[contains(text(), 'search')]")
	if err != nil {
		log.Fatal(err)
	}
	searchBtn.Click()

	resultSpan, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='query-results ng-star-inserted']/*[contains(text(), '"+osdbVersion+"')]")
	if err != nil {
		log.Fatal(err)
	}
	resultSpan.Click()

	linkBinaryBtn, err := wd.FindElement(webdriver.ByXPATH, "//button/*[contains(text(), 'LINK BINARY')]")
	if err != nil {
		log.Fatal(err)
	}
	linkBinaryBtn.Click()

	//Binary linked successfully
	if err = wd.WaitWithTimeout(binaryLinked, 20*time.Second); err != nil {
		log.Fatal("Binary linked failed...", err)
	} else {
		log.Println("Binary linked successful")
	}

	return nil

}

func binariesTabLoaded(wd webdriver.WebDriver) (bool, error) {

	_, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'DOWNLOAD')]")
	if err != nil {
		return false, err
	}

	return true, nil
}

func uploadBinaries(wd webdriver.WebDriver, asicConf AsicConf) error {

	//goto Binaries tab - click Binaries
	binTab, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'Binaries')]")
	if err != nil {
		log.Fatal(err)
	}
	binTab.Click()

	if err = wd.WaitWithTimeout(binariesTabLoaded, 1*time.Second); err != nil {
		log.Fatal("Binaries tab failed to be loaded")
	}

	//upload VBIOS & OSDB separately
	uploadVbiosBinary(wd, asicConf.VbiosVersion)
	uploadOsdbBinary(wd, asicConf.OsdbVersion)

	return nil
}

func testReportUploaded(wd webdriver.WebDriver) (bool, error) {

	_, err := wd.FindElement(webdriver.ByXPATH, "//span[contains(text(), 'UPLOAD NEW REPORT')]")
	if err != nil {
		return false, err
	}

	return true, nil
}

func uploadTestReport(wd webdriver.WebDriver) error {

	//goto Test Reports tab
	testReportTab, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'Test Reports')]")
	if err != nil {
		log.Fatal(err)
	}
	testReportTab.Click()

	//check test report uploaded or not
	if err = wd.WaitWithTimeout(testReportUploaded, 5*time.Second); err == nil {
		log.Println("- SKIP - Test Report has alreayd been uploaded yet")
		return err
	}

	//add report click
	addReportBtn, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'ADD REPORT')]")
	if err != nil {
		log.Fatal(err)
	}
	addReportBtn.Click()

	//select file
	fileInput, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='upload-file']/input[@type='file']")
	if err != nil {
		log.Fatal(err)
	}
	fileInput.SendKeys(stackConf.StackPath + "/" + stackConf.Version + "/" + stackConf.TestReport)

	//upload click
	uploadBtn, err := wd.FindElement(webdriver.ByXPATH, "//button[@class='save-button mat-button ng-star-inserted']/span[contains(text(), 'UPLOAD')]")
	if err != nil {
		log.Fatal(err)
	}
	uploadBtn.Click()

	//check results
	if err = wd.WaitWithTimeout(testReportUploaded, 5*time.Second); err != nil {
		log.Fatal("Test report upload failed...")
	} else {
		log.Println("Test report upload successful")
	}
	return nil
}

//UpdateNavi10SMT hello world
func UpdateNavi10SMT(wd webdriver.WebDriver) {

	log.Println("Go to stacks")

	for _, entry := range stackConf.LnxStack {

		if found := gotoSpecNavi10Stack(wd, entry.StackName); found == true { //upload binaries if founded
			uploadBinaries(wd, entry)
			uploadTestReport(wd)
		}
	}
}
