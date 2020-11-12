package smtauto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func writeJSONFile(data StackConf) {

	file, _ := json.MarshalIndent(data, "", "    ")

	_ = ioutil.WriteFile("stack_conf.json", file, 0644)

	fmt.Println("Called write Json File ")
}

func calcSmtStackName(vbios string) string {

	var stackName = ""

	exp := `(GLXL)` //D417GLXL
	r := regexp.MustCompile(exp)
	vbios = r.ReplaceAllString(vbios, "02")

	/*exp = `D41205`
	r = regexp.MustCompile(exp)
	if found := r.FindAllString(vbios, 1); found != nil {
		stackName = found[0] + stackConf.Version + "A"
		return stackName
	}

	exp = `D41401`
	r = regexp.MustCompile(exp)
	if found := r.FindAllString(vbios, 1); found != nil {
		stackName = found[0] + stackConf.Version + "A"
		return stackName
	}*/

	exp = `D(\d){5}`
	r = regexp.MustCompile(exp)
	if found := r.FindAllString(vbios, 1); found != nil {
		stackName = found[0] + stackConf.Version
		return stackName
	}

	exp = `D(\d){4}`
	r = regexp.MustCompile(exp)
	if found := r.FindAllString(vbios, 1); found != nil {
		stackName = found[0] + "1" + stackConf.Version
		return stackName
	}

	return stackName
}

func calcVbiosVersion(vbios string) string {

	var vbiosName = ""

	exp := `D(\d)*[a-zA-Z]*[.|_](\d)*`
	r := regexp.MustCompile(exp)
	if found := r.FindAllString(vbios, 1); found != nil {
		vbiosName = found[0]
	}

	exp = `\.`
	r = regexp.MustCompile(exp)
	vbiosName = r.ReplaceAllString(vbiosName, ".")

	return vbiosName
}

func calcAsicName(vbios string) string {

	var asicName = ""

	exp := `D(\w){5}`
	r := regexp.MustCompile(exp)

	if found := r.FindAllString(vbios, 1); found != nil {
		asicName = asicNameMap[found[0]]
	}

	return asicName
}

func calcTargetRelease(vbios string) string {

	var targetRelease = ""

	exp := `D(\d){3}`
	r := regexp.MustCompile(exp)

	if found := r.FindAllString(vbios, 1); found != nil {
		targetRelease = targetReleaseMap[found[0]]
	}

	fmt.Println("Calculate Target Release version - ", targetRelease)
	return targetRelease
}

//OSDBVersion - 20.30-1085420-ubuntu-20.04
func calcOsdbVersion(vbios string, osdbSlice []string) string {

	var osdbName = ""

	fmt.Println("osdbSlice is - ", osdbSlice)

	if targetRelease := calcTargetRelease(vbios); targetRelease != "" {
		for _, osdb := range osdbSlice {
			exp := targetRelease + `-(\d)*`
			r := regexp.MustCompile(exp)
			if found := r.FindAllString(osdb, 1); found != nil {
				osdbName = found[0]
			}
		}
	}

	fmt.Println("Calculate OSDB version - ", osdbName)

	return osdbName
}

func calcOsdbID(vbios string, osdbSlice []string) string {

	var osdbID = ""

	//osdbVersion - "20.30-1085420"
	osdbName := calcOsdbVersion(vbios, osdbSlice)

	fmt.Println("osdbname is - ", osdbName)

	exp := `(\d)*`
	r := regexp.MustCompile(exp)
	if found := r.FindAllString(osdbName, -1); found != nil {
		osdbID = found[2]
	}

	return osdbID
}

//GetProgramID to provide program ID of SMT stack per programName
func GetProgramID(programName string) string {
	//pre-define the format of programName - Navi21

	//get program ID from map
	id := programIDMap[programName]

	//Alert if programID doesn't exist
	if id == "" {
		fmt.Println("Error - Invalid Program Name: ", programName)
		os.Exit(1)
	}

	//return program ID
	return id
}

//PostAsicConf will collect & write ASIC config into stack_conf.JSON
func PostAsicConf(ww string, programName string) {

	//input WW48,...
	stackConf.Version = ww
	stackConf.StackPath = stackConf.StackPath + "/" + programName + "_Stack/"

	vbiosSlice := GetVBIOS()
	osdbSlice := GetOSDB()

	asicConf := make([]AsicConf, len(vbiosSlice))

	i := 0
	for _, raw := range vbiosSlice {
		if raw != "" {
			asicConf[i].ProgramName = programName
			asicConf[i].OsName = osName                   //Ubuntu 20.04.1
			asicConf[i].DistroName = distroName           //ubuntu20.04
			asicConf[i].DistroShortName = distroShortName //u2004_64
			asicConf[i].StackName = calcSmtStackName(raw)
			asicConf[i].VbiosVersion = calcVbiosVersion(raw)
			asicConf[i].VbiosFileName = raw
			asicConf[i].OsdbVersion = calcOsdbVersion(raw, osdbSlice) //"20.30-1085420-ubuntu-20.04"
			asicConf[i].OsdbID = calcOsdbID(raw, osdbSlice)           //"1085420"
			asicConf[i].AsicName = calcAsicName(raw)                  //"D18x"
			asicConf[i].TargetRelease = calcTargetRelease(raw)        //"19.40"
			i++
		}
	}

	fmt.Println("ASIC conf ==> ", asicConf)

	stackConf.TestReport = GetTestReport()
	stackConf.LnxStack = asicConf
	stackConf.ProgramName = programName
	stackConf.ProgramID = GetProgramID(programName)

	writeJSONFile(stackConf)
}
