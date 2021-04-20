package response

const _externalErrorCode = 4000
const _internalErrorCode = 5000

const (
	SuccessCode uint64 = 2000

	ErrInvalidRequestCode uint64 = _externalErrorCode + 1
	ErrInvalidTokenCode   uint64 = _externalErrorCode + 6

	ErrDatabaseCode   uint64 = _internalErrorCode + 1
	ErrThirdPartyCode uint64 = _internalErrorCode + 2
)

const (
	SuccessMessageEN           string = "Success."
	ErrInternalServerMessageEN string = "Internal server error."
	// Authentication
	ErrAuthenticationTokenMessageEN string = "Authentication Token Fail."
	// PushMessage
	SuccessPushMessageMessageEN string = "Push message notification success."
	ErrPushMessageNotiMessageEN string = "Cannot push message notification."
	// ReplyMessage
	SuccessReplyMessageMessageEN string = "Reply message callback success."
	ErrReplyMessageBackMessageEN string = "Cannot reply message callback."
	// Description
	ErrContactAdminDescEN string = "Please contact administrator for more information."
)

const (
	SuccessMessageTH           string = "สำเร็ว."
	ErrInternalServerMessageTH string = "มีข้อผิดพลาดภายในเซิร์ฟเวอร์."
	// Authentication
	ErrAuthenticationTokenMessageTH string = "โทเคนสำหรับยืนยันตัวตนไม่ถูกต้อง."
	// PushMessage
	SuccessPushMessageMessageTH string = "ส่งข้อความแจ้งเตือนสำเร็จ."
	ErrPushMessageNotiMessageTH string = "ไม่สามารถส่งข้อความแจ้งเตือนได้."
	// ReplyMessage
	SuccessReplyMessageMessageTH string = "ส่งข้อความตอบกลับสำเร็จ."
	ErrReplyMessageBackMessageTH string = "ไม่สามารถส่งข้อความตอบกลับได้."
	// Description
	ErrContactAdminDescTH string = "กรุณาติดต่อเจ้าหน้าที่ดูแลระบบเพื่อรับข้อมูลเพิ่มเติม."
)

const (
	ValidateAuthenticationTokenError string = "Invalid Authentication Token"
	ValidateFieldError               string = "Invalid Parameters"
)
