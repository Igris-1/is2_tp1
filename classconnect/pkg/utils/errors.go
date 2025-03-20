package utils

import "github.com/gin-gonic/gin"

// ErrorResponse estructura estándar de error RFC 7807
type ErrorResponse struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

// SendErrorResponse envía errores en formato RFC 7807
func SendErrorResponse(c *gin.Context, status int, title, detail string) {
	errorResponse := ErrorResponse{
		Type:     "about:blank",
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: c.Request.URL.Path,
	}

	c.JSON(status, errorResponse)
}
