package models

type Path struct {
	ID          uint      `gorm:"primary_key"`
	Path        string    `gorm:"unique" json:"path"`
	WorkspaceID Workspace `json:"workspace_id"`
	FormID      Form      `json:"form_id"`
}
