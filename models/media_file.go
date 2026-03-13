/*
	ICI Tech Teknoloji A.Ş.
	User        : ASCI
	Name        : Ali Sinan COBANI
	Date        : 13.03.2026
	Time        : 10:21
	Notes       :


*/

package models

import "github.com/gofrs/uuid/v5"

type MediaPurpose string

const (
	MediaPurposeProfilePhoto MediaPurpose = "ProfilePhoto"
	MediaPurposeLogo         MediaPurpose = "Logo"
	MediaPurposeSyllabus     MediaPurpose = "Syllabus"
	MediaPurposeTranscript   MediaPurpose = "Transcript"
	MediaPurposeDocument     MediaPurpose = "Document"
	MediaPurposeEventFlyer   MediaPurpose = "EventFlyer"
	MediaPurposeAssignment   MediaPurpose = "Assignment"
	MediaPurposeCertificate  MediaPurpose = "Certificate"
)

type MediaFileStatus string

const (
	MediaFileStatusPending  MediaFileStatus = "Pending"
	MediaFileStatusUploaded MediaFileStatus = "Uploaded"
	MediaFileStatusFailed   MediaFileStatus = "Failed"
)

// MediaFile stores metadata for any file uploaded to S3 or a compatible storage bucket.
// OwnerID + OwnerType form a polymorphic association so any model can attach files.
// The actual file is stored in the bucket at the given Key — never store raw file data here.
type MediaFile struct {
	ID        uuid.UUID       `gorm:"primaryKey" json:"id"`
	OwnerID   uuid.UUID       `gorm:"not null;index:idx_media_owner" json:"owner_id"`
	OwnerType string          `gorm:"size:100;not null;index:idx_media_owner" json:"owner_type"`
	Purpose   MediaPurpose    `gorm:"size:50;not null" json:"purpose"`
	Status    MediaFileStatus `gorm:"size:20;not null;default:'Pending'" json:"status"`
	Bucket    string          `gorm:"size:255;not null" json:"bucket"`
	Key       string          `gorm:"size:1000;not null" json:"key"`
	FileName  string          `gorm:"size:255;not null" json:"file_name"`
	MimeType  string          `gorm:"size:100;not null" json:"mime_type"`
	Size      int64           `gorm:"not null;default:0" json:"size"`
	BaseRecordFields
}
