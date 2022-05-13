package exec

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

// GetLocalLangLoc
// Get-Culture List: http://www.codedigest.com/CodeDigest/207-Get-All-Language-Country-Code-List-for-all-Culture-in-C---ASP-Net.aspx
// Chinese (Singapore) zh-SG
// Chinese (People's Republic of China) zh-CN
// Chinese (Hong Kong S.A.R.) zh-HK
// Chinese (Macao S.A.R.) zh-MO
// English (United States) en-US
// Japanese (Japan) ja-JP
func GetLocalLangLoc(defaultLang, defaultLoc string) (string, string) {
	osHost := runtime.GOOS
	switch osHost {
	case "windows":
		// https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.utility/get-culture?view=powershell-7.2
		cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // Hide the window, so you will not see the window flash of PowerShell.
		output, err := cmd.Output()
		if err == nil {
			langLocRaw := strings.TrimSpace(string(output))
			langLoc := strings.Split(langLocRaw, "-")
			return langLoc[0], langLoc[1]
		}
	case "darwin":
		// Exec powershell Get-Culture on Windows.
		cmd := exec.Command("sh", "osascript -e 'user locale of (get system info)'")
		output, err := cmd.Output()
		if err == nil {
			langLocRaw := strings.TrimSpace(string(output))
			langLoc := strings.Split(langLocRaw, "_")
			return langLoc[0], langLoc[1]
		}
	case "linux":
		envLang, ok := os.LookupEnv("LANG")
		if ok {
			langLocRaw := strings.TrimSpace(envLang)
			langLocRaw = strings.Split(envLang, ".")[0]
			langLoc := strings.Split(langLocRaw, "_")
			return langLoc[0], langLoc[1]
		}
	}
	return defaultLang, defaultLoc
}
