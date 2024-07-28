package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Define a struct to represent the request payload
type OrderRequest struct {
	Items int `json:"items"`
}

// Define the response format
type PackResponse struct {
	Packs []Pack `json:"packs"`
}

type Pack struct {
	Size  int `json:"size"`
	Count int `json:"count"`
}

func hello(c echo.Context) error {
	fmt.Println("Hello, World!")
	return c.String(200, "Hello, World!")
}

func orderPacks(c echo.Context) error {
	var orderRequest OrderRequest

	items, err := strconv.Atoi(c.QueryParam("items"))

	if err != nil || items == 0 {
		return c.JSON(http.StatusBadRequest, "Items parameter is required and must be a non-zero integer")
	}

	orderRequest.Items = items

	// Now you can use the values from the struct
	fmt.Println("Receiveddddd", orderRequest.Items, "items")

	// Define your pack sizes and corresponding logic to calculate the number of complete packs
	packSizes := []int{5000, 2000, 1000, 500, 250}
	packs := []Pack{}
	remainingItems := orderRequest.Items

	for _, size := range packSizes {
		if remainingItems >= size {
			count := remainingItems / size
			if count > 0 {
				packs = append(packs, Pack{Size: size, Count: count})
				remainingItems -= count * size
			}
		}
	}

	if remainingItems > 0 {
		for i := len(packSizes) - 1; i >= 0; i-- {
			size := packSizes[i]
			if remainingItems <= size {
				packs = append(packs, Pack{Size: size, Count: 1})
				break
			}
		}
	}
	fmt.Printf("Order: %d, Packs: %v\n", orderRequest.Items, packs)

	return c.JSON(http.StatusOK, PackResponse{Packs: packs})
}

// Register endpoints
func routes(e *echo.Echo) {
	e.GET("/hello", hello)
	e.GET("/order_packs", orderPacks)
}

func main() {
	e := echo.New()
	routes(e)
	e.Start(":8080")
}
