package recipe

import (
	"encoding/json"
	"time"
)

// rawRecipe represents the data from a schema.org Recipe type. It allows many non-string fields to
// be unmarshalled as json.RawMessage, which can then be further refined to create a fully-processed
// Recipe struct
type rawRecipe struct {
	Context            string               `json:"@context"`
	Type               string               `json:"@type"`
	MainEntityOfPage   string               `json:"mainEntityOfPage"` // URL
	Name               string               `json:"name"`
	Image              *json.RawMessage     `json:"image"`         // may be a @type:ImageObject, may be just a URL
	DatePublished      string               `json:"datePublished"` // 2019-04-02T06:49:59.000Z -- may be compatible with time.Time, may not be
	Description        string               `json:"description"`
	PrepTime           string               `json:"prepTime"` // PT20M -- 20 minutes
	CookTime           string               `json:"cookTime"`
	TotalTime          string               `json:"totalTime"`
	RecipeYield        string               `json:"recipeYield"`
	RecipeIngredient   []string             `json:"recipeIngredient"`
	RecipeInstructions []*recipeInstruction `json:"recipeInstructions"`
	RecipeCategory     *json.RawMessage     `json:"recipeCategory"` // may be a single string, may be a []string
	RecipeCuisine      *json.RawMessage     `json:"recipeCuisine"`  // may be a single string, may be a []string
	Author             *author              `json:"author"`
	AggregateRating    *json.RawMessage     `json:"aggregateRating"`
	Nutrition          *json.RawMessage     `json:"nutrition"`
}

// Recipe defines the data for a schema.org Recipe. It only includes fields that are commonly used
// in recipe pages.
type Recipe struct {
	Context            string
	Type               string
	MainEntityOfPage   string
	Name               string
	Image              *image
	DatePublished      time.Time
	Description        string
	PrepTime           string
	CookTime           string
	TotalTime          string
	RecipeYield        string
	RecipeIngredient   []string
	RecipeInstructions []*recipeInstruction
	RecipeCategory     []string
	RecipeCuisine      []string
	Author             *author
	AggregateRating    *aggregateRating
	Nutrition          *nutrition
}

// processRawRecipe calls functions to fully process all of the *json.RawMessage fields in the
// rawRecipe struct to create a Recipe struct
func processRawRecipe(raw *rawRecipe) (*Recipe, error) {
	// the fields that need explicit conversion are Image, DatePublished, RecipeCategory,
	// RecipeCuisine, AggregateRating, and Nutrition
	recipe := &Recipe{}

	// First, copy over uncontroversial fields
	recipe.Context = raw.Context
	recipe.Type = raw.Type
	recipe.MainEntityOfPage = raw.MainEntityOfPage
	recipe.Name = raw.Name
	recipe.Description = raw.Description
	recipe.PrepTime = raw.PrepTime // these time fields may need further processing
	recipe.CookTime = raw.CookTime
	recipe.TotalTime = raw.TotalTime
	recipe.RecipeYield = raw.RecipeYield
	recipe.RecipeIngredient = raw.RecipeIngredient
	recipe.RecipeInstructions = raw.RecipeInstructions

	// Image
	img, err := processRawImageFromJSON(raw.Image)
	if err != nil {
		return nil, err
	}
	recipe.Image = img

	// DatePublished
	date, err := parseDatePublishedFromJSON(raw.DatePublished)
	if err != nil {
		return nil, err
	}
	recipe.DatePublished = date

	// RecipeCategory
	cats, err := parseStringSliceFromJSON(raw.RecipeCategory)
	if err != nil {
		return nil, err
	}
	recipe.RecipeCategory = cats

	// RecipeCuisine
	cuisine, err := parseStringSliceFromJSON(raw.RecipeCuisine)
	if err != nil {
		return nil, err
	}
	recipe.RecipeCuisine = cuisine

	// AggregateRating
	rating, err := processAggregateRatingFromJSON(raw.AggregateRating)
	if err != nil {
		return nil, err
	}
	recipe.AggregateRating = rating

	// Nutrition
	nutrition, err := processNutritionFromJSON(raw.Nutrition)
	if err != nil {
		return nil, err
	}
	recipe.Nutrition = nutrition

	return nil, nil
}
