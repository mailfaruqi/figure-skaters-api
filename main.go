package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Athlete represents a figure skating athlete
type Athlete struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Country     string `json:"country"`
	Discipline  string `json:"discipline"` // e.g., "Men's Sin`gles", "Pairs", "Ice Dance"
	YearsActive string `json:"years_active"`
}

// Category represents skating categories
type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// In-memory data storage
var athletes = []Athlete{
	{ID: 1, Name: "Yuzuru Hanyu", Country: "Japan", Discipline: "Men's Singles", YearsActive: "2010-2022"},
	{ID: 2, Name: "Kim Yuna", Country: "South Korea", Discipline: "Women's Singles", YearsActive: "2006-2014"},
}

var categories = []Category{
	{ID: 1, Name: "Men's Singles", Description: "Individual male figure skaters"},
	{ID: 2, Name: "Women's Singles", Description: "Individual female figure skaters"},
	{ID: 3, Name: "Pairs", Description: "Two skaters performing together"},
	{ID: 4, Name: "Ice Dance", Description: "Couple performing dance movements on ice"},
}

// ==================== ATHLETE HANDLERS ====================

func athleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if !strings.HasPrefix(r.URL.Path, "/api/athletes") {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// GET /api/athletes → list all
		if r.URL.Path == "/api/athletes" {
			json.NewEncoder(w).Encode(athletes)
			return
		}
		// GET /api/athletes/{id} → get by ID
		idStr := strings.TrimPrefix(r.URL.Path, "/api/athletes/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid athlete ID", http.StatusBadRequest)
			return
		}
		for _, a := range athletes {
			if a.ID == id {
				json.NewEncoder(w).Encode(a)
				return
			}
		}
		http.Error(w, "Athlete not found", http.StatusNotFound)

	case http.MethodPost:
		// POST → create new athlete
		var newAthlete Athlete
		err := json.NewDecoder(r.Body).Decode(&newAthlete)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if newAthlete.Name == "" {
			http.Error(w, "Athlete name is required", http.StatusBadRequest)
			return
		}
		newAthlete.ID = len(athletes) + 1
		athletes = append(athletes, newAthlete)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newAthlete)

	case http.MethodDelete:
		// DELETE /api/athletes/{id}
		idStr := strings.TrimPrefix(r.URL.Path, "/api/athletes/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid athlete ID", http.StatusBadRequest)
			return
		}
		for i, a := range athletes {
			if a.ID == id {
				athletes = append(athletes[:i], athletes[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		http.Error(w, "Athlete not found", http.StatusNotFound)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ==================== CATEGORY HANDLERS ====================

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if !strings.HasPrefix(r.URL.Path, "/api/categories") {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// GET /api/categories → list all
		if r.URL.Path == "/api/categories" {
			json.NewEncoder(w).Encode(categories)
			return
		}
		// GET /api/categories/{id} → get by ID
		idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}
		for _, c := range categories {
			if c.ID == id {
				json.NewEncoder(w).Encode(c)
				return
			}
		}
		http.Error(w, "Category not found", http.StatusNotFound)

	case http.MethodPost:
		// POST → create new category
		var newCategory Category
		err := json.NewDecoder(r.Body).Decode(&newCategory)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if newCategory.Name == "" {
			http.Error(w, "Category name is required", http.StatusBadRequest)
			return
		}
		newCategory.ID = len(categories) + 1
		categories = append(categories, newCategory)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newCategory)

	case http.MethodPut:
		// PUT /api/categories/{id} → update category
		idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}
		
		var updatedCategory Category
		err = json.NewDecoder(r.Body).Decode(&updatedCategory)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		for i, c := range categories {
			if c.ID == id {
				// Keep the same ID, update other fields
				updatedCategory.ID = id
				categories[i] = updatedCategory
				json.NewEncoder(w).Encode(updatedCategory)
				return
			}
		}
		http.Error(w, "Category not found", http.StatusNotFound)

	case http.MethodDelete:
		// DELETE /api/categories/{id}
		idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}
		for i, c := range categories {
			if c.ID == id {
				categories = append(categories[:i], categories[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		http.Error(w, "Category not found", http.StatusNotFound)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ==================== HEALTH CHECK ====================

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "OK",
		"message": "Figure Skating API Running",
	})
}

// ==================== MAIN ====================

func main() {
	// Athletes endpoints
	http.HandleFunc("/api/athletes", athleteHandler)
	http.HandleFunc("/api/athletes/", athleteHandler)

	// Categories endpoints
	http.HandleFunc("/api/categories", categoryHandler)
	http.HandleFunc("/api/categories/", categoryHandler)

	// Health check
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Figure Skating API running at http://localhost:8080")
	fmt.Println("Endpoints:")
	fmt.Println("   Athletes: /api/athletes")
	fmt.Println("   Categories: /api/categories")
	fmt.Println("   Health: /health")
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}