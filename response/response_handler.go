package response

import (
	"context"
	"line-notification/common"
)

var (
	EN = Global{
		AuthenticationTokenInvalid: ErrResponse{Code: ErrInvalidTokenCode, Message: ErrAuthenticationTokenMessageEN},
		PushMessageSuccess:         Response{Code: SuccessCode, Message: SuccessPushMessageMessageEN},
		PushMessageValidateReq:     ErrResponse{Code: ErrInvalidRequestCode, Message: ErrPushMessageNotiMessageEN},
		PushMessageThirdParty:      ErrResponse{Code: ErrThirdPartyCode, Message: ErrPushMessageNotiMessageEN},
		ReplyMessageSuccess:        Response{Code: SuccessCode, Message: SuccessReplyMessageMessageEN},
		ReplyMessageValidateReq:    ErrResponse{Code: ErrInvalidRequestCode, Message: ErrReplyMessageBackMessageEN},
		ReplyMessageThirdParty:     ErrResponse{Code: ErrThirdPartyCode, Message: ErrReplyMessageBackMessageEN},
	}

	TH = Global{
		AuthenticationTokenInvalid: ErrResponse{Code: ErrInvalidTokenCode, Message: ErrAuthenticationTokenMessageTH},
		PushMessageSuccess:         Response{Code: SuccessCode, Message: SuccessPushMessageMessageTH},
		PushMessageValidateReq:     ErrResponse{Code: ErrInvalidRequestCode, Message: ErrPushMessageNotiMessageTH},
		PushMessageThirdParty:      ErrResponse{Code: ErrThirdPartyCode, Message: ErrPushMessageNotiMessageTH},
		ReplyMessageSuccess:        Response{Code: SuccessCode, Message: SuccessReplyMessageMessageTH},
		ReplyMessageValidateReq:    ErrResponse{Code: ErrInvalidRequestCode, Message: ErrReplyMessageBackMessageTH},
		ReplyMessageThirdParty:     ErrResponse{Code: ErrThirdPartyCode, Message: ErrReplyMessageBackMessageTH},
	}

	Language = map[interface{}]Global{
		"en": EN,
		"th": TH,
	}
)

type Global struct {
	AuthenticationTokenInvalid ErrResponse
	PushMessageSuccess         Response
	PushMessageValidateReq     ErrResponse
	PushMessageThirdParty      ErrResponse
	ReplyMessageSuccess        Response
	ReplyMessageValidateReq    ErrResponse
	ReplyMessageThirdParty     ErrResponse
}

func ResponseContextLocale(ctx context.Context) *Global {
	v := ctx.Value(common.LocaleKey)
	if v == nil {
		return nil
	}
	l, ok := Language[v]
	if ok {
		return &l
	}
	return &EN
}
