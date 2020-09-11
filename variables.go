package smtauto

//AsicConf define the info required to maintain one ASIC
type AsicConf struct {
	ProgramName     string
	OsName          string
	DistroName      string
	DistroShortName string
	AsicName        string
	StackName       string
	TargetRelease   string
	VbiosVersion    string
	VbiosFileName   string
	OsdbVersion     string
	OsdbID          string
}

//StackConf defines the info required to manage one weekly Linux Stack for Navi1x on SMT website
type StackConf struct {
	Version     string
	StackPath   string
	TestReport  string
	ProgramName string
	ProgramID   string
	LnxStack    []AsicConf
}

//VbiosConf defines the info collected from VBIOS website
type VbiosConf struct {
	Name        string
	Link        string
	Version     string
	LastVersion string
}

//IfwiFirmwareConf defines the firmwares version in IFWI
type IfwiFirmwareConf struct {
	MC          string
	DMUCB       string
	SecPolicyL0 string
	SecPolicyL1 string
	SMU         string
	PspBL       string
	DXIO        string
	VBL         string
}

//GpuDriverFirmwareConf defines the firmwares version from Linux GPU driver
type GpuDriverFirmwareConf struct {
	SDMA string
	ME   string
	MEC  string
	VCN  string
	PFP  string
	RLC  string
	SMC  string
	CE   string
	SOS  string
}

var (

	//used to upload binaries
	osName          = "Ubuntu 20.04.1"
	distroName      = "ubuntu-20.04"
	distroShortName = "u2004_64"

	//programIDMap provides the mapping btw Program Name and Program ID of SMT stack
	programIDMap = map[string]string{
		"Navi21": "1289",
		"Navi22": "1297",
		"Navi10": "1258",
	}

	//get target release per each ASIC - gaming /wks
	targetReleaseMap = map[string]string{
		"D412": "20.45",
		"D414": "20.45",
		"D417": "20.45",
		"D512": "20.45",
		"D511": "20.45",
		"D199": "20.30",
	}

	//get asic Name per each ASIC
	asicNameMap = map[string]string{
		"D412": "D41201 XTX",
		"D414": "D41401 XL",
		"D417": "D41702 Navi21 GL XL",
		"D511": "D51101 XT",
		"D512": "D51201 XTX",
		"D199": "D19901 XT",
	}
	//Not need to maintain - 7/29/2020
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
			"D417GLXL",
		},
	}
)

var (
	vbiosConf = []VbiosConf{}
)
