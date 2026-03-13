/*
	ICI Tech Teknoloji A.Ş.
	User        : ASCI
	Name        : Ali Sinan COBANI
	Date        : 13.03.2026
	Time        : 10:32
	Notes       :


*/

package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type BaseRecordFields struct {
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	CreatedBy     uuid.UUID      `gorm:"type:uuid" json:"createdBy"`
	CreatedByName string         `gorm:"size:100" json:"createdByName"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	UpdatedBy     uuid.UUID      `gorm:"type:uuid" json:"updatedBy"`
	UpdatedByName string         `gorm:"size:100" json:"updatedByName"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
