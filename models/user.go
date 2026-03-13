/*
	ICI Tech Teknoloji A.Ş.
	User        : ASCI
	Name        : Ali Sinan COBANI
	Date        : 13.03.2026
	Time        : 10:05
	Notes       :


*/

package models

import (
	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID                  uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v5(uuid_nil(), gen_random_uuid())"' json:"id"`
	StudNum             string       `gorm:"type:varchar(255);not null" json:"stud_num"`
	Name                string       `gorm:"type:varchar(255);not null" json:"name"`
	SurName             string       `gorm:"type:varchar(255);not null" json:"surname"`
	PhoneNumber1        string       `gorm:"type:varchar(255);not null" json:"phone_number_1"`
	PhoneNumber2        string       `gorm:"type:varchar(255);" json:"phone_number_2"`
	PhoneNumber3        string       `gorm:"type:varchar(255);" json:"phone_number_3"`
	Email1              string       `gorm:"type:varchar(255);not null" json:"email_1"`
	Email2              string       `gorm:"type:varchar(255);" json:"email_2"`
	FamilyAddress       string       `gorm:"type:varchar(255);" json:"family_address"`
	FamilyCity          string       `gorm:"type:varchar(100);" json:"family_city"`
	FamilyState         string       `gorm:"type:varchar(100);" json:"family_state"`
	FamilyPostalCode    string       `gorm:"type:varchar(20);" json:"family_postal_code"`
	FamilyCountry       string       `gorm:"type:varchar(100);" json:"family_country"`
	ResidenceAddress    string       `gorm:"type:varchar(255);" json:"residence_address"`
	ResidenceCity       string       `gorm:"type:varchar(100);" json:"residence_city"`
	ResidenceState      string       `gorm:"type:varchar(100);" json:"residence_state"`
	ResidencePostalCode string       `gorm:"type:varchar(20);" json:"residence_postal_code"`
	ResidenceCountry    string       `gorm:"type:varchar(100);" json:"residence_country"`
	Password            string       `gorm:"type:varchar(255);not null" json:"password"`
	UserRoles           []UserRole   `json:"user_roles,omitempty" gorm:"foreignKey:UserID"`
	Enrollments         []Enrollment `json:"enrollments,omitempty" gorm:"foreignKey:UserID"`
	BaseRecordFields
}
