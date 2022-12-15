package w32

var (
	CLSID_InternetExplorer *GUID
)

func init() {
	CLSID_InternetExplorer = NewGUID("0002DF01-0000-0000-C000-000000000046")
}
