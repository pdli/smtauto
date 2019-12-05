package smtauto

//    "fmt"

type AsicConf struct {
	AsicName      string
	StackName     string
	TargetRelease string
	VbiosVersion  string
	VbiosFileName string
	OsdbVersion   string
	OsdbID        string
}

type StackConf struct {
	Version    string
	StackPath  string
	TestReport string
	LnxStack   []AsicConf
}

type VbiosConf struct {
	Name        string
	Link        string
	Version     string
	LastVersion string
}

var (
	targetReleaseMap = map[string]string{
		"D182": "19.40",
		"D187": "19.40",
		"D188": "19.50",
		"D189": "19.50",
	}

	asicNameMap = map[string]string{
		"D182": "Navi10 XT",
		"D187": "Navi10 XM",
		"D188": "Navi10 Pro-XL",
		"D189": "Navi10 XLE",
	}

	vbiosFileNameList = []string{
		"D1820101",
		"D1870101",
		"D1880201",
		"D1890101",
	}
 )

var (

	vbiosConf = []VbiosConf{}
)
