package types

import (
	"strings"
)

const (
	DEFAULT_ERROR_TYPE_SEP = "|"
)

//-----------------------------------------------------------------------------
type ErrorType int

const (
	ErrorType_OK                    ErrorType = 0       //ok
	ErrorType_DiskMissing           ErrorType = 1 << 0  //disk is missing
	ErrorType_FileMissing           ErrorType = 1 << 1  //file is missing or damage
	ErrorType_TicketExpired         ErrorType = 1 << 2  //ticket expired (PC1+PC2 take a long time more than 25 hours)
	ErrorType_GPU                   ErrorType = 1 << 3  //GPU error
	ErrorType_DiskIO                ErrorType = 1 << 4  //disk input/output error
	ErrorType_TaskAbort             ErrorType = 1 << 5  //task abort
	ErrorType_RustError             ErrorType = 1 << 6  //rust error
	ErrorType_InvalidProof          ErrorType = 1 << 7  //C2 compute proof failed
	ErrorType_WorkerClosed          ErrorType = 1 << 8  //worker closed
	ErrorType_MessageExecFailed     ErrorType = 1 << 9  //message execution failed
	ErrorType_MovingToStorageFailed ErrorType = 1 << 10 //moving sector to storage error
	ErrorType_FileInconsistent      ErrorType = 1 << 11 //file inconsistent
	ErrorType_ComputeFailed         ErrorType = 1 << 12 //consecutive compute fails
	ErrorType_InvalidHashedNode     ErrorType = 1 << 13 //invalid hashed node length
	ErrorType_WebSocketClosed       ErrorType = 1 << 14 //websocket connection closed
	ErrorType_ConnectTimeout        ErrorType = 1 << 15 //connect timeout
	ErrorType_SectorNotFound        ErrorType = 1 << 16 //sector not found
	ErrorType_RpcConnClosed         ErrorType = 1 << 17 //rpc connection close error

	//---------------------------------------------------------------
	ErrorType_Unknown ErrorType = 1 << 30 //unknown error (keep biggest of int)
)

func (t ErrorType) GoString() string {
	return t.String()
}

func (t ErrorType) String() string {
	var errs []string

	if t == ErrorType_OK {
		return "OK"
	}

	if t&ErrorType_Unknown > 0 {
		errs = append(errs, "UnknownError")
	}
	if t&ErrorType_DiskMissing > 0 {
		errs = append(errs, "DiskMissing")
	}
	if t&ErrorType_FileMissing > 0 {
		errs = append(errs, "FileMissing")
	}
	if t&ErrorType_TicketExpired > 0 {
		errs = append(errs, "TicketExpired")
	}
	if t&ErrorType_GPU > 0 {
		errs = append(errs, "GPU")
	}
	if t&ErrorType_DiskIO > 0 {
		errs = append(errs, "DiskIO")
	}
	if t&ErrorType_TaskAbort > 0 {
		errs = append(errs, "TaskAbort")
	}
	if t&ErrorType_RustError > 0 {
		errs = append(errs, "RustError")
	}
	if t&ErrorType_InvalidProof > 0 {
		errs = append(errs, "InvalidProof")
	}
	if t&ErrorType_WorkerClosed > 0 {
		errs = append(errs, "WorkerClosed")
	}
	if t&ErrorType_MessageExecFailed > 0 {
		errs = append(errs, "MessageExecFailed")
	}
	if t&ErrorType_MovingToStorageFailed > 0 {
		errs = append(errs, "MovingToStorageFailed")
	}
	if t&ErrorType_FileInconsistent > 0 {
		errs = append(errs, "FileInconsistent")
	}
	if t&ErrorType_ComputeFailed > 0 {
		errs = append(errs, "ComputeFailed")
	}
	if t&ErrorType_InvalidHashedNode > 0 {
		errs = append(errs, "InvalidHashedNode")
	}
	if t&ErrorType_WebSocketClosed > 0 {
		errs = append(errs, "WebSocketClosed")
	}
	if t&ErrorType_ConnectTimeout > 0 {
		errs = append(errs, "ConnectTimeout")
	}
	if t&ErrorType_SectorNotFound > 0 {
		errs = append(errs, "SectorNotFound")
	}
	if t&ErrorType_RpcConnClosed > 0 {
		errs = append(errs, "RpcConnClosed")
	}
	return strings.Join(errs, DEFAULT_ERROR_TYPE_SEP)
}

func MakeErrorTypeFromName(strType string) (et ErrorType) {

	for _, v := range strings.Split(strType, DEFAULT_ERROR_TYPE_SEP) {

		switch v {
		case "UnknownError":
			et |= ErrorType_Unknown
		case "OK":
			et |= ErrorType_OK
		case "DiskMissing":
			et |= ErrorType_DiskMissing
		case "FileMissing":
			et |= ErrorType_FileMissing
		case "TicketExpired":
			et |= ErrorType_TicketExpired
		case "GPU":
			et |= ErrorType_GPU
		case "DiskIO":
			et |= ErrorType_DiskIO
		case "TaskAbort":
			et |= ErrorType_TaskAbort
		case "RustError":
			et |= ErrorType_RustError
		case "InvalidProof":
			et |= ErrorType_InvalidProof
		case "WorkerClosed":
			et |= ErrorType_WorkerClosed
		case "MessageExecFailed":
			et |= ErrorType_MessageExecFailed
		case "MovingToStorageFailed":
			et |= ErrorType_MovingToStorageFailed
		case "FileInconsistent":
			et |= ErrorType_FileInconsistent
		case "ComputeFailed":
			et |= ErrorType_ComputeFailed
		case "InvalidHashedNode":
			et |= ErrorType_InvalidHashedNode
		case "WebSocketClosed":
			et |= ErrorType_WebSocketClosed
		case "ConnectTimeout":
			et |= ErrorType_ConnectTimeout
		case "SectorNotFound":
			et |= ErrorType_SectorNotFound
		case "RpcConnClosed":
			et |= ErrorType_RpcConnClosed
		}
	}
	return
}

//JSON marshal implement
func (t ErrorType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

//JSON unmarshal implement
func (t *ErrorType) UnmarshalText(value []byte) error {
	*t = MakeErrorTypeFromName(string(value))
	return nil
}
