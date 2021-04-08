package response

import (
	"context"
	"line-notification/common"
)

var (
	EN = Global{}

	TH = Global{}

	Language = map[interface{}]Global{
		"en": EN,
		"th": TH,
	}
)

type Global struct {
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
