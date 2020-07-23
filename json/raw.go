package json

import "encoding/json"

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
	RecipeInstructions []*RecipeInstruction `json:"recipeInstructions"`
	RecipeCategory     *json.RawMessage     `json:"recipeCategory"` // may be a single string, may be a []string
	RecipeCuisine      *json.RawMessage     `json:"recipeCuisine"`  // may be a single string, may be a []string
	Author             *Author              `json:"author"`
	AggregateRating    *json.RawMessage     `json:"aggregateRating"`
	Nutrition          *json.RawMessage     `json:"nutrition"`
}
