package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthController ...
type HealthController struct{}

// Status health check
func (h HealthController) Status(c *gin.Context) {
	c.Status(http.StatusOK)
}
