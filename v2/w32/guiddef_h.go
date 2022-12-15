package w32

type IID GUID

type REFCLSID /*const*/ *IID
type REFFMTID /*const*/ *IID
type REFGUID /*const*/ *GUID
type REFIID /*const*/ *IID // A reference to the interface identifier IID

type CLSID GUID
type LPCLSID *CLSID

var (
	// IID_NULL is null Interface ID, used when no other Interface ID is known.
	IID_NULL = NewGUID("00000000-0000-0000-0000-000000000000")

	// IID_IUnknown is for IUnknown interfaces.
	IID_IUnknown = NewGUID("00000000-0000-0000-C000-000000000046")

	// IID_IDispatch is for IDispatch interfaces.
	IID_IDispatch = NewGUID("00020400-0000-0000-C000-000000000046")

	// IID_IEnumVariant is for IEnumVariant interfaces
	IID_IEnumVariant = NewGUID("00020404-0000-0000-C000-000000000046")

	// IID_IConnectionPointContainer is for IConnectionPointContainer interfaces.
	IID_IConnectionPointContainer = NewGUID("{B196B284-BAB4-101A-B69C-00AA00341D07}")

	// IID_IConnectionPoint is for IConnectionPoint interfaces.
	IID_IConnectionPoint = NewGUID("{B196B286-BAB4-101A-B69C-00AA00341D07}")

	// IID_IInspectable is for IInspectable interfaces.
	IID_IInspectable = NewGUID("{AF86E2E0-B12D-4C6A-9C5A-D7AA65101E90}")

	// IID_IProvideClassInfo is for IProvideClassInfo interfaces.
	IID_IProvideClassInfo = NewGUID("{B196B283-BAB4-101A-B69C-00AA00341D07}")

	IID_IWebBrowser2 = NewGUID("D30C1661-CDAF-11D0-8A3E-00C04FC9E26E")
)
