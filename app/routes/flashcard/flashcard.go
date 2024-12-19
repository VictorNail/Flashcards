package flashcard

import (
	controller "Flashcards/app/controllers/flashcard"
	service "Flashcards/app/services/flashcard"

	"github.com/gin-gonic/gin"
)

func SetupRouter(g *gin.Engine) {

	servicesFlashcard := service.New()
	flashcardController := controller.New(servicesFlashcard)

	v1 := g.Group("/v1")
	{
		flashcards := v1.Group("/flashcards")
		{
			flashcards.POST("", flashcardController.Create)     // Créer une flashcard
			flashcards.POST("/:id", flashcardController.Update) // Mettre à jour une flashcard
			flashcards.GET("", flashcardController.Search)      // Rechercher des flashcards
			flashcards.GET("/:id", flashcardController.GetByID) // Récupérer une flashcard par ID
		}
	}
}
