package domain

func (Instructor) TableName() string {
    return "instructors"
}

type Instructor struct {
	ID           	int    `gorm:"primaryKey" json:"id"`
	Name         	string `json:"name"`
	NameKana  		string `json:"name_kana"`
	Email 			string `json:"email"`
	Tel 			string `json:"tel"`
	Pass 			string `json:"pass"`
	InitialSetup 	string `json:"initial_setup"`
	PrefectureID 	string `json:"prefecture_id"`
	Introduced 		string `json:"introduced"`
	CanEditEvents	string `json:"can_edit_events"`
}

func NewInstructor(
	ID           	int,
	name         	string,
	nameKana  		string,
	email 			string,
	tel 			string,
	pass 			string,
	initialSetup 	string,
	prefectureID 	string,
	introduced 		string,
	canEditEvents	string,
) Instructor {
	return Instructor{
		ID:           	ID,
		Name:         	name,
		NameKana:  		nameKana,
		Email: 			email,
		Tel: 			tel,
		Pass: 			pass,
		InitialSetup: 	initialSetup,
		PrefectureID: 	prefectureID,
		Introduced: 	introduced,
		CanEditEvents:	canEditEvents,
	}
}
