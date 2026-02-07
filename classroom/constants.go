package classroom

import "sync"

var buildingMap = map[string]string{
	"0":  "行政大樓",
	"1":  "教育大樓",
	"3":  "文學大樓",
	"4":  "綜合大樓",
	"5":  "藝術大樓",
	"6":  "活動中心",
	"7":  "愛閱館(原電算中心)",
	"8":  "雋永樓",
	"9":  "研究大樓",
	"CM": "寰宇大樓",
	"TC": "科技大樓",
	"MA": "致理大樓",
	"PH": "高斯大樓",
	"BT": "生科大樓",
	"LI": "圖書資訊大樓",
}

var _floors = []string{"02", "03", "07", "09", "12", "13", "14", "15", "31", "32", "33", "34", "35", "41", "42", "43", "44", "51", "52", "53", "54", "61", "62", "63", "65", "66", "68", "69", "71", "72", "73", "74", "91", "92", "93", "BT1", "BT2", "BT3", "CM1", "CM2", "CM3", "CM4", "CM5", "CM6", "LI1", "MA1", "MA2", "MA3", "MA4", "MA5", "MA6", "MA7", "MA8", "PH1", "PH2", "PH3", "PH4", "PH5", "PH7", "TC1", "TC2", "TC3", "TC4", "TC5"}

type _floorMap map[string]struct{}

var (
	floorMap _floorMap
	once     sync.Once
	mu       sync.RWMutex
)

func initFloorMap() {
	floorMap = make(_floorMap, len(_floors))
	for _, code := range _floors {
		floorMap[code] = struct{}{}
	}
}

func getValidFloorMap() _floorMap {
	once.Do(func() {
		mu.Lock()
		defer mu.Unlock()
		if floorMap == nil {
			initFloorMap()
		}
	})
	return floorMap
}

func HasMap(buildingID, floor string) bool {
	codes := getValidFloorMap()
	mu.RLock()
	defer mu.RUnlock()
	_, exists := codes[buildingID+floor]
	return exists
}
