package json

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
