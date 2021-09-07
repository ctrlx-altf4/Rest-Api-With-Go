package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Order struct{
	OrderID  string `json:"orderId"`
	CustomerName string `json:"customerName`
	OrderedAt string `json:"orderedAt"`
	Items string `json:"items"`
}

type Items struct{
	ItemId string `json:"itemId"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
}

var orders []Order
var prevOrderId =0

func main(){
	router := mux.NewRouter()

	//CREATE
	router.HandleFunc("/orders",createOrder).Methods("POST")

	//READ
	router.HandleFunc("/orders/{orderId}",getOrder).Methods("GET")

	//Read all
	router.HandleFunc("/orders",getOrders).Methods("GET")

	//update
	router.HandleFunc("/orders/{orderId}",updateOrder).Methods("PUT")

	//Delete
	router.HandleFunc("/orders/{orderId}",deleteOrder).Methods("DELETE")

	//Swagger
	// router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080",router))
}

func createOrder(w http.ResponseWriter,r *http.Request	){
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderId++
	order.OrderID = strconv.Itoa(prevOrderId)
	orders = append(orders,order)
	w.Header().Set("Content-type","application/json")

	json.NewEncoder(w).Encode(order)
}

func getOrders(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(orders)
}

func getOrder(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for _, order := range orders{
		if order.OrderID == inputOrderID{
			json.NewEncoder(w).Encode(order)
			return
		}
	}
}

func updateOrder(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	
	for i, order := range orders{
		if inputOrderID == order.OrderID{
			orders = append(orders[:i],orders[i+1:]...)
			var updatedOrder Order
			json.NewDecoder(r.Body).Decode(&updatedOrder)

			orders = append(orders,updatedOrder)
			json.NewEncoder(w).Encode(updatedOrder)
			return

		}
	}
}

func deleteOrder(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]

	for i, order := range orders{
		if inputOrderID == order.OrderID{
			orders = append(orders[:i],orders[i+1:]... )
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}