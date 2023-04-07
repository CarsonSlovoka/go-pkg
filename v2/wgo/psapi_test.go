package wgo_test

import "log"

func ExampleWGO_FindModuleFileName() {
	targets := []string{"notepad.exe", "cmd.exe"}
	m := wGo.FindModuleFileName(targets...)
	if m == nil {
		return
	}
	for _, target := range targets {
		if pID, exists := m[target]; exists {
			if eno := wGo.KillProcessByPID(pID); eno != 0 {
				log.Println(eno)
			} else {
				log.Printf("close %q successful!\n", target)
			}
		}
	}

	// Output:
}
