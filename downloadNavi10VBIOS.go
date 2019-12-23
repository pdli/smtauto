package smtauto

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	//    "encoding/json"
)

func wget(url, filepath string) error {
	//run shell `wget URL -O filepath`
	cmd := exec.Command("wget", url, "-O", filepath, "--user", "amd/paulil", "--ask-password")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func downloadBiosPage(asicName string) {

	url := "http://home.amd.com/VideoBios/Video%20BIOS%20Releases/SingleASICRelease.asp?AsicName=" + asicName
	filepath := "./SingleAsic" + asicName + ".html"
	if err := wget(url, filepath); err != nil {
		log.Fatal(err)
	}
}

func getSpecVbiosLink(asicName string, index int) string {

	var vbiosLink = ""

	//read navi10 bios web page
	raw, err := ioutil.ReadFile("./SingleAsic" + asicName + ".html")
	if err != nil {
		log.Fatal(err)
	}
	weblines := strings.Split(string(raw), "\n")

	//get download link for specific bios
	exp := "http://storeiis2/BIOSTest/SignedBIOS.*_" + biosFileMap[asicName][index] + ".*signed.rom"
	r := regexp.MustCompile(exp)

	for _, entry := range weblines {
		if found := r.FindAllString(entry, 1); found != nil {
			vbiosLink = found[0]
		}
	}

	exp = `'>.*`
	r = regexp.MustCompile(exp)
	vbiosLink = r.ReplaceAllString(vbiosLink, "")

	//display vbioslink, return
	log.Println("  Here is vbios link --", vbiosLink)

	return vbiosLink
}

//GetLatestVbios to list latet VBIOS version per specific ASIC
func GetLatestVbios(asicName string) {

	downloadBiosPage(asicName)

	for index := range biosFileMap[asicName] { //Match / Should omit 2nd valud

		link := getSpecVbiosLink(asicName, index)
		if link != "" {
			vbiosConf = append(vbiosConf, VbiosConf{Name: biosFileMap[asicName][index], Link: link})
		}
	}

	log.Println(vbiosConf)
}
