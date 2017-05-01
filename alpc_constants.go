package w32

const (
	ALPC_PORFLG_ALLOW_LPC_REQUESTS = 0x20000
	ALPC_PORFLG_SYSTEM_PROCESS     = 0x100000
	ALPC_PORFLG_WAITABLE_PORT      = 0x40000
)

const (
	ALPC_MSGFLG_REPLY_MESSAGE   = 0x1
	ALPC_MSGFLG_LPC_MODE        = 0x2     // ?
	ALPC_MSGFLG_RELEASE_MESSAGE = 0x10000 // dbg
	ALPC_MSGFLG_SYNC_REQUEST    = 0x20000 // dbg
	ALPC_MSGFLG_WAIT_USER_MODE  = 0x100000
	ALPC_MSGFLG_WAIT_ALERTABLE  = 0x200000
	ALPC_MSGFLG_WOW64_CALL      = 0x80000000 // dbg
)
const (
	ALPC_MESSAGE_SECURITY_ATTRIBUTE = 0x80000000
	ALPC_MESSAGE_VIEW_ATTRIBUTE     = 0x40000000
	ALPC_MESSAGE_CONTEXT_ATTRIBUTE  = 0x20000000
	ALPC_MESSAGE_HANDLE_ATTRIBUTE   = 0x10000000
)

const (
	OBJ_INHERIT          = 0x00000002
	OBJ_PERMANENT        = 0x00000010
	OBJ_EXCLUSIVE        = 0x00000020
	OBJ_CASE_INSENSITIVE = 0x00000040
	OBJ_OPENIF           = 0x00000080
	OBJ_OPENLINK         = 0x00000100
	OBJ_KERNEL_HANDLE    = 0x00000200
)

const (
	LPC_REQUEST               = 1
	LPC_REPLY                 = 2
	LPC_DATAGRAM              = 3
	LPC_LOST_REPLY            = 4
	LPC_PORT_CLOSED           = 5
	LPC_CLIENT_DIED           = 6
	LPC_EXCEPTION             = 7
	LPC_DEBUG_EVENT           = 8
	LPC_ERROR_EVENT           = 9
	LPC_CONNECTION_REQUEST    = 10
	LPC_CONTINUATION_REQUIRED = 0x2000
)

const (
	SecurityAnonymous      uint32 = 1
	SecurityIdentification uint32 = 2
	SecurityImpersonation  uint32 = 3
	SecurityDelegation     uint32 = 4
)

const (
	SECURITY_DYNAMIC_TRACKING byte = 1
	SECURITY_STATIC_TRACKING  byte = 0
)

const (
	ALPC_SYNC_OBJECT_TYPE   uint32 = 2
	ALPC_THREAD_OBJECT_TYPE uint32 = 4
)
