package exec

import (
	"fmt"
	"os"
	"os/exec"
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

// debug運行可能會刪不掉，要變成執行檔才行
func testListenToDeleteApp(t *testing.T) {
	chanKill := make(chan bool)
	chanQuit := make(chan bool)
	targetExePath := "temp.txt"
	f, _ := os.Create(targetExePath)
	f.Close()

	listenToDelFunc, err := ListenToDelete(targetExePath) // os.Args[0]
	if err != nil {
		t.Fatal(err)
	}
	go listenToDelFunc(chanKill, func(err error) {
		if err != nil {
			fmt.Println(err)
		}
		close(chanQuit)
	})

	go func() {
		needDelete := true
		if needDelete {
			chanKill <- true
			return
		}
		close(chanQuit)
	}()

	select {
	case <-chanQuit:
		return
	}
}
