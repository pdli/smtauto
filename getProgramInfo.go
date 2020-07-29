package smtauto

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	//"encoding/json"
)

func readLinuxStackDir() []string {
	log.Println("stack version and path", stackConf.Version, stackConf.StackPath) 

	filesName, err := ioutil.ReadDir(stackConf.StackPath + "/" + stackConf.Version)
	if err != nil {
		log.Fatal(err)
	}

	y := make([]string, len(filesName))
	for index, f := range filesName {
		fmt.Println("Include file ==> ", f.Name())
		y[index] = f.Name()
	}

	return y
}

//GetVBIOS list VBIOS info from specific folder - /opt/shares/Navi10_Stack/WWxx
func GetVBIOS() []string {

	filesName := readLinuxStackDir()
	vbiosSlice := make([]string, len(filesName))

	fmt.Println("Get VBIOS info -> ")

	exp := `D(\d){4}[0-9a-zA-Z]{3}[_|.]`
	r, err := regexp.Compile(exp)
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for _, f := range filesName {
		if found := r.FindAllString(f, -1); found != nil {
			vbiosSlice[count] = f
			fmt.Println("  ==> ", count, vbiosSlice[count])
			count++
		}
	}

	vbiosSlice = append(vbiosSlice[:count])
	return vbiosSlice
}

//GetOSDB list OSDB info from specific folder - /opt/shares/Navi10_Stack/WWxx
func GetOSDB() []string {

	filesName := readLinuxStackDir()
	osdbSlice := make([]string, len(filesName))

	fmt.Println("Get OSDB info ->")

	exp := `^amdgpu-pro-20.(\d)0-(.)*-ubuntu-20.04.tar.xz`
	r := regexp.MustCompile(exp)

	count := 0
	for _, f := range filesName {
		if found := r.FindAllString(f, -1); found != nil {
			osdbSlice[count] = f
			fmt.Println("  ==> ", count, osdbSlice[count])
			count++
		}
	}

	osdbSlice = append(osdbSlice[:count])
	return osdbSlice
}

//GetTestReport list Test Report info from specific folder - /opt/shares/Navi10_Stack/WWxx
func GetTestReport() string {

	filesName := readLinuxStackDir()
	testReport := ""

	fmt.Println("Get Test Report info ->")

	exp := `Navi(.)*.msg`
	r := regexp.MustCompile(exp)

	for _, f := range filesName {
		if found := r.FindAllString(f, -1); found != nil {
			testReport = f
			fmt.Println("  ==> ", testReport)
			break
		}
	}

	return testReport
}
