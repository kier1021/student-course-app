package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kier1021/student-course-app/api/controllers"
	"github.com/kier1021/student-course-app/api/repositories"
	"github.com/kier1021/student-course-app/api/services"
	"github.com/kier1021/student-course-app/databases"
)

type APIRoutes struct {
	r                 *gin.Engine
	courseController  *controllers.CourseController
	studentController *controllers.StudentController
}

func NewAPIRoutes(r *gin.Engine) (api *APIRoutes, err error) {

	db, err := databases.NewDBSchool()
	if err != nil {
		return nil, err
	}

	// Init Course Repo, Services and Controller
	courseRepo := repositories.NewCourseRepository(db)
	courseSrv := services.NewCourseService(courseRepo)
	courseController := controllers.NewCourseController(courseSrv)

	// Init Student Repo, Services and Controller
	studentRepo := repositories.NewStudentRepository(db)
	studentSrv := services.NewStudentService(studentRepo, courseRepo)
	studentController := controllers.NewStudentController(studentSrv)

	return &APIRoutes{
		r:                 r,
		courseController:  courseController,
		studentController: studentController,
	}, nil
}

func (api *APIRoutes) SetRoutes() {

	api.r.POST("/course", api.courseController.AddCourse())
	api.r.DELETE("/course", api.courseController.DeleteCourse())

	api.r.POST("/student", api.studentController.AddStudent())
	api.r.POST("/courses/student", api.studentController.AddStudentCourses())

	api.r.GET("/students", api.studentController.GetStudents())
	api.r.GET("/course/student", api.studentController.GetStudentByCourseID())
}
