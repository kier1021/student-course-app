package services

import (
	"fmt"

	"github.com/kier1021/student-course-app/api/models"
	"github.com/kier1021/student-course-app/api/repositories"
)

type CourseService struct {
	repo *repositories.CourseRepository
}

func NewCourseService(repo *repositories.CourseRepository) *CourseService {
	return &CourseService{
		repo: repo,
	}
}

func (srv *CourseService) AddCourse(course *models.Course) (resp interface{}, err error) {

	c, err := srv.repo.GetCourseByID(course.ID)
	if err != nil {
		return nil, err
	}

	if c != nil {
		return map[string]interface{}{
			"message": fmt.Sprintf("Course %s already exists", course.ID),
		}, nil
	}

	err = srv.repo.AddCourse(course)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "Course successfully added.",
	}, nil
}

func (srv *CourseService) DeleteCourse(courseID string) (resp interface{}, err error) {

	c, err := srv.repo.GetCourseByID(courseID)
	if err != nil {
		return nil, err
	}

	if c == nil {
		return map[string]interface{}{
			"message": fmt.Sprintf("Course %s not exists", courseID),
		}, nil
	}

	err = srv.repo.DeleteCourse(courseID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "Course successfully deleted.",
	}, nil
}
