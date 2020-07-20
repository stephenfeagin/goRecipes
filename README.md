# goRecipes

This WIP application aims to eventually parse, store, organize, and display recipes.

The component that I'm working on right now is the "parse" bit. The goal is to be able to
turn an HTML file into a defined Recipe struct. I do this using 
[goquery](https://github.com/PuerKitoBio/goquery) to extract recipe data, currently only
supporting JSON-LD that implements the schema.org/Recipe structure. I am still in the
process of completing this step. Immediate next steps include parsing schema.org/Recipe
data in HTML/XML.

