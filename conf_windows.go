package smtauto

import (
	"github.com/radutopala/webdriver/chrome"
)

var (
	stackConf = StackConf{
		Version:   "W2028LNX",
		StackPath: "//pauline.amd.com/shares/Navi21_Stack",
	}

	chrCaps = chrome.Capabilities{
		Args: []string{
			//"--headless",
			//"--disable-gpu",
			"--no-sandbox",
			"--log-level=1",
			"--window-size=1920,1080",
			"--remote-debugging-port=9222",
			"--disable-dev-shm-usage",
		},
	}
)
