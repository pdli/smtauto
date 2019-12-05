package smtauto

import (
	"flag"
	"os"
	"github.com/radutopala/webdriver/chrome"
)

var (
	stackConf = StackConf{
		Version:   "WW47",
		StackPath: "/opt/shares/Navi10_Stack/",
	}

	chromeBinary = flag.String("chrome_binary", "/usr/bin/chromium-browser", "The name of the Chrome binary or the path to it. If name is not an exact path, the PATH will be searched.")
	chrCaps      = chrome.Capabilities{
		Path: *chromeBinary,
		Args: []string{
			//"--headless",
			//"--disable-gpu",
			//"--no-sandbox",
			//"--log-level=1",
			//"--remote-debugging-port=9222",
			//"--disable-dev-shm-usage",
			"--user-data-dir=" + os.Getenv("HOME") + "/.config/chromium/",
		},
	}
)
