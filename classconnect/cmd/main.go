package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Igris-1/classconnect/config"
	"github.com/Igris-1/classconnect/internals/handlers"
	"github.com/Igris-1/classconnect/internals/repositories"
	"github.com/Igris-1/classconnect/internals/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar configuración
	config.LoadEnvVariables()

	// Crear repositorio y servicio
	repo := repositories.NewCourseRepository()
	service := services.NewCourseService(repo)
	handler := handlers.NewCourseHandler(service)

	// Configurar router con Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Definir endpoints
	r.POST("/courses", handler.CreateCourse)
	r.GET("/courses", handler.GetAllCourses)
	r.GET("/courses/:id", handler.GetCourseByID)
	r.DELETE("/courses/:id", handler.DeleteCourse)

	// Obtener puerto de variables de entorno
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor por defecto si no se especifica
	}

	// Iniciar servidor
	log.Printf("Server started and running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
