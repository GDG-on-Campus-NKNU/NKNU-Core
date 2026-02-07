package main

/*
#include <stdlib.h>
*/
import "C"
import "unsafe"

//export Free
func Free(p *C.char) {
	if p == nil {
		return
	}
	C.free(unsafe.Pointer(p))
}

// SSO

//export GetSessionInfo
func GetSessionInfo() *C.char {
	return C.CString(GetSessionInfoApi())
}

//export Login
func Login(aspNetSessionId *C.char, viewState *C.char, account *C.char, password *C.char) *C.char {
	aspNetString := C.GoString(aspNetSessionId)
	viewStateString := C.GoString(viewState)
	parsedAccount := C.GoString(account)
	parsedPassword := C.GoString(password)
	result := LoginApi(aspNetString, viewStateString, parsedAccount, parsedPassword)
	return C.CString(result)
}

//export GetHistoryScore
func GetHistoryScore(aspNetSessionId *C.char, viewState *C.char) *C.char {
	aspString := C.GoString(aspNetSessionId)
	viewStateString := C.GoString(viewState)
	return C.CString(GetHistoryScoreApi(aspString, viewStateString))
}

//export GetMailServiceAccount
func GetMailServiceAccount(aspNetSessionId *C.char) *C.char {
	sessionID := C.GoString(aspNetSessionId)
	return C.CString(GetMailServiceAccountApi(sessionID))
}

// school bus schedule

//export LoadSavedData
func LoadSavedData(toYcData, toHpData *C.char) *C.char {
	toYcString := C.GoString(toYcData)
	toHpString := C.GoString(toHpData)
	return C.CString(LoadSavedDataApi(toYcString, toHpString))
}

//export RefreshSchoolBusData
func RefreshSchoolBusData() *C.char {
	return C.CString(RefreshSchoolBusDataApi())
}

//export GetLastSchoolBusDataFetchTime
func GetLastSchoolBusDataFetchTime() *C.char {
	return C.CString(GetLastSchoolBusDataFetchTimeApi())
}

//export GetYcToHpSchedule
func GetYcToHpSchedule() *C.char {
	return C.CString(GetYcToHpScheduleApi())
}

//export GetHpToYcSchedule
func GetHpToYcSchedule() *C.char {
	return C.CString(GetHpToYcScheduleApi())
}

//export GetYcToHpNextBusNow
func GetYcToHpNextBusNow() *C.char {
	return C.CString(GetYcToHpNextBusNowApi())
}

//export GetHpToYcNextBusNow
func GetHpToYcNextBusNow() *C.char {
	return C.CString(GetHpToYcNextBusNowApi())
}

//export GetYcToHpNextBus
func GetYcToHpNextBus(rawYear, rawMonth, rawDay, rawHour, rawMinute C.int) *C.char {
	year := int(rawYear)
	month := int(rawMonth)
	day := int(rawDay)
	hour := int(rawHour)
	minute := int(rawMinute)
	return C.CString(GetYcToHpNextBusApi(year, month, day, hour, minute))
}

//export GetHpToYcNextBus
func GetHpToYcNextBus(rawYear, rawMonth, rawDay, rawHour, rawMinute C.int) *C.char {
	year := int(rawYear)
	month := int(rawMonth)
	day := int(rawDay)
	hour := int(rawHour)
	minute := int(rawMinute)
	return C.CString(GetHpToYcNextBusApi(year, month, day, hour, minute))
}

//export GetYcToHpBusByIndex
func GetYcToHpBusByIndex(index C.int) *C.char {
	parsedIndex := int(index)
	return C.CString(GetYcToHpBusByIndexApi(parsedIndex))
}

//export GetHpToYcBusByIndex
func GetHpToYcBusByIndex(index C.int) *C.char {
	parsedIndex := int(index)
	return C.CString(GetHpToYcBusByIndexApi(parsedIndex))
}

// school news

//export CountNews
func CountNews() *C.char {
	return C.CString(CountNewsApi())
}

//export CountNewsByCategory
func CountNewsByCategory(category *C.char) *C.char {
	cat := C.GoString(category)
	return C.CString(CountNewsByCategoryApi(cat))
}

//export CountNewsByPublisher
func CountNewsByPublisher(publisher *C.char) *C.char {
	pub := C.GoString(publisher)
	return C.CString(CountNewsByPublisherApi(pub))
}

//export GetNews
func GetNews(startIndex, endIndex C.int) *C.char {
	return C.CString(GetNewsApi(int(startIndex), int(endIndex)))
}

//export GetNewsByCategory
func GetNewsByCategory(category *C.char, startIndex, endIndex C.int) *C.char {
	cat := C.GoString(category)
	return C.CString(GetNewsByCategoryApi(cat, int(startIndex), int(endIndex)))
}

//export GetNewsByPublisher
func GetNewsByPublisher(publisher *C.char, startIndex, endIndex C.int) *C.char {
	pub := C.GoString(publisher)
	return C.CString(GetNewsByPublisherApi(pub, int(startIndex), int(endIndex)))
}

//export ForceRefreshNews
func ForceRefreshNews() *C.char {
	return C.CString(ForceRefreshNewsApi())
}

//export GetLastNewsRefreshTime
func GetLastNewsRefreshTime() *C.char {
	return C.CString(GetLastNewsRefreshTimeApi())
}

// main

func main() {

}
