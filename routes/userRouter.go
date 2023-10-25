package routes

import (
	"github.com/LeonardoMagnani/auth_go/controllers"
	"github.com/LeonardoMagnani/auth_go/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(incomingRouter *gin.Engine) {
	incomingRouter.Use(middleware.Authenticate())

	incomingRouter.GET("/users", controllers.GetUsers)
	incomingRouter.GET("/user/:user_id", controllers.GetUser)
	incomingRouter.POST("/user/verify", controllers.VerifyUser)
}
