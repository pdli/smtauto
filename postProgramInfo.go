package smtauto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
)

func writeJSONFile(data StackConf) {

	file, _ := json.MarshalIndent(data, "", "    ")

	_ = ioutil.WriteFile("stack_conf.json", file, 0644)

	fmt.Println("Called write Json File ")
}

func calcSmtStackName(vbios string) string {

	var stackName = ""

	exp := `D41(\d){2}`
	r := regexp.MustCompile(exp)
	if found := r.FindAllString(vbios, 1); found != nil {
		stackName = found[0] + "1" + stackConf.Version
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

	exp := `D41(\d)`
	r := regexp.MustCompile(exp)

	if found := r.FindAllString(vbios, 1); found != nil {
		asicName = asicNameMap[found[0]]
	}

	return asicName
}

func calcTargetRelease(vbios string) string {

	var targetRelease = ""

	exp := `D41(\d)`
	r := regexp.MustCompile(exp)

	if found := r.FindAllString(vbios, 1); found != nil {
		targetRelease = targetReleaseMap[found[0]]
	}

	return targetRelease
}

//OSDBVersion - 20.30-1085420-ubuntu-20.04
func calcOsdbVersion(vbios string, osdbSlice []string) string {

	var osdbName = ""

	if targetRelease := calcTargetRelease(vbios); targetRelease != "" {
		for _, osdb := range osdbSlice {
			exp := targetRelease + `-(\d)*` + `-([a-z]*)` + `-(\d)*.(\d)*`
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

	osdbName := calcOsdbVersion(vbios, osdbSlice)

	exp := `(\d)*$`
	r := regexp.MustCompile(exp)
	if found := r.FindAllString(osdbName, 1); found != nil {
		osdbID = found[0]
	}

	return osdbID
}

//PostAsicConf will collect & write ASIC config into stack_conf.JSON
func PostAsicConf(ww string) {

	//input WW48,...
	stackConf.Version = ww

	vbiosSlice := GetVBIOS()
	osdbSlice := GetOSDB()

	asicConf := make([]AsicConf, len(vbiosSlice))

	i := 0
	for _, raw := range vbiosSlice {
		if raw != "" {
			asicConf[i].ProgramName = "Navi 21"
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

	writeJSONFile(stackConf)
}
