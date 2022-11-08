package exec

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"testing"
	"time"
)

func ExampleTaskKill() {
	testApp := "powershell.exe"                           // 如果用taskmgr.exe(Task manager)會需要提升管理權限才有辦法運行，否則會有錯誤: The requested operation requires elevation.
	if err := exec.Command(testApp).Start(); err != nil { // please run with admin
		panic(err) // The requested operation requires elevation.
	}
	time.Sleep(250 * time.Millisecond)
	if err := TaskKill(testApp); err != nil {
		panic(err)
	}
	if IsTaskRunning(testApp) {
		panic("The program was still alive.")
	}

	// Output:
}

func ExampleIsTaskRunning() {
	testApp := "taskmgr.exe"                              // Task manager
	if err := exec.Command(testApp).Start(); err != nil { // please run with admin
		panic(err) // The requested operation requires elevation.
	}
	time.Sleep(250 * time.Millisecond)
	if err := TaskKill(testApp); err != nil {
		panic(err)
	}
	if IsTaskRunning(testApp) {
		panic("The program was still alive.")
	}
}

func TestIsSingleInstance(t *testing.T) {
	IsSingleInstance("notepad.exe")
}

func ExampleIsSingleInstance() {
	testApp := "notepad.exe"
	if err := exec.Command(testApp).Start(); err != nil {
		panic(err)
	}
	if !IsSingleInstance(testApp) {
		panic("假設notepad的應用程式對象只有一個，那麼不該有錯誤")
	}

	// 建立第二個notepad
	if err := exec.Command(testApp).Start(); err != nil { // run again
		panic(err)
	}
	if IsSingleInstance(testApp) {
		panic("此時有兩個對象，所以InSingleInstance應該為false")
	}
}

// Delete Self Exe
func ExampleListenToDelete() {
	chanRemoveFile := make(chan string)
	chanQuit := make(chan bool)
	go ListenToDelete(chanRemoveFile, func(curFile string, err error) {
		defer close(chanQuit)
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Printf("remove the file successful: %s\n", curFile) // 由於我們用start去運行，因此沒有錯誤只是代表呼叫成功，但不意味已經刪除該檔案了
		fmt.Printf("call the del command successful: %s\n", curFile)
	})
	chanRemoveFile <- os.Args[0]
	select {
	case <-chanQuit:
		return
	}
}

func ExampleListenToDelete_multiple() {
	chanRemoveFile := make(chan string)
	chanQuit := make(chan bool)
	deleteFiles := []string{
		"temp.txt",
		"temp.exe",
	}
	for _, testFilepath := range deleteFiles {
		f, _ := os.Create(testFilepath)
		if err := f.Close(); err != nil {
			panic(err)
		}
	}
	time.Sleep(time.Second) // just let you see the file created.

	wg := new(sync.WaitGroup)
	wg.Add(len(deleteFiles))
	go ListenToDelete(chanRemoveFile, func(curFile string, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("call the del command successful: %s\n", curFile)
	})

	go func() {
		var curFile string
		for {
			if len(deleteFiles) == 0 {
				close(chanRemoveFile)
				close(chanQuit)
				break
			}
			curFile, deleteFiles = deleteFiles[0], deleteFiles[1:] // pop
			chanRemoveFile <- curFile
		}
	}()

	select {
	case <-chanQuit:
		wg.Wait()
		time.Sleep(time.Second) // goland的debug貌似主程序結束，powershell的呼叫也會被中斷，所以要等待，如果是打包出去的執行檔，不需要特別加上sleep
		for _, filePath := range []string{
			"temp.txt",
			"temp.exe",
		} {
			if _, err := os.Stat(filePath); !os.IsNotExist(err) {
				panic("the file still exists, not deleted. " + err.Error())
			}
		}
		return
	}
}

func ExampleCmdWithoutWindow() {
	cmd := CmdWithoutWindow("powershell", fmt.Sprintf("Get-FileHash %s -Algorithm md5 | select Hash,Path", os.Args[0]))
	rtnBytes, err := cmd.Output()
	if err != nil {
		return // <-- github action 會無法執行成功，為了避免影響測試，故不進行錯誤處理
	}
	log.Println(string(rtnBytes))

	// Output:
}
