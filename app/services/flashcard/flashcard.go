package flashcard

import (
	"Flashcards/app/functions"
	"Flashcards/app/models"
	"Flashcards/app/server"
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

type Flashcard struct {
	validate *validator.Validate
}

func New() *Flashcard {
	return &Flashcard{
		validate: validator.New(),
	}
}

func (f *Flashcard) Create(in *models.FlashcardInput) (*models.Flashcard, error) {
	var flashcard models.Flashcard

	srv := server.GetServer()
	collection := srv.Database.Collection(flashcard.Collection())

	// Validate input
	err := f.validate.Struct(in)
	if err != nil {
		log.Error().Err(err).Msg("Invalid input")
		return nil, err
	}
	/*
		// Convert input to flashcard model
		flashcard.Question = input.Question
		flashcard.Answers = input.Answers
		flashcard.CorrectAnswer = input.CorrectAnswer
		flashcard.Tags = input.Tags
		flashcard.CreatedAt = time.Now()
		flashcard.CustomID = mongodb.NewUUID()

	*/

	err = functions.ConvertInputStructToDataStruct(in, &flashcard)
	if err != nil {
		log.Error().Err(err).Msg("Error converting input to struct")
		return nil, err
	}

	flashcard.CustomID = functions.NewUUID()
	flashcard.CreatedAt = time.Now()

	_, err = collection.InsertOne(context.TODO(), flashcard)
	if err != nil {
		log.Error().Err(err).Msg("Error inserting flashcard into database")
		return nil, err
	}

	return &flashcard, nil
}

func (f *Flashcard) GetByID(id string) (models.Flashcard, error) {
	var flashcard models.Flashcard

	srv := server.GetServer()
	collection := srv.Database.Collection(flashcard.Collection())

	filter := bson.M{"customID": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&flashcard)
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving flashcard")
		return flashcard, err
	}

	return flashcard, nil
}

// Update updates an existing flashcard
func (fs *Flashcard) Update(id string, input *models.FlashcardInput) error {
	// Validate input
	err := fs.validate.Struct(input)
	if err != nil {
		log.Error().Err(err).Msg("Invalid input")
		return err
	}

	// Find and update flashcard
	filter := bson.M{"customID": id}
	update := bson.M{
		"$set": bson.M{
			"answer":           input.Answer,
			"responses":        input.Responses,
			"numRightResponse": input.NumRightResponse,
			"tags":             input.Tags,
		},
	}

	srv := server.GetServer()
	collection := srv.Database.Collection("flashcards")
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Error().Err(err).Msg("Error updating flashcard")
		return err
	}

	return nil
}

func (fs *Flashcard) Search(category string) ([]models.Flashcard, error) {
	var flashcards []models.Flashcard

	srv := server.GetServer()
	collection := srv.Database.Collection("flashcards")

	filter := bson.M{"tags": category}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Error().Err(err).Msg("Error searching flashcards")
		return nil, err
	}

	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var flashcard models.Flashcard
		if err = cursor.Decode(&flashcard); err != nil {
			log.Error().Err(err).Msg("Error decoding flashcard")
			return nil, err
		}
		flashcards = append(flashcards, flashcard)
	}

	return flashcards, nil
}
