package models

type Student struct {
	ID    int    `json:"student_id"`
	Name  string `json:"student_name"`
	Email string `json:"student_email"`
	Phone string `json:"student_phone"`
}

type StudentCourseBinding struct {
	StudentID int `json:"student_id"`
	CourseID  int `json:"course_id"`
}

type StudentWithCourse struct {
	Student
	CourseName string
	Courses    []Course
}
