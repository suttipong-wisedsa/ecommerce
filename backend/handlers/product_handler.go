package handlers

import (
	"net/http"

	"myapi/config"
	"myapi/models"

	"github.com/gin-gonic/gin"
)

// var products = []models.Product{
// 	{ID: 1, Name: "iPhone", Price: 35000, Photo: "https://media.studio7thailand.com/225686/MacBook_13-in_A18_Pro_Citrus_PDP_Image_Position_1__TH-TH-square_medium.jpg"},
// 	{ID: 2, Name: "MacBook", Price: 55000, Photo: "https://media.studio7thailand.com/225686/MacBook_13-in_A18_Pro_Citrus_PDP_Image_Position_1__TH-TH.jpg"},
// 	{ID: 3, Name: "iPhone", Price: 35000, Photo: "https://media.studio7thailand.com/225686/MacBook_13-in_A18_Pro_Citrus_PDP_Image_Position_1__TH-TH-square_medium.jpg"},
// 	{ID: 4, Name: "MacBook", Price: 55000, Photo: "https://media.studio7thailand.com/225686/MacBook_13-in_A18_Pro_Citrus_PDP_Image_Position_1__TH-TH.jpg"},
// 	{ID: 5, Name: "iPhone", Price: 35000, Photo: "https://media.studio7thailand.com/225686/MacBook_13-in_A18_Pro_Citrus_PDP_Image_Position_1__TH-TH-square_medium.jpg"},
// 	{ID: 6, Name: "MacBook", Price: 55000, Photo: "https://media.studio7thailand.com/225686/MacBook_13-in_A18_Pro_Citrus_PDP_Image_Position_1__TH-TH.jpg"},
// 	{ID: 7, Name: "iPhone", Price: 35000, Photo: "https://media.studio7thailand.com/225686/MacBook_13-in_A18_Pro_Citrus_PDP_Image_Position_1__TH-TH-square_medium.jpg"},
// 	{ID: 8, Name: "MacBook", Price: 55000, Photo: "https://media.studio7thailand.com/225686/MacBook_13-in_A18_Pro_Citrus_PDP_Image_Position_1__TH-TH.jpg"},
// }

func GetProduct(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

var cart []models.CartItem

func GetCart(c *gin.Context) {
	c.JSON(http.StatusOK, cart)
}

func AddToCart(c *gin.Context) {
	var item models.CartItem

	// รับ JSON
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 👉 logic: ถ้ามีสินค้าแล้ว → เพิ่ม quantity
	// found := false
	// for i, v := range cart {
	// 	if v.ProductID == item.ProductID {
	// 		cart[i].Quantity += item.Quantity
	// 		found = true
	// 		break
	// 	}
	// }

	// if !found {
	// 	cart = append(cart, item)
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "added to cart",
		"cart":    cart,
	})
}

// GET /users/:id
// func GetUserByID(c *gin.Context) {
// 	idParam := c.Param("id")
// 	id, _ := strconv.Atoi(idParam)

// 	for _, u := range users {
// 		if u.ID == id {
// 			c.JSON(http.StatusOK, u)
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// }

// POST /users
// func CreateUser(c *gin.Context) {
// 	var newUser models.User

// 	if err := c.ShouldBindJSON(&newUser); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	newUser.ID = len(users) + 1
// 	users = append(users, newUser)

// 	c.JSON(http.StatusCreated, newUser)
// }

// // PUT /users/:id
// func UpdateUser(c *gin.Context) {
// 	idParam := c.Param("id")
// 	id, _ := strconv.Atoi(idParam)

// 	var updated models.User
// 	if err := c.ShouldBindJSON(&updated); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	for i, u := range users {
// 		if u.ID == id {
// 			users[i].Name = updated.Name
// 			c.JSON(http.StatusOK, users[i])
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// }

// // DELETE /users/:id
// func DeleteUser(c *gin.Context) {
// 	idParam := c.Param("id")
// 	id, _ := strconv.Atoi(idParam)

// 	for i, u := range users {
// 		if u.ID == id {
// 			users = append(users[:i], users[i+1:]...)
// 			c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// }
