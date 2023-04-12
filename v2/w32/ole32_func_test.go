// go test -v -run=Test_openIE
package w32_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"github.com/CarsonSlovoka/go-pkg/v2/wgo"
	"log"
	"sync"
	"syscall"
	"testing"
	"time"
	"unsafe"
)

func ExampleGUID_String() {
	var guid *w32.GUID
	guid = w32.IID_IUnknown
	fmt.Printf("%s\n", guid.String())
	// Output:
	// {00000000-0000-0000-C000-000000000046}
}

// 開啟IE瀏覽器，找到搜尋欄位，鍵入golang
func Test_openIE(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	w32.SetPreferLCID(w32.LOCALE_SYSTEM_DEFAULT)
	go func() { // Avoid: all goroutines are asleep - deadlock!
		defer wg.Done()

		log.Println(oleDll.CoInitializeEx(0, w32.COINIT_MULTITHREADED))
		defer oleDll.CoUnInitialize()
		unknown, errno := w32.NewIUnknownInstance(oleDll, w32.CLSID_InternetExplorer, w32.CLSCTX_SERVER)
		if errno != 0 {
			log.Printf("%s\n", errno)
			return
		}
		defer unknown.Release()

		var dispatchIE *w32.IDispatch
		dispatchIE, errno = unknown.QueryInterface(w32.IID_IDispatch)
		if errno != 0 {
			log.Printf("%s\n", errno)
			return
		}

		if _, errno = dispatchIE.PropertyPut("Visible", true); errno != 0 {
			log.Printf("%s\n", errno)
			return
		}
		dispatchIE.MustMethod("Navigate", "http://www.google.com")

		for {
			if dispatchIE.MustPropertyGet("Busy").Val == 0 {
				break
			}
		}
		time.Sleep(time.Second) // 太快執行，下面的內容會報錯
		document := dispatchIE.MustPropertyGet("Document").ToIDispatch()
		htmlElems := document.MustMethod("getElementsByName", "q").ToIDispatch()
		htmlInputElem := htmlElems.MustMethod("item", 0).ToIDispatch()
		htmlInputElem.MustPropertyPut("value", "golang")
	}()
	wg.Wait()
}

// 開啟IE瀏覽器，找到搜尋欄位，鍵入golang
// 此範例不使用封裝的方法所實現，會顯得囉嗦很多
func Test_openIE2(t *testing.T) {
	oleDll.CoInitialize(0)
	defer oleDll.CoUnInitialize()

	var clsID w32.GUID
	// https://learn.microsoft.com/en-us/previous-versions/windows/internet-explorer/ie-developer/platform-apis/aa752084(v=vs.85)
	if errno := oleDll.CLSIDFromProgID("InternetExplorer.Application", &clsID); errno != 0 {
		log.Printf("%s\n", errno)
		return
	}
	// {0002DF01-0000-0000-C000-000000000046}
	log.Printf("%s\n", clsID.String())

	unknown, errno := oleDll.CoCreateInstance(
		&clsID,            // w32.CLSID_InternetExplorer,
		w32.CLSCTX_SERVER, // w32.CLSCTX_INPROC_HANDLER, 這個要註冊class才有辦法使用
		w32.IID_IUnknown,
	)
	if errno != 0 {
		log.Fatalf("%s\n", errno)
	}
	defer unknown.Release()

	dispatchIE, errno2 := unknown.QueryInterface(w32.IID_IDispatch)
	if errno2 != 0 {
		log.Printf("%s\n", errno2)
	}
	lcid := kernelDll.GetUserDefaultLCID()

	dispIDs := make([]w32.DISPID, 4)
	for i, name := range []string{
		"Visible",
		"Navigate",
		"Busy",
		"document",
	} {
		tmpDispIDs, errno3 := dispatchIE.GetIDsOfNames(nil,
			[]string{
				name,
				// 不能同時取多個，會報錯
				// "Visible",
			},
			0, lcid)
		if errno3 != 0 {
			log.Fatalf("%s", errno3)
		}
		dispIDs[i] = tmpDispIDs[0]
	}

	dispIDVisible := dispIDs[0]
	nameArgs := [1]w32.DISPID{w32.DISPID_PROPERTYPUT}
	var dispParams *w32.DispParams
	{
		dispParams = &w32.DispParams{
			// 只有DISPATCH_PROPERTYPUT, DISPATCH_PROPERTYPUTREF要額外設定Name相關內容
			NamedArgs:  uintptr(unsafe.Pointer(&nameArgs[0])),
			CNamedArgs: 1,
		}
		params := []any{true} // 方法(這裡的method指的是Visible)所用到的參數
		vargs := make([]w32.VARIANT, len(params))
		oleAutDll.VariantInit(&vargs[0])
		vargs[0] = w32.NewVariant(w32.VT_BOOL, 0xffff /* true */)
		dispParams.CArgs = uint32(len(params))
		dispParams.VArgs = uintptr(unsafe.Pointer(&vargs[0]))
	}

	var exceptInfo w32.EXCEPINFO
	if _, errno = dispatchIE.Invoke(dispIDVisible, w32.IID_NULL, lcid, w32.DISPATCH_PROPERTYPUT, dispParams, &exceptInfo, nil); errno != 0 {
		log.Fatalf("%s\n", errno)
	}

	dispIDNavigate := dispIDs[1]
	nameArgs = [1]w32.DISPID{} // DISPATCH_METHOD 不需要設定
	dispParams = &w32.DispParams{
		// RgdispidNamedArgs: uintptr(unsafe.Pointer(&nameArgs[0])),
		// CNamedArgs:        0,
	}
	params := []any{"http://www.google.com"}
	dispParams.CArgs = uint32(len(params))
	vargs := make([]w32.VARIANT, len(params))
	oleAutDll.VariantInit(&vargs[0])
	vargs[0] = w32.NewVariant(w32.VT_BSTR, int64(uintptr(unsafe.Pointer(oleAutDll.SysAllocStringLen(params[0].(string))))))
	dispParams.VArgs = uintptr(unsafe.Pointer(&vargs[0]))
	if _, errno := dispatchIE.Invoke(dispIDNavigate, w32.IID_NULL, lcid, w32.DISPATCH_METHOD, dispParams, &exceptInfo, nil); errno != 0 {
		log.Printf("%s\n", errno)
	}
	oleAutDll.SysFreeString((*uint16)(unsafe.Pointer(uintptr(vargs[0].Val))))

	dispIDBusy := dispIDs[2]
	nameArgs = [1]w32.DISPID{}
	dispParams = &w32.DispParams{
		// NamedArgs: uintptr(unsafe.Pointer(&nameArgs[0])),
		// CNamedArgs:        0,
		VArgs: 0,
		CArgs: 0, // 沒有參數
	}
	var result *w32.VARIANT
	for {
		if result, errno = dispatchIE.Invoke(dispIDBusy, w32.IID_NULL, lcid, w32.DISPATCH_PROPERTYGET, dispParams, &exceptInfo, nil); errno != 0 {
			log.Printf("%s\n", errno)
			return
		}
		if result.Val == 0 {
			break
		}
		result.ToIDispatch()
	}
	time.Sleep(time.Second) // 太快執行，下面的內容會報錯

	dispIDDocument := dispIDs[3]
	dispParams = &w32.DispParams{} // 都沒有數值需要填還是需要指向一個空內容，不然會錯
	var document *w32.VARIANT
	oleAutDll.VariantInit(document)
	if document, errno = dispatchIE.Invoke(dispIDDocument, w32.IID_NULL, lcid, w32.DISPATCH_PROPERTYGET, dispParams, &exceptInfo, nil); errno != 0 {
		log.Printf("%s\n", errno)
		return
	}
	dispDocument := document.ToIDispatch()

	dispIDs, _ = dispDocument.GetIDsOfNames(nil, []string{"getElementsByName"}, 1, lcid)
	dispIDGetElementByName := dispIDs[0]
	params = []any{"q"} // query
	dispParams = &w32.DispParams{}
	dispParams.CArgs = uint32(len(params))
	vargs = make([]w32.VARIANT, len(params))
	oleAutDll.VariantInit(&vargs[0])
	vargs[0] = w32.NewVariant(w32.VT_BSTR, int64(uintptr(unsafe.Pointer(oleAutDll.SysAllocStringLen(params[0].(string))))))
	dispParams.VArgs = uintptr(unsafe.Pointer(&vargs[0]))

	var htmlElem *w32.VARIANT
	if htmlElem, errno = dispDocument.Invoke(dispIDGetElementByName, w32.IID_NULL, lcid, w32.DISPATCH_METHOD, dispParams, &exceptInfo, nil); errno != 0 {
		log.Printf("%s\n", errno)
		return
	}
	dispatchHtmlElem := htmlElem.ToIDispatch()

	dispIDs, _ = dispatchHtmlElem.GetIDsOfNames(nil, []string{"item"}, 1, lcid)
	dispIDItem := dispIDs[0]
	params = []any{0} // query
	dispParams = &w32.DispParams{}
	vargs = make([]w32.VARIANT, len(params))
	oleAutDll.VariantInit(&vargs[0])
	vargs[0] = w32.NewVariant(w32.VT_I4, int64(params[0].(int)))
	dispParams.CArgs = uint32(len(params))
	dispParams.VArgs = uintptr(unsafe.Pointer(&vargs[0]))

	var query *w32.VARIANT
	if query, errno = dispatchHtmlElem.Invoke(dispIDItem, w32.IID_NULL, lcid, w32.DISPATCH_METHOD, dispParams, &exceptInfo, nil); errno != 0 {
		log.Printf("%s\n", errno)
		return
	}
	dispatchQuery := query.ToIDispatch()

	dispIDs, _ = dispatchQuery.GetIDsOfNames(nil, []string{"value"}, 1, lcid)
	dispIDValue := dispIDs[0]
	nameArgs = [1]w32.DISPID{w32.DISPID_PROPERTYPUT}
	params = []any{"golang"} // query
	dispParams = &w32.DispParams{
		NamedArgs:  uintptr(unsafe.Pointer(&nameArgs[0])),
		CNamedArgs: 1,
	}
	dispParams.CArgs = uint32(len(params))
	vargs = make([]w32.VARIANT, len(params))
	oleAutDll.VariantInit(&vargs[0])
	vargs[0] = w32.NewVariant(w32.VT_BSTR, int64(uintptr(unsafe.Pointer(oleAutDll.SysAllocStringLen(params[0].(string)))))) // 字串都要透過SpyFreeString來釋放記憶體
	dispParams.VArgs = uintptr(unsafe.Pointer(&vargs[0]))

	if _, errno := dispatchQuery.Invoke(dispIDValue, w32.IID_NULL, lcid, w32.DISPATCH_PROPERTYPUT, dispParams, &exceptInfo, nil); errno != 0 {
		log.Printf("%s\n", errno)
		return
	}
	oleAutDll.SysFreeString((*uint16)(unsafe.Pointer(uintptr(vargs[0].Val))))
}

func ExampleOle32DLL_CLSIDFromProgID() {
	oleDll.CoInitialize(0)
	defer oleDll.CoUnInitialize()
	clsID := new(w32.GUID)
	if eno := oleDll.CLSIDFromProgID("Excel.Application", clsID); eno != 0 { // HKEY_LOCAL_MACHINE\SOFTWARE\Classes\Excel.Application\CLSID
		log.Fatal(eno)
		return
	}
	log.Println(clsID.String()) // {00024500-0000-0000-C000-000000000046}

	{
	}

	// Output:
}

func ExampleOle32DLL_CLSIDFromString() {
	oleDll.CoInitialize(0)
	defer oleDll.CoUnInitialize()
	clsID := new(w32.GUID)
	if eno := oleDll.CLSIDFromString("{50AC103F-D235-4598-BBEF-98FE4D1A3AD4}", clsID); eno != 0 {
		log.Fatal(eno)
	}
	fmt.Println(clsID.String())
	fmt.Println(w32.NewGUID("{50AC103F-D235-4598-BBEF-98FE4D1A3AD4}").String())

	// Output:
	// {50AC103F-D235-4598-BBEF-98FE4D1A3AD4}
	// {50AC103F-D235-4598-BBEF-98FE4D1A3AD4}
}

func Test_excel(t *testing.T) {
	oleDll.CoInitialize(0)
	defer oleDll.CoUnInitialize()

	wGo := wgo.NewWGO(kernelDll)

	// 紀錄已經開啟的excel程序
	orgOpenExcelEntrySlice, _ := wGo.GetProcessEntry(func(entry *w32.PROCESSENTRY32W) bool {
		return entry.ExeFileName() == "EXCEL.EXE"
	})

	unknown, errno := w32.NewIUnknownInstance(oleDll, "Excel.Application", w32.CLSCTX_SERVER)
	if errno != 0 {
		log.Printf("%s", errno)
		return
	}
	defer unknown.Release()

	excel, _ := unknown.QueryInterface(w32.IID_IDispatch)

	// 刪除退出之後還沒有被關掉的EXCEL程序
	defer func() {
		needDeleteExcels, _ := wGo.GetProcessEntry(func(entry *w32.PROCESSENTRY32W) bool {
			if entry.ExeFileName() != "EXCEL.EXE" {
				return false
			}
			curPID := entry.Th32ProcessID
			for _, orgEntry := range orgOpenExcelEntrySlice {
				if orgEntry.Th32ProcessID == curPID {
					return false
				}
			}
			return true
		})

		if len(needDeleteExcels) > 0 {
			wGo.KillProcess(needDeleteExcels, func(entry *w32.PROCESSENTRY32W, errno syscall.Errno) {
				if errno != 0 {
					log.Printf("[Kill Process Error--Excel] processID: %d", entry.Th32ProcessID)
				}
			})
		}
	}()

	defer excel.Release()
	_, _ = excel.PropertyPut("Visible", false) // 這種狀態還是能用，只是在背景執行
	workbooks := excel.MustPropertyGet("Workbooks").ToIDispatch()
	workbook := workbooks.MustMethod("Add", nil).ToIDispatch()

	// https://learn.microsoft.com/en-us/office/vba/api/excel.worksheet
	worksheet := workbook.MustPropertyGet("Worksheets", 1).ToIDispatch() // 表示最左邊的工作表

	pageSetup := worksheet.MustPropertyGet("PageSetup").ToIDispatch()
	pageSetup.MustPropertyPut("FitToPagesWide", 5)

	var (
		vCell *w32.VARIANT
		i, j  int
	)
	for i = 1; i < 10; i++ {
		for j = 1; j < 10; j++ {
			// https://learn.microsoft.com/en-us/office/vba/api/excel.worksheet.cells
			vCell, errno = worksheet.PropertyGet("Cells", i, j) // 注意下標是1開始，給0會報錯
			if i == 1 {
				// range.Property_Font.Property_Size = 38
				// vCell.ToIDispatch().MustPropertyGet("Font").ToIDispatch().MustPropertyPut("Size", 38) // 如果只要設定大小，可以一次寫完

				// Font: https://learn.microsoft.com/en-us/office/vba/api/excel.font(object)
				// Color https://learn.microsoft.com/en-us/office/vba/api/excel.font.color
				//  	從color連結可以訪問: https://learn.microsoft.com/en-us/office/vba/language/reference/user-interface-help/rgb-function
				// 		會得知RGB函數要的是一個LONG的型別: Returns a Long whole number representing an RGB color value.
				// Data type summary: https://learn.microsoft.com/en-us/office/vba/language/reference/user-interface-help/data-type-summary
				// 可以得知LONG是一個4byte範圍從-2,147,483,648 to 2,147,483,647
				font := vCell.ToIDispatch().MustPropertyGet("Font").ToIDispatch()
				font.MustPropertyPut("Size", 38)

				font.MustPropertyPut("Color", int32(w32.RGB(0xff, 0xff, 0x00))) // yellow
				// font.MustPropertyPut("Color", int32((0x00<<16)|(0xff<<8)|0xff)) // 同上, bgr
			}
			if errno != 0 {
				log.Printf("%s", errno)
				return
			}
			vCell.ToIDispatch().MustPropertyPut("Value", i*j)
		}
	}

	// range
	{
		// 只能橫著一列一列寫
		rg := worksheet.MustPropertyGet("Range", "A10:C10").ToIDispatch()
		rg.MustPropertyPut("Value", []string{"aa", "2", "c"})

		/* 直的寫入會有問題
		rg2 := worksheet.MustPropertyGet("Range", "E10:F12").ToIDispatch()
		rg2.MustPropertyPut("Value", []string{"a", "b", "c", "d"})
		// output:
		// a b
		// a b

		rg := worksheet.MustPropertyGet("Range", "C10:C12").ToIDispatch()
		rg.MustPropertyPut("Value", []string{"1", "2", "c"})
		output:
		// 1
		// 1
		// 1
		*/
	}

	_, _ = workbook.PropertyPut("Saved", true) // 儲存異動結果(非存檔)
	// workbook.MustMethod("SaveAs", "C:\\myDir\\out.xlsx") // 注意！是用\\而不是/
	_, _ = workbook.Method("Closed", false)
	excel.MustMethod("Quit")
}

// https://learn.microsoft.com/en-us/windows/win32/wmisdk/scripting-api-for-wmi
// Powershell:
//
//	Get-WmiObject -Class Win32_NetworkAdapterConfiguration  | Format-Table
//	Get-WmiObject -Class Win32_NetworkAdapterConfiguration -Filter IPEnabled=TRUE | Format-Table
//	Get-WmiObject -Class Win32_NetworkAdapterConfiguration -Filter IPEnabled=TRUE | Format-Table -Property IPAddress
func Test_wmi(t *testing.T) {
	oleDll.CoInitialize(0)
	defer oleDll.CoUnInitialize()

	// https://learn.microsoft.com/en-us/windows/win32/wmisdk/document-conventions-for-the-scripting-api
	// HKEY_LOCAL_MACHINE\SOFTWARE\Classes\WbemScripting.SWbemLocator
	unknown, errno := w32.NewIUnknownInstance(oleDll, "WbemScripting.SWbemLocator", w32.CLSCTX_SERVER)
	if errno != 0 {
		return
	}
	defer unknown.Release()

	wmi, _ := unknown.QueryInterface(w32.IID_IDispatch)

	// https://learn.microsoft.com/en-us/windows/win32/wmisdk/swbemlocator#methods
	// return: https://learn.microsoft.com/en-us/windows/win32/wmisdk/swbemlocator-connectserver#return-value
	service := wmi.MustMethod("ConnectServer").ToIDispatch()
	defer service.Release()

	// https://learn.microsoft.com/en-us/windows/win32/wmisdk/swbemservices
	// https://learn.microsoft.com/en-us/windows/win32/wmisdk/swbemservices-execquery
	sWbemObjectSet := service.MustMethod("ExecQuery", "SELECT * FROM Win32_NetworkAdapterConfiguration").ToIDispatch()
	defer sWbemObjectSet.Release()

	// https://learn.microsoft.com/en-us/windows/win32/wmisdk/swbemobjectset#properties
	countVar := sWbemObjectSet.MustPropertyGet("Count")
	count := int(countVar.Val)
	fmt.Println(count)
	for i := 0; i < count; i++ {
		// item is a SWbemObject, but really a Win32_Process
		// https://learn.microsoft.com/en-us/windows/win32/wmisdk/swbemobjectset-itemindex
		item := sWbemObjectSet.MustMethod("ItemIndex", i).ToIDispatch()

		// https://learn.microsoft.com/en-us/windows/win32/cimwin32prov/win32-networkadapterconfiguration#syntax
		addr, errno := item.PropertyGet("IPAddress")
		if errno != 0 {
			log.Println(errno.Error())
		} else {
			arr := addr.ToArray()
			if arr != nil {
				// fmt.Println(arr.ToStringArray()) // TODO
			}
		}
		fmt.Println(item.MustPropertyGet("Description").ToString())
		item.Release()
	}
}

func ExampleOle32DLL_CoCreateGuid() {
	// 創建三組GUID，可以發現他們所產生的內容都不一樣
	for i := 0; i < 3; i++ {
		var guid w32.GUID
		if eno := oleDll.CoCreateGuid(&guid); eno != 0 {
			log.Println(syscall.Errno(eno))
		}
		// {46997087-EADB-4C19-844E-D97A5B41D892}
		log.Println(guid.String())
	}
	// Output:
}

func ExampleOle32DLL_CoTaskMemFree() {
	oleDll.CoTaskMemFree(unsafe.Pointer(nil))
	// Output:
}
