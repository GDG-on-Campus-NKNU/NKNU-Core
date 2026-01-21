package sso

type Session struct {
	AspNETSessionId string `json:"aspNETSessionId"`
	ViewState       string `json:"viewState"`
}
