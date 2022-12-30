package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/bontusss/go-ecom/database"
	"github.com/bontusss/go-ecom/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "users")
var ProductCollection *mongo.Collection = database.Productdata(database.Client, "products")
var Validate = Validator.New()

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// Validate.Struct
		count, err := UserCollection.CountDocuments()

	}
}

func Login(c *gin.Context) {}

func ProductViewerAdmin(c *gin.Context) {}

func Searchproduct(c *gin.Context) {}

func SearchproductByQuery(c *gin.Context) {}

func HashPassword(password string) string {
	return ""
}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	return true, ""
}
