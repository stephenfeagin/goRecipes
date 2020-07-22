package recipe

import "encoding/json"

// recipeInstruction defines the data for a schema.org RecipeInstruction
type recipeInstruction struct {
	Type string `json:"@type"` // HowToStep
	Text string `json:"text"`
}

// author defines the data for a schema.org Recipe.Author -- it should be type Person
type author struct {
	Type string `json:"@type"` // Person
	Name string `json:"name"`
}

// rawAggregateRating is a more flexible representation of a schema.org AggregateRating object. It
// allows for number fields to be in string form.
type rawAggregateRating struct {
	Type         string           `json:"@type"` // AggregateRating
	RatingValue  *json.RawMessage `json:"ratingValue"`
	RatingCount  *json.RawMessage `json:"ratingCount"`
	ItemReviewed string           `json:"itemReviewed"` // should match allRecipes.Name
	BestRating   *json.RawMessage `json:"bestRating"`
	WorstRating  *json.RawMessage `json:"worstRating"`
}

// aggregateRating defines the data for a schema.org AggregateRating object. It only includes fields
// that are commonly used in the schema.org Recipe object.
type aggregateRating struct {
	Type         string  `json:"@type"` // AggregateRating
	RatingValue  float64 `json:"ratingValue"`
	RatingCount  int     `json:"ratingCount"`
	ItemReviewed string  `json:"itemReviewed"` // should match allRecipes.Name
	BestRating   int     `json:"bestRating"`
	WorstRating  int     `json:"worstRating"`
}

// nutrition defines the data for a schema.org NutritionInformation type
type nutrition struct {
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
