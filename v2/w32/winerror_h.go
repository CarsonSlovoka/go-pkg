package w32

/************************************************************************
*                                                                       *
*   winerror.h --  error code definitions for the Win32 API functions   *
*                                                                       *
*   Copyright (c) Microsoft Corp. All rights reserved.                  *
*                                                                       *
************************************************************************/

const (
	FACILITY_NULL                                     = 0
	FACILITY_RPC                                      = 1
	FACILITY_DISPATCH                                 = 2
	FACILITY_STORAGE                                  = 3
	FACILITY_ITF                                      = 4
	FACILITY_WIN32                                    = 7
	FACILITY_WINDOWS                                  = 8
	FACILITY_SSPI                                     = 9
	FACILITY_SECURITY                                 = 9
	FACILITY_CONTROL                                  = 10
	FACILITY_CERT                                     = 11
	FACILITY_INTERNET                                 = 12
	FACILITY_MEDIASERVER                              = 13
	FACILITY_MSMQ                                     = 14
	FACILITY_SETUPAPI                                 = 15
	FACILITY_SCARD                                    = 16
	FACILITY_COMPLUS                                  = 17
	FACILITY_AAF                                      = 18
	FACILITY_URT                                      = 19
	FACILITY_ACS                                      = 20
	FACILITY_DPLAY                                    = 21
	FACILITY_UMI                                      = 22
	FACILITY_SXS                                      = 23
	FACILITY_WINDOWS_CE                               = 24
	FACILITY_HTTP                                     = 25
	FACILITY_USERMODE_COMMONLOG                       = 26
	FACILITY_WER                                      = 27
	FACILITY_USERMODE_FILTER_MANAGER                  = 31
	FACILITY_BACKGROUNDCOPY                           = 32
	FACILITY_CONFIGURATION                            = 33
	FACILITY_WIA                                      = 33
	FACILITY_STATE_MANAGEMENT                         = 34
	FACILITY_METADIRECTORY                            = 35
	FACILITY_WINDOWSUPDATE                            = 36
	FACILITY_DIRECTORYSERVICE                         = 37
	FACILITY_GRAPHICS                                 = 38
	FACILITY_SHELL                                    = 39
	FACILITY_NAP                                      = 39
	FACILITY_TPM_SERVICES                             = 40
	FACILITY_TPM_SOFTWARE                             = 41
	FACILITY_UI                                       = 42
	FACILITY_XAML                                     = 43
	FACILITY_ACTION_QUEUE                             = 44
	FACILITY_PLA                                      = 48
	FACILITY_WINDOWS_SETUP                            = 48
	FACILITY_FVE                                      = 49
	FACILITY_FWP                                      = 50
	FACILITY_WINRM                                    = 51
	FACILITY_NDIS                                     = 52
	FACILITY_USERMODE_HYPERVISOR                      = 53
	FACILITY_CMI                                      = 54
	FACILITY_USERMODE_VIRTUALIZATION                  = 55
	FACILITY_USERMODE_VOLMGR                          = 56
	FACILITY_BCD                                      = 57
	FACILITY_USERMODE_VHD                             = 58
	FACILITY_USERMODE_HNS                             = 59
	FACILITY_SDIAG                                    = 60
	FACILITY_WEBSERVICES                              = 61
	FACILITY_WINPE                                    = 61
	FACILITY_WPN                                      = 62
	FACILITY_WINDOWS_STORE                            = 63
	FACILITY_INPUT                                    = 64
	FACILITY_EAP                                      = 66
	FACILITY_WINDOWS_DEFENDER                         = 80
	FACILITY_OPC                                      = 81
	FACILITY_XPS                                      = 82
	FACILITY_MBN                                      = 84
	FACILITY_POWERSHELL                               = 84
	FACILITY_RAS                                      = 83
	FACILITY_P2P_INT                                  = 98
	FACILITY_P2P                                      = 99
	FACILITY_DAF                                      = 100
	FACILITY_BLUETOOTH_ATT                            = 101
	FACILITY_AUDIO                                    = 102
	FACILITY_STATEREPOSITORY                          = 103
	FACILITY_VISUALCPP                                = 109
	FACILITY_SCRIPT                                   = 112
	FACILITY_PARSE                                    = 113
	FACILITY_BLB                                      = 120
	FACILITY_BLB_CLI                                  = 121
	FACILITY_WSBAPP                                   = 122
	FACILITY_BLBUI                                    = 128
	FACILITY_USN                                      = 129
	FACILITY_USERMODE_VOLSNAP                         = 130
	FACILITY_TIERING                                  = 131
	FACILITY_WSB_ONLINE                               = 133
	FACILITY_ONLINE_ID                                = 134
	FACILITY_DEVICE_UPDATE_AGENT                      = 135
	FACILITY_DRVSERVICING                             = 136
	FACILITY_DLS                                      = 153
	FACILITY_DELIVERY_OPTIMIZATION                    = 208
	FACILITY_USERMODE_SPACES                          = 231
	FACILITY_USER_MODE_SECURITY_CORE                  = 232
	FACILITY_USERMODE_LICENSING                       = 234
	FACILITY_SOS                                      = 160
	FACILITY_DEBUGGERS                                = 176
	FACILITY_SPP                                      = 256
	FACILITY_RESTORE                                  = 256
	FACILITY_DMSERVER                                 = 256
	FACILITY_DEPLOYMENT_SERVICES_SERVER               = 257
	FACILITY_DEPLOYMENT_SERVICES_IMAGING              = 258
	FACILITY_DEPLOYMENT_SERVICES_MANAGEMENT           = 259
	FACILITY_DEPLOYMENT_SERVICES_UTIL                 = 260
	FACILITY_DEPLOYMENT_SERVICES_BINLSVC              = 261
	FACILITY_DEPLOYMENT_SERVICES_PXE                  = 263
	FACILITY_DEPLOYMENT_SERVICES_TFTP                 = 264
	FACILITY_DEPLOYMENT_SERVICES_TRANSPORT_MANAGEMENT = 272
	FACILITY_DEPLOYMENT_SERVICES_DRIVER_PROVISIONING  = 278
	FACILITY_DEPLOYMENT_SERVICES_MULTICAST_SERVER     = 289
	FACILITY_DEPLOYMENT_SERVICES_MULTICAST_CLIENT     = 290
	FACILITY_DEPLOYMENT_SERVICES_CONTENT_PROVIDER     = 293
	FACILITY_LINGUISTIC_SERVICES                      = 305
	FACILITY_AUDIOSTREAMING                           = 1094
	FACILITY_ACCELERATOR                              = 1536
	FACILITY_WMAAECMA                                 = 1996
	FACILITY_DIRECTMUSIC                              = 2168
	FACILITY_DIRECT3D10                               = 2169
	FACILITY_DXGI                                     = 2170
	FACILITY_DXGI_DDI                                 = 2171
	FACILITY_DIRECT3D11                               = 2172
	FACILITY_DIRECT3D11_DEBUG                         = 2173
	FACILITY_DIRECT3D12                               = 2174
	FACILITY_DIRECT3D12_DEBUG                         = 2175
	FACILITY_LEAP                                     = 2184
	FACILITY_AUDCLNT                                  = 2185
	FACILITY_WINCODEC_DWRITE_DWM                      = 2200
	FACILITY_WINML                                    = 2192
	FACILITY_DIRECT2D                                 = 2201
	FACILITY_DEFRAG                                   = 2304
	FACILITY_USERMODE_SDBUS                           = 2305
	FACILITY_JSCRIPT                                  = 2306
	FACILITY_PIDGENX                                  = 2561
	FACILITY_EAS                                      = 85
	FACILITY_WEB                                      = 885
	FACILITY_WEB_SOCKET                               = 886
	FACILITY_MOBILE                                   = 1793
	FACILITY_SQLITE                                   = 1967
	FACILITY_UTC                                      = 1989
	FACILITY_WEP                                      = 2049
	FACILITY_SYNCENGINE                               = 2050
	FACILITY_XBOX                                     = 2339
	FACILITY_PIX                                      = 2748

	// Define the severity codes

	// The operation completed successfully.
	ERROR_SUCCESS = 0

	NO_ERROR = 0

	// Incorrect function.
	ERROR_INVALID_FUNCTION = 1 // dderror

	// The system cannot find the file specified.
	ERROR_FILE_NOT_FOUND = 2

	// The system cannot find the path specified.
	ERROR_PATH_NOT_FOUND = 3

	// The system cannot open the file.
	ERROR_TOO_MANY_OPEN_FILES = 4

	// Access is denied.
	ERROR_ACCESS_DENIED = 5

	// The handle is invalid.
	ERROR_INVALID_HANDLE = 6

	// The storage control blocks were destroyed.
	ERROR_ARENA_TRASHED = 7

	// ERROR_NOT_ENOUGH_MEMORY
	// Not enough memory resources are available to process this command.
	ERROR_NOT_ENOUGH_MEMORY = 8 // dderror

	// ERROR_INVALID_BLOCK
	// The storage control block address is invalid.
	ERROR_INVALID_BLOCK = 9

	// ERROR_BAD_ENVIRONMENT
	// The environment is incorrect.
	ERROR_BAD_ENVIRONMENT = 10

	// ERROR_BAD_FORMAT
	// An attempt was made to load a program with an incorrect format.
	ERROR_BAD_FORMAT = 11

	// ERROR_INVALID_ACCESS
	// The access code is invalid.
	ERROR_INVALID_ACCESS = 12

	// ERROR_INVALID_DATA
	// The data is invalid.
	ERROR_INVALID_DATA = 13

	// ERROR_OUTOFMEMORY
	// Not enough memory resources are available to complete this operation.
	ERROR_OUTOFMEMORY = 14

	// ERROR_INVALID_DRIVE
	// The system cannot find the drive specified.
	ERROR_INVALID_DRIVE = 15

	// ERROR_CURRENT_DIRECTORY
	// The directory cannot be removed.
	ERROR_CURRENT_DIRECTORY = 16

	// ERROR_NOT_SAME_DEVICE
	// The system cannot move the file to a different disk drive.
	ERROR_NOT_SAME_DEVICE = 17

	// ERROR_NO_MORE_FILES
	// There are no more files.
	ERROR_NO_MORE_FILES = 18

	// ERROR_WRITE_PROTECT
	// The media is write protected.
	ERROR_WRITE_PROTECT = 19

	// ERROR_BAD_UNIT
	// The system cannot find the device specified.
	ERROR_BAD_UNIT = 20

	// ERROR_NOT_READY
	// The device is not ready.
	ERROR_NOT_READY = 21

	// ERROR_BAD_COMMAND
	// The device does not recognize the command.
	ERROR_BAD_COMMAND = 22

	// ERROR_CRC
	// Data error (cyclic redundancy check).
	ERROR_CRC = 23

	// ERROR_BAD_LENGTH
	// The program issued a command but the command length is incorrect.
	ERROR_BAD_LENGTH = 24

	// ERROR_SEEK
	// The drive cannot locate a specific area or track on the disk.
	ERROR_SEEK = 25

	// ERROR_NOT_DOS_DISK
	// The specified disk or diskette cannot be accessed.
	ERROR_NOT_DOS_DISK = 26

	// ERROR_SECTOR_NOT_FOUND
	// The drive cannot find the sector requested.
	ERROR_SECTOR_NOT_FOUND = 27

	// ERROR_OUT_OF_PAPER
	// The printer is out of paper.
	ERROR_OUT_OF_PAPER = 28

	// ERROR_WRITE_FAULT
	// The system cannot write to the specified device.
	ERROR_WRITE_FAULT = 29

	// ERROR_READ_FAULT
	// The system cannot read from the specified device.
	ERROR_READ_FAULT = 30

	// ERROR_GEN_FAILURE
	// A device attached to the system is not functioning.
	ERROR_GEN_FAILURE = 31

	// ERROR_SHARING_VIOLATION
	// The process cannot access the file because it is being used by another process.
	ERROR_SHARING_VIOLATION = 32

	// ERROR_LOCK_VIOLATION
	// The process cannot access the file because another process has locked a portion of the file.
	ERROR_LOCK_VIOLATION = 33

	// ERROR_WRONG_DISK
	// The wrong diskette is in the drive.
	// Insert %2 (Volume Serial Number: %3) into drive %1.
	ERROR_WRONG_DISK = 34

	// ERROR_SHARING_BUFFER_EXCEEDED
	// Too many files opened for sharing.
	ERROR_SHARING_BUFFER_EXCEEDED = 36

	// ERROR_HANDLE_EOF
	// Reached the end of the file.
	ERROR_HANDLE_EOF = 38

	// ERROR_HANDLE_DISK_FUL
	// The disk is full.
	ERROR_HANDLE_DISK_FULL = 39

	// ERROR_NOT_SUPPORTED
	// The request is not supported.
	ERROR_NOT_SUPPORTED = 50

	// ERROR_REM_NOT_LIST
	// Windows cannot find the network path. Verify that the network path is correct and the destination computer is not busy or turned off. If Windows still cannot find the network path, contact your network administrator.
	ERROR_REM_NOT_LIST = 51

	// ERROR_DUP_NAME
	// You were not connected because a duplicate name exists on the network. If joining a domain, go to System in Control Panel to change the computer name and try again. If joining a workgroup, choose another workgroup name.
	ERROR_DUP_NAME = 52

	// ERROR_BAD_NETPATH
	// The network path was not found.
	ERROR_BAD_NETPATH = 53

	// ERROR_NETWORK_BUSY
	// The network is busy.
	ERROR_NETWORK_BUSY = 54

	// ERROR_DEV_NOT_EXIST
	// The specified network resource or device is no longer available.
	ERROR_DEV_NOT_EXIST = 55

	// ERROR_TOO_MANY_CMDS
	// The network BIOS command limit has been reached.
	ERROR_TOO_MANY_CMDS = 56

	// ERROR_ADAP_HDW_ERR
	// A network adapter hardware error occurred.
	ERROR_ADAP_HDW_ERR = 57

	// ERROR_BAD_NET_RESP
	// The specified server cannot perform the requested operation.
	ERROR_BAD_NET_RESP = 58

	// ERROR_UNEXP_NET_ERR
	// An unexpected network error occurred.
	ERROR_UNEXP_NET_ERR = 59

	// ERROR_BAD_REM_ADAP
	// The remote adapter is not compatible.
	ERROR_BAD_REM_ADAP = 60

	// ERROR_PRINTQ_FUL
	// The printer queue is full.
	ERROR_PRINTQ_FULL = 61

	// ERROR_NO_SPOOL_SPACE
	// Space to store the file waiting to be printed is not available on the server.
	ERROR_NO_SPOOL_SPACE = 62

	// ERROR_PRINT_CANCELLED
	// Your file waiting to be printed was deleted.
	ERROR_PRINT_CANCELLED = 63

	// ERROR_NETNAME_DELETED
	// The specified network name is no longer available.
	ERROR_NETNAME_DELETED = 64

	// ERROR_NETWORK_ACCESS_DENIED
	// Network access is denied.
	ERROR_NETWORK_ACCESS_DENIED = 65

	// ERROR_BAD_DEV_TYPE
	// The network resource type is not correct.
	ERROR_BAD_DEV_TYPE = 66

	// ERROR_BAD_NET_NAME
	// The network name cannot be found.
	ERROR_BAD_NET_NAME = 67

	// ERROR_TOO_MANY_NAMES
	// The name limit for the local computer network adapter card was exceeded.
	ERROR_TOO_MANY_NAMES = 68

	// ERROR_TOO_MANY_SESS
	// The network BIOS session limit was exceeded.
	ERROR_TOO_MANY_SESS = 69

	// ERROR_SHARING_PAUSED
	// The remote server has been paused or is in the process of being started.
	ERROR_SHARING_PAUSED = 70

	// ERROR_REQ_NOT_ACCEP
	// No more connections can be made to this remote computer at this time because there are already as many connections as the computer can accept.
	ERROR_REQ_NOT_ACCEP = 71

	// ERROR_REDIR_PAUSED
	// The specified printer or disk device has been paused.
	ERROR_REDIR_PAUSED = 72

	// ERROR_FILE_EXISTS
	// The file exists.
	ERROR_FILE_EXISTS = 80

	// ERROR_CANNOT_MAKE
	// The directory or file cannot be created.
	ERROR_CANNOT_MAKE = 82

	// ERROR_FAIL_I24
	// Fail on INT=24.
	ERROR_FAIL_I24 = 83

	// ERROR_OUT_OF_STRUCTURES
	// Storage to process this request is not available.
	ERROR_OUT_OF_STRUCTURES = 84

	// ERROR_ALREADY_ASSIGNED
	// The local device name is already in use.
	ERROR_ALREADY_ASSIGNED = 85

	// ERROR_INVALID_PASSWORD
	// The specified network password is not correct.
	ERROR_INVALID_PASSWORD = 86

	// ERROR_INVALID_PARAMETER
	// The parameter is incorrect.
	ERROR_INVALID_PARAMETER = 87

	// ERROR_NET_WRITE_FAULT
	// A write fault occurred on the network.
	ERROR_NET_WRITE_FAULT = 88

	// ERROR_NO_PROC_SLOTS
	// The system cannot start another process at this time.
	ERROR_NO_PROC_SLOTS = 89

	// ERROR_TOO_MANY_SEMAPHORES
	// Cannot create another system semaphore.
	ERROR_TOO_MANY_SEMAPHORES = 100

	// ERROR_EXCL_SEM_ALREADY_OWNED
	// The exclusive semaphore is owned by another process.
	ERROR_EXCL_SEM_ALREADY_OWNED = 101

	// ERROR_SEM_IS_SET
	// The semaphore is set and cannot be closed.
	ERROR_SEM_IS_SET = 102

	// ERROR_TOO_MANY_SEM_REQUESTS
	// The semaphore cannot be set again.
	ERROR_TOO_MANY_SEM_REQUESTS = 103

	// ERROR_INVALID_AT_INTERRUPT_TIME
	// Cannot request exclusive semaphores at interrupt time.
	ERROR_INVALID_AT_INTERRUPT_TIME = 104

	// ERROR_SEM_OWNER_DIED
	// The previous ownership of this semaphore has ended.
	ERROR_SEM_OWNER_DIED = 105

	// ERROR_SEM_USER_LIMIT
	// Insert the diskette for drive %1.
	ERROR_SEM_USER_LIMIT = 106

	// ERROR_DISK_CHANGE
	// The program stopped because an alternate diskette was not inserted.
	ERROR_DISK_CHANGE = 107

	// ERROR_DRIVE_LOCKED
	// The disk is in use or locked by another process.
	ERROR_DRIVE_LOCKED = 108

	// ERROR_BROKEN_PIPE
	// The pipe has been ended.
	ERROR_BROKEN_PIPE = 109

	// ERROR_OPEN_FAILED
	// The system cannot open the device or file specified.
	ERROR_OPEN_FAILED = 110

	// ERROR_BUFFER_OVERFLOW
	// The file name is too long.
	ERROR_BUFFER_OVERFLOW = 111

	// ERROR_DISK_FUL
	// There is not enough space on the disk.
	ERROR_DISK_FULL = 112

	// ERROR_NO_MORE_SEARCH_HANDLES
	// No more internal file identifiers available.
	ERROR_NO_MORE_SEARCH_HANDLES = 113

	// ERROR_INVALID_TARGET_HANDLE
	// The target internal file identifier is incorrect.
	ERROR_INVALID_TARGET_HANDLE = 114

	// ERROR_INVALID_CATEGORY
	// The IOCTL call made by the application program is not correct.
	ERROR_INVALID_CATEGORY = 117

	// ERROR_INVALID_VERIFY_SWITCH
	// The verify-on-write switch parameter value is not correct.
	ERROR_INVALID_VERIFY_SWITCH = 118

	// ERROR_BAD_DRIVER_LEVE
	// The system does not support the command requested.
	ERROR_BAD_DRIVER_LEVEL = 119

	// ERROR_CALL_NOT_IMPLEMENTED
	// This function is not supported on this system.
	ERROR_CALL_NOT_IMPLEMENTED = 120

	// ERROR_SEM_TIMEOUT
	// The semaphore timeout period has expired.
	ERROR_SEM_TIMEOUT = 121

	// ERROR_INSUFFICIENT_BUFFER
	// The data area passed to a system call is too small.
	ERROR_INSUFFICIENT_BUFFER = 122

	// ERROR_INVALID_NAME
	// The filename, directory name, or volume label syntax is incorrect.
	ERROR_INVALID_NAME = 123

	// ERROR_INVALID_LEVE
	// The system call level is not correct.
	ERROR_INVALID_LEVEL = 124

	// ERROR_NO_VOLUME_LABE
	// The disk has no volume label.
	ERROR_NO_VOLUME_LABEL = 125

	// ERROR_MOD_NOT_FOUND
	// The specified module could not be found.
	ERROR_MOD_NOT_FOUND = 126

	// ERROR_PROC_NOT_FOUND
	// The specified procedure could not be found.
	ERROR_PROC_NOT_FOUND = 127

	// ERROR_WAIT_NO_CHILDREN
	// There are no child processes to wait for.
	ERROR_WAIT_NO_CHILDREN = 128

	// ERROR_CHILD_NOT_COMPLETE
	// The %1 application cannot be run in Win32 mode.
	ERROR_CHILD_NOT_COMPLETE = 129

	// ERROR_DIRECT_ACCESS_HANDLE
	// Attempt to use a file handle to an open disk partition for an operation other than raw disk I/O.
	ERROR_DIRECT_ACCESS_HANDLE = 130

	// ERROR_NEGATIVE_SEEK
	// An attempt was made to move the file pointer before the beginning of the file.
	ERROR_NEGATIVE_SEEK = 131

	// ERROR_SEEK_ON_DEVICE
	// The file pointer cannot be set on the specified device or file.
	ERROR_SEEK_ON_DEVICE = 132

	// ERROR_IS_JOIN_TARGET
	// A JOIN or SUBST command cannot be used for a drive that contains previously joined drives.
	ERROR_IS_JOIN_TARGET = 133

	// ERROR_IS_JOINED
	// An attempt was made to use a JOIN or SUBST command on a drive that has already been joined.
	ERROR_IS_JOINED = 134

	// ERROR_IS_SUBSTED
	// An attempt was made to use a JOIN or SUBST command on a drive that has already been substituted.
	ERROR_IS_SUBSTED = 135

	// ERROR_NOT_JOINED
	// The system tried to delete the JOIN of a drive that is not joined.
	ERROR_NOT_JOINED = 136

	// ERROR_NOT_SUBSTED
	// The system tried to delete the substitution of a drive that is not substituted.
	ERROR_NOT_SUBSTED = 137

	// ERROR_JOIN_TO_JOIN
	// The system tried to join a drive to a directory on a joined drive.
	ERROR_JOIN_TO_JOIN = 138

	// ERROR_SUBST_TO_SUBST
	// The system tried to substitute a drive to a directory on a substituted drive.
	ERROR_SUBST_TO_SUBST = 139

	// ERROR_JOIN_TO_SUBST
	// The system tried to join a drive to a directory on a substituted drive.
	ERROR_JOIN_TO_SUBST = 140

	// ERROR_SUBST_TO_JOIN
	// The system tried to SUBST a drive to a directory on a joined drive.
	ERROR_SUBST_TO_JOIN = 141

	// ERROR_BUSY_DRIVE
	// The system cannot perform a JOIN or SUBST at this time.
	ERROR_BUSY_DRIVE = 142

	// ERROR_SAME_DRIVE
	// The system cannot join or substitute a drive to or for a directory on the same drive.
	ERROR_SAME_DRIVE = 143

	// ERROR_DIR_NOT_ROOT
	// The directory is not a subdirectory of the root directory.
	ERROR_DIR_NOT_ROOT = 144

	// ERROR_DIR_NOT_EMPTY
	// The directory is not empty.
	ERROR_DIR_NOT_EMPTY = 145

	// ERROR_IS_SUBST_PATH
	// The path specified is being used in a substitute.
	ERROR_IS_SUBST_PATH = 146

	// ERROR_IS_JOIN_PATH
	// Not enough resources are available to process this command.
	ERROR_IS_JOIN_PATH = 147

	// ERROR_PATH_BUSY
	// The path specified cannot be used at this time.
	ERROR_PATH_BUSY = 148

	// ERROR_IS_SUBST_TARGET
	// An attempt was made to join or substitute a drive for which a directory on the drive is the target of a previous substitute.
	ERROR_IS_SUBST_TARGET = 149

	// ERROR_SYSTEM_TRACE
	// System trace information was not specified in your CONFIG.SYS file, or tracing is disallowed.
	ERROR_SYSTEM_TRACE = 150

	// ERROR_INVALID_EVENT_COUNT
	// The number of specified semaphore events for DosMuxSemWait is not correct.
	ERROR_INVALID_EVENT_COUNT = 151

	// ERROR_TOO_MANY_MUXWAITERS
	// DosMuxSemWait did not execute; too many semaphores are already set.
	ERROR_TOO_MANY_MUXWAITERS = 152

	// ERROR_INVALID_LIST_FORMAT
	// The DosMuxSemWait list is not correct.
	ERROR_INVALID_LIST_FORMAT = 153

	// ERROR_LABEL_TOO_LONG
	// The volume label you entered exceeds the label character limit of the target file system.
	ERROR_LABEL_TOO_LONG = 154

	// ERROR_TOO_MANY_TCBS
	// Cannot create another thread.
	ERROR_TOO_MANY_TCBS = 155

	// ERROR_SIGNAL_REFUSED
	// The recipient process has refused the signal.
	ERROR_SIGNAL_REFUSED = 156

	// ERROR_DISCARDED
	// The segment is already discarded and cannot be locked.
	ERROR_DISCARDED = 157

	// ERROR_NOT_LOCKED
	// The segment is already unlocked.
	ERROR_NOT_LOCKED = 158

	// ERROR_BAD_THREADID_ADDR
	// The address for the thread ID is not correct.
	ERROR_BAD_THREADID_ADDR = 159

	// ERROR_BAD_ARGUMENTS
	// One or more arguments are not correct.
	ERROR_BAD_ARGUMENTS = 160

	// ERROR_BAD_PATHNAME
	// The specified path is invalid.
	ERROR_BAD_PATHNAME = 161

	// ERROR_SIGNAL_PENDING
	// A signal is already pending.
	ERROR_SIGNAL_PENDING = 162

	// ERROR_MAX_THRDS_REACHED
	// No more threads can be created in the system.
	ERROR_MAX_THRDS_REACHED = 164

	// ERROR_LOCK_FAILED
	// Unable to lock a region of a file.
	ERROR_LOCK_FAILED = 167

	// ERROR_BUSY
	// The requested resource is in use.
	ERROR_BUSY = 170

	// ERROR_DEVICE_SUPPORT_IN_PROGRESS
	// Device's command support detection is in progress.
	ERROR_DEVICE_SUPPORT_IN_PROGRESS = 171

	// ERROR_CANCEL_VIOLATION
	// A lock request was not outstanding for the supplied cancel region.
	ERROR_CANCEL_VIOLATION = 173

	// ERROR_ATOMIC_LOCKS_NOT_SUPPORTED
	// The file system does not support atomic changes to the lock type.
	ERROR_ATOMIC_LOCKS_NOT_SUPPORTED = 174

	// ERROR_INVALID_SEGMENT_NUMBER
	// The system detected a segment number that was not correct.
	ERROR_INVALID_SEGMENT_NUMBER = 180

	// ERROR_INVALID_ORDINA
	// The operating system cannot run %1.
	ERROR_INVALID_ORDINAL = 182

	// ERROR_ALREADY_EXISTS
	// Cannot create a file when that file already exists.
	ERROR_ALREADY_EXISTS = 183

	// ERROR_INVALID_FLAG_NUMBER
	// The flag passed is not correct.
	ERROR_INVALID_FLAG_NUMBER = 186

	// ERROR_SEM_NOT_FOUND
	// The specified system semaphore name was not found.
	ERROR_SEM_NOT_FOUND = 187

	// ERROR_INVALID_STARTING_CODESEG
	// The operating system cannot run %1.
	ERROR_INVALID_STARTING_CODESEG = 188

	// ERROR_INVALID_STACKSEG
	// The operating system cannot run %1.
	ERROR_INVALID_STACKSEG = 189

	// ERROR_INVALID_MODULETYPE
	// The operating system cannot run %1.
	ERROR_INVALID_MODULETYPE = 190

	// ERROR_INVALID_EXE_SIGNATURE
	// Cannot run %1 in Win32 mode.
	ERROR_INVALID_EXE_SIGNATURE = 191

	// ERROR_EXE_MARKED_INVALID
	// The operating system cannot run %1.
	ERROR_EXE_MARKED_INVALID = 192

	// ERROR_BAD_EXE_FORMAT
	// %1 is not a valid Win32 application.
	ERROR_BAD_EXE_FORMAT = 193

	// ERROR_ITERATED_DATA_EXCEEDS_64k
	// The operating system cannot run %1.
	ERROR_ITERATED_DATA_EXCEEDS_64k = 194

	// ERROR_INVALID_MINALLOCSIZE
	// The operating system cannot run %1.
	ERROR_INVALID_MINALLOCSIZE = 195

	// ERROR_DYNLINK_FROM_INVALID_RING
	// The operating system cannot run this application program.
	ERROR_DYNLINK_FROM_INVALID_RING = 196

	// ERROR_IOPL_NOT_ENABLED
	// The operating system is not presently configured to run this application.
	ERROR_IOPL_NOT_ENABLED = 197

	// ERROR_INVALID_SEGDP
	// The operating system cannot run %1.
	ERROR_INVALID_SEGDPL = 198

	// ERROR_AUTODATASEG_EXCEEDS_64k
	// The operating system cannot run this application program.
	ERROR_AUTODATASEG_EXCEEDS_64k = 199

	// ERROR_RING2SEG_MUST_BE_MOVABLE
	// The code segment cannot be greater than or equal to=64K.
	ERROR_RING2SEG_MUST_BE_MOVABLE = 200

	// ERROR_RELOC_CHAIN_XEEDS_SEGLIM
	// The operating system cannot run %1.
	ERROR_RELOC_CHAIN_XEEDS_SEGLIM = 201

	// ERROR_INFLOOP_IN_RELOC_CHAIN
	// The operating system cannot run %1.
	ERROR_INFLOOP_IN_RELOC_CHAIN = 202

	// ERROR_ENVVAR_NOT_FOUND
	// The system could not find the environment option that was entered.
	ERROR_ENVVAR_NOT_FOUND = 203

	// ERROR_NO_SIGNAL_SENT
	// No process in the command subtree has a signal handler.
	ERROR_NO_SIGNAL_SENT = 205

	// ERROR_FILENAME_EXCED_RANGE
	// The filename or extension is too long.
	ERROR_FILENAME_EXCED_RANGE = 206

	// ERROR_RING2_STACK_IN_USE
	// The ring=2 stack is in use.
	ERROR_RING2_STACK_IN_USE = 207

	// ERROR_META_EXPANSION_TOO_LONG
	// The global filename characters, * or ?, are entered incorrectly or too many global filename characters are specified.
	ERROR_META_EXPANSION_TOO_LONG = 208

	// ERROR_INVALID_SIGNAL_NUMBER
	// The signal being posted is not correct.
	ERROR_INVALID_SIGNAL_NUMBER = 209

	// ERROR_THREAD_1_INACTIVE
	// The signal handler cannot be set.
	ERROR_THREAD_1_INACTIVE = 210

	// ERROR_LOCKED
	// The segment is locked and cannot be reallocated.
	ERROR_LOCKED = 212

	// ERROR_TOO_MANY_MODULES
	// Too many dynamic-link modules are attached to this program or dynamic-link module.
	ERROR_TOO_MANY_MODULES = 214

	// ERROR_NESTING_NOT_ALLOWED
	// Cannot nest calls to LoadModule.
	ERROR_NESTING_NOT_ALLOWED = 215

	// ERROR_EXE_MACHINE_TYPE_MISMATCH
	// This version of %1 is not compatible with the version of Windows you're running. Check your computer's system information and then contact the software publisher.
	ERROR_EXE_MACHINE_TYPE_MISMATCH = 216

	// ERROR_EXE_CANNOT_MODIFY_SIGNED_BINARY
	// The image file %1 is signed, unable to modify.
	ERROR_EXE_CANNOT_MODIFY_SIGNED_BINARY = 217

	// ERROR_EXE_CANNOT_MODIFY_STRONG_SIGNED_BINARY
	// The image file %1 is strong signed, unable to modify.
	ERROR_EXE_CANNOT_MODIFY_STRONG_SIGNED_BINARY = 218

	// ERROR_FILE_CHECKED_OUT
	// This file is checked out or locked for editing by another user.
	ERROR_FILE_CHECKED_OUT = 220

	// ERROR_CHECKOUT_REQUIRED
	// The file must be checked out before saving changes.
	ERROR_CHECKOUT_REQUIRED = 221

	// ERROR_BAD_FILE_TYPE
	// The file type being saved or retrieved has been blocked.
	ERROR_BAD_FILE_TYPE = 222

	// ERROR_FILE_TOO_LARGE
	// The file size exceeds the limit allowed and cannot be saved.
	ERROR_FILE_TOO_LARGE = 223

	// ERROR_FORMS_AUTH_REQUIRED
	// Access Denied. Before opening files in this location, you must first add the web site to your trusted sites list, browse to the web site, and select the option to login automatically.
	ERROR_FORMS_AUTH_REQUIRED = 224

	// ERROR_VIRUS_INFECTED
	// Operation did not complete successfully because the file contains a virus or potentially unwanted software.
	ERROR_VIRUS_INFECTED = 225

	// ERROR_VIRUS_DELETED
	// This file contains a virus or potentially unwanted software and cannot be opened. Due to the nature of this virus or potentially unwanted software, the file has been removed from this location.
	ERROR_VIRUS_DELETED = 226

	// ERROR_PIPE_LOCA
	// The pipe is local.
	ERROR_PIPE_LOCAL = 229

	// ERROR_BAD_PIPE
	// The pipe state is invalid.
	ERROR_BAD_PIPE = 230

	// ERROR_PIPE_BUSY
	// All pipe instances are busy.
	ERROR_PIPE_BUSY = 231

	// ERROR_NO_DATA
	// The pipe is being closed.
	ERROR_NO_DATA = 232

	// ERROR_PIPE_NOT_CONNECTED
	// No process is on the other end of the pipe.
	ERROR_PIPE_NOT_CONNECTED = 233

	// ERROR_MORE_DATA
	// More data is available.
	ERROR_MORE_DATA = 234

	// ERROR_NO_WORK_DONE
	// The action requested resulted in no work being done. Error-style clean-up has been performed.
	ERROR_NO_WORK_DONE = 235

	// ERROR_VC_DISCONNECTED
	// The session was canceled.
	ERROR_VC_DISCONNECTED = 240

	// ERROR_INVALID_EA_NAME
	// The specified extended attribute name was invalid.
	ERROR_INVALID_EA_NAME = 254

	// ERROR_EA_LIST_INCONSISTENT
	// The extended attributes are inconsistent.
	ERROR_EA_LIST_INCONSISTENT = 255

	// WAIT_TIMEOUT
	// The wait operation timed out.

	WAIT_TIMEOUT = 258

	// ERROR_NO_MORE_ITEMS
	// No more data is available.
	ERROR_NO_MORE_ITEMS = 259

	// ERROR_CANNOT_COPY
	// The copy functions cannot be used.
	ERROR_CANNOT_COPY = 266

	// ERROR_DIRECTORY
	// The directory name is invalid.
	ERROR_DIRECTORY = 267

	// ERROR_EAS_DIDNT_FIT
	// The extended attributes did not fit in the buffer.
	ERROR_EAS_DIDNT_FIT = 275

	// ERROR_EA_FILE_CORRUPT
	// The extended attribute file on the mounted file system is corrupt.
	ERROR_EA_FILE_CORRUPT = 276

	// ERROR_EA_TABLE_FUL
	// The extended attribute table file is full.
	ERROR_EA_TABLE_FULL = 277

	// ERROR_INVALID_EA_HANDLE
	// The specified extended attribute handle is invalid.
	ERROR_INVALID_EA_HANDLE = 278

	// ERROR_EAS_NOT_SUPPORTED
	// The mounted file system does not support extended attributes.
	ERROR_EAS_NOT_SUPPORTED = 282

	// ERROR_NOT_OWNER
	// Attempt to release mutex not owned by caller.
	ERROR_NOT_OWNER = 288

	// ERROR_TOO_MANY_POSTS
	// Too many posts were made to a semaphore.
	ERROR_TOO_MANY_POSTS = 298

	// ERROR_PARTIAL_COPY
	// Only part of a ReadProcessMemory or WriteProcessMemory request was completed.
	ERROR_PARTIAL_COPY = 299

	// ERROR_OPLOCK_NOT_GRANTED
	// The oplock request is denied.
	ERROR_OPLOCK_NOT_GRANTED = 300

	// ERROR_INVALID_OPLOCK_PROTOCO
	// An invalid oplock acknowledgment was received by the system.
	ERROR_INVALID_OPLOCK_PROTOCOL = 301

	// ERROR_DISK_TOO_FRAGMENTED
	// The volume is too fragmented to complete this operation.
	ERROR_DISK_TOO_FRAGMENTED = 302

	// ERROR_DELETE_PENDING
	// The file cannot be opened because it is in the process of being deleted.
	ERROR_DELETE_PENDING = 303

	// ERROR_INCOMPATIBLE_WITH_GLOBAL_SHORT_NAME_REGISTRY_SETTING
	// Short name settings may not be changed on this volume due to the global registry setting.
	ERROR_INCOMPATIBLE_WITH_GLOBAL_SHORT_NAME_REGISTRY_SETTING = 304

	// ERROR_SHORT_NAMES_NOT_ENABLED_ON_VOLUME
	// Short names are not enabled on this volume.
	ERROR_SHORT_NAMES_NOT_ENABLED_ON_VOLUME = 305

	// ERROR_SECURITY_STREAM_IS_INCONSISTENT
	// The security stream for the given volume is in an inconsistent state.
	// Please run CHKDSK on the volume.
	ERROR_SECURITY_STREAM_IS_INCONSISTENT = 306

	// ERROR_INVALID_LOCK_RANGE
	// A requested file lock operation cannot be processed due to an invalid byte range.
	ERROR_INVALID_LOCK_RANGE = 307

	// ERROR_IMAGE_SUBSYSTEM_NOT_PRESENT
	// The subsystem needed to support the image type is not present.
	ERROR_IMAGE_SUBSYSTEM_NOT_PRESENT = 308

	// ERROR_NOTIFICATION_GUID_ALREADY_DEFINED
	// The specified file already has a notification GUID associated with it.
	ERROR_NOTIFICATION_GUID_ALREADY_DEFINED = 309

	// ERROR_INVALID_EXCEPTION_HANDLER
	// An invalid exception handler routine has been detected.
	ERROR_INVALID_EXCEPTION_HANDLER = 310

	// ERROR_DUPLICATE_PRIVILEGES
	// Duplicate privileges were specified for the token.
	ERROR_DUPLICATE_PRIVILEGES = 311

	// ERROR_NO_RANGES_PROCESSED
	// No ranges for the specified operation were able to be processed.
	ERROR_NO_RANGES_PROCESSED = 312

	// ERROR_NOT_ALLOWED_ON_SYSTEM_FILE
	// Operation is not allowed on a file system internal file.
	ERROR_NOT_ALLOWED_ON_SYSTEM_FILE = 313

	// ERROR_DISK_RESOURCES_EXHAUSTED
	// The physical resources of this disk have been exhausted.
	ERROR_DISK_RESOURCES_EXHAUSTED = 314

	// ERROR_INVALID_TOKEN
	// The token representing the data is invalid.
	ERROR_INVALID_TOKEN = 315

	// ERROR_DEVICE_FEATURE_NOT_SUPPORTED
	// The device does not support the command feature.
	ERROR_DEVICE_FEATURE_NOT_SUPPORTED = 316

	// ERROR_MR_MID_NOT_FOUND
	// The system cannot find message text for message number=0x%1 in the message file for %2.
	ERROR_MR_MID_NOT_FOUND = 317

	// ERROR_SCOPE_NOT_FOUND
	// The scope specified was not found.
	ERROR_SCOPE_NOT_FOUND = 318

	// ERROR_UNDEFINED_SCOPE
	// The Central Access Policy specified is not defined on the target machine.
	ERROR_UNDEFINED_SCOPE = 319

	// ERROR_INVALID_CAP
	// The Central Access Policy obtained from Active Directory is invalid.
	ERROR_INVALID_CAP = 320

	// ERROR_DEVICE_UNREACHABLE
	// The device is unreachable.
	ERROR_DEVICE_UNREACHABLE = 321

	// ERROR_DEVICE_NO_RESOURCES
	// The target device has insufficient resources to complete the operation.
	ERROR_DEVICE_NO_RESOURCES = 322

	// ERROR_DATA_CHECKSUM_ERROR
	// A data integrity checksum error occurred. Data in the file stream is corrupt.
	ERROR_DATA_CHECKSUM_ERROR = 323

	// ERROR_INTERMIXED_KERNEL_EA_OPERATION
	// An attempt was made to modify both a KERNEL and normal Extended Attribute (EA) in the same operation.
	ERROR_INTERMIXED_KERNEL_EA_OPERATION = 324

	// ERROR_FILE_LEVEL_TRIM_NOT_SUPPORTED
	// Device does not support file-level TRIM.
	ERROR_FILE_LEVEL_TRIM_NOT_SUPPORTED = 326

	// ERROR_OFFSET_ALIGNMENT_VIOLATION
	// The command specified a data offset that does not align to the device's granularity/alignment.
	ERROR_OFFSET_ALIGNMENT_VIOLATION = 327

	// ERROR_INVALID_FIELD_IN_PARAMETER_LIST
	// The command specified an invalid field in its parameter list.
	ERROR_INVALID_FIELD_IN_PARAMETER_LIST = 328

	// ERROR_OPERATION_IN_PROGRESS
	// An operation is currently in progress with the device.
	ERROR_OPERATION_IN_PROGRESS = 329

	// ERROR_BAD_DEVICE_PATH
	// An attempt was made to send down the command via an invalid path to the target device.
	ERROR_BAD_DEVICE_PATH = 330

	// ERROR_TOO_MANY_DESCRIPTORS
	// The command specified a number of descriptors that exceeded the maximum supported by the device.
	ERROR_TOO_MANY_DESCRIPTORS = 331

	// ERROR_SCRUB_DATA_DISABLED
	// Scrub is disabled on the specified file.
	ERROR_SCRUB_DATA_DISABLED = 332

	// ERROR_NOT_REDUNDANT_STORAGE
	// The storage device does not provide redundancy.
	ERROR_NOT_REDUNDANT_STORAGE = 333

	// ERROR_RESIDENT_FILE_NOT_SUPPORTED
	// An operation is not supported on a resident file.
	ERROR_RESIDENT_FILE_NOT_SUPPORTED = 334

	// ERROR_COMPRESSED_FILE_NOT_SUPPORTED
	// An operation is not supported on a compressed file.
	ERROR_COMPRESSED_FILE_NOT_SUPPORTED = 335

	// ERROR_DIRECTORY_NOT_SUPPORTED
	// An operation is not supported on a directory.
	ERROR_DIRECTORY_NOT_SUPPORTED = 336

	// ERROR_NOT_READ_FROM_COPY
	// The specified copy of the requested data could not be read.
	ERROR_NOT_READ_FROM_COPY = 337

	// ERROR_FT_WRITE_FAILURE
	// The specified data could not be written to any of the copies.
	ERROR_FT_WRITE_FAILURE = 338

	// ERROR_FT_DI_SCAN_REQUIRED
	// One or more copies of data on this device may be out of sync. No writes may be performed until a data integrity scan is completed.
	ERROR_FT_DI_SCAN_REQUIRED = 339

	// ERROR_INVALID_KERNEL_INFO_VERSION
	// The supplied kernel information version is invalid.
	ERROR_INVALID_KERNEL_INFO_VERSION = 340

	// ERROR_INVALID_PEP_INFO_VERSION
	// The supplied PEP information version is invalid.
	ERROR_INVALID_PEP_INFO_VERSION = 341

	// ERROR_OBJECT_NOT_EXTERNALLY_BACKED
	// This object is not externally backed by any provider.
	ERROR_OBJECT_NOT_EXTERNALLY_BACKED = 342

	// ERROR_EXTERNAL_BACKING_PROVIDER_UNKNOWN
	// The external backing provider is not recognized.
	ERROR_EXTERNAL_BACKING_PROVIDER_UNKNOWN = 343

	// ERROR_COMPRESSION_NOT_BENEFICIA
	// Compressing this object would not save space.
	ERROR_COMPRESSION_NOT_BENEFICIAL = 344

	// ERROR_STORAGE_TOPOLOGY_ID_MISMATCH
	// The request failed due to a storage topology ID mismatch.
	ERROR_STORAGE_TOPOLOGY_ID_MISMATCH = 345

	// ERROR_BLOCKED_BY_PARENTAL_CONTROLS
	// The operation was blocked by parental controls.
	ERROR_BLOCKED_BY_PARENTAL_CONTROLS = 346

	// ERROR_BLOCK_TOO_MANY_REFERENCES
	// A file system block being referenced has already reached the maximum reference count and can't be referenced any further.
	ERROR_BLOCK_TOO_MANY_REFERENCES = 347

	// ERROR_MARKED_TO_DISALLOW_WRITES
	// The requested operation failed because the file stream is marked to disallow writes.
	ERROR_MARKED_TO_DISALLOW_WRITES = 348

	// ERROR_ENCLAVE_FAILURE
	// The requested operation failed with an architecture-specific failure code.
	ERROR_ENCLAVE_FAILURE = 349

	// ERROR_FAIL_NOACTION_REBOOT
	// No action was taken as a system reboot is required.
	ERROR_FAIL_NOACTION_REBOOT = 350

	// ERROR_FAIL_SHUTDOWN
	// The shutdown operation failed.
	ERROR_FAIL_SHUTDOWN = 351

	// ERROR_FAIL_RESTART
	// The restart operation failed.
	ERROR_FAIL_RESTART = 352

	// ERROR_MAX_SESSIONS_REACHED
	// The maximum number of sessions has been reached.
	ERROR_MAX_SESSIONS_REACHED = 353

	// ERROR_NETWORK_ACCESS_DENIED_EDP
	// Windows Information Protection policy does not allow access to this network resource.
	ERROR_NETWORK_ACCESS_DENIED_EDP = 354

	// ERROR_DEVICE_HINT_NAME_BUFFER_TOO_SMAL
	// The device hint name buffer is too small to receive the remaining name.
	ERROR_DEVICE_HINT_NAME_BUFFER_TOO_SMALL = 355

	// ERROR_EDP_POLICY_DENIES_OPERATION
	// The requested operation was blocked by Windows Information Protection policy. For more information, contact your system administrator.
	ERROR_EDP_POLICY_DENIES_OPERATION = 356

	// ERROR_EDP_DPL_POLICY_CANT_BE_SATISFIED
	// The requested operation cannot be performed because hardware or software configuration of the device does not comply with Windows Information Protection under Lock policy. Please, verify that user PIN has been created. For more information, contact your system administrator.
	ERROR_EDP_DPL_POLICY_CANT_BE_SATISFIED = 357

	// ERROR_CLOUD_FILE_SYNC_ROOT_METADATA_CORRUPT
	// The cloud sync root metadata is corrupted.
	ERROR_CLOUD_FILE_SYNC_ROOT_METADATA_CORRUPT = 358

	// ERROR_DEVICE_IN_MAINTENANCE
	// The device is in maintenance mode.
	ERROR_DEVICE_IN_MAINTENANCE = 359

	// ERROR_NOT_SUPPORTED_ON_DAX
	// This operation is not supported on a DAX volume.
	ERROR_NOT_SUPPORTED_ON_DAX = 360

	// ERROR_DAX_MAPPING_EXISTS
	// The volume has active DAX mappings.
	ERROR_DAX_MAPPING_EXISTS = 361

	// ERROR_CLOUD_FILE_PROVIDER_NOT_RUNNING
	// The cloud file provider is not running.
	ERROR_CLOUD_FILE_PROVIDER_NOT_RUNNING = 362

	// ERROR_CLOUD_FILE_METADATA_CORRUPT
	// The cloud file metadata is corrupt and unreadable.
	ERROR_CLOUD_FILE_METADATA_CORRUPT = 363

	// ERROR_CLOUD_FILE_METADATA_TOO_LARGE
	// The cloud file metadata is too large.
	ERROR_CLOUD_FILE_METADATA_TOO_LARGE = 364

	// ERROR_CLOUD_FILE_PROPERTY_BLOB_TOO_LARGE
	// The cloud file property is too large.
	ERROR_CLOUD_FILE_PROPERTY_BLOB_TOO_LARGE = 365

	// ERROR_CLOUD_FILE_PROPERTY_BLOB_CHECKSUM_MISMATCH
	// The cloud file property is possibly corrupt. The on-disk checksum does not match the computed checksum.
	ERROR_CLOUD_FILE_PROPERTY_BLOB_CHECKSUM_MISMATCH = 366

	// ERROR_CHILD_PROCESS_BLOCKED
	// The process creation has been blocked.
	ERROR_CHILD_PROCESS_BLOCKED = 367

	// ERROR_STORAGE_LOST_DATA_PERSISTENCE
	// The storage device has lost data or persistence.
	ERROR_STORAGE_LOST_DATA_PERSISTENCE = 368

	// ERROR_FILE_SYSTEM_VIRTUALIZATION_UNAVAILABLE
	// The provider that supports file system virtualization is temporarily unavailable.
	ERROR_FILE_SYSTEM_VIRTUALIZATION_UNAVAILABLE = 369

	// ERROR_FILE_SYSTEM_VIRTUALIZATION_METADATA_CORRUPT
	// The metadata for file system virtualization is corrupt and unreadable.
	ERROR_FILE_SYSTEM_VIRTUALIZATION_METADATA_CORRUPT = 370

	// ERROR_FILE_SYSTEM_VIRTUALIZATION_BUSY
	// The provider that supports file system virtualization is too busy to complete this operation.
	ERROR_FILE_SYSTEM_VIRTUALIZATION_BUSY = 371

	// ERROR_FILE_SYSTEM_VIRTUALIZATION_PROVIDER_UNKNOWN
	// The provider that supports file system virtualization is unknown.
	ERROR_FILE_SYSTEM_VIRTUALIZATION_PROVIDER_UNKNOWN = 372

	// ERROR_GDI_HANDLE_LEAK
	// GDI handles were potentially leaked by the application.
	ERROR_GDI_HANDLE_LEAK = 373

	// ERROR_CLOUD_FILE_TOO_MANY_PROPERTY_BLOBS
	// The maximum number of cloud file properties has been reached.
	ERROR_CLOUD_FILE_TOO_MANY_PROPERTY_BLOBS = 374

	// ERROR_CLOUD_FILE_PROPERTY_VERSION_NOT_SUPPORTED
	// The version of the cloud file property store is not supported.
	ERROR_CLOUD_FILE_PROPERTY_VERSION_NOT_SUPPORTED = 375

	// ERROR_NOT_A_CLOUD_FILE
	// The file is not a cloud file.
	ERROR_NOT_A_CLOUD_FILE = 376

	// ERROR_CLOUD_FILE_NOT_IN_SYNC
	// The file is not in sync with the cloud.
	ERROR_CLOUD_FILE_NOT_IN_SYNC = 377

	// ERROR_CLOUD_FILE_ALREADY_CONNECTED
	// The cloud sync root is already connected with another cloud sync provider.
	ERROR_CLOUD_FILE_ALREADY_CONNECTED = 378

	// ERROR_CLOUD_FILE_NOT_SUPPORTED
	// The operation is not supported by the cloud sync provider.
	ERROR_CLOUD_FILE_NOT_SUPPORTED = 379

	// ERROR_CLOUD_FILE_INVALID_REQUEST
	// The cloud operation is invalid.
	ERROR_CLOUD_FILE_INVALID_REQUEST = 380

	// ERROR_CLOUD_FILE_READ_ONLY_VOLUME
	// The cloud operation is not supported on a read-only volume.
	ERROR_CLOUD_FILE_READ_ONLY_VOLUME = 381

	// ERROR_CLOUD_FILE_CONNECTED_PROVIDER_ONLY
	// The operation is reserved for a connected cloud sync provider.
	ERROR_CLOUD_FILE_CONNECTED_PROVIDER_ONLY = 382

	// ERROR_CLOUD_FILE_VALIDATION_FAILED
	// The cloud sync provider failed to validate the downloaded data.
	ERROR_CLOUD_FILE_VALIDATION_FAILED = 383

	// ERROR_SMB1_NOT_AVAILABLE
	// You can't connect to the file share because it's not secure. This share requires the obsolete SMB1 protocol, which is unsafe and could expose your system to attack.
	// Your system requires SMB2 or higher. For more info on resolving this issue, see: https://go.microsoft.com/fwlink/?linkid=852747
	ERROR_SMB1_NOT_AVAILABLE = 384

	// ERROR_FILE_SYSTEM_VIRTUALIZATION_INVALID_OPERATION
	// The virtualization operation is not allowed on the file in its current state.
	ERROR_FILE_SYSTEM_VIRTUALIZATION_INVALID_OPERATION = 385

	// ERROR_CLOUD_FILE_AUTHENTICATION_FAILED
	// The cloud sync provider failed user authentication.
	ERROR_CLOUD_FILE_AUTHENTICATION_FAILED = 386

	// ERROR_CLOUD_FILE_INSUFFICIENT_RESOURCES
	// The cloud sync provider failed to perform the operation due to low system resources.
	ERROR_CLOUD_FILE_INSUFFICIENT_RESOURCES = 387

	// ERROR_CLOUD_FILE_NETWORK_UNAVAILABLE
	// The cloud sync provider failed to perform the operation due to network being unavailable.
	ERROR_CLOUD_FILE_NETWORK_UNAVAILABLE = 388

	// ERROR_CLOUD_FILE_UNSUCCESSFU
	// The cloud operation was unsuccessful.
	ERROR_CLOUD_FILE_UNSUCCESSFUL = 389

	// ERROR_CLOUD_FILE_NOT_UNDER_SYNC_ROOT
	// The operation is only supported on files under a cloud sync root.
	ERROR_CLOUD_FILE_NOT_UNDER_SYNC_ROOT = 390

	// ERROR_CLOUD_FILE_IN_USE
	// The operation cannot be performed on cloud files in use.
	ERROR_CLOUD_FILE_IN_USE = 391

	// ERROR_CLOUD_FILE_PINNED
	// The operation cannot be performed on pinned cloud files.
	ERROR_CLOUD_FILE_PINNED = 392

	// ERROR_CLOUD_FILE_REQUEST_ABORTED
	// The cloud operation was aborted.
	ERROR_CLOUD_FILE_REQUEST_ABORTED = 393

	// ERROR_CLOUD_FILE_PROPERTY_CORRUPT
	// The cloud file's property store is corrupt.
	ERROR_CLOUD_FILE_PROPERTY_CORRUPT = 394

	// ERROR_CLOUD_FILE_ACCESS_DENIED
	// Access to the cloud file is denied.
	ERROR_CLOUD_FILE_ACCESS_DENIED = 395

	// ERROR_CLOUD_FILE_INCOMPATIBLE_HARDLINKS
	// The cloud operation cannot be performed on a file with incompatible hardlinks.
	ERROR_CLOUD_FILE_INCOMPATIBLE_HARDLINKS = 396

	// ERROR_CLOUD_FILE_PROPERTY_LOCK_CONFLICT
	// The operation failed due to a conflicting cloud file property lock.
	ERROR_CLOUD_FILE_PROPERTY_LOCK_CONFLICT = 397

	// ERROR_CLOUD_FILE_REQUEST_CANCELED
	// The cloud operation was canceled by user.
	ERROR_CLOUD_FILE_REQUEST_CANCELED = 398

	// ERROR_EXTERNAL_SYSKEY_NOT_SUPPORTED
	// An externally encrypted syskey has been configured, but the system no longer supports this feature.  Please see https://go.microsoft.com/fwlink/?linkid=851152 for more information.
	ERROR_EXTERNAL_SYSKEY_NOT_SUPPORTED = 399

	// ERROR_THREAD_MODE_ALREADY_BACKGROUND
	// The thread is already in background processing mode.
	ERROR_THREAD_MODE_ALREADY_BACKGROUND = 400

	// ERROR_THREAD_MODE_NOT_BACKGROUND
	// The thread is not in background processing mode.
	ERROR_THREAD_MODE_NOT_BACKGROUND = 401

	// ERROR_PROCESS_MODE_ALREADY_BACKGROUND
	// The process is already in background processing mode.
	ERROR_PROCESS_MODE_ALREADY_BACKGROUND = 402

	// ERROR_PROCESS_MODE_NOT_BACKGROUND
	// The process is not in background processing mode.
	ERROR_PROCESS_MODE_NOT_BACKGROUND = 403

	// ERROR_CLOUD_FILE_PROVIDER_TERMINATED
	// The cloud file provider exited unexpectedly.
	ERROR_CLOUD_FILE_PROVIDER_TERMINATED = 404

	// ERROR_NOT_A_CLOUD_SYNC_ROOT
	// The file is not a cloud sync root.
	ERROR_NOT_A_CLOUD_SYNC_ROOT = 405

	// ERROR_FILE_PROTECTED_UNDER_DP
	// The read or write operation to an encrypted file could not be completed because the file can only be accessed when the device is unlocked.
	ERROR_FILE_PROTECTED_UNDER_DPL = 406

	// ERROR_VOLUME_NOT_CLUSTER_ALIGNED
	// The volume is not cluster aligned on the disk.
	ERROR_VOLUME_NOT_CLUSTER_ALIGNED = 407

	// ERROR_NO_PHYSICALLY_ALIGNED_FREE_SPACE_FOUND
	// No physically aligned free space was found on the volume.
	ERROR_NO_PHYSICALLY_ALIGNED_FREE_SPACE_FOUND = 408

	// ERROR_APPX_FILE_NOT_ENCRYPTED
	// The APPX file can not be accessed because it is not encrypted as expected.
	ERROR_APPX_FILE_NOT_ENCRYPTED = 409

	// ERROR_RWRAW_ENCRYPTED_FILE_NOT_ENCRYPTED
	// A read or write of raw encrypted data cannot be performed because the file is not encrypted.
	ERROR_RWRAW_ENCRYPTED_FILE_NOT_ENCRYPTED = 410

	// ERROR_RWRAW_ENCRYPTED_INVALID_EDATAINFO_FILEOFFSET
	// An invalid file offset in the encrypted data info block was passed for read or write operation of file's raw encrypted data.
	ERROR_RWRAW_ENCRYPTED_INVALID_EDATAINFO_FILEOFFSET = 411

	// ERROR_RWRAW_ENCRYPTED_INVALID_EDATAINFO_FILERANGE
	// An invalid offset and length combination in the encrypted data info block was passed for read or write operation of file's raw encrypted data.
	ERROR_RWRAW_ENCRYPTED_INVALID_EDATAINFO_FILERANGE = 412

	// ERROR_RWRAW_ENCRYPTED_INVALID_EDATAINFO_PARAMETER
	// An invalid parameter in the encrypted data info block was passed for read or write operation of file's raw encrypted data.
	ERROR_RWRAW_ENCRYPTED_INVALID_EDATAINFO_PARAMETER = 413

	// ERROR_LINUX_SUBSYSTEM_NOT_PRESENT
	// The Windows Subsystem for Linux has not been enabled.
	ERROR_LINUX_SUBSYSTEM_NOT_PRESENT = 414

	// ERROR_FT_READ_FAILURE
	// The specified data could not be read from any of the copies.
	ERROR_FT_READ_FAILURE = 415

	// ERROR_STORAGE_RESERVE_ID_INVALID
	// The specified storage reserve ID is invalid.
	ERROR_STORAGE_RESERVE_ID_INVALID = 416

	// ERROR_STORAGE_RESERVE_DOES_NOT_EXIST
	// The specified storage reserve does not exist.
	ERROR_STORAGE_RESERVE_DOES_NOT_EXIST = 417

	// ERROR_STORAGE_RESERVE_ALREADY_EXISTS
	// The specified storage reserve already exists.
	ERROR_STORAGE_RESERVE_ALREADY_EXISTS = 418

	// ERROR_STORAGE_RESERVE_NOT_EMPTY
	// The specified storage reserve is not empty.
	ERROR_STORAGE_RESERVE_NOT_EMPTY = 419

	// ERROR_NOT_A_DAX_VOLUME
	// This operation requires a DAX volume.
	ERROR_NOT_A_DAX_VOLUME = 420

	// ERROR_NOT_DAX_MAPPABLE
	// This stream is not DAX mappable.
	ERROR_NOT_DAX_MAPPABLE = 421

	// ERROR_TIME_CRITICAL_THREAD
	// Operation cannot be performed on a time critical thread.
	ERROR_TIME_CRITICAL_THREAD = 422

	// ERROR_DPL_NOT_SUPPORTED_FOR_USER
	// User data protection is not supported for the current or provided user.
	ERROR_DPL_NOT_SUPPORTED_FOR_USER = 423

	// ERROR_CASE_DIFFERING_NAMES_IN_DIR
	// This directory contains entries whose names differ only in case.
	ERROR_CASE_DIFFERING_NAMES_IN_DIR = 424

	// **** Available SYSTEM error codes ****

	///////////////////////////////////////////////////
	//                                               //
	//    Capability Authorization Error codes       //
	//                                               //
	//                =0450 to=0460                  //
	///////////////////////////////////////////////////

	// ERROR_CAPAUTHZ_NOT_DEVUNLOCKED
	// Neither developer unlocked mode nor side loading mode is enabled on the device.
	ERROR_CAPAUTHZ_NOT_DEVUNLOCKED = 450

	// ERROR_CAPAUTHZ_CHANGE_TYPE
	// Can not change application type during upgrade or re-provision.
	ERROR_CAPAUTHZ_CHANGE_TYPE = 451

	// ERROR_CAPAUTHZ_NOT_PROVISIONED
	// The application has not been provisioned.
	ERROR_CAPAUTHZ_NOT_PROVISIONED = 452

	// ERROR_CAPAUTHZ_NOT_AUTHORIZED
	// The requested capability can not be authorized for this application.
	ERROR_CAPAUTHZ_NOT_AUTHORIZED = 453

	// ERROR_CAPAUTHZ_NO_POLICY
	// There is no capability authorization policy on the device.
	ERROR_CAPAUTHZ_NO_POLICY = 454

	// ERROR_CAPAUTHZ_DB_CORRUPTED
	// The capability authorization database has been corrupted.
	ERROR_CAPAUTHZ_DB_CORRUPTED = 455

	// ERROR_CAPAUTHZ_SCCD_INVALID_CATALOG
	// The custom capability's SCCD has an invalid catalog.
	ERROR_CAPAUTHZ_SCCD_INVALID_CATALOG = 456

	// ERROR_CAPAUTHZ_SCCD_NO_AUTH_ENTITY
	// None of the authorized entity elements in the SCCD matched the app being installed; either the PFNs don't match, or the element's signature hash doesn't validate.
	ERROR_CAPAUTHZ_SCCD_NO_AUTH_ENTITY = 457

	// ERROR_CAPAUTHZ_SCCD_PARSE_ERROR
	// The custom capability's SCCD failed to parse.
	ERROR_CAPAUTHZ_SCCD_PARSE_ERROR = 458

	// ERROR_CAPAUTHZ_SCCD_DEV_MODE_REQUIRED
	// The custom capability's SCCD requires developer mode.
	ERROR_CAPAUTHZ_SCCD_DEV_MODE_REQUIRED = 459

	// ERROR_CAPAUTHZ_SCCD_NO_CAPABILITY_MATCH
	// There not all declared custom capabilities are found in the SCCD.
	ERROR_CAPAUTHZ_SCCD_NO_CAPABILITY_MATCH = 460

	// **** Available SYSTEM error codes ****

	// ERROR_PNP_QUERY_REMOVE_DEVICE_TIMEOUT
	// The operation timed out waiting for this device to complete a PnP query-remove request due to a potential hang in its device stack. The system may need to be rebooted to complete the request.
	ERROR_PNP_QUERY_REMOVE_DEVICE_TIMEOUT = 480

	// ERROR_PNP_QUERY_REMOVE_RELATED_DEVICE_TIMEOUT
	// The operation timed out waiting for this device to complete a PnP query-remove request due to a potential hang in the device stack of a related device. The system may need to be rebooted to complete the operation.
	ERROR_PNP_QUERY_REMOVE_RELATED_DEVICE_TIMEOUT = 481

	// ERROR_PNP_QUERY_REMOVE_UNRELATED_DEVICE_TIMEOUT
	// The operation timed out waiting for this device to complete a PnP query-remove request due to a potential hang in the device stack of an unrelated device. The system may need to be rebooted to complete the operation.
	ERROR_PNP_QUERY_REMOVE_UNRELATED_DEVICE_TIMEOUT = 482

	// ERROR_DEVICE_HARDWARE_ERROR
	// The request failed due to a fatal device hardware error.
	ERROR_DEVICE_HARDWARE_ERROR = 483

	// ERROR_INVALID_ADDRESS
	// Attempt to access invalid address.
	ERROR_INVALID_ADDRESS = 487

	// ERROR_VRF_CFG_ENABLED
	// Driver Verifier Volatile settings cannot be set when CFG is enabled.
	ERROR_VRF_CFG_ENABLED = 1183

	// ERROR_PARTITION_TERMINATING
	// An attempt was made to access a partition that has begun termination.
	ERROR_PARTITION_TERMINATING = 1184

	// **** Available SYSTEM error codes ****

	// ERROR_USER_PROFILE_LOAD
	// User profile cannot be loaded.
	ERROR_USER_PROFILE_LOAD = 500

	// **** Available SYSTEM error codes ****

	// ERROR_ARITHMETIC_OVERFLOW
	// Arithmetic result exceeded=32 bits.
	ERROR_ARITHMETIC_OVERFLOW = 534

	// ERROR_PIPE_CONNECTED
	// There is a process on other end of the pipe.
	ERROR_PIPE_CONNECTED = 535

	// ERROR_PIPE_LISTENING
	// Waiting for a process to open the other end of the pipe.
	ERROR_PIPE_LISTENING = 536

	// ERROR_VERIFIER_STOP
	// Application verifier has found an error in the current process.
	ERROR_VERIFIER_STOP = 537

	// ERROR_ABIOS_ERROR
	// An error occurred in the ABIOS subsystem.
	ERROR_ABIOS_ERROR = 538

	// ERROR_WX86_WARNING
	// A warning occurred in the WX86 subsystem.
	ERROR_WX86_WARNING = 539

	// ERROR_WX86_ERROR
	// An error occurred in the WX86 subsystem.
	ERROR_WX86_ERROR = 540

	// ERROR_TIMER_NOT_CANCELED
	// An attempt was made to cancel or set a timer that has an associated APC and the subject thread is not the thread that originally set the timer with an associated APC routine.
	ERROR_TIMER_NOT_CANCELED = 541

	// ERROR_UNWIND
	// Unwind exception code.
	ERROR_UNWIND = 542

	// ERROR_BAD_STACK
	// An invalid or unaligned stack was encountered during an unwind operation.
	ERROR_BAD_STACK = 543

	// ERROR_INVALID_UNWIND_TARGET
	// An invalid unwind target was encountered during an unwind operation.
	ERROR_INVALID_UNWIND_TARGET = 544

	// ERROR_INVALID_PORT_ATTRIBUTES
	// Invalid Object Attributes specified to NtCreatePort or invalid Port Attributes specified to NtConnectPort
	ERROR_INVALID_PORT_ATTRIBUTES = 545

	// ERROR_PORT_MESSAGE_TOO_LONG
	// Length of message passed to NtRequestPort or NtRequestWaitReplyPort was longer than the maximum message allowed by the port.
	ERROR_PORT_MESSAGE_TOO_LONG = 546

	// ERROR_INVALID_QUOTA_LOWER
	// An attempt was made to lower a quota limit below the current usage.
	ERROR_INVALID_QUOTA_LOWER = 547

	// ERROR_DEVICE_ALREADY_ATTACHED
	// An attempt was made to attach to a device that was already attached to another device.
	ERROR_DEVICE_ALREADY_ATTACHED = 548

	// ERROR_INSTRUCTION_MISALIGNMENT
	// An attempt was made to execute an instruction at an unaligned address and the host system does not support unaligned instruction references.
	ERROR_INSTRUCTION_MISALIGNMENT = 549

	// ERROR_PROFILING_NOT_STARTED
	// Profiling not started.
	ERROR_PROFILING_NOT_STARTED = 550

	// ERROR_PROFILING_NOT_STOPPED
	// Profiling not stopped.
	ERROR_PROFILING_NOT_STOPPED = 551

	// ERROR_COULD_NOT_INTERPRET
	// The passed ACL did not contain the minimum required information.
	ERROR_COULD_NOT_INTERPRET = 552

	// ERROR_PROFILING_AT_LIMIT
	// The number of active profiling objects is at the maximum and no more may be started.
	ERROR_PROFILING_AT_LIMIT = 553

	// ERROR_CANT_WAIT
	// Used to indicate that an operation cannot continue without blocking for I/O.
	ERROR_CANT_WAIT = 554

	// ERROR_CANT_TERMINATE_SELF
	// Indicates that a thread attempted to terminate itself by default (called NtTerminateThread with NULL) and it was the last thread in the current process.
	ERROR_CANT_TERMINATE_SELF = 555

	// ERROR_UNEXPECTED_MM_CREATE_ERR
	// If an MM error is returned which is not defined in the standard FsRtl filter, it is converted to one of the following errors which is guaranteed to be in the filter.
	// In this case information is lost, however, the filter correctly handles the exception.
	ERROR_UNEXPECTED_MM_CREATE_ERR = 556

	// ERROR_UNEXPECTED_MM_MAP_ERROR
	// If an MM error is returned which is not defined in the standard FsRtl filter, it is converted to one of the following errors which is guaranteed to be in the filter.
	// In this case information is lost, however, the filter correctly handles the exception.
	ERROR_UNEXPECTED_MM_MAP_ERROR = 557

	// ERROR_UNEXPECTED_MM_EXTEND_ERR
	// If an MM error is returned which is not defined in the standard FsRtl filter, it is converted to one of the following errors which is guaranteed to be in the filter.
	// In this case information is lost, however, the filter correctly handles the exception.
	ERROR_UNEXPECTED_MM_EXTEND_ERR = 558

	// ERROR_BAD_FUNCTION_TABLE
	// A malformed function table was encountered during an unwind operation.
	ERROR_BAD_FUNCTION_TABLE = 559

	// ERROR_NO_GUID_TRANSLATION
	// Indicates that an attempt was made to assign protection to a file system file or directory and one of the SIDs in the security descriptor could not be translated into a GUID that could be stored by the file system.
	// This causes the protection attempt to fail, which may cause a file creation attempt to fail.
	ERROR_NO_GUID_TRANSLATION = 560

	// ERROR_INVALID_LDT_SIZE
	// Indicates that an attempt was made to grow an LDT by setting its size, or that the size was not an even number of selectors.
	ERROR_INVALID_LDT_SIZE = 561

	// ERROR_INVALID_LDT_OFFSET
	// Indicates that the starting value for the LDT information was not an integral multiple of the selector size.
	ERROR_INVALID_LDT_OFFSET = 563

	// ERROR_INVALID_LDT_DESCRIPTOR
	// Indicates that the user supplied an invalid descriptor when trying to set up Ldt descriptors.
	ERROR_INVALID_LDT_DESCRIPTOR = 564

	// ERROR_TOO_MANY_THREADS
	// Indicates a process has too many threads to perform the requested action. For example, assignment of a primary token may only be performed when a process has zero or one threads.
	ERROR_TOO_MANY_THREADS = 565

	// ERROR_THREAD_NOT_IN_PROCESS
	// An attempt was made to operate on a thread within a specific process, but the thread specified is not in the process specified.
	ERROR_THREAD_NOT_IN_PROCESS = 566

	// ERROR_PAGEFILE_QUOTA_EXCEEDED
	// Page file quota was exceeded.
	ERROR_PAGEFILE_QUOTA_EXCEEDED = 567

	// ERROR_LOGON_SERVER_CONFLICT
	// The Netlogon service cannot start because another Netlogon service running in the domain conflicts with the specified role.
	ERROR_LOGON_SERVER_CONFLICT = 568

	// ERROR_SYNCHRONIZATION_REQUIRED
	// The SAM database on a Windows Server is significantly out of synchronization with the copy on the Domain Controller. A complete synchronization is required.
	ERROR_SYNCHRONIZATION_REQUIRED = 569

	// ERROR_NET_OPEN_FAILED
	// The NtCreateFile API failed. This error should never be returned to an application, it is a place holder for the Windows Lan Manager Redirector to use in its internal error mapping routines.
	ERROR_NET_OPEN_FAILED = 570

	// ERROR_IO_PRIVILEGE_FAILED
	// {Privilege Failed}
	// The I/O permissions for the process could not be changed.
	ERROR_IO_PRIVILEGE_FAILED = 571

	// ERROR_CONTROL_C_EXIT
	// {Application Exit by CTRL+C}
	// The application terminated as a result of a CTRL+C.
	ERROR_CONTROL_C_EXIT = 572

	// ERROR_MISSING_SYSTEMFILE
	// {Missing System File}
	// The required system file %hs is bad or missing.
	ERROR_MISSING_SYSTEMFILE = 573

	// ERROR_UNHANDLED_EXCEPTION
	// {Application Error}
	// The exception %s (0x%08lx) occurred in the application at location=0x%08lx.
	ERROR_UNHANDLED_EXCEPTION = 574

	// ERROR_APP_INIT_FAILURE
	// {Application Error}
	// The application was unable to start correctly (0x%lx). Click OK to close the application.
	ERROR_APP_INIT_FAILURE = 575

	// ERROR_PAGEFILE_CREATE_FAILED
	// {Unable to Create Paging File}
	// The creation of the paging file %hs failed (%lx). The requested size was %ld.
	ERROR_PAGEFILE_CREATE_FAILED = 576

	// ERROR_INVALID_IMAGE_HASH
	// Windows cannot verify the digital signature for this file. A recent hardware or software change might have installed a file that is signed incorrectly or damaged, or that might be malicious software from an unknown source.
	ERROR_INVALID_IMAGE_HASH = 577

	// ERROR_NO_PAGEFILE
	// {No Paging File Specified}
	// No paging file was specified in the system configuration.
	ERROR_NO_PAGEFILE = 578

	// ERROR_ILLEGAL_FLOAT_CONTEXT
	// {EXCEPTION}
	// A real-mode application issued a floating-point instruction and floating-point hardware is not present.
	ERROR_ILLEGAL_FLOAT_CONTEXT = 579

	// ERROR_NO_EVENT_PAIR
	// An event pair synchronization operation was performed using the thread specific client/server event pair object, but no event pair object was associated with the thread.
	ERROR_NO_EVENT_PAIR = 580

	// ERROR_DOMAIN_CTRLR_CONFIG_ERROR
	// A Windows Server has an incorrect configuration.
	ERROR_DOMAIN_CTRLR_CONFIG_ERROR = 581

	// ERROR_ILLEGAL_CHARACTER
	// An illegal character was encountered. For a multi-byte character set this includes a lead byte without a succeeding trail byte. For the Unicode character set this includes the characters=0xFFFF and=0xFFFE.
	ERROR_ILLEGAL_CHARACTER = 582

	// ERROR_UNDEFINED_CHARACTER
	// The Unicode character is not defined in the Unicode character set installed on the system.
	ERROR_UNDEFINED_CHARACTER = 583

	// ERROR_FLOPPY_VOLUME
	// The paging file cannot be created on a floppy diskette.
	ERROR_FLOPPY_VOLUME = 584

	// ERROR_BIOS_FAILED_TO_CONNECT_INTERRUPT
	// The system BIOS failed to connect a system interrupt to the device or bus for which the device is connected.
	ERROR_BIOS_FAILED_TO_CONNECT_INTERRUPT = 585

	// ERROR_BACKUP_CONTROLLER
	// This operation is only allowed for the Primary Domain Controller of the domain.
	ERROR_BACKUP_CONTROLLER = 586

	// ERROR_MUTANT_LIMIT_EXCEEDED
	// An attempt was made to acquire a mutant such that its maximum count would have been exceeded.
	ERROR_MUTANT_LIMIT_EXCEEDED = 587

	// ERROR_FS_DRIVER_REQUIRED
	// A volume has been accessed for which a file system driver is required that has not yet been loaded.
	ERROR_FS_DRIVER_REQUIRED = 588

	// ERROR_CANNOT_LOAD_REGISTRY_FILE
	// {Registry File Failure}
	// The registry cannot load the hive (file):
	// %hs
	// or its log or alternate.
	// It is corrupt, absent, or not writable.
	ERROR_CANNOT_LOAD_REGISTRY_FILE = 589

	// ERROR_DEBUG_ATTACH_FAILED
	// {Unexpected Failure in DebugActiveProcess}
	// An unexpected failure occurred while processing a DebugActiveProcess API request. You may choose OK to terminate the process, or Cancel to ignore the error.
	ERROR_DEBUG_ATTACH_FAILED = 590

	// ERROR_SYSTEM_PROCESS_TERMINATED
	// {Fatal System Error}
	// The %hs system process terminated unexpectedly with a status of=0x%08x (0x%08x=0x%08x).
	// The system has been shut down.
	ERROR_SYSTEM_PROCESS_TERMINATED = 591

	// ERROR_DATA_NOT_ACCEPTED
	// {Data Not Accepted}
	// The TDI client could not handle the data received during an indication.
	ERROR_DATA_NOT_ACCEPTED = 592

	// ERROR_VDM_HARD_ERROR
	// NTVDM encountered a hard error.
	ERROR_VDM_HARD_ERROR = 593

	// ERROR_DRIVER_CANCEL_TIMEOUT
	// {Cancel Timeout}
	// The driver %hs failed to complete a cancelled I/O request in the allotted time.
	ERROR_DRIVER_CANCEL_TIMEOUT = 594

	// ERROR_REPLY_MESSAGE_MISMATCH
	// {Reply Message Mismatch}
	// An attempt was made to reply to an LPC message, but the thread specified by the client ID in the message was not waiting on that message.
	ERROR_REPLY_MESSAGE_MISMATCH = 595

	// ERROR_LOST_WRITEBEHIND_DATA
	// {Delayed Write Failed}
	// Windows was unable to save all the data for the file %hs. The data has been lost.
	// This error may be caused by a failure of your computer hardware or network connection. Please try to save this file elsewhere.
	ERROR_LOST_WRITEBEHIND_DATA = 596

	// ERROR_CLIENT_SERVER_PARAMETERS_INVALID
	// The parameter(s) passed to the server in the client/server shared memory window were invalid. Too much data may have been put in the shared memory window.
	ERROR_CLIENT_SERVER_PARAMETERS_INVALID = 597

	// ERROR_NOT_TINY_STREAM
	// The stream is not a tiny stream.
	ERROR_NOT_TINY_STREAM = 598

	// ERROR_STACK_OVERFLOW_READ
	// The request must be handled by the stack overflow code.
	ERROR_STACK_OVERFLOW_READ = 599

	// ERROR_CONVERT_TO_LARGE
	// Internal OFS status codes indicating how an allocation operation is handled. Either it is retried after the containing onode is moved or the extent stream is converted to a large stream.
	ERROR_CONVERT_TO_LARGE = 600

	// ERROR_FOUND_OUT_OF_SCOPE
	// The attempt to find the object found an object matching by ID on the volume but it is out of the scope of the handle used for the operation.
	ERROR_FOUND_OUT_OF_SCOPE = 601

	// ERROR_ALLOCATE_BUCKET
	// The bucket array must be grown. Retry transaction after doing so.
	ERROR_ALLOCATE_BUCKET = 602

	// ERROR_MARSHALL_OVERFLOW
	// The user/kernel marshalling buffer has overflowed.
	ERROR_MARSHALL_OVERFLOW = 603

	// ERROR_INVALID_VARIANT
	// The supplied variant structure contains invalid data.
	ERROR_INVALID_VARIANT = 604

	// ERROR_BAD_COMPRESSION_BUFFER
	// The specified buffer contains ill-formed data.
	ERROR_BAD_COMPRESSION_BUFFER = 605

	// ERROR_AUDIT_FAILED
	// {Audit Failed}
	// An attempt to generate a security audit failed.
	ERROR_AUDIT_FAILED = 606

	// ERROR_TIMER_RESOLUTION_NOT_SET
	// The timer resolution was not previously set by the current process.
	ERROR_TIMER_RESOLUTION_NOT_SET = 607

	// ERROR_INSUFFICIENT_LOGON_INFO
	// There is insufficient account information to log you on.
	ERROR_INSUFFICIENT_LOGON_INFO = 608

	// ERROR_BAD_DLL_ENTRYPOINT
	// {Invalid DLL Entrypoint}
	// The dynamic link library %hs is not written correctly. The stack pointer has been left in an inconsistent state.
	// The entrypoint should be declared as WINAPI or STDCALL. Select YES to fail the DLL load. Select NO to continue execution. Selecting NO may cause the application to operate incorrectly.
	ERROR_BAD_DLL_ENTRYPOINT = 609

	// ERROR_BAD_SERVICE_ENTRYPOINT
	// {Invalid Service Callback Entrypoint}
	// The %hs service is not written correctly. The stack pointer has been left in an inconsistent state.
	// The callback entrypoint should be declared as WINAPI or STDCALL. Selecting OK will cause the service to continue operation. However, the service process may operate incorrectly.
	ERROR_BAD_SERVICE_ENTRYPOINT = 610

	// ERROR_IP_ADDRESS_CONFLICT1
	// There is an IP address conflict with another system on the network
	ERROR_IP_ADDRESS_CONFLICT1 = 611

	// ERROR_IP_ADDRESS_CONFLICT2
	// There is an IP address conflict with another system on the network
	ERROR_IP_ADDRESS_CONFLICT2 = 612

	// ERROR_REGISTRY_QUOTA_LIMIT
	// {Low On Registry Space}
	// The system has reached the maximum size allowed for the system part of the registry. Additional storage requests will be ignored.
	ERROR_REGISTRY_QUOTA_LIMIT = 613

	// ERROR_NO_CALLBACK_ACTIVE
	// A callback return system service cannot be executed when no callback is active.
	ERROR_NO_CALLBACK_ACTIVE = 614

	// ERROR_PWD_TOO_SHORT
	// The password provided is too short to meet the policy of your user account.
	// Please choose a longer password.
	ERROR_PWD_TOO_SHORT = 615

	// ERROR_PWD_TOO_RECENT
	// The policy of your user account does not allow you to change passwords too frequently.
	// This is done to prevent users from changing back to a familiar, but potentially discovered, password.
	// If you feel your password has been compromised then please contact your administrator immediately to have a new one assigned.
	ERROR_PWD_TOO_RECENT = 616

	// ERROR_PWD_HISTORY_CONFLICT
	// You have attempted to change your password to one that you have used in the past.
	// The policy of your user account does not allow this. Please select a password that you have not previously used.
	ERROR_PWD_HISTORY_CONFLICT = 617

	// ERROR_UNSUPPORTED_COMPRESSION
	// The specified compression format is unsupported.
	ERROR_UNSUPPORTED_COMPRESSION = 618

	// ERROR_INVALID_HW_PROFILE
	// The specified hardware profile configuration is invalid.
	ERROR_INVALID_HW_PROFILE = 619

	// ERROR_INVALID_PLUGPLAY_DEVICE_PATH
	// The specified Plug and Play registry device path is invalid.
	ERROR_INVALID_PLUGPLAY_DEVICE_PATH = 620

	// ERROR_QUOTA_LIST_INCONSISTENT
	// The specified quota list is internally inconsistent with its descriptor.
	ERROR_QUOTA_LIST_INCONSISTENT = 621

	// ERROR_EVALUATION_EXPIRATION
	// {Windows Evaluation Notification}
	// The evaluation period for this installation of Windows has expired. This system will shutdown in=1 hour. To restore access to this installation of Windows, please upgrade this installation using a licensed distribution of this product.
	ERROR_EVALUATION_EXPIRATION = 622

	// ERROR_ILLEGAL_DLL_RELOCATION
	// {Illegal System DLL Relocation}
	// The system DLL %hs was relocated in memory. The application will not run properly.
	// The relocation occurred because the DLL %hs occupied an address range reserved for Windows system DLLs. The vendor supplying the DLL should be contacted for a new DLL.
	ERROR_ILLEGAL_DLL_RELOCATION = 623

	// ERROR_DLL_INIT_FAILED_LOGOFF
	// {DLL Initialization Failed}
	// The application failed to initialize because the window station is shutting down.
	ERROR_DLL_INIT_FAILED_LOGOFF = 624

	// ERROR_VALIDATE_CONTINUE
	// The validation process needs to continue on to the next step.
	ERROR_VALIDATE_CONTINUE = 625

	// ERROR_NO_MORE_MATCHES
	// There are no more matches for the current index enumeration.
	ERROR_NO_MORE_MATCHES = 626

	// ERROR_RANGE_LIST_CONFLICT
	// The range could not be added to the range list because of a conflict.
	ERROR_RANGE_LIST_CONFLICT = 627

	// ERROR_SERVER_SID_MISMATCH
	// The server process is running under a SID different than that required by client.
	ERROR_SERVER_SID_MISMATCH = 628

	// ERROR_CANT_ENABLE_DENY_ONLY
	// A group marked use for deny only cannot be enabled.
	ERROR_CANT_ENABLE_DENY_ONLY = 629

	// ERROR_FLOAT_MULTIPLE_FAULTS
	// {EXCEPTION}
	// Multiple floating point faults.
	ERROR_FLOAT_MULTIPLE_FAULTS = 630

	// ERROR_FLOAT_MULTIPLE_TRAPS
	// {EXCEPTION}
	// Multiple floating point traps.
	ERROR_FLOAT_MULTIPLE_TRAPS = 631

	// ERROR_NOINTERFACE
	// The requested interface is not supported.
	ERROR_NOINTERFACE = 632

	// ERROR_DRIVER_FAILED_SLEEP
	// {System Standby Failed}
	// The driver %hs does not support standby mode. Updating this driver may allow the system to go to standby mode.
	ERROR_DRIVER_FAILED_SLEEP = 633

	// ERROR_CORRUPT_SYSTEM_FILE
	// The system file %1 has become corrupt and has been replaced.
	ERROR_CORRUPT_SYSTEM_FILE = 634

	// ERROR_COMMITMENT_MINIMUM
	// {Virtual Memory Minimum Too Low}
	// Your system is low on virtual memory. Windows is increasing the size of your virtual memory paging file.
	// During this process, memory requests for some applications may be denied. For more information, see Help.
	ERROR_COMMITMENT_MINIMUM = 635

	// ERROR_PNP_RESTART_ENUMERATION
	// A device was removed so enumeration must be restarted.
	ERROR_PNP_RESTART_ENUMERATION = 636

	// ERROR_SYSTEM_IMAGE_BAD_SIGNATURE
	// {Fatal System Error}
	// The system image %s is not properly signed.
	// The file has been replaced with the signed file.
	// The system has been shut down.
	ERROR_SYSTEM_IMAGE_BAD_SIGNATURE = 637

	// ERROR_PNP_REBOOT_REQUIRED
	// Device will not start without a reboot.
	ERROR_PNP_REBOOT_REQUIRED = 638

	// ERROR_INSUFFICIENT_POWER
	// There is not enough power to complete the requested operation.
	ERROR_INSUFFICIENT_POWER = 639

	// ERROR_MULTIPLE_FAULT_VIOLATION
	//  ERROR_MULTIPLE_FAULT_VIOLATION
	ERROR_MULTIPLE_FAULT_VIOLATION = 640

	// ERROR_SYSTEM_SHUTDOWN
	// The system is in the process of shutting down.
	ERROR_SYSTEM_SHUTDOWN = 641

	// ERROR_PORT_NOT_SET
	// An attempt to remove a processes DebugPort was made, but a port was not already associated with the process.
	ERROR_PORT_NOT_SET = 642

	// ERROR_DS_VERSION_CHECK_FAILURE
	// This version of Windows is not compatible with the behavior version of directory forest, domain or domain controller.
	ERROR_DS_VERSION_CHECK_FAILURE = 643

	// ERROR_RANGE_NOT_FOUND
	// The specified range could not be found in the range list.
	ERROR_RANGE_NOT_FOUND = 644

	// ERROR_NOT_SAFE_MODE_DRIVER
	// The driver was not loaded because the system is booting into safe mode.
	ERROR_NOT_SAFE_MODE_DRIVER = 646

	// ERROR_FAILED_DRIVER_ENTRY
	// The driver was not loaded because it failed its initialization call.
	ERROR_FAILED_DRIVER_ENTRY = 647

	// ERROR_DEVICE_ENUMERATION_ERROR
	// The "%hs" encountered an error while applying power or reading the device configuration.
	// This may be caused by a failure of your hardware or by a poor connection.
	ERROR_DEVICE_ENUMERATION_ERROR = 648

	// ERROR_MOUNT_POINT_NOT_RESOLVED
	// The create operation failed because the name contained at least one mount point which resolves to a volume to which the specified device object is not attached.
	ERROR_MOUNT_POINT_NOT_RESOLVED = 649

	// ERROR_INVALID_DEVICE_OBJECT_PARAMETER
	// The device object parameter is either not a valid device object or is not attached to the volume specified by the file name.
	ERROR_INVALID_DEVICE_OBJECT_PARAMETER = 650

	// ERROR_MCA_OCCURED
	// A Machine Check Error has occurred. Please check the system eventlog for additional information.
	ERROR_MCA_OCCURED = 651

	// ERROR_DRIVER_DATABASE_ERROR
	// There was error [%2] processing the driver database.
	ERROR_DRIVER_DATABASE_ERROR = 652

	// ERROR_SYSTEM_HIVE_TOO_LARGE
	// System hive size has exceeded its limit.
	ERROR_SYSTEM_HIVE_TOO_LARGE = 653

	// ERROR_DRIVER_FAILED_PRIOR_UNLOAD
	// The driver could not be loaded because a previous version of the driver is still in memory.
	ERROR_DRIVER_FAILED_PRIOR_UNLOAD = 654

	// ERROR_VOLSNAP_PREPARE_HIBERNATE
	// {Volume Shadow Copy Service}
	// Please wait while the Volume Shadow Copy Service prepares volume %hs for hibernation.
	ERROR_VOLSNAP_PREPARE_HIBERNATE = 655

	// ERROR_HIBERNATION_FAILURE
	// The system has failed to hibernate (The error code is %hs). Hibernation will be disabled until the system is restarted.
	ERROR_HIBERNATION_FAILURE = 656

	// ERROR_PWD_TOO_LONG
	// The password provided is too long to meet the policy of your user account.
	// Please choose a shorter password.
	ERROR_PWD_TOO_LONG = 657

	// ERROR_FILE_SYSTEM_LIMITATION
	// The requested operation could not be completed due to a file system limitation
	ERROR_FILE_SYSTEM_LIMITATION = 665

	// ERROR_ASSERTION_FAILURE
	// An assertion failure has occurred.
	ERROR_ASSERTION_FAILURE = 668

	// ERROR_ACPI_ERROR
	// An error occurred in the ACPI subsystem.
	ERROR_ACPI_ERROR = 669

	// ERROR_WOW_ASSERTION
	// WOW Assertion Error.
	ERROR_WOW_ASSERTION = 670

	// ERROR_PNP_BAD_MPS_TABLE
	// A device is missing in the system BIOS MPS table. This device will not be used.
	// Please contact your system vendor for system BIOS update.
	ERROR_PNP_BAD_MPS_TABLE = 671

	// ERROR_PNP_TRANSLATION_FAILED
	// A translator failed to translate resources.
	ERROR_PNP_TRANSLATION_FAILED = 672

	// ERROR_PNP_IRQ_TRANSLATION_FAILED
	// A IRQ translator failed to translate resources.
	ERROR_PNP_IRQ_TRANSLATION_FAILED = 673

	// ERROR_PNP_INVALID_ID
	// Driver %2 returned invalid ID for a child device (%3).
	ERROR_PNP_INVALID_ID = 674

	// ERROR_WAKE_SYSTEM_DEBUGGER
	// {Kernel Debugger Awakened}
	// the system debugger was awakened by an interrupt.
	ERROR_WAKE_SYSTEM_DEBUGGER = 675

	// ERROR_HANDLES_CLOSED
	// {Handles Closed}
	// Handles to objects have been automatically closed as a result of the requested operation.
	ERROR_HANDLES_CLOSED = 676

	// ERROR_EXTRANEOUS_INFORMATION
	// {Too Much Information}
	// The specified access control list (ACL) contained more information than was expected.
	ERROR_EXTRANEOUS_INFORMATION = 677

	// ERROR_RXACT_COMMIT_NECESSARY
	// This warning level status indicates that the transaction state already exists for the registry sub-tree, but that a transaction commit was previously aborted.
	// The commit has NOT been completed, but has not been rolled back either (so it may still be committed if desired).
	ERROR_RXACT_COMMIT_NECESSARY = 678

	// ERROR_MEDIA_CHECK
	// {Media Changed}
	// The media may have changed.
	ERROR_MEDIA_CHECK = 679

	// ERROR_GUID_SUBSTITUTION_MADE
	// {GUID Substitution}
	// During the translation of a global identifier (GUID) to a Windows security ID (SID), no administratively-defined GUID prefix was found.
	// A substitute prefix was used, which will not compromise system security. However, this may provide a more restrictive access than intended.
	ERROR_GUID_SUBSTITUTION_MADE = 680

	// ERROR_STOPPED_ON_SYMLINK
	// The create operation stopped after reaching a symbolic link
	ERROR_STOPPED_ON_SYMLINK = 681

	// ERROR_LONGJUMP
	// A long jump has been executed.
	ERROR_LONGJUMP = 682

	// ERROR_PLUGPLAY_QUERY_VETOED
	// The Plug and Play query operation was not successful.
	ERROR_PLUGPLAY_QUERY_VETOED = 683

	// ERROR_UNWIND_CONSOLIDATE
	// A frame consolidation has been executed.
	ERROR_UNWIND_CONSOLIDATE = 684

	// ERROR_REGISTRY_HIVE_RECOVERED
	// {Registry Hive Recovered}
	// Registry hive (file):
	// %hs
	// was corrupted and it has been recovered. Some data might have been lost.
	ERROR_REGISTRY_HIVE_RECOVERED = 685

	// ERROR_DLL_MIGHT_BE_INSECURE
	// The application is attempting to run executable code from the module %hs. This may be insecure. An alternative, %hs, is available. Should the application use the secure module %hs?
	ERROR_DLL_MIGHT_BE_INSECURE = 686

	// ERROR_DLL_MIGHT_BE_INCOMPATIBLE
	// The application is loading executable code from the module %hs. This is secure, but may be incompatible with previous releases of the operating system. An alternative, %hs, is available. Should the application use the secure module %hs?
	ERROR_DLL_MIGHT_BE_INCOMPATIBLE = 687

	// ERROR_DBG_EXCEPTION_NOT_HANDLED
	// Debugger did not handle the exception.
	ERROR_DBG_EXCEPTION_NOT_HANDLED = 688

	// ERROR_DBG_REPLY_LATER
	// Debugger will reply later.
	ERROR_DBG_REPLY_LATER = 689

	// ERROR_DBG_UNABLE_TO_PROVIDE_HANDLE
	// Debugger cannot provide handle.
	ERROR_DBG_UNABLE_TO_PROVIDE_HANDLE = 690

	// ERROR_DBG_TERMINATE_THREAD
	// Debugger terminated thread.
	ERROR_DBG_TERMINATE_THREAD = 691

	// ERROR_DBG_TERMINATE_PROCESS
	// Debugger terminated process.
	ERROR_DBG_TERMINATE_PROCESS = 692

	// ERROR_DBG_CONTROL_C
	// Debugger got control C.
	ERROR_DBG_CONTROL_C = 693

	// ERROR_DBG_PRINTEXCEPTION_C
	// Debugger printed exception on control C.
	ERROR_DBG_PRINTEXCEPTION_C = 694

	// ERROR_DBG_RIPEXCEPTION
	// Debugger received RIP exception.
	ERROR_DBG_RIPEXCEPTION = 695

	// ERROR_DBG_CONTROL_BREAK
	// Debugger received control break.
	ERROR_DBG_CONTROL_BREAK = 696

	// ERROR_DBG_COMMAND_EXCEPTION
	// Debugger command communication exception.
	ERROR_DBG_COMMAND_EXCEPTION = 697

	// ERROR_OBJECT_NAME_EXISTS
	// {Object Exists}
	// An attempt was made to create an object and the object name already existed.
	ERROR_OBJECT_NAME_EXISTS = 698

	// ERROR_THREAD_WAS_SUSPENDED
	// {Thread Suspended}
	// A thread termination occurred while the thread was suspended. The thread was resumed, and termination proceeded.
	ERROR_THREAD_WAS_SUSPENDED = 699

	// ERROR_IMAGE_NOT_AT_BASE
	// {Image Relocated}
	// An image file could not be mapped at the address specified in the image file. Local fixups must be performed on this image.
	ERROR_IMAGE_NOT_AT_BASE = 700

	// ERROR_RXACT_STATE_CREATED
	// This informational level status indicates that a specified registry sub-tree transaction state did not yet exist and had to be created.
	ERROR_RXACT_STATE_CREATED = 701

	// ERROR_SEGMENT_NOTIFICATION
	// {Segment Load}
	// A virtual DOS machine (VDM) is loading, unloading, or moving an MS-DOS or Win16 program segment image.
	// An exception is raised so a debugger can load, unload or track symbols and breakpoints within these=16-bit segments.
	ERROR_SEGMENT_NOTIFICATION = 702

	// ERROR_BAD_CURRENT_DIRECTORY
	// {Invalid Current Directory}
	// The process cannot switch to the startup current directory %hs.
	// Select OK to set current directory to %hs, or select CANCEL to exit.
	ERROR_BAD_CURRENT_DIRECTORY = 703

	// ERROR_FT_READ_RECOVERY_FROM_BACKUP
	// {Redundant Read}
	// To satisfy a read request, the NT fault-tolerant file system successfully read the requested data from a redundant copy.
	// This was done because the file system encountered a failure on a member of the fault-tolerant volume, but was unable to reassign the failing area of the device.
	ERROR_FT_READ_RECOVERY_FROM_BACKUP = 704

	// ERROR_FT_WRITE_RECOVERY
	// {Redundant Write}
	// To satisfy a write request, the NT fault-tolerant file system successfully wrote a redundant copy of the information.
	// This was done because the file system encountered a failure on a member of the fault-tolerant volume, but was not able to reassign the failing area of the device.
	ERROR_FT_WRITE_RECOVERY = 705

	// ERROR_IMAGE_MACHINE_TYPE_MISMATCH
	// {Machine Type Mismatch}
	// The image file %hs is valid, but is for a machine type other than the current machine. Select OK to continue, or CANCEL to fail the DLL load.
	ERROR_IMAGE_MACHINE_TYPE_MISMATCH = 706

	// ERROR_RECEIVE_PARTIA
	// {Partial Data Received}
	// The network transport returned partial data to its client. The remaining data will be sent later.
	ERROR_RECEIVE_PARTIAL = 707

	// ERROR_RECEIVE_EXPEDITED
	// {Expedited Data Received}
	// The network transport returned data to its client that was marked as expedited by the remote system.
	ERROR_RECEIVE_EXPEDITED = 708

	// ERROR_RECEIVE_PARTIAL_EXPEDITED
	// {Partial Expedited Data Received}
	// The network transport returned partial data to its client and this data was marked as expedited by the remote system. The remaining data will be sent later.
	ERROR_RECEIVE_PARTIAL_EXPEDITED = 709

	// ERROR_EVENT_DONE
	// {TDI Event Done}
	// The TDI indication has completed successfully.
	ERROR_EVENT_DONE = 710

	// ERROR_EVENT_PENDING
	// {TDI Event Pending}
	// The TDI indication has entered the pending state.
	ERROR_EVENT_PENDING = 711

	// ERROR_CHECKING_FILE_SYSTEM
	// Checking file system on %wZ
	ERROR_CHECKING_FILE_SYSTEM = 712

	// ERROR_FATAL_APP_EXIT
	// {Fatal Application Exit}
	// %hs
	ERROR_FATAL_APP_EXIT = 713

	// ERROR_PREDEFINED_HANDLE
	// The specified registry key is referenced by a predefined handle.
	ERROR_PREDEFINED_HANDLE = 714

	// ERROR_WAS_UNLOCKED
	// {Page Unlocked}
	// The page protection of a locked page was changed to 'No Access' and the page was unlocked from memory and from the process.
	ERROR_WAS_UNLOCKED = 715

	// ERROR_SERVICE_NOTIFICATION
	// %hs
	ERROR_SERVICE_NOTIFICATION = 716

	// ERROR_WAS_LOCKED
	// {Page Locked}
	// One of the pages to lock was already locked.
	ERROR_WAS_LOCKED = 717

	// ERROR_LOG_HARD_ERROR
	// Application popup: %1 : %2
	ERROR_LOG_HARD_ERROR = 718

	// ERROR_ALREADY_WIN32
	//  ERROR_ALREADY_WIN32
	ERROR_ALREADY_WIN32 = 719

	// ERROR_IMAGE_MACHINE_TYPE_MISMATCH_EXE
	// {Machine Type Mismatch}
	// The image file %hs is valid, but is for a machine type other than the current machine.
	ERROR_IMAGE_MACHINE_TYPE_MISMATCH_EXE = 720

	// ERROR_NO_YIELD_PERFORMED
	// A yield execution was performed and no thread was available to run.
	ERROR_NO_YIELD_PERFORMED = 721

	// ERROR_TIMER_RESUME_IGNORED
	// The resumable flag to a timer API was ignored.
	ERROR_TIMER_RESUME_IGNORED = 722

	// ERROR_ARBITRATION_UNHANDLED
	// The arbiter has deferred arbitration of these resources to its parent
	ERROR_ARBITRATION_UNHANDLED = 723

	// ERROR_CARDBUS_NOT_SUPPORTED
	// The inserted CardBus device cannot be started because of a configuration error on "%hs".
	ERROR_CARDBUS_NOT_SUPPORTED = 724

	// ERROR_MP_PROCESSOR_MISMATCH
	// The CPUs in this multiprocessor system are not all the same revision level. To use all processors the operating system restricts itself to the features of the least capable processor in the system. Should problems occur with this system, contact the CPU manufacturer to see if this mix of processors is supported.
	ERROR_MP_PROCESSOR_MISMATCH = 725

	// ERROR_HIBERNATED
	// The system was put into hibernation.
	ERROR_HIBERNATED = 726

	// ERROR_RESUME_HIBERNATION
	// The system was resumed from hibernation.
	ERROR_RESUME_HIBERNATION = 727

	// ERROR_FIRMWARE_UPDATED
	// Windows has detected that the system firmware (BIOS) was updated [previous firmware date = %2, current firmware date %3].
	ERROR_FIRMWARE_UPDATED = 728

	// ERROR_DRIVERS_LEAKING_LOCKED_PAGES
	// A device driver is leaking locked I/O pages causing system degradation. The system has automatically enabled tracking code in order to try and catch the culprit.
	ERROR_DRIVERS_LEAKING_LOCKED_PAGES = 729

	// ERROR_WAKE_SYSTEM
	// The system has awoken
	ERROR_WAKE_SYSTEM = 730

	// ERROR_WAIT_1
	//  ERROR_WAIT_1
	ERROR_WAIT_1 = 731

	// ERROR_WAIT_2
	//  ERROR_WAIT_2
	ERROR_WAIT_2 = 732

	// ERROR_WAIT_3
	//  ERROR_WAIT_3
	ERROR_WAIT_3 = 733

	// ERROR_WAIT_63
	//  ERROR_WAIT_63
	ERROR_WAIT_63 = 734

	// ERROR_ABANDONED_WAIT_0
	//  ERROR_ABANDONED_WAIT_0
	ERROR_ABANDONED_WAIT_0 = 735

	// ERROR_ABANDONED_WAIT_63
	//  ERROR_ABANDONED_WAIT_63
	ERROR_ABANDONED_WAIT_63 = 736

	// ERROR_USER_APC
	//  ERROR_USER_APC
	ERROR_USER_APC = 737

	// ERROR_KERNEL_APC
	//  ERROR_KERNEL_APC
	ERROR_KERNEL_APC = 738

	// ERROR_ALERTED
	//  ERROR_ALERTED
	ERROR_ALERTED = 739

	// ERROR_ELEVATION_REQUIRED
	// The requested operation requires elevation.
	ERROR_ELEVATION_REQUIRED = 740

	// ERROR_REPARSE
	// A reparse should be performed by the Object Manager since the name of the file resulted in a symbolic link.
	ERROR_REPARSE = 741

	// ERROR_OPLOCK_BREAK_IN_PROGRESS
	// An open/create operation completed while an oplock break is underway.
	ERROR_OPLOCK_BREAK_IN_PROGRESS = 742

	// ERROR_VOLUME_MOUNTED
	// A new volume has been mounted by a file system.
	ERROR_VOLUME_MOUNTED = 743

	// ERROR_RXACT_COMMITTED
	// This success level status indicates that the transaction state already exists for the registry sub-tree, but that a transaction commit was previously aborted.
	// The commit has now been completed.
	ERROR_RXACT_COMMITTED = 744

	// ERROR_NOTIFY_CLEANUP
	// This indicates that a notify change request has been completed due to closing the handle which made the notify change request.
	ERROR_NOTIFY_CLEANUP = 745

	// ERROR_PRIMARY_TRANSPORT_CONNECT_FAILED
	// {Connect Failure on Primary Transport}
	// An attempt was made to connect to the remote server %hs on the primary transport, but the connection failed.
	// The computer WAS able to connect on a secondary transport.
	ERROR_PRIMARY_TRANSPORT_CONNECT_FAILED = 746

	// ERROR_PAGE_FAULT_TRANSITION
	// Page fault was a transition fault.
	ERROR_PAGE_FAULT_TRANSITION = 747

	// ERROR_PAGE_FAULT_DEMAND_ZERO
	// Page fault was a demand zero fault.
	ERROR_PAGE_FAULT_DEMAND_ZERO = 748

	// ERROR_PAGE_FAULT_COPY_ON_WRITE
	// Page fault was a demand zero fault.
	ERROR_PAGE_FAULT_COPY_ON_WRITE = 749

	// ERROR_PAGE_FAULT_GUARD_PAGE
	// Page fault was a demand zero fault.
	ERROR_PAGE_FAULT_GUARD_PAGE = 750

	// ERROR_PAGE_FAULT_PAGING_FILE
	// Page fault was satisfied by reading from a secondary storage device.
	ERROR_PAGE_FAULT_PAGING_FILE = 751

	// ERROR_CACHE_PAGE_LOCKED
	// Cached page was locked during operation.
	ERROR_CACHE_PAGE_LOCKED = 752

	// ERROR_CRASH_DUMP
	// Crash dump exists in paging file.
	ERROR_CRASH_DUMP = 753

	// ERROR_BUFFER_ALL_ZEROS
	// Specified buffer contains all zeros.
	ERROR_BUFFER_ALL_ZEROS = 754

	// ERROR_REPARSE_OBJECT
	// A reparse should be performed by the Object Manager since the name of the file resulted in a symbolic link.
	ERROR_REPARSE_OBJECT = 755

	// ERROR_RESOURCE_REQUIREMENTS_CHANGED
	// The device has succeeded a query-stop and its resource requirements have changed.
	ERROR_RESOURCE_REQUIREMENTS_CHANGED = 756

	// ERROR_TRANSLATION_COMPLETE
	// The translator has translated these resources into the global space and no further translations should be performed.
	ERROR_TRANSLATION_COMPLETE = 757

	// ERROR_NOTHING_TO_TERMINATE
	// A process being terminated has no threads to terminate.
	ERROR_NOTHING_TO_TERMINATE = 758

	// ERROR_PROCESS_NOT_IN_JOB
	// The specified process is not part of a job.
	ERROR_PROCESS_NOT_IN_JOB = 759

	// ERROR_PROCESS_IN_JOB
	// The specified process is part of a job.
	ERROR_PROCESS_IN_JOB = 760

	// ERROR_VOLSNAP_HIBERNATE_READY
	// {Volume Shadow Copy Service}
	// The system is now ready for hibernation.
	ERROR_VOLSNAP_HIBERNATE_READY = 761

	// ERROR_FSFILTER_OP_COMPLETED_SUCCESSFULLY
	// A file system or file system filter driver has successfully completed an FsFilter operation.
	ERROR_FSFILTER_OP_COMPLETED_SUCCESSFULLY = 762

	// ERROR_INTERRUPT_VECTOR_ALREADY_CONNECTED
	// The specified interrupt vector was already connected.
	ERROR_INTERRUPT_VECTOR_ALREADY_CONNECTED = 763

	// ERROR_INTERRUPT_STILL_CONNECTED
	// The specified interrupt vector is still connected.
	ERROR_INTERRUPT_STILL_CONNECTED = 764

	// ERROR_WAIT_FOR_OPLOCK
	// An operation is blocked waiting for an oplock.
	ERROR_WAIT_FOR_OPLOCK = 765

	// ERROR_DBG_EXCEPTION_HANDLED
	// Debugger handled exception
	ERROR_DBG_EXCEPTION_HANDLED = 766

	// ERROR_DBG_CONTINUE
	// Debugger continued
	ERROR_DBG_CONTINUE = 767

	// ERROR_CALLBACK_POP_STACK
	// An exception occurred in a user mode callback and the kernel callback frame should be removed.
	ERROR_CALLBACK_POP_STACK = 768

	// ERROR_COMPRESSION_DISABLED
	// Compression is disabled for this volume.
	ERROR_COMPRESSION_DISABLED = 769

	// ERROR_CANTFETCHBACKWARDS
	// The data provider cannot fetch backwards through a result set.
	ERROR_CANTFETCHBACKWARDS = 770

	// ERROR_CANTSCROLLBACKWARDS
	// The data provider cannot scroll backwards through a result set.
	ERROR_CANTSCROLLBACKWARDS = 771

	// ERROR_ROWSNOTRELEASED
	// The data provider requires that previously fetched data is released before asking for more data.
	ERROR_ROWSNOTRELEASED = 772

	// ERROR_BAD_ACCESSOR_FLAGS
	// The data provider was not able to interpret the flags set for a column binding in an accessor.
	ERROR_BAD_ACCESSOR_FLAGS = 773

	// ERROR_ERRORS_ENCOUNTERED
	// One or more errors occurred while processing the request.
	ERROR_ERRORS_ENCOUNTERED = 774

	// ERROR_NOT_CAPABLE
	// The implementation is not capable of performing the request.
	ERROR_NOT_CAPABLE = 775

	// ERROR_REQUEST_OUT_OF_SEQUENCE
	// The client of a component requested an operation which is not valid given the state of the component instance.
	ERROR_REQUEST_OUT_OF_SEQUENCE = 776

	// ERROR_VERSION_PARSE_ERROR
	// A version number could not be parsed.
	ERROR_VERSION_PARSE_ERROR = 777

	// ERROR_BADSTARTPOSITION
	// The iterator's start position is invalid.
	ERROR_BADSTARTPOSITION = 778

	// ERROR_MEMORY_HARDWARE
	// The hardware has reported an uncorrectable memory error.
	ERROR_MEMORY_HARDWARE = 779

	// ERROR_DISK_REPAIR_DISABLED
	// The attempted operation required self healing to be enabled.
	ERROR_DISK_REPAIR_DISABLED = 780

	// ERROR_INSUFFICIENT_RESOURCE_FOR_SPECIFIED_SHARED_SECTION_SIZE
	// The Desktop heap encountered an error while allocating session memory. There is more information in the system event log.
	ERROR_INSUFFICIENT_RESOURCE_FOR_SPECIFIED_SHARED_SECTION_SIZE = 781

	// ERROR_SYSTEM_POWERSTATE_TRANSITION
	// The system power state is transitioning from %2 to %3.
	ERROR_SYSTEM_POWERSTATE_TRANSITION = 782

	// ERROR_SYSTEM_POWERSTATE_COMPLEX_TRANSITION
	// The system power state is transitioning from %2 to %3 but could enter %4.
	ERROR_SYSTEM_POWERSTATE_COMPLEX_TRANSITION = 783

	// ERROR_MCA_EXCEPTION
	// A thread is getting dispatched with MCA EXCEPTION because of MCA.
	ERROR_MCA_EXCEPTION = 784

	// ERROR_ACCESS_AUDIT_BY_POLICY
	// Access to %1 is monitored by policy rule %2.
	ERROR_ACCESS_AUDIT_BY_POLICY = 785

	// ERROR_ACCESS_DISABLED_NO_SAFER_UI_BY_POLICY
	// Access to %1 has been restricted by your Administrator by policy rule %2.
	ERROR_ACCESS_DISABLED_NO_SAFER_UI_BY_POLICY = 786

	// ERROR_ABANDON_HIBERFILE
	// A valid hibernation file has been invalidated and should be abandoned.
	ERROR_ABANDON_HIBERFILE = 787

	// ERROR_LOST_WRITEBEHIND_DATA_NETWORK_DISCONNECTED
	// {Delayed Write Failed}
	// Windows was unable to save all the data for the file %hs; the data has been lost.
	// This error may be caused by network connectivity issues. Please try to save this file elsewhere.
	ERROR_LOST_WRITEBEHIND_DATA_NETWORK_DISCONNECTED = 788

	// ERROR_LOST_WRITEBEHIND_DATA_NETWORK_SERVER_ERROR
	// {Delayed Write Failed}
	// Windows was unable to save all the data for the file %hs; the data has been lost.
	// This error was returned by the server on which the file exists. Please try to save this file elsewhere.
	ERROR_LOST_WRITEBEHIND_DATA_NETWORK_SERVER_ERROR = 789

	// ERROR_LOST_WRITEBEHIND_DATA_LOCAL_DISK_ERROR
	// {Delayed Write Failed}
	// Windows was unable to save all the data for the file %hs; the data has been lost.
	// This error may be caused if the device has been removed or the media is write-protected.
	ERROR_LOST_WRITEBEHIND_DATA_LOCAL_DISK_ERROR = 790

	// ERROR_BAD_MCFG_TABLE
	// The resources required for this device conflict with the MCFG table.
	ERROR_BAD_MCFG_TABLE = 791

	// ERROR_DISK_REPAIR_REDIRECTED
	// The volume repair could not be performed while it is online.
	// Please schedule to take the volume offline so that it can be repaired.
	ERROR_DISK_REPAIR_REDIRECTED = 792

	// ERROR_DISK_REPAIR_UNSUCCESSFU
	// The volume repair was not successful.
	ERROR_DISK_REPAIR_UNSUCCESSFUL = 793

	// ERROR_CORRUPT_LOG_OVERFUL
	// One of the volume corruption logs is full. Further corruptions that may be detected won't be logged.
	ERROR_CORRUPT_LOG_OVERFULL = 794

	// ERROR_CORRUPT_LOG_CORRUPTED
	// One of the volume corruption logs is internally corrupted and needs to be recreated. The volume may contain undetected corruptions and must be scanned.
	ERROR_CORRUPT_LOG_CORRUPTED = 795

	// ERROR_CORRUPT_LOG_UNAVAILABLE
	// One of the volume corruption logs is unavailable for being operated on.
	ERROR_CORRUPT_LOG_UNAVAILABLE = 796

	// ERROR_CORRUPT_LOG_DELETED_FUL
	// One of the volume corruption logs was deleted while still having corruption records in them. The volume contains detected corruptions and must be scanned.
	ERROR_CORRUPT_LOG_DELETED_FULL = 797

	// ERROR_CORRUPT_LOG_CLEARED
	// One of the volume corruption logs was cleared by chkdsk and no longer contains real corruptions.
	ERROR_CORRUPT_LOG_CLEARED = 798

	// ERROR_ORPHAN_NAME_EXHAUSTED
	// Orphaned files exist on the volume but could not be recovered because no more new names could be created in the recovery directory. Files must be moved from the recovery directory.
	ERROR_ORPHAN_NAME_EXHAUSTED = 799

	// ERROR_OPLOCK_SWITCHED_TO_NEW_HANDLE
	// The oplock that was associated with this handle is now associated with a different handle.
	ERROR_OPLOCK_SWITCHED_TO_NEW_HANDLE = 800

	// ERROR_CANNOT_GRANT_REQUESTED_OPLOCK
	// An oplock of the requested level cannot be granted.  An oplock of a lower level may be available.
	ERROR_CANNOT_GRANT_REQUESTED_OPLOCK = 801

	// ERROR_CANNOT_BREAK_OPLOCK
	// The operation did not complete successfully because it would cause an oplock to be broken. The caller has requested that existing oplocks not be broken.
	ERROR_CANNOT_BREAK_OPLOCK = 802

	// ERROR_OPLOCK_HANDLE_CLOSED
	// The handle with which this oplock was associated has been closed.  The oplock is now broken.
	ERROR_OPLOCK_HANDLE_CLOSED = 803

	// ERROR_NO_ACE_CONDITION
	// The specified access control entry (ACE) does not contain a condition.
	ERROR_NO_ACE_CONDITION = 804

	// ERROR_INVALID_ACE_CONDITION
	// The specified access control entry (ACE) contains an invalid condition.
	ERROR_INVALID_ACE_CONDITION = 805

	// ERROR_FILE_HANDLE_REVOKED
	// Access to the specified file handle has been revoked.
	ERROR_FILE_HANDLE_REVOKED = 806

	// ERROR_IMAGE_AT_DIFFERENT_BASE
	// {Image Relocated}
	// An image file was mapped at a different address from the one specified in the image file but fixups will still be automatically performed on the image.
	ERROR_IMAGE_AT_DIFFERENT_BASE = 807

	// ERROR_ENCRYPTED_IO_NOT_POSSIBLE
	// The read or write operation to an encrypted file could not be completed because the file has not been opened for data access.
	ERROR_ENCRYPTED_IO_NOT_POSSIBLE = 808

	// ERROR_FILE_METADATA_OPTIMIZATION_IN_PROGRESS
	// File metadata optimization is already in progress.
	ERROR_FILE_METADATA_OPTIMIZATION_IN_PROGRESS = 809

	// ERROR_QUOTA_ACTIVITY
	// The requested operation failed due to quota operation is still in progress.
	ERROR_QUOTA_ACTIVITY = 810

	// ERROR_HANDLE_REVOKED
	// Access to the specified handle has been revoked.
	ERROR_HANDLE_REVOKED = 811

	// ERROR_CALLBACK_INVOKE_INLINE
	// The callback function must be invoked inline.
	ERROR_CALLBACK_INVOKE_INLINE = 812

	// ERROR_CPU_SET_INVALID
	// The specified CPU Set IDs are invalid.
	ERROR_CPU_SET_INVALID = 813

	// ERROR_ENCLAVE_NOT_TERMINATED
	// The specified enclave has not yet been terminated.
	ERROR_ENCLAVE_NOT_TERMINATED = 814

	// ERROR_ENCLAVE_VIOLATION
	// An attempt was made to access protected memory in violation of its secure access policy.
	ERROR_ENCLAVE_VIOLATION = 815

	// **** Available SYSTEM error codes ****

	// ERROR_EA_ACCESS_DENIED
	// Access to the extended attribute was denied.
	ERROR_EA_ACCESS_DENIED = 994

	// ERROR_OPERATION_ABORTED
	// The I/O operation has been aborted because of either a thread exit or an application request.
	ERROR_OPERATION_ABORTED = 995

	// ERROR_IO_INCOMPLETE
	// Overlapped I/O event is not in a signaled state.
	ERROR_IO_INCOMPLETE = 996

	// ERROR_IO_PENDING
	// Overlapped I/O operation is in progress.
	ERROR_IO_PENDING = 997

	// ERROR_NOACCESS
	// Invalid access to memory location.
	ERROR_NOACCESS = 998

	// ERROR_SWAPERROR
	// Error performing inpage operation.
	ERROR_SWAPERROR = 999

	// ERROR_STACK_OVERFLOW
	// Recursion too deep; the stack overflowed.
	ERROR_STACK_OVERFLOW = 1001

	// ERROR_INVALID_MESSAGE
	// The window cannot act on the sent message.
	ERROR_INVALID_MESSAGE = 1002

	// ERROR_CAN_NOT_COMPLETE
	// Cannot complete this function.
	ERROR_CAN_NOT_COMPLETE = 1003

	// ERROR_INVALID_FLAGS
	// Invalid flags.
	ERROR_INVALID_FLAGS = 1004

	// ERROR_UNRECOGNIZED_VOLUME
	// The volume does not contain a recognized file system.
	// Please make sure that all required file system drivers are loaded and that the volume is not corrupted.
	ERROR_UNRECOGNIZED_VOLUME = 1005

	// ERROR_FILE_INVALID
	// The volume for a file has been externally altered so that the opened file is no longer valid.
	ERROR_FILE_INVALID = 1006

	// ERROR_FULLSCREEN_MODE
	// The requested operation cannot be performed in full-screen mode.
	ERROR_FULLSCREEN_MODE = 1007

	// ERROR_NO_TOKEN
	// An attempt was made to reference a token that does not exist.
	ERROR_NO_TOKEN = 1008

	// ERROR_BADDB
	// The configuration registry database is corrupt.
	ERROR_BADDB = 1009

	// ERROR_BADKEY
	// The configuration registry key is invalid.
	ERROR_BADKEY = 1010

	// ERROR_CANTOPEN
	// The configuration registry key could not be opened.
	ERROR_CANTOPEN = 1011

	// ERROR_CANTREAD
	// The configuration registry key could not be read.
	ERROR_CANTREAD = 1012

	// ERROR_CANTWRITE
	// The configuration registry key could not be written.
	ERROR_CANTWRITE = 1013

	// ERROR_REGISTRY_RECOVERED
	// One of the files in the registry database had to be recovered by use of a log or alternate copy. The recovery was successful.
	ERROR_REGISTRY_RECOVERED = 1014

	// ERROR_REGISTRY_CORRUPT
	// The registry is corrupted. The structure of one of the files containing registry data is corrupted, or the system's memory image of the file is corrupted, or the file could not be recovered because the alternate copy or log was absent or corrupted.
	ERROR_REGISTRY_CORRUPT = 1015

	// ERROR_REGISTRY_IO_FAILED
	// An I/O operation initiated by the registry failed unrecoverably. The registry could not read in, or write out, or flush, one of the files that contain the system's image of the registry.
	ERROR_REGISTRY_IO_FAILED = 1016

	// ERROR_NOT_REGISTRY_FILE
	// The system has attempted to load or restore a file into the registry, but the specified file is not in a registry file format.
	ERROR_NOT_REGISTRY_FILE = 1017

	// ERROR_KEY_DELETED
	// Illegal operation attempted on a registry key that has been marked for deletion.
	ERROR_KEY_DELETED = 1018

	// ERROR_NO_LOG_SPACE
	// System could not allocate the required space in a registry log.
	ERROR_NO_LOG_SPACE = 1019

	// ERROR_KEY_HAS_CHILDREN
	// Cannot create a symbolic link in a registry key that already has subkeys or values.
	ERROR_KEY_HAS_CHILDREN = 1020

	// ERROR_CHILD_MUST_BE_VOLATILE
	// Cannot create a stable subkey under a volatile parent key.
	ERROR_CHILD_MUST_BE_VOLATILE = 1021

	// ERROR_NOTIFY_ENUM_DIR
	// A notify change request is being completed and the information is not being returned in the caller's buffer. The caller now needs to enumerate the files to find the changes.
	ERROR_NOTIFY_ENUM_DIR = 1022

	// ERROR_DEPENDENT_SERVICES_RUNNING
	// A stop control has been sent to a service that other running services are dependent on.
	ERROR_DEPENDENT_SERVICES_RUNNING = 1051

	// ERROR_INVALID_SERVICE_CONTRO
	// The requested control is not valid for this service.
	ERROR_INVALID_SERVICE_CONTROL = 1052

	// ERROR_SERVICE_REQUEST_TIMEOUT
	// The service did not respond to the start or control request in a timely fashion.
	ERROR_SERVICE_REQUEST_TIMEOUT = 1053

	// ERROR_SERVICE_NO_THREAD
	// A thread could not be created for the service.
	ERROR_SERVICE_NO_THREAD = 1054

	// ERROR_SERVICE_DATABASE_LOCKED
	// The service database is locked.
	ERROR_SERVICE_DATABASE_LOCKED = 1055

	// ERROR_SERVICE_ALREADY_RUNNING
	// An instance of the service is already running.
	ERROR_SERVICE_ALREADY_RUNNING = 1056

	// ERROR_INVALID_SERVICE_ACCOUNT
	// The account name is invalid or does not exist, or the password is invalid for the account name specified.
	ERROR_INVALID_SERVICE_ACCOUNT = 1057

	// ERROR_SERVICE_DISABLED
	// The service cannot be started, either because it is disabled or because it has no enabled devices associated with it.
	ERROR_SERVICE_DISABLED = 1058

	// ERROR_CIRCULAR_DEPENDENCY
	// Circular service dependency was specified.
	ERROR_CIRCULAR_DEPENDENCY = 1059

	// ERROR_SERVICE_DOES_NOT_EXIST
	// The specified service does not exist as an installed service.
	ERROR_SERVICE_DOES_NOT_EXIST = 1060

	// ERROR_SERVICE_CANNOT_ACCEPT_CTR
	// The service cannot accept control messages at this time.
	ERROR_SERVICE_CANNOT_ACCEPT_CTRL = 1061

	// ERROR_SERVICE_NOT_ACTIVE
	// The service has not been started.
	ERROR_SERVICE_NOT_ACTIVE = 1062

	// ERROR_FAILED_SERVICE_CONTROLLER_CONNECT
	// The service process could not connect to the service controller.
	ERROR_FAILED_SERVICE_CONTROLLER_CONNECT = 1063

	// ERROR_EXCEPTION_IN_SERVICE
	// An exception occurred in the service when handling the control request.
	ERROR_EXCEPTION_IN_SERVICE = 1064

	// ERROR_DATABASE_DOES_NOT_EXIST
	// The database specified does not exist.
	ERROR_DATABASE_DOES_NOT_EXIST = 1065

	// ERROR_SERVICE_SPECIFIC_ERROR
	// The service has returned a service-specific error code.
	ERROR_SERVICE_SPECIFIC_ERROR = 1066

	// ERROR_PROCESS_ABORTED
	// The process terminated unexpectedly.
	ERROR_PROCESS_ABORTED = 1067

	// ERROR_SERVICE_DEPENDENCY_FAI
	// The dependency service or group failed to start.
	ERROR_SERVICE_DEPENDENCY_FAIL = 1068

	// ERROR_SERVICE_LOGON_FAILED
	// The service did not start due to a logon failure.
	ERROR_SERVICE_LOGON_FAILED = 1069

	// ERROR_SERVICE_START_HANG
	// After starting, the service hung in a start-pending state.
	ERROR_SERVICE_START_HANG = 1070

	// ERROR_INVALID_SERVICE_LOCK
	// The specified service database lock is invalid.
	ERROR_INVALID_SERVICE_LOCK = 1071

	// ERROR_SERVICE_MARKED_FOR_DELETE
	// The specified service has been marked for deletion.
	ERROR_SERVICE_MARKED_FOR_DELETE = 1072

	// ERROR_SERVICE_EXISTS
	// The specified service already exists.
	ERROR_SERVICE_EXISTS = 1073

	// ERROR_ALREADY_RUNNING_LKG
	// The system is currently running with the last-known-good configuration.
	ERROR_ALREADY_RUNNING_LKG = 1074

	// ERROR_SERVICE_DEPENDENCY_DELETED
	// The dependency service does not exist or has been marked for deletion.
	ERROR_SERVICE_DEPENDENCY_DELETED = 1075

	// ERROR_BOOT_ALREADY_ACCEPTED
	// The current boot has already been accepted for use as the last-known-good control set.
	ERROR_BOOT_ALREADY_ACCEPTED = 1076

	// ERROR_SERVICE_NEVER_STARTED
	// No attempts to start the service have been made since the last boot.
	ERROR_SERVICE_NEVER_STARTED = 1077

	// ERROR_DUPLICATE_SERVICE_NAME
	// The name is already in use as either a service name or a service display name.
	ERROR_DUPLICATE_SERVICE_NAME = 1078

	// ERROR_DIFFERENT_SERVICE_ACCOUNT
	// The account specified for this service is different from the account specified for other services running in the same process.
	ERROR_DIFFERENT_SERVICE_ACCOUNT = 1079

	// ERROR_CANNOT_DETECT_DRIVER_FAILURE
	// Failure actions can only be set for Win32 services, not for drivers.
	ERROR_CANNOT_DETECT_DRIVER_FAILURE = 1080

	// ERROR_CANNOT_DETECT_PROCESS_ABORT
	// This service runs in the same process as the service control manager.
	// Therefore, the service control manager cannot take action if this service's process terminates unexpectedly.
	ERROR_CANNOT_DETECT_PROCESS_ABORT = 1081

	// ERROR_NO_RECOVERY_PROGRAM
	// No recovery program has been configured for this service.
	ERROR_NO_RECOVERY_PROGRAM = 1082

	// ERROR_SERVICE_NOT_IN_EXE
	// The executable program that this service is configured to run in does not implement the service.
	ERROR_SERVICE_NOT_IN_EXE = 1083

	// ERROR_NOT_SAFEBOOT_SERVICE
	// This service cannot be started in Safe Mode
	ERROR_NOT_SAFEBOOT_SERVICE = 1084

	// ERROR_END_OF_MEDIA
	// The physical end of the tape has been reached.
	ERROR_END_OF_MEDIA = 1100

	// ERROR_FILEMARK_DETECTED
	// A tape access reached a filemark.
	ERROR_FILEMARK_DETECTED = 1101

	// ERROR_BEGINNING_OF_MEDIA
	// The beginning of the tape or a partition was encountered.
	ERROR_BEGINNING_OF_MEDIA = 1102

	// ERROR_SETMARK_DETECTED
	// A tape access reached the end of a set of files.
	ERROR_SETMARK_DETECTED = 1103

	// ERROR_NO_DATA_DETECTED
	// No more data is on the tape.
	ERROR_NO_DATA_DETECTED = 1104

	// ERROR_PARTITION_FAILURE
	// Tape could not be partitioned.
	ERROR_PARTITION_FAILURE = 1105

	// ERROR_INVALID_BLOCK_LENGTH
	// When accessing a new tape of a multivolume partition, the current block size is incorrect.
	ERROR_INVALID_BLOCK_LENGTH = 1106

	// ERROR_DEVICE_NOT_PARTITIONED
	// Tape partition information could not be found when loading a tape.
	ERROR_DEVICE_NOT_PARTITIONED = 1107

	// ERROR_UNABLE_TO_LOCK_MEDIA
	// Unable to lock the media eject mechanism.
	ERROR_UNABLE_TO_LOCK_MEDIA = 1108

	// ERROR_UNABLE_TO_UNLOAD_MEDIA
	// Unable to unload the media.
	ERROR_UNABLE_TO_UNLOAD_MEDIA = 1109

	// ERROR_MEDIA_CHANGED
	// The media in the drive may have changed.
	ERROR_MEDIA_CHANGED = 1110

	// ERROR_BUS_RESET
	// The I/O bus was reset.
	ERROR_BUS_RESET = 1111

	// ERROR_NO_MEDIA_IN_DRIVE
	// No media in drive.
	ERROR_NO_MEDIA_IN_DRIVE = 1112

	// ERROR_NO_UNICODE_TRANSLATION
	// No mapping for the Unicode character exists in the target multi-byte code page.
	ERROR_NO_UNICODE_TRANSLATION = 1113

	// ERROR_DLL_INIT_FAILED
	// A dynamic link library (DLL) initialization routine failed.
	ERROR_DLL_INIT_FAILED = 1114

	// ERROR_SHUTDOWN_IN_PROGRESS
	// A system shutdown is in progress.
	ERROR_SHUTDOWN_IN_PROGRESS = 1115

	// ERROR_NO_SHUTDOWN_IN_PROGRESS
	// Unable to abort the system shutdown because no shutdown was in progress.
	ERROR_NO_SHUTDOWN_IN_PROGRESS = 1116

	// ERROR_IO_DEVICE
	// The request could not be performed because of an I/O device error.
	ERROR_IO_DEVICE = 1117

	// ERROR_SERIAL_NO_DEVICE
	// No serial device was successfully initialized. The serial driver will unload.
	ERROR_SERIAL_NO_DEVICE = 1118

	// ERROR_IRQ_BUSY
	// Unable to open a device that was sharing an interrupt request (IRQ) with other devices. At least one other device that uses that IRQ was already opened.
	ERROR_IRQ_BUSY = 1119

	// ERROR_MORE_WRITES
	// A serial I/O operation was completed by another write to the serial port.
	// (The IOCTL_SERIAL_XOFF_COUNTER reached zero.)
	ERROR_MORE_WRITES = 1120

	// ERROR_COUNTER_TIMEOUT
	// A serial I/O operation completed because the timeout period expired.
	// (The IOCTL_SERIAL_XOFF_COUNTER did not reach zero.)
	ERROR_COUNTER_TIMEOUT = 1121

	// ERROR_FLOPPY_ID_MARK_NOT_FOUND
	// No ID address mark was found on the floppy disk.
	ERROR_FLOPPY_ID_MARK_NOT_FOUND = 1122

	// ERROR_FLOPPY_WRONG_CYLINDER
	// Mismatch between the floppy disk sector ID field and the floppy disk controller track address.
	ERROR_FLOPPY_WRONG_CYLINDER = 1123

	// ERROR_FLOPPY_UNKNOWN_ERROR
	// The floppy disk controller reported an error that is not recognized by the floppy disk driver.
	ERROR_FLOPPY_UNKNOWN_ERROR = 1124

	// ERROR_FLOPPY_BAD_REGISTERS
	// The floppy disk controller returned inconsistent results in its registers.
	ERROR_FLOPPY_BAD_REGISTERS = 1125

	// ERROR_DISK_RECALIBRATE_FAILED
	// While accessing the hard disk, a recalibrate operation failed, even after retries.
	ERROR_DISK_RECALIBRATE_FAILED = 1126

	// ERROR_DISK_OPERATION_FAILED
	// While accessing the hard disk, a disk operation failed even after retries.
	ERROR_DISK_OPERATION_FAILED = 1127

	// ERROR_DISK_RESET_FAILED
	// While accessing the hard disk, a disk controller reset was needed, but even that failed.
	ERROR_DISK_RESET_FAILED = 1128

	// ERROR_EOM_OVERFLOW
	// Physical end of tape encountered.
	ERROR_EOM_OVERFLOW = 1129

	// ERROR_NOT_ENOUGH_SERVER_MEMORY
	// Not enough server memory resources are available to process this command.
	ERROR_NOT_ENOUGH_SERVER_MEMORY = 1130

	// ERROR_POSSIBLE_DEADLOCK
	// A potential deadlock condition has been detected.
	ERROR_POSSIBLE_DEADLOCK = 1131

	// ERROR_MAPPED_ALIGNMENT
	// The base address or the file offset specified does not have the proper alignment.
	ERROR_MAPPED_ALIGNMENT = 1132

	// ERROR_SET_POWER_STATE_VETOED
	// An attempt to change the system power state was vetoed by another application or driver.
	ERROR_SET_POWER_STATE_VETOED = 1140

	// ERROR_SET_POWER_STATE_FAILED
	// The system BIOS failed an attempt to change the system power state.
	ERROR_SET_POWER_STATE_FAILED = 1141

	// ERROR_TOO_MANY_LINKS
	// An attempt was made to create more links on a file than the file system supports.
	ERROR_TOO_MANY_LINKS = 1142

	// ERROR_OLD_WIN_VERSION
	// The specified program requires a newer version of Windows.
	ERROR_OLD_WIN_VERSION = 1150

	// ERROR_APP_WRONG_OS
	// The specified program is not a Windows or MS-DOS program.
	ERROR_APP_WRONG_OS = 1151

	// ERROR_SINGLE_INSTANCE_APP
	// Cannot start more than one instance of the specified program.
	ERROR_SINGLE_INSTANCE_APP = 1152

	// ERROR_RMODE_APP
	// The specified program was written for an earlier version of Windows.
	ERROR_RMODE_APP = 1153

	// ERROR_INVALID_DL
	// One of the library files needed to run this application is damaged.
	ERROR_INVALID_DLL = 1154

	// ERROR_NO_ASSOCIATION
	// No application is associated with the specified file for this operation.
	ERROR_NO_ASSOCIATION = 1155

	// ERROR_DDE_FAI
	// An error occurred in sending the command to the application.
	ERROR_DDE_FAIL = 1156

	// ERROR_DLL_NOT_FOUND
	// One of the library files needed to run this application cannot be found.
	ERROR_DLL_NOT_FOUND = 1157

	// ERROR_NO_MORE_USER_HANDLES
	// The current process has used all of its system allowance of handles for Window Manager objects.
	ERROR_NO_MORE_USER_HANDLES = 1158

	// ERROR_MESSAGE_SYNC_ONLY
	// The message can be used only with synchronous operations.
	ERROR_MESSAGE_SYNC_ONLY = 1159

	// ERROR_SOURCE_ELEMENT_EMPTY
	// The indicated source element has no media.
	ERROR_SOURCE_ELEMENT_EMPTY = 1160

	// ERROR_DESTINATION_ELEMENT_FUL
	// The indicated destination element already contains media.
	ERROR_DESTINATION_ELEMENT_FULL = 1161

	// ERROR_ILLEGAL_ELEMENT_ADDRESS
	// The indicated element does not exist.
	ERROR_ILLEGAL_ELEMENT_ADDRESS = 1162

	// ERROR_MAGAZINE_NOT_PRESENT
	// The indicated element is part of a magazine that is not present.
	ERROR_MAGAZINE_NOT_PRESENT = 1163

	// ERROR_DEVICE_REINITIALIZATION_NEEDED
	// The indicated device requires reinitialization due to hardware errors.
	ERROR_DEVICE_REINITIALIZATION_NEEDED = 1164

	// ERROR_DEVICE_REQUIRES_CLEANING
	// The device has indicated that cleaning is required before further operations are attempted.
	ERROR_DEVICE_REQUIRES_CLEANING = 1165

	// ERROR_DEVICE_DOOR_OPEN
	// The device has indicated that its door is open.
	ERROR_DEVICE_DOOR_OPEN = 1166

	// ERROR_DEVICE_NOT_CONNECTED
	// The device is not connected.
	ERROR_DEVICE_NOT_CONNECTED = 1167

	// ERROR_NOT_FOUND
	// Element not found.
	ERROR_NOT_FOUND = 1168

	// ERROR_NO_MATCH
	// There was no match for the specified key in the index.
	ERROR_NO_MATCH = 1169

	// ERROR_SET_NOT_FOUND
	// The property set specified does not exist on the object.
	ERROR_SET_NOT_FOUND = 1170

	// ERROR_POINT_NOT_FOUND
	// The point passed to GetMouseMovePoints is not in the buffer.
	ERROR_POINT_NOT_FOUND = 1171

	// ERROR_NO_TRACKING_SERVICE
	// The tracking (workstation) service is not running.
	ERROR_NO_TRACKING_SERVICE = 1172

	// ERROR_NO_VOLUME_ID
	// The Volume ID could not be found.
	ERROR_NO_VOLUME_ID = 1173

	// ERROR_UNABLE_TO_REMOVE_REPLACED
	// Unable to remove the file to be replaced.
	ERROR_UNABLE_TO_REMOVE_REPLACED = 1175

	// ERROR_UNABLE_TO_MOVE_REPLACEMENT
	// Unable to move the replacement file to the file to be replaced. The file to be replaced has retained its original name.
	ERROR_UNABLE_TO_MOVE_REPLACEMENT = 1176

	// ERROR_UNABLE_TO_MOVE_REPLACEMENT_2
	// Unable to move the replacement file to the file to be replaced. The file to be replaced has been renamed using the backup name.
	ERROR_UNABLE_TO_MOVE_REPLACEMENT_2 = 1177

	// ERROR_JOURNAL_DELETE_IN_PROGRESS
	// The volume change journal is being deleted.
	ERROR_JOURNAL_DELETE_IN_PROGRESS = 1178

	// ERROR_JOURNAL_NOT_ACTIVE
	// The volume change journal is not active.
	ERROR_JOURNAL_NOT_ACTIVE = 1179

	// ERROR_POTENTIAL_FILE_FOUND
	// A file was found, but it may not be the correct file.
	ERROR_POTENTIAL_FILE_FOUND = 1180

	// ERROR_JOURNAL_ENTRY_DELETED
	// The journal entry has been deleted from the journal.
	ERROR_JOURNAL_ENTRY_DELETED = 1181

	// ERROR_SHUTDOWN_IS_SCHEDULED
	// A system shutdown has already been scheduled.
	ERROR_SHUTDOWN_IS_SCHEDULED = 1190

	// ERROR_SHUTDOWN_USERS_LOGGED_ON
	// The system shutdown cannot be initiated because there are other users logged on to the computer.
	ERROR_SHUTDOWN_USERS_LOGGED_ON = 1191

	// ERROR_BAD_DEVICE
	// The specified device name is invalid.
	ERROR_BAD_DEVICE = 1200

	// ERROR_CONNECTION_UNAVAI
	// The device is not currently connected but it is a remembered connection.
	ERROR_CONNECTION_UNAVAIL = 1201

	// ERROR_DEVICE_ALREADY_REMEMBERED
	// The local device name has a remembered connection to another network resource.
	ERROR_DEVICE_ALREADY_REMEMBERED = 1202

	// ERROR_NO_NET_OR_BAD_PATH
	// The network path was either typed incorrectly, does not exist, or the network provider is not currently available. Please try retyping the path or contact your network administrator.
	ERROR_NO_NET_OR_BAD_PATH = 1203

	// ERROR_BAD_PROVIDER
	// The specified network provider name is invalid.
	ERROR_BAD_PROVIDER = 1204

	// ERROR_CANNOT_OPEN_PROFILE
	// Unable to open the network connection profile.
	ERROR_CANNOT_OPEN_PROFILE = 1205

	// ERROR_BAD_PROFILE
	// The network connection profile is corrupted.
	ERROR_BAD_PROFILE = 1206

	// ERROR_NOT_CONTAINER
	// Cannot enumerate a noncontainer.
	ERROR_NOT_CONTAINER = 1207

	// ERROR_EXTENDED_ERROR
	// An extended error has occurred.
	ERROR_EXTENDED_ERROR = 1208

	// ERROR_INVALID_GROUPNAME
	// The format of the specified group name is invalid.
	ERROR_INVALID_GROUPNAME = 1209

	// ERROR_INVALID_COMPUTERNAME
	// The format of the specified computer name is invalid.
	ERROR_INVALID_COMPUTERNAME = 1210

	// ERROR_INVALID_EVENTNAME
	// The format of the specified event name is invalid.
	ERROR_INVALID_EVENTNAME = 1211

	// ERROR_INVALID_DOMAINNAME
	// The format of the specified domain name is invalid.
	ERROR_INVALID_DOMAINNAME = 1212

	// ERROR_INVALID_SERVICENAME
	// The format of the specified service name is invalid.
	ERROR_INVALID_SERVICENAME = 1213

	// ERROR_INVALID_NETNAME
	// The format of the specified network name is invalid.
	ERROR_INVALID_NETNAME = 1214

	// ERROR_INVALID_SHARENAME
	// The format of the specified share name is invalid.
	ERROR_INVALID_SHARENAME = 1215

	// ERROR_INVALID_PASSWORDNAME
	// The format of the specified password is invalid.
	ERROR_INVALID_PASSWORDNAME = 1216

	// ERROR_INVALID_MESSAGENAME
	// The format of the specified message name is invalid.
	ERROR_INVALID_MESSAGENAME = 1217

	// ERROR_INVALID_MESSAGEDEST
	// The format of the specified message destination is invalid.
	ERROR_INVALID_MESSAGEDEST = 1218

	// ERROR_SESSION_CREDENTIAL_CONFLICT
	// Multiple connections to a server or shared resource by the same user, using more than one user name, are not allowed. Disconnect all previous connections to the server or shared resource and try again.
	ERROR_SESSION_CREDENTIAL_CONFLICT = 1219

	// ERROR_REMOTE_SESSION_LIMIT_EXCEEDED
	// An attempt was made to establish a session to a network server, but there are already too many sessions established to that server.
	ERROR_REMOTE_SESSION_LIMIT_EXCEEDED = 1220

	// ERROR_DUP_DOMAINNAME
	// The workgroup or domain name is already in use by another computer on the network.
	ERROR_DUP_DOMAINNAME = 1221

	// ERROR_NO_NETWORK
	// The network is not present or not started.
	ERROR_NO_NETWORK = 1222

	// ERROR_CANCELLED
	// The operation was canceled by the user.
	ERROR_CANCELLED = 1223

	// ERROR_USER_MAPPED_FILE
	// The requested operation cannot be performed on a file with a user-mapped section open.
	ERROR_USER_MAPPED_FILE = 1224

	// ERROR_CONNECTION_REFUSED
	// The remote computer refused the network connection.
	ERROR_CONNECTION_REFUSED = 1225

	// ERROR_GRACEFUL_DISCONNECT
	// The network connection was gracefully closed.
	ERROR_GRACEFUL_DISCONNECT = 1226

	// ERROR_ADDRESS_ALREADY_ASSOCIATED
	// The network transport endpoint already has an address associated with it.
	ERROR_ADDRESS_ALREADY_ASSOCIATED = 1227

	// ERROR_ADDRESS_NOT_ASSOCIATED
	// An address has not yet been associated with the network endpoint.
	ERROR_ADDRESS_NOT_ASSOCIATED = 1228

	// ERROR_CONNECTION_INVALID
	// An operation was attempted on a nonexistent network connection.
	ERROR_CONNECTION_INVALID = 1229

	// ERROR_CONNECTION_ACTIVE
	// An invalid operation was attempted on an active network connection.
	ERROR_CONNECTION_ACTIVE = 1230

	// ERROR_NETWORK_UNREACHABLE
	// The network location cannot be reached. For information about network troubleshooting, see Windows Help.
	ERROR_NETWORK_UNREACHABLE = 1231

	// ERROR_HOST_UNREACHABLE
	// The network location cannot be reached. For information about network troubleshooting, see Windows Help.
	ERROR_HOST_UNREACHABLE = 1232

	// ERROR_PROTOCOL_UNREACHABLE
	// The network location cannot be reached. For information about network troubleshooting, see Windows Help.
	ERROR_PROTOCOL_UNREACHABLE = 1233

	// ERROR_PORT_UNREACHABLE
	// No service is operating at the destination network endpoint on the remote system.
	ERROR_PORT_UNREACHABLE = 1234

	// ERROR_REQUEST_ABORTED
	// The request was aborted.
	ERROR_REQUEST_ABORTED = 1235

	// ERROR_CONNECTION_ABORTED
	// The network connection was aborted by the local system.
	ERROR_CONNECTION_ABORTED = 1236

	// ERROR_RETRY
	// The operation could not be completed. A retry should be performed.
	ERROR_RETRY = 1237

	// ERROR_CONNECTION_COUNT_LIMIT
	// A connection to the server could not be made because the limit on the number of concurrent connections for this account has been reached.
	ERROR_CONNECTION_COUNT_LIMIT = 1238

	// ERROR_LOGIN_TIME_RESTRICTION
	// Attempting to log in during an unauthorized time of day for this account.
	ERROR_LOGIN_TIME_RESTRICTION = 1239

	// ERROR_LOGIN_WKSTA_RESTRICTION
	// The account is not authorized to log in from this station.
	ERROR_LOGIN_WKSTA_RESTRICTION = 1240

	// ERROR_INCORRECT_ADDRESS
	// The network address could not be used for the operation requested.
	ERROR_INCORRECT_ADDRESS = 1241

	// ERROR_ALREADY_REGISTERED
	// The service is already registered.
	ERROR_ALREADY_REGISTERED = 1242

	// ERROR_SERVICE_NOT_FOUND
	// The specified service does not exist.
	ERROR_SERVICE_NOT_FOUND = 1243

	// ERROR_NOT_AUTHENTICATED
	// The operation being requested was not performed because the user has not been authenticated.
	ERROR_NOT_AUTHENTICATED = 1244

	// ERROR_NOT_LOGGED_ON
	// The operation being requested was not performed because the user has not logged on to the network. The specified service does not exist.
	ERROR_NOT_LOGGED_ON = 1245

	// ERROR_CONTINUE
	// Continue with work in progress.
	ERROR_CONTINUE = 1246

	// ERROR_ALREADY_INITIALIZED
	// An attempt was made to perform an initialization operation when initialization has already been completed.
	ERROR_ALREADY_INITIALIZED = 1247

	// ERROR_NO_MORE_DEVICES
	// No more local devices.
	ERROR_NO_MORE_DEVICES = 1248

	// ERROR_NO_SUCH_SITE
	// The specified site does not exist.
	ERROR_NO_SUCH_SITE = 1249

	// ERROR_DOMAIN_CONTROLLER_EXISTS
	// A domain controller with the specified name already exists.
	ERROR_DOMAIN_CONTROLLER_EXISTS = 1250

	// ERROR_ONLY_IF_CONNECTED
	// This operation is supported only when you are connected to the server.
	ERROR_ONLY_IF_CONNECTED = 1251

	// ERROR_OVERRIDE_NOCHANGES
	// The group policy framework should call the extension even if there are no changes.
	ERROR_OVERRIDE_NOCHANGES = 1252

	// ERROR_BAD_USER_PROFILE
	// The specified user does not have a valid profile.
	ERROR_BAD_USER_PROFILE = 1253

	// ERROR_NOT_SUPPORTED_ON_SBS
	// This operation is not supported on a computer running Windows Server=2003 for Small Business Server
	ERROR_NOT_SUPPORTED_ON_SBS = 1254

	// ERROR_SERVER_SHUTDOWN_IN_PROGRESS
	// The server machine is shutting down.
	ERROR_SERVER_SHUTDOWN_IN_PROGRESS = 1255

	// ERROR_HOST_DOWN
	// The remote system is not available. For information about network troubleshooting, see Windows Help.
	ERROR_HOST_DOWN = 1256

	// ERROR_NON_ACCOUNT_SID
	// The security identifier provided is not from an account domain.
	ERROR_NON_ACCOUNT_SID = 1257

	// ERROR_NON_DOMAIN_SID
	// The security identifier provided does not have a domain component.
	ERROR_NON_DOMAIN_SID = 1258

	// ERROR_APPHELP_BLOCK
	// AppHelp dialog canceled thus preventing the application from starting.
	ERROR_APPHELP_BLOCK = 1259

	// ERROR_ACCESS_DISABLED_BY_POLICY
	// This program is blocked by group policy. For more information, contact your system administrator.
	ERROR_ACCESS_DISABLED_BY_POLICY = 1260

	// ERROR_REG_NAT_CONSUMPTION
	// A program attempt to use an invalid register value. Normally caused by an uninitialized register. This error is Itanium specific.
	ERROR_REG_NAT_CONSUMPTION = 1261

	// ERROR_CSCSHARE_OFFLINE
	// The share is currently offline or does not exist.
	ERROR_CSCSHARE_OFFLINE = 1262

	// ERROR_PKINIT_FAILURE
	// The Kerberos protocol encountered an error while validating the KDC certificate during smartcard logon. There is more information in the system event log.
	ERROR_PKINIT_FAILURE = 1263

	// ERROR_SMARTCARD_SUBSYSTEM_FAILURE
	// The Kerberos protocol encountered an error while attempting to utilize the smartcard subsystem.
	ERROR_SMARTCARD_SUBSYSTEM_FAILURE = 1264

	// ERROR_DOWNGRADE_DETECTED
	// The system cannot contact a domain controller to service the authentication request. Please try again later.
	ERROR_DOWNGRADE_DETECTED = 1265

	// Do not use ID's=1266 -=1270 as the symbolicNames have been moved to SEC_E_*

	// ERROR_MACHINE_LOCKED
	// The machine is locked and cannot be shut down without the force option.
	ERROR_MACHINE_LOCKED = 1271

	// ERROR_SMB_GUEST_LOGON_BLOCKED
	// You can't access this shared folder because your organization's security policies block unauthenticated guest access. These policies help protect your PC from unsafe or malicious devices on the network.
	ERROR_SMB_GUEST_LOGON_BLOCKED = 1272

	// ERROR_CALLBACK_SUPPLIED_INVALID_DATA
	// An application-defined callback gave invalid data when called.
	ERROR_CALLBACK_SUPPLIED_INVALID_DATA = 1273

	// ERROR_SYNC_FOREGROUND_REFRESH_REQUIRED
	// The group policy framework should call the extension in the synchronous foreground policy refresh.
	ERROR_SYNC_FOREGROUND_REFRESH_REQUIRED = 1274

	// ERROR_DRIVER_BLOCKED
	// This driver has been blocked from loading
	ERROR_DRIVER_BLOCKED = 1275

	// ERROR_INVALID_IMPORT_OF_NON_DL
	// A dynamic link library (DLL) referenced a module that was neither a DLL nor the process's executable image.
	ERROR_INVALID_IMPORT_OF_NON_DLL = 1276

	// ERROR_ACCESS_DISABLED_WEBBLADE
	// Windows cannot open this program since it has been disabled.
	ERROR_ACCESS_DISABLED_WEBBLADE = 1277

	// ERROR_ACCESS_DISABLED_WEBBLADE_TAMPER
	// Windows cannot open this program because the license enforcement system has been tampered with or become corrupted.
	ERROR_ACCESS_DISABLED_WEBBLADE_TAMPER = 1278

	// ERROR_RECOVERY_FAILURE
	// A transaction recover failed.
	ERROR_RECOVERY_FAILURE = 1279

	// ERROR_ALREADY_FIBER
	// The current thread has already been converted to a fiber.
	ERROR_ALREADY_FIBER = 1280

	// ERROR_ALREADY_THREAD
	// The current thread has already been converted from a fiber.
	ERROR_ALREADY_THREAD = 1281

	// ERROR_STACK_BUFFER_OVERRUN
	// The system detected an overrun of a stack-based buffer in this application. This overrun could potentially allow a malicious user to gain control of this application.
	ERROR_STACK_BUFFER_OVERRUN = 1282

	// ERROR_PARAMETER_QUOTA_EXCEEDED
	// Data present in one of the parameters is more than the function can operate on.
	ERROR_PARAMETER_QUOTA_EXCEEDED = 1283

	// ERROR_DEBUGGER_INACTIVE
	// An attempt to do an operation on a debug object failed because the object is in the process of being deleted.
	ERROR_DEBUGGER_INACTIVE = 1284

	// ERROR_DELAY_LOAD_FAILED
	// An attempt to delay-load a .dll or get a function address in a delay-loaded .dll failed.
	ERROR_DELAY_LOAD_FAILED = 1285

	// ERROR_VDM_DISALLOWED
	// %1 is a=16-bit application. You do not have permissions to execute=16-bit applications. Check your permissions with your system administrator.
	ERROR_VDM_DISALLOWED = 1286

	// ERROR_UNIDENTIFIED_ERROR
	// Insufficient information exists to identify the cause of failure.
	ERROR_UNIDENTIFIED_ERROR = 1287

	// ERROR_INVALID_CRUNTIME_PARAMETER
	// The parameter passed to a C runtime function is incorrect.
	ERROR_INVALID_CRUNTIME_PARAMETER = 1288

	// ERROR_BEYOND_VD
	// The operation occurred beyond the valid data length of the file.
	ERROR_BEYOND_VDL = 1289

	// ERROR_INCOMPATIBLE_SERVICE_SID_TYPE
	// The service start failed since one or more services in the same process have an incompatible service SID type setting. A service with restricted service SID type can only coexist in the same process with other services with a restricted SID type. If the service SID type for this service was just configured, the hosting process must be restarted in order to start this service.
	ERROR_INCOMPATIBLE_SERVICE_SID_TYPE = 1290

	// ERROR_DRIVER_PROCESS_TERMINATED
	// The process hosting the driver for this device has been terminated.
	ERROR_DRIVER_PROCESS_TERMINATED = 1291

	// ERROR_IMPLEMENTATION_LIMIT
	// An operation attempted to exceed an implementation-defined limit.
	ERROR_IMPLEMENTATION_LIMIT = 1292

	// ERROR_PROCESS_IS_PROTECTED
	// Either the target process, or the target thread's containing process, is a protected process.
	ERROR_PROCESS_IS_PROTECTED = 1293

	// ERROR_SERVICE_NOTIFY_CLIENT_LAGGING
	// The service notification client is lagging too far behind the current state of services in the machine.
	ERROR_SERVICE_NOTIFY_CLIENT_LAGGING = 1294

	// ERROR_DISK_QUOTA_EXCEEDED
	// The requested file operation failed because the storage quota was exceeded.
	// To free up disk space, move files to a different location or delete unnecessary files. For more information, contact your system administrator.
	ERROR_DISK_QUOTA_EXCEEDED = 1295

	// ERROR_CONTENT_BLOCKED
	// The requested file operation failed because the storage policy blocks that type of file. For more information, contact your system administrator.
	ERROR_CONTENT_BLOCKED = 1296

	// ERROR_INCOMPATIBLE_SERVICE_PRIVILEGE
	// A privilege that the service requires to function properly does not exist in the service account configuration.
	// You may use the Services Microsoft Management Console (MMC) snap-in (services.msc) and the Local Security Settings MMC snap-in (secpol.msc) to view the service configuration and the account configuration.
	ERROR_INCOMPATIBLE_SERVICE_PRIVILEGE = 1297

	// ERROR_APP_HANG
	// A thread involved in this operation appears to be unresponsive.
	ERROR_APP_HANG = 1298

	///////////////////////////////////////////////////
	//                                               //
	//             SECURITY Error codes              //
	//                                               //
	//                =1299 to=1399                  //
	///////////////////////////////////////////////////

	// ERROR_INVALID_LABE
	// Indicates a particular Security ID may not be assigned as the label of an object.
	ERROR_INVALID_LABEL = 1299

	// ERROR_NOT_ALL_ASSIGNED
	// Not all privileges or groups referenced are assigned to the caller.
	ERROR_NOT_ALL_ASSIGNED = 1300

	// ERROR_SOME_NOT_MAPPED
	// Some mapping between account names and security IDs was not done.
	ERROR_SOME_NOT_MAPPED = 1301

	// ERROR_NO_QUOTAS_FOR_ACCOUNT
	// No system quota limits are specifically set for this account.
	ERROR_NO_QUOTAS_FOR_ACCOUNT = 1302

	// ERROR_LOCAL_USER_SESSION_KEY
	// No encryption key is available. A well-known encryption key was returned.
	ERROR_LOCAL_USER_SESSION_KEY = 1303

	// ERROR_NULL_LM_PASSWORD
	// The password is too complex to be converted to a LAN Manager password. The LAN Manager password returned is a NULL string.
	ERROR_NULL_LM_PASSWORD = 1304

	// ERROR_UNKNOWN_REVISION
	// The revision level is unknown.
	ERROR_UNKNOWN_REVISION = 1305

	// ERROR_REVISION_MISMATCH
	// Indicates two revision levels are incompatible.
	ERROR_REVISION_MISMATCH = 1306

	// ERROR_INVALID_OWNER
	// This security ID may not be assigned as the owner of this object.
	ERROR_INVALID_OWNER = 1307

	// ERROR_INVALID_PRIMARY_GROUP
	// This security ID may not be assigned as the primary group of an object.
	ERROR_INVALID_PRIMARY_GROUP = 1308

	// ERROR_NO_IMPERSONATION_TOKEN
	// An attempt has been made to operate on an impersonation token by a thread that is not currently impersonating a client.
	ERROR_NO_IMPERSONATION_TOKEN = 1309

	// ERROR_CANT_DISABLE_MANDATORY
	// The group may not be disabled.
	ERROR_CANT_DISABLE_MANDATORY = 1310

	// ERROR_NO_LOGON_SERVERS
	// We can't sign you in with this credential because your domain isn't available. Make sure your device is connected to your organization's network and try again. If you previously signed in on this device with another credential, you can sign in with that credential.
	ERROR_NO_LOGON_SERVERS = 1311

	// ERROR_NO_SUCH_LOGON_SESSION
	// A specified logon session does not exist. It may already have been terminated.
	ERROR_NO_SUCH_LOGON_SESSION = 1312

	// ERROR_NO_SUCH_PRIVILEGE
	// A specified privilege does not exist.
	ERROR_NO_SUCH_PRIVILEGE = 1313

	// ERROR_PRIVILEGE_NOT_HELD
	// A required privilege is not held by the client.
	ERROR_PRIVILEGE_NOT_HELD = 1314

	// ERROR_INVALID_ACCOUNT_NAME
	// The name provided is not a properly formed account name.
	ERROR_INVALID_ACCOUNT_NAME = 1315

	// ERROR_USER_EXISTS
	// The specified account already exists.
	ERROR_USER_EXISTS = 1316

	// ERROR_NO_SUCH_USER
	// The specified account does not exist.
	ERROR_NO_SUCH_USER = 1317

	// ERROR_GROUP_EXISTS
	// The specified group already exists.
	ERROR_GROUP_EXISTS = 1318

	// ERROR_NO_SUCH_GROUP
	// The specified group does not exist.
	ERROR_NO_SUCH_GROUP = 1319

	// ERROR_MEMBER_IN_GROUP
	// Either the specified user account is already a member of the specified group, or the specified group cannot be deleted because it contains a member.
	ERROR_MEMBER_IN_GROUP = 1320

	// ERROR_MEMBER_NOT_IN_GROUP
	// The specified user account is not a member of the specified group account.
	ERROR_MEMBER_NOT_IN_GROUP = 1321

	// ERROR_LAST_ADMIN
	// This operation is disallowed as it could result in an administration account being disabled, deleted or unable to logon.
	ERROR_LAST_ADMIN = 1322

	// ERROR_WRONG_PASSWORD
	// Unable to update the password. The value provided as the current password is incorrect.
	ERROR_WRONG_PASSWORD = 1323

	// ERROR_ILL_FORMED_PASSWORD
	// Unable to update the password. The value provided for the new password contains values that are not allowed in passwords.
	ERROR_ILL_FORMED_PASSWORD = 1324

	// ERROR_PASSWORD_RESTRICTION
	// Unable to update the password. The value provided for the new password does not meet the length, complexity, or history requirements of the domain.
	ERROR_PASSWORD_RESTRICTION = 1325

	// ERROR_LOGON_FAILURE
	// The user name or password is incorrect.
	ERROR_LOGON_FAILURE = 1326

	// ERROR_ACCOUNT_RESTRICTION
	// Account restrictions are preventing this user from signing in. For example: blank passwords aren't allowed, sign-in times are limited, or a policy restriction has been enforced.
	ERROR_ACCOUNT_RESTRICTION = 1327

	// ERROR_INVALID_LOGON_HOURS
	// Your account has time restrictions that keep you from signing in right now.
	ERROR_INVALID_LOGON_HOURS = 1328

	// ERROR_INVALID_WORKSTATION
	// This user isn't allowed to sign in to this computer.
	ERROR_INVALID_WORKSTATION = 1329

	// ERROR_PASSWORD_EXPIRED
	// The password for this account has expired.
	ERROR_PASSWORD_EXPIRED = 1330

	// ERROR_ACCOUNT_DISABLED
	// This user can't sign in because this account is currently disabled.
	ERROR_ACCOUNT_DISABLED = 1331

	// ERROR_NONE_MAPPED
	// No mapping between account names and security IDs was done.
	ERROR_NONE_MAPPED = 1332

	// ERROR_TOO_MANY_LUIDS_REQUESTED
	// Too many local user identifiers (LUIDs) were requested at one time.
	ERROR_TOO_MANY_LUIDS_REQUESTED = 1333

	// ERROR_LUIDS_EXHAUSTED
	// No more local user identifiers (LUIDs) are available.
	ERROR_LUIDS_EXHAUSTED = 1334

	// ERROR_INVALID_SUB_AUTHORITY
	// The subauthority part of a security ID is invalid for this particular use.
	ERROR_INVALID_SUB_AUTHORITY = 1335

	// ERROR_INVALID_AC
	// The access control list (ACL) structure is invalid.
	ERROR_INVALID_ACL = 1336

	// ERROR_INVALID_SID
	// The security ID structure is invalid.
	ERROR_INVALID_SID = 1337

	// ERROR_INVALID_SECURITY_DESCR
	// The security descriptor structure is invalid.
	ERROR_INVALID_SECURITY_DESCR = 1338

	// ERROR_BAD_INHERITANCE_AC
	// The inherited access control list (ACL) or access control entry (ACE) could not be built.
	ERROR_BAD_INHERITANCE_ACL = 1340

	// ERROR_SERVER_DISABLED
	// The server is currently disabled.
	ERROR_SERVER_DISABLED = 1341

	// ERROR_SERVER_NOT_DISABLED
	// The server is currently enabled.
	ERROR_SERVER_NOT_DISABLED = 1342

	// ERROR_INVALID_ID_AUTHORITY
	// The value provided was an invalid value for an identifier authority.
	ERROR_INVALID_ID_AUTHORITY = 1343

	// ERROR_ALLOTTED_SPACE_EXCEEDED
	// No more memory is available for security information updates.
	ERROR_ALLOTTED_SPACE_EXCEEDED = 1344

	// ERROR_INVALID_GROUP_ATTRIBUTES
	// The specified attributes are invalid, or incompatible with the attributes for the group as a whole.
	ERROR_INVALID_GROUP_ATTRIBUTES = 1345

	// ERROR_BAD_IMPERSONATION_LEVE
	// Either a required impersonation level was not provided, or the provided impersonation level is invalid.
	ERROR_BAD_IMPERSONATION_LEVEL = 1346

	// ERROR_CANT_OPEN_ANONYMOUS
	// Cannot open an anonymous level security token.
	ERROR_CANT_OPEN_ANONYMOUS = 1347

	// ERROR_BAD_VALIDATION_CLASS
	// The validation information class requested was invalid.
	ERROR_BAD_VALIDATION_CLASS = 1348

	// ERROR_BAD_TOKEN_TYPE
	// The type of the token is inappropriate for its attempted use.
	ERROR_BAD_TOKEN_TYPE = 1349

	// ERROR_NO_SECURITY_ON_OBJECT
	// Unable to perform a security operation on an object that has no associated security.
	ERROR_NO_SECURITY_ON_OBJECT = 1350

	// ERROR_CANT_ACCESS_DOMAIN_INFO
	// Configuration information could not be read from the domain controller, either because the machine is unavailable, or access has been denied.
	ERROR_CANT_ACCESS_DOMAIN_INFO = 1351

	// ERROR_INVALID_SERVER_STATE
	// The security account manager (SAM) or local security authority (LSA) server was in the wrong state to perform the security operation.
	ERROR_INVALID_SERVER_STATE = 1352

	// ERROR_INVALID_DOMAIN_STATE
	// The domain was in the wrong state to perform the security operation.
	ERROR_INVALID_DOMAIN_STATE = 1353

	// ERROR_INVALID_DOMAIN_ROLE
	// This operation is only allowed for the Primary Domain Controller of the domain.
	ERROR_INVALID_DOMAIN_ROLE = 1354

	// ERROR_NO_SUCH_DOMAIN
	// The specified domain either does not exist or could not be contacted.
	ERROR_NO_SUCH_DOMAIN = 1355

	// ERROR_DOMAIN_EXISTS
	// The specified domain already exists.
	ERROR_DOMAIN_EXISTS = 1356

	// ERROR_DOMAIN_LIMIT_EXCEEDED
	// An attempt was made to exceed the limit on the number of domains per server.
	ERROR_DOMAIN_LIMIT_EXCEEDED = 1357

	// ERROR_INTERNAL_DB_CORRUPTION
	// Unable to complete the requested operation because of either a catastrophic media failure or a data structure corruption on the disk.
	ERROR_INTERNAL_DB_CORRUPTION = 1358

	// ERROR_INTERNAL_ERROR
	// An internal error occurred.
	ERROR_INTERNAL_ERROR = 1359

	// ERROR_GENERIC_NOT_MAPPED
	// Generic access types were contained in an access mask which should already be mapped to nongeneric types.
	ERROR_GENERIC_NOT_MAPPED = 1360

	// ERROR_BAD_DESCRIPTOR_FORMAT
	// A security descriptor is not in the right format (absolute or self-relative).
	ERROR_BAD_DESCRIPTOR_FORMAT = 1361

	// ERROR_NOT_LOGON_PROCESS
	// The requested action is restricted for use by logon processes only. The calling process has not registered as a logon process.
	ERROR_NOT_LOGON_PROCESS = 1362

	// ERROR_LOGON_SESSION_EXISTS
	// Cannot start a new logon session with an ID that is already in use.
	ERROR_LOGON_SESSION_EXISTS = 1363

	// ERROR_NO_SUCH_PACKAGE
	// A specified authentication package is unknown.
	ERROR_NO_SUCH_PACKAGE = 1364

	// ERROR_BAD_LOGON_SESSION_STATE
	// The logon session is not in a state that is consistent with the requested operation.
	ERROR_BAD_LOGON_SESSION_STATE = 1365

	// ERROR_LOGON_SESSION_COLLISION
	// The logon session ID is already in use.
	ERROR_LOGON_SESSION_COLLISION = 1366

	// ERROR_INVALID_LOGON_TYPE
	// A logon request contained an invalid logon type value.
	ERROR_INVALID_LOGON_TYPE = 1367

	// ERROR_CANNOT_IMPERSONATE
	// Unable to impersonate using a named pipe until data has been read from that pipe.
	ERROR_CANNOT_IMPERSONATE = 1368

	// ERROR_RXACT_INVALID_STATE
	// The transaction state of a registry subtree is incompatible with the requested operation.
	ERROR_RXACT_INVALID_STATE = 1369

	// ERROR_RXACT_COMMIT_FAILURE
	// An internal security database corruption has been encountered.
	ERROR_RXACT_COMMIT_FAILURE = 1370

	// ERROR_SPECIAL_ACCOUNT
	// Cannot perform this operation on built-in accounts.
	ERROR_SPECIAL_ACCOUNT = 1371

	// ERROR_SPECIAL_GROUP
	// Cannot perform this operation on this built-in special group.
	ERROR_SPECIAL_GROUP = 1372

	// ERROR_SPECIAL_USER
	// Cannot perform this operation on this built-in special user.
	ERROR_SPECIAL_USER = 1373

	// ERROR_MEMBERS_PRIMARY_GROUP
	// The user cannot be removed from a group because the group is currently the user's primary group.
	ERROR_MEMBERS_PRIMARY_GROUP = 1374

	// ERROR_TOKEN_ALREADY_IN_USE
	// The token is already in use as a primary token.
	ERROR_TOKEN_ALREADY_IN_USE = 1375

	// ERROR_NO_SUCH_ALIAS
	// The specified local group does not exist.
	ERROR_NO_SUCH_ALIAS = 1376

	// ERROR_MEMBER_NOT_IN_ALIAS
	// The specified account name is not a member of the group.
	ERROR_MEMBER_NOT_IN_ALIAS = 1377

	// ERROR_MEMBER_IN_ALIAS
	// The specified account name is already a member of the group.
	ERROR_MEMBER_IN_ALIAS = 1378

	// ERROR_ALIAS_EXISTS
	// The specified local group already exists.
	ERROR_ALIAS_EXISTS = 1379

	// ERROR_LOGON_NOT_GRANTED
	// Logon failure: the user has not been granted the requested logon type at this computer.
	ERROR_LOGON_NOT_GRANTED = 1380

	// ERROR_TOO_MANY_SECRETS
	// The maximum number of secrets that may be stored in a single system has been exceeded.
	ERROR_TOO_MANY_SECRETS = 1381

	// ERROR_SECRET_TOO_LONG
	// The length of a secret exceeds the maximum length allowed.
	ERROR_SECRET_TOO_LONG = 1382

	// ERROR_INTERNAL_DB_ERROR
	// The local security authority database contains an internal inconsistency.
	ERROR_INTERNAL_DB_ERROR = 1383

	// ERROR_TOO_MANY_CONTEXT_IDS
	// During a logon attempt, the user's security context accumulated too many security IDs.
	ERROR_TOO_MANY_CONTEXT_IDS = 1384

	// ERROR_LOGON_TYPE_NOT_GRANTED
	// Logon failure: the user has not been granted the requested logon type at this computer.
	ERROR_LOGON_TYPE_NOT_GRANTED = 1385

	// ERROR_NT_CROSS_ENCRYPTION_REQUIRED
	// A cross-encrypted password is necessary to change a user password.
	ERROR_NT_CROSS_ENCRYPTION_REQUIRED = 1386

	// ERROR_NO_SUCH_MEMBER
	// A member could not be added to or removed from the local group because the member does not exist.
	ERROR_NO_SUCH_MEMBER = 1387

	// ERROR_INVALID_MEMBER
	// A new member could not be added to a local group because the member has the wrong account type.
	ERROR_INVALID_MEMBER = 1388

	// ERROR_TOO_MANY_SIDS
	// Too many security IDs have been specified.
	ERROR_TOO_MANY_SIDS = 1389

	// ERROR_LM_CROSS_ENCRYPTION_REQUIRED
	// A cross-encrypted password is necessary to change this user password.
	ERROR_LM_CROSS_ENCRYPTION_REQUIRED = 1390

	// ERROR_NO_INHERITANCE
	// Indicates an ACL contains no inheritable components.
	ERROR_NO_INHERITANCE = 1391

	// ERROR_FILE_CORRUPT
	// The file or directory is corrupted and unreadable.
	ERROR_FILE_CORRUPT = 1392

	// ERROR_DISK_CORRUPT
	// The disk structure is corrupted and unreadable.
	ERROR_DISK_CORRUPT = 1393

	// ERROR_NO_USER_SESSION_KEY
	// There is no user session key for the specified logon session.
	ERROR_NO_USER_SESSION_KEY = 1394

	// ERROR_LICENSE_QUOTA_EXCEEDED
	// The service being accessed is licensed for a particular number of connections. No more connections can be made to the service at this time because there are already as many connections as the service can accept.
	ERROR_LICENSE_QUOTA_EXCEEDED = 1395

	// ERROR_WRONG_TARGET_NAME
	// The target account name is incorrect.
	ERROR_WRONG_TARGET_NAME = 1396

	// ERROR_MUTUAL_AUTH_FAILED
	// Mutual Authentication failed. The server's password is out of date at the domain controller.
	ERROR_MUTUAL_AUTH_FAILED = 1397

	// ERROR_TIME_SKEW
	// There is a time and/or date difference between the client and server.
	ERROR_TIME_SKEW = 1398

	// ERROR_CURRENT_DOMAIN_NOT_ALLOWED
	// This operation cannot be performed on the current domain.
	ERROR_CURRENT_DOMAIN_NOT_ALLOWED = 1399

	///////////////////////////////////////////////////
	//                                               //
	//              WinUser Error codes              //
	//                                               //
	//                =1400 to=1499                  //
	///////////////////////////////////////////////////

	// ERROR_INVALID_WINDOW_HANDLE
	// Invalid window handle.
	ERROR_INVALID_WINDOW_HANDLE = 1400

	// ERROR_INVALID_MENU_HANDLE
	// Invalid menu handle.
	ERROR_INVALID_MENU_HANDLE = 1401

	// ERROR_INVALID_CURSOR_HANDLE
	// Invalid cursor handle.
	ERROR_INVALID_CURSOR_HANDLE = 1402

	// ERROR_INVALID_ACCEL_HANDLE
	// Invalid accelerator table handle.
	ERROR_INVALID_ACCEL_HANDLE = 1403

	// ERROR_INVALID_HOOK_HANDLE
	// Invalid hook handle.
	ERROR_INVALID_HOOK_HANDLE = 1404

	// ERROR_INVALID_DWP_HANDLE
	// Invalid handle to a multiple-window position structure.
	ERROR_INVALID_DWP_HANDLE = 1405

	// ERROR_TLW_WITH_WSCHILD
	// Cannot create a top-level child window.
	ERROR_TLW_WITH_WSCHILD = 1406

	// ERROR_CANNOT_FIND_WND_CLASS
	// Cannot find window class.
	ERROR_CANNOT_FIND_WND_CLASS = 1407

	// ERROR_WINDOW_OF_OTHER_THREAD
	// Invalid window; it belongs to other thread.
	ERROR_WINDOW_OF_OTHER_THREAD = 1408

	// ERROR_HOTKEY_ALREADY_REGISTERED
	// Hot key is already registered.
	ERROR_HOTKEY_ALREADY_REGISTERED = 1409

	// ERROR_CLASS_ALREADY_EXISTS
	// Class already exists.
	ERROR_CLASS_ALREADY_EXISTS = 1410

	// ERROR_CLASS_DOES_NOT_EXIST
	// Class does not exist.
	ERROR_CLASS_DOES_NOT_EXIST = 1411

	// ERROR_CLASS_HAS_WINDOWS
	// Class still has open windows.
	ERROR_CLASS_HAS_WINDOWS = 1412

	// ERROR_INVALID_INDEX
	// Invalid index.
	ERROR_INVALID_INDEX = 1413

	// ERROR_INVALID_ICON_HANDLE
	// Invalid icon handle.
	ERROR_INVALID_ICON_HANDLE = 1414

	// ERROR_PRIVATE_DIALOG_INDEX
	// Using private DIALOG window words.
	ERROR_PRIVATE_DIALOG_INDEX = 1415

	// ERROR_LISTBOX_ID_NOT_FOUND
	// The list box identifier was not found.
	ERROR_LISTBOX_ID_NOT_FOUND = 1416

	// ERROR_NO_WILDCARD_CHARACTERS
	// No wildcards were found.
	ERROR_NO_WILDCARD_CHARACTERS = 1417

	// ERROR_CLIPBOARD_NOT_OPEN
	// Thread does not have a clipboard open.
	ERROR_CLIPBOARD_NOT_OPEN = 1418

	// ERROR_HOTKEY_NOT_REGISTERED
	// Hot key is not registered.
	ERROR_HOTKEY_NOT_REGISTERED = 1419

	// ERROR_WINDOW_NOT_DIALOG
	// The window is not a valid dialog window.
	ERROR_WINDOW_NOT_DIALOG = 1420

	// ERROR_CONTROL_ID_NOT_FOUND
	// Control ID not found.
	ERROR_CONTROL_ID_NOT_FOUND = 1421

	// ERROR_INVALID_COMBOBOX_MESSAGE
	// Invalid message for a combo box because it does not have an edit control.
	ERROR_INVALID_COMBOBOX_MESSAGE = 1422

	// ERROR_WINDOW_NOT_COMBOBOX
	// The window is not a combo box.
	ERROR_WINDOW_NOT_COMBOBOX = 1423

	// ERROR_INVALID_EDIT_HEIGHT
	// Height must be less than=256.
	ERROR_INVALID_EDIT_HEIGHT = 1424

	// ERROR_DC_NOT_FOUND
	// Invalid device context (DC) handle.
	ERROR_DC_NOT_FOUND = 1425

	// ERROR_INVALID_HOOK_FILTER
	// Invalid hook procedure type.
	ERROR_INVALID_HOOK_FILTER = 1426

	// ERROR_INVALID_FILTER_PROC
	// Invalid hook procedure.
	ERROR_INVALID_FILTER_PROC = 1427

	// ERROR_HOOK_NEEDS_HMOD
	// Cannot set nonlocal hook without a module handle.
	ERROR_HOOK_NEEDS_HMOD = 1428

	// ERROR_GLOBAL_ONLY_HOOK
	// This hook procedure can only be set globally.
	ERROR_GLOBAL_ONLY_HOOK = 1429

	// ERROR_JOURNAL_HOOK_SET
	// The journal hook procedure is already installed.
	ERROR_JOURNAL_HOOK_SET = 1430

	// ERROR_HOOK_NOT_INSTALLED
	// The hook procedure is not installed.
	ERROR_HOOK_NOT_INSTALLED = 1431

	// ERROR_INVALID_LB_MESSAGE
	// Invalid message for single-selection list box.
	ERROR_INVALID_LB_MESSAGE = 1432

	// ERROR_SETCOUNT_ON_BAD_LB
	// LB_SETCOUNT sent to non-lazy list box.
	ERROR_SETCOUNT_ON_BAD_LB = 1433

	// ERROR_LB_WITHOUT_TABSTOPS
	// This list box does not support tab stops.
	ERROR_LB_WITHOUT_TABSTOPS = 1434

	// ERROR_DESTROY_OBJECT_OF_OTHER_THREAD
	// Cannot destroy object created by another thread.
	ERROR_DESTROY_OBJECT_OF_OTHER_THREAD = 1435

	// ERROR_CHILD_WINDOW_MENU
	// Child windows cannot have menus.
	ERROR_CHILD_WINDOW_MENU = 1436

	// ERROR_NO_SYSTEM_MENU
	// The window does not have a system menu.
	ERROR_NO_SYSTEM_MENU = 1437

	// ERROR_INVALID_MSGBOX_STYLE
	// Invalid message box style.
	ERROR_INVALID_MSGBOX_STYLE = 1438

	// ERROR_INVALID_SPI_VALUE
	// Invalid system-wide (SPI_*) parameter.
	ERROR_INVALID_SPI_VALUE = 1439

	// ERROR_SCREEN_ALREADY_LOCKED
	// Screen already locked.
	ERROR_SCREEN_ALREADY_LOCKED = 1440

	// ERROR_HWNDS_HAVE_DIFF_PARENT
	// All handles to windows in a multiple-window position structure must have the same parent.
	ERROR_HWNDS_HAVE_DIFF_PARENT = 1441

	// ERROR_NOT_CHILD_WINDOW
	// The window is not a child window.
	ERROR_NOT_CHILD_WINDOW = 1442

	// ERROR_INVALID_GW_COMMAND
	// Invalid GW_* command.
	ERROR_INVALID_GW_COMMAND = 1443

	// ERROR_INVALID_THREAD_ID
	// Invalid thread identifier.
	ERROR_INVALID_THREAD_ID = 1444

	// ERROR_NON_MDICHILD_WINDOW
	// Cannot process a message from a window that is not a multiple document interface (MDI) window.
	ERROR_NON_MDICHILD_WINDOW = 1445

	// ERROR_POPUP_ALREADY_ACTIVE
	// Popup menu already active.
	ERROR_POPUP_ALREADY_ACTIVE = 1446

	// ERROR_NO_SCROLLBARS
	// The window does not have scroll bars.
	ERROR_NO_SCROLLBARS = 1447

	// ERROR_INVALID_SCROLLBAR_RANGE
	// Scroll bar range cannot be greater than MAXLONG.
	ERROR_INVALID_SCROLLBAR_RANGE = 1448

	// ERROR_INVALID_SHOWWIN_COMMAND
	// Cannot show or remove the window in the way specified.
	ERROR_INVALID_SHOWWIN_COMMAND = 1449

	// ERROR_NO_SYSTEM_RESOURCES
	// Insufficient system resources exist to complete the requested service.
	ERROR_NO_SYSTEM_RESOURCES = 1450

	// ERROR_NONPAGED_SYSTEM_RESOURCES
	// Insufficient system resources exist to complete the requested service.
	ERROR_NONPAGED_SYSTEM_RESOURCES = 1451

	// ERROR_PAGED_SYSTEM_RESOURCES
	// Insufficient system resources exist to complete the requested service.
	ERROR_PAGED_SYSTEM_RESOURCES = 1452

	// ERROR_WORKING_SET_QUOTA
	// Insufficient quota to complete the requested service.
	ERROR_WORKING_SET_QUOTA = 1453

	// ERROR_PAGEFILE_QUOTA
	// Insufficient quota to complete the requested service.
	ERROR_PAGEFILE_QUOTA = 1454

	// ERROR_COMMITMENT_LIMIT
	// The paging file is too small for this operation to complete.
	ERROR_COMMITMENT_LIMIT = 1455

	// ERROR_MENU_ITEM_NOT_FOUND
	// A menu item was not found.
	ERROR_MENU_ITEM_NOT_FOUND = 1456

	// ERROR_INVALID_KEYBOARD_HANDLE
	// Invalid keyboard layout handle.
	ERROR_INVALID_KEYBOARD_HANDLE = 1457

	// ERROR_HOOK_TYPE_NOT_ALLOWED
	// Hook type not allowed.
	ERROR_HOOK_TYPE_NOT_ALLOWED = 1458

	// ERROR_REQUIRES_INTERACTIVE_WINDOWSTATION
	// This operation requires an interactive window station.
	ERROR_REQUIRES_INTERACTIVE_WINDOWSTATION = 1459

	// ERROR_TIMEOUT
	// This operation returned because the timeout period expired.
	ERROR_TIMEOUT = 1460

	// ERROR_INVALID_MONITOR_HANDLE
	// Invalid monitor handle.
	ERROR_INVALID_MONITOR_HANDLE = 1461

	// ERROR_INCORRECT_SIZE
	// Incorrect size argument.
	ERROR_INCORRECT_SIZE = 1462

	// ERROR_SYMLINK_CLASS_DISABLED
	// The symbolic link cannot be followed because its type is disabled.
	ERROR_SYMLINK_CLASS_DISABLED = 1463

	// ERROR_SYMLINK_NOT_SUPPORTED
	// This application does not support the current operation on symbolic links.
	ERROR_SYMLINK_NOT_SUPPORTED = 1464

	// ERROR_XML_PARSE_ERROR
	// Windows was unable to parse the requested XML data.
	ERROR_XML_PARSE_ERROR = 1465

	// ERROR_XMLDSIG_ERROR
	// An error was encountered while processing an XML digital signature.
	ERROR_XMLDSIG_ERROR = 1466

	// ERROR_RESTART_APPLICATION
	// This application must be restarted.
	ERROR_RESTART_APPLICATION = 1467

	// ERROR_WRONG_COMPARTMENT
	// The caller made the connection request in the wrong routing compartment.
	ERROR_WRONG_COMPARTMENT = 1468

	// ERROR_AUTHIP_FAILURE
	// There was an AuthIP failure when attempting to connect to the remote host.
	ERROR_AUTHIP_FAILURE = 1469

	// ERROR_NO_NVRAM_RESOURCES
	// Insufficient NVRAM resources exist to complete the requested service. A reboot might be required.
	ERROR_NO_NVRAM_RESOURCES = 1470

	// ERROR_NOT_GUI_PROCESS
	// Unable to finish the requested operation because the specified process is not a GUI process.
	ERROR_NOT_GUI_PROCESS = 1471

	///////////////////////////////////////////////////
	//                                               //
	//             EventLog Error codes              //
	//                                               //
	//                =1500 to=1549                  //
	///////////////////////////////////////////////////

	// ERROR_EVENTLOG_FILE_CORRUPT
	// The event log file is corrupted.
	ERROR_EVENTLOG_FILE_CORRUPT = 1500

	// ERROR_EVENTLOG_CANT_START
	// No event log file could be opened, so the event logging service did not start.
	ERROR_EVENTLOG_CANT_START = 1501

	// ERROR_LOG_FILE_FUL
	// The event log file is full.
	ERROR_LOG_FILE_FULL = 1502

	// ERROR_EVENTLOG_FILE_CHANGED
	// The event log file has changed between read operations.
	ERROR_EVENTLOG_FILE_CHANGED = 1503

	// ERROR_CONTAINER_ASSIGNED
	// The specified Job already has a container assigned to it.
	ERROR_CONTAINER_ASSIGNED = 1504

	// ERROR_JOB_NO_CONTAINER
	// The specified Job does not have a container assigned to it.
	ERROR_JOB_NO_CONTAINER = 1505

	///////////////////////////////////////////////////
	//                                               //
	//            Class Scheduler Error codes        //
	//                                               //
	//                =1550 to=1599                  //
	///////////////////////////////////////////////////

	// ERROR_INVALID_TASK_NAME
	// The specified task name is invalid.
	ERROR_INVALID_TASK_NAME = 1550

	// ERROR_INVALID_TASK_INDEX
	// The specified task index is invalid.
	ERROR_INVALID_TASK_INDEX = 1551

	// ERROR_THREAD_ALREADY_IN_TASK
	// The specified thread is already joining a task.
	ERROR_THREAD_ALREADY_IN_TASK = 1552

	///////////////////////////////////////////////////
	//                                               //
	//                MSI Error codes                //
	//                                               //
	//                =1600 to=1699                  //
	///////////////////////////////////////////////////

	// ERROR_INSTALL_SERVICE_FAILURE
	// The Windows Installer Service could not be accessed. This can occur if the Windows Installer is not correctly installed. Contact your support personnel for assistance.
	ERROR_INSTALL_SERVICE_FAILURE = 1601

	// ERROR_INSTALL_USEREXIT
	// User cancelled installation.
	ERROR_INSTALL_USEREXIT = 1602

	// ERROR_INSTALL_FAILURE
	// Fatal error during installation.
	ERROR_INSTALL_FAILURE = 1603

	// ERROR_INSTALL_SUSPEND
	// Installation suspended, incomplete.
	ERROR_INSTALL_SUSPEND = 1604

	// ERROR_UNKNOWN_PRODUCT
	// This action is only valid for products that are currently installed.
	ERROR_UNKNOWN_PRODUCT = 1605

	// ERROR_UNKNOWN_FEATURE
	// Feature ID not registered.
	ERROR_UNKNOWN_FEATURE = 1606

	// ERROR_UNKNOWN_COMPONENT
	// Component ID not registered.
	ERROR_UNKNOWN_COMPONENT = 1607

	// ERROR_UNKNOWN_PROPERTY
	// Unknown property.
	ERROR_UNKNOWN_PROPERTY = 1608

	// ERROR_INVALID_HANDLE_STATE
	// Handle is in an invalid state.
	ERROR_INVALID_HANDLE_STATE = 1609

	// ERROR_BAD_CONFIGURATION
	// The configuration data for this product is corrupt. Contact your support personnel.
	ERROR_BAD_CONFIGURATION = 1610

	// ERROR_INDEX_ABSENT
	// Component qualifier not present.
	ERROR_INDEX_ABSENT = 1611

	// ERROR_INSTALL_SOURCE_ABSENT
	// The installation source for this product is not available. Verify that the source exists and that you can access it.
	ERROR_INSTALL_SOURCE_ABSENT = 1612

	// ERROR_INSTALL_PACKAGE_VERSION
	// This installation package cannot be installed by the Windows Installer service. You must install a Windows service pack that contains a newer version of the Windows Installer service.
	ERROR_INSTALL_PACKAGE_VERSION = 1613

	// ERROR_PRODUCT_UNINSTALLED
	// Product is uninstalled.
	ERROR_PRODUCT_UNINSTALLED = 1614

	// ERROR_BAD_QUERY_SYNTAX
	// SQL query syntax invalid or unsupported.
	ERROR_BAD_QUERY_SYNTAX = 1615

	// ERROR_INVALID_FIELD
	// Record field does not exist.
	ERROR_INVALID_FIELD = 1616

	// ERROR_DEVICE_REMOVED
	// The device has been removed.
	ERROR_DEVICE_REMOVED = 1617

	// ERROR_INSTALL_ALREADY_RUNNING
	// Another installation is already in progress. Complete that installation before proceeding with this install.
	ERROR_INSTALL_ALREADY_RUNNING = 1618

	// ERROR_INSTALL_PACKAGE_OPEN_FAILED
	// This installation package could not be opened. Verify that the package exists and that you can access it, or contact the application vendor to verify that this is a valid Windows Installer package.
	ERROR_INSTALL_PACKAGE_OPEN_FAILED = 1619

	// ERROR_INSTALL_PACKAGE_INVALID
	// This installation package could not be opened. Contact the application vendor to verify that this is a valid Windows Installer package.
	ERROR_INSTALL_PACKAGE_INVALID = 1620

	// ERROR_INSTALL_UI_FAILURE
	// There was an error starting the Windows Installer service user interface. Contact your support personnel.
	ERROR_INSTALL_UI_FAILURE = 1621

	// ERROR_INSTALL_LOG_FAILURE
	// Error opening installation log file. Verify that the specified log file location exists and that you can write to it.
	ERROR_INSTALL_LOG_FAILURE = 1622

	// ERROR_INSTALL_LANGUAGE_UNSUPPORTED
	// The language of this installation package is not supported by your system.
	ERROR_INSTALL_LANGUAGE_UNSUPPORTED = 1623

	// ERROR_INSTALL_TRANSFORM_FAILURE
	// Error applying transforms. Verify that the specified transform paths are valid.
	ERROR_INSTALL_TRANSFORM_FAILURE = 1624

	// ERROR_INSTALL_PACKAGE_REJECTED
	// This installation is forbidden by system policy. Contact your system administrator.
	ERROR_INSTALL_PACKAGE_REJECTED = 1625

	// ERROR_FUNCTION_NOT_CALLED
	// Function could not be executed.
	ERROR_FUNCTION_NOT_CALLED = 1626

	// ERROR_FUNCTION_FAILED
	// Function failed during execution.
	ERROR_FUNCTION_FAILED = 1627

	// ERROR_INVALID_TABLE
	// Invalid or unknown table specified.
	ERROR_INVALID_TABLE = 1628

	// ERROR_DATATYPE_MISMATCH
	// Data supplied is of wrong type.
	ERROR_DATATYPE_MISMATCH = 1629

	// ERROR_UNSUPPORTED_TYPE
	// Data of this type is not supported.
	ERROR_UNSUPPORTED_TYPE = 1630

	// ERROR_CREATE_FAILED
	// The Windows Installer service failed to start. Contact your support personnel.
	ERROR_CREATE_FAILED = 1631

	// ERROR_INSTALL_TEMP_UNWRITABLE
	// The Temp folder is on a drive that is full or is inaccessible. Free up space on the drive or verify that you have write permission on the Temp folder.
	ERROR_INSTALL_TEMP_UNWRITABLE = 1632

	// ERROR_INSTALL_PLATFORM_UNSUPPORTED
	// This installation package is not supported by this processor type. Contact your product vendor.
	ERROR_INSTALL_PLATFORM_UNSUPPORTED = 1633

	// ERROR_INSTALL_NOTUSED
	// Component not used on this computer.
	ERROR_INSTALL_NOTUSED = 1634

	// ERROR_PATCH_PACKAGE_OPEN_FAILED
	// This update package could not be opened. Verify that the update package exists and that you can access it, or contact the application vendor to verify that this is a valid Windows Installer update package.
	ERROR_PATCH_PACKAGE_OPEN_FAILED = 1635

	// ERROR_PATCH_PACKAGE_INVALID
	// This update package could not be opened. Contact the application vendor to verify that this is a valid Windows Installer update package.
	ERROR_PATCH_PACKAGE_INVALID = 1636

	// ERROR_PATCH_PACKAGE_UNSUPPORTED
	// This update package cannot be processed by the Windows Installer service. You must install a Windows service pack that contains a newer version of the Windows Installer service.
	ERROR_PATCH_PACKAGE_UNSUPPORTED = 1637

	// ERROR_PRODUCT_VERSION
	// Another version of this product is already installed. Installation of this version cannot continue. To configure or remove the existing version of this product, use Add/Remove Programs on the Control Panel.
	ERROR_PRODUCT_VERSION = 1638

	// ERROR_INVALID_COMMAND_LINE
	// Invalid command line argument. Consult the Windows Installer SDK for detailed command line help.
	ERROR_INVALID_COMMAND_LINE = 1639

	// ERROR_INSTALL_REMOTE_DISALLOWED
	// Only administrators have permission to add, remove, or configure server software during a Terminal services remote session. If you want to install or configure software on the server, contact your network administrator.
	ERROR_INSTALL_REMOTE_DISALLOWED = 1640

	// ERROR_SUCCESS_REBOOT_INITIATED
	// The requested operation completed successfully. The system will be restarted so the changes can take effect.
	ERROR_SUCCESS_REBOOT_INITIATED = 1641

	// ERROR_PATCH_TARGET_NOT_FOUND
	// The upgrade cannot be installed by the Windows Installer service because the program to be upgraded may be missing, or the upgrade may update a different version of the program. Verify that the program to be upgraded exists on your computer and that you have the correct upgrade.
	ERROR_PATCH_TARGET_NOT_FOUND = 1642

	// ERROR_PATCH_PACKAGE_REJECTED
	// The update package is not permitted by software restriction policy.
	ERROR_PATCH_PACKAGE_REJECTED = 1643

	// ERROR_INSTALL_TRANSFORM_REJECTED
	// One or more customizations are not permitted by software restriction policy.
	ERROR_INSTALL_TRANSFORM_REJECTED = 1644

	// ERROR_INSTALL_REMOTE_PROHIBITED
	// The Windows Installer does not permit installation from a Remote Desktop Connection.
	ERROR_INSTALL_REMOTE_PROHIBITED = 1645

	// ERROR_PATCH_REMOVAL_UNSUPPORTED
	// Uninstallation of the update package is not supported.
	ERROR_PATCH_REMOVAL_UNSUPPORTED = 1646

	// ERROR_UNKNOWN_PATCH
	// The update is not applied to this product.
	ERROR_UNKNOWN_PATCH = 1647

	// ERROR_PATCH_NO_SEQUENCE
	// No valid sequence could be found for the set of updates.
	ERROR_PATCH_NO_SEQUENCE = 1648

	// ERROR_PATCH_REMOVAL_DISALLOWED
	// Update removal was disallowed by policy.
	ERROR_PATCH_REMOVAL_DISALLOWED = 1649

	// ERROR_INVALID_PATCH_XM
	// The XML update data is invalid.
	ERROR_INVALID_PATCH_XML = 1650

	// ERROR_PATCH_MANAGED_ADVERTISED_PRODUCT
	// Windows Installer does not permit updating of managed advertised products. At least one feature of the product must be installed before applying the update.
	ERROR_PATCH_MANAGED_ADVERTISED_PRODUCT = 1651

	// ERROR_INSTALL_SERVICE_SAFEBOOT
	// The Windows Installer service is not accessible in Safe Mode. Please try again when your computer is not in Safe Mode or you can use System Restore to return your machine to a previous good state.
	ERROR_INSTALL_SERVICE_SAFEBOOT = 1652

	// ERROR_FAIL_FAST_EXCEPTION
	// A fail fast exception occurred. Exception handlers will not be invoked and the process will be terminated immediately.
	ERROR_FAIL_FAST_EXCEPTION = 1653

	// ERROR_INSTALL_REJECTED
	// The app that you are trying to run is not supported on this version of Windows.
	ERROR_INSTALL_REJECTED = 1654

	// ERROR_DYNAMIC_CODE_BLOCKED
	// The operation was blocked as the process prohibits dynamic code generation.
	ERROR_DYNAMIC_CODE_BLOCKED = 1655

	// ERROR_NOT_SAME_OBJECT
	// The objects are not identical.
	ERROR_NOT_SAME_OBJECT = 1656

	// ERROR_STRICT_CFG_VIOLATION
	// The specified image file was blocked from loading because it does not enable a feature required by the process: Control Flow Guard.
	ERROR_STRICT_CFG_VIOLATION = 1657

	// ERROR_SET_CONTEXT_DENIED
	// The thread context could not be updated because this has been restricted for the process.
	ERROR_SET_CONTEXT_DENIED = 1660

	// ERROR_CROSS_PARTITION_VIOLATION
	// An invalid cross-partition private file/section access was attempted.
	ERROR_CROSS_PARTITION_VIOLATION = 1661

	///////////////////////////////////////////////////
	//                                               //
	//               RPC Error codes                 //
	//                                               //
	//                =1700 to=1999                  //
	///////////////////////////////////////////////////

	// RPC_S_INVALID_STRING_BINDING
	// The string binding is invalid.

	RPC_S_INVALID_STRING_BINDING = 1700

	// RPC_S_WRONG_KIND_OF_BINDING
	// The binding handle is not the correct type.

	RPC_S_WRONG_KIND_OF_BINDING = 1701

	// RPC_S_INVALID_BINDING
	// The binding handle is invalid.

	RPC_S_INVALID_BINDING = 1702

	// RPC_S_PROTSEQ_NOT_SUPPORTED
	// The RPC protocol sequence is not supported.

	RPC_S_PROTSEQ_NOT_SUPPORTED = 1703

	// RPC_S_INVALID_RPC_PROTSEQ
	// The RPC protocol sequence is invalid.

	RPC_S_INVALID_RPC_PROTSEQ = 1704

	// RPC_S_INVALID_STRING_UUID
	// The string universal unique identifier (UUID) is invalid.

	RPC_S_INVALID_STRING_UUID = 1705

	// RPC_S_INVALID_ENDPOINT_FORMAT
	// The endpoint format is invalid.

	RPC_S_INVALID_ENDPOINT_FORMAT = 1706

	// RPC_S_INVALID_NET_ADDR
	// The network address is invalid.

	RPC_S_INVALID_NET_ADDR = 1707

	// RPC_S_NO_ENDPOINT_FOUND
	// No endpoint was found.

	RPC_S_NO_ENDPOINT_FOUND = 1708

	// RPC_S_INVALID_TIMEOUT
	// The timeout value is invalid.

	RPC_S_INVALID_TIMEOUT = 1709

	// RPC_S_OBJECT_NOT_FOUND
	// The object universal unique identifier (UUID) was not found.

	RPC_S_OBJECT_NOT_FOUND = 1710

	// RPC_S_ALREADY_REGISTERED
	// The object universal unique identifier (UUID) has already been registered.

	RPC_S_ALREADY_REGISTERED = 1711

	// RPC_S_TYPE_ALREADY_REGISTERED
	// The type universal unique identifier (UUID) has already been registered.

	RPC_S_TYPE_ALREADY_REGISTERED = 1712

	// RPC_S_ALREADY_LISTENING
	// The RPC server is already listening.

	RPC_S_ALREADY_LISTENING = 1713

	// RPC_S_NO_PROTSEQS_REGISTERED
	// No protocol sequences have been registered.

	RPC_S_NO_PROTSEQS_REGISTERED = 1714

	// RPC_S_NOT_LISTENING
	// The RPC server is not listening.

	RPC_S_NOT_LISTENING = 1715

	// RPC_S_UNKNOWN_MGR_TYPE
	// The manager type is unknown.

	RPC_S_UNKNOWN_MGR_TYPE = 1716

	// RPC_S_UNKNOWN_IF
	// The interface is unknown.

	RPC_S_UNKNOWN_IF = 1717

	// RPC_S_NO_BINDINGS
	// There are no bindings.

	RPC_S_NO_BINDINGS = 1718

	// RPC_S_NO_PROTSEQS
	// There are no protocol sequences.

	RPC_S_NO_PROTSEQS = 1719

	// RPC_S_CANT_CREATE_ENDPOINT
	// The endpoint cannot be created.

	RPC_S_CANT_CREATE_ENDPOINT = 1720

	// RPC_S_OUT_OF_RESOURCES
	// Not enough resources are available to complete this operation.

	RPC_S_OUT_OF_RESOURCES = 1721

	// RPC_S_SERVER_UNAVAILABLE
	// The RPC server is unavailable.

	RPC_S_SERVER_UNAVAILABLE = 1722

	// RPC_S_SERVER_TOO_BUSY
	// The RPC server is too busy to complete this operation.

	RPC_S_SERVER_TOO_BUSY = 1723

	// RPC_S_INVALID_NETWORK_OPTIONS
	// The network options are invalid.

	RPC_S_INVALID_NETWORK_OPTIONS = 1724

	// RPC_S_NO_CALL_ACTIVE
	// There are no remote procedure calls active on this thread.

	RPC_S_NO_CALL_ACTIVE = 1725

	// RPC_S_CALL_FAILED
	// The remote procedure call failed.

	RPC_S_CALL_FAILED = 1726

	// RPC_S_CALL_FAILED_DNE
	// The remote procedure call failed and did not execute.

	RPC_S_CALL_FAILED_DNE = 1727

	// RPC_S_PROTOCOL_ERROR
	// A remote procedure call (RPC) protocol error occurred.

	RPC_S_PROTOCOL_ERROR = 1728

	// RPC_S_PROXY_ACCESS_DENIED
	// Access to the HTTP proxy is denied.

	RPC_S_PROXY_ACCESS_DENIED = 1729

	// RPC_S_UNSUPPORTED_TRANS_SYN
	// The transfer syntax is not supported by the RPC server.

	RPC_S_UNSUPPORTED_TRANS_SYN = 1730

	// RPC_S_UNSUPPORTED_TYPE
	// The universal unique identifier (UUID) type is not supported.

	RPC_S_UNSUPPORTED_TYPE = 1732

	// RPC_S_INVALID_TAG
	// The tag is invalid.

	RPC_S_INVALID_TAG = 1733

	// RPC_S_INVALID_BOUND
	// The array bounds are invalid.

	RPC_S_INVALID_BOUND = 1734

	// RPC_S_NO_ENTRY_NAME
	// The binding does not contain an entry name.

	RPC_S_NO_ENTRY_NAME = 1735

	// RPC_S_INVALID_NAME_SYNTAX
	// The name syntax is invalid.

	RPC_S_INVALID_NAME_SYNTAX = 1736

	// RPC_S_UNSUPPORTED_NAME_SYNTAX
	// The name syntax is not supported.

	RPC_S_UNSUPPORTED_NAME_SYNTAX = 1737

	// RPC_S_UUID_NO_ADDRESS
	// No network address is available to use to construct a universal unique identifier (UUID).

	RPC_S_UUID_NO_ADDRESS = 1739

	// RPC_S_DUPLICATE_ENDPOINT
	// The endpoint is a duplicate.

	RPC_S_DUPLICATE_ENDPOINT = 1740

	// RPC_S_UNKNOWN_AUTHN_TYPE
	// The authentication type is unknown.

	RPC_S_UNKNOWN_AUTHN_TYPE = 1741

	// RPC_S_MAX_CALLS_TOO_SMAL
	// The maximum number of calls is too small.

	RPC_S_MAX_CALLS_TOO_SMALL = 1742

	// RPC_S_STRING_TOO_LONG
	// The string is too long.

	RPC_S_STRING_TOO_LONG = 1743

	// RPC_S_PROTSEQ_NOT_FOUND
	// The RPC protocol sequence was not found.

	RPC_S_PROTSEQ_NOT_FOUND = 1744

	// RPC_S_PROCNUM_OUT_OF_RANGE
	// The procedure number is out of range.

	RPC_S_PROCNUM_OUT_OF_RANGE = 1745

	// RPC_S_BINDING_HAS_NO_AUTH
	// The binding does not contain any authentication information.

	RPC_S_BINDING_HAS_NO_AUTH = 1746

	// RPC_S_UNKNOWN_AUTHN_SERVICE
	// The authentication service is unknown.

	RPC_S_UNKNOWN_AUTHN_SERVICE = 1747

	// RPC_S_UNKNOWN_AUTHN_LEVE
	// The authentication level is unknown.

	RPC_S_UNKNOWN_AUTHN_LEVEL = 1748

	// RPC_S_INVALID_AUTH_IDENTITY
	// The security context is invalid.

	RPC_S_INVALID_AUTH_IDENTITY = 1749

	// RPC_S_UNKNOWN_AUTHZ_SERVICE
	// The authorization service is unknown.

	RPC_S_UNKNOWN_AUTHZ_SERVICE = 1750

	// EPT_S_INVALID_ENTRY
	// The entry is invalid.
	EPT_S_INVALID_ENTRY = 1751

	// EPT_S_CANT_PERFORM_OP
	// The server endpoint cannot perform the operation.
	EPT_S_CANT_PERFORM_OP = 1752

	// EPT_S_NOT_REGISTERED
	// There are no more endpoints available from the endpoint mapper.
	EPT_S_NOT_REGISTERED = 1753

	// RPC_S_NOTHING_TO_EXPORT
	// No interfaces have been exported.

	RPC_S_NOTHING_TO_EXPORT = 1754

	// RPC_S_INCOMPLETE_NAME
	// The entry name is incomplete.

	RPC_S_INCOMPLETE_NAME = 1755

	// RPC_S_INVALID_VERS_OPTION
	// The version option is invalid.

	RPC_S_INVALID_VERS_OPTION = 1756

	// RPC_S_NO_MORE_MEMBERS
	// There are no more members.

	RPC_S_NO_MORE_MEMBERS = 1757

	// RPC_S_NOT_ALL_OBJS_UNEXPORTED
	// There is nothing to unexport.

	RPC_S_NOT_ALL_OBJS_UNEXPORTED = 1758

	// RPC_S_INTERFACE_NOT_FOUND
	// The interface was not found.

	RPC_S_INTERFACE_NOT_FOUND = 1759

	// RPC_S_ENTRY_ALREADY_EXISTS
	// The entry already exists.

	RPC_S_ENTRY_ALREADY_EXISTS = 1760

	// RPC_S_ENTRY_NOT_FOUND
	// The entry is not found.

	RPC_S_ENTRY_NOT_FOUND = 1761

	// RPC_S_NAME_SERVICE_UNAVAILABLE
	// The name service is unavailable.

	RPC_S_NAME_SERVICE_UNAVAILABLE = 1762

	// RPC_S_INVALID_NAF_ID
	// The network address family is invalid.

	RPC_S_INVALID_NAF_ID = 1763

	// RPC_S_CANNOT_SUPPORT
	// The requested operation is not supported.

	RPC_S_CANNOT_SUPPORT = 1764

	// RPC_S_NO_CONTEXT_AVAILABLE
	// No security context is available to allow impersonation.

	RPC_S_NO_CONTEXT_AVAILABLE = 1765

	// RPC_S_INTERNAL_ERROR
	// An internal error occurred in a remote procedure call (RPC).

	RPC_S_INTERNAL_ERROR = 1766

	// RPC_S_ZERO_DIVIDE
	// The RPC server attempted an integer division by zero.

	RPC_S_ZERO_DIVIDE = 1767

	// RPC_S_ADDRESS_ERROR
	// An addressing error occurred in the RPC server.

	RPC_S_ADDRESS_ERROR = 1768

	// RPC_S_FP_DIV_ZERO
	// A floating-point operation at the RPC server caused a division by zero.

	RPC_S_FP_DIV_ZERO = 1769

	// RPC_S_FP_UNDERFLOW
	// A floating-point underflow occurred at the RPC server.

	RPC_S_FP_UNDERFLOW = 1770

	// RPC_S_FP_OVERFLOW
	// A floating-point overflow occurred at the RPC server.

	RPC_S_FP_OVERFLOW = 1771

	// RPC_X_NO_MORE_ENTRIES
	// The list of RPC servers available for the binding of auto handles has been exhausted.

	RPC_X_NO_MORE_ENTRIES = 1772

	// RPC_X_SS_CHAR_TRANS_OPEN_FAI
	// Unable to open the character translation table file.

	RPC_X_SS_CHAR_TRANS_OPEN_FAIL = 1773

	// RPC_X_SS_CHAR_TRANS_SHORT_FILE
	// The file containing the character translation table has fewer than=512 bytes.

	RPC_X_SS_CHAR_TRANS_SHORT_FILE = 1774

	// RPC_X_SS_IN_NULL_CONTEXT
	// A null context handle was passed from the client to the host during a remote procedure call.

	RPC_X_SS_IN_NULL_CONTEXT = 1775

	// RPC_X_SS_CONTEXT_DAMAGED
	// The context handle changed during a remote procedure call.

	RPC_X_SS_CONTEXT_DAMAGED = 1777

	// RPC_X_SS_HANDLES_MISMATCH
	// The binding handles passed to a remote procedure call do not match.

	RPC_X_SS_HANDLES_MISMATCH = 1778

	// RPC_X_SS_CANNOT_GET_CALL_HANDLE
	// The stub is unable to get the remote procedure call handle.

	RPC_X_SS_CANNOT_GET_CALL_HANDLE = 1779

	// RPC_X_NULL_REF_POINTER
	// A null reference pointer was passed to the stub.

	RPC_X_NULL_REF_POINTER = 1780

	// RPC_X_ENUM_VALUE_OUT_OF_RANGE
	// The enumeration value is out of range.

	RPC_X_ENUM_VALUE_OUT_OF_RANGE = 1781

	// RPC_X_BYTE_COUNT_TOO_SMAL
	// The byte count is too small.

	RPC_X_BYTE_COUNT_TOO_SMALL = 1782

	// RPC_X_BAD_STUB_DATA
	// The stub received bad data.

	RPC_X_BAD_STUB_DATA = 1783

	// ERROR_INVALID_USER_BUFFER
	// The supplied user buffer is not valid for the requested operation.
	ERROR_INVALID_USER_BUFFER = 1784

	// ERROR_UNRECOGNIZED_MEDIA
	// The disk media is not recognized. It may not be formatted.
	ERROR_UNRECOGNIZED_MEDIA = 1785

	// ERROR_NO_TRUST_LSA_SECRET
	// The workstation does not have a trust secret.
	ERROR_NO_TRUST_LSA_SECRET = 1786

	// ERROR_NO_TRUST_SAM_ACCOUNT
	// The security database on the server does not have a computer account for this workstation trust relationship.
	ERROR_NO_TRUST_SAM_ACCOUNT = 1787

	// ERROR_TRUSTED_DOMAIN_FAILURE
	// The trust relationship between the primary domain and the trusted domain failed.
	ERROR_TRUSTED_DOMAIN_FAILURE = 1788

	// ERROR_TRUSTED_RELATIONSHIP_FAILURE
	// The trust relationship between this workstation and the primary domain failed.
	ERROR_TRUSTED_RELATIONSHIP_FAILURE = 1789

	// ERROR_TRUST_FAILURE
	// The network logon failed.
	ERROR_TRUST_FAILURE = 1790

	// RPC_S_CALL_IN_PROGRESS
	// A remote procedure call is already in progress for this thread.

	RPC_S_CALL_IN_PROGRESS = 1791

	// ERROR_NETLOGON_NOT_STARTED
	// An attempt was made to logon, but the network logon service was not started.
	ERROR_NETLOGON_NOT_STARTED = 1792

	// ERROR_ACCOUNT_EXPIRED
	// The user's account has expired.
	ERROR_ACCOUNT_EXPIRED = 1793

	// ERROR_REDIRECTOR_HAS_OPEN_HANDLES
	// The redirector is in use and cannot be unloaded.
	ERROR_REDIRECTOR_HAS_OPEN_HANDLES = 1794

	// ERROR_PRINTER_DRIVER_ALREADY_INSTALLED
	// The specified printer driver is already installed.
	ERROR_PRINTER_DRIVER_ALREADY_INSTALLED = 1795

	// ERROR_UNKNOWN_PORT
	// The specified port is unknown.
	ERROR_UNKNOWN_PORT = 1796

	// ERROR_UNKNOWN_PRINTER_DRIVER
	// The printer driver is unknown.
	ERROR_UNKNOWN_PRINTER_DRIVER = 1797

	// ERROR_UNKNOWN_PRINTPROCESSOR
	// The print processor is unknown.
	ERROR_UNKNOWN_PRINTPROCESSOR = 1798

	// ERROR_INVALID_SEPARATOR_FILE
	// The specified separator file is invalid.
	ERROR_INVALID_SEPARATOR_FILE = 1799

	// ERROR_INVALID_PRIORITY
	// The specified priority is invalid.
	ERROR_INVALID_PRIORITY = 1800

	// ERROR_INVALID_PRINTER_NAME
	// The printer name is invalid.
	ERROR_INVALID_PRINTER_NAME = 1801

	// ERROR_PRINTER_ALREADY_EXISTS
	// The printer already exists.
	ERROR_PRINTER_ALREADY_EXISTS = 1802

	// ERROR_INVALID_PRINTER_COMMAND
	// The printer command is invalid.
	ERROR_INVALID_PRINTER_COMMAND = 1803

	// ERROR_INVALID_DATATYPE
	// The specified datatype is invalid.
	ERROR_INVALID_DATATYPE = 1804

	// ERROR_INVALID_ENVIRONMENT
	// The environment specified is invalid.
	ERROR_INVALID_ENVIRONMENT = 1805

	// RPC_S_NO_MORE_BINDINGS
	// There are no more bindings.

	RPC_S_NO_MORE_BINDINGS = 1806

	// ERROR_NOLOGON_INTERDOMAIN_TRUST_ACCOUNT
	// The account used is an interdomain trust account. Use your global user account or local user account to access this server.
	ERROR_NOLOGON_INTERDOMAIN_TRUST_ACCOUNT = 1807

	// ERROR_NOLOGON_WORKSTATION_TRUST_ACCOUNT
	// The account used is a computer account. Use your global user account or local user account to access this server.
	ERROR_NOLOGON_WORKSTATION_TRUST_ACCOUNT = 1808

	// ERROR_NOLOGON_SERVER_TRUST_ACCOUNT
	// The account used is a server trust account. Use your global user account or local user account to access this server.
	ERROR_NOLOGON_SERVER_TRUST_ACCOUNT = 1809

	// ERROR_DOMAIN_TRUST_INCONSISTENT
	// The name or security ID (SID) of the domain specified is inconsistent with the trust information for that domain.
	ERROR_DOMAIN_TRUST_INCONSISTENT = 1810

	// ERROR_SERVER_HAS_OPEN_HANDLES
	// The server is in use and cannot be unloaded.
	ERROR_SERVER_HAS_OPEN_HANDLES = 1811

	// ERROR_RESOURCE_DATA_NOT_FOUND
	// The specified image file did not contain a resource section.
	ERROR_RESOURCE_DATA_NOT_FOUND = 1812

	// ERROR_RESOURCE_TYPE_NOT_FOUND
	// The specified resource type cannot be found in the image file.
	ERROR_RESOURCE_TYPE_NOT_FOUND = 1813

	// ERROR_RESOURCE_NAME_NOT_FOUND
	// The specified resource name cannot be found in the image file.
	ERROR_RESOURCE_NAME_NOT_FOUND = 1814

	// ERROR_RESOURCE_LANG_NOT_FOUND
	// The specified resource language ID cannot be found in the image file.
	ERROR_RESOURCE_LANG_NOT_FOUND = 1815

	// ERROR_NOT_ENOUGH_QUOTA
	// Not enough quota is available to process this command.
	ERROR_NOT_ENOUGH_QUOTA = 1816

	// RPC_S_NO_INTERFACES
	// No interfaces have been registered.

	RPC_S_NO_INTERFACES = 1817

	// RPC_S_CALL_CANCELLED
	// The remote procedure call was cancelled.

	RPC_S_CALL_CANCELLED = 1818

	// RPC_S_BINDING_INCOMPLETE
	// The binding handle does not contain all required information.

	RPC_S_BINDING_INCOMPLETE = 1819

	// RPC_S_COMM_FAILURE
	// A communications failure occurred during a remote procedure call.

	RPC_S_COMM_FAILURE = 1820

	// RPC_S_UNSUPPORTED_AUTHN_LEVE
	// The requested authentication level is not supported.

	RPC_S_UNSUPPORTED_AUTHN_LEVEL = 1821

	// RPC_S_NO_PRINC_NAME
	// No principal name registered.

	RPC_S_NO_PRINC_NAME = 1822

	// RPC_S_NOT_RPC_ERROR
	// The error specified is not a valid Windows RPC error code.

	RPC_S_NOT_RPC_ERROR = 1823

	// RPC_S_UUID_LOCAL_ONLY
	// A UUID that is valid only on this computer has been allocated.

	RPC_S_UUID_LOCAL_ONLY = 1824

	// RPC_S_SEC_PKG_ERROR
	// A security package specific error occurred.

	RPC_S_SEC_PKG_ERROR = 1825

	// RPC_S_NOT_CANCELLED
	// Thread is not canceled.

	RPC_S_NOT_CANCELLED = 1826

	// RPC_X_INVALID_ES_ACTION
	// Invalid operation on the encoding/decoding handle.

	RPC_X_INVALID_ES_ACTION = 1827

	// RPC_X_WRONG_ES_VERSION
	// Incompatible version of the serializing package.

	RPC_X_WRONG_ES_VERSION = 1828

	// RPC_X_WRONG_STUB_VERSION
	// Incompatible version of the RPC stub.

	RPC_X_WRONG_STUB_VERSION = 1829

	// RPC_X_INVALID_PIPE_OBJECT
	// The RPC pipe object is invalid or corrupted.

	RPC_X_INVALID_PIPE_OBJECT = 1830

	// RPC_X_WRONG_PIPE_ORDER
	// An invalid operation was attempted on an RPC pipe object.

	RPC_X_WRONG_PIPE_ORDER = 1831

	// RPC_X_WRONG_PIPE_VERSION
	// Unsupported RPC pipe version.

	RPC_X_WRONG_PIPE_VERSION = 1832

	// RPC_S_COOKIE_AUTH_FAILED
	// HTTP proxy server rejected the connection because the cookie authentication failed.

	RPC_S_COOKIE_AUTH_FAILED = 1833

	// RPC_S_DO_NOT_DISTURB
	// The RPC server is suspended, and could not be resumed for this request. The call did not execute.

	RPC_S_DO_NOT_DISTURB = 1834

	// RPC_S_SYSTEM_HANDLE_COUNT_EXCEEDED
	// The RPC call contains too many handles to be transmitted in a single request.

	RPC_S_SYSTEM_HANDLE_COUNT_EXCEEDED = 1835

	// RPC_S_SYSTEM_HANDLE_TYPE_MISMATCH
	// The RPC call contains a handle that differs from the declared handle type.

	RPC_S_SYSTEM_HANDLE_TYPE_MISMATCH = 1836

	// RPC_S_GROUP_MEMBER_NOT_FOUND
	// The group member was not found.

	RPC_S_GROUP_MEMBER_NOT_FOUND = 1898

	// EPT_S_CANT_CREATE
	// The endpoint mapper database entry could not be created.
	EPT_S_CANT_CREATE = 1899

	// RPC_S_INVALID_OBJECT
	// The object universal unique identifier (UUID) is the nil UUID.

	RPC_S_INVALID_OBJECT = 1900

	// ERROR_INVALID_TIME
	// The specified time is invalid.
	ERROR_INVALID_TIME = 1901

	// ERROR_INVALID_FORM_NAME
	// The specified form name is invalid.
	ERROR_INVALID_FORM_NAME = 1902

	// ERROR_INVALID_FORM_SIZE
	// The specified form size is invalid.
	ERROR_INVALID_FORM_SIZE = 1903

	// ERROR_ALREADY_WAITING
	// The specified printer handle is already being waited on
	ERROR_ALREADY_WAITING = 1904

	// ERROR_PRINTER_DELETED
	// The specified printer has been deleted.
	ERROR_PRINTER_DELETED = 1905

	// ERROR_INVALID_PRINTER_STATE
	// The state of the printer is invalid.
	ERROR_INVALID_PRINTER_STATE = 1906

	// ERROR_PASSWORD_MUST_CHANGE
	// The user's password must be changed before signing in.
	ERROR_PASSWORD_MUST_CHANGE = 1907

	// ERROR_DOMAIN_CONTROLLER_NOT_FOUND
	// Could not find the domain controller for this domain.
	ERROR_DOMAIN_CONTROLLER_NOT_FOUND = 1908

	// ERROR_ACCOUNT_LOCKED_OUT
	// The referenced account is currently locked out and may not be logged on to.
	ERROR_ACCOUNT_LOCKED_OUT = 1909

	// OR_INVALID_OXID
	// The object exporter specified was not found.

	OR_INVALID_OXID = 1910

	// OR_INVALID_OID
	// The object specified was not found.

	OR_INVALID_OID = 1911

	// OR_INVALID_SET
	// The object resolver set specified was not found.

	OR_INVALID_SET = 1912

	// RPC_S_SEND_INCOMPLETE
	// Some data remains to be sent in the request buffer.

	RPC_S_SEND_INCOMPLETE = 1913

	// RPC_S_INVALID_ASYNC_HANDLE
	// Invalid asynchronous remote procedure call handle.

	RPC_S_INVALID_ASYNC_HANDLE = 1914

	// RPC_S_INVALID_ASYNC_CAL
	// Invalid asynchronous RPC call handle for this operation.

	RPC_S_INVALID_ASYNC_CALL = 1915

	// RPC_X_PIPE_CLOSED
	// The RPC pipe object has already been closed.

	RPC_X_PIPE_CLOSED = 1916

	// RPC_X_PIPE_DISCIPLINE_ERROR
	// The RPC call completed before all pipes were processed.

	RPC_X_PIPE_DISCIPLINE_ERROR = 1917

	// RPC_X_PIPE_EMPTY
	// No more data is available from the RPC pipe.

	RPC_X_PIPE_EMPTY = 1918

	// ERROR_NO_SITENAME
	// No site name is available for this machine.
	ERROR_NO_SITENAME = 1919

	// ERROR_CANT_ACCESS_FILE
	// The file cannot be accessed by the system.
	ERROR_CANT_ACCESS_FILE = 1920

	// ERROR_CANT_RESOLVE_FILENAME
	// The name of the file cannot be resolved by the system.
	ERROR_CANT_RESOLVE_FILENAME = 1921

	// RPC_S_ENTRY_TYPE_MISMATCH
	// The entry is not of the expected type.

	RPC_S_ENTRY_TYPE_MISMATCH = 1922

	// RPC_S_NOT_ALL_OBJS_EXPORTED
	// Not all object UUIDs could be exported to the specified entry.

	RPC_S_NOT_ALL_OBJS_EXPORTED = 1923

	// RPC_S_INTERFACE_NOT_EXPORTED
	// Interface could not be exported to the specified entry.

	RPC_S_INTERFACE_NOT_EXPORTED = 1924

	// RPC_S_PROFILE_NOT_ADDED
	// The specified profile entry could not be added.

	RPC_S_PROFILE_NOT_ADDED = 1925

	// RPC_S_PRF_ELT_NOT_ADDED
	// The specified profile element could not be added.

	RPC_S_PRF_ELT_NOT_ADDED = 1926

	// RPC_S_PRF_ELT_NOT_REMOVED
	// The specified profile element could not be removed.

	RPC_S_PRF_ELT_NOT_REMOVED = 1927

	// RPC_S_GRP_ELT_NOT_ADDED
	// The group element could not be added.

	RPC_S_GRP_ELT_NOT_ADDED = 1928

	// RPC_S_GRP_ELT_NOT_REMOVED
	// The group element could not be removed.

	RPC_S_GRP_ELT_NOT_REMOVED = 1929

	// ERROR_KM_DRIVER_BLOCKED
	// The printer driver is not compatible with a policy enabled on your computer that blocks NT=4.0 drivers.
	ERROR_KM_DRIVER_BLOCKED = 1930

	// ERROR_CONTEXT_EXPIRED
	// The context has expired and can no longer be used.
	ERROR_CONTEXT_EXPIRED = 1931

	// ERROR_PER_USER_TRUST_QUOTA_EXCEEDED
	// The current user's delegated trust creation quota has been exceeded.
	ERROR_PER_USER_TRUST_QUOTA_EXCEEDED = 1932

	// ERROR_ALL_USER_TRUST_QUOTA_EXCEEDED
	// The total delegated trust creation quota has been exceeded.
	ERROR_ALL_USER_TRUST_QUOTA_EXCEEDED = 1933

	// ERROR_USER_DELETE_TRUST_QUOTA_EXCEEDED
	// The current user's delegated trust deletion quota has been exceeded.
	ERROR_USER_DELETE_TRUST_QUOTA_EXCEEDED = 1934

	// ERROR_AUTHENTICATION_FIREWALL_FAILED
	// The computer you are signing into is protected by an authentication firewall. The specified account is not allowed to authenticate to the computer.
	ERROR_AUTHENTICATION_FIREWALL_FAILED = 1935

	// ERROR_REMOTE_PRINT_CONNECTIONS_BLOCKED
	// Remote connections to the Print Spooler are blocked by a policy set on your machine.
	ERROR_REMOTE_PRINT_CONNECTIONS_BLOCKED = 1936

	// ERROR_NTLM_BLOCKED
	// Authentication failed because NTLM authentication has been disabled.
	ERROR_NTLM_BLOCKED = 1937

	// ERROR_PASSWORD_CHANGE_REQUIRED
	// Logon Failure: EAS policy requires that the user change their password before this operation can be performed.
	ERROR_PASSWORD_CHANGE_REQUIRED = 1938

	// ERROR_LOST_MODE_LOGON_RESTRICTION
	// An administrator has restricted sign in. To sign in, make sure your device is connected to the Internet, and have your administrator sign in first.
	ERROR_LOST_MODE_LOGON_RESTRICTION = 1939

	///////////////////////////////////////////////////
	//                                               //
	//              OpenGL Error codes               //
	//                                               //
	//                =2000 to=2009                  //
	///////////////////////////////////////////////////

	// ERROR_INVALID_PIXEL_FORMAT
	// The pixel format is invalid.
	ERROR_INVALID_PIXEL_FORMAT = 2000

	// ERROR_BAD_DRIVER
	// The specified driver is invalid.
	ERROR_BAD_DRIVER = 2001

	// ERROR_INVALID_WINDOW_STYLE
	// The window style or class attribute is invalid for this operation.
	ERROR_INVALID_WINDOW_STYLE = 2002

	// ERROR_METAFILE_NOT_SUPPORTED
	// The requested metafile operation is not supported.
	ERROR_METAFILE_NOT_SUPPORTED = 2003

	// ERROR_TRANSFORM_NOT_SUPPORTED
	// The requested transformation operation is not supported.
	ERROR_TRANSFORM_NOT_SUPPORTED = 2004

	// ERROR_CLIPPING_NOT_SUPPORTED
	// The requested clipping operation is not supported.
	ERROR_CLIPPING_NOT_SUPPORTED = 2005

	///////////////////////////////////////////////////
	//                                               //
	//       Image Color Management Error codes      //
	//                                               //
	//                =2010 to=2049                  //
	///////////////////////////////////////////////////

	// ERROR_INVALID_CMM
	// The specified color management module is invalid.
	ERROR_INVALID_CMM = 2010

	// ERROR_INVALID_PROFILE
	// The specified color profile is invalid.
	ERROR_INVALID_PROFILE = 2011

	// ERROR_TAG_NOT_FOUND
	// The specified tag was not found.
	ERROR_TAG_NOT_FOUND = 2012

	// ERROR_TAG_NOT_PRESENT
	// A required tag is not present.
	ERROR_TAG_NOT_PRESENT = 2013

	// ERROR_DUPLICATE_TAG
	// The specified tag is already present.
	ERROR_DUPLICATE_TAG = 2014

	// ERROR_PROFILE_NOT_ASSOCIATED_WITH_DEVICE
	// The specified color profile is not associated with the specified device.
	ERROR_PROFILE_NOT_ASSOCIATED_WITH_DEVICE = 2015

	// ERROR_PROFILE_NOT_FOUND
	// The specified color profile was not found.
	ERROR_PROFILE_NOT_FOUND = 2016

	// ERROR_INVALID_COLORSPACE
	// The specified color space is invalid.
	ERROR_INVALID_COLORSPACE = 2017

	// ERROR_ICM_NOT_ENABLED
	// Image Color Management is not enabled.
	ERROR_ICM_NOT_ENABLED = 2018

	// ERROR_DELETING_ICM_XFORM
	// There was an error while deleting the color transform.
	ERROR_DELETING_ICM_XFORM = 2019

	// ERROR_INVALID_TRANSFORM
	// The specified color transform is invalid.
	ERROR_INVALID_TRANSFORM = 2020

	// ERROR_COLORSPACE_MISMATCH
	// The specified transform does not match the bitmap's color space.
	ERROR_COLORSPACE_MISMATCH = 2021

	// ERROR_INVALID_COLORINDEX
	// The specified named color index is not present in the profile.
	ERROR_INVALID_COLORINDEX = 2022

	// ERROR_PROFILE_DOES_NOT_MATCH_DEVICE
	// The specified profile is intended for a device of a different type than the specified device.
	ERROR_PROFILE_DOES_NOT_MATCH_DEVICE = 2023

	///////////////////////////////////////////////////
	//                                               //
	//             Winnet32 Error codes              //
	//                                               //
	//                =2100 to=2999                  //
	//                                               //
	// The range=2100 through=2999 is reserved for   //
	// network status codes. See lmerr.h for a       //
	// complete listing                              //
	///////////////////////////////////////////////////

	// ERROR_CONNECTED_OTHER_PASSWORD
	// The network connection was made successfully, but the user had to be prompted for a password other than the one originally specified.
	ERROR_CONNECTED_OTHER_PASSWORD = 2108

	// ERROR_CONNECTED_OTHER_PASSWORD_DEFAULT
	// The network connection was made successfully using default credentials.
	ERROR_CONNECTED_OTHER_PASSWORD_DEFAULT = 2109

	// ERROR_BAD_USERNAME
	// The specified username is invalid.
	ERROR_BAD_USERNAME = 2202

	// ERROR_NOT_CONNECTED
	// This network connection does not exist.
	ERROR_NOT_CONNECTED = 2250

	// ERROR_OPEN_FILES
	// This network connection has files open or requests pending.
	ERROR_OPEN_FILES = 2401

	// ERROR_ACTIVE_CONNECTIONS
	// Active connections still exist.
	ERROR_ACTIVE_CONNECTIONS = 2402

	// ERROR_DEVICE_IN_USE
	// The device is in use by an active process and cannot be disconnected.
	ERROR_DEVICE_IN_USE = 2404

	///////////////////////////////////////////////////
	//                                               //
	//           Win32 Spooler Error codes           //
	//                                               //
	//                =3000 to=3049                  //
	///////////////////////////////////////////////////

	// ERROR_UNKNOWN_PRINT_MONITOR
	// The specified print monitor is unknown.
	ERROR_UNKNOWN_PRINT_MONITOR = 3000

	// ERROR_PRINTER_DRIVER_IN_USE
	// The specified printer driver is currently in use.
	ERROR_PRINTER_DRIVER_IN_USE = 3001

	// ERROR_SPOOL_FILE_NOT_FOUND
	// The spool file was not found.
	ERROR_SPOOL_FILE_NOT_FOUND = 3002

	// ERROR_SPL_NO_STARTDOC
	// A StartDocPrinter call was not issued.
	ERROR_SPL_NO_STARTDOC = 3003

	// ERROR_SPL_NO_ADDJOB
	// An AddJob call was not issued.
	ERROR_SPL_NO_ADDJOB = 3004

	// ERROR_PRINT_PROCESSOR_ALREADY_INSTALLED
	// The specified print processor has already been installed.
	ERROR_PRINT_PROCESSOR_ALREADY_INSTALLED = 3005

	// ERROR_PRINT_MONITOR_ALREADY_INSTALLED
	// The specified print monitor has already been installed.
	ERROR_PRINT_MONITOR_ALREADY_INSTALLED = 3006

	// ERROR_INVALID_PRINT_MONITOR
	// The specified print monitor does not have the required functions.
	ERROR_INVALID_PRINT_MONITOR = 3007

	// ERROR_PRINT_MONITOR_IN_USE
	// The specified print monitor is currently in use.
	ERROR_PRINT_MONITOR_IN_USE = 3008

	// ERROR_PRINTER_HAS_JOBS_QUEUED
	// The requested operation is not allowed when there are jobs queued to the printer.
	ERROR_PRINTER_HAS_JOBS_QUEUED = 3009

	// ERROR_SUCCESS_REBOOT_REQUIRED
	// The requested operation is successful. Changes will not be effective until the system is rebooted.
	ERROR_SUCCESS_REBOOT_REQUIRED = 3010

	// ERROR_SUCCESS_RESTART_REQUIRED
	// The requested operation is successful. Changes will not be effective until the service is restarted.
	ERROR_SUCCESS_RESTART_REQUIRED = 3011

	// ERROR_PRINTER_NOT_FOUND
	// No printers were found.
	ERROR_PRINTER_NOT_FOUND = 3012

	// ERROR_PRINTER_DRIVER_WARNED
	// The printer driver is known to be unreliable.
	ERROR_PRINTER_DRIVER_WARNED = 3013

	// ERROR_PRINTER_DRIVER_BLOCKED
	// The printer driver is known to harm the system.
	ERROR_PRINTER_DRIVER_BLOCKED = 3014

	// ERROR_PRINTER_DRIVER_PACKAGE_IN_USE
	// The specified printer driver package is currently in use.
	ERROR_PRINTER_DRIVER_PACKAGE_IN_USE = 3015

	// ERROR_CORE_DRIVER_PACKAGE_NOT_FOUND
	// Unable to find a core driver package that is required by the printer driver package.
	ERROR_CORE_DRIVER_PACKAGE_NOT_FOUND = 3016

	// ERROR_FAIL_REBOOT_REQUIRED
	// The requested operation failed. A system reboot is required to roll back changes made.
	ERROR_FAIL_REBOOT_REQUIRED = 3017

	// ERROR_FAIL_REBOOT_INITIATED
	// The requested operation failed. A system reboot has been initiated to roll back changes made.
	ERROR_FAIL_REBOOT_INITIATED = 3018

	// ERROR_PRINTER_DRIVER_DOWNLOAD_NEEDED
	// The specified printer driver was not found on the system and needs to be downloaded.
	ERROR_PRINTER_DRIVER_DOWNLOAD_NEEDED = 3019

	// ERROR_PRINT_JOB_RESTART_REQUIRED
	// The requested print job has failed to print. A print system update requires the job to be resubmitted.
	ERROR_PRINT_JOB_RESTART_REQUIRED = 3020

	// ERROR_INVALID_PRINTER_DRIVER_MANIFEST
	// The printer driver does not contain a valid manifest, or contains too many manifests.
	ERROR_INVALID_PRINTER_DRIVER_MANIFEST = 3021

	// ERROR_PRINTER_NOT_SHAREABLE
	// The specified printer cannot be shared.
	ERROR_PRINTER_NOT_SHAREABLE = 3022

	///////////////////////////////////////////////////
	//                                               //
	//           CopyFile ext. Error codes           //
	//                                               //
	//                =3050 to=3059                  //
	///////////////////////////////////////////////////

	// ERROR_REQUEST_PAUSED
	// The operation was paused.
	ERROR_REQUEST_PAUSED = 3050

	///////////////////////////////////////////////////
	//                                               //
	//           AppExec Error codes                 //
	//                                               //
	//                =3060 to=3079                  //
	///////////////////////////////////////////////////

	// ERROR_APPEXEC_CONDITION_NOT_SATISFIED
	// The condition supplied for the app execution request was not satisfied, so the request was not performed.
	ERROR_APPEXEC_CONDITION_NOT_SATISFIED = 3060

	// ERROR_APPEXEC_HANDLE_INVALIDATED
	// The supplied handle has been invalidated and may not be used for the requested operation.
	ERROR_APPEXEC_HANDLE_INVALIDATED = 3061

	// ERROR_APPEXEC_INVALID_HOST_GENERATION
	// The supplied host generation has been invalidated and may not be used for the requested operation.
	ERROR_APPEXEC_INVALID_HOST_GENERATION = 3062

	// ERROR_APPEXEC_UNEXPECTED_PROCESS_REGISTRATION
	// An attempt to register a process failed because the target host was not in a valid state to receive process registrations.
	ERROR_APPEXEC_UNEXPECTED_PROCESS_REGISTRATION = 3063

	// ERROR_APPEXEC_INVALID_HOST_STATE
	// The host is not in a valid state to support the execution request.
	ERROR_APPEXEC_INVALID_HOST_STATE = 3064

	// ERROR_APPEXEC_NO_DONOR
	// The operation was not completed because a required resource donor was not found for the host.
	ERROR_APPEXEC_NO_DONOR = 3065

	// ERROR_APPEXEC_HOST_ID_MISMATCH
	// The operation was not completed because an unexpected host ID was encountered.
	ERROR_APPEXEC_HOST_ID_MISMATCH = 3066

	// ERROR_APPEXEC_UNKNOWN_USER
	// The operation was not completed because the specified user was not known to the service.
	ERROR_APPEXEC_UNKNOWN_USER = 3067

	///////////////////////////////////////////////////
	//                                               //
	//                  Available                    //
	//                                               //
	//                =3080 to=3199                  //
	///////////////////////////////////////////////////

	//               the message range
	//                =3200 to=3299
	//      is reserved and used in isolation lib

	///////////////////////////////////////////////////
	//                                               //
	//                  Available                    //
	//                                               //
	//                =3300 to=3899                  //
	///////////////////////////////////////////////////

	///////////////////////////////////////////////////
	//                                               //
	//                IO Error Codes                 //
	//                                               //
	//                =3900 to=3999                  //
	///////////////////////////////////////////////////

	// ERROR_IO_REISSUE_AS_CACHED
	// Reissue the given operation as a cached IO operation
	ERROR_IO_REISSUE_AS_CACHED = 3950

	///////////////////////////////////////////////////
	//                                               //
	//                Wins Error codes               //
	//                                               //
	//                =4000 to=4049                  //
	///////////////////////////////////////////////////

	// ERROR_WINS_INTERNA
	// WINS encountered an error while processing the command.
	ERROR_WINS_INTERNAL = 4000

	// ERROR_CAN_NOT_DEL_LOCAL_WINS
	// The local WINS cannot be deleted.
	ERROR_CAN_NOT_DEL_LOCAL_WINS = 4001

	// ERROR_STATIC_INIT
	// The importation from the file failed.
	ERROR_STATIC_INIT = 4002

	// ERROR_INC_BACKUP
	// The backup failed. Was a full backup done before?
	ERROR_INC_BACKUP = 4003

	// ERROR_FULL_BACKUP
	// The backup failed. Check the directory to which you are backing the database.
	ERROR_FULL_BACKUP = 4004

	// ERROR_REC_NON_EXISTENT
	// The name does not exist in the WINS database.
	ERROR_REC_NON_EXISTENT = 4005

	// ERROR_RPL_NOT_ALLOWED
	// Replication with a nonconfigured partner is not allowed.
	ERROR_RPL_NOT_ALLOWED = 4006

	///////////////////////////////////////////////////
	//                                               //
	//              PeerDist Error codes             //
	//                                               //
	//                =4050 to=4099                  //
	///////////////////////////////////////////////////

	// PEERDIST_ERROR_CONTENTINFO_VERSION_UNSUPPORTED
	// The version of the supplied content information is not supported.

	PEERDIST_ERROR_CONTENTINFO_VERSION_UNSUPPORTED = 4050

	// PEERDIST_ERROR_CANNOT_PARSE_CONTENTINFO
	// The supplied content information is malformed.

	PEERDIST_ERROR_CANNOT_PARSE_CONTENTINFO = 4051

	// PEERDIST_ERROR_MISSING_DATA
	// The requested data cannot be found in local or peer caches.

	PEERDIST_ERROR_MISSING_DATA = 4052

	// PEERDIST_ERROR_NO_MORE
	// No more data is available or required.

	PEERDIST_ERROR_NO_MORE = 4053

	// PEERDIST_ERROR_NOT_INITIALIZED
	// The supplied object has not been initialized.

	PEERDIST_ERROR_NOT_INITIALIZED = 4054

	// PEERDIST_ERROR_ALREADY_INITIALIZED
	// The supplied object has already been initialized.

	PEERDIST_ERROR_ALREADY_INITIALIZED = 4055

	// PEERDIST_ERROR_SHUTDOWN_IN_PROGRESS
	// A shutdown operation is already in progress.

	PEERDIST_ERROR_SHUTDOWN_IN_PROGRESS = 4056

	// PEERDIST_ERROR_INVALIDATED
	// The supplied object has already been invalidated.

	PEERDIST_ERROR_INVALIDATED = 4057

	// PEERDIST_ERROR_ALREADY_EXISTS
	// An element already exists and was not replaced.

	PEERDIST_ERROR_ALREADY_EXISTS = 4058

	// PEERDIST_ERROR_OPERATION_NOTFOUND
	// Can not cancel the requested operation as it has already been completed.

	PEERDIST_ERROR_OPERATION_NOTFOUND = 4059

	// PEERDIST_ERROR_ALREADY_COMPLETED
	// Can not perform the requested operation because it has already been carried out.

	PEERDIST_ERROR_ALREADY_COMPLETED = 4060

	// PEERDIST_ERROR_OUT_OF_BOUNDS
	// An operation accessed data beyond the bounds of valid data.

	PEERDIST_ERROR_OUT_OF_BOUNDS = 4061

	// PEERDIST_ERROR_VERSION_UNSUPPORTED
	// The requested version is not supported.

	PEERDIST_ERROR_VERSION_UNSUPPORTED = 4062

	// PEERDIST_ERROR_INVALID_CONFIGURATION
	// A configuration value is invalid.

	PEERDIST_ERROR_INVALID_CONFIGURATION = 4063

	// PEERDIST_ERROR_NOT_LICENSED
	// The SKU is not licensed.

	PEERDIST_ERROR_NOT_LICENSED = 4064

	// PEERDIST_ERROR_SERVICE_UNAVAILABLE
	// PeerDist Service is still initializing and will be available shortly.

	PEERDIST_ERROR_SERVICE_UNAVAILABLE = 4065

	// PEERDIST_ERROR_TRUST_FAILURE
	// Communication with one or more computers will be temporarily blocked due to recent errors.

	PEERDIST_ERROR_TRUST_FAILURE = 4066

	///////////////////////////////////////////////////
	//                                               //
	//               DHCP Error codes                //
	//                                               //
	//                =4100 to=4149                  //
	///////////////////////////////////////////////////

	// ERROR_DHCP_ADDRESS_CONFLICT
	// The DHCP client has obtained an IP address that is already in use on the network. The local interface will be disabled until the DHCP client can obtain a new address.
	ERROR_DHCP_ADDRESS_CONFLICT = 4100

	///////////////////////////////////////////////////
	//                                               //
	//                  Available                    //
	//                                               //
	//                =4150 to=4199                  //
	///////////////////////////////////////////////////

	///////////////////////////////////////////////////
	//                                               //
	//               WMI Error codes                 //
	//                                               //
	//                =4200 to=4249                  //
	///////////////////////////////////////////////////

	// ERROR_WMI_GUID_NOT_FOUND
	// The GUID passed was not recognized as valid by a WMI data provider.
	ERROR_WMI_GUID_NOT_FOUND = 4200

	// ERROR_WMI_INSTANCE_NOT_FOUND
	// The instance name passed was not recognized as valid by a WMI data provider.
	ERROR_WMI_INSTANCE_NOT_FOUND = 4201

	// ERROR_WMI_ITEMID_NOT_FOUND
	// The data item ID passed was not recognized as valid by a WMI data provider.
	ERROR_WMI_ITEMID_NOT_FOUND = 4202

	// ERROR_WMI_TRY_AGAIN
	// The WMI request could not be completed and should be retried.
	ERROR_WMI_TRY_AGAIN = 4203

	// ERROR_WMI_DP_NOT_FOUND
	// The WMI data provider could not be located.
	ERROR_WMI_DP_NOT_FOUND = 4204

	// ERROR_WMI_UNRESOLVED_INSTANCE_REF
	// The WMI data provider references an instance set that has not been registered.
	ERROR_WMI_UNRESOLVED_INSTANCE_REF = 4205

	// ERROR_WMI_ALREADY_ENABLED
	// The WMI data block or event notification has already been enabled.
	ERROR_WMI_ALREADY_ENABLED = 4206

	// ERROR_WMI_GUID_DISCONNECTED
	// The WMI data block is no longer available.
	ERROR_WMI_GUID_DISCONNECTED = 4207

	// ERROR_WMI_SERVER_UNAVAILABLE
	// The WMI data service is not available.
	ERROR_WMI_SERVER_UNAVAILABLE = 4208

	// ERROR_WMI_DP_FAILED
	// The WMI data provider failed to carry out the request.
	ERROR_WMI_DP_FAILED = 4209

	// ERROR_WMI_INVALID_MOF
	// The WMI MOF information is not valid.
	ERROR_WMI_INVALID_MOF = 4210

	// ERROR_WMI_INVALID_REGINFO
	// The WMI registration information is not valid.
	ERROR_WMI_INVALID_REGINFO = 4211

	// ERROR_WMI_ALREADY_DISABLED
	// The WMI data block or event notification has already been disabled.
	ERROR_WMI_ALREADY_DISABLED = 4212

	// ERROR_WMI_READ_ONLY
	// The WMI data item or data block is read only.
	ERROR_WMI_READ_ONLY = 4213

	// ERROR_WMI_SET_FAILURE
	// The WMI data item or data block could not be changed.
	ERROR_WMI_SET_FAILURE = 4214

	///////////////////////////////////////////////////
	//                                               //
	//      app container Specific Error Codes        //
	//                                               //
	//                =4250 to=4299                  //
	///////////////////////////////////////////////////

	// ERROR_NOT_APPCONTAINER
	// This operation is only valid in the context of an app container.
	ERROR_NOT_APPCONTAINER = 4250

	// ERROR_APPCONTAINER_REQUIRED
	// This application can only run in the context of an app container.
	ERROR_APPCONTAINER_REQUIRED = 4251

	// ERROR_NOT_SUPPORTED_IN_APPCONTAINER
	// This functionality is not supported in the context of an app container.
	ERROR_NOT_SUPPORTED_IN_APPCONTAINER = 4252

	// ERROR_INVALID_PACKAGE_SID_LENGTH
	// The length of the SID supplied is not a valid length for app container SIDs.
	ERROR_INVALID_PACKAGE_SID_LENGTH = 4253

	///////////////////////////////////////////////////
	//                                               //
	//        RSM (Media Services) Error codes       //
	//                                               //
	//                =4300 to=4349                  //
	///////////////////////////////////////////////////

	// ERROR_INVALID_MEDIA
	// The media identifier does not represent a valid medium.
	ERROR_INVALID_MEDIA = 4300

	// ERROR_INVALID_LIBRARY
	// The library identifier does not represent a valid library.
	ERROR_INVALID_LIBRARY = 4301

	// ERROR_INVALID_MEDIA_POO
	// The media pool identifier does not represent a valid media pool.
	ERROR_INVALID_MEDIA_POOL = 4302

	// ERROR_DRIVE_MEDIA_MISMATCH
	// The drive and medium are not compatible or exist in different libraries.
	ERROR_DRIVE_MEDIA_MISMATCH = 4303

	// ERROR_MEDIA_OFFLINE
	// The medium currently exists in an offline library and must be online to perform this operation.
	ERROR_MEDIA_OFFLINE = 4304

	// ERROR_LIBRARY_OFFLINE
	// The operation cannot be performed on an offline library.
	ERROR_LIBRARY_OFFLINE = 4305

	// ERROR_EMPTY
	// The library, drive, or media pool is empty.
	ERROR_EMPTY = 4306

	// ERROR_NOT_EMPTY
	// The library, drive, or media pool must be empty to perform this operation.
	ERROR_NOT_EMPTY = 4307

	// ERROR_MEDIA_UNAVAILABLE
	// No media is currently available in this media pool or library.
	ERROR_MEDIA_UNAVAILABLE = 4308

	// ERROR_RESOURCE_DISABLED
	// A resource required for this operation is disabled.
	ERROR_RESOURCE_DISABLED = 4309

	// ERROR_INVALID_CLEANER
	// The media identifier does not represent a valid cleaner.
	ERROR_INVALID_CLEANER = 4310

	// ERROR_UNABLE_TO_CLEAN
	// The drive cannot be cleaned or does not support cleaning.
	ERROR_UNABLE_TO_CLEAN = 4311

	// ERROR_OBJECT_NOT_FOUND
	// The object identifier does not represent a valid object.
	ERROR_OBJECT_NOT_FOUND = 4312

	// ERROR_DATABASE_FAILURE
	// Unable to read from or write to the database.
	ERROR_DATABASE_FAILURE = 4313

	// ERROR_DATABASE_FUL
	// The database is full.
	ERROR_DATABASE_FULL = 4314

	// ERROR_MEDIA_INCOMPATIBLE
	// The medium is not compatible with the device or media pool.
	ERROR_MEDIA_INCOMPATIBLE = 4315

	// ERROR_RESOURCE_NOT_PRESENT
	// The resource required for this operation does not exist.
	ERROR_RESOURCE_NOT_PRESENT = 4316

	// ERROR_INVALID_OPERATION
	// The operation identifier is not valid.
	ERROR_INVALID_OPERATION = 4317

	// ERROR_MEDIA_NOT_AVAILABLE
	// The media is not mounted or ready for use.
	ERROR_MEDIA_NOT_AVAILABLE = 4318

	// ERROR_DEVICE_NOT_AVAILABLE
	// The device is not ready for use.
	ERROR_DEVICE_NOT_AVAILABLE = 4319

	// ERROR_REQUEST_REFUSED
	// The operator or administrator has refused the request.
	ERROR_REQUEST_REFUSED = 4320

	// ERROR_INVALID_DRIVE_OBJECT
	// The drive identifier does not represent a valid drive.
	ERROR_INVALID_DRIVE_OBJECT = 4321

	// ERROR_LIBRARY_FUL
	// Library is full. No slot is available for use.
	ERROR_LIBRARY_FULL = 4322

	// ERROR_MEDIUM_NOT_ACCESSIBLE
	// The transport cannot access the medium.
	ERROR_MEDIUM_NOT_ACCESSIBLE = 4323

	// ERROR_UNABLE_TO_LOAD_MEDIUM
	// Unable to load the medium into the drive.
	ERROR_UNABLE_TO_LOAD_MEDIUM = 4324

	// ERROR_UNABLE_TO_INVENTORY_DRIVE
	// Unable to retrieve the drive status.
	ERROR_UNABLE_TO_INVENTORY_DRIVE = 4325

	// ERROR_UNABLE_TO_INVENTORY_SLOT
	// Unable to retrieve the slot status.
	ERROR_UNABLE_TO_INVENTORY_SLOT = 4326

	// ERROR_UNABLE_TO_INVENTORY_TRANSPORT
	// Unable to retrieve status about the transport.
	ERROR_UNABLE_TO_INVENTORY_TRANSPORT = 4327

	// ERROR_TRANSPORT_FUL
	// Cannot use the transport because it is already in use.
	ERROR_TRANSPORT_FULL = 4328

	// ERROR_CONTROLLING_IEPORT
	// Unable to open or close the inject/eject port.
	ERROR_CONTROLLING_IEPORT = 4329

	// ERROR_UNABLE_TO_EJECT_MOUNTED_MEDIA
	// Unable to eject the medium because it is in a drive.
	ERROR_UNABLE_TO_EJECT_MOUNTED_MEDIA = 4330

	// ERROR_CLEANER_SLOT_SET
	// A cleaner slot is already reserved.
	ERROR_CLEANER_SLOT_SET = 4331

	// ERROR_CLEANER_SLOT_NOT_SET
	// A cleaner slot is not reserved.
	ERROR_CLEANER_SLOT_NOT_SET = 4332

	// ERROR_CLEANER_CARTRIDGE_SPENT
	// The cleaner cartridge has performed the maximum number of drive cleanings.
	ERROR_CLEANER_CARTRIDGE_SPENT = 4333

	// ERROR_UNEXPECTED_OMID
	// Unexpected on-medium identifier.
	ERROR_UNEXPECTED_OMID = 4334

	// ERROR_CANT_DELETE_LAST_ITEM
	// The last remaining item in this group or resource cannot be deleted.
	ERROR_CANT_DELETE_LAST_ITEM = 4335

	// ERROR_MESSAGE_EXCEEDS_MAX_SIZE
	// The message provided exceeds the maximum size allowed for this parameter.
	ERROR_MESSAGE_EXCEEDS_MAX_SIZE = 4336

	// ERROR_VOLUME_CONTAINS_SYS_FILES
	// The volume contains system or paging files.
	ERROR_VOLUME_CONTAINS_SYS_FILES = 4337

	// ERROR_INDIGENOUS_TYPE
	// The media type cannot be removed from this library since at least one drive in the library reports it can support this media type.
	ERROR_INDIGENOUS_TYPE = 4338

	// ERROR_NO_SUPPORTING_DRIVES
	// This offline media cannot be mounted on this system since no enabled drives are present which can be used.
	ERROR_NO_SUPPORTING_DRIVES = 4339

	// ERROR_CLEANER_CARTRIDGE_INSTALLED
	// A cleaner cartridge is present in the tape library.
	ERROR_CLEANER_CARTRIDGE_INSTALLED = 4340

	// ERROR_IEPORT_FUL
	// Cannot use the inject/eject port because it is not empty.
	ERROR_IEPORT_FULL = 4341

	///////////////////////////////////////////////////
	//                                               //
	//       Remote Storage Service Error codes      //
	//                                               //
	//                =4350 to=4389                  //
	///////////////////////////////////////////////////

	// ERROR_FILE_OFFLINE
	// This file is currently not available for use on this computer.
	ERROR_FILE_OFFLINE = 4350

	// ERROR_REMOTE_STORAGE_NOT_ACTIVE
	// The remote storage service is not operational at this time.
	ERROR_REMOTE_STORAGE_NOT_ACTIVE = 4351

	// ERROR_REMOTE_STORAGE_MEDIA_ERROR
	// The remote storage service encountered a media error.
	ERROR_REMOTE_STORAGE_MEDIA_ERROR = 4352

	///////////////////////////////////////////////////
	//                                               //
	//           Reparse Point Error codes           //
	//                                               //
	//                =4390 to=4399                  //
	///////////////////////////////////////////////////

	// ERROR_NOT_A_REPARSE_POINT
	// The file or directory is not a reparse point.
	ERROR_NOT_A_REPARSE_POINT = 4390

	// ERROR_REPARSE_ATTRIBUTE_CONFLICT
	// The reparse point attribute cannot be set because it conflicts with an existing attribute.
	ERROR_REPARSE_ATTRIBUTE_CONFLICT = 4391

	// ERROR_INVALID_REPARSE_DATA
	// The data present in the reparse point buffer is invalid.
	ERROR_INVALID_REPARSE_DATA = 4392

	// ERROR_REPARSE_TAG_INVALID
	// The tag present in the reparse point buffer is invalid.
	ERROR_REPARSE_TAG_INVALID = 4393

	// ERROR_REPARSE_TAG_MISMATCH
	// There is a mismatch between the tag specified in the request and the tag present in the reparse point.
	ERROR_REPARSE_TAG_MISMATCH = 4394

	// ERROR_REPARSE_POINT_ENCOUNTERED
	// The object manager encountered a reparse point while retrieving an object.
	ERROR_REPARSE_POINT_ENCOUNTERED = 4395

	///////////////////////////////////////////////////
	//                                               //
	//         Fast Cache Specific Error Codes       //
	//                                               //
	//                =4400 to=4419                  //
	///////////////////////////////////////////////////

	// ERROR_APP_DATA_NOT_FOUND
	// Fast Cache data not found.
	ERROR_APP_DATA_NOT_FOUND = 4400

	// ERROR_APP_DATA_EXPIRED
	// Fast Cache data expired.
	ERROR_APP_DATA_EXPIRED = 4401

	// ERROR_APP_DATA_CORRUPT
	// Fast Cache data corrupt.
	ERROR_APP_DATA_CORRUPT = 4402

	// ERROR_APP_DATA_LIMIT_EXCEEDED
	// Fast Cache data has exceeded its max size and cannot be updated.
	ERROR_APP_DATA_LIMIT_EXCEEDED = 4403

	// ERROR_APP_DATA_REBOOT_REQUIRED
	// Fast Cache has been ReArmed and requires a reboot until it can be updated.
	ERROR_APP_DATA_REBOOT_REQUIRED = 4404

	///////////////////////////////////////////////////
	//                                               //
	//             SecureBoot Error codes            //
	//                                               //
	//                =4420 to=4439                  //
	///////////////////////////////////////////////////

	// ERROR_SECUREBOOT_ROLLBACK_DETECTED
	// Secure Boot detected that rollback of protected data has been attempted.
	ERROR_SECUREBOOT_ROLLBACK_DETECTED = 4420

	// ERROR_SECUREBOOT_POLICY_VIOLATION
	// The value is protected by Secure Boot policy and cannot be modified or deleted.
	ERROR_SECUREBOOT_POLICY_VIOLATION = 4421

	// ERROR_SECUREBOOT_INVALID_POLICY
	// The Secure Boot policy is invalid.
	ERROR_SECUREBOOT_INVALID_POLICY = 4422

	// ERROR_SECUREBOOT_POLICY_PUBLISHER_NOT_FOUND
	// A new Secure Boot policy did not contain the current publisher on its update list.
	ERROR_SECUREBOOT_POLICY_PUBLISHER_NOT_FOUND = 4423

	// ERROR_SECUREBOOT_POLICY_NOT_SIGNED
	// The Secure Boot policy is either not signed or is signed by a non-trusted signer.
	ERROR_SECUREBOOT_POLICY_NOT_SIGNED = 4424

	// ERROR_SECUREBOOT_NOT_ENABLED
	// Secure Boot is not enabled on this machine.
	ERROR_SECUREBOOT_NOT_ENABLED = 4425

	// ERROR_SECUREBOOT_FILE_REPLACED
	// Secure Boot requires that certain files and drivers are not replaced by other files or drivers.
	ERROR_SECUREBOOT_FILE_REPLACED = 4426

	// ERROR_SECUREBOOT_POLICY_NOT_AUTHORIZED
	// The Secure Boot Supplemental Policy file was not authorized on this machine.
	ERROR_SECUREBOOT_POLICY_NOT_AUTHORIZED = 4427

	// ERROR_SECUREBOOT_POLICY_UNKNOWN
	// The Supplemental Policy is not recognized on this device.
	ERROR_SECUREBOOT_POLICY_UNKNOWN = 4428

	// ERROR_SECUREBOOT_POLICY_MISSING_ANTIROLLBACKVERSION
	// The Antirollback version was not found in the Secure Boot Policy.
	ERROR_SECUREBOOT_POLICY_MISSING_ANTIROLLBACKVERSION = 4429

	// ERROR_SECUREBOOT_PLATFORM_ID_MISMATCH
	// The Platform ID specified in the Secure Boot policy does not match the Platform ID on this device.
	ERROR_SECUREBOOT_PLATFORM_ID_MISMATCH = 4430

	// ERROR_SECUREBOOT_POLICY_ROLLBACK_DETECTED
	// The Secure Boot policy file has an older Antirollback Version than this device.
	ERROR_SECUREBOOT_POLICY_ROLLBACK_DETECTED = 4431

	// ERROR_SECUREBOOT_POLICY_UPGRADE_MISMATCH
	// The Secure Boot policy file does not match the upgraded legacy policy.
	ERROR_SECUREBOOT_POLICY_UPGRADE_MISMATCH = 4432

	// ERROR_SECUREBOOT_REQUIRED_POLICY_FILE_MISSING
	// The Secure Boot policy file is required but could not be found.
	ERROR_SECUREBOOT_REQUIRED_POLICY_FILE_MISSING = 4433

	// ERROR_SECUREBOOT_NOT_BASE_POLICY
	// Supplemental Secure Boot policy file can not be loaded as a base Secure Boot policy.
	ERROR_SECUREBOOT_NOT_BASE_POLICY = 4434

	// ERROR_SECUREBOOT_NOT_SUPPLEMENTAL_POLICY
	// Base Secure Boot policy file can not be loaded as a Supplemental Secure Boot policy.
	ERROR_SECUREBOOT_NOT_SUPPLEMENTAL_POLICY = 4435

	///////////////////////////////////////////////////
	//                                               //
	//       File System Specific Error Codes        //
	//                                               //
	//                =4440 to=4499                  //
	///////////////////////////////////////////////////

	// ERROR_OFFLOAD_READ_FLT_NOT_SUPPORTED
	// The copy offload read operation is not supported by a filter.
	ERROR_OFFLOAD_READ_FLT_NOT_SUPPORTED = 4440

	// ERROR_OFFLOAD_WRITE_FLT_NOT_SUPPORTED
	// The copy offload write operation is not supported by a filter.
	ERROR_OFFLOAD_WRITE_FLT_NOT_SUPPORTED = 4441

	// ERROR_OFFLOAD_READ_FILE_NOT_SUPPORTED
	// The copy offload read operation is not supported for the file.
	ERROR_OFFLOAD_READ_FILE_NOT_SUPPORTED = 4442

	// ERROR_OFFLOAD_WRITE_FILE_NOT_SUPPORTED
	// The copy offload write operation is not supported for the file.
	ERROR_OFFLOAD_WRITE_FILE_NOT_SUPPORTED = 4443

	// ERROR_ALREADY_HAS_STREAM_ID
	// This file is currently associated with a different stream id.
	ERROR_ALREADY_HAS_STREAM_ID = 4444

	// ERROR_SMR_GARBAGE_COLLECTION_REQUIRED
	// The volume must undergo garbage collection.
	ERROR_SMR_GARBAGE_COLLECTION_REQUIRED = 4445

	// ERROR_WOF_WIM_HEADER_CORRUPT
	// The WOF driver encountered a corruption in WIM image's Header.
	ERROR_WOF_WIM_HEADER_CORRUPT = 4446

	// ERROR_WOF_WIM_RESOURCE_TABLE_CORRUPT
	// The WOF driver encountered a corruption in WIM image's Resource Table.
	ERROR_WOF_WIM_RESOURCE_TABLE_CORRUPT = 4447

	// ERROR_WOF_FILE_RESOURCE_TABLE_CORRUPT
	// The WOF driver encountered a corruption in the compressed file's Resource Table.
	ERROR_WOF_FILE_RESOURCE_TABLE_CORRUPT = 4448

	///////////////////////////////////////////////////
	//                                               //
	//    Single Instance Store (SIS) Error codes    //
	//                                               //
	//                =4500 to=4549                  //
	///////////////////////////////////////////////////

	// ERROR_VOLUME_NOT_SIS_ENABLED
	// Single Instance Storage is not available on this volume.
	ERROR_VOLUME_NOT_SIS_ENABLED = 4500

	///////////////////////////////////////////////////
	//                                               //
	//             System Integrity Error codes      //
	//                                               //
	//                =4550 to=4559                  //
	///////////////////////////////////////////////////

	// ERROR_SYSTEM_INTEGRITY_ROLLBACK_DETECTED
	// System Integrity detected that policy rollback has been attempted.
	ERROR_SYSTEM_INTEGRITY_ROLLBACK_DETECTED = 4550

	// ERROR_SYSTEM_INTEGRITY_POLICY_VIOLATION
	// Your organization used Device Guard to block this app. Contact your support person for more info.
	ERROR_SYSTEM_INTEGRITY_POLICY_VIOLATION = 4551

	// ERROR_SYSTEM_INTEGRITY_INVALID_POLICY
	// The System Integrity policy is invalid.
	ERROR_SYSTEM_INTEGRITY_INVALID_POLICY = 4552

	// ERROR_SYSTEM_INTEGRITY_POLICY_NOT_SIGNED
	// The System Integrity policy is either not signed or is signed by a non-trusted signer.
	ERROR_SYSTEM_INTEGRITY_POLICY_NOT_SIGNED = 4553

	///////////////////////////////////////////////////
	//                                               //
	//             VSM Error codes                   //
	//                                               //
	//                =4560 to=4569                  //
	///////////////////////////////////////////////////

	// ERROR_VSM_NOT_INITIALIZED
	// Virtual Secure Mode (VSM) is not initialized. The hypervisor or VSM may not be present or enabled.
	ERROR_VSM_NOT_INITIALIZED = 4560

	// ERROR_VSM_DMA_PROTECTION_NOT_IN_USE
	// The hypervisor is not protecting DMA because an IOMMU is not present or not enabled in the BIOS.
	ERROR_VSM_DMA_PROTECTION_NOT_IN_USE = 4561

	///////////////////////////////////////////////////
	//                                               //
	//         Platform Manifest Error Codes         //
	//                                               //
	//                =4570 to=4579                  //
	///////////////////////////////////////////////////

	// ERROR_PLATFORM_MANIFEST_NOT_AUTHORIZED
	// The Platform Manifest file was not authorized on this machine.
	ERROR_PLATFORM_MANIFEST_NOT_AUTHORIZED = 4570

	// ERROR_PLATFORM_MANIFEST_INVALID
	// The Platform Manifest file was not valid.
	ERROR_PLATFORM_MANIFEST_INVALID = 4571

	// ERROR_PLATFORM_MANIFEST_FILE_NOT_AUTHORIZED
	// The file is not authorized on this platform because an entry was not found in the Platform Manifest.
	ERROR_PLATFORM_MANIFEST_FILE_NOT_AUTHORIZED = 4572

	// ERROR_PLATFORM_MANIFEST_CATALOG_NOT_AUTHORIZED
	// The catalog is not authorized on this platform because an entry was not found in the Platform Manifest.
	ERROR_PLATFORM_MANIFEST_CATALOG_NOT_AUTHORIZED = 4573

	// ERROR_PLATFORM_MANIFEST_BINARY_ID_NOT_FOUND
	// The file is not authorized on this platform because a Binary ID was not found in the embedded signature.
	ERROR_PLATFORM_MANIFEST_BINARY_ID_NOT_FOUND = 4574

	// ERROR_PLATFORM_MANIFEST_NOT_ACTIVE
	// No active Platform Manifest exists on this system.
	ERROR_PLATFORM_MANIFEST_NOT_ACTIVE = 4575

	// ERROR_PLATFORM_MANIFEST_NOT_SIGNED
	// The Platform Manifest file was not properly signed.
	ERROR_PLATFORM_MANIFEST_NOT_SIGNED = 4576

	///////////////////////////////////////////////////
	//                                               //
	//                  Available                    //
	//                                               //
	//                =4580 to=4599                  //
	///////////////////////////////////////////////////

	///////////////////////////////////////////////////
	//                                               //
	//             Cluster Error codes               //
	//                                               //
	//                =5000 to=5999                  //
	///////////////////////////////////////////////////

	// ERROR_DEPENDENT_RESOURCE_EXISTS
	// The operation cannot be completed because other resources are dependent on this resource.
	ERROR_DEPENDENT_RESOURCE_EXISTS = 5001

	// ERROR_DEPENDENCY_NOT_FOUND
	// The cluster resource dependency cannot be found.
	ERROR_DEPENDENCY_NOT_FOUND = 5002

	// ERROR_DEPENDENCY_ALREADY_EXISTS
	// The cluster resource cannot be made dependent on the specified resource because it is already dependent.
	ERROR_DEPENDENCY_ALREADY_EXISTS = 5003

	// ERROR_RESOURCE_NOT_ONLINE
	// The cluster resource is not online.
	ERROR_RESOURCE_NOT_ONLINE = 5004

	// ERROR_HOST_NODE_NOT_AVAILABLE
	// A cluster node is not available for this operation.
	ERROR_HOST_NODE_NOT_AVAILABLE = 5005

	// ERROR_RESOURCE_NOT_AVAILABLE
	// The cluster resource is not available.
	ERROR_RESOURCE_NOT_AVAILABLE = 5006

	// ERROR_RESOURCE_NOT_FOUND
	// The cluster resource could not be found.
	ERROR_RESOURCE_NOT_FOUND = 5007

	// ERROR_SHUTDOWN_CLUSTER
	// The cluster is being shut down.
	ERROR_SHUTDOWN_CLUSTER = 5008

	// ERROR_CANT_EVICT_ACTIVE_NODE
	// A cluster node cannot be evicted from the cluster unless the node is down or it is the last node.
	ERROR_CANT_EVICT_ACTIVE_NODE = 5009

	// ERROR_OBJECT_ALREADY_EXISTS
	// The object already exists.
	ERROR_OBJECT_ALREADY_EXISTS = 5010

	// ERROR_OBJECT_IN_LIST
	// The object is already in the list.
	ERROR_OBJECT_IN_LIST = 5011

	// ERROR_GROUP_NOT_AVAILABLE
	// The cluster group is not available for any new requests.
	ERROR_GROUP_NOT_AVAILABLE = 5012

	// ERROR_GROUP_NOT_FOUND
	// The cluster group could not be found.
	ERROR_GROUP_NOT_FOUND = 5013

	// ERROR_GROUP_NOT_ONLINE
	// The operation could not be completed because the cluster group is not online.
	ERROR_GROUP_NOT_ONLINE = 5014

	// ERROR_HOST_NODE_NOT_RESOURCE_OWNER
	// The operation failed because either the specified cluster node is not the owner of the resource, or the node is not a possible owner of the resource.
	ERROR_HOST_NODE_NOT_RESOURCE_OWNER = 5015

	// ERROR_HOST_NODE_NOT_GROUP_OWNER
	// The operation failed because either the specified cluster node is not the owner of the group, or the node is not a possible owner of the group.
	ERROR_HOST_NODE_NOT_GROUP_OWNER = 5016

	// ERROR_RESMON_CREATE_FAILED
	// The cluster resource could not be created in the specified resource monitor.
	ERROR_RESMON_CREATE_FAILED = 5017

	// ERROR_RESMON_ONLINE_FAILED
	// The cluster resource could not be brought online by the resource monitor.
	ERROR_RESMON_ONLINE_FAILED = 5018

	// ERROR_RESOURCE_ONLINE
	// The operation could not be completed because the cluster resource is online.
	ERROR_RESOURCE_ONLINE = 5019

	// ERROR_QUORUM_RESOURCE
	// The cluster resource could not be deleted or brought offline because it is the quorum resource.
	ERROR_QUORUM_RESOURCE = 5020

	// ERROR_NOT_QUORUM_CAPABLE
	// The cluster could not make the specified resource a quorum resource because it is not capable of being a quorum resource.
	ERROR_NOT_QUORUM_CAPABLE = 5021

	// ERROR_CLUSTER_SHUTTING_DOWN
	// The cluster software is shutting down.
	ERROR_CLUSTER_SHUTTING_DOWN = 5022

	// ERROR_INVALID_STATE
	// The group or resource is not in the correct state to perform the requested operation.
	ERROR_INVALID_STATE = 5023

	// ERROR_RESOURCE_PROPERTIES_STORED
	// The properties were stored but not all changes will take effect until the next time the resource is brought online.
	ERROR_RESOURCE_PROPERTIES_STORED = 5024

	// ERROR_NOT_QUORUM_CLASS
	// The cluster could not make the specified resource a quorum resource because it does not belong to a shared storage class.
	ERROR_NOT_QUORUM_CLASS = 5025

	// ERROR_CORE_RESOURCE
	// The cluster resource could not be deleted since it is a core resource.
	ERROR_CORE_RESOURCE = 5026

	// ERROR_QUORUM_RESOURCE_ONLINE_FAILED
	// The quorum resource failed to come online.
	ERROR_QUORUM_RESOURCE_ONLINE_FAILED = 5027

	// ERROR_QUORUMLOG_OPEN_FAILED
	// The quorum log could not be created or mounted successfully.
	ERROR_QUORUMLOG_OPEN_FAILED = 5028

	// ERROR_CLUSTERLOG_CORRUPT
	// The cluster log is corrupt.
	ERROR_CLUSTERLOG_CORRUPT = 5029

	// ERROR_CLUSTERLOG_RECORD_EXCEEDS_MAXSIZE
	// The record could not be written to the cluster log since it exceeds the maximum size.
	ERROR_CLUSTERLOG_RECORD_EXCEEDS_MAXSIZE = 5030

	// ERROR_CLUSTERLOG_EXCEEDS_MAXSIZE
	// The cluster log exceeds its maximum size.
	ERROR_CLUSTERLOG_EXCEEDS_MAXSIZE = 5031

	// ERROR_CLUSTERLOG_CHKPOINT_NOT_FOUND
	// No checkpoint record was found in the cluster log.
	ERROR_CLUSTERLOG_CHKPOINT_NOT_FOUND = 5032

	// ERROR_CLUSTERLOG_NOT_ENOUGH_SPACE
	// The minimum required disk space needed for logging is not available.
	ERROR_CLUSTERLOG_NOT_ENOUGH_SPACE = 5033

	// ERROR_QUORUM_OWNER_ALIVE
	// The cluster node failed to take control of the quorum resource because the resource is owned by another active node.
	ERROR_QUORUM_OWNER_ALIVE = 5034

	// ERROR_NETWORK_NOT_AVAILABLE
	// A cluster network is not available for this operation.
	ERROR_NETWORK_NOT_AVAILABLE = 5035

	// ERROR_NODE_NOT_AVAILABLE
	// A cluster node is not available for this operation.
	ERROR_NODE_NOT_AVAILABLE = 5036

	// ERROR_ALL_NODES_NOT_AVAILABLE
	// All cluster nodes must be running to perform this operation.
	ERROR_ALL_NODES_NOT_AVAILABLE = 5037

	// ERROR_RESOURCE_FAILED
	// A cluster resource failed.
	ERROR_RESOURCE_FAILED = 5038

	// ERROR_CLUSTER_INVALID_NODE
	// The cluster node is not valid.
	ERROR_CLUSTER_INVALID_NODE = 5039

	// ERROR_CLUSTER_NODE_EXISTS
	// The cluster node already exists.
	ERROR_CLUSTER_NODE_EXISTS = 5040

	// ERROR_CLUSTER_JOIN_IN_PROGRESS
	// A node is in the process of joining the cluster.
	ERROR_CLUSTER_JOIN_IN_PROGRESS = 5041

	// ERROR_CLUSTER_NODE_NOT_FOUND
	// The cluster node was not found.
	ERROR_CLUSTER_NODE_NOT_FOUND = 5042

	// ERROR_CLUSTER_LOCAL_NODE_NOT_FOUND
	// The cluster local node information was not found.
	ERROR_CLUSTER_LOCAL_NODE_NOT_FOUND = 5043

	// ERROR_CLUSTER_NETWORK_EXISTS
	// The cluster network already exists.
	ERROR_CLUSTER_NETWORK_EXISTS = 5044

	// ERROR_CLUSTER_NETWORK_NOT_FOUND
	// The cluster network was not found.
	ERROR_CLUSTER_NETWORK_NOT_FOUND = 5045

	// ERROR_CLUSTER_NETINTERFACE_EXISTS
	// The cluster network interface already exists.
	ERROR_CLUSTER_NETINTERFACE_EXISTS = 5046

	// ERROR_CLUSTER_NETINTERFACE_NOT_FOUND
	// The cluster network interface was not found.
	ERROR_CLUSTER_NETINTERFACE_NOT_FOUND = 5047

	// ERROR_CLUSTER_INVALID_REQUEST
	// The cluster request is not valid for this object.
	ERROR_CLUSTER_INVALID_REQUEST = 5048

	// ERROR_CLUSTER_INVALID_NETWORK_PROVIDER
	// The cluster network provider is not valid.
	ERROR_CLUSTER_INVALID_NETWORK_PROVIDER = 5049

	// ERROR_CLUSTER_NODE_DOWN
	// The cluster node is down.
	ERROR_CLUSTER_NODE_DOWN = 5050

	// ERROR_CLUSTER_NODE_UNREACHABLE
	// The cluster node is not reachable.
	ERROR_CLUSTER_NODE_UNREACHABLE = 5051

	// ERROR_CLUSTER_NODE_NOT_MEMBER
	// The cluster node is not a member of the cluster.
	ERROR_CLUSTER_NODE_NOT_MEMBER = 5052

	// ERROR_CLUSTER_JOIN_NOT_IN_PROGRESS
	// A cluster join operation is not in progress.
	ERROR_CLUSTER_JOIN_NOT_IN_PROGRESS = 5053

	// ERROR_CLUSTER_INVALID_NETWORK
	// The cluster network is not valid.
	ERROR_CLUSTER_INVALID_NETWORK = 5054

	// ERROR_CLUSTER_NODE_UP
	// The cluster node is up.
	ERROR_CLUSTER_NODE_UP = 5056

	// ERROR_CLUSTER_IPADDR_IN_USE
	// The cluster IP address is already in use.
	ERROR_CLUSTER_IPADDR_IN_USE = 5057

	// ERROR_CLUSTER_NODE_NOT_PAUSED
	// The cluster node is not paused.
	ERROR_CLUSTER_NODE_NOT_PAUSED = 5058

	// ERROR_CLUSTER_NO_SECURITY_CONTEXT
	// No cluster security context is available.
	ERROR_CLUSTER_NO_SECURITY_CONTEXT = 5059

	// ERROR_CLUSTER_NETWORK_NOT_INTERNA
	// The cluster network is not configured for internal cluster communication.
	ERROR_CLUSTER_NETWORK_NOT_INTERNAL = 5060

	// ERROR_CLUSTER_NODE_ALREADY_UP
	// The cluster node is already up.
	ERROR_CLUSTER_NODE_ALREADY_UP = 5061

	// ERROR_CLUSTER_NODE_ALREADY_DOWN
	// The cluster node is already down.
	ERROR_CLUSTER_NODE_ALREADY_DOWN = 5062

	// ERROR_CLUSTER_NETWORK_ALREADY_ONLINE
	// The cluster network is already online.
	ERROR_CLUSTER_NETWORK_ALREADY_ONLINE = 5063

	// ERROR_CLUSTER_NETWORK_ALREADY_OFFLINE
	// The cluster network is already offline.
	ERROR_CLUSTER_NETWORK_ALREADY_OFFLINE = 5064

	// ERROR_CLUSTER_NODE_ALREADY_MEMBER
	// The cluster node is already a member of the cluster.
	ERROR_CLUSTER_NODE_ALREADY_MEMBER = 5065

	// ERROR_CLUSTER_LAST_INTERNAL_NETWORK
	// The cluster network is the only one configured for internal cluster communication between two or more active cluster nodes. The internal communication capability cannot be removed from the network.
	ERROR_CLUSTER_LAST_INTERNAL_NETWORK = 5066

	// ERROR_CLUSTER_NETWORK_HAS_DEPENDENTS
	// One or more cluster resources depend on the network to provide service to clients. The client access capability cannot be removed from the network.
	ERROR_CLUSTER_NETWORK_HAS_DEPENDENTS = 5067

	// ERROR_INVALID_OPERATION_ON_QUORUM
	// This operation cannot currently be performed on the cluster group containing the quorum resource.
	ERROR_INVALID_OPERATION_ON_QUORUM = 5068

	// ERROR_DEPENDENCY_NOT_ALLOWED
	// The cluster quorum resource is not allowed to have any dependencies.
	ERROR_DEPENDENCY_NOT_ALLOWED = 5069

	// ERROR_CLUSTER_NODE_PAUSED
	// The cluster node is paused.
	ERROR_CLUSTER_NODE_PAUSED = 5070

	// ERROR_NODE_CANT_HOST_RESOURCE
	// The cluster resource cannot be brought online. The owner node cannot run this resource.
	ERROR_NODE_CANT_HOST_RESOURCE = 5071

	// ERROR_CLUSTER_NODE_NOT_READY
	// The cluster node is not ready to perform the requested operation.
	ERROR_CLUSTER_NODE_NOT_READY = 5072

	// ERROR_CLUSTER_NODE_SHUTTING_DOWN
	// The cluster node is shutting down.
	ERROR_CLUSTER_NODE_SHUTTING_DOWN = 5073

	// ERROR_CLUSTER_JOIN_ABORTED
	// The cluster join operation was aborted.
	ERROR_CLUSTER_JOIN_ABORTED = 5074

	// ERROR_CLUSTER_INCOMPATIBLE_VERSIONS
	// The node failed to join the cluster because the joining node and other nodes in the cluster have incompatible operating system versions. To get more information about operating system versions of the cluster, run the Validate a Configuration Wizard or the Test-Cluster Windows PowerShell cmdlet.
	ERROR_CLUSTER_INCOMPATIBLE_VERSIONS = 5075

	// ERROR_CLUSTER_MAXNUM_OF_RESOURCES_EXCEEDED
	// This resource cannot be created because the cluster has reached the limit on the number of resources it can monitor.
	ERROR_CLUSTER_MAXNUM_OF_RESOURCES_EXCEEDED = 5076

	// ERROR_CLUSTER_SYSTEM_CONFIG_CHANGED
	// The system configuration changed during the cluster join or form operation. The join or form operation was aborted.
	ERROR_CLUSTER_SYSTEM_CONFIG_CHANGED = 5077

	// ERROR_CLUSTER_RESOURCE_TYPE_NOT_FOUND
	// The specified resource type was not found.
	ERROR_CLUSTER_RESOURCE_TYPE_NOT_FOUND = 5078

	// ERROR_CLUSTER_RESTYPE_NOT_SUPPORTED
	// The specified node does not support a resource of this type. This may be due to version inconsistencies or due to the absence of the resource DLL on this node.
	ERROR_CLUSTER_RESTYPE_NOT_SUPPORTED = 5079

	// ERROR_CLUSTER_RESNAME_NOT_FOUND
	// The specified resource name is not supported by this resource DLL. This may be due to a bad (or changed) name supplied to the resource DLL.
	ERROR_CLUSTER_RESNAME_NOT_FOUND = 5080

	// ERROR_CLUSTER_NO_RPC_PACKAGES_REGISTERED
	// No authentication package could be registered with the RPC server.
	ERROR_CLUSTER_NO_RPC_PACKAGES_REGISTERED = 5081

	// ERROR_CLUSTER_OWNER_NOT_IN_PREFLIST
	// You cannot bring the group online because the owner of the group is not in the preferred list for the group. To change the owner node for the group, move the group.
	ERROR_CLUSTER_OWNER_NOT_IN_PREFLIST = 5082

	// ERROR_CLUSTER_DATABASE_SEQMISMATCH
	// The join operation failed because the cluster database sequence number has changed or is incompatible with the locker node. This may happen during a join operation if the cluster database was changing during the join.
	ERROR_CLUSTER_DATABASE_SEQMISMATCH = 5083

	// ERROR_RESMON_INVALID_STATE
	// The resource monitor will not allow the fail operation to be performed while the resource is in its current state. This may happen if the resource is in a pending state.
	ERROR_RESMON_INVALID_STATE = 5084

	// ERROR_CLUSTER_GUM_NOT_LOCKER
	// A non locker code got a request to reserve the lock for making global updates.
	ERROR_CLUSTER_GUM_NOT_LOCKER = 5085

	// ERROR_QUORUM_DISK_NOT_FOUND
	// The quorum disk could not be located by the cluster service.
	ERROR_QUORUM_DISK_NOT_FOUND = 5086

	// ERROR_DATABASE_BACKUP_CORRUPT
	// The backed up cluster database is possibly corrupt.
	ERROR_DATABASE_BACKUP_CORRUPT = 5087

	// ERROR_CLUSTER_NODE_ALREADY_HAS_DFS_ROOT
	// A DFS root already exists in this cluster node.
	ERROR_CLUSTER_NODE_ALREADY_HAS_DFS_ROOT = 5088

	// ERROR_RESOURCE_PROPERTY_UNCHANGEABLE
	// An attempt to modify a resource property failed because it conflicts with another existing property.
	ERROR_RESOURCE_PROPERTY_UNCHANGEABLE = 5089

	// ERROR_NO_ADMIN_ACCESS_POINT
	// This operation is not supported on a cluster without an Administrator Access Point.
	ERROR_NO_ADMIN_ACCESS_POINT = 5090

	/*
	   Codes from=4300 through=5889 overlap with codes in ds\published\inc\apperr2.w.
	   Do not add any more error codes in that range.
	*/

	// ERROR_CLUSTER_MEMBERSHIP_INVALID_STATE
	// An operation was attempted that is incompatible with the current membership state of the node.
	ERROR_CLUSTER_MEMBERSHIP_INVALID_STATE = 5890

	// ERROR_CLUSTER_QUORUMLOG_NOT_FOUND
	// The quorum resource does not contain the quorum log.
	ERROR_CLUSTER_QUORUMLOG_NOT_FOUND = 5891

	// ERROR_CLUSTER_MEMBERSHIP_HALT
	// The membership engine requested shutdown of the cluster service on this node.
	ERROR_CLUSTER_MEMBERSHIP_HALT = 5892

	// ERROR_CLUSTER_INSTANCE_ID_MISMATCH
	// The join operation failed because the cluster instance ID of the joining node does not match the cluster instance ID of the sponsor node.
	ERROR_CLUSTER_INSTANCE_ID_MISMATCH = 5893

	// ERROR_CLUSTER_NETWORK_NOT_FOUND_FOR_IP
	// A matching cluster network for the specified IP address could not be found.
	ERROR_CLUSTER_NETWORK_NOT_FOUND_FOR_IP = 5894

	// ERROR_CLUSTER_PROPERTY_DATA_TYPE_MISMATCH
	// The actual data type of the property did not match the expected data type of the property.
	ERROR_CLUSTER_PROPERTY_DATA_TYPE_MISMATCH = 5895

	// ERROR_CLUSTER_EVICT_WITHOUT_CLEANUP
	// The cluster node was evicted from the cluster successfully, but the node was not cleaned up. To determine what cleanup steps failed and how to recover, see the Failover Clustering application event log using Event Viewer.
	ERROR_CLUSTER_EVICT_WITHOUT_CLEANUP = 5896

	// ERROR_CLUSTER_PARAMETER_MISMATCH
	// Two or more parameter values specified for a resource's properties are in conflict.
	ERROR_CLUSTER_PARAMETER_MISMATCH = 5897

	// ERROR_NODE_CANNOT_BE_CLUSTERED
	// This computer cannot be made a member of a cluster.
	ERROR_NODE_CANNOT_BE_CLUSTERED = 5898

	// ERROR_CLUSTER_WRONG_OS_VERSION
	// This computer cannot be made a member of a cluster because it does not have the correct version of Windows installed.
	ERROR_CLUSTER_WRONG_OS_VERSION = 5899

	// ERROR_CLUSTER_CANT_CREATE_DUP_CLUSTER_NAME
	// A cluster cannot be created with the specified cluster name because that cluster name is already in use. Specify a different name for the cluster.
	ERROR_CLUSTER_CANT_CREATE_DUP_CLUSTER_NAME = 5900

	// ERROR_CLUSCFG_ALREADY_COMMITTED
	// The cluster configuration action has already been committed.
	ERROR_CLUSCFG_ALREADY_COMMITTED = 5901

	// ERROR_CLUSCFG_ROLLBACK_FAILED
	// The cluster configuration action could not be rolled back.
	ERROR_CLUSCFG_ROLLBACK_FAILED = 5902

	// ERROR_CLUSCFG_SYSTEM_DISK_DRIVE_LETTER_CONFLICT
	// The drive letter assigned to a system disk on one node conflicted with the drive letter assigned to a disk on another node.
	ERROR_CLUSCFG_SYSTEM_DISK_DRIVE_LETTER_CONFLICT = 5903

	// ERROR_CLUSTER_OLD_VERSION
	// One or more nodes in the cluster are running a version of Windows that does not support this operation.
	ERROR_CLUSTER_OLD_VERSION = 5904

	// ERROR_CLUSTER_MISMATCHED_COMPUTER_ACCT_NAME
	// The name of the corresponding computer account doesn't match the Network Name for this resource.
	ERROR_CLUSTER_MISMATCHED_COMPUTER_ACCT_NAME = 5905

	// ERROR_CLUSTER_NO_NET_ADAPTERS
	// No network adapters are available.
	ERROR_CLUSTER_NO_NET_ADAPTERS = 5906

	// ERROR_CLUSTER_POISONED
	// The cluster node has been poisoned.
	ERROR_CLUSTER_POISONED = 5907

	// ERROR_CLUSTER_GROUP_MOVING
	// The group is unable to accept the request since it is moving to another node.
	ERROR_CLUSTER_GROUP_MOVING = 5908

	// ERROR_CLUSTER_RESOURCE_TYPE_BUSY
	// The resource type cannot accept the request since is too busy performing another operation.
	ERROR_CLUSTER_RESOURCE_TYPE_BUSY = 5909

	// ERROR_RESOURCE_CALL_TIMED_OUT
	// The call to the cluster resource DLL timed out.
	ERROR_RESOURCE_CALL_TIMED_OUT = 5910

	// ERROR_INVALID_CLUSTER_IPV6_ADDRESS
	// The address is not valid for an IPv6 Address resource. A global IPv6 address is required, and it must match a cluster network. Compatibility addresses are not permitted.
	ERROR_INVALID_CLUSTER_IPV6_ADDRESS = 5911

	// ERROR_CLUSTER_INTERNAL_INVALID_FUNCTION
	// An internal cluster error occurred. A call to an invalid function was attempted.
	ERROR_CLUSTER_INTERNAL_INVALID_FUNCTION = 5912

	// ERROR_CLUSTER_PARAMETER_OUT_OF_BOUNDS
	// A parameter value is out of acceptable range.
	ERROR_CLUSTER_PARAMETER_OUT_OF_BOUNDS = 5913

	// ERROR_CLUSTER_PARTIAL_SEND
	// A network error occurred while sending data to another node in the cluster. The number of bytes transmitted was less than required.
	ERROR_CLUSTER_PARTIAL_SEND = 5914

	// ERROR_CLUSTER_REGISTRY_INVALID_FUNCTION
	// An invalid cluster registry operation was attempted.
	ERROR_CLUSTER_REGISTRY_INVALID_FUNCTION = 5915

	// ERROR_CLUSTER_INVALID_STRING_TERMINATION
	// An input string of characters is not properly terminated.
	ERROR_CLUSTER_INVALID_STRING_TERMINATION = 5916

	// ERROR_CLUSTER_INVALID_STRING_FORMAT
	// An input string of characters is not in a valid format for the data it represents.
	ERROR_CLUSTER_INVALID_STRING_FORMAT = 5917

	// ERROR_CLUSTER_DATABASE_TRANSACTION_IN_PROGRESS
	// An internal cluster error occurred. A cluster database transaction was attempted while a transaction was already in progress.
	ERROR_CLUSTER_DATABASE_TRANSACTION_IN_PROGRESS = 5918

	// ERROR_CLUSTER_DATABASE_TRANSACTION_NOT_IN_PROGRESS
	// An internal cluster error occurred. There was an attempt to commit a cluster database transaction while no transaction was in progress.
	ERROR_CLUSTER_DATABASE_TRANSACTION_NOT_IN_PROGRESS = 5919

	// ERROR_CLUSTER_NULL_DATA
	// An internal cluster error occurred. Data was not properly initialized.
	ERROR_CLUSTER_NULL_DATA = 5920

	// ERROR_CLUSTER_PARTIAL_READ
	// An error occurred while reading from a stream of data. An unexpected number of bytes was returned.
	ERROR_CLUSTER_PARTIAL_READ = 5921

	// ERROR_CLUSTER_PARTIAL_WRITE
	// An error occurred while writing to a stream of data. The required number of bytes could not be written.
	ERROR_CLUSTER_PARTIAL_WRITE = 5922

	// ERROR_CLUSTER_CANT_DESERIALIZE_DATA
	// An error occurred while deserializing a stream of cluster data.
	ERROR_CLUSTER_CANT_DESERIALIZE_DATA = 5923

	// ERROR_DEPENDENT_RESOURCE_PROPERTY_CONFLICT
	// One or more property values for this resource are in conflict with one or more property values associated with its dependent resource(s).
	ERROR_DEPENDENT_RESOURCE_PROPERTY_CONFLICT = 5924

	// ERROR_CLUSTER_NO_QUORUM
	// A quorum of cluster nodes was not present to form a cluster.
	ERROR_CLUSTER_NO_QUORUM = 5925

	// ERROR_CLUSTER_INVALID_IPV6_NETWORK
	// The cluster network is not valid for an IPv6 Address resource, or it does not match the configured address.
	ERROR_CLUSTER_INVALID_IPV6_NETWORK = 5926

	// ERROR_CLUSTER_INVALID_IPV6_TUNNEL_NETWORK
	// The cluster network is not valid for an IPv6 Tunnel resource. Check the configuration of the IP Address resource on which the IPv6 Tunnel resource depends.
	ERROR_CLUSTER_INVALID_IPV6_TUNNEL_NETWORK = 5927

	// ERROR_QUORUM_NOT_ALLOWED_IN_THIS_GROUP
	// Quorum resource cannot reside in the Available Storage group.
	ERROR_QUORUM_NOT_ALLOWED_IN_THIS_GROUP = 5928

	// ERROR_DEPENDENCY_TREE_TOO_COMPLEX
	// The dependencies for this resource are nested too deeply.
	ERROR_DEPENDENCY_TREE_TOO_COMPLEX = 5929

	// ERROR_EXCEPTION_IN_RESOURCE_CAL
	// The call into the resource DLL raised an unhandled exception.
	ERROR_EXCEPTION_IN_RESOURCE_CALL = 5930

	// ERROR_CLUSTER_RHS_FAILED_INITIALIZATION
	// The RHS process failed to initialize.
	ERROR_CLUSTER_RHS_FAILED_INITIALIZATION = 5931

	// ERROR_CLUSTER_NOT_INSTALLED
	// The Failover Clustering feature is not installed on this node.
	ERROR_CLUSTER_NOT_INSTALLED = 5932

	// ERROR_CLUSTER_RESOURCES_MUST_BE_ONLINE_ON_THE_SAME_NODE
	// The resources must be online on the same node for this operation
	ERROR_CLUSTER_RESOURCES_MUST_BE_ONLINE_ON_THE_SAME_NODE = 5933

	// ERROR_CLUSTER_MAX_NODES_IN_CLUSTER
	// A new node can not be added since this cluster is already at its maximum number of nodes.
	ERROR_CLUSTER_MAX_NODES_IN_CLUSTER = 5934

	// ERROR_CLUSTER_TOO_MANY_NODES
	// This cluster can not be created since the specified number of nodes exceeds the maximum allowed limit.
	ERROR_CLUSTER_TOO_MANY_NODES = 5935

	// ERROR_CLUSTER_OBJECT_ALREADY_USED
	// An attempt to use the specified cluster name failed because an enabled computer object with the given name already exists in the domain.
	ERROR_CLUSTER_OBJECT_ALREADY_USED = 5936

	// ERROR_NONCORE_GROUPS_FOUND
	// This cluster cannot be destroyed. It has non-core application groups which must be deleted before the cluster can be destroyed.
	ERROR_NONCORE_GROUPS_FOUND = 5937

	// ERROR_FILE_SHARE_RESOURCE_CONFLICT
	// File share associated with file share witness resource cannot be hosted by this cluster or any of its nodes.
	ERROR_FILE_SHARE_RESOURCE_CONFLICT = 5938

	// ERROR_CLUSTER_EVICT_INVALID_REQUEST
	// Eviction of this node is invalid at this time. Due to quorum requirements node eviction will result in cluster shutdown.
	// If it is the last node in the cluster, destroy cluster command should be used.
	ERROR_CLUSTER_EVICT_INVALID_REQUEST = 5939

	// ERROR_CLUSTER_SINGLETON_RESOURCE
	// Only one instance of this resource type is allowed in the cluster.
	ERROR_CLUSTER_SINGLETON_RESOURCE = 5940

	// ERROR_CLUSTER_GROUP_SINGLETON_RESOURCE
	// Only one instance of this resource type is allowed per resource group.
	ERROR_CLUSTER_GROUP_SINGLETON_RESOURCE = 5941

	// ERROR_CLUSTER_RESOURCE_PROVIDER_FAILED
	// The resource failed to come online due to the failure of one or more provider resources.
	ERROR_CLUSTER_RESOURCE_PROVIDER_FAILED = 5942

	// ERROR_CLUSTER_RESOURCE_CONFIGURATION_ERROR
	// The resource has indicated that it cannot come online on any node.
	ERROR_CLUSTER_RESOURCE_CONFIGURATION_ERROR = 5943

	// ERROR_CLUSTER_GROUP_BUSY
	// The current operation cannot be performed on this group at this time.
	ERROR_CLUSTER_GROUP_BUSY = 5944

	// ERROR_CLUSTER_NOT_SHARED_VOLUME
	// The directory or file is not located on a cluster shared volume.
	ERROR_CLUSTER_NOT_SHARED_VOLUME = 5945

	// ERROR_CLUSTER_INVALID_SECURITY_DESCRIPTOR
	// The Security Descriptor does not meet the requirements for a cluster.
	ERROR_CLUSTER_INVALID_SECURITY_DESCRIPTOR = 5946

	// ERROR_CLUSTER_SHARED_VOLUMES_IN_USE
	// There is one or more shared volumes resources configured in the cluster.
	// Those resources must be moved to available storage in order for operation to succeed.
	ERROR_CLUSTER_SHARED_VOLUMES_IN_USE = 5947

	// ERROR_CLUSTER_USE_SHARED_VOLUMES_API
	// This group or resource cannot be directly manipulated.
	// Use shared volume APIs to perform desired operation.
	ERROR_CLUSTER_USE_SHARED_VOLUMES_API = 5948

	// ERROR_CLUSTER_BACKUP_IN_PROGRESS
	// Back up is in progress. Please wait for backup completion before trying this operation again.
	ERROR_CLUSTER_BACKUP_IN_PROGRESS = 5949

	// ERROR_NON_CSV_PATH
	// The path does not belong to a cluster shared volume.
	ERROR_NON_CSV_PATH = 5950

	// ERROR_CSV_VOLUME_NOT_LOCA
	// The cluster shared volume is not locally mounted on this node.
	ERROR_CSV_VOLUME_NOT_LOCAL = 5951

	// ERROR_CLUSTER_WATCHDOG_TERMINATING
	// The cluster watchdog is terminating.
	ERROR_CLUSTER_WATCHDOG_TERMINATING = 5952

	// ERROR_CLUSTER_RESOURCE_VETOED_MOVE_INCOMPATIBLE_NODES
	// A resource vetoed a move between two nodes because they are incompatible.
	ERROR_CLUSTER_RESOURCE_VETOED_MOVE_INCOMPATIBLE_NODES = 5953

	// ERROR_CLUSTER_INVALID_NODE_WEIGHT
	// The request is invalid either because node weight cannot be changed while the cluster is in disk-only quorum mode, or because changing the node weight would violate the minimum cluster quorum requirements.
	ERROR_CLUSTER_INVALID_NODE_WEIGHT = 5954

	// ERROR_CLUSTER_RESOURCE_VETOED_CAL
	// The resource vetoed the call.
	ERROR_CLUSTER_RESOURCE_VETOED_CALL = 5955

	// ERROR_RESMON_SYSTEM_RESOURCES_LACKING
	// Resource could not start or run because it could not reserve sufficient system resources.
	ERROR_RESMON_SYSTEM_RESOURCES_LACKING = 5956

	// ERROR_CLUSTER_RESOURCE_VETOED_MOVE_NOT_ENOUGH_RESOURCES_ON_DESTINATION
	// A resource vetoed a move between two nodes because the destination currently does not have enough resources to complete the operation.
	ERROR_CLUSTER_RESOURCE_VETOED_MOVE_NOT_ENOUGH_RESOURCES_ON_DESTINATION = 5957

	// ERROR_CLUSTER_RESOURCE_VETOED_MOVE_NOT_ENOUGH_RESOURCES_ON_SOURCE
	// A resource vetoed a move between two nodes because the source currently does not have enough resources to complete the operation.
	ERROR_CLUSTER_RESOURCE_VETOED_MOVE_NOT_ENOUGH_RESOURCES_ON_SOURCE = 5958

	// ERROR_CLUSTER_GROUP_QUEUED
	// The requested operation can not be completed because the group is queued for an operation.
	ERROR_CLUSTER_GROUP_QUEUED = 5959

	// ERROR_CLUSTER_RESOURCE_LOCKED_STATUS
	// The requested operation can not be completed because a resource has locked status.
	ERROR_CLUSTER_RESOURCE_LOCKED_STATUS = 5960

	// ERROR_CLUSTER_SHARED_VOLUME_FAILOVER_NOT_ALLOWED
	// The resource cannot move to another node because a cluster shared volume vetoed the operation.
	ERROR_CLUSTER_SHARED_VOLUME_FAILOVER_NOT_ALLOWED = 5961

	// ERROR_CLUSTER_NODE_DRAIN_IN_PROGRESS
	// A node drain is already in progress.
	ERROR_CLUSTER_NODE_DRAIN_IN_PROGRESS = 5962

	// ERROR_CLUSTER_DISK_NOT_CONNECTED
	// Clustered storage is not connected to the node.
	ERROR_CLUSTER_DISK_NOT_CONNECTED = 5963

	// ERROR_DISK_NOT_CSV_CAPABLE
	// The disk is not configured in a way to be used with CSV. CSV disks must have at least one partition that is formatted with NTFS or REFS.
	ERROR_DISK_NOT_CSV_CAPABLE = 5964

	// ERROR_RESOURCE_NOT_IN_AVAILABLE_STORAGE
	// The resource must be part of the Available Storage group to complete this action.
	ERROR_RESOURCE_NOT_IN_AVAILABLE_STORAGE = 5965

	// ERROR_CLUSTER_SHARED_VOLUME_REDIRECTED
	// CSVFS failed operation as volume is in redirected mode.
	ERROR_CLUSTER_SHARED_VOLUME_REDIRECTED = 5966

	// ERROR_CLUSTER_SHARED_VOLUME_NOT_REDIRECTED
	// CSVFS failed operation as volume is not in redirected mode.
	ERROR_CLUSTER_SHARED_VOLUME_NOT_REDIRECTED = 5967

	// ERROR_CLUSTER_CANNOT_RETURN_PROPERTIES
	// Cluster properties cannot be returned at this time.
	ERROR_CLUSTER_CANNOT_RETURN_PROPERTIES = 5968

	// ERROR_CLUSTER_RESOURCE_CONTAINS_UNSUPPORTED_DIFF_AREA_FOR_SHARED_VOLUMES
	// The clustered disk resource contains software snapshot diff area that are not supported for Cluster Shared Volumes.
	ERROR_CLUSTER_RESOURCE_CONTAINS_UNSUPPORTED_DIFF_AREA_FOR_SHARED_VOLUMES = 5969

	// ERROR_CLUSTER_RESOURCE_IS_IN_MAINTENANCE_MODE
	// The operation cannot be completed because the resource is in maintenance mode.
	ERROR_CLUSTER_RESOURCE_IS_IN_MAINTENANCE_MODE = 5970

	// ERROR_CLUSTER_AFFINITY_CONFLICT
	// The operation cannot be completed because of cluster affinity conflicts
	ERROR_CLUSTER_AFFINITY_CONFLICT = 5971

	// ERROR_CLUSTER_RESOURCE_IS_REPLICA_VIRTUAL_MACHINE
	// The operation cannot be completed because the resource is a replica virtual machine.
	ERROR_CLUSTER_RESOURCE_IS_REPLICA_VIRTUAL_MACHINE = 5972

	// ERROR_CLUSTER_UPGRADE_INCOMPATIBLE_VERSIONS
	// The Cluster Functional Level could not be increased because not all nodes in the cluster support the updated version.
	ERROR_CLUSTER_UPGRADE_INCOMPATIBLE_VERSIONS = 5973

	// ERROR_CLUSTER_UPGRADE_FIX_QUORUM_NOT_SUPPORTED
	// Updating the cluster functional level failed because the cluster is running in fix quorum mode.
	// Start additional nodes which are members of the cluster until the cluster reaches quorum and the cluster will automatically
	// switch out of fix quorum mode, or stop and restart the cluster without the FixQuorum switch. Once the cluster is out
	// of fix quorum mode retry the Update-ClusterFunctionalLevel PowerShell cmdlet to update the cluster functional level.
	ERROR_CLUSTER_UPGRADE_FIX_QUORUM_NOT_SUPPORTED = 5974

	// ERROR_CLUSTER_UPGRADE_RESTART_REQUIRED
	// The cluster functional level has been successfully updated but not all features are available yet. Restart the cluster by
	// using the Stop-Cluster PowerShell cmdlet followed by the Start-Cluster PowerShell cmdlet and all cluster features wil
	// be available.
	ERROR_CLUSTER_UPGRADE_RESTART_REQUIRED = 5975

	// ERROR_CLUSTER_UPGRADE_IN_PROGRESS
	// The cluster is currently performing a version upgrade.
	ERROR_CLUSTER_UPGRADE_IN_PROGRESS = 5976

	// ERROR_CLUSTER_UPGRADE_INCOMPLETE
	// The cluster did not successfully complete the version upgrade.
	ERROR_CLUSTER_UPGRADE_INCOMPLETE = 5977

	// ERROR_CLUSTER_NODE_IN_GRACE_PERIOD
	// The cluster node is in grace period.
	ERROR_CLUSTER_NODE_IN_GRACE_PERIOD = 5978

	// ERROR_CLUSTER_CSV_IO_PAUSE_TIMEOUT
	// The operation has failed because CSV volume was not able to recover in time specified on this file object.
	ERROR_CLUSTER_CSV_IO_PAUSE_TIMEOUT = 5979

	// ERROR_NODE_NOT_ACTIVE_CLUSTER_MEMBER
	// The operation failed because the requested node is not currently part of active cluster membership.
	ERROR_NODE_NOT_ACTIVE_CLUSTER_MEMBER = 5980

	// ERROR_CLUSTER_RESOURCE_NOT_MONITORED
	// The operation failed because the requested cluster resource is currently unmonitored.
	ERROR_CLUSTER_RESOURCE_NOT_MONITORED = 5981

	// ERROR_CLUSTER_RESOURCE_DOES_NOT_SUPPORT_UNMONITORED
	// The operation failed because a resource does not support running in an unmonitored state.
	ERROR_CLUSTER_RESOURCE_DOES_NOT_SUPPORT_UNMONITORED = 5982

	// ERROR_CLUSTER_RESOURCE_IS_REPLICATED
	// The operation cannot be completed because a resource participates in replication.
	ERROR_CLUSTER_RESOURCE_IS_REPLICATED = 5983

	// ERROR_CLUSTER_NODE_ISOLATED
	// The operation failed because the requested cluster node has been isolated
	ERROR_CLUSTER_NODE_ISOLATED = 5984

	// ERROR_CLUSTER_NODE_QUARANTINED
	// The operation failed because the requested cluster node has been quarantined
	ERROR_CLUSTER_NODE_QUARANTINED = 5985

	// ERROR_CLUSTER_DATABASE_UPDATE_CONDITION_FAILED
	// The operation failed because the specified database update condition was not met
	ERROR_CLUSTER_DATABASE_UPDATE_CONDITION_FAILED = 5986

	// ERROR_CLUSTER_SPACE_DEGRADED
	// A clustered space is in a degraded condition and the requested action cannot be completed at this time.
	ERROR_CLUSTER_SPACE_DEGRADED = 5987

	// ERROR_CLUSTER_TOKEN_DELEGATION_NOT_SUPPORTED
	// The operation failed because token delegation for this control is not supported.
	ERROR_CLUSTER_TOKEN_DELEGATION_NOT_SUPPORTED = 5988

	// ERROR_CLUSTER_CSV_INVALID_HANDLE
	// The operation has failed because CSV has invalidated this file object.
	ERROR_CLUSTER_CSV_INVALID_HANDLE = 5989

	// ERROR_CLUSTER_CSV_SUPPORTED_ONLY_ON_COORDINATOR
	// This operation is supported only on the CSV coordinator node.
	ERROR_CLUSTER_CSV_SUPPORTED_ONLY_ON_COORDINATOR = 5990

	// ERROR_GROUPSET_NOT_AVAILABLE
	// The cluster group set is not available for any further requests.
	ERROR_GROUPSET_NOT_AVAILABLE = 5991

	// ERROR_GROUPSET_NOT_FOUND
	// The cluster group set could not be found.
	ERROR_GROUPSET_NOT_FOUND = 5992

	// ERROR_GROUPSET_CANT_PROVIDE
	// The action cannot be completed at this time because the cluster group set would fall below quorum and not be able to act as a provider.
	ERROR_GROUPSET_CANT_PROVIDE = 5993

	// ERROR_CLUSTER_FAULT_DOMAIN_PARENT_NOT_FOUND
	// The specified parent fault domain is not found.
	ERROR_CLUSTER_FAULT_DOMAIN_PARENT_NOT_FOUND = 5994

	// ERROR_CLUSTER_FAULT_DOMAIN_INVALID_HIERARCHY
	// The fault domain cannot be a child of the parent specified.
	ERROR_CLUSTER_FAULT_DOMAIN_INVALID_HIERARCHY = 5995

	// ERROR_CLUSTER_FAULT_DOMAIN_FAILED_S2D_VALIDATION
	// Storage Spaces Direct has rejected the proposed fault domain changes because it impacts the fault tolerance of the storage.
	ERROR_CLUSTER_FAULT_DOMAIN_FAILED_S2D_VALIDATION = 5996

	// ERROR_CLUSTER_FAULT_DOMAIN_S2D_CONNECTIVITY_LOSS
	// Storage Spaces Direct has rejected the proposed fault domain changes because it reduces the storage connected to the system.
	ERROR_CLUSTER_FAULT_DOMAIN_S2D_CONNECTIVITY_LOSS = 5997

	// ERROR_CLUSTER_INVALID_INFRASTRUCTURE_FILESERVER_NAME
	// Cluster infrastructure file server creation failed because a valid non-empty file server name was not provided.
	ERROR_CLUSTER_INVALID_INFRASTRUCTURE_FILESERVER_NAME = 5998

	// ERROR_CLUSTERSET_MANAGEMENT_CLUSTER_UNREACHABLE
	// The action cannot be completed because the cluster set management cluster is unreachable.
	ERROR_CLUSTERSET_MANAGEMENT_CLUSTER_UNREACHABLE = 5999

	///////////////////////////////////////////////////
	//                                               //
	//               EFS Error codes                 //
	//                                               //
	//                =6000 to=6099                  //
	///////////////////////////////////////////////////

	// ERROR_ENCRYPTION_FAILED
	// The specified file could not be encrypted.
	ERROR_ENCRYPTION_FAILED = 6000

	// ERROR_DECRYPTION_FAILED
	// The specified file could not be decrypted.
	ERROR_DECRYPTION_FAILED = 6001

	// ERROR_FILE_ENCRYPTED
	// The specified file is encrypted and the user does not have the ability to decrypt it.
	ERROR_FILE_ENCRYPTED = 6002

	// ERROR_NO_RECOVERY_POLICY
	// There is no valid encryption recovery policy configured for this system.
	ERROR_NO_RECOVERY_POLICY = 6003

	// ERROR_NO_EFS
	// The required encryption driver is not loaded for this system.
	ERROR_NO_EFS = 6004

	// ERROR_WRONG_EFS
	// The file was encrypted with a different encryption driver than is currently loaded.
	ERROR_WRONG_EFS = 6005

	// ERROR_NO_USER_KEYS
	// There are no EFS keys defined for the user.
	ERROR_NO_USER_KEYS = 6006

	// ERROR_FILE_NOT_ENCRYPTED
	// The specified file is not encrypted.
	ERROR_FILE_NOT_ENCRYPTED = 6007

	// ERROR_NOT_EXPORT_FORMAT
	// The specified file is not in the defined EFS export format.
	ERROR_NOT_EXPORT_FORMAT = 6008

	// ERROR_FILE_READ_ONLY
	// The specified file is read only.
	ERROR_FILE_READ_ONLY = 6009

	// ERROR_DIR_EFS_DISALLOWED
	// The directory has been disabled for encryption.
	ERROR_DIR_EFS_DISALLOWED = 6010

	// ERROR_EFS_SERVER_NOT_TRUSTED
	// The server is not trusted for remote encryption operation.
	ERROR_EFS_SERVER_NOT_TRUSTED = 6011

	// ERROR_BAD_RECOVERY_POLICY
	// Recovery policy configured for this system contains invalid recovery certificate.
	ERROR_BAD_RECOVERY_POLICY = 6012

	// ERROR_EFS_ALG_BLOB_TOO_BIG
	// The encryption algorithm used on the source file needs a bigger key buffer than the one on the destination file.
	ERROR_EFS_ALG_BLOB_TOO_BIG = 6013

	// ERROR_VOLUME_NOT_SUPPORT_EFS
	// The disk partition does not support file encryption.
	ERROR_VOLUME_NOT_SUPPORT_EFS = 6014

	// ERROR_EFS_DISABLED
	// This machine is disabled for file encryption.
	ERROR_EFS_DISABLED = 6015

	// ERROR_EFS_VERSION_NOT_SUPPORT
	// A newer system is required to decrypt this encrypted file.
	ERROR_EFS_VERSION_NOT_SUPPORT = 6016

	// ERROR_CS_ENCRYPTION_INVALID_SERVER_RESPONSE
	// The remote server sent an invalid response for a file being opened with Client Side Encryption.
	ERROR_CS_ENCRYPTION_INVALID_SERVER_RESPONSE = 6017

	// ERROR_CS_ENCRYPTION_UNSUPPORTED_SERVER
	// Client Side Encryption is not supported by the remote server even though it claims to support it.
	ERROR_CS_ENCRYPTION_UNSUPPORTED_SERVER = 6018

	// ERROR_CS_ENCRYPTION_EXISTING_ENCRYPTED_FILE
	// File is encrypted and should be opened in Client Side Encryption mode.
	ERROR_CS_ENCRYPTION_EXISTING_ENCRYPTED_FILE = 6019

	// ERROR_CS_ENCRYPTION_NEW_ENCRYPTED_FILE
	// A new encrypted file is being created and a $EFS needs to be provided.
	ERROR_CS_ENCRYPTION_NEW_ENCRYPTED_FILE = 6020

	// ERROR_CS_ENCRYPTION_FILE_NOT_CSE
	// The SMB client requested a CSE FSCTL on a non-CSE file.
	ERROR_CS_ENCRYPTION_FILE_NOT_CSE = 6021

	// ERROR_ENCRYPTION_POLICY_DENIES_OPERATION
	// The requested operation was blocked by policy. For more information, contact your system administrator.
	ERROR_ENCRYPTION_POLICY_DENIES_OPERATION = 6022

	///////////////////////////////////////////////////
	//                                               //
	//              BROWSER Error codes              //
	//                                               //
	//                =6100 to=6199                  //
	///////////////////////////////////////////////////

	// This message number is for historical purposes and cannot be changed or re-used.

	// ERROR_NO_BROWSER_SERVERS_FOUND
	// The list of servers for this workgroup is not currently available
	ERROR_NO_BROWSER_SERVERS_FOUND = 6118

	///////////////////////////////////////////////////
	//                                               //
	//            Task Scheduler Error codes         //
	//            NET START must understand          //
	//                                               //
	//                =6200 to=6249                  //
	///////////////////////////////////////////////////

	// SCHED_E_SERVICE_NOT_LOCALSYSTEM
	// The Task Scheduler service must be configured to run in the System account to function properly. Individual tasks may be configured to run in other accounts.

	SCHED_E_SERVICE_NOT_LOCALSYSTEM = 6200

	///////////////////////////////////////////////////
	//                                               //
	//                  Available                    //
	//                                               //
	//                =6250 to=6599                  //
	///////////////////////////////////////////////////

	///////////////////////////////////////////////////
	//                                               //
	//         Common Log (CLFS) Error codes         //
	//                                               //
	//                =6600 to=6699                  //
	///////////////////////////////////////////////////

	// ERROR_LOG_SECTOR_INVALID
	// Log service encountered an invalid log sector.
	ERROR_LOG_SECTOR_INVALID = 6600

	// ERROR_LOG_SECTOR_PARITY_INVALID
	// Log service encountered a log sector with invalid block parity.
	ERROR_LOG_SECTOR_PARITY_INVALID = 6601

	// ERROR_LOG_SECTOR_REMAPPED
	// Log service encountered a remapped log sector.
	ERROR_LOG_SECTOR_REMAPPED = 6602

	// ERROR_LOG_BLOCK_INCOMPLETE
	// Log service encountered a partial or incomplete log block.
	ERROR_LOG_BLOCK_INCOMPLETE = 6603

	// ERROR_LOG_INVALID_RANGE
	// Log service encountered an attempt access data outside the active log range.
	ERROR_LOG_INVALID_RANGE = 6604

	// ERROR_LOG_BLOCKS_EXHAUSTED
	// Log service user marshalling buffers are exhausted.
	ERROR_LOG_BLOCKS_EXHAUSTED = 6605

	// ERROR_LOG_READ_CONTEXT_INVALID
	// Log service encountered an attempt read from a marshalling area with an invalid read context.
	ERROR_LOG_READ_CONTEXT_INVALID = 6606

	// ERROR_LOG_RESTART_INVALID
	// Log service encountered an invalid log restart area.
	ERROR_LOG_RESTART_INVALID = 6607

	// ERROR_LOG_BLOCK_VERSION
	// Log service encountered an invalid log block version.
	ERROR_LOG_BLOCK_VERSION = 6608

	// ERROR_LOG_BLOCK_INVALID
	// Log service encountered an invalid log block.
	ERROR_LOG_BLOCK_INVALID = 6609

	// ERROR_LOG_READ_MODE_INVALID
	// Log service encountered an attempt to read the log with an invalid read mode.
	ERROR_LOG_READ_MODE_INVALID = 6610

	// ERROR_LOG_NO_RESTART
	// Log service encountered a log stream with no restart area.
	ERROR_LOG_NO_RESTART = 6611

	// ERROR_LOG_METADATA_CORRUPT
	// Log service encountered a corrupted metadata file.
	ERROR_LOG_METADATA_CORRUPT = 6612

	// ERROR_LOG_METADATA_INVALID
	// Log service encountered a metadata file that could not be created by the log file system.
	ERROR_LOG_METADATA_INVALID = 6613

	// ERROR_LOG_METADATA_INCONSISTENT
	// Log service encountered a metadata file with inconsistent data.
	ERROR_LOG_METADATA_INCONSISTENT = 6614

	// ERROR_LOG_RESERVATION_INVALID
	// Log service encountered an attempt to erroneous allocate or dispose reservation space.
	ERROR_LOG_RESERVATION_INVALID = 6615

	// ERROR_LOG_CANT_DELETE
	// Log service cannot delete log file or file system container.
	ERROR_LOG_CANT_DELETE = 6616

	// ERROR_LOG_CONTAINER_LIMIT_EXCEEDED
	// Log service has reached the maximum allowable containers allocated to a log file.
	ERROR_LOG_CONTAINER_LIMIT_EXCEEDED = 6617

	// ERROR_LOG_START_OF_LOG
	// Log service has attempted to read or write backward past the start of the log.
	ERROR_LOG_START_OF_LOG = 6618

	// ERROR_LOG_POLICY_ALREADY_INSTALLED
	// Log policy could not be installed because a policy of the same type is already present.
	ERROR_LOG_POLICY_ALREADY_INSTALLED = 6619

	// ERROR_LOG_POLICY_NOT_INSTALLED
	// Log policy in question was not installed at the time of the request.
	ERROR_LOG_POLICY_NOT_INSTALLED = 6620

	// ERROR_LOG_POLICY_INVALID
	// The installed set of policies on the log is invalid.
	ERROR_LOG_POLICY_INVALID = 6621

	// ERROR_LOG_POLICY_CONFLICT
	// A policy on the log in question prevented the operation from completing.
	ERROR_LOG_POLICY_CONFLICT = 6622

	// ERROR_LOG_PINNED_ARCHIVE_TAI
	// Log space cannot be reclaimed because the log is pinned by the archive tail.
	ERROR_LOG_PINNED_ARCHIVE_TAIL = 6623

	// ERROR_LOG_RECORD_NONEXISTENT
	// Log record is not a record in the log file.
	ERROR_LOG_RECORD_NONEXISTENT = 6624

	// ERROR_LOG_RECORDS_RESERVED_INVALID
	// Number of reserved log records or the adjustment of the number of reserved log records is invalid.
	ERROR_LOG_RECORDS_RESERVED_INVALID = 6625

	// ERROR_LOG_SPACE_RESERVED_INVALID
	// Reserved log space or the adjustment of the log space is invalid.
	ERROR_LOG_SPACE_RESERVED_INVALID = 6626

	// ERROR_LOG_TAIL_INVALID
	// An new or existing archive tail or base of the active log is invalid.
	ERROR_LOG_TAIL_INVALID = 6627

	// ERROR_LOG_FUL
	// Log space is exhausted.
	ERROR_LOG_FULL = 6628

	// ERROR_COULD_NOT_RESIZE_LOG
	// The log could not be set to the requested size.
	ERROR_COULD_NOT_RESIZE_LOG = 6629

	// ERROR_LOG_MULTIPLEXED
	// Log is multiplexed, no direct writes to the physical log is allowed.
	ERROR_LOG_MULTIPLEXED = 6630

	// ERROR_LOG_DEDICATED
	// The operation failed because the log is a dedicated log.
	ERROR_LOG_DEDICATED = 6631

	// ERROR_LOG_ARCHIVE_NOT_IN_PROGRESS
	// The operation requires an archive context.
	ERROR_LOG_ARCHIVE_NOT_IN_PROGRESS = 6632

	// ERROR_LOG_ARCHIVE_IN_PROGRESS
	// Log archival is in progress.
	ERROR_LOG_ARCHIVE_IN_PROGRESS = 6633

	// ERROR_LOG_EPHEMERA
	// The operation requires a non-ephemeral log, but the log is ephemeral.
	ERROR_LOG_EPHEMERAL = 6634

	// ERROR_LOG_NOT_ENOUGH_CONTAINERS
	// The log must have at least two containers before it can be read from or written to.
	ERROR_LOG_NOT_ENOUGH_CONTAINERS = 6635

	// ERROR_LOG_CLIENT_ALREADY_REGISTERED
	// A log client has already registered on the stream.
	ERROR_LOG_CLIENT_ALREADY_REGISTERED = 6636

	// ERROR_LOG_CLIENT_NOT_REGISTERED
	// A log client has not been registered on the stream.
	ERROR_LOG_CLIENT_NOT_REGISTERED = 6637

	// ERROR_LOG_FULL_HANDLER_IN_PROGRESS
	// A request has already been made to handle the log full condition.
	ERROR_LOG_FULL_HANDLER_IN_PROGRESS = 6638

	// ERROR_LOG_CONTAINER_READ_FAILED
	// Log service encountered an error when attempting to read from a log container.
	ERROR_LOG_CONTAINER_READ_FAILED = 6639

	// ERROR_LOG_CONTAINER_WRITE_FAILED
	// Log service encountered an error when attempting to write to a log container.
	ERROR_LOG_CONTAINER_WRITE_FAILED = 6640

	// ERROR_LOG_CONTAINER_OPEN_FAILED
	// Log service encountered an error when attempting open a log container.
	ERROR_LOG_CONTAINER_OPEN_FAILED = 6641

	// ERROR_LOG_CONTAINER_STATE_INVALID
	// Log service encountered an invalid container state when attempting a requested action.
	ERROR_LOG_CONTAINER_STATE_INVALID = 6642

	// ERROR_LOG_STATE_INVALID
	// Log service is not in the correct state to perform a requested action.
	ERROR_LOG_STATE_INVALID = 6643

	// ERROR_LOG_PINNED
	// Log space cannot be reclaimed because the log is pinned.
	ERROR_LOG_PINNED = 6644

	// ERROR_LOG_METADATA_FLUSH_FAILED
	// Log metadata flush failed.
	ERROR_LOG_METADATA_FLUSH_FAILED = 6645

	// ERROR_LOG_INCONSISTENT_SECURITY
	// Security on the log and its containers is inconsistent.
	ERROR_LOG_INCONSISTENT_SECURITY = 6646

	// ERROR_LOG_APPENDED_FLUSH_FAILED
	// Records were appended to the log or reservation changes were made, but the log could not be flushed.
	ERROR_LOG_APPENDED_FLUSH_FAILED = 6647

	// ERROR_LOG_PINNED_RESERVATION
	// The log is pinned due to reservation consuming most of the log space. Free some reserved records to make space available.
	ERROR_LOG_PINNED_RESERVATION = 6648

	///////////////////////////////////////////////////
	//                                               //
	//           Transaction (KTM) Error codes       //
	//                                               //
	//                =6700 to=6799                  //
	///////////////////////////////////////////////////

	// ERROR_INVALID_TRANSACTION
	// The transaction handle associated with this operation is not valid.
	ERROR_INVALID_TRANSACTION = 6700

	// ERROR_TRANSACTION_NOT_ACTIVE
	// The requested operation was made in the context of a transaction that is no longer active.
	ERROR_TRANSACTION_NOT_ACTIVE = 6701

	// ERROR_TRANSACTION_REQUEST_NOT_VALID
	// The requested operation is not valid on the Transaction object in its current state.
	ERROR_TRANSACTION_REQUEST_NOT_VALID = 6702

	// ERROR_TRANSACTION_NOT_REQUESTED
	// The caller has called a response API, but the response is not expected because the TM did not issue the corresponding request to the caller.
	ERROR_TRANSACTION_NOT_REQUESTED = 6703

	// ERROR_TRANSACTION_ALREADY_ABORTED
	// It is too late to perform the requested operation, since the Transaction has already been aborted.
	ERROR_TRANSACTION_ALREADY_ABORTED = 6704

	// ERROR_TRANSACTION_ALREADY_COMMITTED
	// It is too late to perform the requested operation, since the Transaction has already been committed.
	ERROR_TRANSACTION_ALREADY_COMMITTED = 6705

	// ERROR_TM_INITIALIZATION_FAILED
	// The Transaction Manager was unable to be successfully initialized. Transacted operations are not supported.
	ERROR_TM_INITIALIZATION_FAILED = 6706

	// ERROR_RESOURCEMANAGER_READ_ONLY
	// The specified ResourceManager made no changes or updates to the resource under this transaction.
	ERROR_RESOURCEMANAGER_READ_ONLY = 6707

	// ERROR_TRANSACTION_NOT_JOINED
	// The resource manager has attempted to prepare a transaction that it has not successfully joined.
	ERROR_TRANSACTION_NOT_JOINED = 6708

	// ERROR_TRANSACTION_SUPERIOR_EXISTS
	// The Transaction object already has a superior enlistment, and the caller attempted an operation that would have created a new superior. Only a single superior enlistment is allow.
	ERROR_TRANSACTION_SUPERIOR_EXISTS = 6709

	// ERROR_CRM_PROTOCOL_ALREADY_EXISTS
	// The RM tried to register a protocol that already exists.
	ERROR_CRM_PROTOCOL_ALREADY_EXISTS = 6710

	// ERROR_TRANSACTION_PROPAGATION_FAILED
	// The attempt to propagate the Transaction failed.
	ERROR_TRANSACTION_PROPAGATION_FAILED = 6711

	// ERROR_CRM_PROTOCOL_NOT_FOUND
	// The requested propagation protocol was not registered as a CRM.
	ERROR_CRM_PROTOCOL_NOT_FOUND = 6712

	// ERROR_TRANSACTION_INVALID_MARSHALL_BUFFER
	// The buffer passed in to PushTransaction or PullTransaction is not in a valid format.
	ERROR_TRANSACTION_INVALID_MARSHALL_BUFFER = 6713

	// ERROR_CURRENT_TRANSACTION_NOT_VALID
	// The current transaction context associated with the thread is not a valid handle to a transaction object.
	ERROR_CURRENT_TRANSACTION_NOT_VALID = 6714

	// ERROR_TRANSACTION_NOT_FOUND
	// The specified Transaction object could not be opened, because it was not found.
	ERROR_TRANSACTION_NOT_FOUND = 6715

	// ERROR_RESOURCEMANAGER_NOT_FOUND
	// The specified ResourceManager object could not be opened, because it was not found.
	ERROR_RESOURCEMANAGER_NOT_FOUND = 6716

	// ERROR_ENLISTMENT_NOT_FOUND
	// The specified Enlistment object could not be opened, because it was not found.
	ERROR_ENLISTMENT_NOT_FOUND = 6717

	// ERROR_TRANSACTIONMANAGER_NOT_FOUND
	// The specified TransactionManager object could not be opened, because it was not found.
	ERROR_TRANSACTIONMANAGER_NOT_FOUND = 6718

	// ERROR_TRANSACTIONMANAGER_NOT_ONLINE
	// The object specified could not be created or opened, because its associated TransactionManager is not online.  The TransactionManager must be brought fully Online by calling RecoverTransactionManager to recover to the end of its LogFile before objects in its Transaction or ResourceManager namespaces can be opened.  In addition, errors in writing records to its LogFile can cause a TransactionManager to go offline.
	ERROR_TRANSACTIONMANAGER_NOT_ONLINE = 6719

	// ERROR_TRANSACTIONMANAGER_RECOVERY_NAME_COLLISION
	// The specified TransactionManager was unable to create the objects contained in its logfile in the Ob namespace. Therefore, the TransactionManager was unable to recover.
	ERROR_TRANSACTIONMANAGER_RECOVERY_NAME_COLLISION = 6720

	// ERROR_TRANSACTION_NOT_ROOT
	// The call to create a superior Enlistment on this Transaction object could not be completed, because the Transaction object specified for the enlistment is a subordinate branch of the Transaction. Only the root of the Transaction can be enlisted on as a superior.
	ERROR_TRANSACTION_NOT_ROOT = 6721

	// ERROR_TRANSACTION_OBJECT_EXPIRED
	// Because the associated transaction manager or resource manager has been closed, the handle is no longer valid.
	ERROR_TRANSACTION_OBJECT_EXPIRED = 6722

	// ERROR_TRANSACTION_RESPONSE_NOT_ENLISTED
	// The specified operation could not be performed on this Superior enlistment, because the enlistment was not created with the corresponding completion response in the NotificationMask.
	ERROR_TRANSACTION_RESPONSE_NOT_ENLISTED = 6723

	// ERROR_TRANSACTION_RECORD_TOO_LONG
	// The specified operation could not be performed, because the record that would be logged was too long. This can occur because of two conditions: either there are too many Enlistments on this Transaction, or the combined RecoveryInformation being logged on behalf of those Enlistments is too long.
	ERROR_TRANSACTION_RECORD_TOO_LONG = 6724

	// ERROR_IMPLICIT_TRANSACTION_NOT_SUPPORTED
	// Implicit transaction are not supported.
	ERROR_IMPLICIT_TRANSACTION_NOT_SUPPORTED = 6725

	// ERROR_TRANSACTION_INTEGRITY_VIOLATED
	// The kernel transaction manager had to abort or forget the transaction because it blocked forward progress.
	ERROR_TRANSACTION_INTEGRITY_VIOLATED = 6726

	// ERROR_TRANSACTIONMANAGER_IDENTITY_MISMATCH
	// The TransactionManager identity that was supplied did not match the one recorded in the TransactionManager's log file.
	ERROR_TRANSACTIONMANAGER_IDENTITY_MISMATCH = 6727

	// ERROR_RM_CANNOT_BE_FROZEN_FOR_SNAPSHOT
	// This snapshot operation cannot continue because a transactional resource manager cannot be frozen in its current state.  Please try again.
	ERROR_RM_CANNOT_BE_FROZEN_FOR_SNAPSHOT = 6728

	// ERROR_TRANSACTION_MUST_WRITETHROUGH
	// The transaction cannot be enlisted on with the specified EnlistmentMask, because the transaction has already completed the PrePrepare phase.  In order to ensure correctness, the ResourceManager must switch to a write-through mode and cease caching data within this transaction.  Enlisting for only subsequent transaction phases may still succeed.
	ERROR_TRANSACTION_MUST_WRITETHROUGH = 6729

	// ERROR_TRANSACTION_NO_SUPERIOR
	// The transaction does not have a superior enlistment.
	ERROR_TRANSACTION_NO_SUPERIOR = 6730

	// ERROR_HEURISTIC_DAMAGE_POSSIBLE
	// The attempt to commit the Transaction completed, but it is possible that some portion of the transaction tree did not commit successfully due to heuristics.  Therefore it is possible that some data modified in the transaction may not have committed, resulting in transactional inconsistency.  If possible, check the consistency of the associated data.
	ERROR_HEURISTIC_DAMAGE_POSSIBLE = 6731

	///////////////////////////////////////////////////
	//                                               //
	//        Transactional File Services (TxF)      //
	//                  Error codes                  //
	//                                               //
	//                =6800 to=6899                  //
	///////////////////////////////////////////////////

	// ERROR_TRANSACTIONAL_CONFLICT
	// The function attempted to use a name that is reserved for use by another transaction.
	ERROR_TRANSACTIONAL_CONFLICT = 6800

	// ERROR_RM_NOT_ACTIVE
	// Transaction support within the specified resource manager is not started or was shut down due to an error.
	ERROR_RM_NOT_ACTIVE = 6801

	// ERROR_RM_METADATA_CORRUPT
	// The metadata of the RM has been corrupted. The RM will not function.
	ERROR_RM_METADATA_CORRUPT = 6802

	// ERROR_DIRECTORY_NOT_RM
	// The specified directory does not contain a resource manager.
	ERROR_DIRECTORY_NOT_RM = 6803

	// ERROR_TRANSACTIONS_UNSUPPORTED_REMOTE
	// The remote server or share does not support transacted file operations.
	ERROR_TRANSACTIONS_UNSUPPORTED_REMOTE = 6805

	// ERROR_LOG_RESIZE_INVALID_SIZE
	// The requested log size is invalid.
	ERROR_LOG_RESIZE_INVALID_SIZE = 6806

	// ERROR_OBJECT_NO_LONGER_EXISTS
	// The object (file, stream, link) corresponding to the handle has been deleted by a Transaction Savepoint Rollback.
	ERROR_OBJECT_NO_LONGER_EXISTS = 6807

	// ERROR_STREAM_MINIVERSION_NOT_FOUND
	// The specified file miniversion was not found for this transacted file open.
	ERROR_STREAM_MINIVERSION_NOT_FOUND = 6808

	// ERROR_STREAM_MINIVERSION_NOT_VALID
	// The specified file miniversion was found but has been invalidated. Most likely cause is a transaction savepoint rollback.
	ERROR_STREAM_MINIVERSION_NOT_VALID = 6809

	// ERROR_MINIVERSION_INACCESSIBLE_FROM_SPECIFIED_TRANSACTION
	// A miniversion may only be opened in the context of the transaction that created it.
	ERROR_MINIVERSION_INACCESSIBLE_FROM_SPECIFIED_TRANSACTION = 6810

	// ERROR_CANT_OPEN_MINIVERSION_WITH_MODIFY_INTENT
	// It is not possible to open a miniversion with modify access.
	ERROR_CANT_OPEN_MINIVERSION_WITH_MODIFY_INTENT = 6811

	// ERROR_CANT_CREATE_MORE_STREAM_MINIVERSIONS
	// It is not possible to create any more miniversions for this stream.
	ERROR_CANT_CREATE_MORE_STREAM_MINIVERSIONS = 6812

	// ERROR_REMOTE_FILE_VERSION_MISMATCH
	// The remote server sent mismatching version number or Fid for a file opened with transactions.
	ERROR_REMOTE_FILE_VERSION_MISMATCH = 6814

	// ERROR_HANDLE_NO_LONGER_VALID
	// The handle has been invalidated by a transaction. The most likely cause is the presence of memory mapping on a file or an open handle when the transaction ended or rolled back to savepoint.
	ERROR_HANDLE_NO_LONGER_VALID = 6815

	// ERROR_NO_TXF_METADATA
	// There is no transaction metadata on the file.
	ERROR_NO_TXF_METADATA = 6816

	// ERROR_LOG_CORRUPTION_DETECTED
	// The log data is corrupt.
	ERROR_LOG_CORRUPTION_DETECTED = 6817

	// ERROR_CANT_RECOVER_WITH_HANDLE_OPEN
	// The file can't be recovered because there is a handle still open on it.
	ERROR_CANT_RECOVER_WITH_HANDLE_OPEN = 6818

	// ERROR_RM_DISCONNECTED
	// The transaction outcome is unavailable because the resource manager responsible for it has disconnected.
	ERROR_RM_DISCONNECTED = 6819

	// ERROR_ENLISTMENT_NOT_SUPERIOR
	// The request was rejected because the enlistment in question is not a superior enlistment.
	ERROR_ENLISTMENT_NOT_SUPERIOR = 6820

	// ERROR_RECOVERY_NOT_NEEDED
	// The transactional resource manager is already consistent. Recovery is not needed.
	ERROR_RECOVERY_NOT_NEEDED = 6821

	// ERROR_RM_ALREADY_STARTED
	// The transactional resource manager has already been started.
	ERROR_RM_ALREADY_STARTED = 6822

	// ERROR_FILE_IDENTITY_NOT_PERSISTENT
	// The file cannot be opened transactionally, because its identity depends on the outcome of an unresolved transaction.
	ERROR_FILE_IDENTITY_NOT_PERSISTENT = 6823

	// ERROR_CANT_BREAK_TRANSACTIONAL_DEPENDENCY
	// The operation cannot be performed because another transaction is depending on the fact that this property will not change.
	ERROR_CANT_BREAK_TRANSACTIONAL_DEPENDENCY = 6824

	// ERROR_CANT_CROSS_RM_BOUNDARY
	// The operation would involve a single file with two transactional resource managers and is therefore not allowed.
	ERROR_CANT_CROSS_RM_BOUNDARY = 6825

	// ERROR_TXF_DIR_NOT_EMPTY
	// The $Txf directory must be empty for this operation to succeed.
	ERROR_TXF_DIR_NOT_EMPTY = 6826

	// ERROR_INDOUBT_TRANSACTIONS_EXIST
	// The operation would leave a transactional resource manager in an inconsistent state and is therefore not allowed.
	ERROR_INDOUBT_TRANSACTIONS_EXIST = 6827

	// ERROR_TM_VOLATILE
	// The operation could not be completed because the transaction manager does not have a log.
	ERROR_TM_VOLATILE = 6828

	// ERROR_ROLLBACK_TIMER_EXPIRED
	// A rollback could not be scheduled because a previously scheduled rollback has already executed or been queued for execution.
	ERROR_ROLLBACK_TIMER_EXPIRED = 6829

	// ERROR_TXF_ATTRIBUTE_CORRUPT
	// The transactional metadata attribute on the file or directory is corrupt and unreadable.
	ERROR_TXF_ATTRIBUTE_CORRUPT = 6830

	// ERROR_EFS_NOT_ALLOWED_IN_TRANSACTION
	// The encryption operation could not be completed because a transaction is active.
	ERROR_EFS_NOT_ALLOWED_IN_TRANSACTION = 6831

	// ERROR_TRANSACTIONAL_OPEN_NOT_ALLOWED
	// This object is not allowed to be opened in a transaction.
	ERROR_TRANSACTIONAL_OPEN_NOT_ALLOWED = 6832

	// ERROR_LOG_GROWTH_FAILED
	// An attempt to create space in the transactional resource manager's log failed. The failure status has been recorded in the event log.
	ERROR_LOG_GROWTH_FAILED = 6833

	// ERROR_TRANSACTED_MAPPING_UNSUPPORTED_REMOTE
	// Memory mapping (creating a mapped section) a remote file under a transaction is not supported.
	ERROR_TRANSACTED_MAPPING_UNSUPPORTED_REMOTE = 6834

	// ERROR_TXF_METADATA_ALREADY_PRESENT
	// Transaction metadata is already present on this file and cannot be superseded.
	ERROR_TXF_METADATA_ALREADY_PRESENT = 6835

	// ERROR_TRANSACTION_SCOPE_CALLBACKS_NOT_SET
	// A transaction scope could not be entered because the scope handler has not been initialized.
	ERROR_TRANSACTION_SCOPE_CALLBACKS_NOT_SET = 6836

	// ERROR_TRANSACTION_REQUIRED_PROMOTION
	// Promotion was required in order to allow the resource manager to enlist, but the transaction was set to disallow it.
	ERROR_TRANSACTION_REQUIRED_PROMOTION = 6837

	// ERROR_CANNOT_EXECUTE_FILE_IN_TRANSACTION
	// This file is open for modification in an unresolved transaction and may be opened for execute only by a transacted reader.
	ERROR_CANNOT_EXECUTE_FILE_IN_TRANSACTION = 6838

	// ERROR_TRANSACTIONS_NOT_FROZEN
	// The request to thaw frozen transactions was ignored because transactions had not previously been frozen.
	ERROR_TRANSACTIONS_NOT_FROZEN = 6839

	// ERROR_TRANSACTION_FREEZE_IN_PROGRESS
	// Transactions cannot be frozen because a freeze is already in progress.
	ERROR_TRANSACTION_FREEZE_IN_PROGRESS = 6840

	// ERROR_NOT_SNAPSHOT_VOLUME
	// The target volume is not a snapshot volume. This operation is only valid on a volume mounted as a snapshot.
	ERROR_NOT_SNAPSHOT_VOLUME = 6841

	// ERROR_NO_SAVEPOINT_WITH_OPEN_FILES
	// The savepoint operation failed because files are open on the transaction. This is not permitted.
	ERROR_NO_SAVEPOINT_WITH_OPEN_FILES = 6842

	// ERROR_DATA_LOST_REPAIR
	// Windows has discovered corruption in a file, and that file has since been repaired. Data loss may have occurred.
	ERROR_DATA_LOST_REPAIR = 6843

	// ERROR_SPARSE_NOT_ALLOWED_IN_TRANSACTION
	// The sparse operation could not be completed because a transaction is active on the file.
	ERROR_SPARSE_NOT_ALLOWED_IN_TRANSACTION = 6844

	// ERROR_TM_IDENTITY_MISMATCH
	// The call to create a TransactionManager object failed because the Tm Identity stored in the logfile does not match the Tm Identity that was passed in as an argument.
	ERROR_TM_IDENTITY_MISMATCH = 6845

	// ERROR_FLOATED_SECTION
	// I/O was attempted on a section object that has been floated as a result of a transaction ending. There is no valid data.
	ERROR_FLOATED_SECTION = 6846

	// ERROR_CANNOT_ACCEPT_TRANSACTED_WORK
	// The transactional resource manager cannot currently accept transacted work due to a transient condition such as low resources.
	ERROR_CANNOT_ACCEPT_TRANSACTED_WORK = 6847

	// ERROR_CANNOT_ABORT_TRANSACTIONS
	// The transactional resource manager had too many transactions outstanding that could not be aborted. The transactional resource manger has been shut down.
	ERROR_CANNOT_ABORT_TRANSACTIONS = 6848

	// ERROR_BAD_CLUSTERS
	// The operation could not be completed due to bad clusters on disk.
	ERROR_BAD_CLUSTERS = 6849

	// ERROR_COMPRESSION_NOT_ALLOWED_IN_TRANSACTION
	// The compression operation could not be completed because a transaction is active on the file.
	ERROR_COMPRESSION_NOT_ALLOWED_IN_TRANSACTION = 6850

	// ERROR_VOLUME_DIRTY
	// The operation could not be completed because the volume is dirty. Please run chkdsk and try again.
	ERROR_VOLUME_DIRTY = 6851

	// ERROR_NO_LINK_TRACKING_IN_TRANSACTION
	// The link tracking operation could not be completed because a transaction is active.
	ERROR_NO_LINK_TRACKING_IN_TRANSACTION = 6852

	// ERROR_OPERATION_NOT_SUPPORTED_IN_TRANSACTION
	// This operation cannot be performed in a transaction.
	ERROR_OPERATION_NOT_SUPPORTED_IN_TRANSACTION = 6853

	// ERROR_EXPIRED_HANDLE
	// The handle is no longer properly associated with its transaction.  It may have been opened in a transactional resource manager that was subsequently forced to restart.  Please close the handle and open a new one.
	ERROR_EXPIRED_HANDLE = 6854

	// ERROR_TRANSACTION_NOT_ENLISTED
	// The specified operation could not be performed because the resource manager is not enlisted in the transaction.
	ERROR_TRANSACTION_NOT_ENLISTED = 6855

	///////////////////////////////////////////////////
	//                                               //
	//                  Available                    //
	//                                               //
	//                =6900 to=6999                  //
	///////////////////////////////////////////////////

	///////////////////////////////////////////////////
	//                                               //
	//          Terminal Server Error codes          //
	//                                               //
	//                =7000 to=7099                  //
	///////////////////////////////////////////////////

	// ERROR_CTX_WINSTATION_NAME_INVALID
	// The specified session name is invalid.
	ERROR_CTX_WINSTATION_NAME_INVALID = 7001

	// ERROR_CTX_INVALID_PD
	// The specified protocol driver is invalid.
	ERROR_CTX_INVALID_PD = 7002

	// ERROR_CTX_PD_NOT_FOUND
	// The specified protocol driver was not found in the system path.
	ERROR_CTX_PD_NOT_FOUND = 7003

	// ERROR_CTX_WD_NOT_FOUND
	// The specified terminal connection driver was not found in the system path.
	ERROR_CTX_WD_NOT_FOUND = 7004

	// ERROR_CTX_CANNOT_MAKE_EVENTLOG_ENTRY
	// A registry key for event logging could not be created for this session.
	ERROR_CTX_CANNOT_MAKE_EVENTLOG_ENTRY = 7005

	// ERROR_CTX_SERVICE_NAME_COLLISION
	// A service with the same name already exists on the system.
	ERROR_CTX_SERVICE_NAME_COLLISION = 7006

	// ERROR_CTX_CLOSE_PENDING
	// A close operation is pending on the session.
	ERROR_CTX_CLOSE_PENDING = 7007

	// ERROR_CTX_NO_OUTBUF
	// There are no free output buffers available.
	ERROR_CTX_NO_OUTBUF = 7008

	// ERROR_CTX_MODEM_INF_NOT_FOUND
	// The MODEM.INF file was not found.
	ERROR_CTX_MODEM_INF_NOT_FOUND = 7009

	// ERROR_CTX_INVALID_MODEMNAME
	// The modem name was not found in MODEM.INF.
	ERROR_CTX_INVALID_MODEMNAME = 7010

	// ERROR_CTX_MODEM_RESPONSE_ERROR
	// The modem did not accept the command sent to it. Verify that the configured modem name matches the attached modem.
	ERROR_CTX_MODEM_RESPONSE_ERROR = 7011

	// ERROR_CTX_MODEM_RESPONSE_TIMEOUT
	// The modem did not respond to the command sent to it. Verify that the modem is properly cabled and powered on.
	ERROR_CTX_MODEM_RESPONSE_TIMEOUT = 7012

	// ERROR_CTX_MODEM_RESPONSE_NO_CARRIER
	// Carrier detect has failed or carrier has been dropped due to disconnect.
	ERROR_CTX_MODEM_RESPONSE_NO_CARRIER = 7013

	// ERROR_CTX_MODEM_RESPONSE_NO_DIALTONE
	// Dial tone not detected within the required time. Verify that the phone cable is properly attached and functional.
	ERROR_CTX_MODEM_RESPONSE_NO_DIALTONE = 7014

	// ERROR_CTX_MODEM_RESPONSE_BUSY
	// Busy signal detected at remote site on callback.
	ERROR_CTX_MODEM_RESPONSE_BUSY = 7015

	// ERROR_CTX_MODEM_RESPONSE_VOICE
	// Voice detected at remote site on callback.
	ERROR_CTX_MODEM_RESPONSE_VOICE = 7016

	// ERROR_CTX_TD_ERROR
	// Transport driver error
	ERROR_CTX_TD_ERROR = 7017

	// ERROR_CTX_WINSTATION_NOT_FOUND
	// The specified session cannot be found.
	ERROR_CTX_WINSTATION_NOT_FOUND = 7022

	// ERROR_CTX_WINSTATION_ALREADY_EXISTS
	// The specified session name is already in use.
	ERROR_CTX_WINSTATION_ALREADY_EXISTS = 7023

	// ERROR_CTX_WINSTATION_BUSY
	// The task you are trying to do can't be completed because Remote Desktop Services is currently busy. Please try again in a few minutes. Other users should still be able to log on.
	ERROR_CTX_WINSTATION_BUSY = 7024

	// ERROR_CTX_BAD_VIDEO_MODE
	// An attempt has been made to connect to a session whose video mode is not supported by the current client.
	ERROR_CTX_BAD_VIDEO_MODE = 7025

	// ERROR_CTX_GRAPHICS_INVALID
	// The application attempted to enable DOS graphics mode. DOS graphics mode is not supported.
	ERROR_CTX_GRAPHICS_INVALID = 7035

	// ERROR_CTX_LOGON_DISABLED
	// Your interactive logon privilege has been disabled. Please contact your administrator.
	ERROR_CTX_LOGON_DISABLED = 7037

	// ERROR_CTX_NOT_CONSOLE
	// The requested operation can be performed only on the system console. This is most often the result of a driver or system DLL requiring direct console access.
	ERROR_CTX_NOT_CONSOLE = 7038

	// ERROR_CTX_CLIENT_QUERY_TIMEOUT
	// The client failed to respond to the server connect message.
	ERROR_CTX_CLIENT_QUERY_TIMEOUT = 7040

	// ERROR_CTX_CONSOLE_DISCONNECT
	// Disconnecting the console session is not supported.
	ERROR_CTX_CONSOLE_DISCONNECT = 7041

	// ERROR_CTX_CONSOLE_CONNECT
	// Reconnecting a disconnected session to the console is not supported.
	ERROR_CTX_CONSOLE_CONNECT = 7042

	// ERROR_CTX_SHADOW_DENIED
	// The request to control another session remotely was denied.
	ERROR_CTX_SHADOW_DENIED = 7044

	// ERROR_CTX_WINSTATION_ACCESS_DENIED
	// The requested session access is denied.
	ERROR_CTX_WINSTATION_ACCESS_DENIED = 7045

	// ERROR_CTX_INVALID_WD
	// The specified terminal connection driver is invalid.
	ERROR_CTX_INVALID_WD = 7049

	// ERROR_CTX_SHADOW_INVALID
	// The requested session cannot be controlled remotely.
	// This may be because the session is disconnected or does not currently have a user logged on.
	ERROR_CTX_SHADOW_INVALID = 7050

	// ERROR_CTX_SHADOW_DISABLED
	// The requested session is not configured to allow remote control.
	ERROR_CTX_SHADOW_DISABLED = 7051

	// ERROR_CTX_CLIENT_LICENSE_IN_USE
	// Your request to connect to this Terminal Server has been rejected. Your Terminal Server client license number is currently being used by another user. Please call your system administrator to obtain a unique license number.
	ERROR_CTX_CLIENT_LICENSE_IN_USE = 7052

	// ERROR_CTX_CLIENT_LICENSE_NOT_SET
	// Your request to connect to this Terminal Server has been rejected. Your Terminal Server client license number has not been entered for this copy of the Terminal Server client. Please contact your system administrator.
	ERROR_CTX_CLIENT_LICENSE_NOT_SET = 7053

	// ERROR_CTX_LICENSE_NOT_AVAILABLE
	// The number of connections to this computer is limited and all connections are in use right now. Try connecting later or contact your system administrator.
	ERROR_CTX_LICENSE_NOT_AVAILABLE = 7054

	// ERROR_CTX_LICENSE_CLIENT_INVALID
	// The client you are using is not licensed to use this system. Your logon request is denied.
	ERROR_CTX_LICENSE_CLIENT_INVALID = 7055

	// ERROR_CTX_LICENSE_EXPIRED
	// The system license has expired. Your logon request is denied.
	ERROR_CTX_LICENSE_EXPIRED = 7056

	// ERROR_CTX_SHADOW_NOT_RUNNING
	// Remote control could not be terminated because the specified session is not currently being remotely controlled.
	ERROR_CTX_SHADOW_NOT_RUNNING = 7057

	// ERROR_CTX_SHADOW_ENDED_BY_MODE_CHANGE
	// The remote control of the console was terminated because the display mode was changed. Changing the display mode in a remote control session is not supported.
	ERROR_CTX_SHADOW_ENDED_BY_MODE_CHANGE = 7058

	// ERROR_ACTIVATION_COUNT_EXCEEDED
	// Activation has already been reset the maximum number of times for this installation. Your activation timer will not be cleared.
	ERROR_ACTIVATION_COUNT_EXCEEDED = 7059

	// ERROR_CTX_WINSTATIONS_DISABLED
	// Remote logins are currently disabled.
	ERROR_CTX_WINSTATIONS_DISABLED = 7060

	// ERROR_CTX_ENCRYPTION_LEVEL_REQUIRED
	// You do not have the proper encryption level to access this Session.
	ERROR_CTX_ENCRYPTION_LEVEL_REQUIRED = 7061

	// ERROR_CTX_SESSION_IN_USE
	// The user %s\\%s is currently logged on to this computer. Only the current user or an administrator can log on to this computer.
	ERROR_CTX_SESSION_IN_USE = 7062

	// ERROR_CTX_NO_FORCE_LOGOFF
	// The user %s\\%s is already logged on to the console of this computer. You do not have permission to log in at this time. To resolve this issue, contact %s\\%s and have them log off.
	ERROR_CTX_NO_FORCE_LOGOFF = 7063

	// ERROR_CTX_ACCOUNT_RESTRICTION
	// Unable to log you on because of an account restriction.
	ERROR_CTX_ACCOUNT_RESTRICTION = 7064

	// ERROR_RDP_PROTOCOL_ERROR
	// The RDP protocol component %2 detected an error in the protocol stream and has disconnected the client.
	ERROR_RDP_PROTOCOL_ERROR = 7065

	// ERROR_CTX_CDM_CONNECT
	// The Client Drive Mapping Service Has Connected on Terminal Connection.
	ERROR_CTX_CDM_CONNECT = 7066

	// ERROR_CTX_CDM_DISCONNECT
	// The Client Drive Mapping Service Has Disconnected on Terminal Connection.
	ERROR_CTX_CDM_DISCONNECT = 7067

	// ERROR_CTX_SECURITY_LAYER_ERROR
	// The Terminal Server security layer detected an error in the protocol stream and has disconnected the client.
	ERROR_CTX_SECURITY_LAYER_ERROR = 7068

	// ERROR_TS_INCOMPATIBLE_SESSIONS
	// The target session is incompatible with the current session.
	ERROR_TS_INCOMPATIBLE_SESSIONS = 7069

	// ERROR_TS_VIDEO_SUBSYSTEM_ERROR
	// Windows can't connect to your session because a problem occurred in the Windows video subsystem. Try connecting again later, or contact the server administrator for assistance.
	ERROR_TS_VIDEO_SUBSYSTEM_ERROR = 7070

	///////////////////////////////////////////////////
	//                                               //
	//          Windows Fabric Error Codes           //
	//                                               //
	//                =7100 to=7499                  //
	//                                               //
	//          defined in FabricCommon.idl          //
	//                                               //
	///////////////////////////////////////////////////

	///////////////////////////////////////////////////
	//                                                /
	//           Traffic Control Error Codes          /
	//                                                /
	//                 =7500 to=7999                  /
	//                                                /
	//            defined in: tcerror.h               /
	///////////////////////////////////////////////////

	///////////////////////////////////////////////////
	//                                               //
	//           Active Directory Error codes        //
	//                                               //
	//                =8000 to=8999                  //
	///////////////////////////////////////////////////

	// *****************
	// FACILITY_FILE_REPLICATION_SERVICE
	// *****************

	// FRS_ERR_INVALID_API_SEQUENCE
	// The file replication service API was called incorrectly.

	FRS_ERR_INVALID_API_SEQUENCE = 8001

	// FRS_ERR_STARTING_SERVICE
	// The file replication service cannot be started.

	FRS_ERR_STARTING_SERVICE = 8002

	// FRS_ERR_STOPPING_SERVICE
	// The file replication service cannot be stopped.

	FRS_ERR_STOPPING_SERVICE = 8003

	// FRS_ERR_INTERNAL_API
	// The file replication service API terminated the request. The event log may have more information.

	FRS_ERR_INTERNAL_API = 8004

	// FRS_ERR_INTERNA
	// The file replication service terminated the request. The event log may have more information.

	FRS_ERR_INTERNAL = 8005

	// FRS_ERR_SERVICE_COMM
	// The file replication service cannot be contacted. The event log may have more information.

	FRS_ERR_SERVICE_COMM = 8006

	// FRS_ERR_INSUFFICIENT_PRIV
	// The file replication service cannot satisfy the request because the user has insufficient privileges. The event log may have more information.

	FRS_ERR_INSUFFICIENT_PRIV = 8007

	// FRS_ERR_AUTHENTICATION
	// The file replication service cannot satisfy the request because authenticated RPC is not available. The event log may have more information.

	FRS_ERR_AUTHENTICATION = 8008

	// FRS_ERR_PARENT_INSUFFICIENT_PRIV
	// The file replication service cannot satisfy the request because the user has insufficient privileges on the domain controller. The event log may have more information.

	FRS_ERR_PARENT_INSUFFICIENT_PRIV = 8009

	// FRS_ERR_PARENT_AUTHENTICATION
	// The file replication service cannot satisfy the request because authenticated RPC is not available on the domain controller. The event log may have more information.

	FRS_ERR_PARENT_AUTHENTICATION = 8010

	// FRS_ERR_CHILD_TO_PARENT_COMM
	// The file replication service cannot communicate with the file replication service on the domain controller. The event log may have more information.

	FRS_ERR_CHILD_TO_PARENT_COMM = 8011

	// FRS_ERR_PARENT_TO_CHILD_COMM
	// The file replication service on the domain controller cannot communicate with the file replication service on this computer. The event log may have more information.

	FRS_ERR_PARENT_TO_CHILD_COMM = 8012

	// FRS_ERR_SYSVOL_POPULATE
	// The file replication service cannot populate the system volume because of an internal error. The event log may have more information.

	FRS_ERR_SYSVOL_POPULATE = 8013

	// FRS_ERR_SYSVOL_POPULATE_TIMEOUT
	// The file replication service cannot populate the system volume because of an internal timeout. The event log may have more information.

	FRS_ERR_SYSVOL_POPULATE_TIMEOUT = 8014

	// FRS_ERR_SYSVOL_IS_BUSY
	// The file replication service cannot process the request. The system volume is busy with a previous request.

	FRS_ERR_SYSVOL_IS_BUSY = 8015

	// FRS_ERR_SYSVOL_DEMOTE
	// The file replication service cannot stop replicating the system volume because of an internal error. The event log may have more information.

	FRS_ERR_SYSVOL_DEMOTE = 8016

	// FRS_ERR_INVALID_SERVICE_PARAMETER
	// The file replication service detected an invalid parameter.

	FRS_ERR_INVALID_SERVICE_PARAMETER = 8017

	// *****************
	// FACILITY DIRECTORY SERVICE
	// *****************
	DS_S_SUCCESS = NO_ERROR

	// ERROR_DS_NOT_INSTALLED
	// An error occurred while installing the directory service. For more information, see the event log.
	ERROR_DS_NOT_INSTALLED = 8200

	// ERROR_DS_MEMBERSHIP_EVALUATED_LOCALLY
	// The directory service evaluated group memberships locally.
	ERROR_DS_MEMBERSHIP_EVALUATED_LOCALLY = 8201

	// ERROR_DS_NO_ATTRIBUTE_OR_VALUE
	// The specified directory service attribute or value does not exist.
	ERROR_DS_NO_ATTRIBUTE_OR_VALUE = 8202

	// ERROR_DS_INVALID_ATTRIBUTE_SYNTAX
	// The attribute syntax specified to the directory service is invalid.
	ERROR_DS_INVALID_ATTRIBUTE_SYNTAX = 8203

	// ERROR_DS_ATTRIBUTE_TYPE_UNDEFINED
	// The attribute type specified to the directory service is not defined.
	ERROR_DS_ATTRIBUTE_TYPE_UNDEFINED = 8204

	// ERROR_DS_ATTRIBUTE_OR_VALUE_EXISTS
	// The specified directory service attribute or value already exists.
	ERROR_DS_ATTRIBUTE_OR_VALUE_EXISTS = 8205

	// ERROR_DS_BUSY
	// The directory service is busy.
	ERROR_DS_BUSY = 8206

	// ERROR_DS_UNAVAILABLE
	// The directory service is unavailable.
	ERROR_DS_UNAVAILABLE = 8207

	// ERROR_DS_NO_RIDS_ALLOCATED
	// The directory service was unable to allocate a relative identifier.
	ERROR_DS_NO_RIDS_ALLOCATED = 8208

	// ERROR_DS_NO_MORE_RIDS
	// The directory service has exhausted the pool of relative identifiers.
	ERROR_DS_NO_MORE_RIDS = 8209

	// ERROR_DS_INCORRECT_ROLE_OWNER
	// The requested operation could not be performed because the directory service is not the master for that type of operation.
	ERROR_DS_INCORRECT_ROLE_OWNER = 8210

	// ERROR_DS_RIDMGR_INIT_ERROR
	// The directory service was unable to initialize the subsystem that allocates relative identifiers.
	ERROR_DS_RIDMGR_INIT_ERROR = 8211

	// ERROR_DS_OBJ_CLASS_VIOLATION
	// The requested operation did not satisfy one or more constraints associated with the class of the object.
	ERROR_DS_OBJ_CLASS_VIOLATION = 8212

	// ERROR_DS_CANT_ON_NON_LEAF
	// The directory service can perform the requested operation only on a leaf object.
	ERROR_DS_CANT_ON_NON_LEAF = 8213

	// ERROR_DS_CANT_ON_RDN
	// The directory service cannot perform the requested operation on the RDN attribute of an object.
	ERROR_DS_CANT_ON_RDN = 8214

	// ERROR_DS_CANT_MOD_OBJ_CLASS
	// The directory service detected an attempt to modify the object class of an object.
	ERROR_DS_CANT_MOD_OBJ_CLASS = 8215

	// ERROR_DS_CROSS_DOM_MOVE_ERROR
	// The requested cross-domain move operation could not be performed.
	ERROR_DS_CROSS_DOM_MOVE_ERROR = 8216

	// ERROR_DS_GC_NOT_AVAILABLE
	// Unable to contact the global catalog server.
	ERROR_DS_GC_NOT_AVAILABLE = 8217

	// ERROR_SHARED_POLICY
	// The policy object is shared and can only be modified at the root.
	ERROR_SHARED_POLICY = 8218

	// ERROR_POLICY_OBJECT_NOT_FOUND
	// The policy object does not exist.
	ERROR_POLICY_OBJECT_NOT_FOUND = 8219

	// ERROR_POLICY_ONLY_IN_DS
	// The requested policy information is only in the directory service.
	ERROR_POLICY_ONLY_IN_DS = 8220

	// ERROR_PROMOTION_ACTIVE
	// A domain controller promotion is currently active.
	ERROR_PROMOTION_ACTIVE = 8221

	// ERROR_NO_PROMOTION_ACTIVE
	// A domain controller promotion is not currently active
	ERROR_NO_PROMOTION_ACTIVE = 8222

	//=8223 unused

	// ERROR_DS_OPERATIONS_ERROR
	// An operations error occurred.
	ERROR_DS_OPERATIONS_ERROR = 8224

	// ERROR_DS_PROTOCOL_ERROR
	// A protocol error occurred.
	ERROR_DS_PROTOCOL_ERROR = 8225

	// ERROR_DS_TIMELIMIT_EXCEEDED
	// The time limit for this request was exceeded.
	ERROR_DS_TIMELIMIT_EXCEEDED = 8226

	// ERROR_DS_SIZELIMIT_EXCEEDED
	// The size limit for this request was exceeded.
	ERROR_DS_SIZELIMIT_EXCEEDED = 8227

	// ERROR_DS_ADMIN_LIMIT_EXCEEDED
	// The administrative limit for this request was exceeded.
	ERROR_DS_ADMIN_LIMIT_EXCEEDED = 8228

	// ERROR_DS_COMPARE_FALSE
	// The compare response was false.
	ERROR_DS_COMPARE_FALSE = 8229

	// ERROR_DS_COMPARE_TRUE
	// The compare response was true.
	ERROR_DS_COMPARE_TRUE = 8230

	// ERROR_DS_AUTH_METHOD_NOT_SUPPORTED
	// The requested authentication method is not supported by the server.
	ERROR_DS_AUTH_METHOD_NOT_SUPPORTED = 8231

	// ERROR_DS_STRONG_AUTH_REQUIRED
	// A more secure authentication method is required for this server.
	ERROR_DS_STRONG_AUTH_REQUIRED = 8232

	// ERROR_DS_INAPPROPRIATE_AUTH
	// Inappropriate authentication.
	ERROR_DS_INAPPROPRIATE_AUTH = 8233

	// ERROR_DS_AUTH_UNKNOWN
	// The authentication mechanism is unknown.
	ERROR_DS_AUTH_UNKNOWN = 8234

	// ERROR_DS_REFERRA
	// A referral was returned from the server.
	ERROR_DS_REFERRAL = 8235

	// ERROR_DS_UNAVAILABLE_CRIT_EXTENSION
	// The server does not support the requested critical extension.
	ERROR_DS_UNAVAILABLE_CRIT_EXTENSION = 8236

	// ERROR_DS_CONFIDENTIALITY_REQUIRED
	// This request requires a secure connection.
	ERROR_DS_CONFIDENTIALITY_REQUIRED = 8237

	// ERROR_DS_INAPPROPRIATE_MATCHING
	// Inappropriate matching.
	ERROR_DS_INAPPROPRIATE_MATCHING = 8238

	// ERROR_DS_CONSTRAINT_VIOLATION
	// A constraint violation occurred.
	ERROR_DS_CONSTRAINT_VIOLATION = 8239

	// ERROR_DS_NO_SUCH_OBJECT
	// There is no such object on the server.
	ERROR_DS_NO_SUCH_OBJECT = 8240

	// ERROR_DS_ALIAS_PROBLEM
	// There is an alias problem.
	ERROR_DS_ALIAS_PROBLEM = 8241

	// ERROR_DS_INVALID_DN_SYNTAX
	// An invalid dn syntax has been specified.
	ERROR_DS_INVALID_DN_SYNTAX = 8242

	// ERROR_DS_IS_LEAF
	// The object is a leaf object.
	ERROR_DS_IS_LEAF = 8243

	// ERROR_DS_ALIAS_DEREF_PROBLEM
	// There is an alias dereferencing problem.
	ERROR_DS_ALIAS_DEREF_PROBLEM = 8244

	// ERROR_DS_UNWILLING_TO_PERFORM
	// The server is unwilling to process the request.
	ERROR_DS_UNWILLING_TO_PERFORM = 8245

	// ERROR_DS_LOOP_DETECT
	// A loop has been detected.
	ERROR_DS_LOOP_DETECT = 8246

	// ERROR_DS_NAMING_VIOLATION
	// There is a naming violation.
	ERROR_DS_NAMING_VIOLATION = 8247

	// ERROR_DS_OBJECT_RESULTS_TOO_LARGE
	// The result set is too large.
	ERROR_DS_OBJECT_RESULTS_TOO_LARGE = 8248

	// ERROR_DS_AFFECTS_MULTIPLE_DSAS
	// The operation affects multiple DSAs
	ERROR_DS_AFFECTS_MULTIPLE_DSAS = 8249

	// ERROR_DS_SERVER_DOWN
	// The server is not operational.
	ERROR_DS_SERVER_DOWN = 8250

	// ERROR_DS_LOCAL_ERROR
	// A local error has occurred.
	ERROR_DS_LOCAL_ERROR = 8251

	// ERROR_DS_ENCODING_ERROR
	// An encoding error has occurred.
	ERROR_DS_ENCODING_ERROR = 8252

	// ERROR_DS_DECODING_ERROR
	// A decoding error has occurred.
	ERROR_DS_DECODING_ERROR = 8253

	// ERROR_DS_FILTER_UNKNOWN
	// The search filter cannot be recognized.
	ERROR_DS_FILTER_UNKNOWN = 8254

	// ERROR_DS_PARAM_ERROR
	// One or more parameters are illegal.
	ERROR_DS_PARAM_ERROR = 8255

	// ERROR_DS_NOT_SUPPORTED
	// The specified method is not supported.
	ERROR_DS_NOT_SUPPORTED = 8256

	// ERROR_DS_NO_RESULTS_RETURNED
	// No results were returned.
	ERROR_DS_NO_RESULTS_RETURNED = 8257

	// ERROR_DS_CONTROL_NOT_FOUND
	// The specified control is not supported by the server.
	ERROR_DS_CONTROL_NOT_FOUND = 8258

	// ERROR_DS_CLIENT_LOOP
	// A referral loop was detected by the client.
	ERROR_DS_CLIENT_LOOP = 8259

	// ERROR_DS_REFERRAL_LIMIT_EXCEEDED
	// The preset referral limit was exceeded.
	ERROR_DS_REFERRAL_LIMIT_EXCEEDED = 8260

	// ERROR_DS_SORT_CONTROL_MISSING
	// The search requires a SORT control.
	ERROR_DS_SORT_CONTROL_MISSING = 8261

	// ERROR_DS_OFFSET_RANGE_ERROR
	// The search results exceed the offset range specified.
	ERROR_DS_OFFSET_RANGE_ERROR = 8262

	// ERROR_DS_RIDMGR_DISABLED
	// The directory service detected the subsystem that allocates relative identifiers is disabled. This can occur as a protective mechanism when the system determines a significant portion of relative identifiers (RIDs) have been exhausted. Please see http://go.microsoft.com/fwlink/?LinkId=228610 for recommended diagnostic steps and the procedure to re-enable account creation.
	ERROR_DS_RIDMGR_DISABLED = 8263

	// ERROR_DS_ROOT_MUST_BE_NC
	// The root object must be the head of a naming context. The root object cannot have an instantiated parent.
	ERROR_DS_ROOT_MUST_BE_NC = 8301

	// ERROR_DS_ADD_REPLICA_INHIBITED
	// The add replica operation cannot be performed. The naming context must be writeable in order to create the replica.
	ERROR_DS_ADD_REPLICA_INHIBITED = 8302

	// ERROR_DS_ATT_NOT_DEF_IN_SCHEMA
	// A reference to an attribute that is not defined in the schema occurred.
	ERROR_DS_ATT_NOT_DEF_IN_SCHEMA = 8303

	// ERROR_DS_MAX_OBJ_SIZE_EXCEEDED
	// The maximum size of an object has been exceeded.
	ERROR_DS_MAX_OBJ_SIZE_EXCEEDED = 8304

	// ERROR_DS_OBJ_STRING_NAME_EXISTS
	// An attempt was made to add an object to the directory with a name that is already in use.
	ERROR_DS_OBJ_STRING_NAME_EXISTS = 8305

	// ERROR_DS_NO_RDN_DEFINED_IN_SCHEMA
	// An attempt was made to add an object of a class that does not have an RDN defined in the schema.
	ERROR_DS_NO_RDN_DEFINED_IN_SCHEMA = 8306

	// ERROR_DS_RDN_DOESNT_MATCH_SCHEMA
	// An attempt was made to add an object using an RDN that is not the RDN defined in the schema.
	ERROR_DS_RDN_DOESNT_MATCH_SCHEMA = 8307

	// ERROR_DS_NO_REQUESTED_ATTS_FOUND
	// None of the requested attributes were found on the objects.
	ERROR_DS_NO_REQUESTED_ATTS_FOUND = 8308

	// ERROR_DS_USER_BUFFER_TO_SMAL
	// The user buffer is too small.
	ERROR_DS_USER_BUFFER_TO_SMALL = 8309

	// ERROR_DS_ATT_IS_NOT_ON_OBJ
	// The attribute specified in the operation is not present on the object.
	ERROR_DS_ATT_IS_NOT_ON_OBJ = 8310

	// ERROR_DS_ILLEGAL_MOD_OPERATION
	// Illegal modify operation. Some aspect of the modification is not permitted.
	ERROR_DS_ILLEGAL_MOD_OPERATION = 8311

	// ERROR_DS_OBJ_TOO_LARGE
	// The specified object is too large.
	ERROR_DS_OBJ_TOO_LARGE = 8312

	// ERROR_DS_BAD_INSTANCE_TYPE
	// The specified instance type is not valid.
	ERROR_DS_BAD_INSTANCE_TYPE = 8313

	// ERROR_DS_MASTERDSA_REQUIRED
	// The operation must be performed at a master DSA.
	ERROR_DS_MASTERDSA_REQUIRED = 8314

	// ERROR_DS_OBJECT_CLASS_REQUIRED
	// The object class attribute must be specified.
	ERROR_DS_OBJECT_CLASS_REQUIRED = 8315

	// ERROR_DS_MISSING_REQUIRED_ATT
	// A required attribute is missing.
	ERROR_DS_MISSING_REQUIRED_ATT = 8316

	// ERROR_DS_ATT_NOT_DEF_FOR_CLASS
	// An attempt was made to modify an object to include an attribute that is not legal for its class.
	ERROR_DS_ATT_NOT_DEF_FOR_CLASS = 8317

	// ERROR_DS_ATT_ALREADY_EXISTS
	// The specified attribute is already present on the object.
	ERROR_DS_ATT_ALREADY_EXISTS = 8318

	//=8319 unused

	// ERROR_DS_CANT_ADD_ATT_VALUES
	// The specified attribute is not present, or has no values.
	ERROR_DS_CANT_ADD_ATT_VALUES = 8320

	// ERROR_DS_SINGLE_VALUE_CONSTRAINT
	// Multiple values were specified for an attribute that can have only one value.
	ERROR_DS_SINGLE_VALUE_CONSTRAINT = 8321

	// ERROR_DS_RANGE_CONSTRAINT
	// A value for the attribute was not in the acceptable range of values.
	ERROR_DS_RANGE_CONSTRAINT = 8322

	// ERROR_DS_ATT_VAL_ALREADY_EXISTS
	// The specified value already exists.
	ERROR_DS_ATT_VAL_ALREADY_EXISTS = 8323

	// ERROR_DS_CANT_REM_MISSING_ATT
	// The attribute cannot be removed because it is not present on the object.
	ERROR_DS_CANT_REM_MISSING_ATT = 8324

	// ERROR_DS_CANT_REM_MISSING_ATT_VA
	// The attribute value cannot be removed because it is not present on the object.
	ERROR_DS_CANT_REM_MISSING_ATT_VAL = 8325

	// ERROR_DS_ROOT_CANT_BE_SUBREF
	// The specified root object cannot be a subref.
	ERROR_DS_ROOT_CANT_BE_SUBREF = 8326

	// ERROR_DS_NO_CHAINING
	// Chaining is not permitted.
	ERROR_DS_NO_CHAINING = 8327

	// ERROR_DS_NO_CHAINED_EVA
	// Chained evaluation is not permitted.
	ERROR_DS_NO_CHAINED_EVAL = 8328

	// ERROR_DS_NO_PARENT_OBJECT
	// The operation could not be performed because the object's parent is either uninstantiated or deleted.
	ERROR_DS_NO_PARENT_OBJECT = 8329

	// ERROR_DS_PARENT_IS_AN_ALIAS
	// Having a parent that is an alias is not permitted. Aliases are leaf objects.
	ERROR_DS_PARENT_IS_AN_ALIAS = 8330

	// ERROR_DS_CANT_MIX_MASTER_AND_REPS
	// The object and parent must be of the same type, either both masters or both replicas.
	ERROR_DS_CANT_MIX_MASTER_AND_REPS = 8331

	// ERROR_DS_CHILDREN_EXIST
	// The operation cannot be performed because child objects exist. This operation can only be performed on a leaf object.
	ERROR_DS_CHILDREN_EXIST = 8332

	// ERROR_DS_OBJ_NOT_FOUND
	// Directory object not found.
	ERROR_DS_OBJ_NOT_FOUND = 8333

	// ERROR_DS_ALIASED_OBJ_MISSING
	// The aliased object is missing.
	ERROR_DS_ALIASED_OBJ_MISSING = 8334

	// ERROR_DS_BAD_NAME_SYNTAX
	// The object name has bad syntax.
	ERROR_DS_BAD_NAME_SYNTAX = 8335

	// ERROR_DS_ALIAS_POINTS_TO_ALIAS
	// It is not permitted for an alias to refer to another alias.
	ERROR_DS_ALIAS_POINTS_TO_ALIAS = 8336

	// ERROR_DS_CANT_DEREF_ALIAS
	// The alias cannot be dereferenced.
	ERROR_DS_CANT_DEREF_ALIAS = 8337

	// ERROR_DS_OUT_OF_SCOPE
	// The operation is out of scope.
	ERROR_DS_OUT_OF_SCOPE = 8338

	// ERROR_DS_OBJECT_BEING_REMOVED
	// The operation cannot continue because the object is in the process of being removed.
	ERROR_DS_OBJECT_BEING_REMOVED = 8339

	// ERROR_DS_CANT_DELETE_DSA_OBJ
	// The DSA object cannot be deleted.
	ERROR_DS_CANT_DELETE_DSA_OBJ = 8340

	// ERROR_DS_GENERIC_ERROR
	// A directory service error has occurred.
	ERROR_DS_GENERIC_ERROR = 8341

	// ERROR_DS_DSA_MUST_BE_INT_MASTER
	// The operation can only be performed on an internal master DSA object.
	ERROR_DS_DSA_MUST_BE_INT_MASTER = 8342

	// ERROR_DS_CLASS_NOT_DSA
	// The object must be of class DSA.
	ERROR_DS_CLASS_NOT_DSA = 8343

	// ERROR_DS_INSUFF_ACCESS_RIGHTS
	// Insufficient access rights to perform the operation.
	ERROR_DS_INSUFF_ACCESS_RIGHTS = 8344

	// ERROR_DS_ILLEGAL_SUPERIOR
	// The object cannot be added because the parent is not on the list of possible superiors.
	ERROR_DS_ILLEGAL_SUPERIOR = 8345

	// ERROR_DS_ATTRIBUTE_OWNED_BY_SAM
	// Access to the attribute is not permitted because the attribute is owned by the Security Accounts Manager (SAM).
	ERROR_DS_ATTRIBUTE_OWNED_BY_SAM = 8346

	// ERROR_DS_NAME_TOO_MANY_PARTS
	// The name has too many parts.
	ERROR_DS_NAME_TOO_MANY_PARTS = 8347

	// ERROR_DS_NAME_TOO_LONG
	// The name is too long.
	ERROR_DS_NAME_TOO_LONG = 8348

	// ERROR_DS_NAME_VALUE_TOO_LONG
	// The name value is too long.
	ERROR_DS_NAME_VALUE_TOO_LONG = 8349

	// ERROR_DS_NAME_UNPARSEABLE
	// The directory service encountered an error parsing a name.
	ERROR_DS_NAME_UNPARSEABLE = 8350

	// ERROR_DS_NAME_TYPE_UNKNOWN
	// The directory service cannot get the attribute type for a name.
	ERROR_DS_NAME_TYPE_UNKNOWN = 8351

	// ERROR_DS_NOT_AN_OBJECT
	// The name does not identify an object; the name identifies a phantom.
	ERROR_DS_NOT_AN_OBJECT = 8352

	// ERROR_DS_SEC_DESC_TOO_SHORT
	// The security descriptor is too short.
	ERROR_DS_SEC_DESC_TOO_SHORT = 8353

	// ERROR_DS_SEC_DESC_INVALID
	// The security descriptor is invalid.
	ERROR_DS_SEC_DESC_INVALID = 8354

	// ERROR_DS_NO_DELETED_NAME
	// Failed to create name for deleted object.
	ERROR_DS_NO_DELETED_NAME = 8355

	// ERROR_DS_SUBREF_MUST_HAVE_PARENT
	// The parent of a new subref must exist.
	ERROR_DS_SUBREF_MUST_HAVE_PARENT = 8356

	// ERROR_DS_NCNAME_MUST_BE_NC
	// The object must be a naming context.
	ERROR_DS_NCNAME_MUST_BE_NC = 8357

	// ERROR_DS_CANT_ADD_SYSTEM_ONLY
	// It is not permitted to add an attribute which is owned by the system.
	ERROR_DS_CANT_ADD_SYSTEM_ONLY = 8358

	// ERROR_DS_CLASS_MUST_BE_CONCRETE
	// The class of the object must be structural; you cannot instantiate an abstract class.
	ERROR_DS_CLASS_MUST_BE_CONCRETE = 8359

	// ERROR_DS_INVALID_DMD
	// The schema object could not be found.
	ERROR_DS_INVALID_DMD = 8360

	// ERROR_DS_OBJ_GUID_EXISTS
	// A local object with this GUID (dead or alive) already exists.
	ERROR_DS_OBJ_GUID_EXISTS = 8361

	// ERROR_DS_NOT_ON_BACKLINK
	// The operation cannot be performed on a back link.
	ERROR_DS_NOT_ON_BACKLINK = 8362

	// ERROR_DS_NO_CROSSREF_FOR_NC
	// The cross reference for the specified naming context could not be found.
	ERROR_DS_NO_CROSSREF_FOR_NC = 8363

	// ERROR_DS_SHUTTING_DOWN
	// The operation could not be performed because the directory service is shutting down.
	ERROR_DS_SHUTTING_DOWN = 8364

	// ERROR_DS_UNKNOWN_OPERATION
	// The directory service request is invalid.
	ERROR_DS_UNKNOWN_OPERATION = 8365

	// ERROR_DS_INVALID_ROLE_OWNER
	// The role owner attribute could not be read.
	ERROR_DS_INVALID_ROLE_OWNER = 8366

	// ERROR_DS_COULDNT_CONTACT_FSMO
	// The requested FSMO operation failed. The current FSMO holder could not be contacted.
	ERROR_DS_COULDNT_CONTACT_FSMO = 8367

	// ERROR_DS_CROSS_NC_DN_RENAME
	// Modification of a DN across a naming context is not permitted.
	ERROR_DS_CROSS_NC_DN_RENAME = 8368

	// ERROR_DS_CANT_MOD_SYSTEM_ONLY
	// The attribute cannot be modified because it is owned by the system.
	ERROR_DS_CANT_MOD_SYSTEM_ONLY = 8369

	// ERROR_DS_REPLICATOR_ONLY
	// Only the replicator can perform this function.
	ERROR_DS_REPLICATOR_ONLY = 8370

	// ERROR_DS_OBJ_CLASS_NOT_DEFINED
	// The specified class is not defined.
	ERROR_DS_OBJ_CLASS_NOT_DEFINED = 8371

	// ERROR_DS_OBJ_CLASS_NOT_SUBCLASS
	// The specified class is not a subclass.
	ERROR_DS_OBJ_CLASS_NOT_SUBCLASS = 8372

	// ERROR_DS_NAME_REFERENCE_INVALID
	// The name reference is invalid.
	ERROR_DS_NAME_REFERENCE_INVALID = 8373

	// ERROR_DS_CROSS_REF_EXISTS
	// A cross reference already exists.
	ERROR_DS_CROSS_REF_EXISTS = 8374

	// ERROR_DS_CANT_DEL_MASTER_CROSSREF
	// It is not permitted to delete a master cross reference.
	ERROR_DS_CANT_DEL_MASTER_CROSSREF = 8375

	// ERROR_DS_SUBTREE_NOTIFY_NOT_NC_HEAD
	// Subtree notifications are only supported on NC heads.
	ERROR_DS_SUBTREE_NOTIFY_NOT_NC_HEAD = 8376

	// ERROR_DS_NOTIFY_FILTER_TOO_COMPLEX
	// Notification filter is too complex.
	ERROR_DS_NOTIFY_FILTER_TOO_COMPLEX = 8377

	// ERROR_DS_DUP_RDN
	// Schema update failed: duplicate RDN.
	ERROR_DS_DUP_RDN = 8378

	// ERROR_DS_DUP_OID
	// Schema update failed: duplicate OID.
	ERROR_DS_DUP_OID = 8379

	// ERROR_DS_DUP_MAPI_ID
	// Schema update failed: duplicate MAPI identifier.
	ERROR_DS_DUP_MAPI_ID = 8380

	// ERROR_DS_DUP_SCHEMA_ID_GUID
	// Schema update failed: duplicate schema-id GUID.
	ERROR_DS_DUP_SCHEMA_ID_GUID = 8381

	// ERROR_DS_DUP_LDAP_DISPLAY_NAME
	// Schema update failed: duplicate LDAP display name.
	ERROR_DS_DUP_LDAP_DISPLAY_NAME = 8382

	// ERROR_DS_SEMANTIC_ATT_TEST
	// Schema update failed: range-lower less than range upper.
	ERROR_DS_SEMANTIC_ATT_TEST = 8383

	// ERROR_DS_SYNTAX_MISMATCH
	// Schema update failed: syntax mismatch.
	ERROR_DS_SYNTAX_MISMATCH = 8384

	// ERROR_DS_EXISTS_IN_MUST_HAVE
	// Schema deletion failed: attribute is used in must-contain.
	ERROR_DS_EXISTS_IN_MUST_HAVE = 8385

	// ERROR_DS_EXISTS_IN_MAY_HAVE
	// Schema deletion failed: attribute is used in may-contain.
	ERROR_DS_EXISTS_IN_MAY_HAVE = 8386

	// ERROR_DS_NONEXISTENT_MAY_HAVE
	// Schema update failed: attribute in may-contain does not exist.
	ERROR_DS_NONEXISTENT_MAY_HAVE = 8387

	// ERROR_DS_NONEXISTENT_MUST_HAVE
	// Schema update failed: attribute in must-contain does not exist.
	ERROR_DS_NONEXISTENT_MUST_HAVE = 8388

	// ERROR_DS_AUX_CLS_TEST_FAI
	// Schema update failed: class in aux-class list does not exist or is not an auxiliary class.
	ERROR_DS_AUX_CLS_TEST_FAIL = 8389

	// ERROR_DS_NONEXISTENT_POSS_SUP
	// Schema update failed: class in poss-superiors does not exist.
	ERROR_DS_NONEXISTENT_POSS_SUP = 8390

	// ERROR_DS_SUB_CLS_TEST_FAI
	// Schema update failed: class in subclassof list does not exist or does not satisfy hierarchy rules.
	ERROR_DS_SUB_CLS_TEST_FAIL = 8391

	// ERROR_DS_BAD_RDN_ATT_ID_SYNTAX
	// Schema update failed: Rdn-Att-Id has wrong syntax.
	ERROR_DS_BAD_RDN_ATT_ID_SYNTAX = 8392

	// ERROR_DS_EXISTS_IN_AUX_CLS
	// Schema deletion failed: class is used as auxiliary class.
	ERROR_DS_EXISTS_IN_AUX_CLS = 8393

	// ERROR_DS_EXISTS_IN_SUB_CLS
	// Schema deletion failed: class is used as sub class.
	ERROR_DS_EXISTS_IN_SUB_CLS = 8394

	// ERROR_DS_EXISTS_IN_POSS_SUP
	// Schema deletion failed: class is used as poss superior.
	ERROR_DS_EXISTS_IN_POSS_SUP = 8395

	// ERROR_DS_RECALCSCHEMA_FAILED
	// Schema update failed in recalculating validation cache.
	ERROR_DS_RECALCSCHEMA_FAILED = 8396

	// ERROR_DS_TREE_DELETE_NOT_FINISHED
	// The tree deletion is not finished. The request must be made again to continue deleting the tree.
	ERROR_DS_TREE_DELETE_NOT_FINISHED = 8397

	// ERROR_DS_CANT_DELETE
	// The requested delete operation could not be performed.
	ERROR_DS_CANT_DELETE = 8398

	// ERROR_DS_ATT_SCHEMA_REQ_ID
	// Cannot read the governs class identifier for the schema record.
	ERROR_DS_ATT_SCHEMA_REQ_ID = 8399

	// ERROR_DS_BAD_ATT_SCHEMA_SYNTAX
	// The attribute schema has bad syntax.
	ERROR_DS_BAD_ATT_SCHEMA_SYNTAX = 8400

	// ERROR_DS_CANT_CACHE_ATT
	// The attribute could not be cached.
	ERROR_DS_CANT_CACHE_ATT = 8401

	// ERROR_DS_CANT_CACHE_CLASS
	// The class could not be cached.
	ERROR_DS_CANT_CACHE_CLASS = 8402

	// ERROR_DS_CANT_REMOVE_ATT_CACHE
	// The attribute could not be removed from the cache.
	ERROR_DS_CANT_REMOVE_ATT_CACHE = 8403

	// ERROR_DS_CANT_REMOVE_CLASS_CACHE
	// The class could not be removed from the cache.
	ERROR_DS_CANT_REMOVE_CLASS_CACHE = 8404

	// ERROR_DS_CANT_RETRIEVE_DN
	// The distinguished name attribute could not be read.
	ERROR_DS_CANT_RETRIEVE_DN = 8405

	// ERROR_DS_MISSING_SUPREF
	// No superior reference has been configured for the directory service. The directory service is therefore unable to issue referrals to objects outside this forest.
	ERROR_DS_MISSING_SUPREF = 8406

	// ERROR_DS_CANT_RETRIEVE_INSTANCE
	// The instance type attribute could not be retrieved.
	ERROR_DS_CANT_RETRIEVE_INSTANCE = 8407

	// ERROR_DS_CODE_INCONSISTENCY
	// An internal error has occurred.
	ERROR_DS_CODE_INCONSISTENCY = 8408

	// ERROR_DS_DATABASE_ERROR
	// A database error has occurred.
	ERROR_DS_DATABASE_ERROR = 8409

	// ERROR_DS_GOVERNSID_MISSING
	// The attribute GOVERNSID is missing.
	ERROR_DS_GOVERNSID_MISSING = 8410

	// ERROR_DS_MISSING_EXPECTED_ATT
	// An expected attribute is missing.
	ERROR_DS_MISSING_EXPECTED_ATT = 8411

	// ERROR_DS_NCNAME_MISSING_CR_REF
	// The specified naming context is missing a cross reference.
	ERROR_DS_NCNAME_MISSING_CR_REF = 8412

	// ERROR_DS_SECURITY_CHECKING_ERROR
	// A security checking error has occurred.
	ERROR_DS_SECURITY_CHECKING_ERROR = 8413

	// ERROR_DS_SCHEMA_NOT_LOADED
	// The schema is not loaded.
	ERROR_DS_SCHEMA_NOT_LOADED = 8414

	// ERROR_DS_SCHEMA_ALLOC_FAILED
	// Schema allocation failed. Please check if the machine is running low on memory.
	ERROR_DS_SCHEMA_ALLOC_FAILED = 8415

	// ERROR_DS_ATT_SCHEMA_REQ_SYNTAX
	// Failed to obtain the required syntax for the attribute schema.
	ERROR_DS_ATT_SCHEMA_REQ_SYNTAX = 8416

	// ERROR_DS_GCVERIFY_ERROR
	// The global catalog verification failed. The global catalog is not available or does not support the operation. Some part of the directory is currently not available.
	ERROR_DS_GCVERIFY_ERROR = 8417

	// ERROR_DS_DRA_SCHEMA_MISMATCH
	// The replication operation failed because of a schema mismatch between the servers involved.
	ERROR_DS_DRA_SCHEMA_MISMATCH = 8418

	// ERROR_DS_CANT_FIND_DSA_OBJ
	// The DSA object could not be found.
	ERROR_DS_CANT_FIND_DSA_OBJ = 8419

	// ERROR_DS_CANT_FIND_EXPECTED_NC
	// The naming context could not be found.
	ERROR_DS_CANT_FIND_EXPECTED_NC = 8420

	// ERROR_DS_CANT_FIND_NC_IN_CACHE
	// The naming context could not be found in the cache.
	ERROR_DS_CANT_FIND_NC_IN_CACHE = 8421

	// ERROR_DS_CANT_RETRIEVE_CHILD
	// The child object could not be retrieved.
	ERROR_DS_CANT_RETRIEVE_CHILD = 8422

	// ERROR_DS_SECURITY_ILLEGAL_MODIFY
	// The modification was not permitted for security reasons.
	ERROR_DS_SECURITY_ILLEGAL_MODIFY = 8423

	// ERROR_DS_CANT_REPLACE_HIDDEN_REC
	// The operation cannot replace the hidden record.
	ERROR_DS_CANT_REPLACE_HIDDEN_REC = 8424

	// ERROR_DS_BAD_HIERARCHY_FILE
	// The hierarchy file is invalid.
	ERROR_DS_BAD_HIERARCHY_FILE = 8425

	// ERROR_DS_BUILD_HIERARCHY_TABLE_FAILED
	// The attempt to build the hierarchy table failed.
	ERROR_DS_BUILD_HIERARCHY_TABLE_FAILED = 8426

	// ERROR_DS_CONFIG_PARAM_MISSING
	// The directory configuration parameter is missing from the registry.
	ERROR_DS_CONFIG_PARAM_MISSING = 8427

	// ERROR_DS_COUNTING_AB_INDICES_FAILED
	// The attempt to count the address book indices failed.
	ERROR_DS_COUNTING_AB_INDICES_FAILED = 8428

	// ERROR_DS_HIERARCHY_TABLE_MALLOC_FAILED
	// The allocation of the hierarchy table failed.
	ERROR_DS_HIERARCHY_TABLE_MALLOC_FAILED = 8429

	// ERROR_DS_INTERNAL_FAILURE
	// The directory service encountered an internal failure.
	ERROR_DS_INTERNAL_FAILURE = 8430

	// ERROR_DS_UNKNOWN_ERROR
	// The directory service encountered an unknown failure.
	ERROR_DS_UNKNOWN_ERROR = 8431

	// ERROR_DS_ROOT_REQUIRES_CLASS_TOP
	// A root object requires a class of 'top'.
	ERROR_DS_ROOT_REQUIRES_CLASS_TOP = 8432

	// ERROR_DS_REFUSING_FSMO_ROLES
	// This directory server is shutting down, and cannot take ownership of new floating single-master operation roles.
	ERROR_DS_REFUSING_FSMO_ROLES = 8433

	// ERROR_DS_MISSING_FSMO_SETTINGS
	// The directory service is missing mandatory configuration information, and is unable to determine the ownership of floating single-master operation roles.
	ERROR_DS_MISSING_FSMO_SETTINGS = 8434

	// ERROR_DS_UNABLE_TO_SURRENDER_ROLES
	// The directory service was unable to transfer ownership of one or more floating single-master operation roles to other servers.
	ERROR_DS_UNABLE_TO_SURRENDER_ROLES = 8435

	// ERROR_DS_DRA_GENERIC
	// The replication operation failed.
	ERROR_DS_DRA_GENERIC = 8436

	// ERROR_DS_DRA_INVALID_PARAMETER
	// An invalid parameter was specified for this replication operation.
	ERROR_DS_DRA_INVALID_PARAMETER = 8437

	// ERROR_DS_DRA_BUSY
	// The directory service is too busy to complete the replication operation at this time.
	ERROR_DS_DRA_BUSY = 8438

	// ERROR_DS_DRA_BAD_DN
	// The distinguished name specified for this replication operation is invalid.
	ERROR_DS_DRA_BAD_DN = 8439

	// ERROR_DS_DRA_BAD_NC
	// The naming context specified for this replication operation is invalid.
	ERROR_DS_DRA_BAD_NC = 8440

	// ERROR_DS_DRA_DN_EXISTS
	// The distinguished name specified for this replication operation already exists.
	ERROR_DS_DRA_DN_EXISTS = 8441

	// ERROR_DS_DRA_INTERNAL_ERROR
	// The replication system encountered an internal error.
	ERROR_DS_DRA_INTERNAL_ERROR = 8442

	// ERROR_DS_DRA_INCONSISTENT_DIT
	// The replication operation encountered a database inconsistency.
	ERROR_DS_DRA_INCONSISTENT_DIT = 8443

	// ERROR_DS_DRA_CONNECTION_FAILED
	// The server specified for this replication operation could not be contacted.
	ERROR_DS_DRA_CONNECTION_FAILED = 8444

	// ERROR_DS_DRA_BAD_INSTANCE_TYPE
	// The replication operation encountered an object with an invalid instance type.
	ERROR_DS_DRA_BAD_INSTANCE_TYPE = 8445

	// ERROR_DS_DRA_OUT_OF_MEM
	// The replication operation failed to allocate memory.
	ERROR_DS_DRA_OUT_OF_MEM = 8446

	// ERROR_DS_DRA_MAIL_PROBLEM
	// The replication operation encountered an error with the mail system.
	ERROR_DS_DRA_MAIL_PROBLEM = 8447

	// ERROR_DS_DRA_REF_ALREADY_EXISTS
	// The replication reference information for the target server already exists.
	ERROR_DS_DRA_REF_ALREADY_EXISTS = 8448

	// ERROR_DS_DRA_REF_NOT_FOUND
	// The replication reference information for the target server does not exist.
	ERROR_DS_DRA_REF_NOT_FOUND = 8449

	// ERROR_DS_DRA_OBJ_IS_REP_SOURCE
	// The naming context cannot be removed because it is replicated to another server.
	ERROR_DS_DRA_OBJ_IS_REP_SOURCE = 8450

	// ERROR_DS_DRA_DB_ERROR
	// The replication operation encountered a database error.
	ERROR_DS_DRA_DB_ERROR = 8451

	// ERROR_DS_DRA_NO_REPLICA
	// The naming context is in the process of being removed or is not replicated from the specified server.
	ERROR_DS_DRA_NO_REPLICA = 8452

	// ERROR_DS_DRA_ACCESS_DENIED
	// Replication access was denied.
	ERROR_DS_DRA_ACCESS_DENIED = 8453

	// ERROR_DS_DRA_NOT_SUPPORTED
	// The requested operation is not supported by this version of the directory service.
	ERROR_DS_DRA_NOT_SUPPORTED = 8454

	// ERROR_DS_DRA_RPC_CANCELLED
	// The replication remote procedure call was cancelled.
	ERROR_DS_DRA_RPC_CANCELLED = 8455

	// ERROR_DS_DRA_SOURCE_DISABLED
	// The source server is currently rejecting replication requests.
	ERROR_DS_DRA_SOURCE_DISABLED = 8456

	// ERROR_DS_DRA_SINK_DISABLED
	// The destination server is currently rejecting replication requests.
	ERROR_DS_DRA_SINK_DISABLED = 8457

	// ERROR_DS_DRA_NAME_COLLISION
	// The replication operation failed due to a collision of object names.
	ERROR_DS_DRA_NAME_COLLISION = 8458

	// ERROR_DS_DRA_SOURCE_REINSTALLED
	// The replication source has been reinstalled.
	ERROR_DS_DRA_SOURCE_REINSTALLED = 8459

	// ERROR_DS_DRA_MISSING_PARENT
	// The replication operation failed because a required parent object is missing.
	ERROR_DS_DRA_MISSING_PARENT = 8460

	// ERROR_DS_DRA_PREEMPTED
	// The replication operation was preempted.
	ERROR_DS_DRA_PREEMPTED = 8461

	// ERROR_DS_DRA_ABANDON_SYNC
	// The replication synchronization attempt was abandoned because of a lack of updates.
	ERROR_DS_DRA_ABANDON_SYNC = 8462

	// ERROR_DS_DRA_SHUTDOWN
	// The replication operation was terminated because the system is shutting down.
	ERROR_DS_DRA_SHUTDOWN = 8463

	// ERROR_DS_DRA_INCOMPATIBLE_PARTIAL_SET
	// Synchronization attempt failed because the destination DC is currently waiting to synchronize new partial attributes from source. This condition is normal if a recent schema change modified the partial attribute set. The destination partial attribute set is not a subset of source partial attribute set.
	ERROR_DS_DRA_INCOMPATIBLE_PARTIAL_SET = 8464

	// ERROR_DS_DRA_SOURCE_IS_PARTIAL_REPLICA
	// The replication synchronization attempt failed because a master replica attempted to sync from a partial replica.
	ERROR_DS_DRA_SOURCE_IS_PARTIAL_REPLICA = 8465

	// ERROR_DS_DRA_EXTN_CONNECTION_FAILED
	// The server specified for this replication operation was contacted, but that server was unable to contact an additional server needed to complete the operation.
	ERROR_DS_DRA_EXTN_CONNECTION_FAILED = 8466

	// ERROR_DS_INSTALL_SCHEMA_MISMATCH
	// The version of the directory service schema of the source forest is not compatible with the version of directory service on this computer.
	ERROR_DS_INSTALL_SCHEMA_MISMATCH = 8467

	// ERROR_DS_DUP_LINK_ID
	// Schema update failed: An attribute with the same link identifier already exists.
	ERROR_DS_DUP_LINK_ID = 8468

	// ERROR_DS_NAME_ERROR_RESOLVING
	// Name translation: Generic processing error.
	ERROR_DS_NAME_ERROR_RESOLVING = 8469

	// ERROR_DS_NAME_ERROR_NOT_FOUND
	// Name translation: Could not find the name or insufficient right to see name.
	ERROR_DS_NAME_ERROR_NOT_FOUND = 8470

	// ERROR_DS_NAME_ERROR_NOT_UNIQUE
	// Name translation: Input name mapped to more than one output name.
	ERROR_DS_NAME_ERROR_NOT_UNIQUE = 8471

	// ERROR_DS_NAME_ERROR_NO_MAPPING
	// Name translation: Input name found, but not the associated output format.
	ERROR_DS_NAME_ERROR_NO_MAPPING = 8472

	// ERROR_DS_NAME_ERROR_DOMAIN_ONLY
	// Name translation: Unable to resolve completely, only the domain was found.
	ERROR_DS_NAME_ERROR_DOMAIN_ONLY = 8473

	// ERROR_DS_NAME_ERROR_NO_SYNTACTICAL_MAPPING
	// Name translation: Unable to perform purely syntactical mapping at the client without going out to the wire.
	ERROR_DS_NAME_ERROR_NO_SYNTACTICAL_MAPPING = 8474

	// ERROR_DS_CONSTRUCTED_ATT_MOD
	// Modification of a constructed attribute is not allowed.
	ERROR_DS_CONSTRUCTED_ATT_MOD = 8475

	// ERROR_DS_WRONG_OM_OBJ_CLASS
	// The OM-Object-Class specified is incorrect for an attribute with the specified syntax.
	ERROR_DS_WRONG_OM_OBJ_CLASS = 8476

	// ERROR_DS_DRA_REPL_PENDING
	// The replication request has been posted; waiting for reply.
	ERROR_DS_DRA_REPL_PENDING = 8477

	// ERROR_DS_DS_REQUIRED
	// The requested operation requires a directory service, and none was available.
	ERROR_DS_DS_REQUIRED = 8478

	// ERROR_DS_INVALID_LDAP_DISPLAY_NAME
	// The LDAP display name of the class or attribute contains non-ASCII characters.
	ERROR_DS_INVALID_LDAP_DISPLAY_NAME = 8479

	// ERROR_DS_NON_BASE_SEARCH
	// The requested search operation is only supported for base searches.
	ERROR_DS_NON_BASE_SEARCH = 8480

	// ERROR_DS_CANT_RETRIEVE_ATTS
	// The search failed to retrieve attributes from the database.
	ERROR_DS_CANT_RETRIEVE_ATTS = 8481

	// ERROR_DS_BACKLINK_WITHOUT_LINK
	// The schema update operation tried to add a backward link attribute that has no corresponding forward link.
	ERROR_DS_BACKLINK_WITHOUT_LINK = 8482

	// ERROR_DS_EPOCH_MISMATCH
	// Source and destination of a cross-domain move do not agree on the object's epoch number. Either source or destination does not have the latest version of the object.
	ERROR_DS_EPOCH_MISMATCH = 8483

	// ERROR_DS_SRC_NAME_MISMATCH
	// Source and destination of a cross-domain move do not agree on the object's current name. Either source or destination does not have the latest version of the object.
	ERROR_DS_SRC_NAME_MISMATCH = 8484

	// ERROR_DS_SRC_AND_DST_NC_IDENTICA
	// Source and destination for the cross-domain move operation are identical. Caller should use local move operation instead of cross-domain move operation.
	ERROR_DS_SRC_AND_DST_NC_IDENTICAL = 8485

	// ERROR_DS_DST_NC_MISMATCH
	// Source and destination for a cross-domain move are not in agreement on the naming contexts in the forest. Either source or destination does not have the latest version of the Partitions container.
	ERROR_DS_DST_NC_MISMATCH = 8486

	// ERROR_DS_NOT_AUTHORITIVE_FOR_DST_NC
	// Destination of a cross-domain move is not authoritative for the destination naming context.
	ERROR_DS_NOT_AUTHORITIVE_FOR_DST_NC = 8487

	// ERROR_DS_SRC_GUID_MISMATCH
	// Source and destination of a cross-domain move do not agree on the identity of the source object. Either source or destination does not have the latest version of the source object.
	ERROR_DS_SRC_GUID_MISMATCH = 8488

	// ERROR_DS_CANT_MOVE_DELETED_OBJECT
	// Object being moved across-domains is already known to be deleted by the destination server. The source server does not have the latest version of the source object.
	ERROR_DS_CANT_MOVE_DELETED_OBJECT = 8489

	// ERROR_DS_PDC_OPERATION_IN_PROGRESS
	// Another operation which requires exclusive access to the PDC FSMO is already in progress.
	ERROR_DS_PDC_OPERATION_IN_PROGRESS = 8490

	// ERROR_DS_CROSS_DOMAIN_CLEANUP_REQD
	// A cross-domain move operation failed such that two versions of the moved object exist - one each in the source and destination domains. The destination object needs to be removed to restore the system to a consistent state.
	ERROR_DS_CROSS_DOMAIN_CLEANUP_REQD = 8491

	// ERROR_DS_ILLEGAL_XDOM_MOVE_OPERATION
	// This object may not be moved across domain boundaries either because cross-domain moves for this class are disallowed, or the object has some special characteristics, e.g.: trust account or restricted RID, which prevent its move.
	ERROR_DS_ILLEGAL_XDOM_MOVE_OPERATION = 8492

	// ERROR_DS_CANT_WITH_ACCT_GROUP_MEMBERSHPS
	// Can't move objects with memberships across domain boundaries as once moved, this would violate the membership conditions of the account group. Remove the object from any account group memberships and retry.
	ERROR_DS_CANT_WITH_ACCT_GROUP_MEMBERSHPS = 8493

	// ERROR_DS_NC_MUST_HAVE_NC_PARENT
	// A naming context head must be the immediate child of another naming context head, not of an interior node.
	ERROR_DS_NC_MUST_HAVE_NC_PARENT = 8494

	// ERROR_DS_CR_IMPOSSIBLE_TO_VALIDATE
	// The directory cannot validate the proposed naming context name because it does not hold a replica of the naming context above the proposed naming context. Please ensure that the domain naming master role is held by a server that is configured as a global catalog server, and that the server is up to date with its replication partners. (Applies only to Windows=2000 Domain Naming masters)
	ERROR_DS_CR_IMPOSSIBLE_TO_VALIDATE = 8495

	// ERROR_DS_DST_DOMAIN_NOT_NATIVE
	// Destination domain must be in native mode.
	ERROR_DS_DST_DOMAIN_NOT_NATIVE = 8496

	// ERROR_DS_MISSING_INFRASTRUCTURE_CONTAINER
	// The operation cannot be performed because the server does not have an infrastructure container in the domain of interest.
	ERROR_DS_MISSING_INFRASTRUCTURE_CONTAINER = 8497

	// ERROR_DS_CANT_MOVE_ACCOUNT_GROUP
	// Cross-domain move of non-empty account groups is not allowed.
	ERROR_DS_CANT_MOVE_ACCOUNT_GROUP = 8498

	// ERROR_DS_CANT_MOVE_RESOURCE_GROUP
	// Cross-domain move of non-empty resource groups is not allowed.
	ERROR_DS_CANT_MOVE_RESOURCE_GROUP = 8499

	// ERROR_DS_INVALID_SEARCH_FLAG
	// The search flags for the attribute are invalid. The ANR bit is valid only on attributes of Unicode or Teletex strings.
	ERROR_DS_INVALID_SEARCH_FLAG = 8500

	// ERROR_DS_NO_TREE_DELETE_ABOVE_NC
	// Tree deletions starting at an object which has an NC head as a descendant are not allowed.
	ERROR_DS_NO_TREE_DELETE_ABOVE_NC = 8501

	// ERROR_DS_COULDNT_LOCK_TREE_FOR_DELETE
	// The directory service failed to lock a tree in preparation for a tree deletion because the tree was in use.
	ERROR_DS_COULDNT_LOCK_TREE_FOR_DELETE = 8502

	// ERROR_DS_COULDNT_IDENTIFY_OBJECTS_FOR_TREE_DELETE
	// The directory service failed to identify the list of objects to delete while attempting a tree deletion.
	ERROR_DS_COULDNT_IDENTIFY_OBJECTS_FOR_TREE_DELETE = 8503

	// ERROR_DS_SAM_INIT_FAILURE
	// Security Accounts Manager initialization failed because of the following error: %1.
	// Error Status:=0x%2. Please shutdown this system and reboot into Directory Services Restore Mode, check the event log for more detailed information.
	ERROR_DS_SAM_INIT_FAILURE = 8504

	// ERROR_DS_SENSITIVE_GROUP_VIOLATION
	// Only an administrator can modify the membership list of an administrative group.
	ERROR_DS_SENSITIVE_GROUP_VIOLATION = 8505

	// ERROR_DS_CANT_MOD_PRIMARYGROUPID
	// Cannot change the primary group ID of a domain controller account.
	ERROR_DS_CANT_MOD_PRIMARYGROUPID = 8506

	// ERROR_DS_ILLEGAL_BASE_SCHEMA_MOD
	// An attempt is made to modify the base schema.
	ERROR_DS_ILLEGAL_BASE_SCHEMA_MOD = 8507

	// ERROR_DS_NONSAFE_SCHEMA_CHANGE
	// Adding a new mandatory attribute to an existing class, deleting a mandatory attribute from an existing class, or adding an optional attribute to the special class Top that is not a backlink attribute (directly or through inheritance, for example, by adding or deleting an auxiliary class) is not allowed.
	ERROR_DS_NONSAFE_SCHEMA_CHANGE = 8508

	// ERROR_DS_SCHEMA_UPDATE_DISALLOWED
	// Schema update is not allowed on this DC because the DC is not the schema FSMO Role Owner.
	ERROR_DS_SCHEMA_UPDATE_DISALLOWED = 8509

	// ERROR_DS_CANT_CREATE_UNDER_SCHEMA
	// An object of this class cannot be created under the schema container. You can only create attribute-schema and class-schema objects under the schema container.
	ERROR_DS_CANT_CREATE_UNDER_SCHEMA = 8510

	// ERROR_DS_INSTALL_NO_SRC_SCH_VERSION
	// The replica/child install failed to get the objectVersion attribute on the schema container on the source DC. Either the attribute is missing on the schema container or the credentials supplied do not have permission to read it.
	ERROR_DS_INSTALL_NO_SRC_SCH_VERSION = 8511

	// ERROR_DS_INSTALL_NO_SCH_VERSION_IN_INIFILE
	// The replica/child install failed to read the objectVersion attribute in the SCHEMA section of the file schema.ini in the system32 directory.
	ERROR_DS_INSTALL_NO_SCH_VERSION_IN_INIFILE = 8512

	// ERROR_DS_INVALID_GROUP_TYPE
	// The specified group type is invalid.
	ERROR_DS_INVALID_GROUP_TYPE = 8513

	// ERROR_DS_NO_NEST_GLOBALGROUP_IN_MIXEDDOMAIN
	// You cannot nest global groups in a mixed domain if the group is security-enabled.
	ERROR_DS_NO_NEST_GLOBALGROUP_IN_MIXEDDOMAIN = 8514

	// ERROR_DS_NO_NEST_LOCALGROUP_IN_MIXEDDOMAIN
	// You cannot nest local groups in a mixed domain if the group is security-enabled.
	ERROR_DS_NO_NEST_LOCALGROUP_IN_MIXEDDOMAIN = 8515

	// ERROR_DS_GLOBAL_CANT_HAVE_LOCAL_MEMBER
	// A global group cannot have a local group as a member.
	ERROR_DS_GLOBAL_CANT_HAVE_LOCAL_MEMBER = 8516

	// ERROR_DS_GLOBAL_CANT_HAVE_UNIVERSAL_MEMBER
	// A global group cannot have a universal group as a member.
	ERROR_DS_GLOBAL_CANT_HAVE_UNIVERSAL_MEMBER = 8517

	// ERROR_DS_UNIVERSAL_CANT_HAVE_LOCAL_MEMBER
	// A universal group cannot have a local group as a member.
	ERROR_DS_UNIVERSAL_CANT_HAVE_LOCAL_MEMBER = 8518

	// ERROR_DS_GLOBAL_CANT_HAVE_CROSSDOMAIN_MEMBER
	// A global group cannot have a cross-domain member.
	ERROR_DS_GLOBAL_CANT_HAVE_CROSSDOMAIN_MEMBER = 8519

	// ERROR_DS_LOCAL_CANT_HAVE_CROSSDOMAIN_LOCAL_MEMBER
	// A local group cannot have another cross domain local group as a member.
	ERROR_DS_LOCAL_CANT_HAVE_CROSSDOMAIN_LOCAL_MEMBER = 8520

	// ERROR_DS_HAVE_PRIMARY_MEMBERS
	// A group with primary members cannot change to a security-disabled group.
	ERROR_DS_HAVE_PRIMARY_MEMBERS = 8521

	// ERROR_DS_STRING_SD_CONVERSION_FAILED
	// The schema cache load failed to convert the string default SD on a class-schema object.
	ERROR_DS_STRING_SD_CONVERSION_FAILED = 8522

	// ERROR_DS_NAMING_MASTER_GC
	// Only DSAs configured to be Global Catalog servers should be allowed to hold the Domain Naming Master FSMO role. (Applies only to Windows=2000 servers)
	ERROR_DS_NAMING_MASTER_GC = 8523

	// ERROR_DS_DNS_LOOKUP_FAILURE
	// The DSA operation is unable to proceed because of a DNS lookup failure.
	ERROR_DS_DNS_LOOKUP_FAILURE = 8524

	// ERROR_DS_COULDNT_UPDATE_SPNS
	// While processing a change to the DNS Host Name for an object, the Service Principal Name values could not be kept in sync.
	ERROR_DS_COULDNT_UPDATE_SPNS = 8525

	// ERROR_DS_CANT_RETRIEVE_SD
	// The Security Descriptor attribute could not be read.
	ERROR_DS_CANT_RETRIEVE_SD = 8526

	// ERROR_DS_KEY_NOT_UNIQUE
	// The object requested was not found, but an object with that key was found.
	ERROR_DS_KEY_NOT_UNIQUE = 8527

	// ERROR_DS_WRONG_LINKED_ATT_SYNTAX
	// The syntax of the linked attribute being added is incorrect. Forward links can only have syntax=2.5.5.1,=2.5.5.7, and=2.5.5.14, and backlinks can only have syntax=2.5.5.1
	ERROR_DS_WRONG_LINKED_ATT_SYNTAX = 8528

	// ERROR_DS_SAM_NEED_BOOTKEY_PASSWORD
	// Security Account Manager needs to get the boot password.
	ERROR_DS_SAM_NEED_BOOTKEY_PASSWORD = 8529

	// ERROR_DS_SAM_NEED_BOOTKEY_FLOPPY
	// Security Account Manager needs to get the boot key from floppy disk.
	ERROR_DS_SAM_NEED_BOOTKEY_FLOPPY = 8530

	// ERROR_DS_CANT_START
	// Directory Service cannot start.
	ERROR_DS_CANT_START = 8531

	// ERROR_DS_INIT_FAILURE
	// Directory Services could not start.
	ERROR_DS_INIT_FAILURE = 8532

	// ERROR_DS_NO_PKT_PRIVACY_ON_CONNECTION
	// The connection between client and server requires packet privacy or better.
	ERROR_DS_NO_PKT_PRIVACY_ON_CONNECTION = 8533

	// ERROR_DS_SOURCE_DOMAIN_IN_FOREST
	// The source domain may not be in the same forest as destination.
	ERROR_DS_SOURCE_DOMAIN_IN_FOREST = 8534

	// ERROR_DS_DESTINATION_DOMAIN_NOT_IN_FOREST
	// The destination domain must be in the forest.
	ERROR_DS_DESTINATION_DOMAIN_NOT_IN_FOREST = 8535

	// ERROR_DS_DESTINATION_AUDITING_NOT_ENABLED
	// The operation requires that destination domain auditing be enabled.
	ERROR_DS_DESTINATION_AUDITING_NOT_ENABLED = 8536

	// ERROR_DS_CANT_FIND_DC_FOR_SRC_DOMAIN
	// The operation couldn't locate a DC for the source domain.
	ERROR_DS_CANT_FIND_DC_FOR_SRC_DOMAIN = 8537

	// ERROR_DS_SRC_OBJ_NOT_GROUP_OR_USER
	// The source object must be a group or user.
	ERROR_DS_SRC_OBJ_NOT_GROUP_OR_USER = 8538

	// ERROR_DS_SRC_SID_EXISTS_IN_FOREST
	// The source object's SID already exists in destination forest.
	ERROR_DS_SRC_SID_EXISTS_IN_FOREST = 8539

	// ERROR_DS_SRC_AND_DST_OBJECT_CLASS_MISMATCH
	// The source and destination object must be of the same type.
	ERROR_DS_SRC_AND_DST_OBJECT_CLASS_MISMATCH = 8540

	// ERROR_SAM_INIT_FAILURE
	// Security Accounts Manager initialization failed because of the following error: %1.
	// Error Status:=0x%2. Click OK to shut down the system and reboot into Safe Mode. Check the event log for detailed information.
	ERROR_SAM_INIT_FAILURE = 8541

	// ERROR_DS_DRA_SCHEMA_INFO_SHIP
	// Schema information could not be included in the replication request.
	ERROR_DS_DRA_SCHEMA_INFO_SHIP = 8542

	// ERROR_DS_DRA_SCHEMA_CONFLICT
	// The replication operation could not be completed due to a schema incompatibility.
	ERROR_DS_DRA_SCHEMA_CONFLICT = 8543

	// ERROR_DS_DRA_EARLIER_SCHEMA_CONFLICT
	// The replication operation could not be completed due to a previous schema incompatibility.
	ERROR_DS_DRA_EARLIER_SCHEMA_CONFLICT = 8544

	// ERROR_DS_DRA_OBJ_NC_MISMATCH
	// The replication update could not be applied because either the source or the destination has not yet received information regarding a recent cross-domain move operation.
	ERROR_DS_DRA_OBJ_NC_MISMATCH = 8545

	// ERROR_DS_NC_STILL_HAS_DSAS
	// The requested domain could not be deleted because there exist domain controllers that still host this domain.
	ERROR_DS_NC_STILL_HAS_DSAS = 8546

	// ERROR_DS_GC_REQUIRED
	// The requested operation can be performed only on a global catalog server.
	ERROR_DS_GC_REQUIRED = 8547

	// ERROR_DS_LOCAL_MEMBER_OF_LOCAL_ONLY
	// A local group can only be a member of other local groups in the same domain.
	ERROR_DS_LOCAL_MEMBER_OF_LOCAL_ONLY = 8548

	// ERROR_DS_NO_FPO_IN_UNIVERSAL_GROUPS
	// Foreign security principals cannot be members of universal groups.
	ERROR_DS_NO_FPO_IN_UNIVERSAL_GROUPS = 8549

	// ERROR_DS_CANT_ADD_TO_GC
	// The attribute is not allowed to be replicated to the GC because of security reasons.
	ERROR_DS_CANT_ADD_TO_GC = 8550

	// ERROR_DS_NO_CHECKPOINT_WITH_PDC
	// The checkpoint with the PDC could not be taken because there too many modifications being processed currently.
	ERROR_DS_NO_CHECKPOINT_WITH_PDC = 8551

	// ERROR_DS_SOURCE_AUDITING_NOT_ENABLED
	// The operation requires that source domain auditing be enabled.
	ERROR_DS_SOURCE_AUDITING_NOT_ENABLED = 8552

	// ERROR_DS_CANT_CREATE_IN_NONDOMAIN_NC
	// Security principal objects can only be created inside domain naming contexts.
	ERROR_DS_CANT_CREATE_IN_NONDOMAIN_NC = 8553

	// ERROR_DS_INVALID_NAME_FOR_SPN
	// A Service Principal Name (SPN) could not be constructed because the provided hostname is not in the necessary format.
	ERROR_DS_INVALID_NAME_FOR_SPN = 8554

	// ERROR_DS_FILTER_USES_CONTRUCTED_ATTRS
	// A Filter was passed that uses constructed attributes.
	ERROR_DS_FILTER_USES_CONTRUCTED_ATTRS = 8555

	// ERROR_DS_UNICODEPWD_NOT_IN_QUOTES
	// The unicodePwd attribute value must be enclosed in double quotes.
	ERROR_DS_UNICODEPWD_NOT_IN_QUOTES = 8556

	// ERROR_DS_MACHINE_ACCOUNT_QUOTA_EXCEEDED
	// Your computer could not be joined to the domain. You have exceeded the maximum number of computer accounts you are allowed to create in this domain. Contact your system administrator to have this limit reset or increased.
	ERROR_DS_MACHINE_ACCOUNT_QUOTA_EXCEEDED = 8557

	// ERROR_DS_MUST_BE_RUN_ON_DST_DC
	// For security reasons, the operation must be run on the destination DC.
	ERROR_DS_MUST_BE_RUN_ON_DST_DC = 8558

	// ERROR_DS_SRC_DC_MUST_BE_SP4_OR_GREATER
	// For security reasons, the source DC must be NT4SP4 or greater.
	ERROR_DS_SRC_DC_MUST_BE_SP4_OR_GREATER = 8559

	// ERROR_DS_CANT_TREE_DELETE_CRITICAL_OBJ
	// Critical Directory Service System objects cannot be deleted during tree delete operations. The tree delete may have been partially performed.
	ERROR_DS_CANT_TREE_DELETE_CRITICAL_OBJ = 8560

	// ERROR_DS_INIT_FAILURE_CONSOLE
	// Directory Services could not start because of the following error: %1.
	// Error Status:=0x%2. Please click OK to shutdown the system. You can use the recovery console to diagnose the system further.
	ERROR_DS_INIT_FAILURE_CONSOLE = 8561

	// ERROR_DS_SAM_INIT_FAILURE_CONSOLE
	// Security Accounts Manager initialization failed because of the following error: %1.
	// Error Status:=0x%2. Please click OK to shutdown the system. You can use the recovery console to diagnose the system further.
	ERROR_DS_SAM_INIT_FAILURE_CONSOLE = 8562

	// ERROR_DS_FOREST_VERSION_TOO_HIGH
	// The version of the operating system is incompatible with the current AD DS forest functional level or AD LDS Configuration Set functional level. You must upgrade to a new version of the operating system before this server can become an AD DS Domain Controller or add an AD LDS Instance in this AD DS Forest or AD LDS Configuration Set.
	ERROR_DS_FOREST_VERSION_TOO_HIGH = 8563

	// ERROR_DS_DOMAIN_VERSION_TOO_HIGH
	// The version of the operating system installed is incompatible with the current domain functional level. You must upgrade to a new version of the operating system before this server can become a domain controller in this domain.
	ERROR_DS_DOMAIN_VERSION_TOO_HIGH = 8564

	// ERROR_DS_FOREST_VERSION_TOO_LOW
	// The version of the operating system installed on this server no longer supports the current AD DS Forest functional level or AD LDS Configuration Set functional level. You must raise the AD DS Forest functional level or AD LDS Configuration Set functional level before this server can become an AD DS Domain Controller or an AD LDS Instance in this Forest or Configuration Set.
	ERROR_DS_FOREST_VERSION_TOO_LOW = 8565

	// ERROR_DS_DOMAIN_VERSION_TOO_LOW
	// The version of the operating system installed on this server no longer supports the current domain functional level. You must raise the domain functional level before this server can become a domain controller in this domain.
	ERROR_DS_DOMAIN_VERSION_TOO_LOW = 8566

	// ERROR_DS_INCOMPATIBLE_VERSION
	// The version of the operating system installed on this server is incompatible with the functional level of the domain or forest.
	ERROR_DS_INCOMPATIBLE_VERSION = 8567

	// ERROR_DS_LOW_DSA_VERSION
	// The functional level of the domain (or forest) cannot be raised to the requested value, because there exist one or more domain controllers in the domain (or forest) that are at a lower incompatible functional level.
	ERROR_DS_LOW_DSA_VERSION = 8568

	// ERROR_DS_NO_BEHAVIOR_VERSION_IN_MIXEDDOMAIN
	// The forest functional level cannot be raised to the requested value since one or more domains are still in mixed domain mode. All domains in the forest must be in native mode, for you to raise the forest functional level.
	ERROR_DS_NO_BEHAVIOR_VERSION_IN_MIXEDDOMAIN = 8569

	// ERROR_DS_NOT_SUPPORTED_SORT_ORDER
	// The sort order requested is not supported.
	ERROR_DS_NOT_SUPPORTED_SORT_ORDER = 8570

	// ERROR_DS_NAME_NOT_UNIQUE
	// The requested name already exists as a unique identifier.
	ERROR_DS_NAME_NOT_UNIQUE = 8571

	// ERROR_DS_MACHINE_ACCOUNT_CREATED_PRENT4
	// The machine account was created pre-NT4. The account needs to be recreated.
	ERROR_DS_MACHINE_ACCOUNT_CREATED_PRENT4 = 8572

	// ERROR_DS_OUT_OF_VERSION_STORE
	// The database is out of version store.
	ERROR_DS_OUT_OF_VERSION_STORE = 8573

	// ERROR_DS_INCOMPATIBLE_CONTROLS_USED
	// Unable to continue operation because multiple conflicting controls were used.
	ERROR_DS_INCOMPATIBLE_CONTROLS_USED = 8574

	// ERROR_DS_NO_REF_DOMAIN
	// Unable to find a valid security descriptor reference domain for this partition.
	ERROR_DS_NO_REF_DOMAIN = 8575

	// ERROR_DS_RESERVED_LINK_ID
	// Schema update failed: The link identifier is reserved.
	ERROR_DS_RESERVED_LINK_ID = 8576

	// ERROR_DS_LINK_ID_NOT_AVAILABLE
	// Schema update failed: There are no link identifiers available.
	ERROR_DS_LINK_ID_NOT_AVAILABLE = 8577

	// ERROR_DS_AG_CANT_HAVE_UNIVERSAL_MEMBER
	// An account group cannot have a universal group as a member.
	ERROR_DS_AG_CANT_HAVE_UNIVERSAL_MEMBER = 8578

	// ERROR_DS_MODIFYDN_DISALLOWED_BY_INSTANCE_TYPE
	// Rename or move operations on naming context heads or read-only objects are not allowed.
	ERROR_DS_MODIFYDN_DISALLOWED_BY_INSTANCE_TYPE = 8579

	// ERROR_DS_NO_OBJECT_MOVE_IN_SCHEMA_NC
	// Move operations on objects in the schema naming context are not allowed.
	ERROR_DS_NO_OBJECT_MOVE_IN_SCHEMA_NC = 8580

	// ERROR_DS_MODIFYDN_DISALLOWED_BY_FLAG
	// A system flag has been set on the object and does not allow the object to be moved or renamed.
	ERROR_DS_MODIFYDN_DISALLOWED_BY_FLAG = 8581

	// ERROR_DS_MODIFYDN_WRONG_GRANDPARENT
	// This object is not allowed to change its grandparent container. Moves are not forbidden on this object, but are restricted to sibling containers.
	ERROR_DS_MODIFYDN_WRONG_GRANDPARENT = 8582

	// ERROR_DS_NAME_ERROR_TRUST_REFERRA
	// Unable to resolve completely, a referral to another forest is generated.
	ERROR_DS_NAME_ERROR_TRUST_REFERRAL = 8583

	// ERROR_NOT_SUPPORTED_ON_STANDARD_SERVER
	// The requested action is not supported on standard server.
	ERROR_NOT_SUPPORTED_ON_STANDARD_SERVER = 8584

	// ERROR_DS_CANT_ACCESS_REMOTE_PART_OF_AD
	// Could not access a partition of the directory service located on a remote server. Make sure at least one server is running for the partition in question.
	ERROR_DS_CANT_ACCESS_REMOTE_PART_OF_AD = 8585

	// ERROR_DS_CR_IMPOSSIBLE_TO_VALIDATE_V2
	// The directory cannot validate the proposed naming context (or partition) name because it does not hold a replica nor can it contact a replica of the naming context above the proposed naming context. Please ensure that the parent naming context is properly registered in DNS, and at least one replica of this naming context is reachable by the Domain Naming master.
	ERROR_DS_CR_IMPOSSIBLE_TO_VALIDATE_V2 = 8586

	// ERROR_DS_THREAD_LIMIT_EXCEEDED
	// The thread limit for this request was exceeded.
	ERROR_DS_THREAD_LIMIT_EXCEEDED = 8587

	// ERROR_DS_NOT_CLOSEST
	// The Global catalog server is not in the closest site.
	ERROR_DS_NOT_CLOSEST = 8588

	// ERROR_DS_CANT_DERIVE_SPN_WITHOUT_SERVER_REF
	// The DS cannot derive a service principal name (SPN) with which to mutually authenticate the target server because the corresponding server object in the local DS database has no serverReference attribute.
	ERROR_DS_CANT_DERIVE_SPN_WITHOUT_SERVER_REF = 8589

	// ERROR_DS_SINGLE_USER_MODE_FAILED
	// The Directory Service failed to enter single user mode.
	ERROR_DS_SINGLE_USER_MODE_FAILED = 8590

	// ERROR_DS_NTDSCRIPT_SYNTAX_ERROR
	// The Directory Service cannot parse the script because of a syntax error.
	ERROR_DS_NTDSCRIPT_SYNTAX_ERROR = 8591

	// ERROR_DS_NTDSCRIPT_PROCESS_ERROR
	// The Directory Service cannot process the script because of an error.
	ERROR_DS_NTDSCRIPT_PROCESS_ERROR = 8592

	// ERROR_DS_DIFFERENT_REPL_EPOCHS
	// The directory service cannot perform the requested operation because the servers involved are of different replication epochs (which is usually related to a domain rename that is in progress).
	ERROR_DS_DIFFERENT_REPL_EPOCHS = 8593

	// ERROR_DS_DRS_EXTENSIONS_CHANGED
	// The directory service binding must be renegotiated due to a change in the server extensions information.
	ERROR_DS_DRS_EXTENSIONS_CHANGED = 8594

	// ERROR_DS_REPLICA_SET_CHANGE_NOT_ALLOWED_ON_DISABLED_CR
	// Operation not allowed on a disabled cross ref.
	ERROR_DS_REPLICA_SET_CHANGE_NOT_ALLOWED_ON_DISABLED_CR = 8595

	// ERROR_DS_NO_MSDS_INTID
	// Schema update failed: No values for msDS-IntId are available.
	ERROR_DS_NO_MSDS_INTID = 8596

	// ERROR_DS_DUP_MSDS_INTID
	// Schema update failed: Duplicate msDS-INtId. Retry the operation.
	ERROR_DS_DUP_MSDS_INTID = 8597

	// ERROR_DS_EXISTS_IN_RDNATTID
	// Schema deletion failed: attribute is used in rDNAttID.
	ERROR_DS_EXISTS_IN_RDNATTID = 8598

	// ERROR_DS_AUTHORIZATION_FAILED
	// The directory service failed to authorize the request.
	ERROR_DS_AUTHORIZATION_FAILED = 8599

	// ERROR_DS_INVALID_SCRIPT
	// The Directory Service cannot process the script because it is invalid.
	ERROR_DS_INVALID_SCRIPT = 8600

	// ERROR_DS_REMOTE_CROSSREF_OP_FAILED
	// The remote create cross reference operation failed on the Domain Naming Master FSMO. The operation's error is in the extended data.
	ERROR_DS_REMOTE_CROSSREF_OP_FAILED = 8601

	// ERROR_DS_CROSS_REF_BUSY
	// A cross reference is in use locally with the same name.
	ERROR_DS_CROSS_REF_BUSY = 8602

	// ERROR_DS_CANT_DERIVE_SPN_FOR_DELETED_DOMAIN
	// The DS cannot derive a service principal name (SPN) with which to mutually authenticate the target server because the server's domain has been deleted from the forest.
	ERROR_DS_CANT_DERIVE_SPN_FOR_DELETED_DOMAIN = 8603

	// ERROR_DS_CANT_DEMOTE_WITH_WRITEABLE_NC
	// Writeable NCs prevent this DC from demoting.
	ERROR_DS_CANT_DEMOTE_WITH_WRITEABLE_NC = 8604

	// ERROR_DS_DUPLICATE_ID_FOUND
	// The requested object has a non-unique identifier and cannot be retrieved.
	ERROR_DS_DUPLICATE_ID_FOUND = 8605

	// ERROR_DS_INSUFFICIENT_ATTR_TO_CREATE_OBJECT
	// Insufficient attributes were given to create an object. This object may not exist because it may have been deleted and already garbage collected.
	ERROR_DS_INSUFFICIENT_ATTR_TO_CREATE_OBJECT = 8606

	// ERROR_DS_GROUP_CONVERSION_ERROR
	// The group cannot be converted due to attribute restrictions on the requested group type.
	ERROR_DS_GROUP_CONVERSION_ERROR = 8607

	// ERROR_DS_CANT_MOVE_APP_BASIC_GROUP
	// Cross-domain move of non-empty basic application groups is not allowed.
	ERROR_DS_CANT_MOVE_APP_BASIC_GROUP = 8608

	// ERROR_DS_CANT_MOVE_APP_QUERY_GROUP
	// Cross-domain move of non-empty query based application groups is not allowed.
	ERROR_DS_CANT_MOVE_APP_QUERY_GROUP = 8609

	// ERROR_DS_ROLE_NOT_VERIFIED
	// The FSMO role ownership could not be verified because its directory partition has not replicated successfully with at least one replication partner.
	ERROR_DS_ROLE_NOT_VERIFIED = 8610

	// ERROR_DS_WKO_CONTAINER_CANNOT_BE_SPECIA
	// The target container for a redirection of a well known object container cannot already be a special container.
	ERROR_DS_WKO_CONTAINER_CANNOT_BE_SPECIAL = 8611

	// ERROR_DS_DOMAIN_RENAME_IN_PROGRESS
	// The Directory Service cannot perform the requested operation because a domain rename operation is in progress.
	ERROR_DS_DOMAIN_RENAME_IN_PROGRESS = 8612

	// ERROR_DS_EXISTING_AD_CHILD_NC
	// The directory service detected a child partition below the requested partition name. The partition hierarchy must be created in a top down method.
	ERROR_DS_EXISTING_AD_CHILD_NC = 8613

	// ERROR_DS_REPL_LIFETIME_EXCEEDED
	// The directory service cannot replicate with this server because the time since the last replication with this server has exceeded the tombstone lifetime.
	ERROR_DS_REPL_LIFETIME_EXCEEDED = 8614

	// ERROR_DS_DISALLOWED_IN_SYSTEM_CONTAINER
	// The requested operation is not allowed on an object under the system container.
	ERROR_DS_DISALLOWED_IN_SYSTEM_CONTAINER = 8615

	// ERROR_DS_LDAP_SEND_QUEUE_FUL
	// The LDAP servers network send queue has filled up because the client is not processing the results of its requests fast enough. No more requests will be processed until the client catches up. If the client does not catch up then it will be disconnected.
	ERROR_DS_LDAP_SEND_QUEUE_FULL = 8616

	// ERROR_DS_DRA_OUT_SCHEDULE_WINDOW
	// The scheduled replication did not take place because the system was too busy to execute the request within the schedule window. The replication queue is overloaded. Consider reducing the number of partners or decreasing the scheduled replication frequency.
	ERROR_DS_DRA_OUT_SCHEDULE_WINDOW = 8617

	// ERROR_DS_POLICY_NOT_KNOWN
	// At this time, it cannot be determined if the branch replication policy is available on the hub domain controller. Please retry at a later time to account for replication latencies.
	ERROR_DS_POLICY_NOT_KNOWN = 8618

	// ERROR_NO_SITE_SETTINGS_OBJECT
	// The site settings object for the specified site does not exist.
	ERROR_NO_SITE_SETTINGS_OBJECT = 8619

	// ERROR_NO_SECRETS
	// The local account store does not contain secret material for the specified account.
	ERROR_NO_SECRETS = 8620

	// ERROR_NO_WRITABLE_DC_FOUND
	// Could not find a writable domain controller in the domain.
	ERROR_NO_WRITABLE_DC_FOUND = 8621

	// ERROR_DS_NO_SERVER_OBJECT
	// The server object for the domain controller does not exist.
	ERROR_DS_NO_SERVER_OBJECT = 8622

	// ERROR_DS_NO_NTDSA_OBJECT
	// The NTDS Settings object for the domain controller does not exist.
	ERROR_DS_NO_NTDSA_OBJECT = 8623

	// ERROR_DS_NON_ASQ_SEARCH
	// The requested search operation is not supported for ASQ searches.
	ERROR_DS_NON_ASQ_SEARCH = 8624

	// ERROR_DS_AUDIT_FAILURE
	// A required audit event could not be generated for the operation.
	ERROR_DS_AUDIT_FAILURE = 8625

	// ERROR_DS_INVALID_SEARCH_FLAG_SUBTREE
	// The search flags for the attribute are invalid. The subtree index bit is valid only on single valued attributes.
	ERROR_DS_INVALID_SEARCH_FLAG_SUBTREE = 8626

	// ERROR_DS_INVALID_SEARCH_FLAG_TUPLE
	// The search flags for the attribute are invalid. The tuple index bit is valid only on attributes of Unicode strings.
	ERROR_DS_INVALID_SEARCH_FLAG_TUPLE = 8627

	// ERROR_DS_HIERARCHY_TABLE_TOO_DEEP
	// The address books are nested too deeply. Failed to build the hierarchy table.
	ERROR_DS_HIERARCHY_TABLE_TOO_DEEP = 8628

	// ERROR_DS_DRA_CORRUPT_UTD_VECTOR
	// The specified up-to-date-ness vector is corrupt.
	ERROR_DS_DRA_CORRUPT_UTD_VECTOR = 8629

	// ERROR_DS_DRA_SECRETS_DENIED
	// The request to replicate secrets is denied.
	ERROR_DS_DRA_SECRETS_DENIED = 8630

	// ERROR_DS_RESERVED_MAPI_ID
	// Schema update failed: The MAPI identifier is reserved.
	ERROR_DS_RESERVED_MAPI_ID = 8631

	// ERROR_DS_MAPI_ID_NOT_AVAILABLE
	// Schema update failed: There are no MAPI identifiers available.
	ERROR_DS_MAPI_ID_NOT_AVAILABLE = 8632

	// ERROR_DS_DRA_MISSING_KRBTGT_SECRET
	// The replication operation failed because the required attributes of the local krbtgt object are missing.
	ERROR_DS_DRA_MISSING_KRBTGT_SECRET = 8633

	// ERROR_DS_DOMAIN_NAME_EXISTS_IN_FOREST
	// The domain name of the trusted domain already exists in the forest.
	ERROR_DS_DOMAIN_NAME_EXISTS_IN_FOREST = 8634

	// ERROR_DS_FLAT_NAME_EXISTS_IN_FOREST
	// The flat name of the trusted domain already exists in the forest.
	ERROR_DS_FLAT_NAME_EXISTS_IN_FOREST = 8635

	// ERROR_INVALID_USER_PRINCIPAL_NAME
	// The User Principal Name (UPN) is invalid.
	ERROR_INVALID_USER_PRINCIPAL_NAME = 8636

	// ERROR_DS_OID_MAPPED_GROUP_CANT_HAVE_MEMBERS
	// OID mapped groups cannot have members.
	ERROR_DS_OID_MAPPED_GROUP_CANT_HAVE_MEMBERS = 8637

	// ERROR_DS_OID_NOT_FOUND
	// The specified OID cannot be found.
	ERROR_DS_OID_NOT_FOUND = 8638

	// ERROR_DS_DRA_RECYCLED_TARGET
	// The replication operation failed because the target object referred by a link value is recycled.
	ERROR_DS_DRA_RECYCLED_TARGET = 8639

	// ERROR_DS_DISALLOWED_NC_REDIRECT
	// The redirect operation failed because the target object is in a NC different from the domain NC of the current domain controller.
	ERROR_DS_DISALLOWED_NC_REDIRECT = 8640

	// ERROR_DS_HIGH_ADLDS_FF
	// The functional level of the AD LDS configuration set cannot be lowered to the requested value.
	ERROR_DS_HIGH_ADLDS_FFL = 8641

	// ERROR_DS_HIGH_DSA_VERSION
	// The functional level of the domain (or forest) cannot be lowered to the requested value.
	ERROR_DS_HIGH_DSA_VERSION = 8642

	// ERROR_DS_LOW_ADLDS_FF
	// The functional level of the AD LDS configuration set cannot be raised to the requested value, because there exist one or more ADLDS instances that are at a lower incompatible functional level.
	ERROR_DS_LOW_ADLDS_FFL = 8643

	// ERROR_DOMAIN_SID_SAME_AS_LOCAL_WORKSTATION
	// The domain join cannot be completed because the SID of the domain you attempted to join was identical to the SID of this machine. This is a symptom of an improperly cloned operating system install.  You should run sysprep on this machine in order to generate a new machine SID. Please see http://go.microsoft.com/fwlink/?LinkId=168895 for more information.
	ERROR_DOMAIN_SID_SAME_AS_LOCAL_WORKSTATION = 8644

	// ERROR_DS_UNDELETE_SAM_VALIDATION_FAILED
	// The undelete operation failed because the Sam Account Name or Additional Sam Account Name of the object being undeleted conflicts with an existing live object.
	ERROR_DS_UNDELETE_SAM_VALIDATION_FAILED = 8645

	// ERROR_INCORRECT_ACCOUNT_TYPE
	// The system is not authoritative for the specified account and therefore cannot complete the operation. Please retry the operation using the provider associated with this account. If this is an online provider please use the provider's online site.
	ERROR_INCORRECT_ACCOUNT_TYPE = 8646

	// ERROR_DS_SPN_VALUE_NOT_UNIQUE_IN_FOREST
	// The operation failed because SPN value provided for addition/modification is not unique forest-wide.
	ERROR_DS_SPN_VALUE_NOT_UNIQUE_IN_FOREST = 8647

	// ERROR_DS_UPN_VALUE_NOT_UNIQUE_IN_FOREST
	// The operation failed because UPN value provided for addition/modification is not unique forest-wide.
	ERROR_DS_UPN_VALUE_NOT_UNIQUE_IN_FOREST = 8648

	// ERROR_DS_MISSING_FOREST_TRUST
	// The operation failed because the addition/modification referenced an inbound forest-wide trust that is not present.
	ERROR_DS_MISSING_FOREST_TRUST = 8649

	// ERROR_DS_VALUE_KEY_NOT_UNIQUE
	// The link value specified was not found, but a link value with that key was found.
	ERROR_DS_VALUE_KEY_NOT_UNIQUE = 8650

	///////////////////////////////////////////////////
	//                                                /
	//        End of Active Directory Error Codes     /
	//                                                /
	//                 =8000 to =8999                 /
	///////////////////////////////////////////////////

	///////////////////////////////////////////////////
	//                                               //
	//               DNS Error codes                 //
	//                                               //
	//                =9000 to=9999                  //
	///////////////////////////////////////////////////

	// =============================
	// Facility DNS Error Messages
	// =============================

	//  DNS response codes.

	DNS_ERROR_RESPONSE_CODES_BASE = 9000

	DNS_ERROR_RCODE_NO_ERROR = NO_ERROR

	DNS_ERROR_MASK = 0x00002328 //=9000 or DNS_ERROR_RESPONSE_CODES_BASE

	// DNS_ERROR_RCODE_FORMAT_ERROR         =0x00002329

	// DNS_ERROR_RCODE_FORMAT_ERROR
	// DNS server unable to interpret format.

	DNS_ERROR_RCODE_FORMAT_ERROR = 9001

	// DNS_ERROR_RCODE_SERVER_FAILURE       =0x0000232a

	// DNS_ERROR_RCODE_SERVER_FAILURE
	// DNS server failure.

	DNS_ERROR_RCODE_SERVER_FAILURE = 9002

	// DNS_ERROR_RCODE_NAME_ERROR           =0x0000232b

	// DNS_ERROR_RCODE_NAME_ERROR
	// DNS name does not exist.

	DNS_ERROR_RCODE_NAME_ERROR = 9003

	// DNS_ERROR_RCODE_NOT_IMPLEMENTED      =0x0000232c

	// DNS_ERROR_RCODE_NOT_IMPLEMENTED
	// DNS request not supported by name server.

	DNS_ERROR_RCODE_NOT_IMPLEMENTED = 9004

	// DNS_ERROR_RCODE_REFUSED              =0x0000232d

	// DNS_ERROR_RCODE_REFUSED
	// DNS operation refused.

	DNS_ERROR_RCODE_REFUSED = 9005

	// DNS_ERROR_RCODE_YXDOMAIN             =0x0000232e

	// DNS_ERROR_RCODE_YXDOMAIN
	// DNS name that ought not exist, does exist.

	DNS_ERROR_RCODE_YXDOMAIN = 9006

	// DNS_ERROR_RCODE_YXRRSET              =0x0000232f

	// DNS_ERROR_RCODE_YXRRSET
	// DNS RR set that ought not exist, does exist.

	DNS_ERROR_RCODE_YXRRSET = 9007

	// DNS_ERROR_RCODE_NXRRSET              =0x00002330

	// DNS_ERROR_RCODE_NXRRSET
	// DNS RR set that ought to exist, does not exist.

	DNS_ERROR_RCODE_NXRRSET = 9008

	// DNS_ERROR_RCODE_NOTAUTH              =0x00002331

	// DNS_ERROR_RCODE_NOTAUTH
	// DNS server not authoritative for zone.

	DNS_ERROR_RCODE_NOTAUTH = 9009

	// DNS_ERROR_RCODE_NOTZONE              =0x00002332

	// DNS_ERROR_RCODE_NOTZONE
	// DNS name in update or prereq is not in zone.

	DNS_ERROR_RCODE_NOTZONE = 9010

	// DNS_ERROR_RCODE_BADSIG               =0x00002338

	// DNS_ERROR_RCODE_BADSIG
	// DNS signature failed to verify.

	DNS_ERROR_RCODE_BADSIG = 9016

	// DNS_ERROR_RCODE_BADKEY               =0x00002339

	// DNS_ERROR_RCODE_BADKEY
	// DNS bad key.

	DNS_ERROR_RCODE_BADKEY = 9017

	// DNS_ERROR_RCODE_BADTIME              =0x0000233a

	// DNS_ERROR_RCODE_BADTIME
	// DNS signature validity expired.

	DNS_ERROR_RCODE_BADTIME = 9018

	DNS_ERROR_RCODE_LAST = DNS_ERROR_RCODE_BADTIME

	// DNSSEC errors

	DNS_ERROR_DNSSEC_BASE = 9100

	// DNS_ERROR_KEYMASTER_REQUIRED
	// Only the DNS server acting as the key master for the zone may perform this operation.

	DNS_ERROR_KEYMASTER_REQUIRED = 9101

	// DNS_ERROR_NOT_ALLOWED_ON_SIGNED_ZONE
	// This operation is not allowed on a zone that is signed or has signing keys.

	DNS_ERROR_NOT_ALLOWED_ON_SIGNED_ZONE = 9102

	// DNS_ERROR_NSEC3_INCOMPATIBLE_WITH_RSA_SHA1
	// NSEC3 is not compatible with the RSA-SHA-1 algorithm. Choose a different algorithm or use NSEC.

	DNS_ERROR_NSEC3_INCOMPATIBLE_WITH_RSA_SHA1 = 9103

	// DNS_ERROR_NOT_ENOUGH_SIGNING_KEY_DESCRIPTORS
	// The zone does not have enough signing keys. There must be at least one key signing key (KSK) and at least one zone signing key (ZSK).

	DNS_ERROR_NOT_ENOUGH_SIGNING_KEY_DESCRIPTORS = 9104

	// DNS_ERROR_UNSUPPORTED_ALGORITHM
	// The specified algorithm is not supported.

	DNS_ERROR_UNSUPPORTED_ALGORITHM = 9105

	// DNS_ERROR_INVALID_KEY_SIZE
	// The specified key size is not supported.

	DNS_ERROR_INVALID_KEY_SIZE = 9106

	// DNS_ERROR_SIGNING_KEY_NOT_ACCESSIBLE
	// One or more of the signing keys for a zone are not accessible to the DNS server. Zone signing will not be operational until this error is resolved.

	DNS_ERROR_SIGNING_KEY_NOT_ACCESSIBLE = 9107

	// DNS_ERROR_KSP_DOES_NOT_SUPPORT_PROTECTION
	// The specified key storage provider does not support DPAPI++ data protection. Zone signing will not be operational until this error is resolved.

	DNS_ERROR_KSP_DOES_NOT_SUPPORT_PROTECTION = 9108

	// DNS_ERROR_UNEXPECTED_DATA_PROTECTION_ERROR
	// An unexpected DPAPI++ error was encountered. Zone signing will not be operational until this error is resolved.

	DNS_ERROR_UNEXPECTED_DATA_PROTECTION_ERROR = 9109

	// DNS_ERROR_UNEXPECTED_CNG_ERROR
	// An unexpected crypto error was encountered. Zone signing may not be operational until this error is resolved.

	DNS_ERROR_UNEXPECTED_CNG_ERROR = 9110

	// DNS_ERROR_UNKNOWN_SIGNING_PARAMETER_VERSION
	// The DNS server encountered a signing key with an unknown version. Zone signing will not be operational until this error is resolved.

	DNS_ERROR_UNKNOWN_SIGNING_PARAMETER_VERSION = 9111

	// DNS_ERROR_KSP_NOT_ACCESSIBLE
	// The specified key service provider cannot be opened by the DNS server.

	DNS_ERROR_KSP_NOT_ACCESSIBLE = 9112

	// DNS_ERROR_TOO_MANY_SKDS
	// The DNS server cannot accept any more signing keys with the specified algorithm and KSK flag value for this zone.

	DNS_ERROR_TOO_MANY_SKDS = 9113

	// DNS_ERROR_INVALID_ROLLOVER_PERIOD
	// The specified rollover period is invalid.

	DNS_ERROR_INVALID_ROLLOVER_PERIOD = 9114

	// DNS_ERROR_INVALID_INITIAL_ROLLOVER_OFFSET
	// The specified initial rollover offset is invalid.

	DNS_ERROR_INVALID_INITIAL_ROLLOVER_OFFSET = 9115

	// DNS_ERROR_ROLLOVER_IN_PROGRESS
	// The specified signing key is already in process of rolling over keys.

	DNS_ERROR_ROLLOVER_IN_PROGRESS = 9116

	// DNS_ERROR_STANDBY_KEY_NOT_PRESENT
	// The specified signing key does not have a standby key to revoke.

	DNS_ERROR_STANDBY_KEY_NOT_PRESENT = 9117

	// DNS_ERROR_NOT_ALLOWED_ON_ZSK
	// This operation is not allowed on a zone signing key (ZSK).

	DNS_ERROR_NOT_ALLOWED_ON_ZSK = 9118

	// DNS_ERROR_NOT_ALLOWED_ON_ACTIVE_SKD
	// This operation is not allowed on an active signing key.

	DNS_ERROR_NOT_ALLOWED_ON_ACTIVE_SKD = 9119

	// DNS_ERROR_ROLLOVER_ALREADY_QUEUED
	// The specified signing key is already queued for rollover.

	DNS_ERROR_ROLLOVER_ALREADY_QUEUED = 9120

	// DNS_ERROR_NOT_ALLOWED_ON_UNSIGNED_ZONE
	// This operation is not allowed on an unsigned zone.

	DNS_ERROR_NOT_ALLOWED_ON_UNSIGNED_ZONE = 9121

	// DNS_ERROR_BAD_KEYMASTER
	// This operation could not be completed because the DNS server listed as the current key master for this zone is down or misconfigured. Resolve the problem on the current key master for this zone or use another DNS server to seize the key master role.

	DNS_ERROR_BAD_KEYMASTER = 9122

	// DNS_ERROR_INVALID_SIGNATURE_VALIDITY_PERIOD
	// The specified signature validity period is invalid.

	DNS_ERROR_INVALID_SIGNATURE_VALIDITY_PERIOD = 9123

	// DNS_ERROR_INVALID_NSEC3_ITERATION_COUNT
	// The specified NSEC3 iteration count is higher than allowed by the minimum key length used in the zone.

	DNS_ERROR_INVALID_NSEC3_ITERATION_COUNT = 9124

	// DNS_ERROR_DNSSEC_IS_DISABLED
	// This operation could not be completed because the DNS server has been configured with DNSSEC features disabled. Enable DNSSEC on the DNS server.

	DNS_ERROR_DNSSEC_IS_DISABLED = 9125

	// DNS_ERROR_INVALID_XM
	// This operation could not be completed because the XML stream received is empty or syntactically invalid.

	DNS_ERROR_INVALID_XML = 9126

	// DNS_ERROR_NO_VALID_TRUST_ANCHORS
	// This operation completed, but no trust anchors were added because all of the trust anchors received were either invalid, unsupported, expired, or would not become valid in less than=30 days.

	DNS_ERROR_NO_VALID_TRUST_ANCHORS = 9127

	// DNS_ERROR_ROLLOVER_NOT_POKEABLE
	// The specified signing key is not waiting for parental DS update.

	DNS_ERROR_ROLLOVER_NOT_POKEABLE = 9128

	// DNS_ERROR_NSEC3_NAME_COLLISION
	// Hash collision detected during NSEC3 signing. Specify a different user-provided salt, or use a randomly generated salt, and attempt to sign the zone again.

	DNS_ERROR_NSEC3_NAME_COLLISION = 9129

	// DNS_ERROR_NSEC_INCOMPATIBLE_WITH_NSEC3_RSA_SHA1
	// NSEC is not compatible with the NSEC3-RSA-SHA-1 algorithm. Choose a different algorithm or use NSEC3.

	DNS_ERROR_NSEC_INCOMPATIBLE_WITH_NSEC3_RSA_SHA1 = 9130

	// Packet format

	DNS_ERROR_PACKET_FMT_BASE = 9500

	// DNS_INFO_NO_RECORDS                  =0x0000251d

	// DNS_INFO_NO_RECORDS
	// No records found for given DNS query.

	DNS_INFO_NO_RECORDS = 9501

	// DNS_ERROR_BAD_PACKET                 =0x0000251e

	// DNS_ERROR_BAD_PACKET
	// Bad DNS packet.

	DNS_ERROR_BAD_PACKET = 9502

	// DNS_ERROR_NO_PACKET                  =0x0000251f

	// DNS_ERROR_NO_PACKET
	// No DNS packet.

	DNS_ERROR_NO_PACKET = 9503

	// DNS_ERROR_RCODE                      =0x00002520

	// DNS_ERROR_RCODE
	// DNS error, check rcode.

	DNS_ERROR_RCODE = 9504

	// DNS_ERROR_UNSECURE_PACKET            =0x00002521

	// DNS_ERROR_UNSECURE_PACKET
	// Unsecured DNS packet.

	DNS_ERROR_UNSECURE_PACKET = 9505

	DNS_STATUS_PACKET_UNSECURE = DNS_ERROR_UNSECURE_PACKET

	// DNS_REQUEST_PENDING                    =0x00002522

	// DNS_REQUEST_PENDING
	// DNS query request is pending.

	DNS_REQUEST_PENDING = 9506

	// General API errors

	DNS_ERROR_NO_MEMORY    = ERROR_OUTOFMEMORY
	DNS_ERROR_INVALID_NAME = ERROR_INVALID_NAME
	DNS_ERROR_INVALID_DATA = ERROR_INVALID_DATA

	DNS_ERROR_GENERAL_API_BASE = 9550

	// DNS_ERROR_INVALID_TYPE               =0x0000254f

	// DNS_ERROR_INVALID_TYPE
	// Invalid DNS type.

	DNS_ERROR_INVALID_TYPE = 9551

	// DNS_ERROR_INVALID_IP_ADDRESS         =0x00002550

	// DNS_ERROR_INVALID_IP_ADDRESS
	// Invalid IP address.

	DNS_ERROR_INVALID_IP_ADDRESS = 9552

	// DNS_ERROR_INVALID_PROPERTY           =0x00002551

	// DNS_ERROR_INVALID_PROPERTY
	// Invalid property.

	DNS_ERROR_INVALID_PROPERTY = 9553

	// DNS_ERROR_TRY_AGAIN_LATER            =0x00002552

	// DNS_ERROR_TRY_AGAIN_LATER
	// Try DNS operation again later.

	DNS_ERROR_TRY_AGAIN_LATER = 9554

	// DNS_ERROR_NOT_UNIQUE                 =0x00002553

	// DNS_ERROR_NOT_UNIQUE
	// Record for given name and type is not unique.

	DNS_ERROR_NOT_UNIQUE = 9555

	// DNS_ERROR_NON_RFC_NAME               =0x00002554

	// DNS_ERROR_NON_RFC_NAME
	// DNS name does not comply with RFC specifications.

	DNS_ERROR_NON_RFC_NAME = 9556

	// DNS_STATUS_FQDN                      =0x00002555

	// DNS_STATUS_FQDN
	// DNS name is a fully-qualified DNS name.

	DNS_STATUS_FQDN = 9557

	// DNS_STATUS_DOTTED_NAME               =0x00002556

	// DNS_STATUS_DOTTED_NAME
	// DNS name is dotted (multi-label).

	DNS_STATUS_DOTTED_NAME = 9558

	// DNS_STATUS_SINGLE_PART_NAME          =0x00002557

	// DNS_STATUS_SINGLE_PART_NAME
	// DNS name is a single-part name.

	DNS_STATUS_SINGLE_PART_NAME = 9559

	// DNS_ERROR_INVALID_NAME_CHAR          =0x00002558

	// DNS_ERROR_INVALID_NAME_CHAR
	// DNS name contains an invalid character.

	DNS_ERROR_INVALID_NAME_CHAR = 9560

	// DNS_ERROR_NUMERIC_NAME               =0x00002559

	// DNS_ERROR_NUMERIC_NAME
	// DNS name is entirely numeric.

	DNS_ERROR_NUMERIC_NAME = 9561

	// DNS_ERROR_NOT_ALLOWED_ON_ROOT_SERVER =0x0000255A

	// DNS_ERROR_NOT_ALLOWED_ON_ROOT_SERVER
	// The operation requested is not permitted on a DNS root server.

	DNS_ERROR_NOT_ALLOWED_ON_ROOT_SERVER = 9562

	// DNS_ERROR_NOT_ALLOWED_UNDER_DELEGATION =0x0000255B

	// DNS_ERROR_NOT_ALLOWED_UNDER_DELEGATION
	// The record could not be created because this part of the DNS namespace has been delegated to another server.

	DNS_ERROR_NOT_ALLOWED_UNDER_DELEGATION = 9563

	// DNS_ERROR_CANNOT_FIND_ROOT_HINTS =0x0000255C

	// DNS_ERROR_CANNOT_FIND_ROOT_HINTS
	// The DNS server could not find a set of root hints.

	DNS_ERROR_CANNOT_FIND_ROOT_HINTS = 9564

	// DNS_ERROR_INCONSISTENT_ROOT_HINTS =0x0000255D

	// DNS_ERROR_INCONSISTENT_ROOT_HINTS
	// The DNS server found root hints but they were not consistent across all adapters.

	DNS_ERROR_INCONSISTENT_ROOT_HINTS = 9565

	// DNS_ERROR_DWORD_VALUE_TOO_SMALL   =0x0000255E

	// DNS_ERROR_DWORD_VALUE_TOO_SMAL
	// The specified value is too small for this parameter.

	DNS_ERROR_DWORD_VALUE_TOO_SMALL = 9566

	// DNS_ERROR_DWORD_VALUE_TOO_LARGE   =0x0000255F

	// DNS_ERROR_DWORD_VALUE_TOO_LARGE
	// The specified value is too large for this parameter.

	DNS_ERROR_DWORD_VALUE_TOO_LARGE = 9567

	// DNS_ERROR_BACKGROUND_LOADING      =0x00002560

	// DNS_ERROR_BACKGROUND_LOADING
	// This operation is not allowed while the DNS server is loading zones in the background. Please try again later.

	DNS_ERROR_BACKGROUND_LOADING = 9568

	// DNS_ERROR_NOT_ALLOWED_ON_RODC     =0x00002561

	// DNS_ERROR_NOT_ALLOWED_ON_RODC
	// The operation requested is not permitted on against a DNS server running on a read-only DC.

	DNS_ERROR_NOT_ALLOWED_ON_RODC = 9569

	// DNS_ERROR_NOT_ALLOWED_UNDER_DNAME  =0x00002562

	// DNS_ERROR_NOT_ALLOWED_UNDER_DNAME
	// No data is allowed to exist underneath a DNAME record.

	DNS_ERROR_NOT_ALLOWED_UNDER_DNAME = 9570

	// DNS_ERROR_DELEGATION_REQUIRED      =0x00002563

	// DNS_ERROR_DELEGATION_REQUIRED
	// This operation requires credentials delegation.

	DNS_ERROR_DELEGATION_REQUIRED = 9571

	// DNS_ERROR_INVALID_POLICY_TABLE       =0x00002564

	// DNS_ERROR_INVALID_POLICY_TABLE
	// Name resolution policy table has been corrupted. DNS resolution will fail until it is fixed. Contact your network administrator.

	DNS_ERROR_INVALID_POLICY_TABLE = 9572

	// DNS_ERROR_ADDRESS_REQUIRED       =0x00002565

	// DNS_ERROR_ADDRESS_REQUIRED
	// Not allowed to remove all addresses.

	DNS_ERROR_ADDRESS_REQUIRED = 9573

	// Zone errors

	DNS_ERROR_ZONE_BASE = 9600

	// DNS_ERROR_ZONE_DOES_NOT_EXIST        =0x00002581

	// DNS_ERROR_ZONE_DOES_NOT_EXIST
	// DNS zone does not exist.

	DNS_ERROR_ZONE_DOES_NOT_EXIST = 9601

	// DNS_ERROR_NO_ZONE_INFO               =0x00002582

	// DNS_ERROR_NO_ZONE_INFO
	// DNS zone information not available.

	DNS_ERROR_NO_ZONE_INFO = 9602

	// DNS_ERROR_INVALID_ZONE_OPERATION     =0x00002583

	// DNS_ERROR_INVALID_ZONE_OPERATION
	// Invalid operation for DNS zone.

	DNS_ERROR_INVALID_ZONE_OPERATION = 9603

	// DNS_ERROR_ZONE_CONFIGURATION_ERROR   =0x00002584

	// DNS_ERROR_ZONE_CONFIGURATION_ERROR
	// Invalid DNS zone configuration.

	DNS_ERROR_ZONE_CONFIGURATION_ERROR = 9604

	// DNS_ERROR_ZONE_HAS_NO_SOA_RECORD     =0x00002585

	// DNS_ERROR_ZONE_HAS_NO_SOA_RECORD
	// DNS zone has no start of authority (SOA) record.

	DNS_ERROR_ZONE_HAS_NO_SOA_RECORD = 9605

	// DNS_ERROR_ZONE_HAS_NO_NS_RECORDS     =0x00002586

	// DNS_ERROR_ZONE_HAS_NO_NS_RECORDS
	// DNS zone has no Name Server (NS) record.

	DNS_ERROR_ZONE_HAS_NO_NS_RECORDS = 9606

	// DNS_ERROR_ZONE_LOCKED                =0x00002587

	// DNS_ERROR_ZONE_LOCKED
	// DNS zone is locked.

	DNS_ERROR_ZONE_LOCKED = 9607

	// DNS_ERROR_ZONE_CREATION_FAILED       =0x00002588

	// DNS_ERROR_ZONE_CREATION_FAILED
	// DNS zone creation failed.

	DNS_ERROR_ZONE_CREATION_FAILED = 9608

	// DNS_ERROR_ZONE_ALREADY_EXISTS        =0x00002589

	// DNS_ERROR_ZONE_ALREADY_EXISTS
	// DNS zone already exists.

	DNS_ERROR_ZONE_ALREADY_EXISTS = 9609

	// DNS_ERROR_AUTOZONE_ALREADY_EXISTS    =0x0000258a

	// DNS_ERROR_AUTOZONE_ALREADY_EXISTS
	// DNS automatic zone already exists.

	DNS_ERROR_AUTOZONE_ALREADY_EXISTS = 9610

	// DNS_ERROR_INVALID_ZONE_TYPE          =0x0000258b

	// DNS_ERROR_INVALID_ZONE_TYPE
	// Invalid DNS zone type.

	DNS_ERROR_INVALID_ZONE_TYPE = 9611

	// DNS_ERROR_SECONDARY_REQUIRES_MASTER_IP=0x0000258c

	// DNS_ERROR_SECONDARY_REQUIRES_MASTER_IP
	// Secondary DNS zone requires master IP address.

	DNS_ERROR_SECONDARY_REQUIRES_MASTER_IP = 9612

	// DNS_ERROR_ZONE_NOT_SECONDARY         =0x0000258d

	// DNS_ERROR_ZONE_NOT_SECONDARY
	// DNS zone not secondary.

	DNS_ERROR_ZONE_NOT_SECONDARY = 9613

	// DNS_ERROR_NEED_SECONDARY_ADDRESSES   =0x0000258e

	// DNS_ERROR_NEED_SECONDARY_ADDRESSES
	// Need secondary IP address.

	DNS_ERROR_NEED_SECONDARY_ADDRESSES = 9614

	// DNS_ERROR_WINS_INIT_FAILED           =0x0000258f

	// DNS_ERROR_WINS_INIT_FAILED
	// WINS initialization failed.

	DNS_ERROR_WINS_INIT_FAILED = 9615

	// DNS_ERROR_NEED_WINS_SERVERS          =0x00002590

	// DNS_ERROR_NEED_WINS_SERVERS
	// Need WINS servers.

	DNS_ERROR_NEED_WINS_SERVERS = 9616

	// DNS_ERROR_NBSTAT_INIT_FAILED         =0x00002591

	// DNS_ERROR_NBSTAT_INIT_FAILED
	// NBTSTAT initialization call failed.

	DNS_ERROR_NBSTAT_INIT_FAILED = 9617

	// DNS_ERROR_SOA_DELETE_INVALID         =0x00002592

	// DNS_ERROR_SOA_DELETE_INVALID
	// Invalid delete of start of authority (SOA)

	DNS_ERROR_SOA_DELETE_INVALID = 9618

	// DNS_ERROR_FORWARDER_ALREADY_EXISTS   =0x00002593

	// DNS_ERROR_FORWARDER_ALREADY_EXISTS
	// A conditional forwarding zone already exists for that name.

	DNS_ERROR_FORWARDER_ALREADY_EXISTS = 9619

	// DNS_ERROR_ZONE_REQUIRES_MASTER_IP    =0x00002594

	// DNS_ERROR_ZONE_REQUIRES_MASTER_IP
	// This zone must be configured with one or more master DNS server IP addresses.

	DNS_ERROR_ZONE_REQUIRES_MASTER_IP = 9620

	// DNS_ERROR_ZONE_IS_SHUTDOWN           =0x00002595

	// DNS_ERROR_ZONE_IS_SHUTDOWN
	// The operation cannot be performed because this zone is shut down.

	DNS_ERROR_ZONE_IS_SHUTDOWN = 9621

	// DNS_ERROR_ZONE_LOCKED_FOR_SIGNING    =0x00002596

	// DNS_ERROR_ZONE_LOCKED_FOR_SIGNING
	// This operation cannot be performed because the zone is currently being signed. Please try again later.

	DNS_ERROR_ZONE_LOCKED_FOR_SIGNING = 9622

	// Datafile errors

	DNS_ERROR_DATAFILE_BASE = 9650

	// DNS                                  =0x000025b3

	// DNS_ERROR_PRIMARY_REQUIRES_DATAFILE
	// Primary DNS zone requires datafile.

	DNS_ERROR_PRIMARY_REQUIRES_DATAFILE = 9651

	// DNS                                  =0x000025b4

	// DNS_ERROR_INVALID_DATAFILE_NAME
	// Invalid datafile name for DNS zone.

	DNS_ERROR_INVALID_DATAFILE_NAME = 9652

	// DNS                                  =0x000025b5

	// DNS_ERROR_DATAFILE_OPEN_FAILURE
	// Failed to open datafile for DNS zone.

	DNS_ERROR_DATAFILE_OPEN_FAILURE = 9653

	// DNS                                  =0x000025b6

	// DNS_ERROR_FILE_WRITEBACK_FAILED
	// Failed to write datafile for DNS zone.

	DNS_ERROR_FILE_WRITEBACK_FAILED = 9654

	// DNS                                  =0x000025b7

	// DNS_ERROR_DATAFILE_PARSING
	// Failure while reading datafile for DNS zone.

	DNS_ERROR_DATAFILE_PARSING = 9655

	// Database errors

	DNS_ERROR_DATABASE_BASE = 9700

	// DNS_ERROR_RECORD_DOES_NOT_EXIST      =0x000025e5

	// DNS_ERROR_RECORD_DOES_NOT_EXIST
	// DNS record does not exist.

	DNS_ERROR_RECORD_DOES_NOT_EXIST = 9701

	// DNS_ERROR_RECORD_FORMAT              =0x000025e6

	// DNS_ERROR_RECORD_FORMAT
	// DNS record format error.

	DNS_ERROR_RECORD_FORMAT = 9702

	// DNS_ERROR_NODE_CREATION_FAILED       =0x000025e7

	// DNS_ERROR_NODE_CREATION_FAILED
	// Node creation failure in DNS.

	DNS_ERROR_NODE_CREATION_FAILED = 9703

	// DNS_ERROR_UNKNOWN_RECORD_TYPE        =0x000025e8

	// DNS_ERROR_UNKNOWN_RECORD_TYPE
	// Unknown DNS record type.

	DNS_ERROR_UNKNOWN_RECORD_TYPE = 9704

	// DNS_ERROR_RECORD_TIMED_OUT           =0x000025e9

	// DNS_ERROR_RECORD_TIMED_OUT
	// DNS record timed out.

	DNS_ERROR_RECORD_TIMED_OUT = 9705

	// DNS_ERROR_NAME_NOT_IN_ZONE           =0x000025ea

	// DNS_ERROR_NAME_NOT_IN_ZONE
	// Name not in DNS zone.

	DNS_ERROR_NAME_NOT_IN_ZONE = 9706

	// DNS_ERROR_CNAME_LOOP                 =0x000025eb

	// DNS_ERROR_CNAME_LOOP
	// CNAME loop detected.

	DNS_ERROR_CNAME_LOOP = 9707

	// DNS_ERROR_NODE_IS_CNAME              =0x000025ec

	// DNS_ERROR_NODE_IS_CNAME
	// Node is a CNAME DNS record.

	DNS_ERROR_NODE_IS_CNAME = 9708

	// DNS_ERROR_CNAME_COLLISION            =0x000025ed

	// DNS_ERROR_CNAME_COLLISION
	// A CNAME record already exists for given name.

	DNS_ERROR_CNAME_COLLISION = 9709

	// DNS_ERROR_RECORD_ONLY_AT_ZONE_ROOT   =0x000025ee

	// DNS_ERROR_RECORD_ONLY_AT_ZONE_ROOT
	// Record only at DNS zone root.

	DNS_ERROR_RECORD_ONLY_AT_ZONE_ROOT = 9710

	// DNS_ERROR_RECORD_ALREADY_EXISTS      =0x000025ef

	// DNS_ERROR_RECORD_ALREADY_EXISTS
	// DNS record already exists.

	DNS_ERROR_RECORD_ALREADY_EXISTS = 9711

	// DNS_ERROR_SECONDARY_DATA             =0x000025f0

	// DNS_ERROR_SECONDARY_DATA
	// Secondary DNS zone data error.

	DNS_ERROR_SECONDARY_DATA = 9712

	// DNS_ERROR_NO_CREATE_CACHE_DATA       =0x000025f1

	// DNS_ERROR_NO_CREATE_CACHE_DATA
	// Could not create DNS cache data.

	DNS_ERROR_NO_CREATE_CACHE_DATA = 9713

	// DNS_ERROR_NAME_DOES_NOT_EXIST        =0x000025f2

	// DNS_ERROR_NAME_DOES_NOT_EXIST
	// DNS name does not exist.

	DNS_ERROR_NAME_DOES_NOT_EXIST = 9714

	// DNS_WARNING_PTR_CREATE_FAILED        =0x000025f3

	// DNS_WARNING_PTR_CREATE_FAILED
	// Could not create pointer (PTR) record.

	DNS_WARNING_PTR_CREATE_FAILED = 9715

	// DNS_WARNING_DOMAIN_UNDELETED         =0x000025f4

	// DNS_WARNING_DOMAIN_UNDELETED
	// DNS domain was undeleted.

	DNS_WARNING_DOMAIN_UNDELETED = 9716

	// DNS_ERROR_DS_UNAVAILABLE             =0x000025f5

	// DNS_ERROR_DS_UNAVAILABLE
	// The directory service is unavailable.

	DNS_ERROR_DS_UNAVAILABLE = 9717

	// DNS_ERROR_DS_ZONE_ALREADY_EXISTS     =0x000025f6

	// DNS_ERROR_DS_ZONE_ALREADY_EXISTS
	// DNS zone already exists in the directory service.

	DNS_ERROR_DS_ZONE_ALREADY_EXISTS = 9718

	// DNS_ERROR_NO_BOOTFILE_IF_DS_ZONE     =0x000025f7

	// DNS_ERROR_NO_BOOTFILE_IF_DS_ZONE
	// DNS server not creating or reading the boot file for the directory service integrated DNS zone.

	DNS_ERROR_NO_BOOTFILE_IF_DS_ZONE = 9719

	// DNS_ERROR_NODE_IS_DNAME              =0x000025f8

	// DNS_ERROR_NODE_IS_DNAME
	// Node is a DNAME DNS record.

	DNS_ERROR_NODE_IS_DNAME = 9720

	// DNS_ERROR_DNAME_COLLISION            =0x000025f9

	// DNS_ERROR_DNAME_COLLISION
	// A DNAME record already exists for given name.

	DNS_ERROR_DNAME_COLLISION = 9721

	// DNS_ERROR_ALIAS_LOOP                 =0x000025fa

	// DNS_ERROR_ALIAS_LOOP
	// An alias loop has been detected with either CNAME or DNAME records.

	DNS_ERROR_ALIAS_LOOP = 9722

	// Operation errors

	DNS_ERROR_OPERATION_BASE = 9750

	// DNS_INFO_AXFR_COMPLETE               =0x00002617

	// DNS_INFO_AXFR_COMPLETE
	// DNS AXFR (zone transfer) complete.

	DNS_INFO_AXFR_COMPLETE = 9751

	// DNS_ERROR_AXFR                       =0x00002618

	// DNS_ERROR_AXFR
	// DNS zone transfer failed.

	DNS_ERROR_AXFR = 9752

	// DNS_INFO_ADDED_LOCAL_WINS            =0x00002619

	// DNS_INFO_ADDED_LOCAL_WINS
	// Added local WINS server.

	DNS_INFO_ADDED_LOCAL_WINS = 9753

	// Secure update

	DNS_ERROR_SECURE_BASE = 9800

	// DNS_STATUS_CONTINUE_NEEDED           =0x00002649

	// DNS_STATUS_CONTINUE_NEEDED
	// Secure update call needs to continue update request.

	DNS_STATUS_CONTINUE_NEEDED = 9801

	// Setup errors

	DNS_ERROR_SETUP_BASE = 9850

	// DNS_ERROR_NO_TCPIP                   =0x0000267b

	// DNS_ERROR_NO_TCPIP
	// TCP/IP network protocol not installed.

	DNS_ERROR_NO_TCPIP = 9851

	// DNS_ERROR_NO_DNS_SERVERS             =0x0000267c

	// DNS_ERROR_NO_DNS_SERVERS
	// No DNS servers configured for local system.

	DNS_ERROR_NO_DNS_SERVERS = 9852

	// Directory partition (DP) errors

	DNS_ERROR_DP_BASE = 9900

	// DNS_ERROR_DP_DOES_NOT_EXIST          =0x000026ad

	// DNS_ERROR_DP_DOES_NOT_EXIST
	// The specified directory partition does not exist.

	DNS_ERROR_DP_DOES_NOT_EXIST = 9901

	// DNS_ERROR_DP_ALREADY_EXISTS          =0x000026ae

	// DNS_ERROR_DP_ALREADY_EXISTS
	// The specified directory partition already exists.

	DNS_ERROR_DP_ALREADY_EXISTS = 9902

	// DNS_ERROR_DP_NOT_ENLISTED            =0x000026af

	// DNS_ERROR_DP_NOT_ENLISTED
	// This DNS server is not enlisted in the specified directory partition.

	DNS_ERROR_DP_NOT_ENLISTED = 9903

	// DNS_ERROR_DP_ALREADY_ENLISTED        =0x000026b0

	// DNS_ERROR_DP_ALREADY_ENLISTED
	// This DNS server is already enlisted in the specified directory partition.

	DNS_ERROR_DP_ALREADY_ENLISTED = 9904

	// DNS_ERROR_DP_NOT_AVAILABLE           =0x000026b1

	// DNS_ERROR_DP_NOT_AVAILABLE
	// The directory partition is not available at this time. Please wait a few minutes and try again.

	DNS_ERROR_DP_NOT_AVAILABLE = 9905

	// DNS_ERROR_DP_FSMO_ERROR              =0x000026b2

	// DNS_ERROR_DP_FSMO_ERROR
	// The operation failed because the domain naming master FSMO role could not be reached. The domain controller holding the domain naming master FSMO role is down or unable to service the request or is not running Windows Server=2003 or later.

	DNS_ERROR_DP_FSMO_ERROR = 9906

	// DNS RRL errors from=9911 to=9920

	// DNS_ERROR_RRL_NOT_ENABLED=0x000026B7

	// DNS_ERROR_RRL_NOT_ENABLED
	// The RRL is not enabled.

	DNS_ERROR_RRL_NOT_ENABLED = 9911

	// DNS_ERROR_RRL_INVALID_WINDOW_SIZE=0x000026B8

	// DNS_ERROR_RRL_INVALID_WINDOW_SIZE
	// The window size parameter is invalid. It should be greater than or equal to=1.

	DNS_ERROR_RRL_INVALID_WINDOW_SIZE = 9912

	// DNS_ERROR_RRL_INVALID_IPV4_PREFIX=0x000026B9

	// DNS_ERROR_RRL_INVALID_IPV4_PREFIX
	// The IPv4 prefix length parameter is invalid. It should be less than or equal to=32.

	DNS_ERROR_RRL_INVALID_IPV4_PREFIX = 9913

	// DNS_ERROR_RRL_INVALID_IPV6_PREFIX=0x000026BA

	// DNS_ERROR_RRL_INVALID_IPV6_PREFIX
	// The IPv6 prefix length parameter is invalid. It should be less than or equal to=128.

	DNS_ERROR_RRL_INVALID_IPV6_PREFIX = 9914

	// DNS_ERROR_RRL_INVALID_TC_RATE=0x000026BB

	// DNS_ERROR_RRL_INVALID_TC_RATE
	// The TC Rate parameter is invalid. It should be less than=10.

	DNS_ERROR_RRL_INVALID_TC_RATE = 9915

	// DNS_ERROR_RRL_INVALID_LEAK_RATE=0x000026BC

	// DNS_ERROR_RRL_INVALID_LEAK_RATE
	// The Leak Rate parameter is invalid. It should be either=0, or between=2 and=10.

	DNS_ERROR_RRL_INVALID_LEAK_RATE = 9916

	// DNS_ERROR_RRL_LEAK_RATE_LESSTHAN_TC_RATE=0x000026BD

	// DNS_ERROR_RRL_LEAK_RATE_LESSTHAN_TC_RATE
	// The Leak Rate or TC Rate parameter is invalid. Leak Rate should be greater than TC Rate.

	DNS_ERROR_RRL_LEAK_RATE_LESSTHAN_TC_RATE = 9917

	// DNS Virtualization errors from=9921 to=9950

	// DNS_ERROR_VIRTUALIZATION_INSTANCE_ALREADY_EXISTS   =0x000026c1

	// DNS_ERROR_VIRTUALIZATION_INSTANCE_ALREADY_EXISTS
	// The virtualization instance already exists.

	DNS_ERROR_VIRTUALIZATION_INSTANCE_ALREADY_EXISTS = 9921

	// DNS_ERROR_VIRTUALIZATION_INSTANCE_DOES_NOT_EXIST   =0x000026c2

	// DNS_ERROR_VIRTUALIZATION_INSTANCE_DOES_NOT_EXIST
	// The virtualization instance does not exist.

	DNS_ERROR_VIRTUALIZATION_INSTANCE_DOES_NOT_EXIST = 9922

	// DNS_ERROR_VIRTUALIZATION_TREE_LOCKED       =0x000026c3

	// DNS_ERROR_VIRTUALIZATION_TREE_LOCKED
	// The virtualization tree is locked.

	DNS_ERROR_VIRTUALIZATION_TREE_LOCKED = 9923

	// DNS_ERROR_INVAILD_VIRTUALIZATION_INSTANCE_NAME     =0x000026c4

	// DNS_ERROR_INVAILD_VIRTUALIZATION_INSTANCE_NAME
	// Invalid virtualization instance name.

	DNS_ERROR_INVAILD_VIRTUALIZATION_INSTANCE_NAME = 9924

	// DNS_ERROR_DEFAULT_VIRTUALIZATION_INSTANCE  =0x000026c5

	// DNS_ERROR_DEFAULT_VIRTUALIZATION_INSTANCE
	// The default virtualization instance cannot be added, removed or modified.

	DNS_ERROR_DEFAULT_VIRTUALIZATION_INSTANCE = 9925

	// DNS ZoneScope errors from=9951 to=9970

	// DNS_ERROR_ZONESCOPE_ALREADY_EXISTS              =0x000026df

	// DNS_ERROR_ZONESCOPE_ALREADY_EXISTS
	// The scope already exists for the zone.

	DNS_ERROR_ZONESCOPE_ALREADY_EXISTS = 9951

	// DNS_ERROR_ZONESCOPE_DOES_NOT_EXIST      =0x000026e0

	// DNS_ERROR_ZONESCOPE_DOES_NOT_EXIST
	// The scope does not exist for the zone.

	DNS_ERROR_ZONESCOPE_DOES_NOT_EXIST = 9952

	// DNS_ERROR_DEFAULT_ZONESCOPE=0x000026e1

	// DNS_ERROR_DEFAULT_ZONESCOPE
	// The scope is the same as the default zone scope.

	DNS_ERROR_DEFAULT_ZONESCOPE = 9953

	// DNS_ERROR_INVALID_ZONESCOPE_NAME=0x000026e2

	// DNS_ERROR_INVALID_ZONESCOPE_NAME
	// The scope name contains invalid characters.

	DNS_ERROR_INVALID_ZONESCOPE_NAME = 9954

	// DNS_ERROR_NOT_ALLOWED_WITH_ZONESCOPES=0x000026e3

	// DNS_ERROR_NOT_ALLOWED_WITH_ZONESCOPES
	// Operation not allowed when the zone has scopes.

	DNS_ERROR_NOT_ALLOWED_WITH_ZONESCOPES = 9955

	// DNS_ERROR_LOAD_ZONESCOPE_FAILED=0x000026e4

	// DNS_ERROR_LOAD_ZONESCOPE_FAILED
	// Failed to load zone scope.

	DNS_ERROR_LOAD_ZONESCOPE_FAILED = 9956

	// DNS_ERROR_ZONESCOPE_FILE_WRITEBACK_FAILED=0x000026e5

	// DNS_ERROR_ZONESCOPE_FILE_WRITEBACK_FAILED
	// Failed to write data file for DNS zone scope. Please verify the file exists and is writable.

	DNS_ERROR_ZONESCOPE_FILE_WRITEBACK_FAILED = 9957

	// DNS_ERROR_INVALID_SCOPE_NAME=0x000026e6

	// DNS_ERROR_INVALID_SCOPE_NAME
	// The scope name contains invalid characters.

	DNS_ERROR_INVALID_SCOPE_NAME = 9958

	// DNS_ERROR_SCOPE_DOES_NOT_EXIST      =0x000026e7

	// DNS_ERROR_SCOPE_DOES_NOT_EXIST
	// The scope does not exist.

	DNS_ERROR_SCOPE_DOES_NOT_EXIST = 9959

	// DNS_ERROR_DEFAULT_SCOPE=0x000026e8

	// DNS_ERROR_DEFAULT_SCOPE
	// The scope is the same as the default scope.

	DNS_ERROR_DEFAULT_SCOPE = 9960

	// DNS_ERROR_INVALID_SCOPE_OPERATION=0x000026e9

	// DNS_ERROR_INVALID_SCOPE_OPERATION
	// The operation is invalid on the scope.

	DNS_ERROR_INVALID_SCOPE_OPERATION = 9961

	// DNS_ERROR_SCOPE_LOCKED=0x000026ea

	// DNS_ERROR_SCOPE_LOCKED
	// The scope is locked.

	DNS_ERROR_SCOPE_LOCKED = 9962

	// DNS_ERROR_SCOPE_ALREADY_EXISTS=0x000026eb

	// DNS_ERROR_SCOPE_ALREADY_EXISTS
	// The scope already exists.

	DNS_ERROR_SCOPE_ALREADY_EXISTS = 9963

	// DNS Policy errors from=9971 to=9999

	// DNS_ERROR_POLICY_ALREADY_EXISTS=0x000026f3

	// DNS_ERROR_POLICY_ALREADY_EXISTS
	// A policy with the same name already exists on this level (server level or zone level) on the DNS server.

	DNS_ERROR_POLICY_ALREADY_EXISTS = 9971

	// DNS_ERROR_POLICY_DOES_NOT_EXIST=0x000026f4

	// DNS_ERROR_POLICY_DOES_NOT_EXIST
	// No policy with this name exists on this level (server level or zone level) on the DNS server.

	DNS_ERROR_POLICY_DOES_NOT_EXIST = 9972

	// DNS_ERROR_POLICY_INVALID_CRITERIA=0x000026f5

	// DNS_ERROR_POLICY_INVALID_CRITERIA
	// The criteria provided in the policy are invalid.

	DNS_ERROR_POLICY_INVALID_CRITERIA = 9973

	// DNS_ERROR_POLICY_INVALID_SETTINGS=0x000026f6

	// DNS_ERROR_POLICY_INVALID_SETTINGS
	// At least one of the settings of this policy is invalid.

	DNS_ERROR_POLICY_INVALID_SETTINGS = 9974

	// DNS_ERROR_CLIENT_SUBNET_IS_ACCESSED=0x000026f7

	// DNS_ERROR_CLIENT_SUBNET_IS_ACCESSED
	// The client subnet cannot be deleted while it is being accessed by a policy.

	DNS_ERROR_CLIENT_SUBNET_IS_ACCESSED = 9975

	// DNS_ERROR_CLIENT_SUBNET_DOES_NOT_EXIST=0x000026f8

	// DNS_ERROR_CLIENT_SUBNET_DOES_NOT_EXIST
	// The client subnet does not exist on the DNS server.

	DNS_ERROR_CLIENT_SUBNET_DOES_NOT_EXIST = 9976

	// DNS_ERROR_CLIENT_SUBNET_ALREADY_EXISTS=0x000026f9

	// DNS_ERROR_CLIENT_SUBNET_ALREADY_EXISTS
	// A client subnet with this name already exists on the DNS server.

	DNS_ERROR_CLIENT_SUBNET_ALREADY_EXISTS = 9977

	// DNS_ERROR_SUBNET_DOES_NOT_EXIST=0x000026fa

	// DNS_ERROR_SUBNET_DOES_NOT_EXIST
	// The IP subnet specified does not exist in the client subnet.

	DNS_ERROR_SUBNET_DOES_NOT_EXIST = 9978

	// DNS_ERROR_SUBNET_ALREADY_EXISTS=0x000026fb

	// DNS_ERROR_SUBNET_ALREADY_EXISTS
	// The IP subnet that is being added, already exists in the client subnet.

	DNS_ERROR_SUBNET_ALREADY_EXISTS = 9979

	// DNS_ERROR_POLICY_LOCKED=0x000026fc

	// DNS_ERROR_POLICY_LOCKED
	// The policy is locked.

	DNS_ERROR_POLICY_LOCKED = 9980

	// DNS_ERROR_POLICY_INVALID_WEIGHT=0x000026fd

	// DNS_ERROR_POLICY_INVALID_WEIGHT
	// The weight of the scope in the policy is invalid.

	DNS_ERROR_POLICY_INVALID_WEIGHT = 9981

	// DNS_ERROR_POLICY_INVALID_NAME=0x000026fe

	// DNS_ERROR_POLICY_INVALID_NAME
	// The DNS policy name is invalid.

	DNS_ERROR_POLICY_INVALID_NAME = 9982

	// DNS_ERROR_POLICY_MISSING_CRITERIA=0x000026ff

	// DNS_ERROR_POLICY_MISSING_CRITERIA
	// The policy is missing criteria.

	DNS_ERROR_POLICY_MISSING_CRITERIA = 9983

	// DNS_ERROR_INVALID_CLIENT_SUBNET_NAME=0x00002700

	// DNS_ERROR_INVALID_CLIENT_SUBNET_NAME
	// The name of the the client subnet record is invalid.

	DNS_ERROR_INVALID_CLIENT_SUBNET_NAME = 9984

	// DNS_ERROR_POLICY_PROCESSING_ORDER_INVALID=0x00002701

	// DNS_ERROR_POLICY_PROCESSING_ORDER_INVALID
	// Invalid policy processing order.

	DNS_ERROR_POLICY_PROCESSING_ORDER_INVALID = 9985

	// DNS_ERROR_POLICY_SCOPE_MISSING=0x00002702

	// DNS_ERROR_POLICY_SCOPE_MISSING
	// The scope information has not been provided for a policy that requires it.

	DNS_ERROR_POLICY_SCOPE_MISSING = 9986

	// DNS_ERROR_POLICY_SCOPE_NOT_ALLOWED=0x00002703

	// DNS_ERROR_POLICY_SCOPE_NOT_ALLOWED
	// The scope information has been provided for a policy that does not require it.

	DNS_ERROR_POLICY_SCOPE_NOT_ALLOWED = 9987

	// DNS_ERROR_SERVERSCOPE_IS_REFERENCED=0x00002704

	// DNS_ERROR_SERVERSCOPE_IS_REFERENCED
	// The server scope cannot be deleted because it is referenced by a DNS Policy.

	DNS_ERROR_SERVERSCOPE_IS_REFERENCED = 9988

	// DNS_ERROR_ZONESCOPE_IS_REFERENCED=0x00002705

	// DNS_ERROR_ZONESCOPE_IS_REFERENCED
	// The zone scope cannot be deleted because it is referenced by a DNS Policy.

	DNS_ERROR_ZONESCOPE_IS_REFERENCED = 9989

	// DNS_ERROR_POLICY_INVALID_CRITERIA_CLIENT_SUBNET=0x00002706

	// DNS_ERROR_POLICY_INVALID_CRITERIA_CLIENT_SUBNET
	// The criterion client subnet provided in the policy is invalid.

	DNS_ERROR_POLICY_INVALID_CRITERIA_CLIENT_SUBNET = 9990

	// DNS_ERROR_POLICY_INVALID_CRITERIA_TRANSPORT_PROTOCOL=0x00002707

	// DNS_ERROR_POLICY_INVALID_CRITERIA_TRANSPORT_PROTOCO
	// The criterion transport protocol provided in the policy is invalid.

	DNS_ERROR_POLICY_INVALID_CRITERIA_TRANSPORT_PROTOCOL = 9991

	// DNS_ERROR_POLICY_INVALID_CRITERIA_NETWORK_PROTOCOL=0x00002708

	// DNS_ERROR_POLICY_INVALID_CRITERIA_NETWORK_PROTOCO
	// The criterion network protocol provided in the policy is invalid.

	DNS_ERROR_POLICY_INVALID_CRITERIA_NETWORK_PROTOCOL = 9992

	// DNS_ERROR_POLICY_INVALID_CRITERIA_INTERFACE=0x00002709

	// DNS_ERROR_POLICY_INVALID_CRITERIA_INTERFACE
	// The criterion interface provided in the policy is invalid.

	DNS_ERROR_POLICY_INVALID_CRITERIA_INTERFACE = 9993

	// DNS_ERROR_POLICY_INVALID_CRITERIA_FQDN=0x0000270A

	// DNS_ERROR_POLICY_INVALID_CRITERIA_FQDN
	// The criterion FQDN provided in the policy is invalid.

	DNS_ERROR_POLICY_INVALID_CRITERIA_FQDN = 9994

	// DNS_ERROR_POLICY_INVALID_CRITERIA_QUERY_TYPE=0x0000270B

	// DNS_ERROR_POLICY_INVALID_CRITERIA_QUERY_TYPE
	// The criterion query type provided in the policy is invalid.

	DNS_ERROR_POLICY_INVALID_CRITERIA_QUERY_TYPE = 9995

	// DNS_ERROR_POLICY_INVALID_CRITERIA_TIME_OF_DAY=0x0000270C

	// DNS_ERROR_POLICY_INVALID_CRITERIA_TIME_OF_DAY
	// The criterion time of day provided in the policy is invalid.

	DNS_ERROR_POLICY_INVALID_CRITERIA_TIME_OF_DAY = 9996

	///////////////////////////////////////////////////
	//                                               //
	//             End of DNS Error Codes            //
	//                                               //
	//                 =9000 to=9999                 //
	///////////////////////////////////////////////////

	///////////////////////////////////////////////////
	//                                               //
	//               WinSock Error Codes             //
	//                                               //
	//                =10000 to=11999                //
	///////////////////////////////////////////////////

	// WinSock error codes are also defined in WinSock.h
	// and WinSock2.h, hence the IFDEF

	WSABASEERR = 10000

	// WSAEINTR
	// A blocking operation was interrupted by a call to WSACancelBlockingCall.

	WSAEINTR = 10004

	// WSAEBADF
	// The file handle supplied is not valid.

	WSAEBADF = 10009

	// WSAEACCES
	// An attempt was made to access a socket in a way forbidden by its access permissions.

	WSAEACCES = 10013

	// WSAEFAULT
	// The system detected an invalid pointer address in attempting to use a pointer argument in a call.

	WSAEFAULT = 10014

	// WSAEINVA
	// An invalid argument was supplied.

	WSAEINVAL = 10022

	// WSAEMFILE
	// Too many open sockets.

	WSAEMFILE = 10024

	// WSAEWOULDBLOCK
	// A non-blocking socket operation could not be completed immediately.

	WSAEWOULDBLOCK = 10035

	// WSAEINPROGRESS
	// A blocking operation is currently executing.

	WSAEINPROGRESS = 10036

	// WSAEALREADY
	// An operation was attempted on a non-blocking socket that already had an operation in progress.

	WSAEALREADY = 10037

	// WSAENOTSOCK
	// An operation was attempted on something that is not a socket.

	WSAENOTSOCK = 10038

	// WSAEDESTADDRREQ
	// A required address was omitted from an operation on a socket.

	WSAEDESTADDRREQ = 10039

	// WSAEMSGSIZE
	// A message sent on a datagram socket was larger than the internal message buffer or some other network limit, or the buffer used to receive a datagram into was smaller than the datagram itself.

	WSAEMSGSIZE = 10040

	// WSAEPROTOTYPE
	// A protocol was specified in the socket function call that does not support the semantics of the socket type requested.

	WSAEPROTOTYPE = 10041

	// WSAENOPROTOOPT
	// An unknown, invalid, or unsupported option or level was specified in a getsockopt or setsockopt call.

	WSAENOPROTOOPT = 10042

	// WSAEPROTONOSUPPORT
	// The requested protocol has not been configured into the system, or no implementation for it exists.

	WSAEPROTONOSUPPORT = 10043

	// WSAESOCKTNOSUPPORT
	// The support for the specified socket type does not exist in this address family.

	WSAESOCKTNOSUPPORT = 10044

	// WSAEOPNOTSUPP
	// The attempted operation is not supported for the type of object referenced.

	WSAEOPNOTSUPP = 10045

	// WSAEPFNOSUPPORT
	// The protocol family has not been configured into the system or no implementation for it exists.

	WSAEPFNOSUPPORT = 10046

	// WSAEAFNOSUPPORT
	// An address incompatible with the requested protocol was used.

	WSAEAFNOSUPPORT = 10047

	// WSAEADDRINUSE
	// Only one usage of each socket address (protocol/network address/port) is normally permitted.

	WSAEADDRINUSE = 10048

	// WSAEADDRNOTAVAI
	// The requested address is not valid in its context.

	WSAEADDRNOTAVAIL = 10049

	// WSAENETDOWN
	// A socket operation encountered a dead network.

	WSAENETDOWN = 10050

	// WSAENETUNREACH
	// A socket operation was attempted to an unreachable network.

	WSAENETUNREACH = 10051

	// WSAENETRESET
	// The connection has been broken due to keep-alive activity detecting a failure while the operation was in progress.

	WSAENETRESET = 10052

	// WSAECONNABORTED
	// An established connection was aborted by the software in your host machine.

	WSAECONNABORTED = 10053

	// WSAECONNRESET
	// An existing connection was forcibly closed by the remote host.

	WSAECONNRESET = 10054

	// WSAENOBUFS
	// An operation on a socket could not be performed because the system lacked sufficient buffer space or because a queue was full.

	WSAENOBUFS = 10055

	// WSAEISCONN
	// A connect request was made on an already connected socket.

	WSAEISCONN = 10056

	// WSAENOTCONN
	// A request to send or receive data was disallowed because the socket is not connected and (when sending on a datagram socket using a sendto call) no address was supplied.

	WSAENOTCONN = 10057

	// WSAESHUTDOWN
	// A request to send or receive data was disallowed because the socket had already been shut down in that direction with a previous shutdown call.

	WSAESHUTDOWN = 10058

	// WSAETOOMANYREFS
	// Too many references to some kernel object.

	WSAETOOMANYREFS = 10059

	// WSAETIMEDOUT
	// A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.

	WSAETIMEDOUT = 10060

	// WSAECONNREFUSED
	// No connection could be made because the target machine actively refused it.

	WSAECONNREFUSED = 10061

	// WSAELOOP
	// Cannot translate name.

	WSAELOOP = 10062

	// WSAENAMETOOLONG
	// Name component or name was too long.

	WSAENAMETOOLONG = 10063

	// WSAEHOSTDOWN
	// A socket operation failed because the destination host was down.

	WSAEHOSTDOWN = 10064

	// WSAEHOSTUNREACH
	// A socket operation was attempted to an unreachable host.

	WSAEHOSTUNREACH = 10065

	// WSAENOTEMPTY
	// Cannot remove a directory that is not empty.

	WSAENOTEMPTY = 10066

	// WSAEPROCLIM
	// A Windows Sockets implementation may have a limit on the number of applications that may use it simultaneously.

	WSAEPROCLIM = 10067

	// WSAEUSERS
	// Ran out of quota.

	WSAEUSERS = 10068

	// WSAEDQUOT
	// Ran out of disk quota.

	WSAEDQUOT = 10069

	// WSAESTALE
	// File handle reference is no longer available.

	WSAESTALE = 10070

	// WSAEREMOTE
	// Item is not available locally.

	WSAEREMOTE = 10071

	// WSASYSNOTREADY
	// WSAStartup cannot function at this time because the underlying system it uses to provide network services is currently unavailable.

	WSASYSNOTREADY = 10091

	// WSAVERNOTSUPPORTED
	// The Windows Sockets version requested is not supported.

	WSAVERNOTSUPPORTED = 10092

	// WSANOTINITIALISED
	// Either the application has not called WSAStartup, or WSAStartup failed.

	WSANOTINITIALISED = 10093

	// WSAEDISCON
	// Returned by WSARecv or WSARecvFrom to indicate the remote party has initiated a graceful shutdown sequence.

	WSAEDISCON = 10101

	// WSAENOMORE
	// No more results can be returned by WSALookupServiceNext.

	WSAENOMORE = 10102

	// WSAECANCELLED
	// A call to WSALookupServiceEnd was made while this call was still processing. The call has been canceled.

	WSAECANCELLED = 10103

	// WSAEINVALIDPROCTABLE
	// The procedure call table is invalid.

	WSAEINVALIDPROCTABLE = 10104

	// WSAEINVALIDPROVIDER
	// The requested service provider is invalid.

	WSAEINVALIDPROVIDER = 10105

	// WSAEPROVIDERFAILEDINIT
	// The requested service provider could not be loaded or initialized.

	WSAEPROVIDERFAILEDINIT = 10106

	// WSASYSCALLFAILURE
	// A system call has failed.

	WSASYSCALLFAILURE = 10107

	// WSASERVICE_NOT_FOUND
	// No such service is known. The service cannot be found in the specified name space.

	WSASERVICE_NOT_FOUND = 10108

	// WSATYPE_NOT_FOUND
	// The specified class was not found.

	WSATYPE_NOT_FOUND = 10109

	// WSA_E_NO_MORE
	// No more results can be returned by WSALookupServiceNext.

	WSA_E_NO_MORE = 10110

	// WSA_E_CANCELLED
	// A call to WSALookupServiceEnd was made while this call was still processing. The call has been canceled.

	WSA_E_CANCELLED = 10111

	// WSAEREFUSED
	// A database query failed because it was actively refused.

	WSAEREFUSED = 10112

	// WSAHOST_NOT_FOUND
	// No such host is known.

	WSAHOST_NOT_FOUND = 11001

	// WSATRY_AGAIN
	// This is usually a temporary error during hostname resolution and means that the local server did not receive a response from an authoritative server.

	WSATRY_AGAIN = 11002

	// WSANO_RECOVERY
	// A non-recoverable error occurred during a database lookup.

	WSANO_RECOVERY = 11003

	// WSANO_DATA
	// The requested name is valid, but no data of the requested type was found.

	WSANO_DATA = 11004

	// WSA_QOS_RECEIVERS
	// At least one reserve has arrived.

	WSA_QOS_RECEIVERS = 11005

	// WSA_QOS_SENDERS
	// At least one path has arrived.

	WSA_QOS_SENDERS = 11006

	// WSA_QOS_NO_SENDERS
	// There are no senders.

	WSA_QOS_NO_SENDERS = 11007

	// WSA_QOS_NO_RECEIVERS
	// There are no receivers.

	WSA_QOS_NO_RECEIVERS = 11008

	// WSA_QOS_REQUEST_CONFIRMED
	// Reserve has been confirmed.

	WSA_QOS_REQUEST_CONFIRMED = 11009

	// WSA_QOS_ADMISSION_FAILURE
	// Error due to lack of resources.

	WSA_QOS_ADMISSION_FAILURE = 11010

	// WSA_QOS_POLICY_FAILURE
	// Rejected for administrative reasons - bad credentials.

	WSA_QOS_POLICY_FAILURE = 11011

	// WSA_QOS_BAD_STYLE
	// Unknown or conflicting style.

	WSA_QOS_BAD_STYLE = 11012

	// WSA_QOS_BAD_OBJECT
	// Problem with some part of the filterspec or providerspecific buffer in general.

	WSA_QOS_BAD_OBJECT = 11013

	// WSA_QOS_TRAFFIC_CTRL_ERROR
	// Problem with some part of the flowspec.

	WSA_QOS_TRAFFIC_CTRL_ERROR = 11014

	// WSA_QOS_GENERIC_ERROR
	// General QOS error.

	WSA_QOS_GENERIC_ERROR = 11015

	// WSA_QOS_ESERVICETYPE
	// An invalid or unrecognized service type was found in the flowspec.

	WSA_QOS_ESERVICETYPE = 11016

	// WSA_QOS_EFLOWSPEC
	// An invalid or inconsistent flowspec was found in the QOS structure.

	WSA_QOS_EFLOWSPEC = 11017

	// WSA_QOS_EPROVSPECBUF
	// Invalid QOS provider-specific buffer.

	WSA_QOS_EPROVSPECBUF = 11018

	// WSA_QOS_EFILTERSTYLE
	// An invalid QOS filter style was used.

	WSA_QOS_EFILTERSTYLE = 11019

	// WSA_QOS_EFILTERTYPE
	// An invalid QOS filter type was used.

	WSA_QOS_EFILTERTYPE = 11020

	// WSA_QOS_EFILTERCOUNT
	// An incorrect number of QOS FILTERSPECs were specified in the FLOWDESCRIPTOR.

	WSA_QOS_EFILTERCOUNT = 11021

	// WSA_QOS_EOBJLENGTH
	// An object with an invalid ObjectLength field was specified in the QOS provider-specific buffer.

	WSA_QOS_EOBJLENGTH = 11022

	// WSA_QOS_EFLOWCOUNT
	// An incorrect number of flow descriptors was specified in the QOS structure.

	WSA_QOS_EFLOWCOUNT = 11023

	// WSA_QOS_EUNKOWNPSOBJ
	// An unrecognized object was found in the QOS provider-specific buffer.

	WSA_QOS_EUNKOWNPSOBJ = 11024

	// WSA_QOS_EPOLICYOBJ
	// An invalid policy object was found in the QOS provider-specific buffer.

	WSA_QOS_EPOLICYOBJ = 11025

	// WSA_QOS_EFLOWDESC
	// An invalid QOS flow descriptor was found in the flow descriptor list.

	WSA_QOS_EFLOWDESC = 11026

	// WSA_QOS_EPSFLOWSPEC
	// An invalid or inconsistent flowspec was found in the QOS provider specific buffer.

	WSA_QOS_EPSFLOWSPEC = 11027

	// WSA_QOS_EPSFILTERSPEC
	// An invalid FILTERSPEC was found in the QOS provider-specific buffer.

	WSA_QOS_EPSFILTERSPEC = 11028

	// WSA_QOS_ESDMODEOBJ
	// An invalid shape discard mode object was found in the QOS provider specific buffer.

	WSA_QOS_ESDMODEOBJ = 11029

	// WSA_QOS_ESHAPERATEOBJ
	// An invalid shaping rate object was found in the QOS provider-specific buffer.

	WSA_QOS_ESHAPERATEOBJ = 11030

	// WSA_QOS_RESERVED_PETYPE
	// A reserved policy element was found in the QOS provider-specific buffer.

	WSA_QOS_RESERVED_PETYPE = 11031

	// WSA_SECURE_HOST_NOT_FOUND
	// No such host is known securely.

	WSA_SECURE_HOST_NOT_FOUND = 11032

	// WSA_IPSEC_NAME_POLICY_ERROR
	// Name based IPSEC policy could not be added.

	WSA_IPSEC_NAME_POLICY_ERROR = 11033

	///////////////////////////////////////////////////
	//                                               //
	//           End of WinSock Error Codes          //
	//                                               //
	//                =10000 to=11999                //
	///////////////////////////////////////////////////

	///////////////////////////////////////////////////
	//                                               //
	//                  Available                    //
	//                                               //
	//                =12000 to=12999                //
	///////////////////////////////////////////////////

	///////////////////////////////////////////////////
	//                                               //
	//           Start of IPSec Error codes          //
	//                                               //
	//                =13000 to=13999                //
	///////////////////////////////////////////////////

	// ERROR_IPSEC_QM_POLICY_EXISTS
	// The specified quick mode policy already exists.
	ERROR_IPSEC_QM_POLICY_EXISTS = 13000

	// ERROR_IPSEC_QM_POLICY_NOT_FOUND
	// The specified quick mode policy was not found.
	ERROR_IPSEC_QM_POLICY_NOT_FOUND = 13001

	// ERROR_IPSEC_QM_POLICY_IN_USE
	// The specified quick mode policy is being used.
	ERROR_IPSEC_QM_POLICY_IN_USE = 13002

	// ERROR_IPSEC_MM_POLICY_EXISTS
	// The specified main mode policy already exists.
	ERROR_IPSEC_MM_POLICY_EXISTS = 13003

	// ERROR_IPSEC_MM_POLICY_NOT_FOUND
	// The specified main mode policy was not found
	ERROR_IPSEC_MM_POLICY_NOT_FOUND = 13004

	// ERROR_IPSEC_MM_POLICY_IN_USE
	// The specified main mode policy is being used.
	ERROR_IPSEC_MM_POLICY_IN_USE = 13005

	// ERROR_IPSEC_MM_FILTER_EXISTS
	// The specified main mode filter already exists.
	ERROR_IPSEC_MM_FILTER_EXISTS = 13006

	// ERROR_IPSEC_MM_FILTER_NOT_FOUND
	// The specified main mode filter was not found.
	ERROR_IPSEC_MM_FILTER_NOT_FOUND = 13007

	// ERROR_IPSEC_TRANSPORT_FILTER_EXISTS
	// The specified transport mode filter already exists.
	ERROR_IPSEC_TRANSPORT_FILTER_EXISTS = 13008

	// ERROR_IPSEC_TRANSPORT_FILTER_NOT_FOUND
	// The specified transport mode filter does not exist.
	ERROR_IPSEC_TRANSPORT_FILTER_NOT_FOUND = 13009

	// ERROR_IPSEC_MM_AUTH_EXISTS
	// The specified main mode authentication list exists.
	ERROR_IPSEC_MM_AUTH_EXISTS = 13010

	// ERROR_IPSEC_MM_AUTH_NOT_FOUND
	// The specified main mode authentication list was not found.
	ERROR_IPSEC_MM_AUTH_NOT_FOUND = 13011

	// ERROR_IPSEC_MM_AUTH_IN_USE
	// The specified main mode authentication list is being used.
	ERROR_IPSEC_MM_AUTH_IN_USE = 13012

	// ERROR_IPSEC_DEFAULT_MM_POLICY_NOT_FOUND
	// The specified default main mode policy was not found.
	ERROR_IPSEC_DEFAULT_MM_POLICY_NOT_FOUND = 13013

	// ERROR_IPSEC_DEFAULT_MM_AUTH_NOT_FOUND
	// The specified default main mode authentication list was not found.
	ERROR_IPSEC_DEFAULT_MM_AUTH_NOT_FOUND = 13014

	// ERROR_IPSEC_DEFAULT_QM_POLICY_NOT_FOUND
	// The specified default quick mode policy was not found.
	ERROR_IPSEC_DEFAULT_QM_POLICY_NOT_FOUND = 13015

	// ERROR_IPSEC_TUNNEL_FILTER_EXISTS
	// The specified tunnel mode filter exists.
	ERROR_IPSEC_TUNNEL_FILTER_EXISTS = 13016

	// ERROR_IPSEC_TUNNEL_FILTER_NOT_FOUND
	// The specified tunnel mode filter was not found.
	ERROR_IPSEC_TUNNEL_FILTER_NOT_FOUND = 13017

	// ERROR_IPSEC_MM_FILTER_PENDING_DELETION
	// The Main Mode filter is pending deletion.
	ERROR_IPSEC_MM_FILTER_PENDING_DELETION = 13018

	// ERROR_IPSEC_TRANSPORT_FILTER_PENDING_DELETION
	// The transport filter is pending deletion.
	ERROR_IPSEC_TRANSPORT_FILTER_PENDING_DELETION = 13019

	// ERROR_IPSEC_TUNNEL_FILTER_PENDING_DELETION
	// The tunnel filter is pending deletion.
	ERROR_IPSEC_TUNNEL_FILTER_PENDING_DELETION = 13020

	// ERROR_IPSEC_MM_POLICY_PENDING_DELETION
	// The Main Mode policy is pending deletion.
	ERROR_IPSEC_MM_POLICY_PENDING_DELETION = 13021

	// ERROR_IPSEC_MM_AUTH_PENDING_DELETION
	// The Main Mode authentication bundle is pending deletion.
	ERROR_IPSEC_MM_AUTH_PENDING_DELETION = 13022

	// ERROR_IPSEC_QM_POLICY_PENDING_DELETION
	// The Quick Mode policy is pending deletion.
	ERROR_IPSEC_QM_POLICY_PENDING_DELETION = 13023

	// WARNING_IPSEC_MM_POLICY_PRUNED
	// The Main Mode policy was successfully added, but some of the requested offers are not supported.

	WARNING_IPSEC_MM_POLICY_PRUNED = 13024

	// WARNING_IPSEC_QM_POLICY_PRUNED
	// The Quick Mode policy was successfully added, but some of the requested offers are not supported.

	WARNING_IPSEC_QM_POLICY_PRUNED = 13025

	// ERROR_IPSEC_IKE_NEG_STATUS_BEGIN
	//  ERROR_IPSEC_IKE_NEG_STATUS_BEGIN
	ERROR_IPSEC_IKE_NEG_STATUS_BEGIN = 13800

	// ERROR_IPSEC_IKE_AUTH_FAI
	// IKE authentication credentials are unacceptable
	ERROR_IPSEC_IKE_AUTH_FAIL = 13801

	// ERROR_IPSEC_IKE_ATTRIB_FAI
	// IKE security attributes are unacceptable
	ERROR_IPSEC_IKE_ATTRIB_FAIL = 13802

	// ERROR_IPSEC_IKE_NEGOTIATION_PENDING
	// IKE Negotiation in progress
	ERROR_IPSEC_IKE_NEGOTIATION_PENDING = 13803

	// ERROR_IPSEC_IKE_GENERAL_PROCESSING_ERROR
	// General processing error
	ERROR_IPSEC_IKE_GENERAL_PROCESSING_ERROR = 13804

	// ERROR_IPSEC_IKE_TIMED_OUT
	// Negotiation timed out
	ERROR_IPSEC_IKE_TIMED_OUT = 13805

	// ERROR_IPSEC_IKE_NO_CERT
	// IKE failed to find valid machine certificate. Contact your Network Security Administrator about installing a valid certificate in the appropriate Certificate Store.
	ERROR_IPSEC_IKE_NO_CERT = 13806

	// ERROR_IPSEC_IKE_SA_DELETED
	// IKE SA deleted by peer before establishment completed
	ERROR_IPSEC_IKE_SA_DELETED = 13807

	// ERROR_IPSEC_IKE_SA_REAPED
	// IKE SA deleted before establishment completed
	ERROR_IPSEC_IKE_SA_REAPED = 13808

	// ERROR_IPSEC_IKE_MM_ACQUIRE_DROP
	// Negotiation request sat in Queue too long
	ERROR_IPSEC_IKE_MM_ACQUIRE_DROP = 13809

	// ERROR_IPSEC_IKE_QM_ACQUIRE_DROP
	// Negotiation request sat in Queue too long
	ERROR_IPSEC_IKE_QM_ACQUIRE_DROP = 13810

	// ERROR_IPSEC_IKE_QUEUE_DROP_MM
	// Negotiation request sat in Queue too long
	ERROR_IPSEC_IKE_QUEUE_DROP_MM = 13811

	// ERROR_IPSEC_IKE_QUEUE_DROP_NO_MM
	// Negotiation request sat in Queue too long
	ERROR_IPSEC_IKE_QUEUE_DROP_NO_MM = 13812

	// ERROR_IPSEC_IKE_DROP_NO_RESPONSE
	// No response from peer
	ERROR_IPSEC_IKE_DROP_NO_RESPONSE = 13813

	// ERROR_IPSEC_IKE_MM_DELAY_DROP
	// Negotiation took too long
	ERROR_IPSEC_IKE_MM_DELAY_DROP = 13814

	// ERROR_IPSEC_IKE_QM_DELAY_DROP
	// Negotiation took too long
	ERROR_IPSEC_IKE_QM_DELAY_DROP = 13815

	// ERROR_IPSEC_IKE_ERROR
	// Unknown error occurred
	ERROR_IPSEC_IKE_ERROR = 13816

	// ERROR_IPSEC_IKE_CRL_FAILED
	// Certificate Revocation Check failed
	ERROR_IPSEC_IKE_CRL_FAILED = 13817

	// ERROR_IPSEC_IKE_INVALID_KEY_USAGE
	// Invalid certificate key usage
	ERROR_IPSEC_IKE_INVALID_KEY_USAGE = 13818

	// ERROR_IPSEC_IKE_INVALID_CERT_TYPE
	// Invalid certificate type
	ERROR_IPSEC_IKE_INVALID_CERT_TYPE = 13819

	// ERROR_IPSEC_IKE_NO_PRIVATE_KEY
	// IKE negotiation failed because the machine certificate used does not have a private key. IPsec certificates require a private key. Contact your Network Security administrator about replacing with a certificate that has a private key.
	ERROR_IPSEC_IKE_NO_PRIVATE_KEY = 13820

	// ERROR_IPSEC_IKE_SIMULTANEOUS_REKEY
	// Simultaneous rekeys were detected.
	ERROR_IPSEC_IKE_SIMULTANEOUS_REKEY = 13821

	// ERROR_IPSEC_IKE_DH_FAI
	// Failure in Diffie-Hellman computation
	ERROR_IPSEC_IKE_DH_FAIL = 13822

	// ERROR_IPSEC_IKE_CRITICAL_PAYLOAD_NOT_RECOGNIZED
	// Don't know how to process critical payload
	ERROR_IPSEC_IKE_CRITICAL_PAYLOAD_NOT_RECOGNIZED = 13823

	// ERROR_IPSEC_IKE_INVALID_HEADER
	// Invalid header
	ERROR_IPSEC_IKE_INVALID_HEADER = 13824

	// ERROR_IPSEC_IKE_NO_POLICY
	// No policy configured
	ERROR_IPSEC_IKE_NO_POLICY = 13825

	// ERROR_IPSEC_IKE_INVALID_SIGNATURE
	// Failed to verify signature
	ERROR_IPSEC_IKE_INVALID_SIGNATURE = 13826

	// ERROR_IPSEC_IKE_KERBEROS_ERROR
	// Failed to authenticate using Kerberos
	ERROR_IPSEC_IKE_KERBEROS_ERROR = 13827

	// ERROR_IPSEC_IKE_NO_PUBLIC_KEY
	// Peer's certificate did not have a public key
	ERROR_IPSEC_IKE_NO_PUBLIC_KEY = 13828

	// These must stay as a unit.

	// ERROR_IPSEC_IKE_PROCESS_ERR
	// Error processing error payload
	ERROR_IPSEC_IKE_PROCESS_ERR = 13829

	// ERROR_IPSEC_IKE_PROCESS_ERR_SA
	// Error processing SA payload
	ERROR_IPSEC_IKE_PROCESS_ERR_SA = 13830

	// ERROR_IPSEC_IKE_PROCESS_ERR_PROP
	// Error processing Proposal payload
	ERROR_IPSEC_IKE_PROCESS_ERR_PROP = 13831

	// ERROR_IPSEC_IKE_PROCESS_ERR_TRANS
	// Error processing Transform payload
	ERROR_IPSEC_IKE_PROCESS_ERR_TRANS = 13832

	// ERROR_IPSEC_IKE_PROCESS_ERR_KE
	// Error processing KE payload
	ERROR_IPSEC_IKE_PROCESS_ERR_KE = 13833

	// ERROR_IPSEC_IKE_PROCESS_ERR_ID
	// Error processing ID payload
	ERROR_IPSEC_IKE_PROCESS_ERR_ID = 13834

	// ERROR_IPSEC_IKE_PROCESS_ERR_CERT
	// Error processing Cert payload
	ERROR_IPSEC_IKE_PROCESS_ERR_CERT = 13835

	// ERROR_IPSEC_IKE_PROCESS_ERR_CERT_REQ
	// Error processing Certificate Request payload
	ERROR_IPSEC_IKE_PROCESS_ERR_CERT_REQ = 13836

	// ERROR_IPSEC_IKE_PROCESS_ERR_HASH
	// Error processing Hash payload
	ERROR_IPSEC_IKE_PROCESS_ERR_HASH = 13837

	// ERROR_IPSEC_IKE_PROCESS_ERR_SIG
	// Error processing Signature payload
	ERROR_IPSEC_IKE_PROCESS_ERR_SIG = 13838

	// ERROR_IPSEC_IKE_PROCESS_ERR_NONCE
	// Error processing Nonce payload
	ERROR_IPSEC_IKE_PROCESS_ERR_NONCE = 13839

	// ERROR_IPSEC_IKE_PROCESS_ERR_NOTIFY
	// Error processing Notify payload
	ERROR_IPSEC_IKE_PROCESS_ERR_NOTIFY = 13840

	// ERROR_IPSEC_IKE_PROCESS_ERR_DELETE
	// Error processing Delete Payload
	ERROR_IPSEC_IKE_PROCESS_ERR_DELETE = 13841

	// ERROR_IPSEC_IKE_PROCESS_ERR_VENDOR
	// Error processing VendorId payload
	ERROR_IPSEC_IKE_PROCESS_ERR_VENDOR = 13842

	// ERROR_IPSEC_IKE_INVALID_PAYLOAD
	// Invalid payload received
	ERROR_IPSEC_IKE_INVALID_PAYLOAD = 13843

	// ERROR_IPSEC_IKE_LOAD_SOFT_SA
	// Soft SA loaded
	ERROR_IPSEC_IKE_LOAD_SOFT_SA = 13844

	// ERROR_IPSEC_IKE_SOFT_SA_TORN_DOWN
	// Soft SA torn down
	ERROR_IPSEC_IKE_SOFT_SA_TORN_DOWN = 13845

	// ERROR_IPSEC_IKE_INVALID_COOKIE
	// Invalid cookie received.
	ERROR_IPSEC_IKE_INVALID_COOKIE = 13846

	// ERROR_IPSEC_IKE_NO_PEER_CERT
	// Peer failed to send valid machine certificate
	ERROR_IPSEC_IKE_NO_PEER_CERT = 13847

	// ERROR_IPSEC_IKE_PEER_CRL_FAILED
	// Certification Revocation check of peer's certificate failed
	ERROR_IPSEC_IKE_PEER_CRL_FAILED = 13848

	// ERROR_IPSEC_IKE_POLICY_CHANGE
	// New policy invalidated SAs formed with old policy
	ERROR_IPSEC_IKE_POLICY_CHANGE = 13849

	// ERROR_IPSEC_IKE_NO_MM_POLICY
	// There is no available Main Mode IKE policy.
	ERROR_IPSEC_IKE_NO_MM_POLICY = 13850

	// ERROR_IPSEC_IKE_NOTCBPRIV
	// Failed to enabled TCB privilege.
	ERROR_IPSEC_IKE_NOTCBPRIV = 13851

	// ERROR_IPSEC_IKE_SECLOADFAI
	// Failed to load SECURITY.DLL.
	ERROR_IPSEC_IKE_SECLOADFAIL = 13852

	// ERROR_IPSEC_IKE_FAILSSPINIT
	// Failed to obtain security function table dispatch address from SSPI.
	ERROR_IPSEC_IKE_FAILSSPINIT = 13853

	// ERROR_IPSEC_IKE_FAILQUERYSSP
	// Failed to query Kerberos package to obtain max token size.
	ERROR_IPSEC_IKE_FAILQUERYSSP = 13854

	// ERROR_IPSEC_IKE_SRVACQFAI
	// Failed to obtain Kerberos server credentials for ISAKMP/ERROR_IPSEC_IKE service. Kerberos authentication will not function. The most likely reason for this is lack of domain membership. This is normal if your computer is a member of a workgroup.
	ERROR_IPSEC_IKE_SRVACQFAIL = 13855

	// ERROR_IPSEC_IKE_SRVQUERYCRED
	// Failed to determine SSPI principal name for ISAKMP/ERROR_IPSEC_IKE service (QueryCredentialsAttributes).
	ERROR_IPSEC_IKE_SRVQUERYCRED = 13856

	// ERROR_IPSEC_IKE_GETSPIFAI
	// Failed to obtain new SPI for the inbound SA from IPsec driver. The most common cause for this is that the driver does not have the correct filter. Check your policy to verify the filters.
	ERROR_IPSEC_IKE_GETSPIFAIL = 13857

	// ERROR_IPSEC_IKE_INVALID_FILTER
	// Given filter is invalid
	ERROR_IPSEC_IKE_INVALID_FILTER = 13858

	// ERROR_IPSEC_IKE_OUT_OF_MEMORY
	// Memory allocation failed.
	ERROR_IPSEC_IKE_OUT_OF_MEMORY = 13859

	// ERROR_IPSEC_IKE_ADD_UPDATE_KEY_FAILED
	// Failed to add Security Association to IPsec Driver. The most common cause for this is if the IKE negotiation took too long to complete. If the problem persists, reduce the load on the faulting machine.
	ERROR_IPSEC_IKE_ADD_UPDATE_KEY_FAILED = 13860

	// ERROR_IPSEC_IKE_INVALID_POLICY
	// Invalid policy
	ERROR_IPSEC_IKE_INVALID_POLICY = 13861

	// ERROR_IPSEC_IKE_UNKNOWN_DOI
	// Invalid DOI
	ERROR_IPSEC_IKE_UNKNOWN_DOI = 13862

	// ERROR_IPSEC_IKE_INVALID_SITUATION
	// Invalid situation
	ERROR_IPSEC_IKE_INVALID_SITUATION = 13863

	// ERROR_IPSEC_IKE_DH_FAILURE
	// Diffie-Hellman failure
	ERROR_IPSEC_IKE_DH_FAILURE = 13864

	// ERROR_IPSEC_IKE_INVALID_GROUP
	// Invalid Diffie-Hellman group
	ERROR_IPSEC_IKE_INVALID_GROUP = 13865

	// ERROR_IPSEC_IKE_ENCRYPT
	// Error encrypting payload
	ERROR_IPSEC_IKE_ENCRYPT = 13866

	// ERROR_IPSEC_IKE_DECRYPT
	// Error decrypting payload
	ERROR_IPSEC_IKE_DECRYPT = 13867

	// ERROR_IPSEC_IKE_POLICY_MATCH
	// Policy match error
	ERROR_IPSEC_IKE_POLICY_MATCH = 13868

	// ERROR_IPSEC_IKE_UNSUPPORTED_ID
	// Unsupported ID
	ERROR_IPSEC_IKE_UNSUPPORTED_ID = 13869

	// ERROR_IPSEC_IKE_INVALID_HASH
	// Hash verification failed
	ERROR_IPSEC_IKE_INVALID_HASH = 13870

	// ERROR_IPSEC_IKE_INVALID_HASH_ALG
	// Invalid hash algorithm
	ERROR_IPSEC_IKE_INVALID_HASH_ALG = 13871

	// ERROR_IPSEC_IKE_INVALID_HASH_SIZE
	// Invalid hash size
	ERROR_IPSEC_IKE_INVALID_HASH_SIZE = 13872

	// ERROR_IPSEC_IKE_INVALID_ENCRYPT_ALG
	// Invalid encryption algorithm
	ERROR_IPSEC_IKE_INVALID_ENCRYPT_ALG = 13873

	// ERROR_IPSEC_IKE_INVALID_AUTH_ALG
	// Invalid authentication algorithm
	ERROR_IPSEC_IKE_INVALID_AUTH_ALG = 13874

	// ERROR_IPSEC_IKE_INVALID_SIG
	// Invalid certificate signature
	ERROR_IPSEC_IKE_INVALID_SIG = 13875

	// ERROR_IPSEC_IKE_LOAD_FAILED
	// Load failed
	ERROR_IPSEC_IKE_LOAD_FAILED = 13876

	// ERROR_IPSEC_IKE_RPC_DELETE
	// Deleted via RPC cal
	ERROR_IPSEC_IKE_RPC_DELETE = 13877

	// ERROR_IPSEC_IKE_BENIGN_REINIT
	// Temporary state created to perform reinitialization. This is not a real failure.
	ERROR_IPSEC_IKE_BENIGN_REINIT = 13878

	// ERROR_IPSEC_IKE_INVALID_RESPONDER_LIFETIME_NOTIFY
	// The lifetime value received in the Responder Lifetime Notify is below the Windows=2000 configured minimum value. Please fix the policy on the peer machine.
	ERROR_IPSEC_IKE_INVALID_RESPONDER_LIFETIME_NOTIFY = 13879

	// ERROR_IPSEC_IKE_INVALID_MAJOR_VERSION
	// The recipient cannot handle version of IKE specified in the header.
	ERROR_IPSEC_IKE_INVALID_MAJOR_VERSION = 13880

	// ERROR_IPSEC_IKE_INVALID_CERT_KEYLEN
	// Key length in certificate is too small for configured security requirements.
	ERROR_IPSEC_IKE_INVALID_CERT_KEYLEN = 13881

	// ERROR_IPSEC_IKE_MM_LIMIT
	// Max number of established MM SAs to peer exceeded.
	ERROR_IPSEC_IKE_MM_LIMIT = 13882

	// ERROR_IPSEC_IKE_NEGOTIATION_DISABLED
	// IKE received a policy that disables negotiation.
	ERROR_IPSEC_IKE_NEGOTIATION_DISABLED = 13883

	// ERROR_IPSEC_IKE_QM_LIMIT
	// Reached maximum quick mode limit for the main mode. New main mode will be started.
	ERROR_IPSEC_IKE_QM_LIMIT = 13884

	// ERROR_IPSEC_IKE_MM_EXPIRED
	// Main mode SA lifetime expired or peer sent a main mode delete.
	ERROR_IPSEC_IKE_MM_EXPIRED = 13885

	// ERROR_IPSEC_IKE_PEER_MM_ASSUMED_INVALID
	// Main mode SA assumed to be invalid because peer stopped responding.
	ERROR_IPSEC_IKE_PEER_MM_ASSUMED_INVALID = 13886

	// ERROR_IPSEC_IKE_CERT_CHAIN_POLICY_MISMATCH
	// Certificate doesn't chain to a trusted root in IPsec policy.
	ERROR_IPSEC_IKE_CERT_CHAIN_POLICY_MISMATCH = 13887

	// ERROR_IPSEC_IKE_UNEXPECTED_MESSAGE_ID
	// Received unexpected message ID.
	ERROR_IPSEC_IKE_UNEXPECTED_MESSAGE_ID = 13888

	// ERROR_IPSEC_IKE_INVALID_AUTH_PAYLOAD
	// Received invalid authentication offers.
	ERROR_IPSEC_IKE_INVALID_AUTH_PAYLOAD = 13889

	// ERROR_IPSEC_IKE_DOS_COOKIE_SENT
	// Sent DoS cookie notify to initiator.
	ERROR_IPSEC_IKE_DOS_COOKIE_SENT = 13890

	// ERROR_IPSEC_IKE_SHUTTING_DOWN
	// IKE service is shutting down.
	ERROR_IPSEC_IKE_SHUTTING_DOWN = 13891

	// ERROR_IPSEC_IKE_CGA_AUTH_FAILED
	// Could not verify binding between CGA address and certificate.
	ERROR_IPSEC_IKE_CGA_AUTH_FAILED = 13892

	// ERROR_IPSEC_IKE_PROCESS_ERR_NATOA
	// Error processing NatOA payload.
	ERROR_IPSEC_IKE_PROCESS_ERR_NATOA = 13893

	// ERROR_IPSEC_IKE_INVALID_MM_FOR_QM
	// Parameters of the main mode are invalid for this quick mode.
	ERROR_IPSEC_IKE_INVALID_MM_FOR_QM = 13894

	// ERROR_IPSEC_IKE_QM_EXPIRED
	// Quick mode SA was expired by IPsec driver.
	ERROR_IPSEC_IKE_QM_EXPIRED = 13895

	// ERROR_IPSEC_IKE_TOO_MANY_FILTERS
	// Too many dynamically added IKEEXT filters were detected.
	ERROR_IPSEC_IKE_TOO_MANY_FILTERS = 13896

	// Do NOT change this final value.  It is used in a public API structure

	// ERROR_IPSEC_IKE_NEG_STATUS_END
	//  ERROR_IPSEC_IKE_NEG_STATUS_END
	ERROR_IPSEC_IKE_NEG_STATUS_END = 13897

	// ERROR_IPSEC_IKE_KILL_DUMMY_NAP_TUNNE
	// NAP reauth succeeded and must delete the dummy NAP IKEv2 tunnel.
	ERROR_IPSEC_IKE_KILL_DUMMY_NAP_TUNNEL = 13898

	// ERROR_IPSEC_IKE_INNER_IP_ASSIGNMENT_FAILURE
	// Error in assigning inner IP address to initiator in tunnel mode.
	ERROR_IPSEC_IKE_INNER_IP_ASSIGNMENT_FAILURE = 13899

	// ERROR_IPSEC_IKE_REQUIRE_CP_PAYLOAD_MISSING
	// Require configuration payload missing.
	ERROR_IPSEC_IKE_REQUIRE_CP_PAYLOAD_MISSING = 13900

	// ERROR_IPSEC_KEY_MODULE_IMPERSONATION_NEGOTIATION_PENDING
	// A negotiation running as the security principle who issued the connection is in progress
	ERROR_IPSEC_KEY_MODULE_IMPERSONATION_NEGOTIATION_PENDING = 13901

	// ERROR_IPSEC_IKE_COEXISTENCE_SUPPRESS
	// SA was deleted due to IKEv1/AuthIP co-existence suppress check.
	ERROR_IPSEC_IKE_COEXISTENCE_SUPPRESS = 13902

	// ERROR_IPSEC_IKE_RATELIMIT_DROP
	// Incoming SA request was dropped due to peer IP address rate limiting.
	ERROR_IPSEC_IKE_RATELIMIT_DROP = 13903

	// ERROR_IPSEC_IKE_PEER_DOESNT_SUPPORT_MOBIKE
	// Peer does not support MOBIKE.
	ERROR_IPSEC_IKE_PEER_DOESNT_SUPPORT_MOBIKE = 13904

	// ERROR_IPSEC_IKE_AUTHORIZATION_FAILURE
	// SA establishment is not authorized.
	ERROR_IPSEC_IKE_AUTHORIZATION_FAILURE = 13905

	// ERROR_IPSEC_IKE_STRONG_CRED_AUTHORIZATION_FAILURE
	// SA establishment is not authorized because there is not a sufficiently strong PKINIT-based credential.
	ERROR_IPSEC_IKE_STRONG_CRED_AUTHORIZATION_FAILURE = 13906

	// ERROR_IPSEC_IKE_AUTHORIZATION_FAILURE_WITH_OPTIONAL_RETRY
	// SA establishment is not authorized.  You may need to enter updated or different credentials such as a smartcard.
	ERROR_IPSEC_IKE_AUTHORIZATION_FAILURE_WITH_OPTIONAL_RETRY = 13907

	// ERROR_IPSEC_IKE_STRONG_CRED_AUTHORIZATION_AND_CERTMAP_FAILURE
	// SA establishment is not authorized because there is not a sufficiently strong PKINIT-based credential. This might be related to certificate-to-account mapping failure for the SA.
	ERROR_IPSEC_IKE_STRONG_CRED_AUTHORIZATION_AND_CERTMAP_FAILURE = 13908

	// Extended upper bound for IKE errors to accommodate new errors

	// ERROR_IPSEC_IKE_NEG_STATUS_EXTENDED_END
	//  ERROR_IPSEC_IKE_NEG_STATUS_EXTENDED_END
	ERROR_IPSEC_IKE_NEG_STATUS_EXTENDED_END = 13909

	// Following error codes are returned by IPsec kernel.

	// ERROR_IPSEC_BAD_SPI
	// The SPI in the packet does not match a valid IPsec SA.
	ERROR_IPSEC_BAD_SPI = 13910

	// ERROR_IPSEC_SA_LIFETIME_EXPIRED
	// Packet was received on an IPsec SA whose lifetime has expired.
	ERROR_IPSEC_SA_LIFETIME_EXPIRED = 13911

	// ERROR_IPSEC_WRONG_SA
	// Packet was received on an IPsec SA that does not match the packet characteristics.
	ERROR_IPSEC_WRONG_SA = 13912

	// ERROR_IPSEC_REPLAY_CHECK_FAILED
	// Packet sequence number replay check failed.
	ERROR_IPSEC_REPLAY_CHECK_FAILED = 13913

	// ERROR_IPSEC_INVALID_PACKET
	// IPsec header and/or trailer in the packet is invalid.
	ERROR_IPSEC_INVALID_PACKET = 13914

	// ERROR_IPSEC_INTEGRITY_CHECK_FAILED
	// IPsec integrity check failed.
	ERROR_IPSEC_INTEGRITY_CHECK_FAILED = 13915

	// ERROR_IPSEC_CLEAR_TEXT_DROP
	// IPsec dropped a clear text packet.
	ERROR_IPSEC_CLEAR_TEXT_DROP = 13916

	// ERROR_IPSEC_AUTH_FIREWALL_DROP
	// IPsec dropped an incoming ESP packet in authenticated firewall mode. This drop is benign.
	ERROR_IPSEC_AUTH_FIREWALL_DROP = 13917

	// ERROR_IPSEC_THROTTLE_DROP
	// IPsec dropped a packet due to DoS throttling.
	ERROR_IPSEC_THROTTLE_DROP = 13918

	// ERROR_IPSEC_DOSP_BLOCK
	// IPsec DoS Protection matched an explicit block rule.
	ERROR_IPSEC_DOSP_BLOCK = 13925

	// ERROR_IPSEC_DOSP_RECEIVED_MULTICAST
	// IPsec DoS Protection received an IPsec specific multicast packet which is not allowed.
	ERROR_IPSEC_DOSP_RECEIVED_MULTICAST = 13926

	// ERROR_IPSEC_DOSP_INVALID_PACKET
	// IPsec DoS Protection received an incorrectly formatted packet.
	ERROR_IPSEC_DOSP_INVALID_PACKET = 13927

	// ERROR_IPSEC_DOSP_STATE_LOOKUP_FAILED
	// IPsec DoS Protection failed to look up state.
	ERROR_IPSEC_DOSP_STATE_LOOKUP_FAILED = 13928

	// ERROR_IPSEC_DOSP_MAX_ENTRIES
	// IPsec DoS Protection failed to create state because the maximum number of entries allowed by policy has been reached.
	ERROR_IPSEC_DOSP_MAX_ENTRIES = 13929

	// ERROR_IPSEC_DOSP_KEYMOD_NOT_ALLOWED
	// IPsec DoS Protection received an IPsec negotiation packet for a keying module which is not allowed by policy.
	ERROR_IPSEC_DOSP_KEYMOD_NOT_ALLOWED = 13930

	// ERROR_IPSEC_DOSP_NOT_INSTALLED
	// IPsec DoS Protection has not been enabled.
	ERROR_IPSEC_DOSP_NOT_INSTALLED = 13931

	// ERROR_IPSEC_DOSP_MAX_PER_IP_RATELIMIT_QUEUES
	// IPsec DoS Protection failed to create a per internal IP rate limit queue because the maximum number of queues allowed by policy has been reached.
	ERROR_IPSEC_DOSP_MAX_PER_IP_RATELIMIT_QUEUES = 13932

	///////////////////////////////////////////////////
	//                                               //
	//           End of IPSec Error codes            //
	//                                               //
	//                =13000 to=13999                //
	///////////////////////////////////////////////////

	///////////////////////////////////////////////////
	//                                               //
	//         Start of Side By Side Error Codes     //
	//                                               //
	//                =14000 to=14999                //
	///////////////////////////////////////////////////

	// ERROR_SXS_SECTION_NOT_FOUND
	// The requested section was not present in the activation context.
	ERROR_SXS_SECTION_NOT_FOUND = 14000

	// ERROR_SXS_CANT_GEN_ACTCTX
	// The application has failed to start because its side-by-side configuration is incorrect. Please see the application event log or use the command-line sxstrace.exe tool for more detail.
	ERROR_SXS_CANT_GEN_ACTCTX = 14001

	// ERROR_SXS_INVALID_ACTCTXDATA_FORMAT
	// The application binding data format is invalid.
	ERROR_SXS_INVALID_ACTCTXDATA_FORMAT = 14002

	// ERROR_SXS_ASSEMBLY_NOT_FOUND
	// The referenced assembly is not installed on your system.
	ERROR_SXS_ASSEMBLY_NOT_FOUND = 14003

	// ERROR_SXS_MANIFEST_FORMAT_ERROR
	// The manifest file does not begin with the required tag and format information.
	ERROR_SXS_MANIFEST_FORMAT_ERROR = 14004

	// ERROR_SXS_MANIFEST_PARSE_ERROR
	// The manifest file contains one or more syntax errors.
	ERROR_SXS_MANIFEST_PARSE_ERROR = 14005

	// ERROR_SXS_ACTIVATION_CONTEXT_DISABLED
	// The application attempted to activate a disabled activation context.
	ERROR_SXS_ACTIVATION_CONTEXT_DISABLED = 14006

	// ERROR_SXS_KEY_NOT_FOUND
	// The requested lookup key was not found in any active activation context.
	ERROR_SXS_KEY_NOT_FOUND = 14007

	// ERROR_SXS_VERSION_CONFLICT
	// A component version required by the application conflicts with another component version already active.
	ERROR_SXS_VERSION_CONFLICT = 14008

	// ERROR_SXS_WRONG_SECTION_TYPE
	// The type requested activation context section does not match the query API used.
	ERROR_SXS_WRONG_SECTION_TYPE = 14009

	// ERROR_SXS_THREAD_QUERIES_DISABLED
	// Lack of system resources has required isolated activation to be disabled for the current thread of execution.
	ERROR_SXS_THREAD_QUERIES_DISABLED = 14010

	// ERROR_SXS_PROCESS_DEFAULT_ALREADY_SET
	// An attempt to set the process default activation context failed because the process default activation context was already set.
	ERROR_SXS_PROCESS_DEFAULT_ALREADY_SET = 14011

	// ERROR_SXS_UNKNOWN_ENCODING_GROUP
	// The encoding group identifier specified is not recognized.
	ERROR_SXS_UNKNOWN_ENCODING_GROUP = 14012

	// ERROR_SXS_UNKNOWN_ENCODING
	// The encoding requested is not recognized.
	ERROR_SXS_UNKNOWN_ENCODING = 14013

	// ERROR_SXS_INVALID_XML_NAMESPACE_URI
	// The manifest contains a reference to an invalid URI.
	ERROR_SXS_INVALID_XML_NAMESPACE_URI = 14014

	// ERROR_SXS_ROOT_MANIFEST_DEPENDENCY_NOT_INSTALLED
	// The application manifest contains a reference to a dependent assembly which is not installed
	ERROR_SXS_ROOT_MANIFEST_DEPENDENCY_NOT_INSTALLED = 14015

	// ERROR_SXS_LEAF_MANIFEST_DEPENDENCY_NOT_INSTALLED
	// The manifest for an assembly used by the application has a reference to a dependent assembly which is not installed
	ERROR_SXS_LEAF_MANIFEST_DEPENDENCY_NOT_INSTALLED = 14016

	// ERROR_SXS_INVALID_ASSEMBLY_IDENTITY_ATTRIBUTE
	// The manifest contains an attribute for the assembly identity which is not valid.
	ERROR_SXS_INVALID_ASSEMBLY_IDENTITY_ATTRIBUTE = 14017

	// ERROR_SXS_MANIFEST_MISSING_REQUIRED_DEFAULT_NAMESPACE
	// The manifest is missing the required default namespace specification on the assembly element.
	ERROR_SXS_MANIFEST_MISSING_REQUIRED_DEFAULT_NAMESPACE = 14018

	// ERROR_SXS_MANIFEST_INVALID_REQUIRED_DEFAULT_NAMESPACE
	// The manifest has a default namespace specified on the assembly element but its value is not "urn:schemas-microsoft-com:asm.v1".
	ERROR_SXS_MANIFEST_INVALID_REQUIRED_DEFAULT_NAMESPACE = 14019

	// ERROR_SXS_PRIVATE_MANIFEST_CROSS_PATH_WITH_REPARSE_POINT
	// The private manifest probed has crossed a path with an unsupported reparse point.
	ERROR_SXS_PRIVATE_MANIFEST_CROSS_PATH_WITH_REPARSE_POINT = 14020

	// ERROR_SXS_DUPLICATE_DLL_NAME
	// Two or more components referenced directly or indirectly by the application manifest have files by the same name.
	ERROR_SXS_DUPLICATE_DLL_NAME = 14021

	// ERROR_SXS_DUPLICATE_WINDOWCLASS_NAME
	// Two or more components referenced directly or indirectly by the application manifest have window classes with the same name.
	ERROR_SXS_DUPLICATE_WINDOWCLASS_NAME = 14022

	// ERROR_SXS_DUPLICATE_CLSID
	// Two or more components referenced directly or indirectly by the application manifest have the same COM server CLSIDs.
	ERROR_SXS_DUPLICATE_CLSID = 14023

	// ERROR_SXS_DUPLICATE_IID
	// Two or more components referenced directly or indirectly by the application manifest have proxies for the same COM interface IIDs.
	ERROR_SXS_DUPLICATE_IID = 14024

	// ERROR_SXS_DUPLICATE_TLBID
	// Two or more components referenced directly or indirectly by the application manifest have the same COM type library TLBIDs.
	ERROR_SXS_DUPLICATE_TLBID = 14025

	// ERROR_SXS_DUPLICATE_PROGID
	// Two or more components referenced directly or indirectly by the application manifest have the same COM ProgIDs.
	ERROR_SXS_DUPLICATE_PROGID = 14026

	// ERROR_SXS_DUPLICATE_ASSEMBLY_NAME
	// Two or more components referenced directly or indirectly by the application manifest are different versions of the same component which is not permitted.
	ERROR_SXS_DUPLICATE_ASSEMBLY_NAME = 14027

	// ERROR_SXS_FILE_HASH_MISMATCH
	// A component's file does not match the verification information present in the component manifest.
	ERROR_SXS_FILE_HASH_MISMATCH = 14028

	// ERROR_SXS_POLICY_PARSE_ERROR
	// The policy manifest contains one or more syntax errors.
	ERROR_SXS_POLICY_PARSE_ERROR = 14029

	// ERROR_SXS_XML_E_MISSINGQUOTE
	// Manifest Parse Error : A string literal was expected, but no opening quote character was found.
	ERROR_SXS_XML_E_MISSINGQUOTE = 14030

	// ERROR_SXS_XML_E_COMMENTSYNTAX
	// Manifest Parse Error : Incorrect syntax was used in a comment.
	ERROR_SXS_XML_E_COMMENTSYNTAX = 14031

	// ERROR_SXS_XML_E_BADSTARTNAMECHAR
	// Manifest Parse Error : A name was started with an invalid character.
	ERROR_SXS_XML_E_BADSTARTNAMECHAR = 14032

	// ERROR_SXS_XML_E_BADNAMECHAR
	// Manifest Parse Error : A name contained an invalid character.
	ERROR_SXS_XML_E_BADNAMECHAR = 14033

	// ERROR_SXS_XML_E_BADCHARINSTRING
	// Manifest Parse Error : A string literal contained an invalid character.
	ERROR_SXS_XML_E_BADCHARINSTRING = 14034

	// ERROR_SXS_XML_E_XMLDECLSYNTAX
	// Manifest Parse Error : Invalid syntax for an xml declaration.
	ERROR_SXS_XML_E_XMLDECLSYNTAX = 14035

	// ERROR_SXS_XML_E_BADCHARDATA
	// Manifest Parse Error : An Invalid character was found in text content.
	ERROR_SXS_XML_E_BADCHARDATA = 14036

	// ERROR_SXS_XML_E_MISSINGWHITESPACE
	// Manifest Parse Error : Required white space was missing.
	ERROR_SXS_XML_E_MISSINGWHITESPACE = 14037

	// ERROR_SXS_XML_E_EXPECTINGTAGEND
	// Manifest Parse Error : The character '>' was expected.
	ERROR_SXS_XML_E_EXPECTINGTAGEND = 14038

	// ERROR_SXS_XML_E_MISSINGSEMICOLON
	// Manifest Parse Error : A semi colon character was expected.
	ERROR_SXS_XML_E_MISSINGSEMICOLON = 14039

	// ERROR_SXS_XML_E_UNBALANCEDPAREN
	// Manifest Parse Error : Unbalanced parentheses.
	ERROR_SXS_XML_E_UNBALANCEDPAREN = 14040

	// ERROR_SXS_XML_E_INTERNALERROR
	// Manifest Parse Error : Internal error.
	ERROR_SXS_XML_E_INTERNALERROR = 14041

	// ERROR_SXS_XML_E_UNEXPECTED_WHITESPACE
	// Manifest Parse Error : Whitespace is not allowed at this location.
	ERROR_SXS_XML_E_UNEXPECTED_WHITESPACE = 14042

	// ERROR_SXS_XML_E_INCOMPLETE_ENCODING
	// Manifest Parse Error : End of file reached in invalid state for current encoding.
	ERROR_SXS_XML_E_INCOMPLETE_ENCODING = 14043

	// ERROR_SXS_XML_E_MISSING_PAREN
	// Manifest Parse Error : Missing parenthesis.
	ERROR_SXS_XML_E_MISSING_PAREN = 14044

	// ERROR_SXS_XML_E_EXPECTINGCLOSEQUOTE
	// Manifest Parse Error : A single or double closing quote character (\' or \") is missing.
	ERROR_SXS_XML_E_EXPECTINGCLOSEQUOTE = 14045

	// ERROR_SXS_XML_E_MULTIPLE_COLONS
	// Manifest Parse Error : Multiple colons are not allowed in a name.
	ERROR_SXS_XML_E_MULTIPLE_COLONS = 14046

	// ERROR_SXS_XML_E_INVALID_DECIMA
	// Manifest Parse Error : Invalid character for decimal digit.
	ERROR_SXS_XML_E_INVALID_DECIMAL = 14047

	// ERROR_SXS_XML_E_INVALID_HEXIDECIMA
	// Manifest Parse Error : Invalid character for hexadecimal digit.
	ERROR_SXS_XML_E_INVALID_HEXIDECIMAL = 14048

	// ERROR_SXS_XML_E_INVALID_UNICODE
	// Manifest Parse Error : Invalid unicode character value for this platform.
	ERROR_SXS_XML_E_INVALID_UNICODE = 14049

	// ERROR_SXS_XML_E_WHITESPACEORQUESTIONMARK
	// Manifest Parse Error : Expecting whitespace or '?'.
	ERROR_SXS_XML_E_WHITESPACEORQUESTIONMARK = 14050

	// ERROR_SXS_XML_E_UNEXPECTEDENDTAG
	// Manifest Parse Error : End tag was not expected at this location.
	ERROR_SXS_XML_E_UNEXPECTEDENDTAG = 14051

	// ERROR_SXS_XML_E_UNCLOSEDTAG
	// Manifest Parse Error : The following tags were not closed: %1.
	ERROR_SXS_XML_E_UNCLOSEDTAG = 14052

	// ERROR_SXS_XML_E_DUPLICATEATTRIBUTE
	// Manifest Parse Error : Duplicate attribute.
	ERROR_SXS_XML_E_DUPLICATEATTRIBUTE = 14053

	// ERROR_SXS_XML_E_MULTIPLEROOTS
	// Manifest Parse Error : Only one top level element is allowed in an XML document.
	ERROR_SXS_XML_E_MULTIPLEROOTS = 14054

	// ERROR_SXS_XML_E_INVALIDATROOTLEVE
	// Manifest Parse Error : Invalid at the top level of the document.
	ERROR_SXS_XML_E_INVALIDATROOTLEVEL = 14055

	// ERROR_SXS_XML_E_BADXMLDEC
	// Manifest Parse Error : Invalid xml declaration.
	ERROR_SXS_XML_E_BADXMLDECL = 14056

	// ERROR_SXS_XML_E_MISSINGROOT
	// Manifest Parse Error : XML document must have a top level element.
	ERROR_SXS_XML_E_MISSINGROOT = 14057

	// ERROR_SXS_XML_E_UNEXPECTEDEOF
	// Manifest Parse Error : Unexpected end of file.
	ERROR_SXS_XML_E_UNEXPECTEDEOF = 14058

	// ERROR_SXS_XML_E_BADPEREFINSUBSET
	// Manifest Parse Error : Parameter entities cannot be used inside markup declarations in an internal subset.
	ERROR_SXS_XML_E_BADPEREFINSUBSET = 14059

	// ERROR_SXS_XML_E_UNCLOSEDSTARTTAG
	// Manifest Parse Error : Element was not closed.
	ERROR_SXS_XML_E_UNCLOSEDSTARTTAG = 14060

	// ERROR_SXS_XML_E_UNCLOSEDENDTAG
	// Manifest Parse Error : End element was missing the character '>'.
	ERROR_SXS_XML_E_UNCLOSEDENDTAG = 14061

	// ERROR_SXS_XML_E_UNCLOSEDSTRING
	// Manifest Parse Error : A string literal was not closed.
	ERROR_SXS_XML_E_UNCLOSEDSTRING = 14062

	// ERROR_SXS_XML_E_UNCLOSEDCOMMENT
	// Manifest Parse Error : A comment was not closed.
	ERROR_SXS_XML_E_UNCLOSEDCOMMENT = 14063

	// ERROR_SXS_XML_E_UNCLOSEDDEC
	// Manifest Parse Error : A declaration was not closed.
	ERROR_SXS_XML_E_UNCLOSEDDECL = 14064

	// ERROR_SXS_XML_E_UNCLOSEDCDATA
	// Manifest Parse Error : A CDATA section was not closed.
	ERROR_SXS_XML_E_UNCLOSEDCDATA = 14065

	// ERROR_SXS_XML_E_RESERVEDNAMESPACE
	// Manifest Parse Error : The namespace prefix is not allowed to start with the reserved string "xml".
	ERROR_SXS_XML_E_RESERVEDNAMESPACE = 14066

	// ERROR_SXS_XML_E_INVALIDENCODING
	// Manifest Parse Error : System does not support the specified encoding.
	ERROR_SXS_XML_E_INVALIDENCODING = 14067

	// ERROR_SXS_XML_E_INVALIDSWITCH
	// Manifest Parse Error : Switch from current encoding to specified encoding not supported.
	ERROR_SXS_XML_E_INVALIDSWITCH = 14068

	// ERROR_SXS_XML_E_BADXMLCASE
	// Manifest Parse Error : The name 'xml' is reserved and must be lower case.
	ERROR_SXS_XML_E_BADXMLCASE = 14069

	// ERROR_SXS_XML_E_INVALID_STANDALONE
	// Manifest Parse Error : The standalone attribute must have the value 'yes' or 'no'.
	ERROR_SXS_XML_E_INVALID_STANDALONE = 14070

	// ERROR_SXS_XML_E_UNEXPECTED_STANDALONE
	// Manifest Parse Error : The standalone attribute cannot be used in external entities.
	ERROR_SXS_XML_E_UNEXPECTED_STANDALONE = 14071

	// ERROR_SXS_XML_E_INVALID_VERSION
	// Manifest Parse Error : Invalid version number.
	ERROR_SXS_XML_E_INVALID_VERSION = 14072

	// ERROR_SXS_XML_E_MISSINGEQUALS
	// Manifest Parse Error : Missing equals sign between attribute and attribute value.
	ERROR_SXS_XML_E_MISSINGEQUALS = 14073

	// ERROR_SXS_PROTECTION_RECOVERY_FAILED
	// Assembly Protection Error : Unable to recover the specified assembly.
	ERROR_SXS_PROTECTION_RECOVERY_FAILED = 14074

	// ERROR_SXS_PROTECTION_PUBLIC_KEY_TOO_SHORT
	// Assembly Protection Error : The public key for an assembly was too short to be allowed.
	ERROR_SXS_PROTECTION_PUBLIC_KEY_TOO_SHORT = 14075

	// ERROR_SXS_PROTECTION_CATALOG_NOT_VALID
	// Assembly Protection Error : The catalog for an assembly is not valid, or does not match the assembly's manifest.
	ERROR_SXS_PROTECTION_CATALOG_NOT_VALID = 14076

	// ERROR_SXS_UNTRANSLATABLE_HRESULT
	// An HRESULT could not be translated to a corresponding Win32 error code.
	ERROR_SXS_UNTRANSLATABLE_HRESULT = 14077

	// ERROR_SXS_PROTECTION_CATALOG_FILE_MISSING
	// Assembly Protection Error : The catalog for an assembly is missing.
	ERROR_SXS_PROTECTION_CATALOG_FILE_MISSING = 14078

	// ERROR_SXS_MISSING_ASSEMBLY_IDENTITY_ATTRIBUTE
	// The supplied assembly identity is missing one or more attributes which must be present in this context.
	ERROR_SXS_MISSING_ASSEMBLY_IDENTITY_ATTRIBUTE = 14079

	// ERROR_SXS_INVALID_ASSEMBLY_IDENTITY_ATTRIBUTE_NAME
	// The supplied assembly identity has one or more attribute names that contain characters not permitted in XML names.
	ERROR_SXS_INVALID_ASSEMBLY_IDENTITY_ATTRIBUTE_NAME = 14080

	// ERROR_SXS_ASSEMBLY_MISSING
	// The referenced assembly could not be found.
	ERROR_SXS_ASSEMBLY_MISSING = 14081

	// ERROR_SXS_CORRUPT_ACTIVATION_STACK
	// The activation context activation stack for the running thread of execution is corrupt.
	ERROR_SXS_CORRUPT_ACTIVATION_STACK = 14082

	// ERROR_SXS_CORRUPTION
	// The application isolation metadata for this process or thread has become corrupt.
	ERROR_SXS_CORRUPTION = 14083

	// ERROR_SXS_EARLY_DEACTIVATION
	// The activation context being deactivated is not the most recently activated one.
	ERROR_SXS_EARLY_DEACTIVATION = 14084

	// ERROR_SXS_INVALID_DEACTIVATION
	// The activation context being deactivated is not active for the current thread of execution.
	ERROR_SXS_INVALID_DEACTIVATION = 14085

	// ERROR_SXS_MULTIPLE_DEACTIVATION
	// The activation context being deactivated has already been deactivated.
	ERROR_SXS_MULTIPLE_DEACTIVATION = 14086

	// ERROR_SXS_PROCESS_TERMINATION_REQUESTED
	// A component used by the isolation facility has requested to terminate the process.
	ERROR_SXS_PROCESS_TERMINATION_REQUESTED = 14087

	// ERROR_SXS_RELEASE_ACTIVATION_CONTEXT
	// A kernel mode component is releasing a reference on an activation context.
	ERROR_SXS_RELEASE_ACTIVATION_CONTEXT = 14088

	// ERROR_SXS_SYSTEM_DEFAULT_ACTIVATION_CONTEXT_EMPTY
	// The activation context of system default assembly could not be generated.
	ERROR_SXS_SYSTEM_DEFAULT_ACTIVATION_CONTEXT_EMPTY = 14089

	// ERROR_SXS_INVALID_IDENTITY_ATTRIBUTE_VALUE
	// The value of an attribute in an identity is not within the legal range.
	ERROR_SXS_INVALID_IDENTITY_ATTRIBUTE_VALUE = 14090

	// ERROR_SXS_INVALID_IDENTITY_ATTRIBUTE_NAME
	// The name of an attribute in an identity is not within the legal range.
	ERROR_SXS_INVALID_IDENTITY_ATTRIBUTE_NAME = 14091

	// ERROR_SXS_IDENTITY_DUPLICATE_ATTRIBUTE
	// An identity contains two definitions for the same attribute.
	ERROR_SXS_IDENTITY_DUPLICATE_ATTRIBUTE = 14092

	// ERROR_SXS_IDENTITY_PARSE_ERROR
	// The identity string is malformed. This may be due to a trailing comma, more than two unnamed attributes, missing attribute name or missing attribute value.
	ERROR_SXS_IDENTITY_PARSE_ERROR = 14093

	// ERROR_MALFORMED_SUBSTITUTION_STRING
	// A string containing localized substitutable content was malformed. Either a dollar sign ($) was followed by something other than a left parenthesis or another dollar sign or an substitution's right parenthesis was not found.
	ERROR_MALFORMED_SUBSTITUTION_STRING = 14094

	// ERROR_SXS_INCORRECT_PUBLIC_KEY_TOKEN
	// The public key token does not correspond to the public key specified.
	ERROR_SXS_INCORRECT_PUBLIC_KEY_TOKEN = 14095

	// ERROR_UNMAPPED_SUBSTITUTION_STRING
	// A substitution string had no mapping.
	ERROR_UNMAPPED_SUBSTITUTION_STRING = 14096

	// ERROR_SXS_ASSEMBLY_NOT_LOCKED
	// The component must be locked before making the request.
	ERROR_SXS_ASSEMBLY_NOT_LOCKED = 14097

	// ERROR_SXS_COMPONENT_STORE_CORRUPT
	// The component store has been corrupted.
	ERROR_SXS_COMPONENT_STORE_CORRUPT = 14098

	// ERROR_ADVANCED_INSTALLER_FAILED
	// An advanced installer failed during setup or servicing.
	ERROR_ADVANCED_INSTALLER_FAILED = 14099

	// ERROR_XML_ENCODING_MISMATCH
	// The character encoding in the XML declaration did not match the encoding used in the document.
	ERROR_XML_ENCODING_MISMATCH = 14100

	// ERROR_SXS_MANIFEST_IDENTITY_SAME_BUT_CONTENTS_DIFFERENT
	// The identities of the manifests are identical but their contents are different.
	ERROR_SXS_MANIFEST_IDENTITY_SAME_BUT_CONTENTS_DIFFERENT = 14101

	// ERROR_SXS_IDENTITIES_DIFFERENT
	// The component identities are different.
	ERROR_SXS_IDENTITIES_DIFFERENT = 14102

	// ERROR_SXS_ASSEMBLY_IS_NOT_A_DEPLOYMENT
	// The assembly is not a deployment.
	ERROR_SXS_ASSEMBLY_IS_NOT_A_DEPLOYMENT = 14103

	// ERROR_SXS_FILE_NOT_PART_OF_ASSEMBLY
	// The file is not a part of the assembly.
	ERROR_SXS_FILE_NOT_PART_OF_ASSEMBLY = 14104

	// ERROR_SXS_MANIFEST_TOO_BIG
	// The size of the manifest exceeds the maximum allowed.
	ERROR_SXS_MANIFEST_TOO_BIG = 14105

	// ERROR_SXS_SETTING_NOT_REGISTERED
	// The setting is not registered.
	ERROR_SXS_SETTING_NOT_REGISTERED = 14106

	// ERROR_SXS_TRANSACTION_CLOSURE_INCOMPLETE
	// One or more required members of the transaction are not present.
	ERROR_SXS_TRANSACTION_CLOSURE_INCOMPLETE = 14107

	// ERROR_SMI_PRIMITIVE_INSTALLER_FAILED
	// The SMI primitive installer failed during setup or servicing.
	ERROR_SMI_PRIMITIVE_INSTALLER_FAILED = 14108

	// ERROR_GENERIC_COMMAND_FAILED
	// A generic command executable returned a result that indicates failure.
	ERROR_GENERIC_COMMAND_FAILED = 14109

	// ERROR_SXS_FILE_HASH_MISSING
	// A component is missing file verification information in its manifest.
	ERROR_SXS_FILE_HASH_MISSING = 14110

	///////////////////////////////////////////////////
	//                                               //
	//           End of Side By Side Error Codes     //
	//                                               //
	//                =14000 to=14999                //
	///////////////////////////////////////////////////

	///////////////////////////////////////////////////
	//                                               //
	//           Start of WinEvt Error codes         //
	//                                               //
	//                =15000 to=15079                //
	///////////////////////////////////////////////////

	// ERROR_EVT_INVALID_CHANNEL_PATH
	// The specified channel path is invalid.
	ERROR_EVT_INVALID_CHANNEL_PATH = 15000

	// ERROR_EVT_INVALID_QUERY
	// The specified query is invalid.
	ERROR_EVT_INVALID_QUERY = 15001

	// ERROR_EVT_PUBLISHER_METADATA_NOT_FOUND
	// The publisher metadata cannot be found in the resource.
	ERROR_EVT_PUBLISHER_METADATA_NOT_FOUND = 15002

	// ERROR_EVT_EVENT_TEMPLATE_NOT_FOUND
	// The template for an event definition cannot be found in the resource (error = %1).
	ERROR_EVT_EVENT_TEMPLATE_NOT_FOUND = 15003

	// ERROR_EVT_INVALID_PUBLISHER_NAME
	// The specified publisher name is invalid.
	ERROR_EVT_INVALID_PUBLISHER_NAME = 15004

	// ERROR_EVT_INVALID_EVENT_DATA
	// The event data raised by the publisher is not compatible with the event template definition in the publisher's manifest.
	ERROR_EVT_INVALID_EVENT_DATA = 15005

	// ERROR_EVT_CHANNEL_NOT_FOUND
	// The specified channel could not be found.
	ERROR_EVT_CHANNEL_NOT_FOUND = 15007

	// ERROR_EVT_MALFORMED_XML_TEXT
	// The specified XML text was not well-formed. See Extended Error for more details.
	ERROR_EVT_MALFORMED_XML_TEXT = 15008

	// ERROR_EVT_SUBSCRIPTION_TO_DIRECT_CHANNE
	// The events for a direct channel go directly to a log file and cannot be subscribed to.
	ERROR_EVT_SUBSCRIPTION_TO_DIRECT_CHANNEL = 15009

	// ERROR_EVT_CONFIGURATION_ERROR
	// Configuration error.
	ERROR_EVT_CONFIGURATION_ERROR = 15010

	// ERROR_EVT_QUERY_RESULT_STALE
	// The query result is stale or invalid and must be recreated. This may be due to the log being cleared or rolling over after the query result was created.
	ERROR_EVT_QUERY_RESULT_STALE = 15011

	// ERROR_EVT_QUERY_RESULT_INVALID_POSITION
	// The query result is currently at an invalid position.
	ERROR_EVT_QUERY_RESULT_INVALID_POSITION = 15012

	// ERROR_EVT_NON_VALIDATING_MSXM
	// Registered MSXML doesn't support validation.
	ERROR_EVT_NON_VALIDATING_MSXML = 15013

	// ERROR_EVT_FILTER_ALREADYSCOPED
	// An expression can only be followed by a change-of-scope operation if the expression evaluates to a node set and is not already part of another change-of-scope operation.
	ERROR_EVT_FILTER_ALREADYSCOPED = 15014

	// ERROR_EVT_FILTER_NOTELTSET
	// Cannot perform a step operation from a term that does not represent an element set.
	ERROR_EVT_FILTER_NOTELTSET = 15015

	// ERROR_EVT_FILTER_INVARG
	// Left-hand side arguments to binary operators must be either attributes, nodes or variables. Right-hand side arguments must be constants.
	ERROR_EVT_FILTER_INVARG = 15016

	// ERROR_EVT_FILTER_INVTEST
	// A step operation must involve a node test or, in the case of a predicate, an algebraic expression against which to test each node in the preceeding node set.
	ERROR_EVT_FILTER_INVTEST = 15017

	// ERROR_EVT_FILTER_INVTYPE
	// This data type is currently unsupported.
	ERROR_EVT_FILTER_INVTYPE = 15018

	// ERROR_EVT_FILTER_PARSEERR
	// A syntax error occurred at position %1!d!
	ERROR_EVT_FILTER_PARSEERR = 15019

	// ERROR_EVT_FILTER_UNSUPPORTEDOP
	// This operator is unsupported by this implementation of the filter.
	ERROR_EVT_FILTER_UNSUPPORTEDOP = 15020

	// ERROR_EVT_FILTER_UNEXPECTEDTOKEN
	// An unexpected token was encountered.
	ERROR_EVT_FILTER_UNEXPECTEDTOKEN = 15021

	// ERROR_EVT_INVALID_OPERATION_OVER_ENABLED_DIRECT_CHANNE
	// The requested operation cannot be performed over an enabled direct channel. The channel must first be disabled.
	ERROR_EVT_INVALID_OPERATION_OVER_ENABLED_DIRECT_CHANNEL = 15022

	// ERROR_EVT_INVALID_CHANNEL_PROPERTY_VALUE
	// Channel property %1!s! contains an invalid value. The value has an invalid type, is outside of its valid range, cannot be changed, or is not supported by this type of channel.
	ERROR_EVT_INVALID_CHANNEL_PROPERTY_VALUE = 15023

	// ERROR_EVT_INVALID_PUBLISHER_PROPERTY_VALUE
	// Publisher property %1!s! contains an invalid value. The value has an invalid type, is outside of its valid range, cannot be changed, or is not supported by this type of publisher.
	ERROR_EVT_INVALID_PUBLISHER_PROPERTY_VALUE = 15024

	// ERROR_EVT_CHANNEL_CANNOT_ACTIVATE
	// The channel failed to activate.
	ERROR_EVT_CHANNEL_CANNOT_ACTIVATE = 15025

	// ERROR_EVT_FILTER_TOO_COMPLEX
	// The XPath expression exceeded the supported complexity. Simplify the expression or split it into multiple expressions.
	ERROR_EVT_FILTER_TOO_COMPLEX = 15026

	// ERROR_EVT_MESSAGE_NOT_FOUND
	// The message resource is present but the message was not found in the message table.
	ERROR_EVT_MESSAGE_NOT_FOUND = 15027

	// ERROR_EVT_MESSAGE_ID_NOT_FOUND
	// The message ID for the desired message could not be found.
	ERROR_EVT_MESSAGE_ID_NOT_FOUND = 15028

	// ERROR_EVT_UNRESOLVED_VALUE_INSERT
	// The substitution string for insert index (%1) could not be found.
	ERROR_EVT_UNRESOLVED_VALUE_INSERT = 15029

	// ERROR_EVT_UNRESOLVED_PARAMETER_INSERT
	// The description string for parameter reference (%1) could not be found.
	ERROR_EVT_UNRESOLVED_PARAMETER_INSERT = 15030

	// ERROR_EVT_MAX_INSERTS_REACHED
	// The maximum number of replacements has been reached.
	ERROR_EVT_MAX_INSERTS_REACHED = 15031

	// ERROR_EVT_EVENT_DEFINITION_NOT_FOUND
	// The event definition could not be found for event ID (%1).
	ERROR_EVT_EVENT_DEFINITION_NOT_FOUND = 15032

	// ERROR_EVT_MESSAGE_LOCALE_NOT_FOUND
	// The locale specific resource for the desired message is not present.
	ERROR_EVT_MESSAGE_LOCALE_NOT_FOUND = 15033

	// ERROR_EVT_VERSION_TOO_OLD
	// The resource is too old and is not supported.
	ERROR_EVT_VERSION_TOO_OLD = 15034

	// ERROR_EVT_VERSION_TOO_NEW
	// The resource is too new and is not supported.
	ERROR_EVT_VERSION_TOO_NEW = 15035

	// ERROR_EVT_CANNOT_OPEN_CHANNEL_OF_QUERY
	// The channel at index %1!d! of the query can't be opened.
	ERROR_EVT_CANNOT_OPEN_CHANNEL_OF_QUERY = 15036

	// ERROR_EVT_PUBLISHER_DISABLED
	// The publisher has been disabled and its resource is not available. This usually occurs when the publisher is in the process of being uninstalled or upgraded.
	ERROR_EVT_PUBLISHER_DISABLED = 15037

	// ERROR_EVT_FILTER_OUT_OF_RANGE
	// Attempted to create a numeric type that is outside of its valid range.
	ERROR_EVT_FILTER_OUT_OF_RANGE = 15038

	///////////////////////////////////////////////////
	//                                               //
	//           Start of Wecsvc Error codes         //
	//                                               //
	//                =15080 to=15099                //
	///////////////////////////////////////////////////

	// ERROR_EC_SUBSCRIPTION_CANNOT_ACTIVATE
	// The subscription fails to activate.
	ERROR_EC_SUBSCRIPTION_CANNOT_ACTIVATE = 15080

	// ERROR_EC_LOG_DISABLED
	// The log of the subscription is in disabled state, and can not be used to forward events to. The log must first be enabled before the subscription can be activated.
	ERROR_EC_LOG_DISABLED = 15081

	// ERROR_EC_CIRCULAR_FORWARDING
	// When forwarding events from local machine to itself, the query of the subscription can't contain target log of the subscription.
	ERROR_EC_CIRCULAR_FORWARDING = 15082

	// ERROR_EC_CREDSTORE_FUL
	// The credential store that is used to save credentials is full.
	ERROR_EC_CREDSTORE_FULL = 15083

	// ERROR_EC_CRED_NOT_FOUND
	// The credential used by this subscription can't be found in credential store.
	ERROR_EC_CRED_NOT_FOUND = 15084

	// ERROR_EC_NO_ACTIVE_CHANNE
	// No active channel is found for the query.
	ERROR_EC_NO_ACTIVE_CHANNEL = 15085

	///////////////////////////////////////////////////
	//                                               //
	//           Start of MUI Error codes            //
	//                                               //
	//                =15100 to=15199                //
	///////////////////////////////////////////////////

	// ERROR_MUI_FILE_NOT_FOUND
	// The resource loader failed to find MUI file.
	ERROR_MUI_FILE_NOT_FOUND = 15100

	// ERROR_MUI_INVALID_FILE
	// The resource loader failed to load MUI file because the file fail to pass validation.
	ERROR_MUI_INVALID_FILE = 15101

	// ERROR_MUI_INVALID_RC_CONFIG
	// The RC Manifest is corrupted with garbage data or unsupported version or missing required item.
	ERROR_MUI_INVALID_RC_CONFIG = 15102

	// ERROR_MUI_INVALID_LOCALE_NAME
	// The RC Manifest has invalid culture name.
	ERROR_MUI_INVALID_LOCALE_NAME = 15103

	// ERROR_MUI_INVALID_ULTIMATEFALLBACK_NAME
	// The RC Manifest has invalid ultimatefallback name.
	ERROR_MUI_INVALID_ULTIMATEFALLBACK_NAME = 15104

	// ERROR_MUI_FILE_NOT_LOADED
	// The resource loader cache doesn't have loaded MUI entry.
	ERROR_MUI_FILE_NOT_LOADED = 15105

	// ERROR_RESOURCE_ENUM_USER_STOP
	// User stopped resource enumeration.
	ERROR_RESOURCE_ENUM_USER_STOP = 15106

	// ERROR_MUI_INTLSETTINGS_UILANG_NOT_INSTALLED
	// UI language installation failed.
	ERROR_MUI_INTLSETTINGS_UILANG_NOT_INSTALLED = 15107

	// ERROR_MUI_INTLSETTINGS_INVALID_LOCALE_NAME
	// Locale installation failed.
	ERROR_MUI_INTLSETTINGS_INVALID_LOCALE_NAME = 15108

	// ERROR_MRM_RUNTIME_NO_DEFAULT_OR_NEUTRAL_RESOURCE
	// A resource does not have default or neutral value.
	ERROR_MRM_RUNTIME_NO_DEFAULT_OR_NEUTRAL_RESOURCE = 15110

	// ERROR_MRM_INVALID_PRICONFIG
	// Invalid PRI config file.
	ERROR_MRM_INVALID_PRICONFIG = 15111

	// ERROR_MRM_INVALID_FILE_TYPE
	// Invalid file type.
	ERROR_MRM_INVALID_FILE_TYPE = 15112

	// ERROR_MRM_UNKNOWN_QUALIFIER
	// Unknown qualifier.
	ERROR_MRM_UNKNOWN_QUALIFIER = 15113

	// ERROR_MRM_INVALID_QUALIFIER_VALUE
	// Invalid qualifier value.
	ERROR_MRM_INVALID_QUALIFIER_VALUE = 15114

	// ERROR_MRM_NO_CANDIDATE
	// No Candidate found.
	ERROR_MRM_NO_CANDIDATE = 15115

	// ERROR_MRM_NO_MATCH_OR_DEFAULT_CANDIDATE
	// The ResourceMap or NamedResource has an item that does not have default or neutral resource..
	ERROR_MRM_NO_MATCH_OR_DEFAULT_CANDIDATE = 15116

	// ERROR_MRM_RESOURCE_TYPE_MISMATCH
	// Invalid ResourceCandidate type.
	ERROR_MRM_RESOURCE_TYPE_MISMATCH = 15117

	// ERROR_MRM_DUPLICATE_MAP_NAME
	// Duplicate Resource Map.
	ERROR_MRM_DUPLICATE_MAP_NAME = 15118

	// ERROR_MRM_DUPLICATE_ENTRY
	// Duplicate Entry.
	ERROR_MRM_DUPLICATE_ENTRY = 15119

	// ERROR_MRM_INVALID_RESOURCE_IDENTIFIER
	// Invalid Resource Identifier.
	ERROR_MRM_INVALID_RESOURCE_IDENTIFIER = 15120

	// ERROR_MRM_FILEPATH_TOO_LONG
	// Filepath too long.
	ERROR_MRM_FILEPATH_TOO_LONG = 15121

	// ERROR_MRM_UNSUPPORTED_DIRECTORY_TYPE
	// Unsupported directory type.
	ERROR_MRM_UNSUPPORTED_DIRECTORY_TYPE = 15122

	// ERROR_MRM_INVALID_PRI_FILE
	// Invalid PRI File.
	ERROR_MRM_INVALID_PRI_FILE = 15126

	// ERROR_MRM_NAMED_RESOURCE_NOT_FOUND
	// NamedResource Not Found.
	ERROR_MRM_NAMED_RESOURCE_NOT_FOUND = 15127

	// ERROR_MRM_MAP_NOT_FOUND
	// ResourceMap Not Found.
	ERROR_MRM_MAP_NOT_FOUND = 15135

	// ERROR_MRM_UNSUPPORTED_PROFILE_TYPE
	// Unsupported MRT profile type.
	ERROR_MRM_UNSUPPORTED_PROFILE_TYPE = 15136

	// ERROR_MRM_INVALID_QUALIFIER_OPERATOR
	// Invalid qualifier operator.
	ERROR_MRM_INVALID_QUALIFIER_OPERATOR = 15137

	// ERROR_MRM_INDETERMINATE_QUALIFIER_VALUE
	// Unable to determine qualifier value or qualifier value has not been set.
	ERROR_MRM_INDETERMINATE_QUALIFIER_VALUE = 15138

	// ERROR_MRM_AUTOMERGE_ENABLED
	// Automerge is enabled in the PRI file.
	ERROR_MRM_AUTOMERGE_ENABLED = 15139

	// ERROR_MRM_TOO_MANY_RESOURCES
	// Too many resources defined for package.
	ERROR_MRM_TOO_MANY_RESOURCES = 15140

	// ERROR_MRM_UNSUPPORTED_FILE_TYPE_FOR_MERGE
	// Resource File can not be used for merge operation.
	ERROR_MRM_UNSUPPORTED_FILE_TYPE_FOR_MERGE = 15141

	// ERROR_MRM_UNSUPPORTED_FILE_TYPE_FOR_LOAD_UNLOAD_PRI_FILE
	// Load/UnloadPriFiles cannot be used with resource packages.
	ERROR_MRM_UNSUPPORTED_FILE_TYPE_FOR_LOAD_UNLOAD_PRI_FILE = 15142

	// ERROR_MRM_NO_CURRENT_VIEW_ON_THREAD
	// Resource Contexts may not be created on threads that do not have a CoreWindow.
	ERROR_MRM_NO_CURRENT_VIEW_ON_THREAD = 15143

	// ERROR_DIFFERENT_PROFILE_RESOURCE_MANAGER_EXIST
	// The singleton Resource Manager with different profile is already created.
	ERROR_DIFFERENT_PROFILE_RESOURCE_MANAGER_EXIST = 15144

	// ERROR_OPERATION_NOT_ALLOWED_FROM_SYSTEM_COMPONENT
	// The system component cannot operate given API operation
	ERROR_OPERATION_NOT_ALLOWED_FROM_SYSTEM_COMPONENT = 15145

	// ERROR_MRM_DIRECT_REF_TO_NON_DEFAULT_RESOURCE
	// The resource is a direct reference to a non-default resource candidate.
	ERROR_MRM_DIRECT_REF_TO_NON_DEFAULT_RESOURCE = 15146

	// ERROR_MRM_GENERATION_COUNT_MISMATCH
	// Resource Map has been re-generated and the query string is not valid anymore.
	ERROR_MRM_GENERATION_COUNT_MISMATCH = 15147

	// ERROR_PRI_MERGE_VERSION_MISMATCH
	// The PRI files to be merged have incompatible versions.
	ERROR_PRI_MERGE_VERSION_MISMATCH = 15148

	// ERROR_PRI_MERGE_MISSING_SCHEMA
	// The primary PRI files to be merged does not contain a schema.
	ERROR_PRI_MERGE_MISSING_SCHEMA = 15149

	// ERROR_PRI_MERGE_LOAD_FILE_FAILED
	// Unable to load one of the PRI files to be merged.
	ERROR_PRI_MERGE_LOAD_FILE_FAILED = 15150

	// ERROR_PRI_MERGE_ADD_FILE_FAILED
	// Unable to add one of the PRI files to the merged file.
	ERROR_PRI_MERGE_ADD_FILE_FAILED = 15151

	// ERROR_PRI_MERGE_WRITE_FILE_FAILED
	// Unable to create the merged PRI file.
	ERROR_PRI_MERGE_WRITE_FILE_FAILED = 15152

	// ERROR_PRI_MERGE_MULTIPLE_PACKAGE_FAMILIES_NOT_ALLOWED
	// Packages for a PRI file merge must all be from the same package family.
	ERROR_PRI_MERGE_MULTIPLE_PACKAGE_FAMILIES_NOT_ALLOWED = 15153

	// ERROR_PRI_MERGE_MULTIPLE_MAIN_PACKAGES_NOT_ALLOWED
	// Packages for a PRI file merge must not include multiple main packages.
	ERROR_PRI_MERGE_MULTIPLE_MAIN_PACKAGES_NOT_ALLOWED = 15154

	// ERROR_PRI_MERGE_BUNDLE_PACKAGES_NOT_ALLOWED
	// Packages for a PRI file merge must not include bundle packages.
	ERROR_PRI_MERGE_BUNDLE_PACKAGES_NOT_ALLOWED = 15155

	// ERROR_PRI_MERGE_MAIN_PACKAGE_REQUIRED
	// Packages for a PRI file merge must include one main package.
	ERROR_PRI_MERGE_MAIN_PACKAGE_REQUIRED = 15156

	// ERROR_PRI_MERGE_RESOURCE_PACKAGE_REQUIRED
	// Packages for a PRI file merge must include at least one resource package.
	ERROR_PRI_MERGE_RESOURCE_PACKAGE_REQUIRED = 15157

	// ERROR_PRI_MERGE_INVALID_FILE_NAME
	// Invalid name supplied for a canonical merged PRI file.
	ERROR_PRI_MERGE_INVALID_FILE_NAME = 15158

	// ERROR_MRM_PACKAGE_NOT_FOUND
	// Unable to find the specified package.
	ERROR_MRM_PACKAGE_NOT_FOUND = 15159

	///////////////////////////////////////////////////
	//                                               //
	// Start of Monitor Configuration API error codes//
	//                                               //
	//                =15200 to=15249                //
	///////////////////////////////////////////////////

	// ERROR_MCA_INVALID_CAPABILITIES_STRING
	// The monitor returned a DDC/CI capabilities string that did not comply with the ACCESS.bus=3.0, DDC/CI=1.1 or MCCS=2 Revision=1 specification.
	ERROR_MCA_INVALID_CAPABILITIES_STRING = 15200

	// ERROR_MCA_INVALID_VCP_VERSION
	// The monitor's VCP Version (0xDF) VCP code returned an invalid version value.
	ERROR_MCA_INVALID_VCP_VERSION = 15201

	// ERROR_MCA_MONITOR_VIOLATES_MCCS_SPECIFICATION
	// The monitor does not comply with the MCCS specification it claims to support.
	ERROR_MCA_MONITOR_VIOLATES_MCCS_SPECIFICATION = 15202

	// ERROR_MCA_MCCS_VERSION_MISMATCH
	// The MCCS version in a monitor's mccs_ver capability does not match the MCCS version the monitor reports when the VCP Version (0xDF) VCP code is used.
	ERROR_MCA_MCCS_VERSION_MISMATCH = 15203

	// ERROR_MCA_UNSUPPORTED_MCCS_VERSION
	// The Monitor Configuration API only works with monitors that support the MCCS=1.0 specification, MCCS=2.0 specification or the MCCS=2.0 Revision=1 specification.
	ERROR_MCA_UNSUPPORTED_MCCS_VERSION = 15204

	// ERROR_MCA_INTERNAL_ERROR
	// An internal Monitor Configuration API error occurred.
	ERROR_MCA_INTERNAL_ERROR = 15205

	// ERROR_MCA_INVALID_TECHNOLOGY_TYPE_RETURNED
	// The monitor returned an invalid monitor technology type. CRT, Plasma and LCD (TFT) are examples of monitor technology types. This error implies that the monitor violated the MCCS=2.0 or MCCS=2.0 Revision=1 specification.
	ERROR_MCA_INVALID_TECHNOLOGY_TYPE_RETURNED = 15206

	// ERROR_MCA_UNSUPPORTED_COLOR_TEMPERATURE
	// The caller of SetMonitorColorTemperature specified a color temperature that the current monitor did not support. This error implies that the monitor violated the MCCS=2.0 or MCCS=2.0 Revision=1 specification.
	ERROR_MCA_UNSUPPORTED_COLOR_TEMPERATURE = 15207

	//////////////////////////////////////////////////
	//                                              //
	// End of Monitor Configuration API error codes //
	//                                              //
	//               =15200 to=15249                //
	//                                              //
	//////////////////////////////////////////////////
	//////////////////////////////////////////////////
	//                                              //
	//         Start of Syspart error codes         //
	//               =15250 -=15299                 //
	//                                              //
	//////////////////////////////////////////////////

	// ERROR_AMBIGUOUS_SYSTEM_DEVICE
	// The requested system device cannot be identified due to multiple indistinguishable devices potentially matching the identification criteria.
	ERROR_AMBIGUOUS_SYSTEM_DEVICE = 15250

	// ERROR_SYSTEM_DEVICE_NOT_FOUND
	// The requested system device cannot be found.
	ERROR_SYSTEM_DEVICE_NOT_FOUND = 15299

	//////////////////////////////////////////////////
	//                                              //
	//         Start of Vortex error codes          //
	//               =15300 -=15320                 //
	//                                              //
	//////////////////////////////////////////////////

	// ERROR_HASH_NOT_SUPPORTED
	// Hash generation for the specified hash version and hash type is not enabled on the server.
	ERROR_HASH_NOT_SUPPORTED = 15300

	// ERROR_HASH_NOT_PRESENT
	// The hash requested from the server is not available or no longer valid.
	ERROR_HASH_NOT_PRESENT = 15301

	//////////////////////////////////////////////////
	//                                              //
	//         Start of GPIO error codes            //
	//               =15321 -=15340                 //
	//                                              //
	//////////////////////////////////////////////////

	// ERROR_SECONDARY_IC_PROVIDER_NOT_REGISTERED
	// The secondary interrupt controller instance that manages the specified interrupt is not registered.
	ERROR_SECONDARY_IC_PROVIDER_NOT_REGISTERED = 15321

	// ERROR_GPIO_CLIENT_INFORMATION_INVALID
	// The information supplied by the GPIO client driver is invalid.
	ERROR_GPIO_CLIENT_INFORMATION_INVALID = 15322

	// ERROR_GPIO_VERSION_NOT_SUPPORTED
	// The version specified by the GPIO client driver is not supported.
	ERROR_GPIO_VERSION_NOT_SUPPORTED = 15323

	// ERROR_GPIO_INVALID_REGISTRATION_PACKET
	// The registration packet supplied by the GPIO client driver is not valid.
	ERROR_GPIO_INVALID_REGISTRATION_PACKET = 15324

	// ERROR_GPIO_OPERATION_DENIED
	// The requested operation is not supported for the specified handle.
	ERROR_GPIO_OPERATION_DENIED = 15325

	// ERROR_GPIO_INCOMPATIBLE_CONNECT_MODE
	// The requested connect mode conflicts with an existing mode on one or more of the specified pins.
	ERROR_GPIO_INCOMPATIBLE_CONNECT_MODE = 15326

	// ERROR_GPIO_INTERRUPT_ALREADY_UNMASKED
	// The interrupt requested to be unmasked is not masked.
	ERROR_GPIO_INTERRUPT_ALREADY_UNMASKED = 15327

	//////////////////////////////////////////////////
	//                                              //
	//         Start of Run Level error codes       //
	//               =15400 -=15500                 //
	//                                              //
	//////////////////////////////////////////////////

	// ERROR_CANNOT_SWITCH_RUNLEVE
	// The requested run level switch cannot be completed successfully.
	ERROR_CANNOT_SWITCH_RUNLEVEL = 15400

	// ERROR_INVALID_RUNLEVEL_SETTING
	// The service has an invalid run level setting. The run level for a service
	// must not be higher than the run level of its dependent services.
	ERROR_INVALID_RUNLEVEL_SETTING = 15401

	// ERROR_RUNLEVEL_SWITCH_TIMEOUT
	// The requested run level switch cannot be completed successfully since
	// one or more services will not stop or restart within the specified timeout.
	ERROR_RUNLEVEL_SWITCH_TIMEOUT = 15402

	// ERROR_RUNLEVEL_SWITCH_AGENT_TIMEOUT
	// A run level switch agent did not respond within the specified timeout.
	ERROR_RUNLEVEL_SWITCH_AGENT_TIMEOUT = 15403

	// ERROR_RUNLEVEL_SWITCH_IN_PROGRESS
	// A run level switch is currently in progress.
	ERROR_RUNLEVEL_SWITCH_IN_PROGRESS = 15404

	// ERROR_SERVICES_FAILED_AUTOSTART
	// One or more services failed to start during the service startup phase of a run level switch.
	ERROR_SERVICES_FAILED_AUTOSTART = 15405

	//////////////////////////////////////////////////
	//                                              //
	//         Start of Com Task error codes        //
	//               =15501 -=15510                 //
	//                                              //
	//////////////////////////////////////////////////

	// ERROR_COM_TASK_STOP_PENDING
	// The task stop request cannot be completed immediately since
	// task needs more time to shutdown.
	ERROR_COM_TASK_STOP_PENDING = 15501

	////////////////////////////////////////
	//                                    //
	// APPX Caller Visible Error Codes    //
	//         =15600-15699               //
	////////////////////////////////////////

	// ERROR_INSTALL_OPEN_PACKAGE_FAILED
	// Package could not be opened.
	ERROR_INSTALL_OPEN_PACKAGE_FAILED = 15600

	// ERROR_INSTALL_PACKAGE_NOT_FOUND
	// Package was not found.
	ERROR_INSTALL_PACKAGE_NOT_FOUND = 15601

	// ERROR_INSTALL_INVALID_PACKAGE
	// Package data is invalid.
	ERROR_INSTALL_INVALID_PACKAGE = 15602

	// ERROR_INSTALL_RESOLVE_DEPENDENCY_FAILED
	// Package failed updates, dependency or conflict validation.
	ERROR_INSTALL_RESOLVE_DEPENDENCY_FAILED = 15603

	// ERROR_INSTALL_OUT_OF_DISK_SPACE
	// There is not enough disk space on your computer. Please free up some space and try again.
	ERROR_INSTALL_OUT_OF_DISK_SPACE = 15604

	// ERROR_INSTALL_NETWORK_FAILURE
	// There was a problem downloading your product.
	ERROR_INSTALL_NETWORK_FAILURE = 15605

	// ERROR_INSTALL_REGISTRATION_FAILURE
	// Package could not be registered.
	ERROR_INSTALL_REGISTRATION_FAILURE = 15606

	// ERROR_INSTALL_DEREGISTRATION_FAILURE
	// Package could not be unregistered.
	ERROR_INSTALL_DEREGISTRATION_FAILURE = 15607

	// ERROR_INSTALL_CANCE
	// User cancelled the install request.
	ERROR_INSTALL_CANCEL = 15608

	// ERROR_INSTALL_FAILED
	// Install failed. Please contact your software vendor.
	ERROR_INSTALL_FAILED = 15609

	// ERROR_REMOVE_FAILED
	// Removal failed. Please contact your software vendor.
	ERROR_REMOVE_FAILED = 15610

	// ERROR_PACKAGE_ALREADY_EXISTS
	// The provided package is already installed, and reinstallation of the package was blocked. Check the AppXDeployment-Server event log for details.
	ERROR_PACKAGE_ALREADY_EXISTS = 15611

	// ERROR_NEEDS_REMEDIATION
	// The application cannot be started. Try reinstalling the application to fix the problem.
	ERROR_NEEDS_REMEDIATION = 15612

	// ERROR_INSTALL_PREREQUISITE_FAILED
	// A Prerequisite for an install could not be satisfied.
	ERROR_INSTALL_PREREQUISITE_FAILED = 15613

	// ERROR_PACKAGE_REPOSITORY_CORRUPTED
	// The package repository is corrupted.
	ERROR_PACKAGE_REPOSITORY_CORRUPTED = 15614

	// ERROR_INSTALL_POLICY_FAILURE
	// To install this application you need either a Windows developer license or a sideloading-enabled system.
	ERROR_INSTALL_POLICY_FAILURE = 15615

	// ERROR_PACKAGE_UPDATING
	// The application cannot be started because it is currently updating.
	ERROR_PACKAGE_UPDATING = 15616

	// ERROR_DEPLOYMENT_BLOCKED_BY_POLICY
	// The package deployment operation is blocked by policy. Please contact your system administrator.
	ERROR_DEPLOYMENT_BLOCKED_BY_POLICY = 15617

	// ERROR_PACKAGES_IN_USE
	// The package could not be installed because resources it modifies are currently in use.
	ERROR_PACKAGES_IN_USE = 15618

	// ERROR_RECOVERY_FILE_CORRUPT
	// The package could not be recovered because necessary data for recovery have been corrupted.
	ERROR_RECOVERY_FILE_CORRUPT = 15619

	// ERROR_INVALID_STAGED_SIGNATURE
	// The signature is invalid. To register in developer mode, AppxSignature.p7x and AppxBlockMap.xml must be valid or should not be present.
	ERROR_INVALID_STAGED_SIGNATURE = 15620

	// ERROR_DELETING_EXISTING_APPLICATIONDATA_STORE_FAILED
	// An error occurred while deleting the package's previously existing application data.
	ERROR_DELETING_EXISTING_APPLICATIONDATA_STORE_FAILED = 15621

	// ERROR_INSTALL_PACKAGE_DOWNGRADE
	// The package could not be installed because a higher version of this package is already installed.
	ERROR_INSTALL_PACKAGE_DOWNGRADE = 15622

	// ERROR_SYSTEM_NEEDS_REMEDIATION
	// An error in a system binary was detected. Try refreshing the PC to fix the problem.
	ERROR_SYSTEM_NEEDS_REMEDIATION = 15623

	// ERROR_APPX_INTEGRITY_FAILURE_CLR_NGEN
	// A corrupted CLR NGEN binary was detected on the system.
	ERROR_APPX_INTEGRITY_FAILURE_CLR_NGEN = 15624

	// ERROR_RESILIENCY_FILE_CORRUPT
	// The operation could not be resumed because necessary data for recovery have been corrupted.
	ERROR_RESILIENCY_FILE_CORRUPT = 15625

	// ERROR_INSTALL_FIREWALL_SERVICE_NOT_RUNNING
	// The package could not be installed because the Windows Firewall service is not running. Enable the Windows Firewall service and try again.
	ERROR_INSTALL_FIREWALL_SERVICE_NOT_RUNNING = 15626

	// ERROR_PACKAGE_MOVE_FAILED
	// Package move failed.
	ERROR_PACKAGE_MOVE_FAILED = 15627

	// ERROR_INSTALL_VOLUME_NOT_EMPTY
	// The deployment operation failed because the volume is not empty.
	ERROR_INSTALL_VOLUME_NOT_EMPTY = 15628

	// ERROR_INSTALL_VOLUME_OFFLINE
	// The deployment operation failed because the volume is offline.
	ERROR_INSTALL_VOLUME_OFFLINE = 15629

	// ERROR_INSTALL_VOLUME_CORRUPT
	// The deployment operation failed because the specified volume is corrupt.
	ERROR_INSTALL_VOLUME_CORRUPT = 15630

	// ERROR_NEEDS_REGISTRATION
	// The deployment operation failed because the specified application needs to be registered first.
	ERROR_NEEDS_REGISTRATION = 15631

	// ERROR_INSTALL_WRONG_PROCESSOR_ARCHITECTURE
	// The deployment operation failed because the package targets the wrong processor architecture.
	ERROR_INSTALL_WRONG_PROCESSOR_ARCHITECTURE = 15632

	// ERROR_DEV_SIDELOAD_LIMIT_EXCEEDED
	// You have reached the maximum number of developer sideloaded packages allowed on this device. Please uninstall a sideloaded package and try again.
	ERROR_DEV_SIDELOAD_LIMIT_EXCEEDED = 15633

	// ERROR_INSTALL_OPTIONAL_PACKAGE_REQUIRES_MAIN_PACKAGE
	// A main app package is required to install this optional package.  Install the main package first and try again.
	ERROR_INSTALL_OPTIONAL_PACKAGE_REQUIRES_MAIN_PACKAGE = 15634

	// ERROR_PACKAGE_NOT_SUPPORTED_ON_FILESYSTEM
	// This app package type is not supported on this filesystem
	ERROR_PACKAGE_NOT_SUPPORTED_ON_FILESYSTEM = 15635

	// ERROR_PACKAGE_MOVE_BLOCKED_BY_STREAMING
	// Package move operation is blocked until the application has finished streaming
	ERROR_PACKAGE_MOVE_BLOCKED_BY_STREAMING = 15636

	// ERROR_INSTALL_OPTIONAL_PACKAGE_APPLICATIONID_NOT_UNIQUE
	// A main or another optional app package has the same application ID as this optional package.  Change the application ID for the optional package to avoid conflicts.
	ERROR_INSTALL_OPTIONAL_PACKAGE_APPLICATIONID_NOT_UNIQUE = 15637

	// ERROR_PACKAGE_STAGING_ONHOLD
	// This staging session has been held to allow another staging operation to be prioritized.
	ERROR_PACKAGE_STAGING_ONHOLD = 15638

	// ERROR_INSTALL_INVALID_RELATED_SET_UPDATE
	// A related set cannot be updated because the updated set is invalid. All packages in the related set must be updated at the same time.
	ERROR_INSTALL_INVALID_RELATED_SET_UPDATE = 15639

	// ERROR_INSTALL_OPTIONAL_PACKAGE_REQUIRES_MAIN_PACKAGE_FULLTRUST_CAPABILITY
	// An optional package with a FullTrust entry point requires the main package to have the runFullTrust capability.
	ERROR_INSTALL_OPTIONAL_PACKAGE_REQUIRES_MAIN_PACKAGE_FULLTRUST_CAPABILITY = 15640

	// ERROR_DEPLOYMENT_BLOCKED_BY_USER_LOG_OFF
	// An error occurred because a user was logged off.
	ERROR_DEPLOYMENT_BLOCKED_BY_USER_LOG_OFF = 15641

	// ERROR_PROVISION_OPTIONAL_PACKAGE_REQUIRES_MAIN_PACKAGE_PROVISIONED
	// An optional package provision requires the dependency main package to also be provisioned.
	ERROR_PROVISION_OPTIONAL_PACKAGE_REQUIRES_MAIN_PACKAGE_PROVISIONED = 15642

	// ERROR_PACKAGES_REPUTATION_CHECK_FAILED
	// The packages failed the SmartScreen reputation check.
	ERROR_PACKAGES_REPUTATION_CHECK_FAILED = 15643

	// ERROR_PACKAGES_REPUTATION_CHECK_TIMEDOUT
	// The SmartScreen reputation check operation timed out.
	ERROR_PACKAGES_REPUTATION_CHECK_TIMEDOUT = 15644

	// ERROR_DEPLOYMENT_OPTION_NOT_SUPPORTED
	// The current deployment option is not supported.
	ERROR_DEPLOYMENT_OPTION_NOT_SUPPORTED = 15645

	// ERROR_APPINSTALLER_ACTIVATION_BLOCKED
	// Activation is blocked due to the .appinstaller update settings for this app.
	ERROR_APPINSTALLER_ACTIVATION_BLOCKED = 15646

	// ERROR_REGISTRATION_FROM_REMOTE_DRIVE_NOT_SUPPORTED
	// Remote drives are not supported; use \\server\share to register a remote package.
	ERROR_REGISTRATION_FROM_REMOTE_DRIVE_NOT_SUPPORTED = 15647

	//////////////////////////
	//                      //
	// AppModel Error Codes //
	//    =15700-15720      //
	//                      //
	//////////////////////////

	// APPMODEL_ERROR_NO_PACKAGE
	// The process has no package identity.

	APPMODEL_ERROR_NO_PACKAGE = 15700

	// APPMODEL_ERROR_PACKAGE_RUNTIME_CORRUPT
	// The package runtime information is corrupted.

	APPMODEL_ERROR_PACKAGE_RUNTIME_CORRUPT = 15701

	// APPMODEL_ERROR_PACKAGE_IDENTITY_CORRUPT
	// The package identity is corrupted.

	APPMODEL_ERROR_PACKAGE_IDENTITY_CORRUPT = 15702

	// APPMODEL_ERROR_NO_APPLICATION
	// The process has no application identity.

	APPMODEL_ERROR_NO_APPLICATION = 15703

	// APPMODEL_ERROR_DYNAMIC_PROPERTY_READ_FAILED
	// One or more AppModel Runtime group policy values could not be read. Please contact your system administrator with the contents of your AppModel Runtime event log.

	APPMODEL_ERROR_DYNAMIC_PROPERTY_READ_FAILED = 15704

	// APPMODEL_ERROR_DYNAMIC_PROPERTY_INVALID
	// One or more AppModel Runtime group policy values are invalid. Please contact your system administrator with the contents of your AppModel Runtime event log.

	APPMODEL_ERROR_DYNAMIC_PROPERTY_INVALID = 15705

	// APPMODEL_ERROR_PACKAGE_NOT_AVAILABLE
	// The package is currently not available.

	APPMODEL_ERROR_PACKAGE_NOT_AVAILABLE = 15706

	/////////////////////////////
	//                         //
	// Appx StateManager Codes //
	//    =15800-15840         //
	//                         //
	/////////////////////////////

	// ERROR_STATE_LOAD_STORE_FAILED
	// Loading the state store failed.
	ERROR_STATE_LOAD_STORE_FAILED = 15800

	// ERROR_STATE_GET_VERSION_FAILED
	// Retrieving the state version for the application failed.
	ERROR_STATE_GET_VERSION_FAILED = 15801

	// ERROR_STATE_SET_VERSION_FAILED
	// Setting the state version for the application failed.
	ERROR_STATE_SET_VERSION_FAILED = 15802

	// ERROR_STATE_STRUCTURED_RESET_FAILED
	// Resetting the structured state of the application failed.
	ERROR_STATE_STRUCTURED_RESET_FAILED = 15803

	// ERROR_STATE_OPEN_CONTAINER_FAILED
	// State Manager failed to open the container.
	ERROR_STATE_OPEN_CONTAINER_FAILED = 15804

	// ERROR_STATE_CREATE_CONTAINER_FAILED
	// State Manager failed to create the container.
	ERROR_STATE_CREATE_CONTAINER_FAILED = 15805

	// ERROR_STATE_DELETE_CONTAINER_FAILED
	// State Manager failed to delete the container.
	ERROR_STATE_DELETE_CONTAINER_FAILED = 15806

	// ERROR_STATE_READ_SETTING_FAILED
	// State Manager failed to read the setting.
	ERROR_STATE_READ_SETTING_FAILED = 15807

	// ERROR_STATE_WRITE_SETTING_FAILED
	// State Manager failed to write the setting.
	ERROR_STATE_WRITE_SETTING_FAILED = 15808

	// ERROR_STATE_DELETE_SETTING_FAILED
	// State Manager failed to delete the setting.
	ERROR_STATE_DELETE_SETTING_FAILED = 15809

	// ERROR_STATE_QUERY_SETTING_FAILED
	// State Manager failed to query the setting.
	ERROR_STATE_QUERY_SETTING_FAILED = 15810

	// ERROR_STATE_READ_COMPOSITE_SETTING_FAILED
	// State Manager failed to read the composite setting.
	ERROR_STATE_READ_COMPOSITE_SETTING_FAILED = 15811

	// ERROR_STATE_WRITE_COMPOSITE_SETTING_FAILED
	// State Manager failed to write the composite setting.
	ERROR_STATE_WRITE_COMPOSITE_SETTING_FAILED = 15812

	// ERROR_STATE_ENUMERATE_CONTAINER_FAILED
	// State Manager failed to enumerate the containers.
	ERROR_STATE_ENUMERATE_CONTAINER_FAILED = 15813

	// ERROR_STATE_ENUMERATE_SETTINGS_FAILED
	// State Manager failed to enumerate the settings.
	ERROR_STATE_ENUMERATE_SETTINGS_FAILED = 15814

	// ERROR_STATE_COMPOSITE_SETTING_VALUE_SIZE_LIMIT_EXCEEDED
	// The size of the state manager composite setting value has exceeded the limit.
	ERROR_STATE_COMPOSITE_SETTING_VALUE_SIZE_LIMIT_EXCEEDED = 15815

	// ERROR_STATE_SETTING_VALUE_SIZE_LIMIT_EXCEEDED
	// The size of the state manager setting value has exceeded the limit.
	ERROR_STATE_SETTING_VALUE_SIZE_LIMIT_EXCEEDED = 15816

	// ERROR_STATE_SETTING_NAME_SIZE_LIMIT_EXCEEDED
	// The length of the state manager setting name has exceeded the limit.
	ERROR_STATE_SETTING_NAME_SIZE_LIMIT_EXCEEDED = 15817

	// ERROR_STATE_CONTAINER_NAME_SIZE_LIMIT_EXCEEDED
	// The length of the state manager container name has exceeded the limit.
	ERROR_STATE_CONTAINER_NAME_SIZE_LIMIT_EXCEEDED = 15818

	/////////////////////////////////
	//                             //
	// Application Partition Codes //
	//    =15841-15860             //
	//                             //
	/////////////////////////////////

	// ERROR_API_UNAVAILABLE
	// This API cannot be used in the context of the caller's application type.
	ERROR_API_UNAVAILABLE = 15841

	/////////////////////////////////
	//                             //
	// Windows Store Codes         //
	//    =15861-15880             //
	//                             //
	/////////////////////////////////

	// STORE_ERROR_UNLICENSED
	// This PC does not have a valid license for the application or product.

	STORE_ERROR_UNLICENSED = 15861

	// STORE_ERROR_UNLICENSED_USER
	// The authenticated user does not have a valid license for the application or product.

	STORE_ERROR_UNLICENSED_USER = 15862

	// STORE_ERROR_PENDING_COM_TRANSACTION
	// The commerce transaction associated with this license is still pending verification.

	STORE_ERROR_PENDING_COM_TRANSACTION = 15863

	// STORE_ERROR_LICENSE_REVOKED
	// The license has been revoked for this user.
	STORE_ERROR_LICENSE_REVOKED = 15864
)
