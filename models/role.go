package models

import "github.com/gofrs/uuid/v5"

// Role represents a role in the system (e.g., admin, user, moderator)
type Role struct {
	ID              uuid.UUID        `gorm:"primary_key" json:"id"`
	Name            string           `json:"name" gorm:"unique;not null"` // e.g., "admin", "user", "moderator"
	DisplayName     string           `json:"display_name"`                // e.g., "Administrator", "Regular User"
	Description     string           `json:"description"`                 // Role description
	Active          bool             `json:"active" gorm:"default:true"`
	RolePermissions []RolePermission `json:"role_permissions,omitempty" gorm:"foreignKey:RoleID"`
	BaseRecordFields
}

// Permission represents a specific permission (e.g., user.create, event.delete)
type Permission struct {
	ID          uuid.UUID `gorm:"primary_key" json:"id"`
	Name        string    `json:"name" gorm:"unique;not null"` // e.g., "user.create", "event.delete"
	Resource    string    `json:"resource"`                    // e.g., "user", "event", "reservation"
	Action      string    `json:"action"`                      // e.g., "create", "read", "update", "delete"
	Description string    `json:"description"`
	Active      bool      `json:"active" gorm:"default:true"`
	BaseRecordFields
}

// RolePermission represents the many-to-many relationship between roles and permissions
type RolePermission struct {
	ID           uuid.UUID  `gorm:"primary_key" json:"id"`
	RoleID       uuid.UUID  `json:"role_id"`
	Role         Role       `json:"role" gorm:"foreignKey:RoleID"`
	PermissionID uuid.UUID  `json:"permission_id"`
	Permission   Permission `json:"permission" gorm:"foreignKey:PermissionID"`
	BaseRecordFields
}

// UserRole represents the many-to-many relationship between users and roles
type UserRole struct {
	ID     uuid.UUID `gorm:"primary_key" json:"id"`
	UserID uuid.UUID `json:"user_id"`
	User   User      `json:"user" gorm:"foreignKey:UserID"`
	RoleID uuid.UUID `json:"role_id"`
	Role   Role      `json:"role" gorm:"foreignKey:RoleID"`
	Active bool      `json:"active" gorm:"default:true"`
	BaseRecordFields
}
