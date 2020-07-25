package ldjson

import (
	"encoding/json"

	kb "github.com/stephenfeagin/kitchenbox"
)

// rawRecipe represents the data from a schema.org Recipe type. It allows many non-string fields to
// be unmarshalled as json.RawMessage, which can then be further refined to create a fully-processed
// Recipe struct
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
