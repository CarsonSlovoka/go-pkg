//go:build windows

package exec

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

// ListenToDelete return a function that start another process by Powershell to delete the file.
// If you want to delete the executable (self), pass os.Args[0] to filepath
//
// 如果您選擇的是刪除自身執行檔，那麼callbackFunc，不應該寫得太複雜，因為結束動作是直接打開powershell運行，
// 因此當您的callback要花很多時間，有可能還沒跑完就刪除了
func ListenToDelete(filePath string) (func(ch chan bool, callbackFunc func(error)), error) {
	var err error
	if filePath, err = filepath.Abs(filePath); err != nil {
		return nil, err
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("filepath not found:%s", filePath)
	}
	return func(ch chan bool, callbackFunc func(error)) {
		select {
		case killNow, _ := <-ch:
			if killNow {
				cmd := exec.Command("powershell", "del", filePath)
				cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				err = cmd.Start()
				callbackFunc(err)
			}
		}
	}, nil
}
