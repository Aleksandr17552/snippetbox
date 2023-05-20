package main

import "github.com/Aleksandr17552/snippetbox/pkg/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
