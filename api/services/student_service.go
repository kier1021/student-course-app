package services

import (
	"fmt"
	"strings"

	"github.com/kier1021/student-course-app/api/models"
	"github.com/kier1021/student-course-app/api/repositories"
)

type StudentService struct {
	repo       *repositories.StudentRepository
	courseRepo *repositories.CourseRepository
}

func NewStudentService(repo *repositories.StudentRepository, courseRepo *repositories.CourseRepository) *StudentService {
	return &StudentService{
		repo:       repo,
		courseRepo: courseRepo,
	}
}

func (srv *StudentService) AddStudent(student *models.Student) (resp interface{}, err error) {

	lastInsertID, err := srv.repo.AddStudent(student)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message":      "Student successfully added.",
		"student_id":   lastInsertID,
		"student_name": student.Name,
	}, nil
}

func (srv *StudentService) AddStudentCourses(studentID int, courseID []string) (resp interface{}, err error) {
	student, err := srv.repo.GetStudentByID(studentID)
	if err != nil {
		return nil, err
	}

	if student == nil {
		return map[string]interface{}{
			"message": "Student not exists",
		}, nil
	}

	for _, id := range courseID {
		course, err := srv.courseRepo.GetCourseByID(id)
		if err != nil {
			return nil, err
		}

		if course == nil {
			return map[string]interface{}{
				"message": fmt.Sprintf("Course %s not exists", id),
			}, nil
		}

		studCourse, err := srv.repo.GetStudentsByStudentAndCourseID(studentID, id)
		if err != nil {
			return nil, err
		}

		if len(studCourse) != 0 {
			return map[string]interface{}{
				"message": fmt.Sprintf("Course %s already added to student with ID %d", id, studentID),
			}, nil
		}
	}

	err = srv.repo.AddStudentCourses(studentID, courseID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "Courses successully added to student",
	}, nil
}

func (srv *StudentService) GetStudents() (students []map[string]interface{}, err error) {
	results, err := srv.repo.GetStudents()
	if err != nil {
		return nil, err
	}

	studentMap := map[int]map[string]interface{}{}
	for _, result := range results {
		if student, ok := studentMap[result.ID]; ok {
			studentCourse, _ := student["courses"].(string)
			studentCourses := strings.Split(studentCourse, ", ")

			if result.CourseName != "" {
				student["courses"] = strings.Join(append(studentCourses, result.CourseName), ", ")
			}
		} else {
			resMap := make(map[string]interface{})
			resMap["student_id"] = result.ID
			resMap["name"] = result.Name
			resMap["email"] = result.Email
			resMap["phone_number"] = result.Phone
			resMap["courses"] = result.CourseName
			studentMap[result.ID] = resMap
		}
	}

	resList := []map[string]interface{}{}

	for _, student := range studentMap {
		resList = append(resList, student)
	}

	return resList, nil
}

func (srv *StudentService) GetStudentByCourseID(courseID string) (students []map[string]interface{}, err error) {

	results, err := srv.repo.GetStudentsByCourseID(courseID)
	if err != nil {
		return nil, err
	}

	resList := []map[string]interface{}{}
	for _, result := range results {

		resMap := make(map[string]interface{})
		resMap["student_id"] = result.ID
		resMap["name"] = result.Name
		resMap["email"] = result.Email
		resMap["phone_number"] = result.Phone

		resList = append(resList, resMap)
	}

	return resList, nil

}
