package interfaces

import (
	"github.com/gin-gonic/gin"
)

type UserService interface {
	Create(context *gin.Context)
	GetAll(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}
