package main

import "github.com/viamAhmadi/mars/pkg/models"

type templateData struct {
	Post  *models.Post
	Posts []*models.Post
}
