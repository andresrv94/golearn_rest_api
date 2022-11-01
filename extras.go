package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
