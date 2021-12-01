package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kier1021/student-course-app/api/models"
	"github.com/kier1021/student-course-app/api/services"
)

type StudentCourseIDParam struct {
	StudentID int      `json:"student_id"`
	CourseIDs []string `json:"course_ids"`
}

type StudentController struct {
	srv *services.StudentService
}

func NewStudentController(srv *services.StudentService) *StudentController {
	return &StudentController{
		srv: srv,
	}
}

func (ctrl *StudentController) AddStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var student models.Student

		if err := c.Bind(&student); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "Cannot process the request body."})
			return
		}

		res, err := ctrl.srv.AddStudent(&student)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": fmt.Sprintf("Internal server error occured. %s", err.Error())})
			return
		}

		c.JSON(http.StatusCreated, res)
	}
}

func (ctrl *StudentController) AddStudentCourses() gin.HandlerFunc {
	return func(c *gin.Context) {
		var idParam StudentCourseIDParam

		if err := c.Bind(&idParam); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "Cannot process the request body."})
			return
		}

		res, err := ctrl.srv.AddStudentCourses(idParam.StudentID, idParam.CourseIDs)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": fmt.Sprintf("Internal server error occured. %s", err.Error())})
			return
		}

		c.JSON(http.StatusCreated, res)
	}
}

func (ctrl *StudentController) GetStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := ctrl.srv.GetStudents()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": fmt.Sprintf("Internal server error occured. %s", err.Error())})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func (ctrl *StudentController) GetStudentByCourseID() gin.HandlerFunc {
	return func(c *gin.Context) {

		courseID, _ := c.GetQuery("course_id")

		res, err := ctrl.srv.GetStudentByCourseID(courseID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": fmt.Sprintf("Internal server error occured. %s", err.Error())})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
