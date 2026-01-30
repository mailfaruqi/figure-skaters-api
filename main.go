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
// )

// type Config struct {
// 	Port   string
// 	DBConn string
// }

// func getEnv(key, fallback string) string {
// 	if value := os.Getenv(key); value != "" {
// 		return value
// 	}
// 	return fallback
// }

// func loadConfig() Config {
// 	dbConn := os.Getenv("DB_CONN")
// 	port := os.Getenv("PORT")

// 	if dbConn == "" || port == "" {
// 		log.Println("Environment variables not found, trying .env file...")
// 		if err := loadEnvFile(".env"); err == nil {
// 			if dbConn == "" {
// 				dbConn = os.Getenv("DB_CONN")
// 			}
// 			if port == "" {
// 				port = os.Getenv("PORT")
// 			}
// 		}
// 	}

// 	if port == "" {
// 		port = "8080"
// 	}

// 	if dbConn == "" {
// 		log.Fatal("❌ DB_CONN environment variable is not set!")
// 	}

// 	log.Printf("✅ PORT: %s", port)
// 	log.Printf("✅ DB_CONN: %s", maskConnectionString(dbConn))

// 	return Config{
// 		Port:   port,
// 		DBConn: dbConn,
// 	}
// }

// func loadEnvFile(filename string) error {
// 	data, err := os.ReadFile(filename)
// 	if err != nil {
// 		return err
// 	}

// 	lines := strings.Split(string(data), "\n")
// 	for _, line := range lines {
// 		line = strings.TrimSpace(line)
// 		if line == "" || strings.HasPrefix(line, "#") {
// 			continue
// 		}
// 		parts := strings.SplitN(line, "=", 2)
// 		if len(parts) == 2 {
// 			key := strings.TrimSpace(parts[0])
// 			value := strings.TrimSpace(parts[1])
// 			os.Setenv(key, value)
// 		}
// 	}

// 	return nil
// }

// func maskConnectionString(connStr string) string {
// 	if len(connStr) < 30 {
// 		return "[HIDDEN]"
// 	}
// 	return connStr[:25] + "***" + connStr[len(connStr)-15:]
// }

// func main() {
// 	log.Println("🚀 Starting Figure Skating API...")

// 	config := loadConfig()

// 	log.Println("📦 Connecting to database...")
// 	db, err := database.InitDB(config.DBConn)
// 	if err != nil {
// 		log.Fatalf("❌ Failed to initialize database: %v", err)
// 	}
// 	defer db.Close()
// 	log.Println("✅ Database connected successfully")

// 	categoryRepo := repositories.NewCategoryRepository(db)
// 	elementRepo := repositories.NewElementRepository(db)

// 	categoryService := services.NewCategoryService(categoryRepo)
// 	elementService := services.NewElementService(elementRepo)

// 	categoryHandler := handlers.NewCategoryHandler(categoryService)
// 	elementHandler := handlers.NewElementHandler(elementService)

// 	http.HandleFunc("/api/categories", categoryHandler.HandleCategories)
// 	http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)
// 	http.HandleFunc("/api/elements", elementHandler.HandleElements)
// 	http.HandleFunc("/api/elements/", elementHandler.HandleElementByID)

// 	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(map[string]string{
// 			"status":  "OK",
// 			"message": "Figure Skating Elements API Running",
// 		})
// 	})

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(map[string]string{
// 			"message": "Welcome to Figure Skating API",
// 			"version": "1.0.0",
// 		})
// 	})

// 	log.Printf("✅ Server starting on port %s", config.Port)
// 	fmt.Printf("🌐 Server running on port %s\n", config.Port)

// 	if err := http.ListenAndServe("0.0.0.0:"+config.Port, nil); err != nil {
// 		log.Fatalf("❌ Failed to start server: %v", err)
// 	}
// }
