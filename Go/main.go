package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Function to be called with a 1-second delay
func callfunctions() {
	// Your function logic here
	// For demonstration, we'll just print something
	log.Println("Function called")
	time.Sleep(1 * time.Second)
}

// Handler for the API endpoint
func callFunctionHandler(c *gin.Context) {
	// Get the number of times from the query parameter
	countStr := c.Query("count")
	count, err := strconv.Atoi(countStr)
	if err != nil || count < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid count parameter. Must be a positive integer.",
		})
		return
	}

	// Call the function the specified number of times
	for i := 0; i < count; i++ {
		callfunctions()
		functionCallCounter.Inc()
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Function called successfully",
		"count":   count,
	})
}

// Handler to serve the HTML content
func serveHtmlHandler(c *gin.Context) {
	log.Println("Serving home.html")
	c.File("home.html")
}

// Prometheus metrics
var (
	functionCallCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "my_function_calls_total",
			Help: "Total number of calls to myFunction",
		},
	)
)

func init() {
	// Register the function call counter with Prometheus
	prometheus.MustRegister(functionCallCounter)
}

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Define the endpoint to call the function
	r.GET("/callFunction", callFunctionHandler)

	// Define the endpoint to serve the HTML content
	r.GET("/", serveHtmlHandler)

	// Define the endpoint to expose Prometheus metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Start the server
	log.Println("Starting server on 0.0.0.0:8080")
	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
