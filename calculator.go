package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
