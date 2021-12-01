package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kier1021/student-course-app/api/models"
	"github.com/kier1021/student-course-app/api/services"
)

type DeleteCourseParam struct {
	CourseID string `json:"course_id"`
}

type CourseController struct {
	srv *services.CourseService
}

func NewCourseController(srv *services.CourseService) *CourseController {
	return &CourseController{
		srv: srv,
	}
}

func (ctrl *CourseController) AddCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var course models.Course

		if err := c.Bind(&course); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "Cannot process the request body."})
			return
		}

		res, err := ctrl.srv.AddCourse(&course)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": fmt.Sprintf("Internal server error occured. %s", err.Error())})
			return
		}

		c.JSON(http.StatusCreated, res)
	}
}

func (ctrl *CourseController) DeleteCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var deleteParam DeleteCourseParam

		if err := c.Bind(&deleteParam); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "Cannot process the request body."})
			return
		}

		res, err := ctrl.srv.DeleteCourse(deleteParam.CourseID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": fmt.Sprintf("Internal server error occured. %s", err.Error())})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
