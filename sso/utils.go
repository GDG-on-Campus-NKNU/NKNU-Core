package sso

import (
	error2 "NKNU-Core/sso/error"
	"strings"
)

func isSessionValidate(htmlContent string) error {
	for _, text := range []string{
		"輸入錯誤:「帳號」請輸入資料!",
		"輸入錯誤:「密碼」請輸入資料!",
		"帳號不存在或是密碼錯誤，請確定您輸入的帳號密碼正確或使用「忘記密碼」功能。",
		"您尚未輸入帳號",
	} {
		if strings.Contains(htmlContent, text) {
			return &error2.NoLoginError{
				Title:   text,
				Message: text,
			}
		}
	}
	return nil
}
