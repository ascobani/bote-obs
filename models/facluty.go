/*
	ICI Tech Teknoloji A.Ş.
	User        : ASCI
	Name        : Ali Sinan COBANI
	Date        : 13.03.2026
	Time        : 10:16
	Notes       :


*/

package models

import "github.com/gofrs/uuid/v5"

type Faculty struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string    `gorm:"type:varchar(255);not null"`

	// Relation Ships
	Departments []Departments `gorm:"foreignKey:FacultyID"`
	BaseRecordFields
}
