package exec

import (
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
