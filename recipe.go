package goRecipes

import (
	"time"

	"github.com/stephenfeagin/goRecipes/json"
)

// Recipe defines the data for a schema.org Recipe. It only includes fields that are commonly used
// in recipe pages.
type Recipe struct {
	Context            string
	Type               string
	MainEntityOfPage   string
	Name               string
	Image              *json.image
	DatePublished      time.Time
	Description        string
	PrepTime           string
	CookTime           string
	TotalTime          string
	RecipeYield        string
	RecipeIngredient   []string
	RecipeInstructions []*json.RecipeInstruction
	RecipeCategory     []string
	RecipeCuisine      []string
	Author             *json.Author
	AggregateRating    *json.AggregateRating
	Nutrition          *json.Nutrition
}

// processRawRecipe calls functions to fully process all of the *json.RawMessage fields in the
// rawRecipe struct to create a Recipe struct
func processRawRecipe(raw *json.rawRecipe) (*Recipe, error) {
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
