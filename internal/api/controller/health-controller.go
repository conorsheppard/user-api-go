package controller

import (
	i "github.com/conorsheppard/user-api-go/internal/api/service/interface"
	"github.com/gin-gonic/gin"
)

type HealthController struct {
	healthService i.HealthService
}

func (hc HealthController) SetUpHealthRoute(router *gin.Engine) {
	router.GET("/health", func(context *gin.Context) {
		hc.healthService.Health(context)
	})
}

func NewHealthController(healthService i.HealthService) *HealthController {
	return &HealthController{
		healthService: healthService,
	}
}
