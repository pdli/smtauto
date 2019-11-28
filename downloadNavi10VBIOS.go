package smtauto

import(

    "log"
    "io/ioutil"
    "strings"
    "regexp"
    "os"
    "os/exec"
//    "encoding/json"
)

func wget(url, filepath string) (error){
  //run shell `wget URL -O filepath`
  cmd := exec.Command("wget", url, "-O", filepath, "--user", "amd/paulil", "--ask-password")
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  return cmd.Run()
}

func downloadBiosPage(){

    url := "http://home.amd.com/VideoBios/Video%20BIOS%20Releases/SingleASICRelease.asp?AsicName=navi10"
    filepath := "./SingleAsicNavi10.html"
    if err := wget(url, filepath); err != nil{
       log.Fatal( err )
    }
}

func getSpecVbiosLink( index int )(string) {

    var vbiosLink = ""

    //read navi10 bios web page
    raw, err := ioutil.ReadFile("./SingleAsicNavi10.html")
    if err != nil{
        log.Fatal( err )
    }
    weblines := strings.Split(string(raw), "\n")

    //get download link for specific bios
    exp := "http://storeiis2/BIOSTest/SignedBIOS.*_"+ vbiosFileNameList[ index ] +".*signed.rom"
    r := regexp.MustCompile( exp )

    for _, entry := range weblines {
        if found := r.FindAllString( entry , 1); found != nil {
            vbiosLink = found[0]
        }
    }

    exp = `'>.*`
    r = regexp.MustCompile( exp )
    vbiosLink = r.ReplaceAllString( vbiosLink, "")

    //display vbioslink, return
    log.Println( "  Here is vbios link --", vbiosLink )

    return vbiosLink
}

func GetLatestVbios() {

    downloadBiosPage()

    for index, _ := range vbiosFileNameList {

        _ = getSpecVbiosLink( index )
    }
}
