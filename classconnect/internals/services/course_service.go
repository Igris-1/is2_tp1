package services

import (
	"github.com/Igris-1/classconnect/internals/models"
	"github.com/Igris-1/classconnect/internals/repositories"
)

// CourseService maneja la lógica de negocio
type CourseService struct {
	repo *repositories.CourseRepository
}

// NewCourseService crea un nuevo servicio de cursos
func NewCourseService(repo *repositories.CourseRepository) *CourseService {
	return &CourseService{repo: repo}
}

// CreateCourse crea un curso nuevo
func (s *CourseService) CreateCourse(title, description string) models.Course {
	course := models.Course{Title: title, Description: description}
	return s.repo.Save(course)
}

// GetAllCourses devuelve todos los cursos
func (s *CourseService) GetAllCourses() []models.Course {
	return s.repo.FindAll()
}

// GetCourseByID devuelve un curso por su ID
func (s *CourseService) GetCourseByID(id int) (models.Course, bool) {
	return s.repo.FindByID(id)
}

// DeleteCourse elimina un curso
func (s *CourseService) DeleteCourse(id int) bool {
	return s.repo.Delete(id)
}
