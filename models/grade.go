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

type LetterGrade string

const (
	LetterGradeAA LetterGrade = "AA" // 4.0 — 90-100
	LetterGradeBA LetterGrade = "BA" // 3.5 — 85-89
	LetterGradeBB LetterGrade = "BB" // 3.0 — 75-84
	LetterGradeCB LetterGrade = "CB" // 2.5 — 70-74
	LetterGradeCC LetterGrade = "CC" // 2.0 — 60-69
	LetterGradeDC LetterGrade = "DC" // 1.5 — 55-59
	LetterGradeDD LetterGrade = "DD" // 1.0 — 50-54
	LetterGradeFD LetterGrade = "FD" // 0.5 — 40-49
	LetterGradeFF LetterGrade = "FF" // 0.0 — 0-39
	LetterGradeNA LetterGrade = "F0" // 0.0 — not attended
)

// GradeWeightConfig defines the exam structure for a DepartmentCourse.
// The sum of all child ExamComponent weights must equal 100.
type GradeWeightConfig struct {
	ID                 uuid.UUID `gorm:"primaryKey" json:"id"`
	DepartmentCourseID uuid.UUID `gorm:"not null;uniqueIndex" json:"department_course_id"`

	DepartmentCourse DepartmentCourse `gorm:"foreignKey:DepartmentCourseID" json:"department_course,omitempty"`
	Components       []ExamComponent  `gorm:"foreignKey:GradeWeightConfigID" json:"components,omitempty"`
	BaseRecordFields
}

// ExamComponent defines a single exam entry (e.g. Midterm 1, Final, Makeup).
// Weight is a percentage — all components on a config must sum to 100.
// IsMakeup marks the component as a makeup exam.
// ReplacesComponentID points to the component this makeup replaces (nullable).
type ExamComponent struct {
	ID                  uuid.UUID  `gorm:"primaryKey" json:"id"`
	GradeWeightConfigID uuid.UUID  `gorm:"not null" json:"grade_weight_config_id"`
	Name                string     `gorm:"size:100;not null" json:"name"`
	Weight              float64    `gorm:"not null" json:"weight"`
	DisplayOrder        int        `gorm:"not null;default:0" json:"display_order"`
	IsMakeup            bool       `gorm:"not null;default:false" json:"is_makeup"`
	ReplacesComponentID *uuid.UUID `gorm:"default:null" json:"replaces_component_id"`

	GradeWeightConfig GradeWeightConfig `gorm:"foreignKey:GradeWeightConfigID" json:"grade_weight_config,omitempty"`
	BaseRecordFields
}

// Grade is the top-level grade record for a user's enrolled course.
// RawScore, LetterGrade and GradePoint are stored (not computed on the fly) for fast GPA queries.
// Scores for each exam component are stored as child ExamScore rows.
type Grade struct {
	ID                     uuid.UUID   `gorm:"primaryKey" json:"id"`
	UserCourseEnrollmentID uuid.UUID   `gorm:"not null;uniqueIndex" json:"user_course_enrollment_id"`
	RawScore               float64     `gorm:"not null;default:0" json:"raw_score"`
	LetterGrade            LetterGrade `gorm:"size:5;not null" json:"letter_grade"`
	GradePoint             float64     `gorm:"not null;default:0" json:"grade_point"`

	UserCourseEnrollment UserCourseEnrollment `gorm:"foreignKey:UserCourseEnrollmentID" json:"user_course_enrollment,omitempty"`
	Scores               []ExamScore          `gorm:"foreignKey:GradeID" json:"scores,omitempty"`
	BaseRecordFields
}

// ExamScore stores the score a user received for a specific ExamComponent.
// If the component IsMakeup and score is present, the service layer uses it
// instead of the replaced component's score when calculating RawScore.
type ExamScore struct {
	ID              uuid.UUID `gorm:"primaryKey" json:"id"`
	GradeID         uuid.UUID `gorm:"not null;uniqueIndex:idx_grade_component" json:"grade_id"`
	ExamComponentID uuid.UUID `gorm:"not null;uniqueIndex:idx_grade_component" json:"exam_component_id"`
	Score           *float64  `gorm:"default:null" json:"score"`

	Grade         Grade         `gorm:"foreignKey:GradeID" json:"grade,omitempty"`
	ExamComponent ExamComponent `gorm:"foreignKey:ExamComponentID" json:"exam_component,omitempty"`
	BaseRecordFields
}
