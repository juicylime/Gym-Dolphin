package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

// Serve the react build
func serveReactApp(e *echo.Echo) {
	// Serve static files
	e.Static("/static", "gym-dolphin/build/static")
	e.File("/", "gym-dolphin/build/index.html")
	e.File("/favicon.ico", "gym-dolphin/build/favicon.ico")
	e.File("/manifest.json", "gym-dolphin/build/manifest.json")
	e.File("/logo192.png", "gym-dolphin/build/logo192.png")
	e.File("/logo512.png", "gym-dolphin/build/logo512.png")
	e.File("/robots.txt", "gym-dolphin/build/robots.txt")
	e.File("/white_shirt.jpg", "gym-dolphin/build/white_shirt.jpg")

	// Serve index.html for any other routes except for known static files
	e.File("/*", "gym-dolphin/build/index.html")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "https://gym-dolphin.onrender.com"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	routes(e)
	serveReactApp(e)

	e.Logger.Fatal(e.Start(":8080"))
}
