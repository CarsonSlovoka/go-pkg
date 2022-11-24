package w32

import "syscall"

// FONTENUMPROC https://learn.microsoft.com/en-us/previous-versions/dd162623(v=vs.85)
type FONTENUMPROC func(logFont *LOGFONT, textmetric *TEXTMETRIC, fontType uint32, lParam LPARAM) int32

// EnumFontFamProc https://learn.microsoft.com/en-us/previous-versions/dd162621(v=vs.85)
type EnumFontFamProc func(logFont *ENUMLOGFONT, textmetric *TEXTMETRIC, fontType uint32, lparam LPARAM) int32

// ENUMLOGFONT https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-enumlogfontw?redirectedfrom=MSDN
type ENUMLOGFONT struct {
	LogFont  LOGFONT
	FullName [LF_FULLFACESIZE]uint16 // For example, ABCD Font Company TrueType Bold Italic Sans Serif. // NameID=4
	Style    [LF_FACESIZE]uint16     // For example, Bold Italic.
}

func (e *ENUMLOGFONT) GetFullName() string {
	return syscall.UTF16ToString(e.FullName[:])
}

func (e *ENUMLOGFONT) GetStyle() string {
	return syscall.UTF16ToString(e.Style[:])
}

/*
// ENUMLOGFONTEX https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-enumlogfontexw
type ENUMLOGFONTEX struct {
	LogFont  LOGFONT
	FullName [LF_FULLFACESIZE]uint16
	Style    [LF_FACESIZE]uint16
	Script   [LF_FACESIZE]uint16 // The script, that is, the character set, of the font. For example, Cyrillic.
}

func (e *ENUMLOGFONTEX) GetFullName() string {
	return syscall.UTF16ToString(e.FullName[:])
}

func (e *ENUMLOGFONTEX) GetStyle() string {
	return syscall.UTF16ToString(e.Style[:])
}

func (e *ENUMLOGFONTEX) GetScript() string {
	return syscall.UTF16ToString(e.Script[:])
}

// ENUMLOGFONTEXDV https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-enumlogfontexdvw?redirectedfrom=MSDN
type ENUMLOGFONTEXDV struct {
	ElfEnumLogfontEx ENUMLOGFONTEX
	ElfDesignVector  DESIGNVECTOR
}
*/

// DESIGNVECTOR https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-designvector
type DESIGNVECTOR struct {
	DvReserved uint32
	DvNumAxes  uint32
	DvValues   [MM_MAX_NUMAXES]int32
}
