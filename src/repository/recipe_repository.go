package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/dbiagi/gororoba/src/domain"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

const (
	RecipeTable = "Recipe"
)

type RecipeRepository struct {
	*dynamodb.DynamoDB
}

func NewRecipeRepository(db *dynamodb.DynamoDB) RecipeRepository {
	return RecipeRepository{DynamoDB: db}
}

func (r *RecipeRepository) GetRecipesByCategory(category string) []domain.Recipe {
	return []domain.Recipe{}
}

func (r *RecipeRepository) CreateRecipe(recipe *domain.Recipe) *domain.Error {
	recipe.Id = uuid.New().String()
	recipe.CreatedAt = time.Now()
	recipe.UpdatedAt = time.Now()
	recipe.IdAndUpdatedAt = recipe.Id + "#" + recipe.UpdatedAt.Format(time.RFC3339)

	marshalledItem, marshallError := dynamodbattribute.MarshalMap(recipe)

	if marshallError != nil {
		slog.Error("Error marshalling recipe", slog.String("error", marshallError.Error()))
		return &domain.Error{
			Message: "Error marshalling recipe. Message: " + marshallError.Error(),
			Cause:   marshallError,
		}
	}
	_, putError := r.PutItem(&dynamodb.PutItemInput{
		TableName: getTableName(),
		Item:      marshalledItem,
	})

	if putError != nil {
		slog.Error("Error putting recipe", slog.String("error", putError.Error()))
		return &domain.Error{
			Message: "Error putting recipe. Message: " + putError.Error(),
			Cause:   putError,
		}
	}

	return nil
}

func getTableName() *string {
	return aws.String(RecipeTable)
}
