package contact

type Contact struct {
	Name   string   `json:"name"`
	Phones []string `json:"phones"`
	Note   string   `json:"note,omitempty"`
}

type CampusContacts struct {
	Peace   []Contact `json:"peace"`
	Yanchao []Contact `json:"yanchao"`
}

type Contacts struct {
	Emergency  []Contact      `json:"emergency"`
	Security   CampusContacts `json:"security"`
	Guard      CampusContacts `json:"guard"`
	Dorms      CampusContacts `json:"dorms"`
	Life       CampusContacts `json:"life"`
	Counselor  Contact        `json:"counselor"`
	Health     CampusContacts `json:"health"`
	Indigenous Contact        `json:"indigenous"`
}
