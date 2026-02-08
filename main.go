package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main(){
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
 		Port: viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Root endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Selamat datang di Kasir API",
			"version": "1.0.0",
			"documentation": "/docs/swagger.json",
			"endpoints": map[string]interface{}{
				"health": "/health",
				"products": map[string]string{
					"list":   "GET /api/produk",
					"search": "GET /api/produk?name={keyword}",
					"create": "POST /api/produk",
					"detail": "GET /api/produk/{id}",
					"update": "PUT /api/produk/{id}",
					"delete": "DELETE /api/produk/{id}",
				},
				"categories": map[string]string{
					"list":   "GET /api/categories",
					"create": "POST /api/categories",
					"detail": "GET /api/categories/{id}",
					"update": "PUT /api/categories/{id}",
					"delete": "DELETE /api/categories/{id}",
				},
				"transactions": map[string]string{
					"checkout": "POST /api/checkout",
				},
				"reports": map[string]string{
					"today":      "GET /api/report/hari-ini",
					"date_range": "GET /api/report?start_date={YYYY-MM-DD}&end_date={YYYY-MM-DD}",
				},
			},
		})
	})

	// GET localhost:8080/api/produk/{id}
	// PUT localhost:8080/api/produk/{id}
	// DELETE localhost:8080/api/produk/{id}
	// POST localhost:8080/api/produk
	// GET localhost:8080/api/produk
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	http.HandleFunc("/api/produk", productHandler.HandleProducts)
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)

	// Category routes dengan layered architecture
	// GET/POST localhost:8080/api/categories
	// GET/PUT/DELETE localhost:8080/api/categories/{id}
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	http.HandleFunc("/api/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)

	// Transaction
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	http.HandleFunc("/api/checkout", transactionHandler.HandleCheckout) // POST
	http.HandleFunc("/api/report/hari-ini", transactionHandler.HandleSalesReport) // GET
	http.HandleFunc("/api/report", transactionHandler.HandleReportByDateRange) // GET with date range

	// Serve Swagger documentation
	http.HandleFunc("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, "docs/swagger.json")
	})

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "OK",
			"message": "Server is healthy",
		})
	})

	addr := fmt.Sprintf(":%s", config.Port)
	fmt.Println("Berhasil running server di port", config.Port)

	err = http.ListenAndServe(addr, nil)

	if err != nil{
		fmt.Println("Gagal memulai server:", err)
	}
}