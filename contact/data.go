package contact

var Data = Contacts{
	Emergency: []Contact{
		{Name: "學校總機", Phones: []string{"07-7172930"}},
	},
	Security: CampusContacts{
		Peace: []Contact{
			{Name: "校安中心值勤室", Phones: []string{"07-7172930#1537", "0910-783-882"}, Note: "上班時間轉分機"},
		},
		Yanchao: []Contact{
			{Name: "校安中心值勤室", Phones: []string{"07-6051116", "07-7172930#6531", "0910-783-992"}, Note: "上班時間轉分機"},
		},
	},
	Guard: CampusContacts{
		Peace: []Contact{
			{Name: "警衛室", Phones: []string{"07-7172930#1119"}},
		},
		Yanchao: []Contact{
			{Name: "警衛室", Phones: []string{"07-7172930#6119"}},
		},
	},
	Dorms: CampusContacts{
		Peace: []Contact{
			{Name: "男舍-逸清", Phones: []string{"07-7175200", "07-7172930#3713"}},
			{Name: "女舍-芝心", Phones: []string{"07-7175143", "07-7172930#3723"}},
			{Name: "女舍-蘭苑", Phones: []string{"07-7175001", "07-7172930#3743"}},
			{Name: "女舍-涵泳", Phones: []string{"07-7175601", "07-7172930#3733"}},
		},
		Yanchao: []Contact{
			{Name: "男舍-霽遠樓", Phones: []string{"07-6052123"}},
			{Name: "女舍-詠絮樓", Phones: []string{"07-6052149"}},
			{Name: "燕窩男舍", Phones: []string{"07-6052722"}},
			{Name: "燕窩女舍", Phones: []string{"07-6052708"}},
		},
	},
	Life: CampusContacts{
		Peace: []Contact{
			{Name: "生活輔導組 (組長)", Phones: []string{"07-7172930#1230"}},
		},
		Yanchao: []Contact{
			{Name: "生活輔導組 (值勤)", Phones: []string{"07-7172930#6531"}},
			{Name: "生活輔導組 (宿舍管理)", Phones: []string{"07-7172930#6532"}},
			{Name: "生活輔導組 (宿舍修繕)", Phones: []string{"07-7172930#6533"}},
		},
	},
	Counselor: Contact{
		Name: "學生輔導中心專線", Phones: []string{"07-7219054"},
	},
	Health: CampusContacts{
		Peace: []Contact{
			{Name: "衛生保健組", Phones: []string{"07-7172930#1291", "07-7172930#1292"}},
		},
		Yanchao: []Contact{
			{Name: "衛生保健組", Phones: []string{"07-7172930#6291", "07-7172930#6292"}},
		},
	},
	Indigenous: Contact{
		Name: "原住民族學生資源中心", Phones: []string{"07-7172930#1259"},
	},
}
