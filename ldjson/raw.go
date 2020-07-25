package ldjson

import (
	"encoding/json"

	kb "github.com/stephenfeagin/kitchenbox"
)

// rawRecipe represents the data from a schema.org Recipe type. It allows many non-string fields to
// be unmarshalled as json.RawMessage, which can then be further refined to create a fully
// processed Recipe struct
type rawRecipe struct {
	Context          string `json:"@context"`
	Type             string `json:"@type"`
	MainEntityOfPage string `json:"mainEntityOfPage"` // URL
	Name             string `json:"name"`

	// may be a @type:ImageObject, may be just a URL
	Image *json.RawMessage `json:"image"`

	// 2019-04-02T06:49:59.000Z -- may be incompatible with RFC3339 due to time/timezone
	DatePublished string `json:"datePublished"`
	Description   string `json:"description"`

	// PT20M -- 20 minutes
	PrepTime           string                  `json:"prepTime"`
	CookTime           string                  `json:"cookTime"`
	TotalTime          string                  `json:"totalTime"`
	RecipeYield        string                  `json:"recipeYield"`
	RecipeIngredient   []string                `json:"recipeIngredient"`
	RecipeInstructions []*kb.RecipeInstruction `json:"recipeInstructions"`

	// RecipeCategory and RecipeCuisine may be a single string or a string array
	RecipeCategory  *json.RawMessage `json:"recipeCategory"`
	RecipeCuisine   *json.RawMessage `json:"recipeCuisine"`
	Author          *kb.Author       `json:"author"`
	AggregateRating *json.RawMessage `json:"aggregateRating"`
	Nutrition       *json.RawMessage `json:"nutrition"`
}

// processRecipe converts a partially-processed rawRecipe into a fully processed, usable kb.Recipe
func processRecipe(raw *rawRecipe) (*kb.Recipe, error) {
	rec := &kb.Recipe{}

	// Copy over the already processed fields
	rec.Context = raw.Context
	rec.Type = raw.Type
	rec.MainEntityOfPage = raw.MainEntityOfPage
	rec.Name = raw.Name
	rec.Description = raw.Description
	rec.PrepTime = raw.PrepTime
	rec.CookTime = raw.CookTime
	rec.TotalTime = raw.TotalTime
	rec.RecipeYield = raw.RecipeYield
	rec.RecipeIngredient = raw.RecipeIngredient[:]
	rec.RecipeInstructions = raw.RecipeInstructions
	rec.Author = raw.Author

	// unmarshal raw fields
	img, err := unmarshalImage(raw.Image)
	if err != nil {
		return nil, err
	}
	rec.Image = img

	recipeCategory, err := unmarshalToStringSlice(raw.RecipeCategory)
	if err != nil {
		return nil, err
	}
	rec.RecipeCategory = recipeCategory

	recipeCuisine, err := unmarshalToStringSlice(raw.RecipeCuisine)
	if err != nil {
		return nil, err
	}
	rec.RecipeCuisine = recipeCuisine

	aggregateRating, err := unmarshalAggregateRating(raw.AggregateRating)
	if err != nil {
		return nil, err
	}
	rec.AggregateRating = aggregateRating

	nutrition, err := unmarshalNutrition(raw.Nutrition)
	if err != nil {
		return nil, err
	}
	rec.Nutrition = nutrition

	return rec, nil
}
