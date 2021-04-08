package response

const _externalErrorCode = 4000
const _internalErrorCode = 5000

const (
	SuccessCode uint64 = 2000

	ErrInvalidRequestCode uint64 = _externalErrorCode + 1

	ErrDatabaseCode uint64 = _internalErrorCode + 1
)

const (
	SuccessMessageEN           string = "Success."
	ErrInvalidRequestMessageEN string = "Input request error."
)

const (
	SuccessMessageTH           string = "สำเร็ว."
	ErrInvalidRequestMessageTH string = "ข้อมูลที่นำเข้าไม่ถูกต้อง."
)

const (
	ValidateFieldError string = "Invalid Parameters"
)
