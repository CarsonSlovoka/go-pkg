package main

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"syscall"
	"unsafe"
)

func main() {
	// 取得目標進程的主要令牌
	targetProcessHandle, err := syscall.OpenProcess(syscall.PROCESS_QUERY_INFORMATION, false, 1234) // 這邊的1234是進程ID，需替換成實際值
	if err != nil {
		fmt.Println("OpenProcess failed:", err)
		return
	}
	var targetToken syscall.Token
	err = syscall.OpenProcessToken(targetProcessHandle, syscall.TOKEN_DUPLICATE|syscall.TOKEN_QUERY|syscall.TOKEN_ASSIGN_PRIMARY, &targetToken)
	if err != nil {
		fmt.Println("OpenProcessToken failed:", err)
		return
	}
	defer targetToken.Close()

	// 設置受限制的安全描述符
	sa := &syscall.SecurityAttributes{}
	sid, _ := syscall.StringToSid("S-1-5-32-544") // Administrators組的SID
	dacl := &syscall.ACL{}
	dacl.Initialize(1, 1)                                                    // 指定只有一個ACL
	dacl.AddAccessAllowedAce(syscall.ACL_REVISION, syscall.GENERIC_ALL, sid) // 允許Administrators組的所有權限
	sa.SecurityDescriptor = &syscall.SecurityDescriptor{}
	sa.SecurityDescriptor.SetSecurityDescriptorDacl(true, dacl, false)

	var acl w32.AccessAllowedAce
	adApiDll.AddAccessAllowedAce(&acl, syscall.ACL_REVISION)

	// 創建受限制的令牌
	var restrictedToken syscall.Token
	err = syscall.CreateRestrictedToken(targetToken, syscall.DISABLE_MAX_PRIVILEGE, 0, nil, 0, nil, 1, &syscall.RestrictedSidsAndAttributes{Sid: sid, Attributes: 0}, 0, nil, &restrictedToken)
	if err != nil {
		fmt.Println("CreateRestrictedToken failed:", err)
		return
	}
	defer restrictedToken.Close()

	// 在新進程中啟動程序
	cmd := "C:\\Windows\\System32\\notepad.exe" // 可替換為實際的執行檔路徑
	var si syscall.StartupInfo
	si.Cb = uint32(unsafe.Sizeof(si))
	si.Flags = syscall.STARTF_USESHOWWINDOW
	si.ShowWindow = syscall.SW_SHOWNORMAL
	err = syscall.CreateProcessAsUser(restrictedToken, nil, &syscall.StringToUTF16(cmd)[0], nil, nil, false, syscall.CREATE_UNICODE_ENVIRONMENT, nil, nil, &si, &procInfo)
	if err != nil {
		fmt.Println("CreateProcessAsUser failed:", err)
		return
	}
	defer syscall.CloseHandle(procInfo.Process)
}
