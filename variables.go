
package smtauto

import (
//    "fmt"

)

type AsicConf struct {
    AsicName string
    StackName string
    TargetRelease string
    VbiosVersion string
    VbiosFileName string
    OsdbVersion string
    OsdbID string
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

    stackConf = StackConf{
        Version: "WW47",
        StackPath: "/opt/shares/Navi10_Stack/",
    }
)

