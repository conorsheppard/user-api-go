package impl

import (
	i "github.com/conorsheppard/user-api-go/internal/api/service/interface"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthServiceImpl struct{}

func (h HealthServiceImpl) Health(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func NewHealthService() i.HealthService {
	return &HealthServiceImpl{}
}
