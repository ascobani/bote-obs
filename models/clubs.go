/*
	ICI Tech Teknoloji A.Ş.
	User        : ASCI
	Name        : Ali Sinan COBANI
	Date        : 13.03.2026
	Time        : 11:23
	Notes       :


*/

package models

import "github.com/gofrs/uuid/v5"

type ClubAdminRole string

const (
	ClubAdminRolePresident     ClubAdminRole = "President"
	ClubAdminRoleVicePresident ClubAdminRole = "VicePresident"
	ClubAdminRoleSecretary     ClubAdminRole = "Secretary"
	ClubAdminRoleAccountant    ClubAdminRole = "Accountant"
	ClubAdminRoleAdvisor       ClubAdminRole = "Advisor"
	ClubAdminFullMember        ClubAdminRole = "FullMember"
)

type ClubMemberStatus string

const (
	ClubMemberStatusActive   ClubMemberStatus = "Active"
	ClubMemberStatusInactive ClubMemberStatus = "Inactive"
	ClubMemberStatusPending  ClubMemberStatus = "Pending"
	ClubMemberStatusBanned   ClubMemberStatus = "Banned"
)

type Club struct {
	ID                  uuid.UUID `gorm:"primaryKey"`
	Name                string    `gorm:"size:255;not null"`
	CommunicationNumber string    `gorm:"size:50"`
	Email               string    `gorm:"size:255"`
	Active              bool      `gorm:"not null;default:true"`

	Administrators []ClubAdministrators `gorm:"foreignKey:ClubID"`
	Members        []ClubMember         `gorm:"foreignKey:ClubID"`
	BaseRecordFields
}

type ClubMember struct {
	ID     uuid.UUID        `gorm:"primaryKey"`
	ClubID uuid.UUID        `gorm:"not null;uniqueIndex:idx_club_member"`
	UserID uuid.UUID        `gorm:"not null;uniqueIndex:idx_club_member"`
	Status ClubMemberStatus `gorm:"size:20;not null;default:'Pending'"`

	Club Club `gorm:"foreignKey:ClubID"`
	User User `gorm:"foreignKey:UserID"`
	BaseRecordFields
}

type ClubAdministrators struct {
	ID     uuid.UUID     `gorm:"primaryKey"`
	ClubID uuid.UUID     `gorm:"not null"`
	UserID uuid.UUID     `gorm:"not null"`
	Role   ClubAdminRole `gorm:"size:20;not null"`

	Club Club `gorm:"foreignKey:ClubID"`
	User User `gorm:"foreignKey:UserID"`
	BaseRecordFields
}
