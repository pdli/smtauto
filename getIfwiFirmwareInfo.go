package smtauto

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func readReleaseNotesOfIFWI() []string {

	// Read file line by line
	file, err := os.Open(stackConf.StackPath + "/" + stackConf.Version + "/" + "FirmwareReleaseNote.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	//Supposed there are only the first 20 lines will be useful
	return txtlines[0:19]
}

func findComponentFromList(list []string, comp string) (bool, string) {

	//Format of FW version in release note - UMC  v87.03.15.00[2020-06-05]
	//										 SEC  v0B.21.00.1C[2020-06-14]
	//Return - found, value
	re := regexp.MustCompile(comp + `\s*v(\w{2}(\.\w{2})*)`)

	for _, value := range list {
		str := re.FindStringSubmatch(value)
		if len(str) > 2 {
			return true, str[1]
		}
	}

	fmt.Println("**** WARNING : Failed to find firmware from IFWI release notes - ", comp)

	return false, ""
}

//GetIfwiComponentsForStack to get IFWI firmwares version for SMT Stack
func GetIfwiComponentsForStack() IfwiFirmwareConf {

	var ifwiConf IfwiFirmwareConf

	fwConfList := readReleaseNotesOfIFWI()

	fmt.Println("=====")
	_, ifwiConf.MC = findComponentFromList(fwConfList, "UMC")
	_, ifwiConf.DMUCB = findComponentFromList(fwConfList, "DMCUB")
	_, ifwiConf.SecPolicyL0 = findComponentFromList(fwConfList, "SEC POLICY L0/L1")
	_, ifwiConf.SecPolicyL1 = findComponentFromList(fwConfList, "SEC POLICY L0/L1")
	_, ifwiConf.SMU = findComponentFromList(fwConfList, "Mini-PMFW")
	_, ifwiConf.PspBL = findComponentFromList(fwConfList, "PSP BL")
	_, ifwiConf.DXIO = findComponentFromList(fwConfList, "DXIO")
	_, ifwiConf.VBL = findComponentFromList(fwConfList, "PSP VBL")

	fmt.Println("++++++")
	fmt.Println(ifwiConf)

	return ifwiConf
}
