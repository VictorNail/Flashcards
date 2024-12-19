package flashcard

import (
	"Flashcards/app/models"
	"Flashcards/app/services/flashcard"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Flashcard struct {
	FlashcardService *flashcard.Flashcard
}

func New(flashcardService *flashcard.Flashcard) *Flashcard {
	return &Flashcard{
		FlashcardService: flashcardService,
	}
}

// Create handles the creation of a new flashcard
func (fc *Flashcard) Create(ctx *gin.Context) {
	var input models.FlashcardInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	flashcard, err := fc.FlashcardService.Create(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create flashcard"})
		return
	}

	ctx.JSON(http.StatusCreated, flashcard)
}

// Update handles updating an existing flashcard
func (fc *Flashcard) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var input models.FlashcardInput

	if err := c.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := fc.FlashcardService.Update(id, &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update flashcard"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Flashcard updated successfully"})
}

// Search handles searching flashcards by tags or categories
func (fc *Flashcard) Search(ctx *gin.Context) {
	category := ctx.Query("category")

	flashcards, err := fc.FlashcardService.Search(category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to search flashcards"})
		return
	}

	ctx.JSON(http.StatusOK, flashcards)
}

// GetByID handles retrieving a flashcard by its ID
func (fc *Flashcard) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	flashcard, err := fc.FlashcardService.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Flashcard not found"})
		return
	}

	ctx.JSON(http.StatusOK, flashcard)
}
