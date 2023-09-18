package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Item struct {
	ShortDescription string `json:"short_description"`
	Price            string `json:"price"`
}

type Receipt struct {
	ID           string `json:"id"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchase_date"`
	PurchaseTime string `json:"purchase_time"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
	Points       int
}

// Map of our receipts
var receipts map[string]Receipt

func countAlphanumeric(s string) int {
	re := regexp.MustCompile("[[:alnum:]]")
	matches := re.FindAllString(s, -1)
	return len(matches)
}

func endsInDoubleZero(s string) bool {
	return strings.HasSuffix(s, ".00")
}

func isMultipleOfQuarter(s string) bool {
	// Convert the string to a float
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}

	// Divide by 0.25 and check if the remainder is 0
	return math.Mod(f, 0.25) == 0
}

func calculateDescLength(l []Item) int {
	total := 0
	for i := 0; i < len(l); i++ {
		b := l[i]
		if len(strings.TrimSpace(b.ShortDescription))%3 == 0 {
			val, err := strconv.ParseFloat(b.Price, 32)
			if err == nil {
				total += int(math.Ceil(val * 0.2))
			}
		}
	}
	return total
}

func isOddDay(date string) int {
	if len(date) < 10 {
		return 0
	}
	val, err := strconv.Atoi(string(date[9]))
	if err != nil {
		return 0
	}
	if val%2 == 1 {
		return 6
	}
	return 0
}

// Calculate the points for a receipt
func calculatePoints(receipt Receipt) int {
	var points = 0

	// 1 point for every alphanumeric character in the retailer name
	fmt.Println("len:", countAlphanumeric(receipt.Retailer))
	points += countAlphanumeric(receipt.Retailer)

	// 50 points if the total is a round dollar amount with no cents
	if endsInDoubleZero(receipt.Total) {
		fmt.Println("ends in double zero, +50")
		points += 50
	}

	// 25 points if the total is a multiple of 0.25
	if isMultipleOfQuarter(receipt.Total) {
		fmt.Println("is multiple of quarter, +25")
		points += 25
	}

	// 5 points for every two items on the receipt
	fmt.Println("5 for every two", int(math.Floor(float64(len(receipt.Items))/2.0))*5)
	points += int(math.Floor(float64(len(receipt.Items))/2.0)) * 5

	// If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	// The result is the number of points earned
	fmt.Println("desc length", calculateDescLength(receipt.Items))
	points += calculateDescLength(receipt.Items)

	// 6 points if the day in the purchase date is odd
	fmt.Println("odd day +6", isOddDay(receipt.PurchaseDate))
	points += isOddDay(receipt.PurchaseDate)

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	if receipt.PurchaseTime > "14:00" && receipt.PurchaseTime < "16:00" {
		fmt.Println("between 2 and 4 +10")
		points += 10
	}

	return points
}

// Create a new receipt when a POST request is made to /receipts/process
func generateReceipt(c *gin.Context) {
	var newReceipt Receipt
	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}

	newReceipt.Points = calculatePoints(newReceipt)
	fmt.Println(newReceipt.Points)

	// Generate a new UUID for the receipt
	id := (uuid.New().String())
	receipts[id] = newReceipt

	// Return the ID of the receipt
	c.IndentedJSON(http.StatusCreated, gin.H{"id": id})
}

// Get a receipt when a GET request is made to /receipts/:id/points
func getReceipt(c *gin.Context) {
	id := c.Param("id")
	for key, value := range receipts {
		if key == id {
			c.IndentedJSON(http.StatusOK, gin.H{"points": value.Points})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "receipt not found"})
}

func main() {
	receipts = make(map[string]Receipt)
	router := gin.Default()
	router.POST("/receipts/process", generateReceipt)
	router.GET("/receipts/:id/points", getReceipt)
	log.Fatal(router.Run(":8080"))
}
