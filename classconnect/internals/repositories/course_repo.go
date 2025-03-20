package repositories

import (
	"github.com/Igris-1/classconnect/internals/models"
	"sync"
)

// CourseRepository maneja la persistencia en memoria
type CourseRepository struct {
	mu      sync.Mutex
	courses map[int]models.Course
	nextID  int
}

// NewCourseRepository crea una nueva instancia del repositorio
func NewCourseRepository() *CourseRepository {
	return &CourseRepository{
		courses: make(map[int]models.Course),
		nextID:  1,
	}
}

// Save guarda un nuevo curso
func (r *CourseRepository) Save(course models.Course) models.Course {
	r.mu.Lock()
	defer r.mu.Unlock()

	course.ID = r.nextID
	r.courses[r.nextID] = course
	r.nextID++

	return course
}

// FindAll retorna todos los cursos
func (r *CourseRepository) FindAll() []models.Course {
	r.mu.Lock()
	defer r.mu.Unlock()

	var courses []models.Course
	for _, course := range r.courses {
		courses = append(courses, course)
	}

	return courses
}

// FindByID busca un curso por ID
func (r *CourseRepository) FindByID(id int) (models.Course, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	course, exists := r.courses[id]
	return course, exists
}

// Delete elimina un curso por ID
func (r *CourseRepository) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.courses[id]; !exists {
		return false
	}
	delete(r.courses, id)
	return true
}
