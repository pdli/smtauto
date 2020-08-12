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

	_, err := wd.FindElement(webdriver.ByXPATH, "(//p-panel)[1]//*[contains(text(), 'Status: Uploaded')]")
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
	selectActBtn, err := wd.FindElement(webdriver.ByXPATH, "(//div[@class='row'])[1]/div[@class='element-name']/button[@class='mat-raised-button']")
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

	resultSpan, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='query-results ng-star-inserted'][last()]/*[contains(text(), '"+biosVersion+"')]")
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
		log.Fatal("VBIOS Binary linked failed...", err)
	} else {
		log.Println("VBIOS Binary linked successful")
	}

	return nil
}

func gotoSpecSMTStack(wd webdriver.WebDriver, stackName string) bool {

	var loaded = true //false if not loaded

	log.Println("***** Preparing to update Stack - ", stackName)

	err := wd.Get("http://smt.amd.com/#/view/program/1289")
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second) //wait for stacks page loading

	lnxStackSpan, err := wd.FindElement(webdriver.ByXPATH, "//h4[contains(text(), '"+stackName+"')]")
	if err != nil {
		log.Println("- SKIP - Cant' find the spec Linux Stack - ", stackName)
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

	_, err := wd.FindElement(webdriver.ByXPATH, "(//p-panel)[2]//*[contains(text(), 'Linux-GPU-Driver')]")
	if err != nil {
		log.Println("- SKIP - OSDB binary is not requried by SMT.")
		return true, nil
	}

	_, err = wd.FindElement(webdriver.ByXPATH, "(//p-panel)[2]//*[contains(text(), 'Status: Uploaded')]")
	if err != nil {
		return false, err
	}

	return true, nil
}

func uploadOsdbBinary(wd webdriver.WebDriver, osdbVersion string) error {

	//check if updated
	if err := wd.WaitWithTimeout(osdbUpdated, 5*time.Second); err == nil {
		log.Println("- SKIP - OSDB has already been updated - ", osdbVersion)
		return nil
	}

	//*****update vbios since it is not updated
	//select Action
	selectActBtn, err := wd.FindElement(webdriver.ByXPATH, "(//div[@class='row'])[2]/div[@class='element-name']/button[@class='mat-raised-button']")
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
	if err = wd.WaitWithTimeout(binaryLinked, 30*time.Second); err != nil {
		log.Fatal("OSDB Binary linked failed...", err)
	} else {
		log.Println("OSDB Binary linked successful")
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

func newStackPageLoaded(wd webdriver.WebDriver) (bool, error) {
	_, err := wd.FindElement(webdriver.ByXPATH, "//span[contains(text(), 'SKU LIST')]")
	if err != nil {
		return false, err
	}

	return true, nil
}

func skuListLoaded(wd webdriver.WebDriver) (bool, error) {
	_, err := wd.FindElement(webdriver.ByXPATH, "(//div[@class='cdk-overlay-pane'])[2]")
	if err != nil {
		return false, err
	}

	return true, nil
}

func createStacks(wd webdriver.WebDriver, asicConf AsicConf) error {

	log.Println("==> Start to create New stack: ", asicConf.StackName)

	//check if stack existed
	_, err := wd.FindElement(webdriver.ByXPATH, "//h4[contains(text(), '"+asicConf.StackName+"')]")
	if err == nil {
		log.Println("==> Linux Stack has been created: ", asicConf.StackName)
		return nil
	}

	log.Println("==> Need to create a new Linux Stack: ", asicConf.StackName)
	//create a new stack if not existed
	//click new stack
	newBtn, err := wd.FindElement(webdriver.ByXPATH, "(//button[@class='add-stack mat-button ng-star-inserted'])[1]")
	if err != nil {
		log.Fatal(err)
	}
	newBtn.Click()

	//wait for new stack page loaded
	if err = wd.WaitWithTimeout(newStackPageLoaded, 10*time.Second); err != nil {
		log.Fatal("New stack page failed to be loaded")
		return nil
	}

	//click SKU List
	skuListBtn, err := wd.FindElement(webdriver.ByXPATH, "//span[contains(text(), 'SKU LIST')]")
	if err != nil {
		log.Fatal(err)
	}
	skuListBtn.Click()

	//wait for SKU list displayed
	if err = wd.WaitWithTimeout(skuListLoaded, 10*time.Second); err != nil {
		log.Fatal("Failed to load SKU List")
	}

	//choose SKU
	skuOpt, err := wd.FindElement(webdriver.ByXPATH, "//mat-option//span[contains(text(), '"+asicConf.AsicName+"')]")
	if err != nil {
		log.Fatal(err)
	}
	skuOpt.Click()

	//choose start date
	startBtn, err := wd.FindElement(webdriver.ByXPATH, "//span[contains(text(), 'Start Date')]")
	if err != nil {
		log.Fatal(err)
	}
	startBtn.Click()

	startTodayBtn, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='mat-calendar-body-cell-content mat-calendar-body-today']")
	if err != nil {
		log.Fatal(err)
	}
	startTodayBtn.Click()

	//choose end date
	endBtn, err := wd.FindElement(webdriver.ByXPATH, "//span[contains(text(), 'Release Date')]")
	if err != nil {
		log.Fatal(err)
	}
	endBtn.Click()

	endDateBtn, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='mat-calendar-body-cell-content mat-calendar-body-today']")
	if err != nil {
		log.Fatal(err)
	}
	endDateBtn.Click()

	//Set version sufix =LNX
	sufixInput, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='suffix small ng-star-inserted']/input[@type='text']")
	if err != nil {
		log.Fatal(err)
	}
	sufixInput.SendKeys("LNX")

	//choose IFWI element
	ifwiCheckbox, err := wd.FindElement(webdriver.ByXPATH, "(//div[@class='sub-content']//mat-checkbox)[2]")
	if err != nil {
		log.Fatal(err)
	}
	ifwiCheckbox.Click()

	//choose GPU element
	gpuCheckbox, err := wd.FindElement(webdriver.ByXPATH, "(//div[@class='sub-content']//mat-checkbox)[4]")
	if err != nil {
		log.Fatal(err)
	}
	gpuCheckbox.Click()

	//un-click Testing
	testCheckbox, err := wd.FindElement(webdriver.ByXPATH, "//mat-checkbox[@class='mat-checkbox mat-accent ng-untouched ng-pristine ng-valid mat-checkbox-checked']")
	if err != nil {
		log.Fatal(err)
	}
	testCheckbox.Click()

	//create Stack
	createBtn, err := wd.FindElement(webdriver.ByXPATH, "//button[@class='save-button mat-button ng-star-inserted']")
	if err != nil {
		log.Fatal(err)
	}
	createBtn.Click()

	//Wait for the completion of creation
	time.Sleep(10 * time.Second)

	return nil
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

	//Wait for the completion of upload
	time.Sleep(10 * time.Second)

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
	log.Println("Add report")
	addReportBtn, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'ADD REPORT')]")
	if err != nil {
		log.Fatal(err)
	}
	addReportBtn.Click()

	//select file
	log.Println("To select Files - " + stackConf.StackPath + "/" + stackConf.Version + "/" + stackConf.TestReport)
	fileInput, err := wd.FindElement(webdriver.ByXPATH, "//div[@class='upload-file']/input[@type='file']")
	if err != nil {
		log.Fatal(err)
	}
	fileInput.SendKeys(stackConf.StackPath + "/" + stackConf.Version + "/" + stackConf.TestReport)

	//upload click
	log.Println("To click upload")
	uploadBtn, err := wd.FindElement(webdriver.ByXPATH, "//button[@class='save-button mat-button ng-star-inserted']/span[contains(text(), 'UPLOAD')]")
	if err != nil {
		log.Fatal(err)
	}
	uploadBtn.Click()

	//check results
	log.Println("To check result of test report")
	if err = wd.WaitWithTimeout(testReportUploaded, 60*time.Second); err != nil {
		log.Fatal("Test report upload failed...")
	} else {
		log.Println("Test report upload successful")
	}
	return nil
}

//updateStackComponents - update IFWI / Linux-GPU-Driver
func updateStackComponents(wd webdriver.WebDriver) {

	// get Firmwre info of IFWI
	ifwiConf := GetIfwiComponentsForStack()

	// get Firmware info of Linux GPU Driver
	gpuDriverConf := GetGpuDriverComponentsForStack()

	//goto EDIT web page
	editBtn, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'EDIT')]")
	if err != nil {
		log.Fatal(err)
	}
	editBtn.Click()

	//goto Firmware tab
	firmwaresTab, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'Firmwares')]")
	if err != nil {
		log.Fatal(err)
	}
	firmwaresTab.Click()

	// 1- Update Firmware info of IFWI
	//input MC
	mcInput, err := wd.FindElement(webdriver.ByXPATH, "//app-firmware-select/div/div/table/tbody/tr[1]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	mcInput.Clear()
	mcInput.SendKeys(ifwiConf.MC)

	//input DMUCB
	dmucbInput, err := wd.FindElement(webdriver.ByXPATH, "//app-firmware-select/div/div/table/tbody/tr[2]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	dmucbInput.Clear()
	dmucbInput.SendKeys(ifwiConf.DMUCB)

	//input SEC Policy L0
	policyL0Input, err := wd.FindElement(webdriver.ByXPATH, "//app-firmware-select/div/div/table/tbody/tr[3]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	policyL0Input.Clear()
	policyL0Input.SendKeys(ifwiConf.SecPolicyL0)

	//input SEC Policy L1
	policyL1Input, err := wd.FindElement(webdriver.ByXPATH, "//app-firmware-select/div/div/table/tbody/tr[4]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	policyL1Input.Clear()
	policyL1Input.SendKeys(ifwiConf.SecPolicyL1)

	//input SMU
	smuInput, err := wd.FindElement(webdriver.ByXPATH, "//app-firmware-select/div/div/table/tbody/tr[5]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	smuInput.Clear()
	smuInput.SendKeys(ifwiConf.SMU)

	//input PSP-BL
	pspBLInput, err := wd.FindElement(webdriver.ByXPATH, "//app-firmware-select/div/div/table/tbody/tr[6]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	pspBLInput.Clear()
	pspBLInput.SendKeys(ifwiConf.PspBL)

	//input DXIO
	dxioInput, err := wd.FindElement(webdriver.ByXPATH, "//app-firmware-select/div/div/table/tbody/tr[7]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	dxioInput.Clear()
	dxioInput.SendKeys(ifwiConf.DXIO)

	//input VBL
	vblInput, err := wd.FindElement(webdriver.ByXPATH, "//app-firmware-select/div/div/table/tbody/tr[8]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	vblInput.Clear()
	vblInput.SendKeys(ifwiConf.VBL)

	// 2- Update firmware version of Linux GPU driver
	//input SDMA
	sdmalInput, err := wd.FindElement(webdriver.ByXPATH, "//div[4]//app-firmware-select/div/div/table/tbody/tr[1]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	sdmalInput.Clear()
	sdmalInput.SendKeys(gpuDriverConf.SDMA)

	//input ME
	meInput, err := wd.FindElement(webdriver.ByXPATH, "//div[4]//app-firmware-select/div/div/table/tbody/tr[2]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	meInput.Clear()
	meInput.SendKeys(gpuDriverConf.ME)

	//input MEC
	mecInput, err := wd.FindElement(webdriver.ByXPATH, "//div[4]//app-firmware-select/div/div/table/tbody/tr[3]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	mecInput.Clear()
	mecInput.SendKeys(gpuDriverConf.MEC)

	//input VCN
	vncInput, err := wd.FindElement(webdriver.ByXPATH, "//div[4]//app-firmware-select/div/div/table/tbody/tr[4]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	vncInput.Clear()
	vncInput.SendKeys(gpuDriverConf.VCN)

	//input PFP
	pfpInput, err := wd.FindElement(webdriver.ByXPATH, "//div[4]//app-firmware-select/div/div/table/tbody/tr[5]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	pfpInput.Clear()
	pfpInput.SendKeys(gpuDriverConf.PFP)

	//input RLC
	rlcInput, err := wd.FindElement(webdriver.ByXPATH, "//div[4]//app-firmware-select/div/div/table/tbody/tr[6]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	rlcInput.Clear()
	rlcInput.SendKeys(gpuDriverConf.RLC)

	//input SMC
	smcInput, err := wd.FindElement(webdriver.ByXPATH, "//div[4]//app-firmware-select/div/div/table/tbody/tr[7]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	smcInput.Clear()
	smcInput.SendKeys(gpuDriverConf.SMC)

	//input CE
	ceInput, err := wd.FindElement(webdriver.ByXPATH, "//div[4]//app-firmware-select/div/div/table/tbody/tr[8]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	ceInput.Clear()
	ceInput.SendKeys(gpuDriverConf.CE)

	//input SOS
	sosInput, err := wd.FindElement(webdriver.ByXPATH, "//div[4]//app-firmware-select/div/div/table/tbody/tr[9]/td[2]/input[2]")
	if err != nil {
		log.Fatal(err)
	}
	sosInput.Clear()
	sosInput.SendKeys(gpuDriverConf.SOS)

	//click Save Stack
	saveStackBtn, err := wd.FindElement(webdriver.ByXPATH, "//*[contains(text(), 'SAVE STACK')]")
	if err != nil {
		log.Fatal(err)
	}
	saveStackBtn.Click()
}

//UpdateSMTforLinux to update SMT
func UpdateSMTforLinux(wd webdriver.WebDriver, disableReport bool) {

	log.Println("Go to stacks")

	for _, entry := range stackConf.LnxStack {

		createStacks(wd, entry)
		if found := gotoSpecSMTStack(wd, entry.StackName); found == true { //upload binaries if founded
			if entry.ProgramName == "Navi21" {
				updateStackComponents(wd)
			}

			uploadBinaries(wd, entry)

			if disableReport == false {
				uploadTestReport(wd)
			}
		}
	}
}
