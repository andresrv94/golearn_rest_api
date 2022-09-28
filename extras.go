package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type btcResponse struct {
	Status map[string]interface{} `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getBtcPrice(c *gin.Context) {
	response, err := http.Get("https://data.messari.io/api/v1/assets/btc/metrics/market-data")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(responseData))

	data_obj := btcResponse{}
	jsonErr := json.Unmarshal(responseData, &data_obj)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	// fmt.Println("Printing unmarshalled values:")
	// fmt.Println("Data: ", data_obj.Data)
	// fmt.Println("Status: ", data_obj.Status)
	marketData := data_obj.Data["market_data"].(map[string]interface{})

	c.IndentedJSON(http.StatusOK, marketData["price_usd"])
}

func calculator2(c *gin.Context) {
	sn1 := c.PostForm("n1")
	sn2 := c.PostForm("n2")
	op := c.PostForm("op")

	n1, err := strconv.Atoi(sn1)
	if err != nil {
		// fmt.Println("Error during conversion")
		// return
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"status": "Error",
			"result": "n1 is not a valid integer",
		})
		fmt.Println("Error during n1 conversion")
		return
	}
	n2, err := strconv.Atoi(sn2)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"status": "Error",
			"result": "n2 is not a valid integer",
		})
		fmt.Println("Error during n2 conversion")
		return
	}
	if op == "+" {
		result := sumNumbers(n1, n2)
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "OK",
			"result": result,
		})
	} else if op == "-" {
		result := substractNumbers(n1, n2)
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "OK",
			"result": result,
		})
	} else if op == "*" {
		result := multiplyNumbers(n1, n2)
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "OK",
			"result": result,
		})
	} else if op == "/" {
		result := divideNumbers(n1, n2)
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "OK",
			"result": result,
		})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "Error",
			"result": "Operation not recognized",
		})
	}

}

func sumNumbers(x int, y int) int {
	return x + y
}

func substractNumbers(x int, y int) int {
	return x - y
}

func multiplyNumbers(x int, y int) int {
	return x * y
}

func divideNumbers(x int, y int) int {
	return x / y
}
