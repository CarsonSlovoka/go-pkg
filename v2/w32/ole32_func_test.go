// go test -v -run=Test_openIE
package w32_test

import (
	"fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/w32"
	"log"
	"sync"
	"testing"
	"time"
	"unsafe"
)

var (
	ole    *w32.Ole32DLL
	kernel *w32.Kernel32DLL
	oleAut *w32.OleAut32DLL
)

func init() {
	ole = w32.NewOle32DLL()
	kernel = w32.NewKernel32DLL()
	oleAut = w32.OleAutDll
}

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

		log.Println(ole.CoInitializeEx(0, w32.COINIT_MULTITHREADED))
		defer ole.CoUnInitialize()
		unknown, errno := w32.NewIUnknownInstance(ole, w32.CLSID_InternetExplorer, w32.CLSCTX_SERVER)
		if errno != 0 {
			log.Printf("%s\n", errno)
			return
		}

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
		wg.Done()
	}()
	wg.Wait()

}

// 開啟IE瀏覽器，找到搜尋欄位，鍵入golang
// 此範例不使用封裝的方法所實現，會顯得囉嗦很多
func Test_openIE2(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUnInitialize()

	var clsID w32.GUID
	// https://learn.microsoft.com/en-us/previous-versions/windows/internet-explorer/ie-developer/platform-apis/aa752084(v=vs.85)
	if errno := ole.CLSIDFromProgID("InternetExplorer.Application", &clsID); errno != 0 {
		log.Printf("%s\n", errno)
		return
	}
	// {0002DF01-0000-0000-C000-000000000046}
	log.Printf("%s\n", clsID.String())

	unknown, errno := ole.CoCreateInstance(
		&clsID,            // w32.CLSID_InternetExplorer,
		w32.CLSCTX_SERVER, // w32.CLSCTX_INPROC_HANDLER, 這個要註冊class才有辦法使用
		w32.IID_IUnknown,
	)
	if errno != 0 {
		log.Fatalf("%s\n", errno)
	}

	dispatchIE, errno2 := unknown.QueryInterface(w32.IID_IDispatch)
	if errno2 != 0 {
		log.Printf("%s\n", errno2)
	}
	lcid := kernel.GetUserDefaultLCID()

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
		oleAut.VariantInit(&vargs[0])
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
	oleAut.VariantInit(&vargs[0])
	vargs[0] = w32.NewVariant(w32.VT_BSTR, int64(uintptr(unsafe.Pointer(oleAut.SysAllocStringLen(params[0].(string))))))
	dispParams.VArgs = uintptr(unsafe.Pointer(&vargs[0]))
	if _, errno := dispatchIE.Invoke(dispIDNavigate, w32.IID_NULL, lcid, w32.DISPATCH_METHOD, dispParams, &exceptInfo, nil); errno != 0 {
		log.Printf("%s\n", errno)
	}
	oleAut.SysFreeString((*uint16)(unsafe.Pointer(uintptr(vargs[0].Val))))

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
	oleAut.VariantInit(document)
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
	oleAut.VariantInit(&vargs[0])
	vargs[0] = w32.NewVariant(w32.VT_BSTR, int64(uintptr(unsafe.Pointer(oleAut.SysAllocStringLen(params[0].(string))))))
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
	oleAut.VariantInit(&vargs[0])
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
	oleAut.VariantInit(&vargs[0])
	vargs[0] = w32.NewVariant(w32.VT_BSTR, int64(uintptr(unsafe.Pointer(oleAut.SysAllocStringLen(params[0].(string)))))) // 字串都要透過SpyFreeString來釋放記憶體
	dispParams.VArgs = uintptr(unsafe.Pointer(&vargs[0]))

	if _, errno := dispatchQuery.Invoke(dispIDValue, w32.IID_NULL, lcid, w32.DISPATCH_PROPERTYPUT, dispParams, &exceptInfo, nil); errno != 0 {
		log.Printf("%s\n", errno)
		return
	}
	oleAut.SysFreeString((*uint16)(unsafe.Pointer(uintptr(vargs[0].Val))))
}
