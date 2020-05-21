package smtauto

//AsicConf define the info required to maintain one ASIC
type AsicConf struct {
	AsicName      string
	StackName     string
	TargetRelease string
	VbiosVersion  string
	VbiosFileName string
	OsdbVersion   string
	OsdbID        string
}

//StackConf defines the info required to manage one weekly Linux Stack for Navi1x on SMT website
type StackConf struct {
	Version    string
	StackPath  string
	TestReport string
	LnxStack   []AsicConf
}

//VbiosConf defines the info collected from VBIOS website
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

	biosFileMap = map[string][]string{
		"navi10": []string{
			"D1820101",
			"D1870101",
			"D1880201",
			"D1890101",
		},
		"navi14": []string{
			"D3220500",
			"D3221500",
			"D3221600",
			"D3250100",
			"D3231000",
		},
        "navi21": []string{
            "D4120XTX",
            "D4140XTX",
            "D41711XT",
        },
	}
)

var (
	vbiosConf = []VbiosConf{}
)
