//go:build windows

package exec

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

func countTask(exeName string) int {
	cmdTimer := exec.Command("TASKLIST", "/FI", fmt.Sprintf("IMAGENAME eq %s", exeName))
	cmdTimer.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	rtnBytes, _ := cmdTimer.Output()
	return strings.Count(strings.ToUpper(string(rtnBytes)), fmt.Sprintf("%s", strings.ToUpper(exeName)))
}

func IsSingleInstance(curExeName string) bool {
	if countTask(curExeName) > 1 { // 當前啟動的程式會佔用一個，所以要大於1個才是真的有重複
		return false
	}
	return true
}

func IsTaskRunning(exeName string) bool {
	if countTask(exeName) >= 1 {
		return true
	}
	return false
}

// TaskKill 強制關閉程式, 可透過tasklist查看task清單
func TaskKill(exeName string) error {
	if IsTaskRunning(exeName) {
		cmd := exec.Command("taskkill",
			"/IM", exeName,
			"/F")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		return cmd.Run()
	}
	return nil
}
