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
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
Port   string `mapstructure:"PORT"`
DBConn string `mapstructure:"DB_CONN"`
}

func main() {
viper.AutomaticEnv()
viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

if _, err := os.Stat(".env"); err == nil {
viper.SetConfigFile(".env")
_ = viper.ReadInConfig()
}

config := Config{
Port:   viper.GetString("PORT"),
DBConn: viper.GetString("DB_CONN"),
}

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

fmt.Println("Server running at localhost:" + config.Port)
err = http.ListenAndServe(":"+config.Port, nil)
if err != nil {
fmt.Println("failed running server")
}
}