/*
	ICI Tech Teknoloji A.Ş.
	User        : ASCI
	Name        : Ali Sinan COBANI
	Date        : 13.03.2026
	Time        : 10:19
	Notes       :


*/

package models

import "github.com/gofrs/uuid/v5"

type Departments struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string    `gorm:"type:varchar(255);not null"`

	FacultyID uuid.UUID `gorm:"not null"`
	Faculty   Faculty   `gorm:"foreignKey:FacultyID"`
	BaseRecordFields
}
