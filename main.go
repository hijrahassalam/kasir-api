package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Produk struct {
	ID int `json:"id"`
	Nama string `json:"nama"`
	Harga int `json:"harga"`
	Stok int `json:"stok"`
}

var produk = []Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 20},
	{ID: 2, Nama: "Aquaviva 2000ml", Harga: 4000, Stok: 15},
	{ID: 3, Nama: "Chitato", Harga: 8000, Stok: 10},
}

func getProdukByID(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			// ini set header agar response nya json
			w.Header().Set("Content-Type", "application/json")
			// ini set status code nya jadi 400 Bad Request
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Invalid product ID",
			})
			return
		}

		for _, p := range produk {
			if p.ID == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(p)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Product not found",
		})
}

func updateProdukByID(w http.ResponseWriter, r *http.Request) {
	// ambil ID dari request URL
	// ganti jadi int
	// loop data produk, cari yang sesuai ID
	// ganti sesuai dengan data dari request body
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid product ID",
		})
		return
	}

	var updatedProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updatedProduk)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid request payload",
		})
		return
	}

	for i := range produk {
		if produk[i].ID == id {
			produk[i] = updatedProduk
			produk[i].ID = id
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk[i])
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func deleteProdukByID(w http.ResponseWriter, r *http.Request) {
	// get id 
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	// ganti id jadi int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid product ID",
		})
		return
	}

	// loop cari id di data produk
	for i, p := range produk {
		if p.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah data yang dihapus
			produk = append(produk[:i], produk[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Product deleted successfully",
			})
			return 
		}
	}


	http.Error(w, "Product not found", http.StatusNotFound)
}
	


func main(){
	// GET localhost:8080/api/produk/{id}
	// PUT localhost:8080/api/produk/{id}
	// DELETE localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/",	 func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukByID(w, r)
		} else if r.Method == "PUT" {
			updateProdukByID(w, r)
		} else if r.Method == "DELETE" {
			deleteProdukByID(w, r)
		}

	})	

	// POST localhost:8080/api/produk
	// GET localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid request payload", http.StatusBadRequest)
				return
			}

			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(produkBaru)
 		}
	})

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "OK",
			"message": "Server is healthy",
		})
	})

	fmt.Println("Berhasil running server di Localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil{
		fmt.Println("Gagal memulai server:", err)
	}
}