package mysql

import (
	"database/sql"
	"github.com/viamAhmadi/mars/pkg/models"
)

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *PostModel) Get(id int) (*models.Post, error) {
	return nil, nil
}

func (m *PostModel) Latest() ([]*models.Post, error) {
	return nil, nil
}
