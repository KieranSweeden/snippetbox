package main

import "snippetbox.kieransweeden.dev/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
