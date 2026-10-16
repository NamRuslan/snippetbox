package main

import "snippetbox.namruslan.com/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
