package smtauto

import (
    "fmt"
    "io/ioutil"
    "log"
    "regexp"
    //"encoding/json"
)

func Test(x string) (string) {

    fmt.Println("X is -- ", x)

    return  x
}

func readNavi10StackDir() ([]string) {

    filesName, err := ioutil.ReadDir( stackConf.StackPath + "/" + stackConf.Version)
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

func GetVBIOS() ([]string){

    filesName :=  readNavi10StackDir()
    vbiosSlice := make([]string, len(filesName))

    fmt.Println("Get VBIOS info -> ")

    exp := `D18(\d){3}01[_|.]`
    r, err := regexp.Compile( exp )
    if err != nil {
        log.Fatal( err )
    }

    count := 0
    for _, f := range filesName {
        if found := r.FindAllString( f, -1 ); found != nil {
          vbiosSlice[count] =  f
          fmt.Println("  ==> ", count, vbiosSlice[count])
          count ++
        }
    }

    vbiosSlice = append( vbiosSlice[:count])
    return vbiosSlice
}

func GetOSDB() ([]string){

    filesName := readNavi10StackDir()
    osdbSlice := make([]string, len(filesName))

    fmt.Println("Get OSDB info ->")

    exp := `^amdgpu-pro-19.(\d)0-(.)*-ubuntu-18.04.tar.xz`
    r := regexp.MustCompile( exp )

    count := 0
    for _, f := range filesName {
        if found := r.FindAllString( f, -1); found != nil {
            osdbSlice[ count ] = f
            fmt.Println("  ==> ", count, osdbSlice[count])
            count ++
        }
    }

    osdbSlice = append( osdbSlice[ :count])
    return osdbSlice
}

func GetTestReport() (string) {

    filesName := readNavi10StackDir()
    testReport := ""

    exp := `Navi10 Linux SW Stack Test Report (.)*.msg`
    r := regexp.MustCompile( exp )

    for _, f := range filesName {
        if found := r.FindAllString( f, -1); found != nil {
            testReport =  f
            fmt.Println("Get test report ==> ", testReport)
            break
        }
    }

    return testReport
}
