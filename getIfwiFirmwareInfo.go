package smtauto

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func readFileLineByLine(fileName string) []string {

	// Read file line by line
	file, err := os.Open(stackConf.StackPath + "/" + stackConf.Version + "/" + fileName)

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

	return txtlines
}

func findGpuDriverComponentFromList(list []string, comp string) (bool, string) {

	//Format of FW version - Firmware,"MC: <br/>VCE: 0.0<br/>UVD: 0.0<br/> ....
	//									RLC: 80.1<br/>RLC SRLC: 0.0
	re := regexp.MustCompile("Firmware,.*<br/>" + comp + ": " + `(\w+\.\w+)(<br.*)`)

	for _, value := range list {
		str := re.FindStringSubmatch(value)
		if len(str) > 2 {
			return true, str[1]
		}
	}

	fmt.Println("**** WARNING : Failed to find firmware from Linux GPU driver - ", comp)

	return false, ""
}

func findIfwiComponentFromList(list []string, comp string) (bool, string) {

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

	fwConfList := readFileLineByLine("FirmwareReleaseNote.txt")

	fmt.Println("==== Get IFWI Firmware Config ====")
	_, ifwiConf.UMC = findIfwiComponentFromList(fwConfList, "UMC")
	_, ifwiConf.DMUCB = findIfwiComponentFromList(fwConfList, "DMCUB")
	_, ifwiConf.SecPolicyL0 = findIfwiComponentFromList(fwConfList, "SEC POLICY L0/L1")
	_, ifwiConf.SecPolicyL1 = findIfwiComponentFromList(fwConfList, "SEC POLICY L0/L1")
	_, ifwiConf.SMU = findIfwiComponentFromList(fwConfList, "Mini-PMFW")
	_, ifwiConf.PspBL = findIfwiComponentFromList(fwConfList, "PSP BL")
	_, ifwiConf.DXIO = findIfwiComponentFromList(fwConfList, "DXIO")
	_, ifwiConf.VBL = findIfwiComponentFromList(fwConfList, "PSP VBL")

	fmt.Println(ifwiConf)

	return ifwiConf
}

//GetGpuDriverComponentsForStack to get Linux GPU driver firmwares version for SMT stack
func GetGpuDriverComponentsForStack() GpuDriverFirmwareConf {

	var gpuDriverFWConf GpuDriverFirmwareConf

	fwConfList := readFileLineByLine("sw_info.csv")

	_, gpuDriverFWConf.SDMA = findGpuDriverComponentFromList(fwConfList, "SDMA0")
	_, gpuDriverFWConf.ME = findGpuDriverComponentFromList(fwConfList, "ME")
	_, gpuDriverFWConf.MEC = findGpuDriverComponentFromList(fwConfList, "MEC")
	_, gpuDriverFWConf.VCN = findGpuDriverComponentFromList(fwConfList, "VCN")
	_, gpuDriverFWConf.PFP = findGpuDriverComponentFromList(fwConfList, "PFP")
	_, gpuDriverFWConf.RLC = findGpuDriverComponentFromList(fwConfList, "RLC")
	_, gpuDriverFWConf.SMC = findGpuDriverComponentFromList(fwConfList, "SMC")
	_, gpuDriverFWConf.CE = findGpuDriverComponentFromList(fwConfList, "CE")
	_, gpuDriverFWConf.SOS = findGpuDriverComponentFromList(fwConfList, "SOS")

	fmt.Println("==== Get GPU Driver Firmware Config ====")
	fmt.Println(gpuDriverFWConf)

	return gpuDriverFWConf

}
