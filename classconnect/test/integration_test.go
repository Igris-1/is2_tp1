package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Igris-1/classconnect/internals/handlers"
	"github.com/Igris-1/classconnect/internals/repositories"
	"github.com/Igris-1/classconnect/internals/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setupRouter configura el router para las pruebas
func setupRouter() *gin.Engine {
	// Inicializa dependencias
	repo := repositories.NewCourseRepository()
	service := services.NewCourseService(repo)
	handler := handlers.NewCourseHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/courses", handler.CreateCourse)
	r.GET("/courses", handler.GetAllCourses)
	r.GET("/courses/:id", handler.GetCourseByID)
	r.DELETE("/courses/:id", handler.DeleteCourse)

	return r
}

// TestCreateCourse verifica que se pueda crear un curso correctamente
func TestCreateCourse(t *testing.T) {
	r := setupRouter()

	// Simula una solicitud POST con JSON
	course := map[string]string{"title": "Go Basics", "description": "Learn Golang"}
	jsonValue, _ := json.Marshal(course)
	req, _ := http.NewRequest("POST", "/courses", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Ejecuta la solicitud en un servidor de prueba
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verifica la respuesta
	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Go Basics", response["data"].(map[string]interface{})["title"])
}

// TestGetAllCourses verifica que se puedan recuperar todos los cursos
func TestGetAllCourses(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/courses", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCourseByID(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/courses/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// StatusNoContent
func TestDeleteCourse(t *testing.T) {
	r := setupRouter()
	createReq, _ := http.NewRequest("POST", "/courses", bytes.NewBuffer([]byte(`{"id": 1, "name": "Test Course"}`)))
	createReq.Header.Set("Content-Type", "application/json")
	createW := httptest.NewRecorder()
	r.ServeHTTP(createW, createReq)

	// Ahora intenta eliminarlo
	req, _ := http.NewRequest("DELETE", "/courses/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

}

// TODO: Agregar mas tests
