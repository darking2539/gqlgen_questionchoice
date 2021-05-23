package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	database "github.com/darking2539/gqlgen/db"
	"github.com/darking2539/gqlgen/graph/generated"
	"github.com/darking2539/gqlgen/graph/model"
	"github.com/rs/xid"
)

func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.QuestionInput) (*model.Question, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}

	fmt.Println("input", input.QuestionText, input.PubDate)
	question := model.Question{}
	question.QuestionText = input.QuestionText
	question.PubDate = input.PubDate
	question.ID = xid.New().String()
	db.Create(&question)
	return &question, nil
}

func (r *mutationResolver) CreateChoice(ctx context.Context, input *model.ChoiceInput) (*model.Choice, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	
	fmt.Println("input", input.QuestionID, input.ChoiceText)
	choice := model.Choice{}
	question := model.Question{}
	choice.QuestionID = input.QuestionID
	choice.ChoiceText = input.ChoiceText
	choice.ID = xid.New().String()

	
	
	db.Where("id = ?", choice.QuestionID).First(&question)
	//fmt.Println("&question ID = ", question.ID)
	if question.ID != choice.QuestionID {
		log.Println("Question ID dosen't math from DB", err)
		err := fmt.Errorf("Question ID dosen't math from DB")
		return nil, err
	}
		
	choice.Question = &question
	db.Select("id", "question_id", "choice_text").Create(&choice)

	return &choice, nil
}

func (r *queryResolver) Questions(ctx context.Context) ([]*model.Question, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}	
	db.Find(&r.questions)
	
	for _, question := range r.questions {
		var choices []*model.Choice
		db.Where(&model.Choice{QuestionID: question.ID}).Find(&choices)
		question.Choices = choices
	}

	return r.questions, nil
}

func (r *queryResolver) Questionq(ctx context.Context, input model.QuestionQuery) ([]*model.Question, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}

	question := model.Question{}
	question.QuestionText = input.QuestionText
	
	db.Where("question_text LIKE ?", question.QuestionText).Find(&r.questions)

	for _, question := range r.questions {
		var choices []*model.Choice
		db.Where(&model.Choice{QuestionID: question.ID}).Find(&choices)
		question.Choices = choices
	}

	return r.questions, nil
}

func (r *queryResolver) Choices(ctx context.Context) ([]*model.Choice, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
