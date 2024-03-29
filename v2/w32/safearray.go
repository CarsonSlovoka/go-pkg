// Package is meant to retrieve and process safe array data returned from COM.

package w32

// SafeArrayBound defines the SafeArray boundaries.
type SafeArrayBound struct {
	Elements   uint32
	LowerBound int32
}

// SafeArray is how COM handles arrays.
// https://learn.microsoft.com/en-us/archive/msdn-magazine/2017/march/introducing-the-safearray-data-structure#introducing-the-safearray-data-structure-1
type SafeArray struct {
	Dimensions   uint16
	FeaturesFlag uint16
	ElementsSize uint32
	LocksAmount  uint32
	Data         uint32
	Bounds       [16]byte
}

// SAFEARRAY is obsolete, exists for backwards compatibility.
// Use SafeArray
type SAFEARRAY SafeArray

// SAFEARRAYBOUND is obsolete, exists for backwards compatibility.
// Use SafeArrayBound
type SAFEARRAYBOUND SafeArrayBound
