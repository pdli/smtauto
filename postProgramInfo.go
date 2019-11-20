package smtauto

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "regexp"
)

type AsicConf struct {
    AsicName string
    StackName string
    TargetRelease string
    VBIOS string
    OSDB string
}

type StackConf struct {
    Version string
    StackPath string
    TestReport string
    LnxStack []AsicConf
}

var (
    targetReleaseMap = map[string]string {
        "D182": "19.40",
        "D187": "19.40",
        "D188": "19.50",
        "D189": "19.50",
    }

    asicNameMap = map[string]string {
        "D182": "Navi10 XT",
        "D187": "Navi10 XM",
        "D188": "Navi10 Pro-XL",
        "D189": "Navi10 XLE",
    }

    asicConf = make([]AsicConf, 10)

    stackConf = StackConf{
        Version: "WW46",
        StackPath: "/opt/shares/Navi10_Stack/",
    }
)

func writeJsonFile( data StackConf ) {

    file, _ := json.MarshalIndent( data, "", "    ")

    _ = ioutil.WriteFile("test.json", file, 0644)

    fmt.Println("Called write Json File ")
}

func calcStackName(vbios string) (string) {

    var stackName = ""

    exp := `D18(\d)`
    r := regexp.MustCompile( exp )
    if found := r.FindAllString( vbios, 1); found != nil {
        stackName = found[0] + "01W19" + "46" + "LN5" 
    }

    return stackName
}

func calcVBIOS(vbios string) (string) {

    var vbiosName = ""

    exp := `D(\d)*[.|_](\d)*`
    r := regexp.MustCompile( exp )
    if found := r.FindAllString( vbios, 1); found != nil {
        vbiosName = found[0]
    }

    exp = `\.`
    r = regexp.MustCompile( exp )
    vbiosName = r.ReplaceAllString( vbiosName, "_")

    return vbiosName
}

func calcAsicName(vbios string) (string) {

    var asicName = ""

    exp := `D18(\d)`
    r := regexp.MustCompile( exp )

    if found := r.FindAllString( vbios, 1); found != nil {
        asicName = asicNameMap[found[0]]
    }

   return asicName
}

func calcTargetRelease(vbios string) (string) {

    var targetRelease = ""

    exp := `D18(\d)`
    r := regexp.MustCompile( exp )

    if found := r.FindAllString( vbios, 1); found != nil {
        targetRelease = targetReleaseMap[found[0]]
    }

    return targetRelease
}


func calcOSDB(vbios string, osdbSlice []string) (string) {

    var osdbName = ""

    fmt.Println("Print OSDB slice - ", osdbSlice)

    if targetRelease := calcTargetRelease( vbios ); targetRelease != "" {
        for  _, osdb := range osdbSlice {
            exp := targetRelease + `-(\d)*`
            r := regexp.MustCompile( exp )
            if found := r.FindAllString( osdb, 1); found != nil {
                osdbName = found[0]
            }
        }
    }

    return osdbName
}


func PostAsicConf() {

    vbiosSlice := GetVBIOS()
    osdbSlice := GetOSDB()

    i := 0
    for _, raw := range vbiosSlice{
        if raw != "" {
            asicConf[i].StackName = calcStackName( raw )
            asicConf[i].VBIOS = calcVBIOS( raw )
            asicConf[i].OSDB = calcOSDB( raw, osdbSlice )//"amdgpu-pro-19.40"
            asicConf[i].AsicName = calcAsicName ( raw ) //"D18x"
            asicConf[i].TargetRelease = calcTargetRelease( raw ) //"19.40"
            i ++
        }
    }

    fmt.Println("ASIC conf ==> ", asicConf)

    stackConf.TestReport = GetTestReport()
    stackConf.LnxStack = asicConf

    writeJsonFile( stackConf )
}
