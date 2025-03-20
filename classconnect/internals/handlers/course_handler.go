package handlers

import (
	"github.com/Igris-1/classconnect/internals/services"
	"github.com/Igris-1/classconnect/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CourseHandler maneja las solicitudes HTTP
type CourseHandler struct {
	service *services.CourseService
}

// NewCourseHandler crea un nuevo handler
func NewCourseHandler(service *services.CourseService) *CourseHandler {
	return &CourseHandler{service: service}
}

// CreateCourse maneja la creación de cursos
func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid Request", "The request body is invalid or missing fields.")
		return
	}

	course := h.service.CreateCourse(req.Title, req.Description)
	c.JSON(http.StatusCreated, gin.H{"data": course})
}

// GetAllCourses devuelve todos los cursos
func (h *CourseHandler) GetAllCourses(c *gin.Context) {
	courses := h.service.GetAllCourses()
	c.JSON(http.StatusOK, gin.H{"data": courses})
}

// GetCourseByID obtiene un curso específico
func (h *CourseHandler) GetCourseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid ID", "The course ID must be an integer.")
		return
	}

	course, found := h.service.GetCourseByID(id)
	if !found {
		utils.SendErrorResponse(c, http.StatusNotFound, "Course Not Found", "The requested course does not exist.")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": course})
}

// DeleteCourse elimina un curso
func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid ID", "The course ID must be an integer.")
		return
	}

	if !h.service.DeleteCourse(id) {
		utils.SendErrorResponse(c, http.StatusNotFound, "Course Not Found", "The requested course does not exist.")
		return
	}

	c.Status(http.StatusNoContent)
}
