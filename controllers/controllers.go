package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/bontusss/go-ecom/database"
	"github.com/bontusss/go-ecom/models"
	"github.com/bontusss/go-ecom/tokens"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var ProductCollection *mongo.Collection = database.Productdata(database.Client, "Products")
var Validate = validator.New()

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
		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validationErr,
			})
		}
		count, err := UserCollection.CountDocuments(ctx, bson.M{
			"email": user.Email,
		})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "User already exist",
			})
		}

		count, err = UserCollection.CountDocuments(ctx, bson.M{
			"phone": user.Phone,
		})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Phone number already exist",
			})
			return
		}

		password := HashPassword(user.Password)
		user.Password = password
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		token, refreshToken, _ := tokens.TokenGenerator(user.Email, user.First_Name, user.Last_Name, string(user.User_ID))
		user.Token = token
		user.Refresh_Token = refreshToken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Detail = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)

		_, insertErr := UserCollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not create User",
			})
			return
		}
		defer cancel()

		c.JSON(http.StatusCreated, gin.H{
			"success": "User created successfully",
		})
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
