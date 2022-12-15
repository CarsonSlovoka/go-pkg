//go:build windows

package w32

// https://learn.microsoft.com/en-us/dotnet/api/microsoft.visualstudio.ole.interop.clsctx?view=visualstudiosdk-2022#--
const (
	CLSCTX_INPROC_SERVER  = 1 // The code that creates and manages objects of this class is a DLL that runs in the same process as the caller of the function specifying the class context.
	CLSCTX_INPROC_HANDLER = 2 // Indicates a handler dll, which runs on the same process as the caller.
	CLSCTX_LOCAL_SERVER   = 4 // Indicates a server executable, which runs on the same machine but on a different process than the caller.

	// Deprecated: Use CLSCTX_REMOTE_SERVER instead.
	CLSCTX_INPROC_SERVER16      = 8
	CLSCTX_REMOTE_SERVER        = 16 // Indicates a server executable, which runs on a different machine than the caller.
	CLSCTX_INPROC_HANDLER16     = 32
	CLSCTX_NO_CODE_DOWNLOAD     = 1024 // Indicates that code should not be allowed to be downloaded from the Directory Service (if any) or the Internet.
	CLSCTX_ENABLE_CODE_DOWNLOAD = 8192 // Indicates that code should be allowed to be downloaded from the Directory Service (if any) or the Internet.
	CLSCTX_ALL                  = CLSCTX_INPROC_SERVER | CLSCTX_INPROC_HANDLER | CLSCTX_LOCAL_SERVER
	CLSCTX_INPROC               = CLSCTX_INPROC_SERVER | CLSCTX_INPROC_HANDLER
	CLSCTX_SERVER               = CLSCTX_INPROC_SERVER | CLSCTX_LOCAL_SERVER | CLSCTX_REMOTE_SERVER
)
