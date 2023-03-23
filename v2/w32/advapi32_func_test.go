package w32_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"log"
	"testing"
	"unsafe"
)

// IsUserAdmin: https://learn.microsoft.com/en-us/windows/win32/api/securitybaseapi/nf-securitybaseapi-checktokenmembership#examples
func ExampleAdApiDLL_CheckTokenMembership() {
	hToken, eno := adApiDll.OpenProcessToken(kernelDll.GetCurrentProcess(), w32.TOKEN_QUERY|w32.TOKEN_DUPLICATE) // 如果少了TOKEN_DUPLICATE，在使用DuplicateToken會出現Access is denied.的錯誤
	if eno != 0 {
		log.Fatal(eno)
	}
	defer func() {
		if eno = kernelDll.CloseHandle(hToken); eno != 0 {
			log.Fatal(eno)
		}
	}()

	// 生成SID
	pSid, eno := adApiDll.AllocateAndInitializeSid(&w32.SECURITY_NT_AUTHORITY, 2,
		w32.SECURITY_BUILTIN_DOMAIN_RID, w32.DOMAIN_ALIAS_RID_ADMINS, 0, 0, 0, 0, 0, 0,
	)
	if eno != 0 {
		log.Fatal(eno)
	}
	sidStr, _ := adApiDll.ConvertSidToStringSid(pSid)
	fmt.Println(sidStr) // S-1-5-32-544

	// pSid, _ := adApiDll.ConvertStringSidToSid("S-1-5-32-544") // 也可以用這種方式生成

	defer func() {
		if r := adApiDll.FreeSid(pSid); r != 0 {
			log.Fatal("FreeSid error")
		}
	}()

	copyHToken, eno := adApiDll.DuplicateToken(hToken, w32.SecurityIdentification)
	if eno != 0 {
		log.Fatal(eno)
	}

	var isUserAdmin bool
	// isUserAdmin, eno = adApiDll.CheckTokenMembership(hToken, pSid) // 如果用原token會遇到錯誤: An attempt has been made to operate on an impersonation token by a thread that is not currently impersonating a client.
	isUserAdmin, eno = adApiDll.CheckTokenMembership(copyHToken, pSid)
	if eno != 0 {
		log.Fatal(eno)
	}

	log.Println(isUserAdmin)
	fmt.Println("ok")
	// Output:
	// S-1-5-32-544
	// ok
}

func ExampleAdApiDLL_ConvertStringSidToSid() {
	sid, eno := adApiDll.ConvertStringSidToSid("S-1-5-32-544")
	if eno != 0 {
		log.Fatal(eno)
	}
	fmt.Println(sid.SubAuthority())

	if _, eno = kernelDll.LocalFree(uintptr(unsafe.Pointer(sid))); eno != 0 {
		log.Println(eno)
	}

	sid2, _ := adApiDll.ConvertStringSidToSid("S-1-5-21-3051027765-3782066248-1388807790-500")
	fmt.Println(sid2.SubAuthority())
	defer func() {
		_, _ = kernelDll.LocalFree(uintptr(unsafe.Pointer(sid2)))
	}()

	var sidStr string
	if sidStr, eno = adApiDll.ConvertSidToStringSid(sid2); eno != 0 {
		log.Fatal(eno)
	}
	fmt.Println(sidStr)

	// Output:
	// [32 544]
	// [21 3051027765 3782066248 1388807790 500]
	// S-1-5-21-3051027765-3782066248-1388807790-500
}

func ExampleAdApiDLL_GetSidSubAuthorityCount() {
	var count byte
	if sid, eno := adApiDll.ConvertStringSidToSid("S-1-5-32-544"); eno == 0 {
		count, eno = adApiDll.GetSidSubAuthorityCount(sid)
		if eno != 0 {
			log.Fatal(eno)
		}
		fmt.Println(count)
	}

	// Output:
	// 2
}

func TestAdApiDLL_IsValidSid(t *testing.T) {
	sid, _ := adApiDll.ConvertStringSidToSid("S-987-5-32-544") // 其實如果不正確，ConvertStringSidToSid他的eno也會和IsValidSid是相同的
	if _, eno := adApiDll.IsValidSid(sid); eno == 0 {
		t.Fatal("should error")
	} else {
		fmt.Println(eno) // The security ID structure is invalid.
	}
}
