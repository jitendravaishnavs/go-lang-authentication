package controllers

import (
	"authentication/config"
	"authentication/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gingonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

var validate = validator.New()
var userCollection = config.OpenCollection("users")

func Signup() gin.HandleFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

		defer cancel()

		var user models.User

		//get User input
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		//validate user input
		if validationErr := validate.Struct(user); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{
			"$or": []bson.M{
				{"email": user.email},
				{"phone": user.phone},
			},
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email or phone number already exist"})
			return
		}

		user.Password = HashPassword(user.Password)
		user.Created_at = time.Now()
		user.Updated_at = time.Now()

	}

}

func Login() {

}

func GetUsers() {}

func GetUser() {

}
