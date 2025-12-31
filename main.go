package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
	AIEnabled   bool    `json:"aiEnabled,omitempty"`
}

var productList []Product

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT,DELETE,PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Zubaer")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data) // fix here
}

// Root handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, `
		<h1>Welcome to Product API</h1>
		<p>Available endpoints:</p>
		<ul>
			<li><a href="/products">GET /products</a> - List all products</li>
			<li>POST /create-products - Add new product</li>
		</ul>
		<p>Frontend should be running on <a href="http://localhost:3000">http://localhost:3000</a></p>
	`)
}

func getProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		return
	}
	
	sendData(w, productList, http.StatusOK) // use send data here
} // 

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	sendData(w, newProduct, http.StatusCreated)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)  
	mux.HandleFunc("/products", enableCORS(getProducts))
	mux.HandleFunc("/create-products", enableCORS(createProduct))

	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("Visit: http://localhost:8080/products to see products")
	http.ListenAndServe(":8080", mux)
}

func init() {
	productList = append(productList,
		Product{
			ID:          1,
			Title:       "Orange",
			Description: "Orange is red, I love orange",
			Price:       100,
			ImgURL:      "https://imgs.search.brave.com/1xgw_-dkLyHeEr3d8AnzeEgEY1gT_0EOKXOodKdzruA/rs:fit:500:0:1/g:ce/aHR0cHM6Ly9zdGF0/aWMudmVjdGVlenkuY29tL3N5c3RlbS9y/ZXNvdXJjZXMvdGh1/bWJuYWlscy8wNjYv/MjQ2Lzc2NS9zbWFs/bC9vcmFuZ2UtZnJ1/aXQtc2xpY2UtaWxs/dXN0cmF0aW9uLXBu/Zy5wbmc",
		},
		Product{
			ID:          2,
			Title:       "Apple",
			Description: "Apple is red, I love Apple",
			Price:       100,
			ImgURL:      "https://cdn.pixabay.com/photo/2023/01/16/10/28/apple-7722078_640.jpg",
		},
		Product{
			ID:          3,
			Title:       "Banana",
			Description: "Banana is yellow, I love Banana",
			Price:       100,
			ImgURL:      "https://render.fineartamerica.com/images/images-profile-flow/400/images-medium-large-5/bananas-splashed-into-water-henrik-sorensen.jpg",
		},
	)
}
