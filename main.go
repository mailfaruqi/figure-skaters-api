// package main

// import (
// 	"encoding/json"
// 	"figure-skaters-api/database"
// 	"figure-skaters-api/handlers"
// 	"figure-skaters-api/repositories"
// 	"figure-skaters-api/services"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strings"

// 	"github.com/spf13/viper"
// )

// type Config struct {
// Port   string `mapstructure:"PORT"`
// DBConn string `mapstructure:"DB_CONN"`
// }

// func main() {
// viper.AutomaticEnv()
// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

// if _, err := os.Stat(".env"); err == nil {
// viper.SetConfigFile(".env")
// _ = viper.ReadInConfig()
// }

// config := Config{
// Port:   viper.GetString("PORT"),
// DBConn: viper.GetString("DB_CONN"),
// }

// // Setup database
// db, err := database.InitDB(config.DBConn)
// if err != nil {
// log.Fatal("Failed to initialize database:", err)
// }
// defer db.Close()

// // Setup Category
// categoryRepo := repositories.NewCategoryRepository(db)
// categoryService := services.NewCategoryService(categoryRepo)
// categoryHandler := handlers.NewCategoryHandler(categoryService)

// // Setup Element
// elementRepo := repositories.NewElementRepository(db)
// elementService := services.NewElementService(elementRepo)
// elementHandler := handlers.NewElementHandler(elementService)

// // Setup routes for categories
// http.HandleFunc("/api/categories", categoryHandler.HandleCategories)
// http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)

// // Setup routes for elements
// http.HandleFunc("/api/elements", elementHandler.HandleElements)
// http.HandleFunc("/api/elements/", elementHandler.HandleElementByID)

// // Health check
// http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
// w.Header().Set("Content-Type", "application/json")
// json.NewEncoder(w).Encode(map[string]string{
// "status":  "OK",
// "message": "Figure Skating Elements API Running",
// })
// })

// fmt.Println("Server running at localhost:" + config.Port)
// err = http.ListenAndServe(":"+config.Port, nil)
// if err != nil {
// fmt.Println("failed running server")
// }
// }

package main

import (
	"encoding/json"
	"figure-skaters-api/database"
	"figure-skaters-api/handlers"
	"figure-skaters-api/repositories"
	"figure-skaters-api/services"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	// Setup Viper
	viper.AutomaticEnv()
	
	// Bind specific env vars
	viper.BindEnv("PORT")
	viper.BindEnv("DB_CONN")
	
	// Read .env if exists (for local development)
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	// Get config with fallback defaults
	port := viper.GetString("PORT")
	if port == "" {
		port = "8080"
	}
	
	dbConn := viper.GetString("DB_CONN")
	if dbConn == "" {
		log.Fatal("DB_CONN environment variable is required")
	}

	config := Config{
		Port:   port,
		DBConn: dbConn,
	}

	// Debug log (remove in production)
	log.Printf("PORT: %s\n", config.Port)
	log.Printf("DB_CONN: %s\n", maskConnectionString(config.DBConn))

	// Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Setup Category
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Setup Element
	elementRepo := repositories.NewElementRepository(db)
	elementService := services.NewElementService(elementRepo)
	elementHandler := handlers.NewElementHandler(elementService)

	// Setup routes for categories
	http.HandleFunc("/api/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)

	// Setup routes for elements
	http.HandleFunc("/api/elements", elementHandler.HandleElements)
	http.HandleFunc("/api/elements/", elementHandler.HandleElementByID)

	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Figure Skating Elements API Running",
		})
	})

	fmt.Println("Server running on port " + config.Port)
	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// Helper function to mask sensitive data in logs
func maskConnectionString(connStr string) string {
	if len(connStr) < 30 {
		return "***"
	}
	return connStr[:20] + "***" + connStr[len(connStr)-10:]
}
