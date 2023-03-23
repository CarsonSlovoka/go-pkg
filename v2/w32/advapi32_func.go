//go:build windows

package w32

import (
	"syscall"
	"unsafe"
)

const (
	PNAddAccessAllowedAce ProcName = "AddAccessAllowedAce"

	PNAllocateAndInitializeSid ProcName = "AllocateAndInitializeSid"

	PNCheckTokenMembership ProcName = "CheckTokenMembership"

	PNConvertSidToStringSid ProcName = "ConvertSidToStringSidW"
	PNConvertStringSidToSid ProcName = "ConvertStringSidToSidW"

	PNCreateRestrictedToken ProcName = "CreateRestrictedToken"

	PNDuplicateToken ProcName = "DuplicateToken"

	PNFreeSid ProcName = "FreeSid"

	PNGetSidSubAuthorityCount ProcName = "GetSidSubAuthorityCount"

	PNIsValidSid ProcName = "IsValidSid"

	PNOpenProcessToken ProcName = "OpenProcessToken"
)

type AdApiDLL struct {
	*dLL
}

func NewAdApi32DLL(procList ...ProcName) *AdApiDLL {
	if len(procList) == 0 {
		procList = []ProcName{
			PNAddAccessAllowedAce,

			PNAllocateAndInitializeSid,

			PNCheckTokenMembership,

			PNConvertSidToStringSid,
			PNConvertStringSidToSid,

			PNCreateRestrictedToken,

			PNDuplicateToken,

			PNFreeSid,

			PNGetSidSubAuthorityCount,

			PNIsValidSid,

			PNOpenProcessToken,
		}
	}
	dll := newDll(DNAdApi32, procList)
	return &AdApiDLL{dll}
}

// AddAccessAllowedAce https://learn.microsoft.com/en-us/windows/win32/api/securitybaseapi/nf-securitybaseapi-addaccessallowedace
// errCode: ERROR_ALLOTTED_SPACE_EXCEEDED, ERROR_INVALID_ACL, ERROR_INVALID_SID, ERROR_REVISION_MISMATCH, ERROR_SUCCESS
func (dll *AdApiDLL) AddAccessAllowedAce(acl *AccessAllowedAce, aceRevision uint32, accessMask uint32, pSid *SID) syscall.Errno {
	proc := dll.mustProc(PNAddAccessAllowedAce)
	_, _, eno := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(&acl)),
		uintptr(aceRevision),
		uintptr(accessMask),
		uintptr(unsafe.Pointer(pSid)),
	)
	return eno
}

// AllocateAndInitializeSid https://learn.microsoft.com/en-us/windows/win32/api/securitybaseapi/nf-securitybaseapi-allocateandinitializesid
// 此函數可以幫你產生SID，如果你對字串比較熟悉可以用 ConvertStringSidToSid 也能產生SID
func (dll *AdApiDLL) AllocateAndInitializeSid(identifierAuthority *SidIdentifierAuthority,
	nSubAuthorityCount byte,
	nSubAuthority0 uint32,
	nSubAuthority1 uint32,
	nSubAuthority2 uint32,
	nSubAuthority3 uint32,
	nSubAuthority4 uint32,
	nSubAuthority5 uint32,
	nSubAuthority6 uint32,
	nSubAuthority7 uint32,
) (*SID, syscall.Errno) {
	var sid *SID
	proc := dll.mustProc(PNAllocateAndInitializeSid)
	_, _, eno := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(identifierAuthority)),
		uintptr(nSubAuthorityCount),
		uintptr(nSubAuthority0),
		uintptr(nSubAuthority1),
		uintptr(nSubAuthority2),
		uintptr(nSubAuthority3),
		uintptr(nSubAuthority4),
		uintptr(nSubAuthority5),
		uintptr(nSubAuthority6),
		uintptr(nSubAuthority7),
		uintptr(unsafe.Pointer(&sid)),
	)
	return sid, eno
}

// CheckTokenMembership https://learn.microsoft.com/en-us/windows/win32/api/securitybaseapi/nf-securitybaseapi-checktokenmembership
func (dll *AdApiDLL) CheckTokenMembership(tokenHandle HANDLE, sidToCheck *SID) (bool, syscall.Errno) {
	var isMember bool
	proc := dll.mustProc(PNCheckTokenMembership)
	_, _, eno := syscall.SyscallN(proc.Addr(),
		uintptr(tokenHandle),
		uintptr(unsafe.Pointer(sidToCheck)),
		uintptr(unsafe.Pointer(&isMember)),
	)
	return isMember, eno
}

// ConvertSidToStringSid https://learn.microsoft.com/en-us/windows/win32/api/sddl/nf-sddl-convertsidtostringsidw
func (dll *AdApiDLL) ConvertSidToStringSid(sid *SID) (string, syscall.Errno) {
	proc := dll.mustProc(PNConvertSidToStringSid)
	out := make([]uint16, 1024)
	_, _, eno := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(sid)),
		uintptr(unsafe.Pointer(&out)),
	)
	if eno != 0 {
		return "", eno // ERROR_NOT_ENOUGH_MEMORY, ERROR_INVALID_SID, ERROR_INVALID_PARAMETER
	}
	return syscall.UTF16ToString(out), 0
}

// ConvertStringSidToSid https://learn.microsoft.com/en-us/windows/win32/api/sddl/nf-sddl-convertstringsidtosidw
func (dll *AdApiDLL) ConvertStringSidToSid(sidStr string) (*SID, syscall.Errno) {
	var sid **SID
	sid = new(*SID)
	proc := dll.mustProc(PNConvertStringSidToSid)
	_, _, eno := syscall.SyscallN(proc.Addr(),
		UintptrFromStr(sidStr),
		uintptr(unsafe.Pointer(sid)),
	)
	return *sid, eno
}

// CreateRestrictedToken https://learn.microsoft.com/en-us/windows/win32/api/securitybaseapi/nf-securitybaseapi-createrestrictedtoken
func (dll *AdApiDLL) CreateRestrictedToken(existingTokenHandle HANDLE,
	flags uint32,
	disableSidCount uint32, sidsToDisable *SidAndAttributes,
	deletePrivilegeCount uint32, privilegesToDelete *LuidAndAttributes,
	restrictedSidCount uint32, sidsToRestrict *SidAndAttributes,
) (HANDLE, syscall.Errno) {
	var out *uintptr
	proc := dll.mustProc(PNCreateRestrictedToken)
	_, _, eno := syscall.SyscallN(proc.Addr(),
		uintptr(existingTokenHandle),
		uintptr(flags),
		uintptr(disableSidCount),
		uintptr(unsafe.Pointer(sidsToDisable)),
		uintptr(deletePrivilegeCount),
		uintptr(unsafe.Pointer(privilegesToDelete)),
		uintptr(restrictedSidCount),
		uintptr(unsafe.Pointer(sidsToRestrict)),
		uintptr(unsafe.Pointer(out)),
	)
	return HANDLE(*out), eno
}

// DuplicateToken https://learn.microsoft.com/en-us/windows/win32/api/securitybaseapi/nf-securitybaseapi-duplicatetoken
func (dll *AdApiDLL) DuplicateToken(existingTokenHandle HANDLE, impersonationLevel SECURITY_IMPERSONATION_LEVEL,
) (HANDLE, syscall.Errno) {
	var duplicateTokenHandle HANDLE
	proc := dll.mustProc(PNDuplicateToken)
	_, _, eno := syscall.SyscallN(proc.Addr(),
		uintptr(existingTokenHandle),
		uintptr(impersonationLevel),
		uintptr(unsafe.Pointer(&duplicateTokenHandle)),
	)
	return duplicateTokenHandle, eno
}

// FreeSid https://learn.microsoft.com/en-us/windows/win32/api/securitybaseapi/nf-securitybaseapi-freesid
// If the function succeeds, the function returns NULL.
// If the function fails, it returns a pointer to the SID structure represented by the pSid parameter.
func (dll *AdApiDLL) FreeSid(pSid *SID) uintptr {
	proc := dll.mustProc(PNFreeSid)
	r, _, _ := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(pSid)),
	)
	return r
}

// GetSidSubAuthorityCount https://learn.microsoft.com/en-us/windows/win32/api/securitybaseapi/nf-securitybaseapi-getsidsubauthoritycount
// This function does not handle SID structures that are not valid. Call the IsValidSid function to verify that the SID structure is valid before you call this function.
func (dll *AdApiDLL) GetSidSubAuthorityCount(pSid *SID) (byte, syscall.Errno) {
	proc := dll.mustProc(PNGetSidSubAuthorityCount)
	pCount, _, eno := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(pSid)),
	)
	// pCount其實就是&pSid.SubAuthorityCount，也就是pSid.SubAuthorityCount的位址
	return *(*byte)(unsafe.Pointer(pCount)), eno
}

// IsValidSid https://learn.microsoft.com/en-us/windows/win32/api/securitybaseapi/nf-securitybaseapi-isvalidsid
func (dll *AdApiDLL) IsValidSid(pSid *SID) (bool, syscall.Errno) {
	proc := dll.mustProc(PNIsValidSid)
	r, _, eno := syscall.SyscallN(proc.Addr(),
		uintptr(unsafe.Pointer(pSid)),
	)
	return r != 0, eno
}

// OpenProcessToken https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-openprocesstoken
// 🧙 Call CloseHandle when you are not used.
func (dll *AdApiDLL) OpenProcessToken(processHandle HANDLE, desiredAccess uint32) (HANDLE, syscall.Errno) {
	var h HANDLE
	proc := dll.mustProc(PNOpenProcessToken)
	_, _, eno := syscall.SyscallN(proc.Addr(),
		uintptr(processHandle),
		uintptr(desiredAccess),
		uintptr(unsafe.Pointer(&h)),
	)
	return h, eno
}
