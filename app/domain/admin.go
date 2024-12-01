package domain

func (Admin) TableName() string {
    return "admins"
}

type Admin struct {
	ID           	int    `gorm:"primaryKey" json:"id"`
	Name         	string `json:"name"`
	NameKana  		string `json:"name_kana"`
	Email 			string `json:"email"`
	Role 	 		string `json:"role"`
}


func NewAdmin(
	ID 				int,
	name			string,
	nameKana 		string,
	email 			string,
	role 			string,
) Admin {
	return Admin{
		ID: 			ID,
		Name:			name,
		NameKana: 		nameKana,
		Email: 			email,
		Role: 			role,
	}
}
