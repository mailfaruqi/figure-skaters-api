package models

import "time"

// Category model
type Category struct {
ID          int       `json:"id"`
Name        string    `json:"name"`
Description string    `json:"description"`
CreatedAt   time.Time `json:"created_at"`
}

// Element model (skating_elements table)
type Element struct {
ID              int     `json:"id"`
Name            string  `json:"name"`
Code            string  `json:"code"`
CategoryID      int     `json:"category_id"`
BaseValue       float64 `json:"base_value"`
DifficultyLevel string  `json:"difficulty_level"`
CreatedAt       time.Time `json:"created_at"`
}

// ElementDetail model with category name (for JOIN)
type ElementDetail struct {
ID              int       `json:"id"`
Name            string    `json:"name"`
Code            string    `json:"code"`
CategoryID      int       `json:"category_id"`
CategoryName    string    `json:"category_name"`
BaseValue       float64   `json:"base_value"`
DifficultyLevel string    `json:"difficulty_level"`
CreatedAt       time.Time `json:"created_at"`
}