package interfaces

import "github.com/gin-gonic/gin"

type HealthService interface {
	Health(context *gin.Context)
}
