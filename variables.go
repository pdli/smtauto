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
		"D412": "20.30",
		"D414": "20.30",
		"D417": "20.30",
	}

	asicNameMap = map[string]string{
		"D412": "Navi21 XTX",
		"D414": "Navi21 XL",
		"D417": "Navi21 GL-XTA",
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
