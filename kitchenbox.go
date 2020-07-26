package kitchenbox

import (
	"time"
)

// Recipe defines the data for a schema.org Recipe. It only includes fields that are commonly used
// in recipe pages.
type Recipe struct {
	Context            string              `json:"context"`
	Type               string              `json:"@type"`
	MainEntityOfPage   string              `json:"mainEntityOfPage"`
	Name               string              `json:"name"`
	Image              Image               `json:"image"`
	DatePublished      time.Time           `json:"datePublished"`
	Description        string              `json:"description"`
	PrepTime           string              `json:"prepTime"`
	CookTime           string              `json:"cookTime"`
	TotalTime          string              `json:"totalTime"`
	RecipeYield        string              `json:"recipeYield"`
	RecipeIngredient   []string            `json:"recipeIngredient"`
	RecipeInstructions []RecipeInstruction `json:"recipeInstruction"`
	RecipeCategory     []string            `json:"recipeCategory"`
	RecipeCuisine      []string            `json:"recipeCuisine"`
	Author             Author              `json:"author"`
	AggregateRating    AggregateRating     `json:"aggregateRating"`
	Nutrition          Nutrition           `json:"nutrition"`
}

// Nutrition defines the data for a schema.org NutritionInformation type
type Nutrition struct {
	Type                  string `json:"@type"`               // NutritionInformation
	Calories              string `json:"calories"`            // "233 calories"
	CarbohydrateContent   string `json:"carbohydrateContent"` // "2.4 g"
	CholesterolContent    string `json:"cholesterolContent"`
	FatContent            string `json:"fatContent"`
	FiberContent          string `json:"fiberContent"`
	ProteinContent        string `json:"proteinContent"`
	SaturatedFatContent   string `json:"saturatedFatContent"`
	ServingSize           string `json:"servingSize"`
	SodiumContent         string `json:"sodiumContent"`
	SugarContent          string `json:"sugarContent"`
	TransFatContent       string `json:"transFatContent"`
	UnsaturatedFatContent string `json:"unsaturatedFatContent"`
}

// AggregateRating defines the data for a schema.org AggregateRating object. It only includes fields
// that are commonly used in the schema.org Recipe object.
type AggregateRating struct {
	Type         string  `json:"@type"` // AggregateRating
	RatingValue  float64 `json:"ratingValue"`
	RatingCount  int     `json:"ratingCount"`
	ItemReviewed string  `json:"itemReviewed"` // should match allRecipes.Name
	BestRating   int     `json:"bestRating"`
	WorstRating  int     `json:"worstRating"`
}

// Image defines the data for a schema.org Image. It only includes fields that are frequently
// included in Recipe models
type Image struct {
	Type    string `json:"@type"` // should be ImageObject
	URL     string `json:"url"`
	Width   string `json:"width"`
	Height  string `json:"height"`
	Caption string `json:"caption"`
}

// recipeInstruction defines the data for a schema.org RecipeInstruction
type RecipeInstruction struct {
	Type string `json:"@type"` // HowToStep
	Text string `json:"text"`
}

// author defines the data for a schema.org Recipe.Author -- it should be type Person
type Author struct {
	Type string `json:"@type"` // Person
	Name string `json:"name"`
}
