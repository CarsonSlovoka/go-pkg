package w32

/* Flags for IDispatch::Invoke */

type DispatchFlag uint16

const (
	DISPATCH_METHOD         DispatchFlag = 0x1
	DISPATCH_PROPERTYGET    DispatchFlag = 0x2
	DISPATCH_PROPERTYPUT    DispatchFlag = 0x4
	DISPATCH_PROPERTYPUTREF DispatchFlag = 0x8
)
