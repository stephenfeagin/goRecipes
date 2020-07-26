![Go](https://github.com/stephenfeagin/goRecipes/workflows/Go/badge.svg)

# goRecipes

This WIP application aims to eventually parse, store, organize, display, and export recipes.

The component that I'm working on right now is the "parse" bit. The goal is to be able to
turn an HTML file into a defined Recipe struct. I do this using 
[goquery](https://github.com/PuerKitoBio/goquery) to extract recipe data, currently only
supporting JSON-LD that implements the schema.org/Recipe structure. I am still in the
process of completing this step. Immediate next steps include parsing schema.org/Recipe
data in HTML microdata.

## TODO

- [ ] Parse `rawRecipe` into `Recipe` from JSON-LD  
    - [ ] Create valid `image` struct
	- Still need to accommodate multiple image URLs
    - [x] Create valid `aggregateRating` struct
	- Involves accommodating number fields input as either strings or JSON numbers
    - [x] Create valid `nutrition` struct
	- Involves accommodating JSON object (empty or not) or an empty string or array
    - [x] Create valid `DatePublished` field
	- Parse into a `time.Time`
	- May or may not have TZ info, which may have to be inferred/imputed
    - [x] Create valid `RecipeCategory` field
	- May be input as a string or as a JSON array
    - [x] Create valid `RecipeCuisine` field
	- May be input as a string or as a JSON array
- [ ] Parse `rawRecipe` into `Recipe` from HTML microdata

