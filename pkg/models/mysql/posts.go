package mysql

import (
	"database/sql"
	"github.com/viamAhmadi/mars/pkg/models"
)

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO posts (title, content, created, expires) VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *PostModel) Get(id int) (*models.Post, error) {
	p := &models.Post{}
	row := m.DB.QueryRow(`SELECT id, title, content, created, expires FROM posts WHERE expires > UTC_TIMESTAMP() AND id=?`, id)
	err := row.Scan(&p.ID, &p.Title, &p.Content, &p.Created, &p.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return p, nil
}

func (m *PostModel) Latest() ([]*models.Post, error) {
	stmt := `SELECT id, title, content, created, expires FROM posts WHERE expires>UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []*models.Post

	for rows.Next() {
		p := &models.Post{}
		err = rows.Scan(&p.ID, &p.Title, &p.Content, &p.Created, &p.Expires)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}
