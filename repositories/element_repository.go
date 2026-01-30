package repositories

import (
	"database/sql"
	"errors"
	"figure-skaters-api/models"
)

type ElementRepository struct {
db *sql.DB
}

func NewElementRepository(db *sql.DB) *ElementRepository {
return &ElementRepository{db: db}
}

func (repo *ElementRepository) GetAll() ([]models.Element, error) {
query := "SELECT id, name, code, category_id, base_value, difficulty_level, created_at FROM skating_elements ORDER BY id"
rows, err := repo.db.Query(query)
if err != nil {
return nil, err
}
defer rows.Close()

elements := make([]models.Element, 0)
for rows.Next() {
var e models.Element
err := rows.Scan(&e.ID, &e.Name, &e.Code, &e.CategoryID, &e.BaseValue, &e.DifficultyLevel, &e.CreatedAt)
if err != nil {
return nil, err
}
elements = append(elements, e)
}
return elements, nil
}

// GetByID with JOIN to get category name
func (repo *ElementRepository) GetByID(id int) (*models.ElementDetail, error) {
query := `
SELECT e.id, e.name, e.code, e.category_id, c.name as category_name, e.base_value, e.difficulty_level, e.created_at
FROM skating_elements e
LEFT JOIN categories c ON e.category_id = c.id
WHERE e.id = $1
`
var e models.ElementDetail
err := repo.db.QueryRow(query, id).Scan(
&e.ID, &e.Name, &e.Code, &e.CategoryID, &e.CategoryName, &e.BaseValue, &e.DifficultyLevel, &e.CreatedAt,
)
if err == sql.ErrNoRows {
return nil, errors.New("element not found")
}
if err != nil {
return nil, err
}
return &e, nil
}

func (repo *ElementRepository) Create(element *models.Element) error {
query := "INSERT INTO skating_elements (name, code, category_id, base_value, difficulty_level) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at"
err := repo.db.QueryRow(query, element.Name, element.Code, element.CategoryID, element.BaseValue, element.DifficultyLevel).Scan(&element.ID, &element.CreatedAt)
return err
}

func (repo *ElementRepository) Update(element *models.Element) error {
query := "UPDATE skating_elements SET name = $1, code = $2, category_id = $3, base_value = $4, difficulty_level = $5 WHERE id = $6"
result, err := repo.db.Exec(query, element.Name, element.Code, element.CategoryID, element.BaseValue, element.DifficultyLevel, element.ID)
if err != nil {
return err
}
rows, err := result.RowsAffected()
if err != nil {
return err
}
if rows == 0 {
return errors.New("element not found")
}
return nil
}

func (repo *ElementRepository) Delete(id int) error {
query := "DELETE FROM skating_elements WHERE id = $1"
result, err := repo.db.Exec(query, id)
if err != nil {
return err
}
rows, err := result.RowsAffected()
if err != nil {
return err
}
if rows == 0 {
return errors.New("element not found")
}
return nil
}