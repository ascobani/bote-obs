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

type CourseEnrollmentStatus string

const (
	CourseEnrollmentStatusEnrolled CourseEnrollmentStatus = "Enrolled"
	CourseEnrollmentStatusDropped  CourseEnrollmentStatus = "Dropped"
	CourseEnrollmentStatusPassed   CourseEnrollmentStatus = "Passed"
	CourseEnrollmentStatusFailed   CourseEnrollmentStatus = "Failed"
)

type UserCourseEnrollment struct {
	ID                 uuid.UUID              `gorm:"primaryKey" json:"id"`
	UserID             uuid.UUID              `gorm:"not null;uniqueIndex:idx_user_course_semester" json:"user_id"`
	DepartmentCourseID uuid.UUID              `gorm:"not null;uniqueIndex:idx_user_course_semester" json:"department_course_id"`
	TimetableID        uuid.UUID              `gorm:"not null" json:"timetable_id"`
	Semester           int                    `gorm:"not null;uniqueIndex:idx_user_course_semester" json:"semester"`
	Year               int                    `gorm:"not null" json:"year"`
	Status             CourseEnrollmentStatus `gorm:"size:20;not null;default:'Enrolled'" json:"status"`

	User             User             `gorm:"foreignKey:UserID" json:"user,omitempty"`
	DepartmentCourse DepartmentCourse `gorm:"foreignKey:DepartmentCourseID" json:"department_course,omitempty"`
	Timetable        Timetable        `gorm:"foreignKey:TimetableID" json:"timetable,omitempty"`
	BaseRecordFields
}
