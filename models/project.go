package models

type Project struct {
	Activitys    []Activity `json:"-" gorm:"ForeignKey:ProjectID"`
	ID           uint
	UserID       uint   `json:"userid"`
	Project_name string `json:"project_name"`
}
