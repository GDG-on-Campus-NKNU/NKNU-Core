package main

import (
	"C"
	"encoding/json"
	"nknu-core/contact"
	"nknu-core/utils"
)

func GetContactEmergencyApi() string {
	dataBytes, err := json.Marshal(contact.Data.Emergency)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), nil)
}

func GetContactSecurityApi() string {
	dataBytes, err := json.Marshal(contact.Data.Security)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), nil)
}

func GetContactGuardApi() string {
	dataBytes, err := json.Marshal(contact.Data.Guard)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), nil)
}

func GetContactDormsApi() string {
	dataBytes, err := json.Marshal(contact.Data.Dorms)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), nil)
}

func GetContactLifeApi() string {
	dataBytes, err := json.Marshal(contact.Data.Life)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), nil)
}

func GetContactCounselorApi() string {
	dataBytes, err := json.Marshal(contact.Data.Counselor)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), nil)
}

func GetContactHealthApi() string {
	dataBytes, err := json.Marshal(contact.Data.Health)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), nil)
}

func GetContactIndigenousApi() string {
	dataBytes, err := json.Marshal(contact.Data.Indigenous)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), nil)
}

func GetContactAllApi() string {
	dataBytes, err := json.Marshal(contact.Data)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), nil)
}
