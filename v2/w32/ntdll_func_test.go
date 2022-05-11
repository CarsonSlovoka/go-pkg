package w32_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"testing"
)

// Operating System Version: https://docs.microsoft.com/en-us/windows/win32/sysinfo/operating-system-version?redirectedfrom=MSDN
// https://en.wikipedia.org/wiki/List_of_Microsoft_Windows_versions (看不出Major, Minor
// Version Number: https://www.lifewire.com/windows-version-numbers-2625171
// 所謂的 21H2 指的是 20"21"
// H1, H2指的是上半年或者是下半年所發布
// 10.0.19044.1202
// Major = 10
// Minor = 0
// Build = 19044
// Rev = 1202 (修訂版)
func TestGetVersion(t *testing.T) {

	ntDll := w32.NewNtDLL([]w32.ProcName{
		w32.PNRtlGetVersion,
	})
	user32Dll := w32.NewUser32DLL([]w32.ProcName{
		w32.PNGetSystemMetrics,
	})
	kernel32Dll := w32.NewKernel32DLL([]w32.ProcName{
		w32.PNGetNativeSystemInfo,
	})

	osV := ntDll.RtlGetVersion()
	fmt.Printf("%+v\n", osV)

	versionID := fmt.Sprintf("%d_%d_%d", osV.MajorVersion, osV.MinorVersion, osV.BuildNumber)

	// https://docs.microsoft.com/en-us/windows/win32/sysinfo/operating-system-version?redirectedfrom=MSDN
	// https://en.wikipedia.org/wiki/List_of_Microsoft_Windows_versions 其實這個連結隱約可以知道major和minor的版號，一直到Windows 10 version 1507之後才無法得知
	OsNameMap := map[string]string{
		"10_0_22000": "Windows 11",

		"10_0_19044": "Windows 10",
		"10_0_19043": "Windows 10",
		"10_0_19042": "Windows 10",
		"10_0_19041": "Windows 10",
		"10_0_18362": "Windows 10",
		"10_0_18363": "Windows 10",
		"10_0_17763": "Windows 10",
		"10_0_17134": "Windows 10",
		"10_0_16299": "Windows 10",
		"10_0_15063": "Windows 10",
		"10_0_14393": "Windows 10",
		"10_0_10586": "Windows 10",
		"10_0_10240": "Windows 10",

		"6_3_9600": "Windows 8.1",

		"6_2_9200": "Windows 8",

		"6_1_7601": "Windows 7",

		"6_0_6002": "Windows Vista",

		"5_2_3790": "Windows XP 64-Bit Edition", // 可能要檢驗是否為64位元版本，不過因為剩下的5.2的版本都是Server，所以排除之後就可以確定是這個

		"5_1_3790": "Windows XP",
		"5_1_2710": "Windows XP",
		"5_1_2700": "Windows XP",
		"5_1_2600": "Windows XP",

		"5_0_2195": "Windows 2000",
	}

	// https://docs.microsoft.com/en-us/windows-hardware/drivers/ddi/wdm/ns-wdm-_osversioninfoexw#remarks
	OsServerNameMap := map[string]string{
		"10_0_20348": "Windows Server 2022",
		"10_0_17763": "Windows Server 2019",
		"10_0_14393": "Windows Server 2016",

		"6_3_9600": "Windows Server 2012 R2",
		"6_2_9200": "Windows Server 2012",
		"6_1_7601": "Windows Server 2008 R2",
		"6_0_6003": "Windows Server 2008",

		// 5.2是特例，另外判斷
		"5_2_3790_":                "Windows Server 2003",
		"5_2_3790_r2":              "Windows Server 2003 R2",
		"5_2_3790_x64":             "Windows XP Professional x64 Edition", // 這個雖然沒寫Server但確實是server
		"5_2_3790_suite_wh_server": "Windows Home Server",
	}

	versionName := ""
	if osV.ProductType == w32.VER_NT_WORKSTATION || osV.SuiteMask != w32.VER_SUITE_WH_SERVER {
		if name, exists := OsNameMap[versionID]; exists {
			versionName = name
		}
	} else { // Server
		if name, exists := OsServerNameMap[versionID]; exists {
			versionName = name
		} else if osV.MajorVersion == 5 && osV.MinorVersion == 2 {
			arch := kernel32Dll.GetNativeSystemInfo().ProcessorArchitecture
			is64arch := arch == w32.PROCESSOR_ARCHITECTURE_AMD64 || arch == w32.PROCESSOR_ARCHITECTURE_ARM64 || arch == w32.PROCESSOR_ARCHITECTURE_IA64
			if osV.ProductType == w32.VER_NT_WORKSTATION && is64arch {
				versionName = "Windows XP Professional x64 Edition"
			} else if osV.SuiteMask == w32.VER_SUITE_WH_SERVER {
				versionName = "Windows Home Server"
			} else if user32Dll.GetSystemMetrics(w32.SM_SERVERR2) != 0 { // The build number if the system is Windows Server 2003 R2; otherwise, 0.
				versionName = "Windows Server 2003 R2"
			} else {
				versionName = "Windows Server 2003"
			}
		}
	}
	fmt.Println(versionName)
}
