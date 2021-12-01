package models

type Course struct {
	ID            string `json:"course_id"`
	Name          string `json:"course_name"`
	ProfessorName string `json:"professor_name"`
	Description   string `json:"description"`
}
