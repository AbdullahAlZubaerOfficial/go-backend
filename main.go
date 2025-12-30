package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Zubaer. I'm Web Developer.")
}

type Product struct {
	ID          int       `json:"id"`
	Title       string     `json:"title"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	ImgURL      string       `json:"imageUrl"`
}

var productList []Product


// get products
func getProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Content-Type", "application/json")


	if r.Method != "GET" {
		http.Error(w, "Please give me GET request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productList)
}


// create products
func createProduct(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Content-Type", "application/json")


	if r.Method != "POST" {
		http.Error(w, "Please give me POST request", http.StatusBadRequest)
		return
	}

	// description
	// imageUrl
	// price
	// title

	// r.Body => description, imageUrl, price, title => Product er ekta instance => productList => append

	/*

	1. take body information (description, imageUrl, price , title) from r.Body
	2. create an instance using product struct with the body information 
	3. append the instance into productList

	*/

	var newProduct Product 

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newProduct)
	err := decoder.Decode(&newProduct)
	
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json",400)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getProducts)

	fmt.Println("Server is running on :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}


func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red, I love orange",
		Price:       100,
		ImgURL:      "https://imgs.search.brave.com/1xgw_-dkLyHeEr3d8AnzeEgEY1gT_0EOKXOodKdzruA/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly9zdGF0/aWMudmVjdGVlenku/Y29tL3N5c3RlbS9y/ZXNvdXJjZXMvdGh1/bWJuYWlscy8wNjYv/MjQ2Lzc2NS9zbWFs/bC9vcmFuZ2UtZnJ1/aXQtc2xpY2UtaWxs/dXN0cmF0aW9uLXBu/Zy5wbmc",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is red, I love Apple",
		Price:       100,
		ImgURL:      "https://imgs.search.brave.com/W2GSmeX06IzepHl2uEEBKvRBwl9xr0rMfesJZ_XWgUw/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly9jZG4u/cGl4YWJheS5jb20v/cGhvdG8vMjAyMy8w/MS8xNi8xMC8yOC9h/cHBsZS03NzIyMDc4/XzY0MC5qcGc",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is yellow, I love Banana",
		Price:       100,
		ImgURL:      "https://imgs.search.brave.com/8NqROa-wFwylIqSirGnNdTt47c4pA9ZBPBnVDItuaDA/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly9yZW5k/ZXIuZmluZWFydGFt/ZXJpY2EuY29tL2lt/YWdlcy9pbWFnZXMt/cHJvZmlsZS1mbG93/LzQwMC9pbWFnZXMt/bWVkaXVtLWxhcmdl/LTUvYmFuYW5hcy1z/cGxhc2hlZC1pbnRv/LXdhdGVyLWhlbnJp/ay1zb3JlbnNlbi5q/cGc",
	}



	productList = append(productList, prd1, prd2, prd3)
}


/*

          [] => list
		  {} => object 

		  JSON => javascript object notation

*/



