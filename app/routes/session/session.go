package session

import (
	controller "Flashcards/app/controllers/session"
	service "Flashcards/app/services/session"

	"github.com/gin-gonic/gin"
)

func SetupRouter(g *gin.Engine) {

	servicesSession := service.New()
	sessionController := controller.New(servicesSession)

	v1 := g.Group("/v1")
	{
		sessions := v1.Group("/sessions")
		{
			// Créer une session (nécessite studentID & catégorie)
			sessions.POST("", sessionController.Create)
			sessions.GET("/:id/state", sessionController.GetState) // Params : idSession
			sessions.POST("/:id/answer", sessionController.Answer) // Params : idSession, Body : idCard, numeroReponse
		}
	}
}
