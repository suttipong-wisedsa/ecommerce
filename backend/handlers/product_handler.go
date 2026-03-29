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
	keyword := c.Query("q")
	if keyword != "" {
		config.DB.
			Where("name ILIKE ?", "%"+keyword+"%").
			Find(&products)
	} else {
		config.DB.Find(&products)
	}
	c.JSON(http.StatusOK, products)
}

var cart []models.CartItem

func GetCart(c *gin.Context) {
	c.JSON(http.StatusOK, cart)
}

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		authHeader := c.GetHeader("Authorization")

// 		if authHeader == "" {
// 			c.AbortWithStatusJSON(401, gin.H{"error": "no token"})
// 			return
// 		}

// 		// 🔥 ตัด Bearer ออก
// 		tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)

// 		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 			return jwtKey, nil
// 		})

// 		if err != nil || !token.Valid {
// 			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
// 			return
// 		}

// 		claims := token.Claims.(jwt.MapClaims)

// 		// 🔥 set user_id
// 		c.Set("user_id", uint(claims["user_id"].(float64)))

// 		c.Next()
// 	}
// }

// func AddToCart(c *gin.Context) {
// 	userID := c.MustGet("user_id").(uint)
// 	var body struct {
// 		ProductID uint `json:"product_id"`
// 		Qty       int  `json:"qty"`
// 	}

// 	c.BindJSON(&body)

// 	var cart models.Cart

// 	// 1. หา cart ของ user
// 	if err := config.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
// 		// ถ้าไม่มี → create ใหม่
// 		cart = models.Cart{UserID: 1}
// 		config.DB.Create(&cart)
// 	}

// 	var item models.CartItem

// 	// 2. check ว่ามีสินค้านี้ใน cart ไหม
// 	err := config.DB.
// 		Where("cart_id = ? AND product_id = ?", cart.ID, body.ProductID).
// 		First(&item).Error

// 	if err == nil {
// 		// มีแล้ว → เพิ่มจำนวน
// 		item.Quantity += body.Qty
// 		config.DB.Save(&item)
// 	} else {
// 		// ยังไม่มี → create ใหม่
// 		item = models.CartItem{
// 			CartID:    cart.ID,
// 			ProductID: body.ProductID,
// 			Quantity:  body.Qty,
// 		}
// 		config.DB.Create(&item)
// 	}

// 	c.JSON(200, gin.H{"message": "added to cart"})
// }

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
