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

// Deprecated: I do not recommend that you use this method. Use "CreateMutex" instead, see below,
// https://github.com/CarsonSlovoka/go-pkg/blob/4b6ee040d9e5d9831740d20918e992a260594e80/v2/w32/kernel32_func_test.go#L9-L35
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

// ListenToDelete Send the file path which you want to delete to the channel, and then it will start another process by Powershell to delete the file.
// If you want to delete the executable (self), you can send os.Args[0] "ch"
// 正常而言您不需要透過此函數來刪除檔案，僅需要`os.Remove()`即可。
// 此函數的特色:能刪除掉自身執行檔
// 如果您選擇的是刪除自身執行檔，那麼callbackFunc，不應該寫得太複雜，因為結束動作是直接打開powershell運行，
// 因此當您的callback要花很多時間，有可能還沒跑完就刪除了
func ListenToDelete(ch chan string, cbFunc func(string, error)) {
	var err error
	for {
		select {
		case filePath, isOpen := <-ch:
			if !isOpen {
				// fmt.Println("Stop listen")
				return
			}
			if filePath, err = filepath.Abs(filePath); err != nil {
				cbFunc(filePath, err)
				continue
			}
			if _, err = os.Stat(filePath); os.IsNotExist(err) {
				cbFunc(filePath, fmt.Errorf("filepath not found:%s", filePath))
				continue
			}
			cmd := exec.Command("powershell", "del", filePath)
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			cbFunc(filePath, cmd.Start())
		}
	}
}

func CmdWithoutWindow(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}
