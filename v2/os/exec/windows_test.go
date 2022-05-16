package exec

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"testing"
	"time"
)

func TestTaskKill(t *testing.T) {
	testApp := "taskmgr.exe"                              // Task manager
	if err := exec.Command(testApp).Start(); err != nil { // please run with admin
		t.Fatalf(err.Error()) // The requested operation requires elevation.
	}
	time.Sleep(250 * time.Millisecond)
	if err := TaskKill(testApp); err != nil {
		t.Fatalf(err.Error())
	}
	if IsTaskRunning(testApp) {
		t.Fatalf("The program was still alive.")
	}
}

func TestIsSingleInstance(t *testing.T) {
	testApp := "notepad.exe"
	if err := exec.Command(testApp).Start(); err != nil {
		t.Fatalf(err.Error())
	}
	if !IsSingleInstance(testApp) {
		t.FailNow()
	}
	if err := exec.Command(testApp).Start(); err != nil { // run again
		t.Fatalf(err.Error())
	}
	if IsSingleInstance(testApp) {
		t.FailNow()
	}
}

func TestListenToDeleteSelfExe(t *testing.T) {
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

func TestListenToDeleteMultipleFile(t *testing.T) {
	chanRemoveFile := make(chan string)
	chanQuit := make(chan bool)
	deleteFiles := []string{
		"temp.txt",
		"temp.exe",
	}
	for _, testFilepath := range deleteFiles {
		f, _ := os.Create(testFilepath)
		if err := f.Close(); err != nil {
			t.Fatal(err)
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
				t.Fatal("the file still exists, not deleted.")
			}
		}
		return
	}
}
