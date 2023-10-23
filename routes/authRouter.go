package routes

import (
	controller "github.com/LeonardoMagnani/auth_go/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(incomingRouter *gin.Engine) {
	incomingRouter.POST("/users/signup", controller.Signup)
	incomingRouter.POST("/users/login", controller.Login)
}
